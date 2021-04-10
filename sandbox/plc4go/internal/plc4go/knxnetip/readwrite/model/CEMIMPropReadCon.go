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
	"plc4x.apache.org/plc4go-modbus-driver/0.8.0/internal/plc4go/spi"
)

// The data-structure of this message
type CEMIMPropReadCon struct {
	interfaceObjectType uint16
	objectInstance      uint8
	propertyId          uint8
	numberOfElements    uint8
	startIndex          uint16
	unknown             uint16
	CEMI
}

// The corresponding interface
type ICEMIMPropReadCon interface {
	ICEMI
	Serialize(io spi.WriteBuffer)
}

// Accessors for discriminator values.
func (m CEMIMPropReadCon) MessageCode() uint8 {
	return 0xFB
}

func (m CEMIMPropReadCon) initialize() spi.Message {
	return m
}

func NewCEMIMPropReadCon(interfaceObjectType uint16, objectInstance uint8, propertyId uint8, numberOfElements uint8, startIndex uint16, unknown uint16) CEMIInitializer {
	return &CEMIMPropReadCon{interfaceObjectType: interfaceObjectType, objectInstance: objectInstance, propertyId: propertyId, numberOfElements: numberOfElements, startIndex: startIndex, unknown: unknown}
}

func CastICEMIMPropReadCon(structType interface{}) ICEMIMPropReadCon {
	castFunc := func(typ interface{}) ICEMIMPropReadCon {
		if iCEMIMPropReadCon, ok := typ.(ICEMIMPropReadCon); ok {
			return iCEMIMPropReadCon
		}
		return nil
	}
	return castFunc(structType)
}

func CastCEMIMPropReadCon(structType interface{}) CEMIMPropReadCon {
	castFunc := func(typ interface{}) CEMIMPropReadCon {
		if sCEMIMPropReadCon, ok := typ.(CEMIMPropReadCon); ok {
			return sCEMIMPropReadCon
		}
		return CEMIMPropReadCon{}
	}
	return castFunc(structType)
}

func (m CEMIMPropReadCon) LengthInBits() uint16 {
	var lengthInBits uint16 = m.CEMI.LengthInBits()

	// Simple field (interfaceObjectType)
	lengthInBits += 16

	// Simple field (objectInstance)
	lengthInBits += 8

	// Simple field (propertyId)
	lengthInBits += 8

	// Simple field (numberOfElements)
	lengthInBits += 4

	// Simple field (startIndex)
	lengthInBits += 12

	// Simple field (unknown)
	lengthInBits += 16

	return lengthInBits
}

func (m CEMIMPropReadCon) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func CEMIMPropReadConParse(io *spi.ReadBuffer) (CEMIInitializer, error) {

	// Simple Field (interfaceObjectType)
	interfaceObjectType, _interfaceObjectTypeErr := io.ReadUint16(16)
	if _interfaceObjectTypeErr != nil {
		return nil, errors.New("Error parsing 'interfaceObjectType' field " + _interfaceObjectTypeErr.Error())
	}

	// Simple Field (objectInstance)
	objectInstance, _objectInstanceErr := io.ReadUint8(8)
	if _objectInstanceErr != nil {
		return nil, errors.New("Error parsing 'objectInstance' field " + _objectInstanceErr.Error())
	}

	// Simple Field (propertyId)
	propertyId, _propertyIdErr := io.ReadUint8(8)
	if _propertyIdErr != nil {
		return nil, errors.New("Error parsing 'propertyId' field " + _propertyIdErr.Error())
	}

	// Simple Field (numberOfElements)
	numberOfElements, _numberOfElementsErr := io.ReadUint8(4)
	if _numberOfElementsErr != nil {
		return nil, errors.New("Error parsing 'numberOfElements' field " + _numberOfElementsErr.Error())
	}

	// Simple Field (startIndex)
	startIndex, _startIndexErr := io.ReadUint16(12)
	if _startIndexErr != nil {
		return nil, errors.New("Error parsing 'startIndex' field " + _startIndexErr.Error())
	}

	// Simple Field (unknown)
	unknown, _unknownErr := io.ReadUint16(16)
	if _unknownErr != nil {
		return nil, errors.New("Error parsing 'unknown' field " + _unknownErr.Error())
	}

	// Create the instance
	return NewCEMIMPropReadCon(interfaceObjectType, objectInstance, propertyId, numberOfElements, startIndex, unknown), nil
}

func (m CEMIMPropReadCon) Serialize(io spi.WriteBuffer) {
	ser := func() {

		// Simple Field (interfaceObjectType)
		interfaceObjectType := uint16(m.interfaceObjectType)
		io.WriteUint16(16, (interfaceObjectType))

		// Simple Field (objectInstance)
		objectInstance := uint8(m.objectInstance)
		io.WriteUint8(8, (objectInstance))

		// Simple Field (propertyId)
		propertyId := uint8(m.propertyId)
		io.WriteUint8(8, (propertyId))

		// Simple Field (numberOfElements)
		numberOfElements := uint8(m.numberOfElements)
		io.WriteUint8(4, (numberOfElements))

		// Simple Field (startIndex)
		startIndex := uint16(m.startIndex)
		io.WriteUint16(12, (startIndex))

		// Simple Field (unknown)
		unknown := uint16(m.unknown)
		io.WriteUint16(16, (unknown))

	}
	CEMISerialize(io, m.CEMI, CastICEMI(m), ser)
}
