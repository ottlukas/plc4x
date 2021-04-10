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
type RelativeTimestamp struct {
	timestamp uint16
}

// The corresponding interface
type IRelativeTimestamp interface {
	spi.Message
	Serialize(io spi.WriteBuffer)
}

func NewRelativeTimestamp(timestamp uint16) spi.Message {
	return &RelativeTimestamp{timestamp: timestamp}
}

func CastIRelativeTimestamp(structType interface{}) IRelativeTimestamp {
	castFunc := func(typ interface{}) IRelativeTimestamp {
		if iRelativeTimestamp, ok := typ.(IRelativeTimestamp); ok {
			return iRelativeTimestamp
		}
		return nil
	}
	return castFunc(structType)
}

func CastRelativeTimestamp(structType interface{}) RelativeTimestamp {
	castFunc := func(typ interface{}) RelativeTimestamp {
		if sRelativeTimestamp, ok := typ.(RelativeTimestamp); ok {
			return sRelativeTimestamp
		}
		return RelativeTimestamp{}
	}
	return castFunc(structType)
}

func (m RelativeTimestamp) LengthInBits() uint16 {
	var lengthInBits uint16 = 0

	// Simple field (timestamp)
	lengthInBits += 16

	return lengthInBits
}

func (m RelativeTimestamp) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func RelativeTimestampParse(io spi.ReadBuffer) (spi.Message, error) {

	// Simple Field (timestamp)
	var timestamp uint16 = io.ReadUint16(16)

	// Create the instance
	return NewRelativeTimestamp(timestamp), nil
}

func (m RelativeTimestamp) Serialize(io spi.WriteBuffer) {

	// Simple Field (timestamp)
	timestamp := uint16(m.timestamp)
	io.WriteUint16(16, (timestamp))
}
