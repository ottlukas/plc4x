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
type KnxNetRemoteConfigurationAndDiagnosis struct {
	version uint8
	ServiceId
}

// The corresponding interface
type IKnxNetRemoteConfigurationAndDiagnosis interface {
	IServiceId
	Serialize(io spi.WriteBuffer)
}

// Accessors for discriminator values.
func (m KnxNetRemoteConfigurationAndDiagnosis) ServiceType() uint8 {
	return 0x07
}

func (m KnxNetRemoteConfigurationAndDiagnosis) initialize() spi.Message {
	return m
}

func NewKnxNetRemoteConfigurationAndDiagnosis(version uint8) ServiceIdInitializer {
	return &KnxNetRemoteConfigurationAndDiagnosis{version: version}
}

func (m KnxNetRemoteConfigurationAndDiagnosis) LengthInBits() uint16 {
	var lengthInBits uint16 = m.ServiceId.LengthInBits()

	// Simple field (version)
	lengthInBits += 8

	return lengthInBits
}

func (m KnxNetRemoteConfigurationAndDiagnosis) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func KnxNetRemoteConfigurationAndDiagnosisParse(io spi.ReadBuffer) (ServiceIdInitializer, error) {

	// Simple Field (version)
	var version uint8 = io.ReadUint8(8)

	// Create the instance
	return NewKnxNetRemoteConfigurationAndDiagnosis(version), nil
}

func (m KnxNetRemoteConfigurationAndDiagnosis) Serialize(io spi.WriteBuffer) {
	serializeFunc := func(typ interface{}) {
		if _, ok := typ.(IKnxNetRemoteConfigurationAndDiagnosis); ok {

			// Simple Field (version)
			var version uint8 = m.version
			io.WriteUint8(8, (version))
		}
	}
	serializeFunc(m)
}
