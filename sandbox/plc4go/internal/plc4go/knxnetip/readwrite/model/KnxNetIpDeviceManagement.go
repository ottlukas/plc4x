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
)

// The data-structure of this message
type KnxNetIpDeviceManagement struct {
	version uint8
	ServiceId
}

// The corresponding interface
type IKnxNetIpDeviceManagement interface {
	IServiceId
	Serialize(io spi.WriteBuffer)
}

// Accessors for discriminator values.
func (m KnxNetIpDeviceManagement) ServiceType() uint8 {
	return 0x03
}

func (m KnxNetIpDeviceManagement) initialize() spi.Message {
	return m
}

func NewKnxNetIpDeviceManagement(version uint8) ServiceIdInitializer {
	return &KnxNetIpDeviceManagement{version: version}
}

func CastIKnxNetIpDeviceManagement(structType interface{}) IKnxNetIpDeviceManagement {
	castFunc := func(typ interface{}) IKnxNetIpDeviceManagement {
		if iKnxNetIpDeviceManagement, ok := typ.(IKnxNetIpDeviceManagement); ok {
			return iKnxNetIpDeviceManagement
		}
		return nil
	}
	return castFunc(structType)
}

func CastKnxNetIpDeviceManagement(structType interface{}) KnxNetIpDeviceManagement {
	castFunc := func(typ interface{}) KnxNetIpDeviceManagement {
		if sKnxNetIpDeviceManagement, ok := typ.(KnxNetIpDeviceManagement); ok {
			return sKnxNetIpDeviceManagement
		}
		return KnxNetIpDeviceManagement{}
	}
	return castFunc(structType)
}

func (m KnxNetIpDeviceManagement) LengthInBits() uint16 {
	var lengthInBits uint16 = m.ServiceId.LengthInBits()

	// Simple field (version)
	lengthInBits += 8

	return lengthInBits
}

func (m KnxNetIpDeviceManagement) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func KnxNetIpDeviceManagementParse(io *spi.ReadBuffer) (ServiceIdInitializer, error) {

	// Simple Field (version)
	version, _versionErr := io.ReadUint8(8)
	if _versionErr != nil {
		return nil, errors.New("Error parsing 'version' field " + _versionErr.Error())
	}

	// Create the instance
	return NewKnxNetIpDeviceManagement(version), nil
}

func (m KnxNetIpDeviceManagement) Serialize(io spi.WriteBuffer) {
	ser := func() {

		// Simple Field (version)
		version := uint8(m.version)
		io.WriteUint8(8, (version))

	}
	ServiceIdSerialize(io, m.ServiceId, CastIServiceId(m), ser)
}
