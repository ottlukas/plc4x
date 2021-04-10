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
	log "github.com/sirupsen/logrus"
	"plc4x.apache.org/plc4go-modbus-driver/0.8.0/internal/plc4go/spi"
	"reflect"
	"strconv"
)

// Constant values.
const TPKTPacket_PROTOCOLID uint8 = 0x03

// The data-structure of this message
type TPKTPacket struct {
	payload ICOTPPacket
}

// The corresponding interface
type ITPKTPacket interface {
	spi.Message
	Serialize(io spi.WriteBuffer)
}

func NewTPKTPacket(payload ICOTPPacket) spi.Message {
	return &TPKTPacket{payload: payload}
}

func CastITPKTPacket(structType interface{}) ITPKTPacket {
	castFunc := func(typ interface{}) ITPKTPacket {
		if iTPKTPacket, ok := typ.(ITPKTPacket); ok {
			return iTPKTPacket
		}
		return nil
	}
	return castFunc(structType)
}

func CastTPKTPacket(structType interface{}) TPKTPacket {
	castFunc := func(typ interface{}) TPKTPacket {
		if sTPKTPacket, ok := typ.(TPKTPacket); ok {
			return sTPKTPacket
		}
		return TPKTPacket{}
	}
	return castFunc(structType)
}

func (m TPKTPacket) LengthInBits() uint16 {
	var lengthInBits uint16 = 0

	// Const Field (protocolId)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 8

	// Implicit Field (len)
	lengthInBits += 16

	// Simple field (payload)
	lengthInBits += m.payload.LengthInBits()

	return lengthInBits
}

func (m TPKTPacket) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func TPKTPacketParse(io spi.ReadBuffer) (spi.Message, error) {

	// Const Field (protocolId)
	protocolId, _protocolIdErr := io.ReadUint8(8)
	if _protocolIdErr != nil {
		return nil, errors.New("Error parsing 'protocolId' field " + _protocolIdErr.Error())
	}
	if protocolId != TPKTPacket_PROTOCOLID {
		return nil, errors.New("Expected constant value " + strconv.Itoa(int(TPKTPacket_PROTOCOLID)) + " but got " + strconv.Itoa(int(protocolId)))
	}

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := io.ReadUint8(8)
		if _err != nil {
			return nil, errors.New("Error parsing 'reserved' field " + _err.Error())
		}
		if reserved != uint8(0x00) {
			log.WithFields(log.Fields{
				"expected value": uint8(0x00),
				"got value":      reserved,
			}).Info("Got unexpected response.")
		}
	}

	// Implicit Field (len) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	len, _lenErr := io.ReadUint16(16)
	if _lenErr != nil {
		return nil, errors.New("Error parsing 'len' field " + _lenErr.Error())
	}

	// Simple Field (payload)
	_payloadMessage, _err := COTPPacketParse(io, uint16(len)-uint16(uint16(4)))
	if _err != nil {
		return nil, errors.New("Error parsing simple field 'payload'. " + _err.Error())
	}
	var payload ICOTPPacket
	payload, _payloadOk := _payloadMessage.(ICOTPPacket)
	if !_payloadOk {
		return nil, errors.New("Couldn't cast message of type " + reflect.TypeOf(_payloadMessage).Name() + " to ICOTPPacket")
	}

	// Create the instance
	return NewTPKTPacket(payload), nil
}

func (m TPKTPacket) Serialize(io spi.WriteBuffer) {

	// Const Field (protocolId)
	io.WriteUint8(8, 0x03)

	// Reserved Field (reserved)
	io.WriteUint8(8, uint8(0x00))

	// Implicit Field (len) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	len := uint16(uint16(m.payload.LengthInBytes()) + uint16(uint16(4)))
	io.WriteUint16(16, (len))

	// Simple Field (payload)
	payload := ICOTPPacket(m.payload)
	payload.Serialize(io)
}
