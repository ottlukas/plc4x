/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/
package org.apache.plc4x.java.ads.protocol;

import org.apache.plc4x.java.ads.readwrite.types.ReturnCode;
import org.apache.plc4x.java.ads.configuration.AdsConfiguration;
import org.apache.plc4x.java.ads.field.AdsField;
import org.apache.plc4x.java.ads.field.DirectAdsField;
import org.apache.plc4x.java.ads.field.SymbolicAdsField;
import org.apache.plc4x.java.ads.readwrite.*;
import org.apache.plc4x.java.ads.readwrite.io.DataItemIO;
import org.apache.plc4x.java.ads.readwrite.types.CommandId;
import org.apache.plc4x.java.ads.readwrite.types.ReservedIndexGroups;
import org.apache.plc4x.java.api.exceptions.PlcException;
import org.apache.plc4x.java.api.exceptions.PlcRuntimeException;
import org.apache.plc4x.java.api.messages.PlcReadRequest;
import org.apache.plc4x.java.api.messages.PlcReadResponse;
import org.apache.plc4x.java.api.messages.PlcWriteRequest;
import org.apache.plc4x.java.api.messages.PlcWriteResponse;
import org.apache.plc4x.java.api.model.PlcField;
import org.apache.plc4x.java.api.types.PlcResponseCode;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.ConversationContext;
import org.apache.plc4x.java.spi.Plc4xProtocolBase;
import org.apache.plc4x.java.spi.configuration.HasConfiguration;
import org.apache.plc4x.java.spi.generation.ParseException;
import org.apache.plc4x.java.spi.generation.ReadBuffer;
import org.apache.plc4x.java.spi.generation.WriteBuffer;
import org.apache.plc4x.java.spi.messages.DefaultPlcReadResponse;
import org.apache.plc4x.java.spi.messages.DefaultPlcWriteResponse;
import org.apache.plc4x.java.spi.messages.InternalPlcReadRequest;
import org.apache.plc4x.java.spi.messages.InternalPlcWriteRequest;
import org.apache.plc4x.java.spi.messages.utils.ResponseItem;
import org.apache.plc4x.java.spi.transaction.RequestTransactionManager;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.time.Duration;
import java.util.*;
import java.util.concurrent.*;
import java.util.concurrent.atomic.AtomicLong;
import java.util.stream.Collectors;
import java.util.stream.IntStream;

public class AdsProtocolLogic extends Plc4xProtocolBase<AmsTCPPacket> implements HasConfiguration<AdsConfiguration> {

    private static final Logger LOGGER = LoggerFactory.getLogger(AdsProtocolLogic.class);

    private AdsConfiguration configuration;
    public static final State DEFAULT_COMMAND_STATE = new State(
        false, false, false, false, false, true, false, false, false);

    private ConversationContext<AmsTCPPacket> adsDriverContext;
    private final AtomicLong invokeIdGenerator = new AtomicLong(1);
    private RequestTransactionManager tm;

    private ConcurrentHashMap<SymbolicAdsField, DirectAdsField> symbolicFieldMapping;
    private ConcurrentHashMap<SymbolicAdsField, CompletableFuture<Void>> pendingResolutionRequests;

    public AdsProtocolLogic() {
        symbolicFieldMapping = new ConcurrentHashMap<>();
        pendingResolutionRequests = new ConcurrentHashMap<>();

        // Initialize Transaction Manager.
        // Until the number of concurrent requests is successfully negotiated we set it to a
        // maximum of only one request being able to be sent at a time. During the login process
        // No concurrent requests can be sent anyway. It will be updated when receiving the
        // S7ParameterSetupCommunication response.
        this.tm = new RequestTransactionManager(1);
    }

    @Override
    public void setConfiguration(AdsConfiguration configuration) {
        this.configuration = configuration;
    }

    @Override
    public void setContext(ConversationContext<AmsTCPPacket> context) {
        super.setContext(context);
        adsDriverContext = context;
    }

    @Override
    public void close(ConversationContext<AmsTCPPacket> context) {

    }

    @Override
    public void onConnect(ConversationContext<AmsTCPPacket> context) {
        // AMS/ADS doesn't know a concept of a connect.
        context.fireConnected();
    }

    @Override
    public void onDisconnect(ConversationContext<AmsTCPPacket> context) {
        super.onDisconnect(context);
        // TODO: Here we have to clean up all of the handles this connection acquired.
    }

