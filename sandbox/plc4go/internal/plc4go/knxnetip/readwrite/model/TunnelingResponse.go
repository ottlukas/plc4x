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
type TunnelingResponse struct {
	tunnelingResponseDataBlock TunnelingResponseDataBlock
	KNXNetIPMessage
}

// The corresponding interface
type ITunnelingResponse interface {
	IKNXNetIPMessage
	Serialize(io spi.WriteBuffer)
}

// Accessors for discriminator values.
func (m TunnelingResponse) MsgType() uint16 {
	return 0x0421
}

func (m TunnelingResponse) initialize() spi.Message {
	return m
}

func NewTunnelingResponse(tunnelingResponseDataBlock TunnelingResponseDataBlock) KNXNetIPMessageInitializer {
	return &TunnelingResponse{tunnelingResponseDataBlock: tunnelingResponseDataBlock}
}

func CastITunnelingResponse(structType interface{}) ITunnelingResponse {
	castFunc := func(typ interface{}) ITunnelingResponse {
		if iTunnelingResponse, ok := typ.(ITunnelingResponse); ok {
			return iTunnelingResponse
		}
		return nil
	}
	return castFunc(structType)
}

func CastTunnelingResponse(structType interface{}) TunnelingResponse {
	castFunc := func(typ interface{}) TunnelingResponse {
		if sTunnelingResponse, ok := typ.(TunnelingResponse); ok {
			return sTunnelingResponse
		}
		return TunnelingResponse{}
	}
	return castFunc(structType)
}

func (m TunnelingResponse) LengthInBits() uint16 {
	var lengthInBits uint16 = m.KNXNetIPMessage.LengthInBits()

	// Simple field (tunnelingResponseDataBlock)
	lengthInBits += m.tunnelingResponseDataBlock.LengthInBits()

	return lengthInBits
}

func (m TunnelingResponse) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func TunnelingResponseParse(io spi.ReadBuffer) (KNXNetIPMessageInitializer, error) {

	// Simple Field (tunnelingResponseDataBlock)
	_tunnelingResponseDataBlockMessage, _err := TunnelingResponseDataBlockParse(io)
	if _err != nil {
		return nil, errors.New("Error parsing simple field 'tunnelingResponseDataBlock'. " + _err.Error())
	}
	var tunnelingResponseDataBlock TunnelingResponseDataBlock
	tunnelingResponseDataBlock, _tunnelingResponseDataBlockOk := _tunnelingResponseDataBlockMessage.(TunnelingResponseDataBlock)
	if !_tunnelingResponseDataBlockOk {
		return nil, errors.New("Couldn't cast message of type " + reflect.TypeOf(_tunnelingResponseDataBlockMessage).Name() + " to TunnelingResponseDataBlock")
	}

	// Create the instance
	return NewTunnelingResponse(tunnelingResponseDataBlock), nil
}

func (m TunnelingResponse) Serialize(io spi.WriteBuffer) {

	// Simple Field (tunnelingResponseDataBlock)
	tunnelingResponseDataBlock := TunnelingResponseDataBlock(m.tunnelingResponseDataBlock)
	tunnelingResponseDataBlock.Serialize(io)
}
