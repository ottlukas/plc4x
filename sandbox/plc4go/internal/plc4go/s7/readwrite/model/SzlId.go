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
type SzlId struct {
	typeClass      SzlModuleTypeClass
	sublistExtract uint8
	sublistList    SzlSublist
}

// The corresponding interface
type ISzlId interface {
	spi.Message
	Serialize(io spi.WriteBuffer)
}

func NewSzlId(typeClass SzlModuleTypeClass, sublistExtract uint8, sublistList SzlSublist) spi.Message {
	return &SzlId{typeClass: typeClass, sublistExtract: sublistExtract, sublistList: sublistList}
}

func (m SzlId) LengthInBits() uint16 {
	var lengthInBits uint16 = 0

	// Enum Field (typeClass)
	lengthInBits += 4

	// Simple field (sublistExtract)
	lengthInBits += 4

	// Enum Field (sublistList)
	lengthInBits += 8

	return lengthInBits
}

func (m SzlId) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func SzlIdParse(io spi.ReadBuffer) (spi.Message, error) {

	// Enum field (typeClass)
	typeClass, _typeClassErr := SzlModuleTypeClassParse(io)
	if _typeClassErr != nil {
		return nil, errors.New("Error parsing 'typeClass' field " + _typeClassErr.Error())
	}

	// Simple Field (sublistExtract)
	var sublistExtract uint8 = io.ReadUint8(4)

	// Enum field (sublistList)
	sublistList, _sublistListErr := SzlSublistParse(io)
	if _sublistListErr != nil {
		return nil, errors.New("Error parsing 'sublistList' field " + _sublistListErr.Error())
	}

	// Create the instance
	return NewSzlId(typeClass, sublistExtract, sublistList), nil
}

func (m SzlId) Serialize(io spi.WriteBuffer) {
	serializeFunc := func(typ interface{}) {
		if _, ok := typ.(ISzlId); ok {

			// Enum field (typeClass)
			typeClass := m.typeClass
			typeClass.Serialize(io)

			// Simple Field (sublistExtract)
			var sublistExtract uint8 = m.sublistExtract
			io.WriteUint8(4, (sublistExtract))

			// Enum field (sublistList)
			sublistList := m.sublistList
			sublistList.Serialize(io)
		}
	}
	serializeFunc(m)
}
