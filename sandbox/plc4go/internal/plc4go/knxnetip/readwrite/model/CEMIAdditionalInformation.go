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
type CEMIAdditionalInformation struct {
}

// The corresponding interface
type ICEMIAdditionalInformation interface {
	spi.Message
	AdditionalInformationType() uint8
	Serialize(io spi.WriteBuffer)
}

type CEMIAdditionalInformationInitializer interface {
	initialize() spi.Message
}

func CEMIAdditionalInformationAdditionalInformationType(m ICEMIAdditionalInformation) uint8 {
	return m.AdditionalInformationType()
}

func CastICEMIAdditionalInformation(structType interface{}) ICEMIAdditionalInformation {
	castFunc := func(typ interface{}) ICEMIAdditionalInformation {
		if iCEMIAdditionalInformation, ok := typ.(ICEMIAdditionalInformation); ok {
			return iCEMIAdditionalInformation
		}
		return nil
	}
	return castFunc(structType)
}

func CastCEMIAdditionalInformation(structType interface{}) CEMIAdditionalInformation {
	castFunc := func(typ interface{}) CEMIAdditionalInformation {
		if sCEMIAdditionalInformation, ok := typ.(CEMIAdditionalInformation); ok {
			return sCEMIAdditionalInformation
		}
		return CEMIAdditionalInformation{}
	}
	return castFunc(structType)
}

func (m CEMIAdditionalInformation) LengthInBits() uint16 {
	var lengthInBits uint16 = 0

	// Discriminator Field (additionalInformationType)
	lengthInBits += 8

	// Length of sub-type elements will be added by sub-type...

	return lengthInBits
}

func (m CEMIAdditionalInformation) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func CEMIAdditionalInformationParse(io spi.ReadBuffer) (spi.Message, error) {

	// Discriminator Field (additionalInformationType) (Used as input to a switch field)
	additionalInformationType, _additionalInformationTypeErr := io.ReadUint8(8)
	if _additionalInformationTypeErr != nil {
		return nil, errors.New("Error parsing 'additionalInformationType' field " + _additionalInformationTypeErr.Error())
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var initializer CEMIAdditionalInformationInitializer
	var typeSwitchError error
	switch {
	case additionalInformationType == 0x03:
		initializer, typeSwitchError = CEMIAdditionalInformationBusmonitorInfoParse(io)
	case additionalInformationType == 0x04:
		initializer, typeSwitchError = CEMIAdditionalInformationRelativeTimestampParse(io)
	}
	if typeSwitchError != nil {
		return nil, errors.New("Error parsing sub-type for type-switch. " + typeSwitchError.Error())
	}

	// Create the instance
	return initializer.initialize(), nil
}

func CEMIAdditionalInformationSerialize(io spi.WriteBuffer, m CEMIAdditionalInformation, i ICEMIAdditionalInformation, childSerialize func()) {

	// Discriminator Field (additionalInformationType) (Used as input to a switch field)
	additionalInformationType := uint8(i.AdditionalInformationType())
	io.WriteUint8(8, (additionalInformationType))

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	childSerialize()

}
