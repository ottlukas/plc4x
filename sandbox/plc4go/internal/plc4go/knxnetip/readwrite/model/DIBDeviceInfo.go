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
type DIBDeviceInfo struct {
	descriptionType                uint8
	knxMedium                      uint8
	deviceStatus                   IDeviceStatus
	knxAddress                     IKNXAddress
	projectInstallationIdentifier  IProjectInstallationIdentifier
	knxNetIpDeviceSerialNumber     []int8
	knxNetIpDeviceMulticastAddress IIPAddress
	knxNetIpDeviceMacAddress       IMACAddress
	deviceFriendlyName             []int8
}

// The corresponding interface
type IDIBDeviceInfo interface {
	spi.Message
	Serialize(io spi.WriteBuffer) error
}

func NewDIBDeviceInfo(descriptionType uint8, knxMedium uint8, deviceStatus IDeviceStatus, knxAddress IKNXAddress, projectInstallationIdentifier IProjectInstallationIdentifier, knxNetIpDeviceSerialNumber []int8, knxNetIpDeviceMulticastAddress IIPAddress, knxNetIpDeviceMacAddress IMACAddress, deviceFriendlyName []int8) spi.Message {
	return &DIBDeviceInfo{descriptionType: descriptionType, knxMedium: knxMedium, deviceStatus: deviceStatus, knxAddress: knxAddress, projectInstallationIdentifier: projectInstallationIdentifier, knxNetIpDeviceSerialNumber: knxNetIpDeviceSerialNumber, knxNetIpDeviceMulticastAddress: knxNetIpDeviceMulticastAddress, knxNetIpDeviceMacAddress: knxNetIpDeviceMacAddress, deviceFriendlyName: deviceFriendlyName}
}

func CastIDIBDeviceInfo(structType interface{}) IDIBDeviceInfo {
	castFunc := func(typ interface{}) IDIBDeviceInfo {
		if iDIBDeviceInfo, ok := typ.(IDIBDeviceInfo); ok {
			return iDIBDeviceInfo
		}
		return nil
	}
	return castFunc(structType)
}

func CastDIBDeviceInfo(structType interface{}) DIBDeviceInfo {
	castFunc := func(typ interface{}) DIBDeviceInfo {
		if sDIBDeviceInfo, ok := typ.(DIBDeviceInfo); ok {
			return sDIBDeviceInfo
		}
		return DIBDeviceInfo{}
	}
	return castFunc(structType)
}

func (m DIBDeviceInfo) LengthInBits() uint16 {
	var lengthInBits uint16 = 0

	// Implicit Field (structureLength)
	lengthInBits += 8

	// Simple field (descriptionType)
	lengthInBits += 8

	// Simple field (knxMedium)
	lengthInBits += 8

	// Simple field (deviceStatus)
	lengthInBits += m.deviceStatus.LengthInBits()

	// Simple field (knxAddress)
	lengthInBits += m.knxAddress.LengthInBits()

	// Simple field (projectInstallationIdentifier)
	lengthInBits += m.projectInstallationIdentifier.LengthInBits()

	// Array field
	if len(m.knxNetIpDeviceSerialNumber) > 0 {
		lengthInBits += 8 * uint16(len(m.knxNetIpDeviceSerialNumber))
	}

	// Simple field (knxNetIpDeviceMulticastAddress)
	lengthInBits += m.knxNetIpDeviceMulticastAddress.LengthInBits()

	// Simple field (knxNetIpDeviceMacAddress)
	lengthInBits += m.knxNetIpDeviceMacAddress.LengthInBits()

	// Array field
	if len(m.deviceFriendlyName) > 0 {
		lengthInBits += 8 * uint16(len(m.deviceFriendlyName))
	}

	return lengthInBits
}

