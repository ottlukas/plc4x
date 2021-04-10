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
	"reflect"
	"strconv"
)

// Constant values.
const BACnetConfirmedServiceRequestWriteProperty_OBJECTIDENTIFIERHEADER uint8 = 0x0C
const BACnetConfirmedServiceRequestWriteProperty_PROPERTYIDENTIFIERHEADER uint8 = 0x03
const BACnetConfirmedServiceRequestWriteProperty_OPENINGTAG uint8 = 0x3E
const BACnetConfirmedServiceRequestWriteProperty_CLOSINGTAG uint8 = 0x3F

// The data-structure of this message
type BACnetConfirmedServiceRequestWriteProperty struct {
	objectType               uint16
	objectInstanceNumber     uint32
	propertyIdentifierLength uint8
	propertyIdentifier       []int8
	value                    BACnetTag
	priority                 *BACnetTag
	BACnetConfirmedServiceRequest
}

// The corresponding interface
type IBACnetConfirmedServiceRequestWriteProperty interface {
	IBACnetConfirmedServiceRequest
	Serialize(io spi.WriteBuffer)
}

// Accessors for discriminator values.
func (m BACnetConfirmedServiceRequestWriteProperty) ServiceChoice() uint8 {
	return 0x0F
}

func (m BACnetConfirmedServiceRequestWriteProperty) initialize() spi.Message {
	return m
}

func NewBACnetConfirmedServiceRequestWriteProperty(objectType uint16, objectInstanceNumber uint32, propertyIdentifierLength uint8, propertyIdentifier []int8, value BACnetTag, priority *BACnetTag) BACnetConfirmedServiceRequestInitializer {
	return &BACnetConfirmedServiceRequestWriteProperty{objectType: objectType, objectInstanceNumber: objectInstanceNumber, propertyIdentifierLength: propertyIdentifierLength, propertyIdentifier: propertyIdentifier, value: value, priority: priority}
}

func (m BACnetConfirmedServiceRequestWriteProperty) LengthInBits() uint16 {
	var lengthInBits uint16 = m.BACnetConfirmedServiceRequest.LengthInBits()

	// Const Field (objectIdentifierHeader)
	lengthInBits += 8

	// Simple field (objectType)
	lengthInBits += 10

	// Simple field (objectInstanceNumber)
	lengthInBits += 22

	// Const Field (propertyIdentifierHeader)
	lengthInBits += 5

	// Simple field (propertyIdentifierLength)
	lengthInBits += 3

	// Array field
	if len(m.propertyIdentifier) > 0 {
		lengthInBits += 8 * uint16(len(m.propertyIdentifier))
	}

	// Const Field (openingTag)
	lengthInBits += 8

	// Simple field (value)
	lengthInBits += m.value.LengthInBits()

	// Const Field (closingTag)
	lengthInBits += 8

	// Optional Field (priority)
	if m.priority != nil {
		lengthInBits += m.priority.LengthInBits()
	}

	return lengthInBits
}

