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
type DIBSuppSvcFamilies struct {
	descriptionType uint8
	serviceIds      []ServiceId
}

// The corresponding interface
type IDIBSuppSvcFamilies interface {
	spi.Message
	Serialize(io spi.WriteBuffer)
}

func NewDIBSuppSvcFamilies(descriptionType uint8, serviceIds []ServiceId) spi.Message {
	return &DIBSuppSvcFamilies{descriptionType: descriptionType, serviceIds: serviceIds}
}

func (m DIBSuppSvcFamilies) LengthInBits() uint16 {
	var lengthInBits uint16 = 0

	// Implicit Field (structureLength)
	lengthInBits += 8

	// Simple field (descriptionType)
	lengthInBits += 8

	// Array field
	if len(m.serviceIds) > 0 {
		for _, element := range m.serviceIds {
			lengthInBits += element.LengthInBits()
		}
	}

	return lengthInBits
}

func (m DIBSuppSvcFamilies) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func DIBSuppSvcFamiliesParse(io spi.ReadBuffer) (spi.Message, error) {

	// Implicit Field (structureLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	var _ uint8 = io.ReadUint8(8)

	// Simple Field (descriptionType)
	var descriptionType uint8 = io.ReadUint8(8)

	// Array field (serviceIds)
	var serviceIds []ServiceId
	// Count array
	{
		serviceIds := make([]ServiceId, 3)
		for curItem := uint16(0); curItem < uint16(3); curItem++ {

			_message, _err := ServiceIdParse(io)
			if _err != nil {
				return nil, errors.New("Error parsing 'serviceIds' field " + _err.Error())
			}
			var _item ServiceId
			_item, _ok := _message.(ServiceId)
			if !_ok {
				return nil, errors.New("Couldn't cast message of type " + reflect.TypeOf(_item).Name() + " to ServiceId")
			}
			serviceIds = append(serviceIds, _item)
		}
	}

	// Create the instance
	return NewDIBSuppSvcFamilies(descriptionType, serviceIds), nil
}

func (m DIBSuppSvcFamilies) Serialize(io spi.WriteBuffer) {
	serializeFunc := func(typ interface{}) {
		if _, ok := typ.(IDIBSuppSvcFamilies); ok {

			// Implicit Field (structureLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
			structureLength := uint8(m.LengthInBytes())
			io.WriteUint8(8, (structureLength))

			// Simple Field (descriptionType)
			var descriptionType uint8 = m.descriptionType
			io.WriteUint8(8, (descriptionType))

			// Array Field (serviceIds)
			if m.serviceIds != nil {
				for _, _element := range m.serviceIds {
					_element.Serialize(io)
				}
			}
		}
	}
	serializeFunc(m)
}