    @Override
    public CompletableFuture<PlcReadResponse> read(PlcReadRequest readRequest) {
        // Get all ADS addresses in their resolved state.
        final CompletableFuture<List<DirectAdsField>> directAdsFieldsFuture =
            getDirectAddresses(readRequest.getFields());

        // If all addresses were already resolved we can send the request immediately.
        if(directAdsFieldsFuture.isDone()) {
            final List<DirectAdsField> fields = directAdsFieldsFuture.getNow(null);
            if(fields != null) {
                return executeRead(readRequest, fields);
            } else {
                final CompletableFuture<PlcReadResponse> errorFuture = new CompletableFuture<>();
                errorFuture.completeExceptionally(new PlcException("Error"));
                return errorFuture;
            }
        }
        // If there are still symbolic addresses that have to be resolved, send the
        // request as soon as the resolution is done.
        // In order to instantly be able to return a future, for the final result we have to
        // create a new one which is then completed later on. Unfortunately as soon as the
        // directAdsFieldsFuture is completed we still don't have the end result, but we can
        // now actually send the delayed read request ... as soon as that future completes
        // we can complete the initial one.
        else {
            CompletableFuture<PlcReadResponse> delayedRead = new CompletableFuture<>();
            directAdsFieldsFuture.handle((directAdsFields, throwable) -> {
                if(directAdsFields != null) {
                    final CompletableFuture<PlcReadResponse> delayedResponse =
                        executeRead(readRequest, directAdsFields);
                    delayedResponse.handle((plcReadResponse, throwable1) -> {
                        if (plcReadResponse != null) {
                            delayedRead.complete(plcReadResponse);
                        } else {
                            delayedRead.completeExceptionally(throwable1);
                        }
                        return this;
                    });
                } else {
                    delayedRead.completeExceptionally(throwable);
                }
                return this;
            });
            return delayedRead;
        }
    }

    protected CompletableFuture<PlcReadResponse> executeRead(PlcReadRequest readRequest,
                                                             List<DirectAdsField> directAdsFields) {
        // Depending on the number of fields, use a single item request or a sum-request
        if (directAdsFields.size() == 1) {
            // Do a normal (single item) ADS Read Request
            return singleRead(readRequest, directAdsFields.get(0));
        } else {
            // TODO: Check if the version of the remote station is at least TwinCAT v2.11 Build >= 1550 otherwise split up into single item requests.
            // Do a ADS-Sum Read Request.
            return multiRead(readRequest, directAdsFields);
        }
    }

    protected CompletableFuture<PlcReadResponse> singleRead(PlcReadRequest readRequest, DirectAdsField directAdsField) {
        CompletableFuture<PlcReadResponse> future = new CompletableFuture<>();

        int size = directAdsField.getAdsDataType().getNumBytes() * directAdsField.getNumberOfElements();
        AdsData adsData = new AdsReadRequest(directAdsField.getIndexGroup(), directAdsField.getIndexOffset(), size);
        AmsPacket amsPacket = new AmsPacket(configuration.getTargetAmsNetId(), configuration.getTargetAmsPort(),
            configuration.getSourceAmsNetId(), configuration.getSourceAmsPort(),
            CommandId.ADS_READ, DEFAULT_COMMAND_STATE, 0, getInvokeId(), adsData);
        AmsTCPPacket amsTCPPacket = new AmsTCPPacket(amsPacket);

        // Start a new request-transaction (Is ended in the response-handler)
        RequestTransactionManager.RequestTransaction transaction = tm.startRequest();
        transaction.submit(() -> context.sendRequest(amsTCPPacket)
            .expectResponse(AmsTCPPacket.class, Duration.ofMillis(configuration.getTimeoutRequest()))
            .onTimeout(future::completeExceptionally)
            .onError((p, e) -> future.completeExceptionally(e))
            .check(responseAmsPacket -> responseAmsPacket.getUserdata().getInvokeId() == amsPacket.getInvokeId())
            .unwrap(response -> (AdsReadResponse) response.getUserdata().getData())
            .handle(responseAdsData -> {
                if(responseAdsData.getResult() == ReturnCode.OK) {
                    final PlcReadResponse plcReadResponse = convertToPlc4xReadResponse(readRequest, responseAdsData);
                    // Convert the response from the PLC into a PLC4X Response ...
                    future.complete(plcReadResponse);
                } else {
                    // TODO: Implement this correctly.
                    future.completeExceptionally(new PlcException("Error"));
                }
                // Finish the request-transaction.
                transaction.endRequest();
            }));
        return future;
    }

