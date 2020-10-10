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
type BVLCOriginalUnicastNPDU struct {
	npdu NPDU
	BVLC
}

// The corresponding interface
type IBVLCOriginalUnicastNPDU interface {
	IBVLC
	Serialize(io spi.WriteBuffer)
}

// Accessors for discriminator values.
func (m BVLCOriginalUnicastNPDU) BvlcFunction() uint8 {
	return 0x0A
}

func (m BVLCOriginalUnicastNPDU) initialize() spi.Message {
	return m
}

func NewBVLCOriginalUnicastNPDU(npdu NPDU) BVLCInitializer {
	return &BVLCOriginalUnicastNPDU{npdu: npdu}
}

func (m BVLCOriginalUnicastNPDU) LengthInBits() uint16 {
	var lengthInBits uint16 = m.BVLC.LengthInBits()

	// Simple field (npdu)
	lengthInBits += m.npdu.LengthInBits()

	return lengthInBits
}

func (m BVLCOriginalUnicastNPDU) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BVLCOriginalUnicastNPDUParse(io spi.ReadBuffer, bvlcLength uint16) (BVLCInitializer, error) {

	// Simple Field (npdu)
	_npduMessage, _err := NPDUParse(io, uint16((bvlcLength)-(4)))
	if _err != nil {
		return nil, errors.New("Error parsing simple field 'npdu'. " + _err.Error())
	}
	var npdu NPDU
	npdu, _npduOk := _npduMessage.(NPDU)
	if !_npduOk {
		return nil, errors.New("Couldn't cast message of type " + reflect.TypeOf(_npduMessage).Name() + " to NPDU")
	}

	// Create the instance
	return NewBVLCOriginalUnicastNPDU(npdu), nil
}

func (m BVLCOriginalUnicastNPDU) Serialize(io spi.WriteBuffer) {
	serializeFunc := func(typ interface{}) {
		if _, ok := typ.(IBVLCOriginalUnicastNPDU); ok {

			// Simple Field (npdu)
			var npdu NPDU = m.npdu
			npdu.Serialize(io)
		}
	}
	serializeFunc(m)
}
