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
const BACnetTagWithContent_OPENTAG uint8 = 0x2e
const BACnetTagWithContent_CLOSINGTAG uint8 = 0x2f

// The data-structure of this message
type BACnetTagWithContent struct {
	typeOrTagNumber    uint8
	contextSpecificTag uint8
	lengthValueType    uint8
	extTagNumber       *uint8
	extLength          *uint8
	propertyIdentifier []uint8
	value              IBACnetTag
}

// The corresponding interface
type IBACnetTagWithContent interface {
	spi.Message
	Serialize(io spi.WriteBuffer) error
}

func NewBACnetTagWithContent(typeOrTagNumber uint8, contextSpecificTag uint8, lengthValueType uint8, extTagNumber *uint8, extLength *uint8, propertyIdentifier []uint8, value IBACnetTag) spi.Message {
	return &BACnetTagWithContent{typeOrTagNumber: typeOrTagNumber, contextSpecificTag: contextSpecificTag, lengthValueType: lengthValueType, extTagNumber: extTagNumber, extLength: extLength, propertyIdentifier: propertyIdentifier, value: value}
}

func CastIBACnetTagWithContent(structType interface{}) IBACnetTagWithContent {
	castFunc := func(typ interface{}) IBACnetTagWithContent {
		if iBACnetTagWithContent, ok := typ.(IBACnetTagWithContent); ok {
			return iBACnetTagWithContent
		}
		return nil
	}
	return castFunc(structType)
}

func CastBACnetTagWithContent(structType interface{}) BACnetTagWithContent {
	castFunc := func(typ interface{}) BACnetTagWithContent {
		if sBACnetTagWithContent, ok := typ.(BACnetTagWithContent); ok {
			return sBACnetTagWithContent
		}
		return BACnetTagWithContent{}
	}
	return castFunc(structType)
}

func (m BACnetTagWithContent) LengthInBits() uint16 {
	var lengthInBits uint16 = 0

	// Simple field (typeOrTagNumber)
	lengthInBits += 4

	// Simple field (contextSpecificTag)
	lengthInBits += 1

	// Simple field (lengthValueType)
	lengthInBits += 3

	// Optional Field (extTagNumber)
	if m.extTagNumber != nil {
		lengthInBits += 8
	}

	// Optional Field (extLength)
	if m.extLength != nil {
		lengthInBits += 8
	}

	// Array field
	if len(m.propertyIdentifier) > 0 {
		lengthInBits += 8 * uint16(len(m.propertyIdentifier))
	}

	// Const Field (openTag)
	lengthInBits += 8

	// Simple field (value)
	lengthInBits += m.value.LengthInBits()

	// Const Field (closingTag)
	lengthInBits += 8

	return lengthInBits
}

func (m BACnetTagWithContent) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetTagWithContentParse(io *spi.ReadBuffer) (spi.Message, error) {

	// Simple Field (typeOrTagNumber)
	typeOrTagNumber, _typeOrTagNumberErr := io.ReadUint8(4)
	if _typeOrTagNumberErr != nil {
		return nil, errors.New("Error parsing 'typeOrTagNumber' field " + _typeOrTagNumberErr.Error())
	}

	// Simple Field (contextSpecificTag)
	contextSpecificTag, _contextSpecificTagErr := io.ReadUint8(1)
	if _contextSpecificTagErr != nil {
		return nil, errors.New("Error parsing 'contextSpecificTag' field " + _contextSpecificTagErr.Error())
	}

	// Simple Field (lengthValueType)
	lengthValueType, _lengthValueTypeErr := io.ReadUint8(3)
	if _lengthValueTypeErr != nil {
		return nil, errors.New("Error parsing 'lengthValueType' field " + _lengthValueTypeErr.Error())
	}

	// Optional Field (extTagNumber) (Can be skipped, if a given expression evaluates to false)
	var extTagNumber *uint8 = nil
	if bool((typeOrTagNumber) == (15)) {
		_val, _err := io.ReadUint8(8)
		if _err != nil {
			return nil, errors.New("Error parsing 'extTagNumber' field " + _err.Error())
		}

		extTagNumber = &_val
	}

	// Optional Field (extLength) (Can be skipped, if a given expression evaluates to false)
	var extLength *uint8 = nil
	if bool((lengthValueType) == (5)) {
		_val, _err := io.ReadUint8(8)
		if _err != nil {
			return nil, errors.New("Error parsing 'extLength' field " + _err.Error())
		}

		extLength = &_val
	}

	// Array field (propertyIdentifier)
	// Length array
	propertyIdentifier := make([]uint8, 0)
	_propertyIdentifierLength := spi.InlineIf(bool(bool((lengthValueType) == (5))), uint16((*extLength)), uint16(lengthValueType))
	_propertyIdentifierEndPos := io.GetPos() + uint16(_propertyIdentifierLength)
	for io.GetPos() < _propertyIdentifierEndPos {
		_item, _err := io.ReadUint8(8)
		if _err != nil {
			return nil, errors.New("Error parsing 'propertyIdentifier' field " + _err.Error())
		}
		propertyIdentifier = append(propertyIdentifier, _item)
	}

	// Const Field (openTag)
	openTag, _openTagErr := io.ReadUint8(8)
	if _openTagErr != nil {
		return nil, errors.New("Error parsing 'openTag' field " + _openTagErr.Error())
	}
	if openTag != BACnetTagWithContent_OPENTAG {
		return nil, errors.New("Expected constant value " + strconv.Itoa(int(BACnetTagWithContent_OPENTAG)) + " but got " + strconv.Itoa(int(openTag)))
	}

	// Simple Field (value)
	_valueMessage, _err := BACnetTagParse(io)
	if _err != nil {
		return nil, errors.New("Error parsing simple field 'value'. " + _err.Error())
	}
	var value IBACnetTag
	value, _valueOk := _valueMessage.(IBACnetTag)
	if !_valueOk {
		return nil, errors.New("Couldn't cast message of type " + reflect.TypeOf(_valueMessage).Name() + " to IBACnetTag")
	}

	// Const Field (closingTag)
	closingTag, _closingTagErr := io.ReadUint8(8)
	if _closingTagErr != nil {
		return nil, errors.New("Error parsing 'closingTag' field " + _closingTagErr.Error())
	}
	if closingTag != BACnetTagWithContent_CLOSINGTAG {
		return nil, errors.New("Expected constant value " + strconv.Itoa(int(BACnetTagWithContent_CLOSINGTAG)) + " but got " + strconv.Itoa(int(closingTag)))
	}

	// Create the instance
	return NewBACnetTagWithContent(typeOrTagNumber, contextSpecificTag, lengthValueType, extTagNumber, extLength, propertyIdentifier, value), nil
}

