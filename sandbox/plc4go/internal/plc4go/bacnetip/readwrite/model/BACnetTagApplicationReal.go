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
type BACnetTagApplicationReal struct {
	value float32
	BACnetTag
}

// The corresponding interface
type IBACnetTagApplicationReal interface {
	IBACnetTag
	Serialize(io spi.WriteBuffer)
}

// Accessors for discriminator values.
func (m BACnetTagApplicationReal) ContextSpecificTag() uint8 {
	return 0
}

func (m BACnetTagApplicationReal) initialize(typeOrTagNumber uint8, lengthValueType uint8, extTagNumber *uint8, extLength *uint8) spi.Message {
	m.typeOrTagNumber = typeOrTagNumber
	m.lengthValueType = lengthValueType
	m.extTagNumber = extTagNumber
	m.extLength = extLength
	return m
}

func NewBACnetTagApplicationReal(value float32) BACnetTagInitializer {
	return &BACnetTagApplicationReal{value: value}
}

func (m BACnetTagApplicationReal) LengthInBits() uint16 {
	var lengthInBits uint16 = m.BACnetTag.LengthInBits()

	// Simple field (value)
	lengthInBits += 32

	return lengthInBits
}

func (m BACnetTagApplicationReal) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetTagApplicationRealParse(io spi.ReadBuffer, lengthValueType uint8, extLength uint8) (BACnetTagInitializer, error) {

	// Simple Field (value)
	var value float32 = io.ReadFloat32(32)

	// Create the instance
	return NewBACnetTagApplicationReal(value), nil
}

func (m BACnetTagApplicationReal) Serialize(io spi.WriteBuffer) {
	serializeFunc := func(typ interface{}) {
		if _, ok := typ.(IBACnetTagApplicationReal); ok {

			// Simple Field (value)
			var value float32 = m.value
			io.WriteFloat32(32, (value))
		}
	}
	serializeFunc(m)
}
