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
type BACnetConfirmedServiceRequestVTData struct {
	BACnetConfirmedServiceRequest
}

// The corresponding interface
type IBACnetConfirmedServiceRequestVTData interface {
	IBACnetConfirmedServiceRequest
	Serialize(io spi.WriteBuffer) error
}

// Accessors for discriminator values.
func (m BACnetConfirmedServiceRequestVTData) ServiceChoice() uint8 {
	return 0x17
}

func (m BACnetConfirmedServiceRequestVTData) initialize() spi.Message {
	return m
}

func NewBACnetConfirmedServiceRequestVTData() BACnetConfirmedServiceRequestInitializer {
	return &BACnetConfirmedServiceRequestVTData{}
}

func CastIBACnetConfirmedServiceRequestVTData(structType interface{}) IBACnetConfirmedServiceRequestVTData {
	castFunc := func(typ interface{}) IBACnetConfirmedServiceRequestVTData {
		if iBACnetConfirmedServiceRequestVTData, ok := typ.(IBACnetConfirmedServiceRequestVTData); ok {
			return iBACnetConfirmedServiceRequestVTData
		}
		return nil
	}
	return castFunc(structType)
}

func CastBACnetConfirmedServiceRequestVTData(structType interface{}) BACnetConfirmedServiceRequestVTData {
	castFunc := func(typ interface{}) BACnetConfirmedServiceRequestVTData {
		if sBACnetConfirmedServiceRequestVTData, ok := typ.(BACnetConfirmedServiceRequestVTData); ok {
			return sBACnetConfirmedServiceRequestVTData
		}
		return BACnetConfirmedServiceRequestVTData{}
	}
	return castFunc(structType)
}

func (m BACnetConfirmedServiceRequestVTData) LengthInBits() uint16 {
	var lengthInBits = m.BACnetConfirmedServiceRequest.LengthInBits()

	return lengthInBits
}

func (m BACnetConfirmedServiceRequestVTData) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetConfirmedServiceRequestVTDataParse(io *spi.ReadBuffer) (BACnetConfirmedServiceRequestInitializer, error) {

	// Create the instance
	return NewBACnetConfirmedServiceRequestVTData(), nil
}

func (m BACnetConfirmedServiceRequestVTData) Serialize(io spi.WriteBuffer) error {
	ser := func() error {

		return nil
	}
	return BACnetConfirmedServiceRequestSerialize(io, m.BACnetConfirmedServiceRequest, CastIBACnetConfirmedServiceRequest(m), ser)
}
