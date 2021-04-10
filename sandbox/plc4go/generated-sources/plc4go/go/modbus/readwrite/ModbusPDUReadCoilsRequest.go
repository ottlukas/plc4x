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
	"math"
	"plc4x.apache.org/plc4go-modbus-driver/0.8.0/src/plc4go/spi"
)

type ModbusPDUReadCoilsRequest struct {
	startingAddress uint16
	quantity        uint16
	ModbusPDU
}

func (m ModbusPDUReadCoilsRequest) initialize() ModbusPDU {
	return m.ModbusPDU
}

func NewModbusPDUReadCoilsRequest(startingAddress uint16, quantity uint16) ModbusPDUInitializer {
	return &ModbusPDUReadCoilsRequest{startingAddress: startingAddress, quantity: quantity}
}

func (m ModbusPDUReadCoilsRequest) LengthInBits() uint16 {
	var lengthInBits uint16 = m.ModbusPDU.LengthInBits()

	// Simple field (startingAddress)
	lengthInBits += 16

	// Simple field (quantity)
	lengthInBits += 16

	return lengthInBits
}

func (m ModbusPDUReadCoilsRequest) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ModbusPDUReadCoilsRequestParse(io spi.ReadBuffer) ModbusPDUInitializer {
	var startPos = io.GetPos()
	var curPos uint16

	// Simple Field (startingAddress)
	var startingAddress uint16 = io.ReadUint16(16)

	// Simple Field (quantity)
	var quantity uint16 = io.ReadUint16(16)

	// Create the instance
	return NewModbusPDUReadCoilsRequest(startingAddress, quantity)
}
