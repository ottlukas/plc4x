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
type BACnetErrorGetAlarmSummary struct {
	BACnetError
}

// The corresponding interface
type IBACnetErrorGetAlarmSummary interface {
	IBACnetError
	Serialize(io spi.WriteBuffer) error
}

// Accessors for discriminator values.
func (m BACnetErrorGetAlarmSummary) ServiceChoice() uint8 {
	return 0x03
}

func (m BACnetErrorGetAlarmSummary) initialize() spi.Message {
	return m
}

func NewBACnetErrorGetAlarmSummary() BACnetErrorInitializer {
	return &BACnetErrorGetAlarmSummary{}
}

func CastIBACnetErrorGetAlarmSummary(structType interface{}) IBACnetErrorGetAlarmSummary {
	castFunc := func(typ interface{}) IBACnetErrorGetAlarmSummary {
		if iBACnetErrorGetAlarmSummary, ok := typ.(IBACnetErrorGetAlarmSummary); ok {
			return iBACnetErrorGetAlarmSummary
		}
		return nil
	}
	return castFunc(structType)
}

func CastBACnetErrorGetAlarmSummary(structType interface{}) BACnetErrorGetAlarmSummary {
	castFunc := func(typ interface{}) BACnetErrorGetAlarmSummary {
		if sBACnetErrorGetAlarmSummary, ok := typ.(BACnetErrorGetAlarmSummary); ok {
			return sBACnetErrorGetAlarmSummary
		}
		return BACnetErrorGetAlarmSummary{}
	}
	return castFunc(structType)
}

func (m BACnetErrorGetAlarmSummary) LengthInBits() uint16 {
	var lengthInBits uint16 = m.BACnetError.LengthInBits()

	return lengthInBits
}

func (m BACnetErrorGetAlarmSummary) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetErrorGetAlarmSummaryParse(io *spi.ReadBuffer) (BACnetErrorInitializer, error) {

	// Create the instance
	return NewBACnetErrorGetAlarmSummary(), nil
}

func (m BACnetErrorGetAlarmSummary) Serialize(io spi.WriteBuffer) error {
	ser := func() error {

		return nil
	}
	return BACnetErrorSerialize(io, m.BACnetError, CastIBACnetError(m), ser)
}
