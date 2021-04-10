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
type BACnetConfirmedServiceRequestAcknowledgeAlarm struct {
	BACnetConfirmedServiceRequest
}

// The corresponding interface
type IBACnetConfirmedServiceRequestAcknowledgeAlarm interface {
	IBACnetConfirmedServiceRequest
	Serialize(io spi.WriteBuffer)
}

// Accessors for discriminator values.
func (m BACnetConfirmedServiceRequestAcknowledgeAlarm) ServiceChoice() uint8 {
	return 0x00
}

func (m BACnetConfirmedServiceRequestAcknowledgeAlarm) initialize() spi.Message {
	return m
}

func NewBACnetConfirmedServiceRequestAcknowledgeAlarm() BACnetConfirmedServiceRequestInitializer {
	return &BACnetConfirmedServiceRequestAcknowledgeAlarm{}
}

func CastIBACnetConfirmedServiceRequestAcknowledgeAlarm(structType interface{}) IBACnetConfirmedServiceRequestAcknowledgeAlarm {
	castFunc := func(typ interface{}) IBACnetConfirmedServiceRequestAcknowledgeAlarm {
		if iBACnetConfirmedServiceRequestAcknowledgeAlarm, ok := typ.(IBACnetConfirmedServiceRequestAcknowledgeAlarm); ok {
			return iBACnetConfirmedServiceRequestAcknowledgeAlarm
		}
		return nil
	}
	return castFunc(structType)
}

func CastBACnetConfirmedServiceRequestAcknowledgeAlarm(structType interface{}) BACnetConfirmedServiceRequestAcknowledgeAlarm {
	castFunc := func(typ interface{}) BACnetConfirmedServiceRequestAcknowledgeAlarm {
		if sBACnetConfirmedServiceRequestAcknowledgeAlarm, ok := typ.(BACnetConfirmedServiceRequestAcknowledgeAlarm); ok {
			return sBACnetConfirmedServiceRequestAcknowledgeAlarm
		}
		return BACnetConfirmedServiceRequestAcknowledgeAlarm{}
	}
	return castFunc(structType)
}

func (m BACnetConfirmedServiceRequestAcknowledgeAlarm) LengthInBits() uint16 {
	var lengthInBits uint16 = m.BACnetConfirmedServiceRequest.LengthInBits()

	return lengthInBits
}

func (m BACnetConfirmedServiceRequestAcknowledgeAlarm) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetConfirmedServiceRequestAcknowledgeAlarmParse(io spi.ReadBuffer) (BACnetConfirmedServiceRequestInitializer, error) {

	// Create the instance
	return NewBACnetConfirmedServiceRequestAcknowledgeAlarm(), nil
}

func (m BACnetConfirmedServiceRequestAcknowledgeAlarm) Serialize(io spi.WriteBuffer) {
	ser := func() {

	}
	BACnetConfirmedServiceRequestSerialize(io, m.BACnetConfirmedServiceRequest, CastIBACnetConfirmedServiceRequest(m), ser)
}
