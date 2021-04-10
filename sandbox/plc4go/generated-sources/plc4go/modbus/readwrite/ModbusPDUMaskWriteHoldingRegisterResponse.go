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
	"plc4x.apache.org/plc4go-modbus-driver/0.8.0/src/plc4go/spi"
)

type ModbusPDUMaskWriteHoldingRegisterResponse struct {
	referenceAddress uint16
	andMask          uint16
	orMask           uint16
	ModbusPDU
}

type IModbusPDUMaskWriteHoldingRegisterResponse interface {
	IModbusPDU
	Serialize(io spi.WriteBuffer)
}

// Accessors for discriminator values.
func (m ModbusPDUMaskWriteHoldingRegisterResponse) ErrorFlag() bool {
	return false
}

func (m ModbusPDUMaskWriteHoldingRegisterResponse) FunctionFlag() uint8 {
	return 0x16
}

func (m ModbusPDUMaskWriteHoldingRegisterResponse) Response() bool {
	return true
}

func (m ModbusPDUMaskWriteHoldingRegisterResponse) initialize() spi.Message {
	return spi.Message(m)
}

func NewModbusPDUMaskWriteHoldingRegisterResponse(referenceAddress uint16, andMask uint16, orMask uint16) ModbusPDUInitializer {
	return &ModbusPDUMaskWriteHoldingRegisterResponse{referenceAddress: referenceAddress, andMask: andMask, orMask: orMask}
}

func (m ModbusPDUMaskWriteHoldingRegisterResponse) LengthInBits() uint16 {
	var lengthInBits uint16 = m.ModbusPDU.LengthInBits()

	// Simple field (referenceAddress)
	lengthInBits += 16

	// Simple field (andMask)
	lengthInBits += 16

	// Simple field (orMask)
	lengthInBits += 16

	return lengthInBits
}

func (m ModbusPDUMaskWriteHoldingRegisterResponse) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ModbusPDUMaskWriteHoldingRegisterResponseParse(io spi.ReadBuffer) (ModbusPDUInitializer, error) {

	// Simple Field (referenceAddress)
	var referenceAddress uint16 = io.ReadUint16(16)

	// Simple Field (andMask)
	var andMask uint16 = io.ReadUint16(16)

	// Simple Field (orMask)
	var orMask uint16 = io.ReadUint16(16)

	// Create the instance
	return NewModbusPDUMaskWriteHoldingRegisterResponse(referenceAddress, andMask, orMask), nil
}

func (m ModbusPDUMaskWriteHoldingRegisterResponse) Serialize(io spi.WriteBuffer) {

	// Simple Field (referenceAddress)
	var referenceAddress uint16 = m.referenceAddress
	io.WriteUint16(16, (referenceAddress))

	// Simple Field (andMask)
	var andMask uint16 = m.andMask
	io.WriteUint16(16, (andMask))

	// Simple Field (orMask)
	var orMask uint16 = m.orMask
	io.WriteUint16(16, (orMask))
}