    protected CompletableFuture<PlcReadResponse> multiRead(PlcReadRequest readRequest, List<DirectAdsField> directAdsFields) {
        CompletableFuture<PlcReadResponse> future = new CompletableFuture<>();

        // Calculate the size of all fields together.
        // Calculate the expected size of the response data.
        long expectedResponseDataSize = directAdsFields.stream().mapToLong(
            field -> (field.getAdsDataType().getNumBytes() * field.getNumberOfElements()) + 4).sum();

        // With multi-requests, the index-group is fixed and the index offset indicates the number of elements.
        AdsData adsData = new AdsReadWriteRequest(
            ReservedIndexGroups.ADSIGRP_MULTIPLE_READ.getValue(), directAdsFields.size(), expectedResponseDataSize,
            directAdsFields.stream().map(directAdsField -> new AdsMultiRequestItemRead(
                directAdsField.getIndexGroup(), directAdsField.getIndexOffset(),
                (directAdsField.getAdsDataType().getNumBytes() * directAdsField.getNumberOfElements())))
                .toArray(AdsMultiRequestItem[]::new), null);

        AmsPacket amsPacket = new AmsPacket(configuration.getTargetAmsNetId(), configuration.getTargetAmsPort(),
            configuration.getSourceAmsNetId(), configuration.getSourceAmsPort(),
            CommandId.ADS_READ_WRITE, DEFAULT_COMMAND_STATE, 0, getInvokeId(), adsData);
        AmsTCPPacket amsTCPPacket = new AmsTCPPacket(amsPacket);

        // Start a new request-transaction (Is ended in the response-handler)
        RequestTransactionManager.RequestTransaction transaction = tm.startRequest();
        transaction.submit(() -> context.sendRequest(amsTCPPacket)
            .expectResponse(AmsTCPPacket.class, Duration.ofMillis(configuration.getTimeoutRequest()))
            .onTimeout(future::completeExceptionally)
            .onError((p, e) -> future.completeExceptionally(e))
            .check(responseAmsPacket -> responseAmsPacket.getUserdata().getInvokeId() == amsPacket.getInvokeId())
            .unwrap(response -> (AdsReadWriteResponse) response.getUserdata().getData())
            .handle(responseAdsData -> {
                if(responseAdsData.getResult() == ReturnCode.OK) {
                    final PlcReadResponse plcReadResponse = convertToPlc4xReadResponse(readRequest, responseAdsData);
                    // Convert the response from the PLC into a PLC4X Response ...
                    future.complete(plcReadResponse);
                } else {
                    // TODO: Implement this correctly.
                    future.completeExceptionally(new PlcException("Error"));
                }
                // Finish the request-transaction.
                transaction.endRequest();
            }));
        return future;
    }

    protected PlcReadResponse convertToPlc4xReadResponse(PlcReadRequest readRequest, AdsData adsData) {
        ReadBuffer readBuffer = null;
        Map<String, PlcResponseCode> responseCodes = new HashMap<>();
        if (adsData instanceof AdsReadResponse) {
            AdsReadResponse adsReadResponse = (AdsReadResponse) adsData;
            readBuffer = new ReadBuffer(adsReadResponse.getData(), true);
            responseCodes.put(readRequest.getFieldNames().stream().findFirst().orElse(""),
                parsePlcResponseCode(adsReadResponse.getResult()));
        } else if (adsData instanceof AdsReadWriteResponse) {
            AdsReadWriteResponse adsReadWriteResponse = (AdsReadWriteResponse) adsData;
            readBuffer = new ReadBuffer(adsReadWriteResponse.getData(), true);
            // When parsing a multi-item response, the error codes of each items come
            // in sequence and then come the values.
            for (String fieldName : readRequest.getFieldNames()) {
                try {
                    final ReturnCode result = ReturnCode.valueOf(readBuffer.readUnsignedLong(32));
                    responseCodes.put(fieldName, parsePlcResponseCode(result));
                } catch (ParseException e) {
                    responseCodes.put(fieldName, PlcResponseCode.INTERNAL_ERROR);
                }
            }
        }
        if(readBuffer != null) {
            Map<String, ResponseItem<PlcValue>> values = new HashMap<>();
            for (String fieldName : readRequest.getFieldNames()) {
                AdsField field = (AdsField) readRequest.getField(fieldName);
                // If the response-code was anything but OK, we don't need to parse the payload.
                if(responseCodes.get(fieldName) != PlcResponseCode.OK) {
                    values.put(fieldName, new ResponseItem<>(responseCodes.get(fieldName), null));
                }
                // If the response-code was ok, parse the data returned.
                else {
                    values.put(fieldName, parsePlcValue(field, readBuffer));
                }
            }
            return new DefaultPlcReadResponse((InternalPlcReadRequest) readRequest, values);
        }
        return null;
    }

