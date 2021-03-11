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
	"encoding/xml"
	"errors"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"io"
)

// The data-structure of this message
type AdsWriteControlResponse struct {
	Result ReturnCode
	Parent *AdsData
	IAdsWriteControlResponse
}

// The corresponding interface
type IAdsWriteControlResponse interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *AdsWriteControlResponse) CommandId() CommandId {
	return CommandId_ADS_WRITE_CONTROL
}

func (m *AdsWriteControlResponse) Response() bool {
	return true
}

func (m *AdsWriteControlResponse) InitializeParent(parent *AdsData) {
}

func NewAdsWriteControlResponse(result ReturnCode) *AdsData {
	child := &AdsWriteControlResponse{
		Result: result,
		Parent: NewAdsData(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastAdsWriteControlResponse(structType interface{}) *AdsWriteControlResponse {
	castFunc := func(typ interface{}) *AdsWriteControlResponse {
		if casted, ok := typ.(AdsWriteControlResponse); ok {
			return &casted
		}
		if casted, ok := typ.(*AdsWriteControlResponse); ok {
			return casted
		}
		if casted, ok := typ.(AdsData); ok {
			return CastAdsWriteControlResponse(casted.Child)
		}
		if casted, ok := typ.(*AdsData); ok {
			return CastAdsWriteControlResponse(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *AdsWriteControlResponse) GetTypeName() string {
	return "AdsWriteControlResponse"
}

func (m *AdsWriteControlResponse) LengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Enum Field (result)
	lengthInBits += 32

	return lengthInBits
}

func (m *AdsWriteControlResponse) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func AdsWriteControlResponseParse(io *utils.ReadBuffer) (*AdsData, error) {

	// Enum field (result)
	result, _resultErr := ReturnCodeParse(io)
	if _resultErr != nil {
		return nil, errors.New("Error parsing 'result' field " + _resultErr.Error())
	}

	// Create a partially initialized instance
	_child := &AdsWriteControlResponse{
		Result: result,
		Parent: &AdsData{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *AdsWriteControlResponse) Serialize(io utils.WriteBuffer) error {
	ser := func() error {

		// Enum field (result)
		result := CastReturnCode(m.Result)
		_resultErr := result.Serialize(io)
		if _resultErr != nil {
			return errors.New("Error serializing 'result' field " + _resultErr.Error())
		}

		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *AdsWriteControlResponse) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	token = start
	for {
		switch token.(type) {
		case xml.StartElement:
			tok := token.(xml.StartElement)
			switch tok.Name.Local {
			case "result":
				var data ReturnCode
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.Result = data
			}
		}
		token, err = d.Token()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
}

func (m *AdsWriteControlResponse) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.Result, xml.StartElement{Name: xml.Name{Local: "result"}}); err != nil {
		return err
	}
	return nil
}
