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

type ModbusPDUReadFifoQueueRequest struct {
	fifoPointerAddress uint16
	ModbusPDU
}

func (m ModbusPDUReadFifoQueueRequest) initialize() spi.Message {
	return spi.Message(m)
}

func NewModbusPDUReadFifoQueueRequest(fifoPointerAddress uint16) ModbusPDUInitializer {
	return &ModbusPDUReadFifoQueueRequest{fifoPointerAddress: fifoPointerAddress}
}

func (m ModbusPDUReadFifoQueueRequest) LengthInBits() uint16 {
	var lengthInBits uint16 = m.ModbusPDU.LengthInBits()

	// Simple field (fifoPointerAddress)
	lengthInBits += 16

	return lengthInBits
}

func (m ModbusPDUReadFifoQueueRequest) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ModbusPDUReadFifoQueueRequestParse(io spi.ReadBuffer) (ModbusPDUInitializer, error) {

	// Simple Field (fifoPointerAddress)
	var fifoPointerAddress uint16 = io.ReadUint16(16)

	// Create the instance
	return NewModbusPDUReadFifoQueueRequest(fifoPointerAddress), nil
}