    private PlcResponseCode parsePlcResponseCode(ReturnCode adsResult) {
        if (adsResult == ReturnCode.OK) {
            return PlcResponseCode.OK;
        } else {
            // TODO: Implement this a little more ...
            return PlcResponseCode.INTERNAL_ERROR;
        }
    }

    private ResponseItem<PlcValue> parsePlcValue(AdsField field, ReadBuffer readBuffer) {
        try {
            if (field.getNumberOfElements() == 1) {
                return new ResponseItem<>(PlcResponseCode.OK,
                    DataItemIO.staticParse(readBuffer, field.getAdsDataType()));
            } else {
                // Fetch all
                final PlcValue[] resultItems = IntStream.range(0, field.getNumberOfElements()).mapToObj(i -> {
                    try {
                        return DataItemIO.staticParse(readBuffer, field.getAdsDataType());
                    } catch (ParseException e) {
                        LOGGER.warn("Error parsing field item of type: '{}' (at position {}})", field.getAdsDataType(), i, e);
                    }
                    return null;
                }).toArray(PlcValue[]::new);
                return new ResponseItem<>(PlcResponseCode.OK, PlcValues.of(resultItems));
            }
        } catch (Exception e) {
            LOGGER.warn(String.format("Error parsing field item of type: '%s'", field.getAdsDataType()), e);
            return new ResponseItem<>(PlcResponseCode.INTERNAL_ERROR, null);
        }
    }

    @Override
    public CompletableFuture<PlcWriteResponse> write(PlcWriteRequest writeRequest) {
        // Get all ADS addresses in their resolved state.
        final CompletableFuture<List<DirectAdsField>> directAdsFieldsFuture =
            getDirectAddresses(writeRequest.getFields());

        // If all addresses were already resolved we can send the request immediately.
        if(directAdsFieldsFuture.isDone()) {
            final List<DirectAdsField> fields = directAdsFieldsFuture.getNow(null);
            if(fields != null) {
                return executeWrite((InternalPlcWriteRequest) writeRequest, fields);
            } else {
                final CompletableFuture<PlcWriteResponse> errorFuture = new CompletableFuture<>();
                errorFuture.completeExceptionally(new PlcException("Error"));
                return errorFuture;
            }
        }
        // If there are still symbolic addresses that have to be resolved, send the
        // request as soon as the resolution is done.
        // In order to instantly be able to return a future, for the final result we have to
        // create a new one which is then completed later on. Unfortunately as soon as the
        // directAdsFieldsFuture is completed we still don't have the end result, but we can
        // now actually send the delayed read request ... as soon as that future completes
        // we can complete the initial one.
        else {
            CompletableFuture<PlcWriteResponse> delayedWrite = new CompletableFuture<>();
            directAdsFieldsFuture.handle((directAdsFields, throwable) -> {
                if(directAdsFields != null) {
                    final CompletableFuture<PlcWriteResponse> delayedResponse =
                        executeWrite((InternalPlcWriteRequest) writeRequest, directAdsFields);
                    delayedResponse.handle((plcReadResponse, throwable1) -> {
                        if (plcReadResponse != null) {
                            delayedWrite.complete(plcReadResponse);
                        } else {
                            delayedWrite.completeExceptionally(throwable1);
                        }
                        return this;
                    });
                } else {
                    delayedWrite.completeExceptionally(throwable);
                }
                return this;
            });
            return delayedWrite;
        }
    }

    protected CompletableFuture<PlcWriteResponse> executeWrite(InternalPlcWriteRequest writeRequest,
                                                             List<DirectAdsField> directAdsFields) {
        // Depending on the number of fields, use a single item request or a sum-request
        if (directAdsFields.size() == 1) {
            // Do a normal (single item) ADS Write Request
            return singleWrite(writeRequest, directAdsFields.get(0));
        } else {
            // TODO: Check if the version of the remote station is at least TwinCAT v2.11 Build >= 1550 otherwise split up into single item requests.
            // Do a ADS-Sum Read Request.
            return multiWrite(writeRequest, directAdsFields);
        }
    }

