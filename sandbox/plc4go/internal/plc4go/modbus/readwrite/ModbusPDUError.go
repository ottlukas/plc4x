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
package readwrite

import (
	"plc4x.apache.org/plc4go-modbus-driver/0.8.0/internal/plc4go/spi"
)

// The data-structure of this message
type ModbusPDUError struct {
	exceptionCode uint8
	ModbusPDU
}

// The corresponding interface
type IModbusPDUError interface {
	IModbusPDU
	Serialize(io spi.WriteBuffer)
}

// Accessors for discriminator values.
func (m ModbusPDUError) ErrorFlag() bool {
	return true
}

func (m ModbusPDUError) FunctionFlag() uint8 {
	return 0
}

func (m ModbusPDUError) Response() bool {
	return false
}

func (m ModbusPDUError) initialize() spi.Message {
	return spi.Message(m)
}

func NewModbusPDUError(exceptionCode uint8) ModbusPDUInitializer {
	return &ModbusPDUError{exceptionCode: exceptionCode}
}

func (m ModbusPDUError) LengthInBits() uint16 {
	var lengthInBits uint16 = m.ModbusPDU.LengthInBits()

	// Simple field (exceptionCode)
	lengthInBits += 8

	return lengthInBits
}

func (m ModbusPDUError) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ModbusPDUErrorParse(io spi.ReadBuffer) (ModbusPDUInitializer, error) {

	// Simple Field (exceptionCode)
	var exceptionCode uint8 = io.ReadUint8(8)

	// Create the instance
	return NewModbusPDUError(exceptionCode), nil
}

func (m ModbusPDUError) Serialize(io spi.WriteBuffer) {
	serializeFunc := func(typ interface{}) {
		if _, ok := typ.(IModbusPDU); ok {

			// Simple Field (exceptionCode)
			var exceptionCode uint8 = m.exceptionCode
			io.WriteUint8(8, (exceptionCode))
		}
	}
	serializeFunc(m)
}
