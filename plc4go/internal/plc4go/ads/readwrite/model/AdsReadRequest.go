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
type AdsReadRequest struct {
	IndexGroup  uint32
	IndexOffset uint32
	Length      uint32
	Parent      *AdsData
	IAdsReadRequest
}

// The corresponding interface
type IAdsReadRequest interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *AdsReadRequest) CommandId() CommandId {
	return CommandId_ADS_READ
}

func (m *AdsReadRequest) Response() bool {
	return false
}

func (m *AdsReadRequest) InitializeParent(parent *AdsData) {
}

func NewAdsReadRequest(indexGroup uint32, indexOffset uint32, length uint32) *AdsData {
	child := &AdsReadRequest{
		IndexGroup:  indexGroup,
		IndexOffset: indexOffset,
		Length:      length,
		Parent:      NewAdsData(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastAdsReadRequest(structType interface{}) *AdsReadRequest {
	castFunc := func(typ interface{}) *AdsReadRequest {
		if casted, ok := typ.(AdsReadRequest); ok {
			return &casted
		}
		if casted, ok := typ.(*AdsReadRequest); ok {
			return casted
		}
		if casted, ok := typ.(AdsData); ok {
			return CastAdsReadRequest(casted.Child)
		}
		if casted, ok := typ.(*AdsData); ok {
			return CastAdsReadRequest(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *AdsReadRequest) GetTypeName() string {
	return "AdsReadRequest"
}

func (m *AdsReadRequest) LengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Simple field (indexGroup)
	lengthInBits += 32

	// Simple field (indexOffset)
	lengthInBits += 32

	// Simple field (length)
	lengthInBits += 32

	return lengthInBits
}

func (m *AdsReadRequest) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func AdsReadRequestParse(io *utils.ReadBuffer) (*AdsData, error) {

	// Simple Field (indexGroup)
	indexGroup, _indexGroupErr := io.ReadUint32(32)
	if _indexGroupErr != nil {
		return nil, errors.New("Error parsing 'indexGroup' field " + _indexGroupErr.Error())
	}

	// Simple Field (indexOffset)
	indexOffset, _indexOffsetErr := io.ReadUint32(32)
	if _indexOffsetErr != nil {
		return nil, errors.New("Error parsing 'indexOffset' field " + _indexOffsetErr.Error())
	}

	// Simple Field (length)
	length, _lengthErr := io.ReadUint32(32)
	if _lengthErr != nil {
		return nil, errors.New("Error parsing 'length' field " + _lengthErr.Error())
	}

	// Create a partially initialized instance
	_child := &AdsReadRequest{
		IndexGroup:  indexGroup,
		IndexOffset: indexOffset,
		Length:      length,
		Parent:      &AdsData{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *AdsReadRequest) Serialize(io utils.WriteBuffer) error {
	ser := func() error {

		// Simple Field (indexGroup)
		indexGroup := uint32(m.IndexGroup)
		_indexGroupErr := io.WriteUint32(32, (indexGroup))
		if _indexGroupErr != nil {
			return errors.New("Error serializing 'indexGroup' field " + _indexGroupErr.Error())
		}

		// Simple Field (indexOffset)
		indexOffset := uint32(m.IndexOffset)
		_indexOffsetErr := io.WriteUint32(32, (indexOffset))
		if _indexOffsetErr != nil {
			return errors.New("Error serializing 'indexOffset' field " + _indexOffsetErr.Error())
		}

		// Simple Field (length)
		length := uint32(m.Length)
		_lengthErr := io.WriteUint32(32, (length))
		if _lengthErr != nil {
			return errors.New("Error serializing 'length' field " + _lengthErr.Error())
		}

		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *AdsReadRequest) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	token = start
	for {
		switch token.(type) {
		case xml.StartElement:
			tok := token.(xml.StartElement)
			switch tok.Name.Local {
			case "indexGroup":
				var data uint32
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.IndexGroup = data
			case "indexOffset":
				var data uint32
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.IndexOffset = data
			case "length":
				var data uint32
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.Length = data
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

func (m *AdsReadRequest) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.IndexGroup, xml.StartElement{Name: xml.Name{Local: "indexGroup"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.IndexOffset, xml.StartElement{Name: xml.Name{Local: "indexOffset"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.Length, xml.StartElement{Name: xml.Name{Local: "length"}}); err != nil {
		return err
	}
	return nil
}