    protected CompletableFuture<PlcWriteResponse> singleWrite(InternalPlcWriteRequest writeRequest, DirectAdsField directAdsField) {
        CompletableFuture<PlcWriteResponse> future = new CompletableFuture<>();

        final String fieldName = writeRequest.getFieldNames().iterator().next();
        final AdsField plcField = (AdsField) writeRequest.getField(fieldName);
        final PlcValue plcValue = writeRequest.getPlcValue(fieldName);
        try {
            WriteBuffer writeBuffer = DataItemIO.staticSerialize(plcValue, plcField.getAdsDataType(), true);
            AdsData adsData = new AdsWriteRequest(
                directAdsField.getIndexGroup(), directAdsField.getIndexOffset(), writeBuffer.getData());
            AmsPacket amsPacket = new AmsPacket(configuration.getTargetAmsNetId(), configuration.getTargetAmsPort(),
                configuration.getSourceAmsNetId(), configuration.getSourceAmsPort(),
                CommandId.ADS_WRITE, DEFAULT_COMMAND_STATE, 0, getInvokeId(), adsData);
            AmsTCPPacket amsTCPPacket = new AmsTCPPacket(amsPacket);

            // Start a new request-transaction (Is ended in the response-handler)
            RequestTransactionManager.RequestTransaction transaction = tm.startRequest();
            transaction.submit(() -> context.sendRequest(amsTCPPacket)
                .expectResponse(AmsTCPPacket.class, Duration.ofMillis(configuration.getTimeoutRequest()))
                .onTimeout(future::completeExceptionally)
                .onError((p, e) -> future.completeExceptionally(e))
                .check(responseAmsPacket -> responseAmsPacket.getUserdata().getInvokeId() == amsPacket.getInvokeId())
                .unwrap(response -> (AdsWriteResponse) response.getUserdata().getData())
                .handle(responseAdsData -> {
                    if (responseAdsData.getResult() == ReturnCode.OK) {
                        final PlcWriteResponse plcWriteResponse = convertToPlc4xWriteResponse(writeRequest, responseAdsData);
                        // Convert the response from the PLC into a PLC4X Response ...
                        future.complete(plcWriteResponse);
                    } else {
                        // TODO: Implement this correctly.
                        future.completeExceptionally(new PlcException("Error"));
                    }
                    // Finish the request-transaction.
                    transaction.endRequest();
                }));
        } catch (Exception e) {
            future.completeExceptionally(new PlcException("Error"));
        }
        return future;
    }

    protected CompletableFuture<PlcWriteResponse> multiWrite(InternalPlcWriteRequest writeRequest, List<DirectAdsField> directAdsFields) {
        CompletableFuture<PlcWriteResponse> future = new CompletableFuture<>();

        // Calculate the size of all fields together.
        // Calculate the expected size of the response data.
        int expectedRequestDataSize = directAdsFields.stream().mapToInt(
            field -> field.getAdsDataType().getNumBytes() * field.getNumberOfElements()).sum();
        byte[] writeBuffer = new byte[expectedRequestDataSize];
        int pos = 0;
        for (String fieldName : writeRequest.getFieldNames()) {
            final AdsField field = (AdsField) writeRequest.getField(fieldName);
            final PlcValue plcValue = writeRequest.getPlcValue(fieldName);
            try {
                final WriteBuffer itemWriteBuffer = DataItemIO.staticSerialize(plcValue, field.getAdsDataType(), true);
                int numBytes = itemWriteBuffer.getPos();
                System.arraycopy(itemWriteBuffer.getData(), 0, writeBuffer, pos, numBytes);
                pos += numBytes;
            } catch (Exception e) {
                throw new PlcRuntimeException("Error serializing data", e);
            }
        }

        // With multi-requests, the index-group is fixed and the index offset indicates the number of elements.
        AdsData adsData = new AdsReadWriteRequest(
            ReservedIndexGroups.ADSIGRP_MULTIPLE_WRITE.getValue(), directAdsFields.size(), directAdsFields.size() * 4,
            directAdsFields.stream().map(directAdsField -> new AdsMultiRequestItemWrite(
                directAdsField.getIndexGroup(), directAdsField.getIndexOffset(),
                (directAdsField.getAdsDataType().getNumBytes() * directAdsField.getNumberOfElements())))
                .toArray(AdsMultiRequestItem[]::new), writeBuffer);

        AmsPacket amsPacket = new AmsPacket(configuration.getTargetAmsNetId(), configuration.getTargetAmsPort(),
            configuration.getSourceAmsNetId(), configuration.getSourceAmsPort(),
            CommandId.ADS_READ_WRITE, DEFAULT_COMMAND_STATE, 0, getInvokeId(), adsData);
        AmsTCPPacket amsTCPPacket = new AmsTCPPacket(amsPacket);

        // Start a new request-transaction (Is ended in the response-handler)
        RequestTransactionManager.RequestTransaction transaction = tm.startRequest();
        transaction.submit(() -> context.sendRequest(amsTCPPacket)
            .expectResponse(AmsTCPPacket.class, Duration.ofMillis(configuration.getTimeoutRequest()))
            .onTimeout(future::completeExceptionally)
            .onError((p, e) -> future.completeExceptionally(e))
            .check(responseAmsPacket -> responseAmsPacket.getUserdata().getInvokeId() == amsPacket.getInvokeId())
            .unwrap(response -> (AdsReadWriteResponse) response.getUserdata().getData())
            .handle(responseAdsData -> {
                if(responseAdsData.getResult() == ReturnCode.OK) {
                    final PlcWriteResponse plcWriteResponse = convertToPlc4xWriteResponse(writeRequest, responseAdsData);
                    // Convert the response from the PLC into a PLC4X Response ...
                    future.complete(plcWriteResponse);
                } else {
                    // TODO: Implement this correctly.
                    future.completeExceptionally(new PlcException("Error"));
                }
                // Finish the request-transaction.
                transaction.endRequest();
            }));
        return future;
    }