func (m DIBDeviceInfo) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func DIBDeviceInfoParse(io *spi.ReadBuffer) (spi.Message, error) {

	// Implicit Field (structureLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	_, _structureLengthErr := io.ReadUint8(8)
	if _structureLengthErr != nil {
		return nil, errors.New("Error parsing 'structureLength' field " + _structureLengthErr.Error())
	}

	// Simple Field (descriptionType)
	descriptionType, _descriptionTypeErr := io.ReadUint8(8)
	if _descriptionTypeErr != nil {
		return nil, errors.New("Error parsing 'descriptionType' field " + _descriptionTypeErr.Error())
	}

	// Simple Field (knxMedium)
	knxMedium, _knxMediumErr := io.ReadUint8(8)
	if _knxMediumErr != nil {
		return nil, errors.New("Error parsing 'knxMedium' field " + _knxMediumErr.Error())
	}

	// Simple Field (deviceStatus)
	_deviceStatusMessage, _err := DeviceStatusParse(io)
	if _err != nil {
		return nil, errors.New("Error parsing simple field 'deviceStatus'. " + _err.Error())
	}
	var deviceStatus IDeviceStatus
	deviceStatus, _deviceStatusOk := _deviceStatusMessage.(IDeviceStatus)
	if !_deviceStatusOk {
		return nil, errors.New("Couldn't cast message of type " + reflect.TypeOf(_deviceStatusMessage).Name() + " to IDeviceStatus")
	}

	// Simple Field (knxAddress)
	_knxAddressMessage, _err := KNXAddressParse(io)
	if _err != nil {
		return nil, errors.New("Error parsing simple field 'knxAddress'. " + _err.Error())
	}
	var knxAddress IKNXAddress
	knxAddress, _knxAddressOk := _knxAddressMessage.(IKNXAddress)
	if !_knxAddressOk {
		return nil, errors.New("Couldn't cast message of type " + reflect.TypeOf(_knxAddressMessage).Name() + " to IKNXAddress")
	}

	// Simple Field (projectInstallationIdentifier)
	_projectInstallationIdentifierMessage, _err := ProjectInstallationIdentifierParse(io)
	if _err != nil {
		return nil, errors.New("Error parsing simple field 'projectInstallationIdentifier'. " + _err.Error())
	}
	var projectInstallationIdentifier IProjectInstallationIdentifier
	projectInstallationIdentifier, _projectInstallationIdentifierOk := _projectInstallationIdentifierMessage.(IProjectInstallationIdentifier)
	if !_projectInstallationIdentifierOk {
		return nil, errors.New("Couldn't cast message of type " + reflect.TypeOf(_projectInstallationIdentifierMessage).Name() + " to IProjectInstallationIdentifier")
	}

	// Array field (knxNetIpDeviceSerialNumber)
	// Count array
	knxNetIpDeviceSerialNumber := make([]int8, uint16(6))
	for curItem := uint16(0); curItem < uint16(uint16(6)); curItem++ {

		_item, _err := io.ReadInt8(8)
		if _err != nil {
			return nil, errors.New("Error parsing 'knxNetIpDeviceSerialNumber' field " + _err.Error())
		}
		knxNetIpDeviceSerialNumber[curItem] = _item
	}

	// Simple Field (knxNetIpDeviceMulticastAddress)
	_knxNetIpDeviceMulticastAddressMessage, _err := IPAddressParse(io)
	if _err != nil {
		return nil, errors.New("Error parsing simple field 'knxNetIpDeviceMulticastAddress'. " + _err.Error())
	}
	var knxNetIpDeviceMulticastAddress IIPAddress
	knxNetIpDeviceMulticastAddress, _knxNetIpDeviceMulticastAddressOk := _knxNetIpDeviceMulticastAddressMessage.(IIPAddress)
	if !_knxNetIpDeviceMulticastAddressOk {
		return nil, errors.New("Couldn't cast message of type " + reflect.TypeOf(_knxNetIpDeviceMulticastAddressMessage).Name() + " to IIPAddress")
	}

	// Simple Field (knxNetIpDeviceMacAddress)
	_knxNetIpDeviceMacAddressMessage, _err := MACAddressParse(io)
	if _err != nil {
		return nil, errors.New("Error parsing simple field 'knxNetIpDeviceMacAddress'. " + _err.Error())
	}
	var knxNetIpDeviceMacAddress IMACAddress
	knxNetIpDeviceMacAddress, _knxNetIpDeviceMacAddressOk := _knxNetIpDeviceMacAddressMessage.(IMACAddress)
	if !_knxNetIpDeviceMacAddressOk {
		return nil, errors.New("Couldn't cast message of type " + reflect.TypeOf(_knxNetIpDeviceMacAddressMessage).Name() + " to IMACAddress")
	}

	// Array field (deviceFriendlyName)
	// Count array
	deviceFriendlyName := make([]int8, uint16(30))
	for curItem := uint16(0); curItem < uint16(uint16(30)); curItem++ {

		_item, _err := io.ReadInt8(8)
		if _err != nil {
			return nil, errors.New("Error parsing 'deviceFriendlyName' field " + _err.Error())
		}
		deviceFriendlyName[curItem] = _item
	}

	// Create the instance
	return NewDIBDeviceInfo(descriptionType, knxMedium, deviceStatus, knxAddress, projectInstallationIdentifier, knxNetIpDeviceSerialNumber, knxNetIpDeviceMulticastAddress, knxNetIpDeviceMacAddress, deviceFriendlyName), nil
}

