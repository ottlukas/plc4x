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
type BVLCReadForeignDeviceTableAck struct {
	BVLC
}

// The corresponding interface
type IBVLCReadForeignDeviceTableAck interface {
	IBVLC
	Serialize(io spi.WriteBuffer) error
}

// Accessors for discriminator values.
func (m BVLCReadForeignDeviceTableAck) BvlcFunction() uint8 {
	return 0x07
}

func (m BVLCReadForeignDeviceTableAck) initialize() spi.Message {
	return m
}

func NewBVLCReadForeignDeviceTableAck() BVLCInitializer {
	return &BVLCReadForeignDeviceTableAck{}
}

func CastIBVLCReadForeignDeviceTableAck(structType interface{}) IBVLCReadForeignDeviceTableAck {
	castFunc := func(typ interface{}) IBVLCReadForeignDeviceTableAck {
		if iBVLCReadForeignDeviceTableAck, ok := typ.(IBVLCReadForeignDeviceTableAck); ok {
			return iBVLCReadForeignDeviceTableAck
		}
		return nil
	}
	return castFunc(structType)
}

func CastBVLCReadForeignDeviceTableAck(structType interface{}) BVLCReadForeignDeviceTableAck {
	castFunc := func(typ interface{}) BVLCReadForeignDeviceTableAck {
		if sBVLCReadForeignDeviceTableAck, ok := typ.(BVLCReadForeignDeviceTableAck); ok {
			return sBVLCReadForeignDeviceTableAck
		}
		return BVLCReadForeignDeviceTableAck{}
	}
	return castFunc(structType)
}

func (m BVLCReadForeignDeviceTableAck) LengthInBits() uint16 {
	var lengthInBits = m.BVLC.LengthInBits()

	return lengthInBits
}

func (m BVLCReadForeignDeviceTableAck) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BVLCReadForeignDeviceTableAckParse(io *spi.ReadBuffer) (BVLCInitializer, error) {

	// Create the instance
	return NewBVLCReadForeignDeviceTableAck(), nil
}

func (m BVLCReadForeignDeviceTableAck) Serialize(io spi.WriteBuffer) error {
	ser := func() error {

		return nil
	}
	return BVLCSerialize(io, m.BVLC, CastIBVLC(m), ser)
}
