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
	"reflect"
)

// The data-structure of this message
type BVLCOriginalBroadcastNPDU struct {
	npdu INPDU
	BVLC
}

// The corresponding interface
type IBVLCOriginalBroadcastNPDU interface {
	IBVLC
	Serialize(io spi.WriteBuffer)
}

// Accessors for discriminator values.
func (m BVLCOriginalBroadcastNPDU) BvlcFunction() uint8 {
	return 0x0B
}

func (m BVLCOriginalBroadcastNPDU) initialize() spi.Message {
	return m
}

func NewBVLCOriginalBroadcastNPDU(npdu INPDU) BVLCInitializer {
	return &BVLCOriginalBroadcastNPDU{npdu: npdu}
}

func CastIBVLCOriginalBroadcastNPDU(structType interface{}) IBVLCOriginalBroadcastNPDU {
	castFunc := func(typ interface{}) IBVLCOriginalBroadcastNPDU {
		if iBVLCOriginalBroadcastNPDU, ok := typ.(IBVLCOriginalBroadcastNPDU); ok {
			return iBVLCOriginalBroadcastNPDU
		}
		return nil
	}
	return castFunc(structType)
}

func CastBVLCOriginalBroadcastNPDU(structType interface{}) BVLCOriginalBroadcastNPDU {
	castFunc := func(typ interface{}) BVLCOriginalBroadcastNPDU {
		if sBVLCOriginalBroadcastNPDU, ok := typ.(BVLCOriginalBroadcastNPDU); ok {
			return sBVLCOriginalBroadcastNPDU
		}
		return BVLCOriginalBroadcastNPDU{}
	}
	return castFunc(structType)
}

func (m BVLCOriginalBroadcastNPDU) LengthInBits() uint16 {
	var lengthInBits uint16 = m.BVLC.LengthInBits()

	// Simple field (npdu)
	lengthInBits += m.npdu.LengthInBits()

	return lengthInBits
}

func (m BVLCOriginalBroadcastNPDU) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BVLCOriginalBroadcastNPDUParse(io spi.ReadBuffer, bvlcLength uint16) (BVLCInitializer, error) {

	// Simple Field (npdu)
	_npduMessage, _err := NPDUParse(io, uint16(bvlcLength)-uint16(uint16(4)))
	if _err != nil {
		return nil, errors.New("Error parsing simple field 'npdu'. " + _err.Error())
	}
	var npdu INPDU
	npdu, _npduOk := _npduMessage.(INPDU)
	if !_npduOk {
		return nil, errors.New("Couldn't cast message of type " + reflect.TypeOf(_npduMessage).Name() + " to INPDU")
	}

	// Create the instance
	return NewBVLCOriginalBroadcastNPDU(npdu), nil
}

func (m BVLCOriginalBroadcastNPDU) Serialize(io spi.WriteBuffer) {

	// Simple Field (npdu)
	npdu := INPDU(m.npdu)
	npdu.Serialize(io)
}
