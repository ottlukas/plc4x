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
type HPAIDiscoveryEndpoint struct {
	hostProtocolCode IHostProtocolCode
	ipAddress        IIPAddress
	ipPort           uint16
}

// The corresponding interface
type IHPAIDiscoveryEndpoint interface {
	spi.Message
	Serialize(io spi.WriteBuffer)
}

func NewHPAIDiscoveryEndpoint(hostProtocolCode IHostProtocolCode, ipAddress IIPAddress, ipPort uint16) spi.Message {
	return &HPAIDiscoveryEndpoint{hostProtocolCode: hostProtocolCode, ipAddress: ipAddress, ipPort: ipPort}
}

func CastIHPAIDiscoveryEndpoint(structType interface{}) IHPAIDiscoveryEndpoint {
	castFunc := func(typ interface{}) IHPAIDiscoveryEndpoint {
		if iHPAIDiscoveryEndpoint, ok := typ.(IHPAIDiscoveryEndpoint); ok {
			return iHPAIDiscoveryEndpoint
		}
		return nil
	}
	return castFunc(structType)
}

func CastHPAIDiscoveryEndpoint(structType interface{}) HPAIDiscoveryEndpoint {
	castFunc := func(typ interface{}) HPAIDiscoveryEndpoint {
		if sHPAIDiscoveryEndpoint, ok := typ.(HPAIDiscoveryEndpoint); ok {
			return sHPAIDiscoveryEndpoint
		}
		return HPAIDiscoveryEndpoint{}
	}
	return castFunc(structType)
}

func (m HPAIDiscoveryEndpoint) LengthInBits() uint16 {
	var lengthInBits uint16 = 0

	// Implicit Field (structureLength)
	lengthInBits += 8

	// Enum Field (hostProtocolCode)
	lengthInBits += 8

	// Simple field (ipAddress)
	lengthInBits += m.ipAddress.LengthInBits()

	// Simple field (ipPort)
	lengthInBits += 16

	return lengthInBits
}

func (m HPAIDiscoveryEndpoint) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func HPAIDiscoveryEndpointParse(io spi.ReadBuffer) (spi.Message, error) {

	// Implicit Field (structureLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	_, _structureLengthErr := io.ReadUint8(8)
	if _structureLengthErr != nil {
		return nil, errors.New("Error parsing 'structureLength' field " + _structureLengthErr.Error())
	}

	// Enum field (hostProtocolCode)
	hostProtocolCode, _hostProtocolCodeErr := HostProtocolCodeParse(io)
	if _hostProtocolCodeErr != nil {
		return nil, errors.New("Error parsing 'hostProtocolCode' field " + _hostProtocolCodeErr.Error())
	}

	// Simple Field (ipAddress)
	_ipAddressMessage, _err := IPAddressParse(io)
	if _err != nil {
		return nil, errors.New("Error parsing simple field 'ipAddress'. " + _err.Error())
	}
	var ipAddress IIPAddress
	ipAddress, _ipAddressOk := _ipAddressMessage.(IIPAddress)
	if !_ipAddressOk {
		return nil, errors.New("Couldn't cast message of type " + reflect.TypeOf(_ipAddressMessage).Name() + " to IIPAddress")
	}

	// Simple Field (ipPort)
	ipPort, _ipPortErr := io.ReadUint16(16)
	if _ipPortErr != nil {
		return nil, errors.New("Error parsing 'ipPort' field " + _ipPortErr.Error())
	}

	// Create the instance
	return NewHPAIDiscoveryEndpoint(hostProtocolCode, ipAddress, ipPort), nil
}

func (m HPAIDiscoveryEndpoint) Serialize(io spi.WriteBuffer) {

	// Implicit Field (structureLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	structureLength := uint8(uint8(m.LengthInBytes()))
	io.WriteUint8(8, (structureLength))

	// Enum field (hostProtocolCode)
	hostProtocolCode := IHostProtocolCode(m.hostProtocolCode)
	hostProtocolCode.Serialize(io)

	// Simple Field (ipAddress)
	ipAddress := IIPAddress(m.ipAddress)
	ipAddress.Serialize(io)

	// Simple Field (ipPort)
	ipPort := uint16(m.ipPort)
	io.WriteUint16(16, (ipPort))
}