func (m DIBDeviceInfo) Serialize(io spi.WriteBuffer) error {

	// Implicit Field (structureLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	structureLength := uint8(uint8(m.LengthInBytes()))
	_structureLengthErr := io.WriteUint8(8, (structureLength))
	if _structureLengthErr != nil {
		return errors.New("Error serializing 'structureLength' field " + _structureLengthErr.Error())
	}

	// Simple Field (descriptionType)
	descriptionType := uint8(m.descriptionType)
	_descriptionTypeErr := io.WriteUint8(8, (descriptionType))
	if _descriptionTypeErr != nil {
		return errors.New("Error serializing 'descriptionType' field " + _descriptionTypeErr.Error())
	}

	// Simple Field (knxMedium)
	knxMedium := uint8(m.knxMedium)
	_knxMediumErr := io.WriteUint8(8, (knxMedium))
	if _knxMediumErr != nil {
		return errors.New("Error serializing 'knxMedium' field " + _knxMediumErr.Error())
	}

	// Simple Field (deviceStatus)
	deviceStatus := CastIDeviceStatus(m.deviceStatus)
	_deviceStatusErr := deviceStatus.Serialize(io)
	if _deviceStatusErr != nil {
		return errors.New("Error serializing 'deviceStatus' field " + _deviceStatusErr.Error())
	}

	// Simple Field (knxAddress)
	knxAddress := CastIKNXAddress(m.knxAddress)
	_knxAddressErr := knxAddress.Serialize(io)
	if _knxAddressErr != nil {
		return errors.New("Error serializing 'knxAddress' field " + _knxAddressErr.Error())
	}

	// Simple Field (projectInstallationIdentifier)
	projectInstallationIdentifier := CastIProjectInstallationIdentifier(m.projectInstallationIdentifier)
	_projectInstallationIdentifierErr := projectInstallationIdentifier.Serialize(io)
	if _projectInstallationIdentifierErr != nil {
		return errors.New("Error serializing 'projectInstallationIdentifier' field " + _projectInstallationIdentifierErr.Error())
	}

	// Array Field (knxNetIpDeviceSerialNumber)
	if m.knxNetIpDeviceSerialNumber != nil {
		for _, _element := range m.knxNetIpDeviceSerialNumber {
			_elementErr := io.WriteInt8(8, _element)
			if _elementErr != nil {
				return errors.New("Error serializing 'knxNetIpDeviceSerialNumber' field " + _elementErr.Error())
			}
		}
	}

	// Simple Field (knxNetIpDeviceMulticastAddress)
	knxNetIpDeviceMulticastAddress := CastIIPAddress(m.knxNetIpDeviceMulticastAddress)
	_knxNetIpDeviceMulticastAddressErr := knxNetIpDeviceMulticastAddress.Serialize(io)
	if _knxNetIpDeviceMulticastAddressErr != nil {
		return errors.New("Error serializing 'knxNetIpDeviceMulticastAddress' field " + _knxNetIpDeviceMulticastAddressErr.Error())
	}

	// Simple Field (knxNetIpDeviceMacAddress)
	knxNetIpDeviceMacAddress := CastIMACAddress(m.knxNetIpDeviceMacAddress)
	_knxNetIpDeviceMacAddressErr := knxNetIpDeviceMacAddress.Serialize(io)
	if _knxNetIpDeviceMacAddressErr != nil {
		return errors.New("Error serializing 'knxNetIpDeviceMacAddress' field " + _knxNetIpDeviceMacAddressErr.Error())
	}

	// Array Field (deviceFriendlyName)
	if m.deviceFriendlyName != nil {
		for _, _element := range m.deviceFriendlyName {
			_elementErr := io.WriteInt8(8, _element)
			if _elementErr != nil {
				return errors.New("Error serializing 'deviceFriendlyName' field " + _elementErr.Error())
			}
		}
	}

	return nil
}
