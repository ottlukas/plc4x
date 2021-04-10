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
type BACnetServiceAckGetAlarmSummary struct {
	BACnetServiceAck
}

// The corresponding interface
type IBACnetServiceAckGetAlarmSummary interface {
	IBACnetServiceAck
	Serialize(io spi.WriteBuffer) error
}

// Accessors for discriminator values.
func (m BACnetServiceAckGetAlarmSummary) ServiceChoice() uint8 {
	return 0x03
}

func (m BACnetServiceAckGetAlarmSummary) initialize() spi.Message {
	return m
}

func NewBACnetServiceAckGetAlarmSummary() BACnetServiceAckInitializer {
	return &BACnetServiceAckGetAlarmSummary{}
}

func CastIBACnetServiceAckGetAlarmSummary(structType interface{}) IBACnetServiceAckGetAlarmSummary {
	castFunc := func(typ interface{}) IBACnetServiceAckGetAlarmSummary {
		if iBACnetServiceAckGetAlarmSummary, ok := typ.(IBACnetServiceAckGetAlarmSummary); ok {
			return iBACnetServiceAckGetAlarmSummary
		}
		return nil
	}
	return castFunc(structType)
}

func CastBACnetServiceAckGetAlarmSummary(structType interface{}) BACnetServiceAckGetAlarmSummary {
	castFunc := func(typ interface{}) BACnetServiceAckGetAlarmSummary {
		if sBACnetServiceAckGetAlarmSummary, ok := typ.(BACnetServiceAckGetAlarmSummary); ok {
			return sBACnetServiceAckGetAlarmSummary
		}
		return BACnetServiceAckGetAlarmSummary{}
	}
	return castFunc(structType)
}

func (m BACnetServiceAckGetAlarmSummary) LengthInBits() uint16 {
	var lengthInBits = m.BACnetServiceAck.LengthInBits()

	return lengthInBits
}

func (m BACnetServiceAckGetAlarmSummary) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetServiceAckGetAlarmSummaryParse(io *spi.ReadBuffer) (BACnetServiceAckInitializer, error) {

	// Create the instance
	return NewBACnetServiceAckGetAlarmSummary(), nil
}

func (m BACnetServiceAckGetAlarmSummary) Serialize(io spi.WriteBuffer) error {
	ser := func() error {

		return nil
	}
	return BACnetServiceAckSerialize(io, m.BACnetServiceAck, CastIBACnetServiceAck(m), ser)
}
