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
)

// The data-structure of this message
type ModbusPDUReadInputRegistersRequest struct {
	startingAddress uint16
	quantity        uint16
	ModbusPDU
}

// The corresponding interface
type IModbusPDUReadInputRegistersRequest interface {
	IModbusPDU
	Serialize(io spi.WriteBuffer)
}

// Accessors for discriminator values.
func (m ModbusPDUReadInputRegistersRequest) ErrorFlag() bool {
	return false
}

func (m ModbusPDUReadInputRegistersRequest) FunctionFlag() uint8 {
	return 0x04
}

func (m ModbusPDUReadInputRegistersRequest) Response() bool {
	return false
}

func (m ModbusPDUReadInputRegistersRequest) initialize() spi.Message {
	return m
}

func NewModbusPDUReadInputRegistersRequest(startingAddress uint16, quantity uint16) ModbusPDUInitializer {
	return &ModbusPDUReadInputRegistersRequest{startingAddress: startingAddress, quantity: quantity}
}

func CastIModbusPDUReadInputRegistersRequest(structType interface{}) IModbusPDUReadInputRegistersRequest {
	castFunc := func(typ interface{}) IModbusPDUReadInputRegistersRequest {
		if iModbusPDUReadInputRegistersRequest, ok := typ.(IModbusPDUReadInputRegistersRequest); ok {
			return iModbusPDUReadInputRegistersRequest
		}
		return nil
	}
	return castFunc(structType)
}

func CastModbusPDUReadInputRegistersRequest(structType interface{}) ModbusPDUReadInputRegistersRequest {
	castFunc := func(typ interface{}) ModbusPDUReadInputRegistersRequest {
		if sModbusPDUReadInputRegistersRequest, ok := typ.(ModbusPDUReadInputRegistersRequest); ok {
			return sModbusPDUReadInputRegistersRequest
		}
		return ModbusPDUReadInputRegistersRequest{}
	}
	return castFunc(structType)
}

func (m ModbusPDUReadInputRegistersRequest) LengthInBits() uint16 {
	var lengthInBits uint16 = m.ModbusPDU.LengthInBits()

	// Simple field (startingAddress)
	lengthInBits += 16

	// Simple field (quantity)
	lengthInBits += 16

	return lengthInBits
}

func (m ModbusPDUReadInputRegistersRequest) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ModbusPDUReadInputRegistersRequestParse(io spi.ReadBuffer) (ModbusPDUInitializer, error) {

	// Simple Field (startingAddress)
	startingAddress, _startingAddressErr := io.ReadUint16(16)
	if _startingAddressErr != nil {
		return nil, errors.New("Error parsing 'startingAddress' field " + _startingAddressErr.Error())
	}

	// Simple Field (quantity)
	quantity, _quantityErr := io.ReadUint16(16)
	if _quantityErr != nil {
		return nil, errors.New("Error parsing 'quantity' field " + _quantityErr.Error())
	}

	// Create the instance
	return NewModbusPDUReadInputRegistersRequest(startingAddress, quantity), nil
}

func (m ModbusPDUReadInputRegistersRequest) Serialize(io spi.WriteBuffer) {
	ser := func() {

		// Simple Field (startingAddress)
		startingAddress := uint16(m.startingAddress)
		io.WriteUint16(16, (startingAddress))

		// Simple Field (quantity)
		quantity := uint16(m.quantity)
		io.WriteUint16(16, (quantity))

	}
	ModbusPDUSerialize(io, m.ModbusPDU, CastIModbusPDU(m), ser)
}