    protected PlcWriteResponse convertToPlc4xWriteResponse(PlcWriteRequest writeRequest, AdsData adsData) {
        Map<String, PlcResponseCode> responseCodes = new HashMap<>();
        if (adsData instanceof AdsWriteResponse) {
            AdsWriteResponse adsWriteResponse = (AdsWriteResponse) adsData;
            responseCodes.put(writeRequest.getFieldNames().stream().findFirst().orElse(""),
                parsePlcResponseCode(adsWriteResponse.getResult()));
        } else if (adsData instanceof AdsReadWriteResponse) {
            AdsReadWriteResponse adsReadWriteResponse = (AdsReadWriteResponse) adsData;
            ReadBuffer readBuffer = new ReadBuffer(adsReadWriteResponse.getData(), true);
            // When parsing a multi-item response, the error codes of each items come
            // in sequence and then come the values.
            for (String fieldName : writeRequest.getFieldNames()) {
                try {
                    final ReturnCode result = ReturnCode.valueOf(readBuffer.readUnsignedLong(32));
                    responseCodes.put(fieldName, parsePlcResponseCode(result));
                } catch (ParseException e) {
                    responseCodes.put(fieldName, PlcResponseCode.INTERNAL_ERROR);
                }
            }
        }

        return new DefaultPlcWriteResponse((InternalPlcWriteRequest) writeRequest, responseCodes);
    }

    @Override
    protected void decode(ConversationContext<AmsTCPPacket> context, AmsTCPPacket msg) throws Exception {
        super.decode(context, msg);
    }

    protected CompletableFuture<List<DirectAdsField>> getDirectAddresses(List<PlcField> fields) {
        CompletableFuture<List<DirectAdsField>> future = new CompletableFuture<>();

        // Get all symbolic fields from the current request.
        // These potentially need to be resolved to direct addresses, if this has not been done before.
        final List<SymbolicAdsField> referencedSymbolicFields = fields.stream()
            .filter(plcField -> plcField instanceof SymbolicAdsField)
            .map(plcField -> (SymbolicAdsField) plcField).collect(Collectors.toList());

        // Find out for which of these symbolic addresses no resolution has been initiated.
        final List<SymbolicAdsField> symbolicFieldsNeedingResolution = referencedSymbolicFields.stream()
            .filter(symbolicAdsField -> !symbolicFieldMapping.containsKey(symbolicAdsField))
            .collect(Collectors.toList());

        // If there are unresolved symbolic addresses, initiate the resolution
        if (!symbolicFieldsNeedingResolution.isEmpty()) {
            // Get a list of symbolic addresses for which no resolution request has been sent yet
            // (A parallel request initiated a bit earlier might have already initiated a resolution
            // which has not yet been completed)
            final List<SymbolicAdsField> requiredResolutionFields =
                symbolicFieldsNeedingResolution.stream().filter(symbolicAdsField ->
                    !pendingResolutionRequests.containsKey(symbolicAdsField)).collect(Collectors.toList());
            // If there are fields for which no resolution request has been sent yet,
            // send a request.
            if (!requiredResolutionFields.isEmpty()) {
                CompletableFuture<Void> resolutionFuture;
                // Create a future which will be completed as soon as the
                // resolution result has been added to the map.
                if (requiredResolutionFields.size() == 1) {
                    SymbolicAdsField symbolicAdsField = requiredResolutionFields.get(0);
                    resolutionFuture = resolveSingleSymbolicAddress(requiredResolutionFields.get(0));
                    pendingResolutionRequests.put(symbolicAdsField, resolutionFuture);
                } else {
                    resolutionFuture = resolveMultipleSymbolicAddresses(requiredResolutionFields);
                    for (SymbolicAdsField symbolicAdsField : requiredResolutionFields) {
                        pendingResolutionRequests.put(symbolicAdsField, resolutionFuture);
                    }
                }
            }

            // Create a global future which is completed as soon as all sub-futures for this request are completed.
            final CompletableFuture<Void> resolutionComplete =
                CompletableFuture.allOf(symbolicFieldsNeedingResolution.stream()
                    .map(symbolicAdsField -> pendingResolutionRequests.get(symbolicAdsField))
                    .toArray(CompletableFuture[]::new));

            // Complete the future asynchronously as soon as all fields are resolved.
            resolutionComplete.handleAsync((unused, throwable) -> {
                return future.complete(fields.stream().map(plcField -> {
                    if (plcField instanceof SymbolicAdsField) {
                        return symbolicFieldMapping.get(plcField);
                    } else {
                        return (DirectAdsField) plcField;
                    }
                }).collect(Collectors.toList()));
            });
        } else {
            // If all fields were resolved, we can continue instantly.
            future.complete(fields.stream().map(plcField -> {
                if (plcField instanceof SymbolicAdsField) {
                    return symbolicFieldMapping.get(plcField);
                } else {
                    return (DirectAdsField) plcField;
                }
            }).collect(Collectors.toList()));
        }
        return future;
    }