func (m BACnetTagWithContent) Serialize(io spi.WriteBuffer) error {

	// Simple Field (typeOrTagNumber)
	typeOrTagNumber := uint8(m.typeOrTagNumber)
	_typeOrTagNumberErr := io.WriteUint8(4, (typeOrTagNumber))
	if _typeOrTagNumberErr != nil {
		return errors.New("Error serializing 'typeOrTagNumber' field " + _typeOrTagNumberErr.Error())
	}

	// Simple Field (contextSpecificTag)
	contextSpecificTag := uint8(m.contextSpecificTag)
	_contextSpecificTagErr := io.WriteUint8(1, (contextSpecificTag))
	if _contextSpecificTagErr != nil {
		return errors.New("Error serializing 'contextSpecificTag' field " + _contextSpecificTagErr.Error())
	}

	// Simple Field (lengthValueType)
	lengthValueType := uint8(m.lengthValueType)
	_lengthValueTypeErr := io.WriteUint8(3, (lengthValueType))
	if _lengthValueTypeErr != nil {
		return errors.New("Error serializing 'lengthValueType' field " + _lengthValueTypeErr.Error())
	}

	// Optional Field (extTagNumber) (Can be skipped, if the value is null)
	var extTagNumber *uint8 = nil
	if m.extTagNumber != nil {
		extTagNumber = m.extTagNumber
		_extTagNumberErr := io.WriteUint8(8, *(extTagNumber))
		if _extTagNumberErr != nil {
			return errors.New("Error serializing 'extTagNumber' field " + _extTagNumberErr.Error())
		}
	}

	// Optional Field (extLength) (Can be skipped, if the value is null)
	var extLength *uint8 = nil
	if m.extLength != nil {
		extLength = m.extLength
		_extLengthErr := io.WriteUint8(8, *(extLength))
		if _extLengthErr != nil {
			return errors.New("Error serializing 'extLength' field " + _extLengthErr.Error())
		}
	}

	// Array Field (propertyIdentifier)
	if m.propertyIdentifier != nil {
		for _, _element := range m.propertyIdentifier {
			_elementErr := io.WriteUint8(8, _element)
			if _elementErr != nil {
				return errors.New("Error serializing 'propertyIdentifier' field " + _elementErr.Error())
			}
		}
	}

	// Const Field (openTag)
	_openTagErr := io.WriteUint8(8, 0x2e)
	if _openTagErr != nil {
		return errors.New("Error serializing 'openTag' field " + _openTagErr.Error())
	}

	// Simple Field (value)
	value := CastIBACnetTag(m.value)
	_valueErr := value.Serialize(io)
	if _valueErr != nil {
		return errors.New("Error serializing 'value' field " + _valueErr.Error())
	}

	// Const Field (closingTag)
	_closingTagErr := io.WriteUint8(8, 0x2f)
	if _closingTagErr != nil {
		return errors.New("Error serializing 'closingTag' field " + _closingTagErr.Error())
	}

	return nil
}
