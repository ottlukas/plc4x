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
type ModbusPDUWriteSingleCoilRequest struct {
	address uint16
	value   uint16
	ModbusPDU
}

// The corresponding interface
type IModbusPDUWriteSingleCoilRequest interface {
	IModbusPDU
	Serialize(io spi.WriteBuffer)
}

// Accessors for discriminator values.
func (m ModbusPDUWriteSingleCoilRequest) ErrorFlag() bool {
	return false
}

func (m ModbusPDUWriteSingleCoilRequest) FunctionFlag() uint8 {
	return 0x05
}

func (m ModbusPDUWriteSingleCoilRequest) Response() bool {
	return false
}

func (m ModbusPDUWriteSingleCoilRequest) initialize() spi.Message {
	return m
}

func NewModbusPDUWriteSingleCoilRequest(address uint16, value uint16) ModbusPDUInitializer {
	return &ModbusPDUWriteSingleCoilRequest{address: address, value: value}
}

func (m ModbusPDUWriteSingleCoilRequest) LengthInBits() uint16 {
	var lengthInBits uint16 = m.ModbusPDU.LengthInBits()

	// Simple field (address)
	lengthInBits += 16

	// Simple field (value)
	lengthInBits += 16

	return lengthInBits
}

func (m ModbusPDUWriteSingleCoilRequest) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ModbusPDUWriteSingleCoilRequestParse(io spi.ReadBuffer) (ModbusPDUInitializer, error) {

	// Simple Field (address)
	var address uint16 = io.ReadUint16(16)

	// Simple Field (value)
	var value uint16 = io.ReadUint16(16)

	// Create the instance
	return NewModbusPDUWriteSingleCoilRequest(address, value), nil
}

func (m ModbusPDUWriteSingleCoilRequest) Serialize(io spi.WriteBuffer) {
	serializeFunc := func(typ interface{}) {
		if _, ok := typ.(IModbusPDUWriteSingleCoilRequest); ok {

			// Simple Field (address)
			var address uint16 = m.address
			io.WriteUint16(16, (address))

			// Simple Field (value)
			var value uint16 = m.value
			io.WriteUint16(16, (value))
		}
	}
	serializeFunc(m)
}
