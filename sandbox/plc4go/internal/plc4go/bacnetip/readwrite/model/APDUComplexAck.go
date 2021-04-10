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
	log "github.com/sirupsen/logrus"
	"plc4x.apache.org/plc4go-modbus-driver/0.8.0/internal/plc4go/spi"
	"reflect"
)

// The data-structure of this message
type APDUComplexAck struct {
	segmentedMessage   bool
	moreFollows        bool
	originalInvokeId   uint8
	sequenceNumber     *uint8
	proposedWindowSize *uint8
	serviceAck         IBACnetServiceAck
	APDU
}

// The corresponding interface
type IAPDUComplexAck interface {
	IAPDU
	Serialize(io spi.WriteBuffer)
}

// Accessors for discriminator values.
func (m APDUComplexAck) ApduType() uint8 {
	return 0x3
}

func (m APDUComplexAck) initialize() spi.Message {
	return m
}

func NewAPDUComplexAck(segmentedMessage bool, moreFollows bool, originalInvokeId uint8, sequenceNumber *uint8, proposedWindowSize *uint8, serviceAck IBACnetServiceAck) APDUInitializer {
	return &APDUComplexAck{segmentedMessage: segmentedMessage, moreFollows: moreFollows, originalInvokeId: originalInvokeId, sequenceNumber: sequenceNumber, proposedWindowSize: proposedWindowSize, serviceAck: serviceAck}
}

func CastIAPDUComplexAck(structType interface{}) IAPDUComplexAck {
	castFunc := func(typ interface{}) IAPDUComplexAck {
		if iAPDUComplexAck, ok := typ.(IAPDUComplexAck); ok {
			return iAPDUComplexAck
		}
		return nil
	}
	return castFunc(structType)
}

func CastAPDUComplexAck(structType interface{}) APDUComplexAck {
	castFunc := func(typ interface{}) APDUComplexAck {
		if sAPDUComplexAck, ok := typ.(APDUComplexAck); ok {
			return sAPDUComplexAck
		}
		return APDUComplexAck{}
	}
	return castFunc(structType)
}

func (m APDUComplexAck) LengthInBits() uint16 {
	var lengthInBits uint16 = m.APDU.LengthInBits()

	// Simple field (segmentedMessage)
	lengthInBits += 1

	// Simple field (moreFollows)
	lengthInBits += 1

	// Reserved Field (reserved)
	lengthInBits += 2

	// Simple field (originalInvokeId)
	lengthInBits += 8

	// Optional Field (sequenceNumber)
	if m.sequenceNumber != nil {
		lengthInBits += 8
	}

	// Optional Field (proposedWindowSize)
	if m.proposedWindowSize != nil {
		lengthInBits += 8
	}

	// Simple field (serviceAck)
	lengthInBits += m.serviceAck.LengthInBits()

	return lengthInBits
}

func (m APDUComplexAck) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func APDUComplexAckParse(io *spi.ReadBuffer) (APDUInitializer, error) {

	// Simple Field (segmentedMessage)
	segmentedMessage, _segmentedMessageErr := io.ReadBit()
	if _segmentedMessageErr != nil {
		return nil, errors.New("Error parsing 'segmentedMessage' field " + _segmentedMessageErr.Error())
	}

	// Simple Field (moreFollows)
	moreFollows, _moreFollowsErr := io.ReadBit()
	if _moreFollowsErr != nil {
		return nil, errors.New("Error parsing 'moreFollows' field " + _moreFollowsErr.Error())
	}

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := io.ReadUint8(2)
		if _err != nil {
			return nil, errors.New("Error parsing 'reserved' field " + _err.Error())
		}
		if reserved != uint8(0) {
			log.WithFields(log.Fields{
				"expected value": uint8(0),
				"got value":      reserved,
			}).Info("Got unexpected response.")
		}
	}

	// Simple Field (originalInvokeId)
	originalInvokeId, _originalInvokeIdErr := io.ReadUint8(8)
	if _originalInvokeIdErr != nil {
		return nil, errors.New("Error parsing 'originalInvokeId' field " + _originalInvokeIdErr.Error())
	}

	// Optional Field (sequenceNumber) (Can be skipped, if a given expression evaluates to false)
	var sequenceNumber *uint8 = nil
	if segmentedMessage {
		_val, _err := io.ReadUint8(8)
		if _err != nil {
			return nil, errors.New("Error parsing 'sequenceNumber' field " + _err.Error())
		}

		sequenceNumber = &_val
	}

	// Optional Field (proposedWindowSize) (Can be skipped, if a given expression evaluates to false)
	var proposedWindowSize *uint8 = nil
	if segmentedMessage {
		_val, _err := io.ReadUint8(8)
		if _err != nil {
			return nil, errors.New("Error parsing 'proposedWindowSize' field " + _err.Error())
		}

		proposedWindowSize = &_val
	}

	// Simple Field (serviceAck)
	_serviceAckMessage, _err := BACnetServiceAckParse(io)
	if _err != nil {
		return nil, errors.New("Error parsing simple field 'serviceAck'. " + _err.Error())
	}
	var serviceAck IBACnetServiceAck
	serviceAck, _serviceAckOk := _serviceAckMessage.(IBACnetServiceAck)
	if !_serviceAckOk {
		return nil, errors.New("Couldn't cast message of type " + reflect.TypeOf(_serviceAckMessage).Name() + " to IBACnetServiceAck")
	}

	// Create the instance
	return NewAPDUComplexAck(segmentedMessage, moreFollows, originalInvokeId, sequenceNumber, proposedWindowSize, serviceAck), nil
}

func (m APDUComplexAck) Serialize(io spi.WriteBuffer) {
	ser := func() {

		// Simple Field (segmentedMessage)
		segmentedMessage := bool(m.segmentedMessage)
		io.WriteBit((bool)(segmentedMessage))

		// Simple Field (moreFollows)
		moreFollows := bool(m.moreFollows)
		io.WriteBit((bool)(moreFollows))

		// Reserved Field (reserved)
		io.WriteUint8(2, uint8(0))

		// Simple Field (originalInvokeId)
		originalInvokeId := uint8(m.originalInvokeId)
		io.WriteUint8(8, (originalInvokeId))

		// Optional Field (sequenceNumber) (Can be skipped, if the value is null)
		var sequenceNumber *uint8 = nil
		if m.sequenceNumber != nil {
			sequenceNumber = m.sequenceNumber
			io.WriteUint8(8, *(sequenceNumber))
		}

		// Optional Field (proposedWindowSize) (Can be skipped, if the value is null)
		var proposedWindowSize *uint8 = nil
		if m.proposedWindowSize != nil {
			proposedWindowSize = m.proposedWindowSize
			io.WriteUint8(8, *(proposedWindowSize))
		}

		// Simple Field (serviceAck)
		serviceAck := CastIBACnetServiceAck(m.serviceAck)
		serviceAck.Serialize(io)

	}
	APDUSerialize(io, m.APDU, CastIAPDU(m), ser)
}
