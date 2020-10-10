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
type S7ParameterReadVarResponse struct {
	numItems uint8
	S7Parameter
}

// The corresponding interface
type IS7ParameterReadVarResponse interface {
	IS7Parameter
	Serialize(io spi.WriteBuffer)
}

// Accessors for discriminator values.
func (m S7ParameterReadVarResponse) ParameterType() uint8 {
	return 0x04
}

func (m S7ParameterReadVarResponse) MessageType() uint8 {
	return 0x03
}

func (m S7ParameterReadVarResponse) initialize() spi.Message {
	return m
}

func NewS7ParameterReadVarResponse(numItems uint8) S7ParameterInitializer {
	return &S7ParameterReadVarResponse{numItems: numItems}
}

func (m S7ParameterReadVarResponse) LengthInBits() uint16 {
	var lengthInBits uint16 = m.S7Parameter.LengthInBits()

	// Simple field (numItems)
	lengthInBits += 8

	return lengthInBits
}

func (m S7ParameterReadVarResponse) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func S7ParameterReadVarResponseParse(io spi.ReadBuffer) (S7ParameterInitializer, error) {

	// Simple Field (numItems)
	var numItems uint8 = io.ReadUint8(8)

	// Create the instance
	return NewS7ParameterReadVarResponse(numItems), nil
}

func (m S7ParameterReadVarResponse) Serialize(io spi.WriteBuffer) {
	serializeFunc := func(typ interface{}) {
		if _, ok := typ.(IS7ParameterReadVarResponse); ok {

			// Simple Field (numItems)
			var numItems uint8 = m.numItems
			io.WriteUint8(8, (numItems))
		}
	}
	serializeFunc(m)
}
