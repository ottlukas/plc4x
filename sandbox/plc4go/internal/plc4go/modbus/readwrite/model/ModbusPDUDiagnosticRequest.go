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
type ModbusPDUDiagnosticRequest struct {
	status     uint16
	eventCount uint16
	ModbusPDU
}

// The corresponding interface
type IModbusPDUDiagnosticRequest interface {
	IModbusPDU
	Serialize(io spi.WriteBuffer)
}

// Accessors for discriminator values.
func (m ModbusPDUDiagnosticRequest) ErrorFlag() bool {
	return false
}

func (m ModbusPDUDiagnosticRequest) FunctionFlag() uint8 {
	return 0x08
}

func (m ModbusPDUDiagnosticRequest) Response() bool {
	return false
}

func (m ModbusPDUDiagnosticRequest) initialize() spi.Message {
	return m
}

func NewModbusPDUDiagnosticRequest(status uint16, eventCount uint16) ModbusPDUInitializer {
	return &ModbusPDUDiagnosticRequest{status: status, eventCount: eventCount}
}

func (m ModbusPDUDiagnosticRequest) LengthInBits() uint16 {
	var lengthInBits uint16 = m.ModbusPDU.LengthInBits()

	// Simple field (status)
	lengthInBits += 16

	// Simple field (eventCount)
	lengthInBits += 16

	return lengthInBits
}

func (m ModbusPDUDiagnosticRequest) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ModbusPDUDiagnosticRequestParse(io spi.ReadBuffer) (ModbusPDUInitializer, error) {

	// Simple Field (status)
	var status uint16 = io.ReadUint16(16)

	// Simple Field (eventCount)
	var eventCount uint16 = io.ReadUint16(16)

	// Create the instance
	return NewModbusPDUDiagnosticRequest(status, eventCount), nil
}

func (m ModbusPDUDiagnosticRequest) Serialize(io spi.WriteBuffer) {
	serializeFunc := func(typ interface{}) {
		if _, ok := typ.(IModbusPDUDiagnosticRequest); ok {

			// Simple Field (status)
			var status uint16 = m.status
			io.WriteUint16(16, (status))

			// Simple Field (eventCount)
			var eventCount uint16 = m.eventCount
			io.WriteUint16(16, (eventCount))
		}
	}
	serializeFunc(m)
}
