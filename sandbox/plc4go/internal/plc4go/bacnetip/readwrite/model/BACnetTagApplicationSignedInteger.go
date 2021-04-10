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
	"plc4x.apache.org/plc4go-modbus-driver/0.8.0/internal/plc4go/spi"
)

// The data-structure of this message
type BACnetTagApplicationSignedInteger struct {
	data []int8
	BACnetTag
}

// The corresponding interface
type IBACnetTagApplicationSignedInteger interface {
	IBACnetTag
	Serialize(io spi.WriteBuffer)
}

// Accessors for discriminator values.
func (m BACnetTagApplicationSignedInteger) ContextSpecificTag() uint8 {
	return 0
}

func (m BACnetTagApplicationSignedInteger) initialize(typeOrTagNumber uint8, lengthValueType uint8, extTagNumber *uint8, extLength *uint8) spi.Message {
	m.typeOrTagNumber = typeOrTagNumber
	m.lengthValueType = lengthValueType
	m.extTagNumber = extTagNumber
	m.extLength = extLength
	return m
}

func NewBACnetTagApplicationSignedInteger(data []int8) BACnetTagInitializer {
	return &BACnetTagApplicationSignedInteger{data: data}
}

func (m BACnetTagApplicationSignedInteger) LengthInBits() uint16 {
	var lengthInBits uint16 = m.BACnetTag.LengthInBits()

	// Array field
	if len(m.data) > 0 {
		lengthInBits += 8 * uint16(len(m.data))
	}

	return lengthInBits
}

func (m BACnetTagApplicationSignedInteger) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetTagApplicationSignedIntegerParse(io spi.ReadBuffer, lengthValueType uint8, extLength uint8) (BACnetTagInitializer, error) {

	// Array field (data)
	var data []int8
	// Length array
	_dataLength := uint16(spi.InlineIf(((lengthValueType) == (5)), uint16(extLength), uint16(lengthValueType)))
	_dataEndPos := io.GetPos() + _dataLength
	for io.GetPos() < _dataEndPos {
		data = append(data, io.ReadInt8(8))
	}

	// Create the instance
	return NewBACnetTagApplicationSignedInteger(data), nil
}

func (m BACnetTagApplicationSignedInteger) Serialize(io spi.WriteBuffer) {
	serializeFunc := func(typ interface{}) {
		if _, ok := typ.(IBACnetTagApplicationSignedInteger); ok {

			// Array Field (data)
			if m.data != nil {
				for _, _element := range m.data {
					io.WriteInt8(8, _element)
				}
			}
		}
	}
	serializeFunc(m)
}
