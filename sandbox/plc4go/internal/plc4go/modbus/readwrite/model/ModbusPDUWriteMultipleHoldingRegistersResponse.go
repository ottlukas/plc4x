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
type ModbusPDUWriteMultipleHoldingRegistersResponse struct {
	startingAddress uint16
	quantity        uint16
	ModbusPDU
}

// The corresponding interface
type IModbusPDUWriteMultipleHoldingRegistersResponse interface {
	IModbusPDU
	Serialize(io spi.WriteBuffer) error
}

// Accessors for discriminator values.
func (m ModbusPDUWriteMultipleHoldingRegistersResponse) ErrorFlag() bool {
	return false
}

func (m ModbusPDUWriteMultipleHoldingRegistersResponse) FunctionFlag() uint8 {
	return 0x10
}

func (m ModbusPDUWriteMultipleHoldingRegistersResponse) Response() bool {
	return true
}

func (m ModbusPDUWriteMultipleHoldingRegistersResponse) initialize() spi.Message {
	return m
}

func NewModbusPDUWriteMultipleHoldingRegistersResponse(startingAddress uint16, quantity uint16) ModbusPDUInitializer {
	return &ModbusPDUWriteMultipleHoldingRegistersResponse{startingAddress: startingAddress, quantity: quantity}
}

func CastIModbusPDUWriteMultipleHoldingRegistersResponse(structType interface{}) IModbusPDUWriteMultipleHoldingRegistersResponse {
	castFunc := func(typ interface{}) IModbusPDUWriteMultipleHoldingRegistersResponse {
		if iModbusPDUWriteMultipleHoldingRegistersResponse, ok := typ.(IModbusPDUWriteMultipleHoldingRegistersResponse); ok {
			return iModbusPDUWriteMultipleHoldingRegistersResponse
		}
		return nil
	}
	return castFunc(structType)
}

func CastModbusPDUWriteMultipleHoldingRegistersResponse(structType interface{}) ModbusPDUWriteMultipleHoldingRegistersResponse {
	castFunc := func(typ interface{}) ModbusPDUWriteMultipleHoldingRegistersResponse {
		if sModbusPDUWriteMultipleHoldingRegistersResponse, ok := typ.(ModbusPDUWriteMultipleHoldingRegistersResponse); ok {
			return sModbusPDUWriteMultipleHoldingRegistersResponse
		}
		return ModbusPDUWriteMultipleHoldingRegistersResponse{}
	}
	return castFunc(structType)
}

func (m ModbusPDUWriteMultipleHoldingRegistersResponse) LengthInBits() uint16 {
	var lengthInBits uint16 = m.ModbusPDU.LengthInBits()

	// Simple field (startingAddress)
	lengthInBits += 16

	// Simple field (quantity)
	lengthInBits += 16

	return lengthInBits
}

func (m ModbusPDUWriteMultipleHoldingRegistersResponse) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ModbusPDUWriteMultipleHoldingRegistersResponseParse(io *spi.ReadBuffer) (ModbusPDUInitializer, error) {

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
	return NewModbusPDUWriteMultipleHoldingRegistersResponse(startingAddress, quantity), nil
}

func (m ModbusPDUWriteMultipleHoldingRegistersResponse) Serialize(io spi.WriteBuffer) error {
	ser := func() error {

		// Simple Field (startingAddress)
		startingAddress := uint16(m.startingAddress)
		_startingAddressErr := io.WriteUint16(16, (startingAddress))
		if _startingAddressErr != nil {
			return errors.New("Error serializing 'startingAddress' field " + _startingAddressErr.Error())
		}

		// Simple Field (quantity)
		quantity := uint16(m.quantity)
		_quantityErr := io.WriteUint16(16, (quantity))
		if _quantityErr != nil {
			return errors.New("Error serializing 'quantity' field " + _quantityErr.Error())
		}

		return nil
	}
	return ModbusPDUSerialize(io, m.ModbusPDU, CastIModbusPDU(m), ser)
}
