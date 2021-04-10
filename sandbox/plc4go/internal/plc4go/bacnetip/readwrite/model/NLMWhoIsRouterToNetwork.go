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
type NLMWhoIsRouterToNetwork struct {
	destinationNetworkAddress []uint16
	NLM
}

// The corresponding interface
type INLMWhoIsRouterToNetwork interface {
	INLM
	Serialize(io spi.WriteBuffer)
}

// Accessors for discriminator values.
func (m NLMWhoIsRouterToNetwork) MessageType() uint8 {
	return 0x0
}

func (m NLMWhoIsRouterToNetwork) initialize(vendorId *uint16) spi.Message {
	m.vendorId = vendorId
	return m
}

func NewNLMWhoIsRouterToNetwork(destinationNetworkAddress []uint16) NLMInitializer {
	return &NLMWhoIsRouterToNetwork{destinationNetworkAddress: destinationNetworkAddress}
}

func (m NLMWhoIsRouterToNetwork) LengthInBits() uint16 {
	var lengthInBits uint16 = m.NLM.LengthInBits()

	// Array field
	if len(m.destinationNetworkAddress) > 0 {
		lengthInBits += 16 * uint16(len(m.destinationNetworkAddress))
	}

	return lengthInBits
}

func (m NLMWhoIsRouterToNetwork) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func NLMWhoIsRouterToNetworkParse(io spi.ReadBuffer, apduLength uint16, messageType uint8) (NLMInitializer, error) {

	// Array field (destinationNetworkAddress)
	var destinationNetworkAddress []uint16
	// Length array
	_destinationNetworkAddressLength := uint16((apduLength) - (spi.InlineIf((((messageType) >= (128)) && ((messageType) <= (255))), uint16(3), uint16(1))))
	_destinationNetworkAddressEndPos := io.GetPos() + _destinationNetworkAddressLength
	for io.GetPos() < _destinationNetworkAddressEndPos {
		destinationNetworkAddress = append(destinationNetworkAddress, io.ReadUint16(16))
	}

	// Create the instance
	return NewNLMWhoIsRouterToNetwork(destinationNetworkAddress), nil
}

func (m NLMWhoIsRouterToNetwork) Serialize(io spi.WriteBuffer) {
	serializeFunc := func(typ interface{}) {
		if _, ok := typ.(INLMWhoIsRouterToNetwork); ok {

			// Array Field (destinationNetworkAddress)
			if m.destinationNetworkAddress != nil {
				for _, _element := range m.destinationNetworkAddress {
					io.WriteUint16(16, _element)
				}
			}
		}
	}
	serializeFunc(m)
}
