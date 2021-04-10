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

import "plc4x.apache.org/plc4go-modbus-driver/0.8.0/internal/plc4go/spi"

type DeviceGroup int8

type IDeviceGroup interface {
	spi.Message
	Serialize(io spi.WriteBuffer)
}

const (
	DeviceGroup_PG_OR_PC DeviceGroup = 0x01
	DeviceGroup_OS       DeviceGroup = 0x02
	DeviceGroup_OTHERS   DeviceGroup = 0x03
)

func CastDeviceGroup(structType interface{}) DeviceGroup {
	castFunc := func(typ interface{}) DeviceGroup {
		if sDeviceGroup, ok := typ.(DeviceGroup); ok {
			return sDeviceGroup
		}
		return 0
	}
	return castFunc(structType)
}

func (m DeviceGroup) LengthInBits() uint16 {
	return 8
}

func (m DeviceGroup) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func DeviceGroupParse(io *spi.ReadBuffer) (DeviceGroup, error) {
	// TODO: Implement ...
	return 0, nil
}

func (e DeviceGroup) Serialize(io spi.WriteBuffer) {
	// TODO: Implement ...
}