    protected CompletableFuture<Void> resolveSingleSymbolicAddress(SymbolicAdsField symbolicAdsField) {
        CompletableFuture<Void> future = new CompletableFuture<>();

        AdsData adsData = new AdsReadWriteRequest(ReservedIndexGroups.ADSIGRP_SYM_HNDBYNAME.getValue(), 0,
            4, null,
            getNullByteTerminatedArray(symbolicAdsField.getSymbolicField()));
        AmsPacket amsPacket = new AmsPacket(configuration.getTargetAmsNetId(), configuration.getTargetAmsPort(),
            configuration.getSourceAmsNetId(), configuration.getSourceAmsPort(),
            CommandId.ADS_READ_WRITE, DEFAULT_COMMAND_STATE, 0, getInvokeId(), adsData);
        AmsTCPPacket amsTCPPacket = new AmsTCPPacket(amsPacket);

        // Start a new request-transaction (Is ended in the response-handler)
        RequestTransactionManager.RequestTransaction transaction = tm.startRequest();
        transaction.submit(() -> context.sendRequest(amsTCPPacket)
            .expectResponse(AmsTCPPacket.class, Duration.ofMillis(configuration.getTimeoutRequest()))
            .onTimeout(future::completeExceptionally)
            .onError((p, e) -> future.completeExceptionally(e))
            .check(responseAmsPacket -> responseAmsPacket.getUserdata().getInvokeId() == amsPacket.getInvokeId())
            .unwrap(response -> response.getUserdata().getData())
            .check(adsDataResponse -> adsDataResponse instanceof AdsReadWriteResponse)
            .unwrap(adsDataResponse -> (AdsReadWriteResponse) adsDataResponse)
            .handle(responseAdsData -> {
                if(responseAdsData.getResult() != ReturnCode.OK) {
                    future.completeExceptionally(new PlcException("Couldn't retrieve handle for symbolic field " +
                        symbolicAdsField.getSymbolicField() + " got return code " + responseAdsData.getResult().name()));
                } else {
                    ReadBuffer readBuffer = new ReadBuffer(responseAdsData.getData(), true);
                    try {
                        // Read the handle.
                        long handle = readBuffer.readUnsignedLong(32);

                        DirectAdsField directAdsField = new DirectAdsField(
                            ReservedIndexGroups.ADSIGRP_SYM_VALBYHND.getValue(), handle,
                            symbolicAdsField.getAdsDataType(), symbolicAdsField.getNumberOfElements());
                        symbolicFieldMapping.put(symbolicAdsField, directAdsField);
                        future.complete(null);
                    } catch (ParseException e) {
                        future.completeExceptionally(e);
                    }
                }
                transaction.endRequest();
            }));
        return future;
    }

