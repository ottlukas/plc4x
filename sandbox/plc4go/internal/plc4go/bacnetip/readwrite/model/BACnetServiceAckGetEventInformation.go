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
type BACnetServiceAckGetEventInformation struct {
	BACnetServiceAck
}

// The corresponding interface
type IBACnetServiceAckGetEventInformation interface {
	IBACnetServiceAck
	Serialize(io spi.WriteBuffer)
}

// Accessors for discriminator values.
func (m BACnetServiceAckGetEventInformation) ServiceChoice() uint8 {
	return 0x1D
}

func (m BACnetServiceAckGetEventInformation) initialize() spi.Message {
	return m
}

func NewBACnetServiceAckGetEventInformation() BACnetServiceAckInitializer {
	return &BACnetServiceAckGetEventInformation{}
}

func CastIBACnetServiceAckGetEventInformation(structType interface{}) IBACnetServiceAckGetEventInformation {
	castFunc := func(typ interface{}) IBACnetServiceAckGetEventInformation {
		if iBACnetServiceAckGetEventInformation, ok := typ.(IBACnetServiceAckGetEventInformation); ok {
			return iBACnetServiceAckGetEventInformation
		}
		return nil
	}
	return castFunc(structType)
}

func CastBACnetServiceAckGetEventInformation(structType interface{}) BACnetServiceAckGetEventInformation {
	castFunc := func(typ interface{}) BACnetServiceAckGetEventInformation {
		if sBACnetServiceAckGetEventInformation, ok := typ.(BACnetServiceAckGetEventInformation); ok {
			return sBACnetServiceAckGetEventInformation
		}
		return BACnetServiceAckGetEventInformation{}
	}
	return castFunc(structType)
}

func (m BACnetServiceAckGetEventInformation) LengthInBits() uint16 {
	var lengthInBits uint16 = m.BACnetServiceAck.LengthInBits()

	return lengthInBits
}

func (m BACnetServiceAckGetEventInformation) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetServiceAckGetEventInformationParse(io *spi.ReadBuffer) (BACnetServiceAckInitializer, error) {

	// Create the instance
	return NewBACnetServiceAckGetEventInformation(), nil
}

func (m BACnetServiceAckGetEventInformation) Serialize(io spi.WriteBuffer) {
	ser := func() {

	}
	BACnetServiceAckSerialize(io, m.BACnetServiceAck, CastIBACnetServiceAck(m), ser)
}
