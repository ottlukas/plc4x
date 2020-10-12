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
type CEMIDataReq struct {
	additionalInformationLength uint8
	additionalInformation       []CEMIAdditionalInformation
	cemiDataFrame               CEMIDataFrame
	CEMI
}

// The corresponding interface
type ICEMIDataReq interface {
	ICEMI
	Serialize(io spi.WriteBuffer)
}

// Accessors for discriminator values.
func (m CEMIDataReq) MessageCode() uint8 {
	return 0x11
}

func (m CEMIDataReq) initialize() spi.Message {
	return m
}

func NewCEMIDataReq(additionalInformationLength uint8, additionalInformation []CEMIAdditionalInformation, cemiDataFrame CEMIDataFrame) CEMIInitializer {
	return &CEMIDataReq{additionalInformationLength: additionalInformationLength, additionalInformation: additionalInformation, cemiDataFrame: cemiDataFrame}
}

func CastICEMIDataReq(structType interface{}) ICEMIDataReq {
	castFunc := func(typ interface{}) ICEMIDataReq {
		if iCEMIDataReq, ok := typ.(ICEMIDataReq); ok {
			return iCEMIDataReq
		}
		return nil
	}
	return castFunc(structType)
}

func CastCEMIDataReq(structType interface{}) CEMIDataReq {
	castFunc := func(typ interface{}) CEMIDataReq {
		if sCEMIDataReq, ok := typ.(CEMIDataReq); ok {
			return sCEMIDataReq
		}
		return CEMIDataReq{}
	}
	return castFunc(structType)
}

func (m CEMIDataReq) LengthInBits() uint16 {
	var lengthInBits uint16 = m.CEMI.LengthInBits()

	// Simple field (additionalInformationLength)
	lengthInBits += 8

	// Array field
	if len(m.additionalInformation) > 0 {
		for _, element := range m.additionalInformation {
			lengthInBits += element.LengthInBits()
		}
	}

	// Simple field (cemiDataFrame)
	lengthInBits += m.cemiDataFrame.LengthInBits()

	return lengthInBits
}

func (m CEMIDataReq) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func CEMIDataReqParse(io spi.ReadBuffer) (CEMIInitializer, error) {

	// Simple Field (additionalInformationLength)
	var additionalInformationLength uint8 = io.ReadUint8(8)

	// Array field (additionalInformation)
	var additionalInformation []CEMIAdditionalInformation
	// Length array
	_additionalInformationLength := additionalInformationLength
	_additionalInformationEndPos := io.GetPos() + uint16(_additionalInformationLength)
	for io.GetPos() < _additionalInformationEndPos {
		_message, _err := CEMIAdditionalInformationParse(io)
		if _err != nil {
			return nil, errors.New("Error parsing 'additionalInformation' field " + _err.Error())
		}
		var _item CEMIAdditionalInformation
		_item, _ok := _message.(CEMIAdditionalInformation)
		if !_ok {
			return nil, errors.New("Couldn't cast message of type " + reflect.TypeOf(_item).Name() + " to CEMIAdditionalInformation")
		}
		additionalInformation = append(additionalInformation, _item)
	}

	// Simple Field (cemiDataFrame)
	_cemiDataFrameMessage, _err := CEMIDataFrameParse(io)
	if _err != nil {
		return nil, errors.New("Error parsing simple field 'cemiDataFrame'. " + _err.Error())
	}
	var cemiDataFrame CEMIDataFrame
	cemiDataFrame, _cemiDataFrameOk := _cemiDataFrameMessage.(CEMIDataFrame)
	if !_cemiDataFrameOk {
		return nil, errors.New("Couldn't cast message of type " + reflect.TypeOf(_cemiDataFrameMessage).Name() + " to CEMIDataFrame")
	}

	// Create the instance
	return NewCEMIDataReq(additionalInformationLength, additionalInformation, cemiDataFrame), nil
}

func (m CEMIDataReq) Serialize(io spi.WriteBuffer) {

	// Simple Field (additionalInformationLength)
	additionalInformationLength := uint8(m.additionalInformationLength)
	io.WriteUint8(8, (additionalInformationLength))

	// Array Field (additionalInformation)
	if m.additionalInformation != nil {
		for _, _element := range m.additionalInformation {
			_element.Serialize(io)
		}
	}

	// Simple Field (cemiDataFrame)
	cemiDataFrame := CEMIDataFrame(m.cemiDataFrame)
	cemiDataFrame.Serialize(io)
}