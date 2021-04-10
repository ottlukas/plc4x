//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//
package model

import (
    "errors"
    "plc4x.apache.org/plc4go-modbus-driver/v0/internal/plc4go/spi"
    "plc4x.apache.org/plc4go-modbus-driver/v0/internal/plc4go/utils"
)

type BacnetipParserHelper struct {
}

func (m BacnetipParserHelper) Parse(typeName string, arguments []string, io *utils.ReadBuffer) (spi.Message, error) {
    switch typeName {
    case "APDU":
        apduLength, err := utils.StrToUint16(arguments[0])
        if err != nil {
            return nil, err
        }
        return APDUParse(io, apduLength)
    case "BACnetTag":
        return BACnetTagParse(io)
    case "BACnetTagWithContent":
        return BACnetTagWithContentParse(io)
    case "BACnetError":
        return BACnetErrorParse(io)
    case "NLM":
        apduLength, err := utils.StrToUint16(arguments[0])
        if err != nil {
            return nil, err
        }
        return NLMParse(io, apduLength)
    case "BACnetConfirmedServiceRequest":
        len, err := utils.StrToUint16(arguments[0])
        if err != nil {
            return nil, err
        }
        return BACnetConfirmedServiceRequestParse(io, len)
    case "BACnetAddress":
        return BACnetAddressParse(io)
    case "BACnetConfirmedServiceACK":
        return BACnetConfirmedServiceACKParse(io)
    case "BACnetUnconfirmedServiceRequest":
        len, err := utils.StrToUint16(arguments[0])
        if err != nil {
            return nil, err
        }
        return BACnetUnconfirmedServiceRequestParse(io, len)
    case "BACnetServiceAck":
        return BACnetServiceAckParse(io)
    case "BVLC":
        return BVLCParse(io)
    case "NPDU":
        npduLength, err := utils.StrToUint16(arguments[0])
        if err != nil {
            return nil, err
        }
        return NPDUParse(io, npduLength)
    }
    return nil, errors.New("Unsupported type " + typeName)
}