func (m BACnetConfirmedServiceRequestWriteProperty) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetConfirmedServiceRequestWritePropertyParse(io spi.ReadBuffer, len uint16) (BACnetConfirmedServiceRequestInitializer, error) {
	var startPos = io.GetPos()
	var curPos uint16

	// Const Field (objectIdentifierHeader)
	var objectIdentifierHeader uint8 = io.ReadUint8(8)
	if objectIdentifierHeader != BACnetConfirmedServiceRequestWriteProperty_OBJECTIDENTIFIERHEADER {
		return nil, errors.New("Expected constant value " + strconv.Itoa(int(BACnetConfirmedServiceRequestWriteProperty_OBJECTIDENTIFIERHEADER)) + " but got " + strconv.Itoa(int(objectIdentifierHeader)))
	}

	// Simple Field (objectType)
	var objectType uint16 = io.ReadUint16(10)

	// Simple Field (objectInstanceNumber)
	var objectInstanceNumber uint32 = io.ReadUint32(22)

	// Const Field (propertyIdentifierHeader)
	var propertyIdentifierHeader uint8 = io.ReadUint8(5)
	if propertyIdentifierHeader != BACnetConfirmedServiceRequestWriteProperty_PROPERTYIDENTIFIERHEADER {
		return nil, errors.New("Expected constant value " + strconv.Itoa(int(BACnetConfirmedServiceRequestWriteProperty_PROPERTYIDENTIFIERHEADER)) + " but got " + strconv.Itoa(int(propertyIdentifierHeader)))
	}

	// Simple Field (propertyIdentifierLength)
	var propertyIdentifierLength uint8 = io.ReadUint8(3)

	// Array field (propertyIdentifier)
	var propertyIdentifier []int8
	// Count array
	{
		propertyIdentifier := make([]int8, propertyIdentifierLength)
		for curItem := uint16(0); curItem < uint16(propertyIdentifierLength); curItem++ {

			propertyIdentifier = append(propertyIdentifier, io.ReadInt8(8))
		}
	}

	// Const Field (openingTag)
	var openingTag uint8 = io.ReadUint8(8)
	if openingTag != BACnetConfirmedServiceRequestWriteProperty_OPENINGTAG {
		return nil, errors.New("Expected constant value " + strconv.Itoa(int(BACnetConfirmedServiceRequestWriteProperty_OPENINGTAG)) + " but got " + strconv.Itoa(int(openingTag)))
	}

	// Simple Field (value)
	_valueMessage, _err := BACnetTagParse(io)
	if _err != nil {
		return nil, errors.New("Error parsing simple field 'value'. " + _err.Error())
	}
	var value BACnetTag
	value, _valueOk := _valueMessage.(BACnetTag)
	if !_valueOk {
		return nil, errors.New("Couldn't cast message of type " + reflect.TypeOf(_valueMessage).Name() + " to BACnetTag")
	}

	// Const Field (closingTag)
	var closingTag uint8 = io.ReadUint8(8)
	if closingTag != BACnetConfirmedServiceRequestWriteProperty_CLOSINGTAG {
		return nil, errors.New("Expected constant value " + strconv.Itoa(int(BACnetConfirmedServiceRequestWriteProperty_CLOSINGTAG)) + " but got " + strconv.Itoa(int(closingTag)))
	}

	// Optional Field (priority) (Can be skipped, if a given expression evaluates to false)
	curPos = io.GetPos() - startPos
	var priority *BACnetTag = nil
	if (curPos) < ((len) - (1)) {
		_message, _err := BACnetTagParse(io)
		if _err != nil {
			return nil, errors.New("Error parsing 'priority' field " + _err.Error())
		}
		var _item BACnetTag
		_item, _ok := _message.(BACnetTag)
		if !_ok {
			return nil, errors.New("Couldn't cast message of type " + reflect.TypeOf(_item).Name() + " to BACnetTag")
		}
		priority = &_item
	}

	// Create the instance
	return NewBACnetConfirmedServiceRequestWriteProperty(objectType, objectInstanceNumber, propertyIdentifierLength, propertyIdentifier, value, priority), nil
}

func (m BACnetConfirmedServiceRequestWriteProperty) Serialize(io spi.WriteBuffer) {
	serializeFunc := func(typ interface{}) {
		if _, ok := typ.(IBACnetConfirmedServiceRequestWriteProperty); ok {

			// Const Field (objectIdentifierHeader)
			io.WriteUint8(8, 0x0C)

			// Simple Field (objectType)
			var objectType uint16 = m.objectType
			io.WriteUint16(10, (objectType))

			// Simple Field (objectInstanceNumber)
			var objectInstanceNumber uint32 = m.objectInstanceNumber
			io.WriteUint32(22, (objectInstanceNumber))

			// Const Field (propertyIdentifierHeader)
			io.WriteUint8(5, 0x03)

			// Simple Field (propertyIdentifierLength)
			var propertyIdentifierLength uint8 = m.propertyIdentifierLength
			io.WriteUint8(3, (propertyIdentifierLength))

			// Array Field (propertyIdentifier)
			if m.propertyIdentifier != nil {
				for _, _element := range m.propertyIdentifier {
					io.WriteInt8(8, _element)
				}
			}

			// Const Field (openingTag)
			io.WriteUint8(8, 0x3E)

			// Simple Field (value)
			var value BACnetTag = m.value
			value.Serialize(io)

			// Const Field (closingTag)
			io.WriteUint8(8, 0x3F)

			// Optional Field (priority) (Can be skipped, if the value is null)
			var priority *BACnetTag = nil
			if m.priority != nil {
				priority = m.priority
				priority.Serialize(io)
			}
		}
	}
	serializeFunc(m)
}
