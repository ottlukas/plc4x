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
type BACnetConfirmedServiceACKConfirmedPrivateTransfer struct {
	BACnetConfirmedServiceACK
}

// The corresponding interface
type IBACnetConfirmedServiceACKConfirmedPrivateTransfer interface {
	IBACnetConfirmedServiceACK
	Serialize(io spi.WriteBuffer) error
}

// Accessors for discriminator values.
func (m BACnetConfirmedServiceACKConfirmedPrivateTransfer) ServiceChoice() uint8 {
	return 0x12
}

func (m BACnetConfirmedServiceACKConfirmedPrivateTransfer) initialize() spi.Message {
	return m
}

func NewBACnetConfirmedServiceACKConfirmedPrivateTransfer() BACnetConfirmedServiceACKInitializer {
	return &BACnetConfirmedServiceACKConfirmedPrivateTransfer{}
}

func CastIBACnetConfirmedServiceACKConfirmedPrivateTransfer(structType interface{}) IBACnetConfirmedServiceACKConfirmedPrivateTransfer {
	castFunc := func(typ interface{}) IBACnetConfirmedServiceACKConfirmedPrivateTransfer {
		if iBACnetConfirmedServiceACKConfirmedPrivateTransfer, ok := typ.(IBACnetConfirmedServiceACKConfirmedPrivateTransfer); ok {
			return iBACnetConfirmedServiceACKConfirmedPrivateTransfer
		}
		return nil
	}
	return castFunc(structType)
}

func CastBACnetConfirmedServiceACKConfirmedPrivateTransfer(structType interface{}) BACnetConfirmedServiceACKConfirmedPrivateTransfer {
	castFunc := func(typ interface{}) BACnetConfirmedServiceACKConfirmedPrivateTransfer {
		if sBACnetConfirmedServiceACKConfirmedPrivateTransfer, ok := typ.(BACnetConfirmedServiceACKConfirmedPrivateTransfer); ok {
			return sBACnetConfirmedServiceACKConfirmedPrivateTransfer
		}
		return BACnetConfirmedServiceACKConfirmedPrivateTransfer{}
	}
	return castFunc(structType)
}

func (m BACnetConfirmedServiceACKConfirmedPrivateTransfer) LengthInBits() uint16 {
	var lengthInBits = m.BACnetConfirmedServiceACK.LengthInBits()

	return lengthInBits
}

func (m BACnetConfirmedServiceACKConfirmedPrivateTransfer) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetConfirmedServiceACKConfirmedPrivateTransferParse(io *spi.ReadBuffer) (BACnetConfirmedServiceACKInitializer, error) {

	// Create the instance
	return NewBACnetConfirmedServiceACKConfirmedPrivateTransfer(), nil
}

func (m BACnetConfirmedServiceACKConfirmedPrivateTransfer) Serialize(io spi.WriteBuffer) error {
	ser := func() error {

		return nil
	}
	return BACnetConfirmedServiceACKSerialize(io, m.BACnetConfirmedServiceACK, CastIBACnetConfirmedServiceACK(m), ser)
}
