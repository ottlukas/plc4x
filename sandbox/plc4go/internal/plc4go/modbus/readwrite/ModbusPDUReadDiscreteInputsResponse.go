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
type ModbusPDUReadDiscreteInputsResponse struct {
	value []int8
	ModbusPDU
}

// The corresponding interface
type IModbusPDUReadDiscreteInputsResponse interface {
	IModbusPDU
	Serialize(io spi.WriteBuffer)
}

// Accessors for discriminator values.
func (m ModbusPDUReadDiscreteInputsResponse) ErrorFlag() bool {
	return false
}

func (m ModbusPDUReadDiscreteInputsResponse) FunctionFlag() uint8 {
	return 0x02
}

func (m ModbusPDUReadDiscreteInputsResponse) Response() bool {
	return true
}

func (m ModbusPDUReadDiscreteInputsResponse) initialize() spi.Message {
	return spi.Message(m)
}

func NewModbusPDUReadDiscreteInputsResponse(value []int8) ModbusPDUInitializer {
	return &ModbusPDUReadDiscreteInputsResponse{value: value}
}

func (m ModbusPDUReadDiscreteInputsResponse) LengthInBits() uint16 {
	var lengthInBits uint16 = m.ModbusPDU.LengthInBits()

	// Implicit Field (byteCount)
	lengthInBits += 8

	// Array field
	if len(m.value) > 0 {
		lengthInBits += 8 * uint16(len(m.value))
	}

	return lengthInBits
}

func (m ModbusPDUReadDiscreteInputsResponse) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ModbusPDUReadDiscreteInputsResponseParse(io spi.ReadBuffer) (ModbusPDUInitializer, error) {

	// Implicit Field (byteCount) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	var byteCount uint8 = io.ReadUint8(8)

	// Array field (value)
	var value []int8
	// Count array
	{
		value := make([]int8, byteCount)
		for curItem := uint16(0); curItem < uint16(byteCount); curItem++ {

			value[curItem] = io.ReadInt8(8)
		}
	}

	// Create the instance
	return NewModbusPDUReadDiscreteInputsResponse(value), nil
}

func (m ModbusPDUReadDiscreteInputsResponse) Serialize(io spi.WriteBuffer) {
	serializeFunc := func(typ interface{}) {
		if _, ok := typ.(IModbusPDU); ok {

			// Implicit Field (byteCount) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
			var byteCount uint8 = (uint8(len(m.value)))
			io.WriteUint8(8, (byteCount))

			// Array Field (value)
			if m.value != nil {
				for _, _element := range m.value {
					io.WriteInt8(8, _element)
				}
			}
		}
	}
	serializeFunc(m)
}