    protected CompletableFuture<Void> resolveMultipleSymbolicAddresses(List<SymbolicAdsField> symbolicAdsFields) {
        CompletableFuture<Void> future = new CompletableFuture<>();

        // The expected response for every symbolic address is 12 bytes (8 bytes header and 4 bytes for the handle)
        long expectedResponseDataSize = symbolicAdsFields.size() * 12;
        // Concatenate the string part of each symbolic address into one concattenated string and get the bytes.
        byte[] addressData = symbolicAdsFields.stream().map(
            SymbolicAdsField::getSymbolicField).collect(Collectors.joining("")).getBytes();
        AdsData adsData = new AdsReadWriteRequest(ReservedIndexGroups.ADSIGRP_MULTIPLE_READ_WRITE.getValue(),
            symbolicAdsFields.size(), expectedResponseDataSize, symbolicAdsFields.stream().map(symbolicAdsField ->
            new AdsMultiRequestItemReadWrite(ReservedIndexGroups.ADSIGRP_SYM_HNDBYNAME.getValue(), 0,
                4, symbolicAdsField.getSymbolicField().length())
        ).toArray(AdsMultiRequestItem[]::new), addressData);
        AmsPacket amsPacket = new AmsPacket(configuration.getTargetAmsNetId(), configuration.getTargetAmsPort(),
            configuration.getSourceAmsNetId(), configuration.getSourceAmsPort(),
            CommandId.ADS_READ_WRITE, DEFAULT_COMMAND_STATE, 0, getInvokeId(), adsData);
        AmsTCPPacket amsTCPPacket = new AmsTCPPacket(amsPacket);

        // Start a new request-transaction (Is ended in the response-handler)
        RequestTransactionManager.RequestTransaction transaction = tm.startRequest();
        transaction.submit(() -> context.sendRequest(amsTCPPacket)
            .expectResponse(AmsTCPPacket.class, Duration.ofMillis(configuration.getTimeoutRequest()))
            .onTimeout(future::completeExceptionally)
            .onError((p, e) -> future.completeExceptionally(e))
            .check(responseAmsPacket -> responseAmsPacket.getUserdata().getInvokeId() == amsPacket.getInvokeId())
            .unwrap(response -> response.getUserdata().getData())
            .check(adsDataResponse -> adsDataResponse instanceof AdsReadWriteResponse)
            .unwrap(adsDataResponse -> (AdsReadWriteResponse) adsDataResponse)
            .handle(responseAdsData -> {
                ReadBuffer readBuffer = new ReadBuffer(responseAdsData.getData(), true);
                Map<SymbolicAdsField, Long> returnCodes = new HashMap<>();
                // In the response first come the return codes and the data-lengths for each item.
                symbolicAdsFields.forEach(symbolicAdsField -> {
                    try {
                        // This should be 0 in the success case.
                        long returnCode = readBuffer.readUnsignedLong(32);
                        // This is always 4
                        long itemLength = readBuffer.readUnsignedLong(32);

                        returnCodes.put(symbolicAdsField, returnCode);
                    } catch (ParseException e) {
                        e.printStackTrace();
                    }
                });
                // After reading the header-information, comes the data itself.
                symbolicAdsFields.forEach(symbolicAdsField -> {
                    try {
                        if (returnCodes.get(symbolicAdsField) == 0) {
                            // Read the handle.
                            long handle = readBuffer.readUnsignedLong(32);

                            DirectAdsField directAdsField = new DirectAdsField(
                                ReservedIndexGroups.ADSIGRP_SYM_VALBYHND.getValue(), handle,
                                symbolicAdsField.getAdsDataType(), symbolicAdsField.getNumberOfElements());
                            symbolicFieldMapping.put(symbolicAdsField, directAdsField);
                        } else {
                            // TODO: Handle the case of unsuccessful resolution ..
                        }
                    } catch (ParseException e) {
                        e.printStackTrace();
                    }
                });
                future.complete(null);
                transaction.endRequest();
            }));
        return future;
    }

    protected long getInvokeId() {
        long invokeId = invokeIdGenerator.getAndIncrement();
        // If we've reached the max value for a 16 bit transaction identifier, reset back to 1
        if(invokeIdGenerator.get() == 0xFFFFFFFF) {
            invokeIdGenerator.set(1);
        }
        return invokeId;
    }

    protected byte[] getNullByteTerminatedArray(String value) {
        byte[] valueBytes = value.getBytes();
        byte[] nullTerminatedBytes = new byte[valueBytes.length + 1];
        System.arraycopy(valueBytes, 0, nullTerminatedBytes, 0, valueBytes.length);
        return nullTerminatedBytes;
    }

}
