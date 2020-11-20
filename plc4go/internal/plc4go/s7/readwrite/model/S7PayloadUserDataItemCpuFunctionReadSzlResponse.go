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
	"fmt"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
	"io"
)

// Code generated by build-utils. DO NOT EDIT.

// Constant values.
const S7PayloadUserDataItemCpuFunctionReadSzlResponse_SZLITEMLENGTH uint16 = 28

// The data-structure of this message
type S7PayloadUserDataItemCpuFunctionReadSzlResponse struct {
	Items  []*SzlDataTreeItem
	Parent *S7PayloadUserDataItem
}

// The corresponding interface
type IS7PayloadUserDataItemCpuFunctionReadSzlResponse interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *S7PayloadUserDataItemCpuFunctionReadSzlResponse) CpuFunctionType() uint8 {
	return 0x08
}

func (m *S7PayloadUserDataItemCpuFunctionReadSzlResponse) InitializeParent(parent *S7PayloadUserDataItem, returnCode DataTransportErrorCode, transportSize DataTransportSize, szlId *SzlId, szlIndex uint16) {
	m.Parent.ReturnCode = returnCode
	m.Parent.TransportSize = transportSize
	m.Parent.SzlId = szlId
	m.Parent.SzlIndex = szlIndex
}

func NewS7PayloadUserDataItemCpuFunctionReadSzlResponse(items []*SzlDataTreeItem, returnCode DataTransportErrorCode, transportSize DataTransportSize, szlId *SzlId, szlIndex uint16) *S7PayloadUserDataItem {
	child := &S7PayloadUserDataItemCpuFunctionReadSzlResponse{
		Items:  items,
		Parent: NewS7PayloadUserDataItem(returnCode, transportSize, szlId, szlIndex),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastS7PayloadUserDataItemCpuFunctionReadSzlResponse(structType interface{}) *S7PayloadUserDataItemCpuFunctionReadSzlResponse {
	castFunc := func(typ interface{}) *S7PayloadUserDataItemCpuFunctionReadSzlResponse {
		if casted, ok := typ.(S7PayloadUserDataItemCpuFunctionReadSzlResponse); ok {
			return &casted
		}
		if casted, ok := typ.(*S7PayloadUserDataItemCpuFunctionReadSzlResponse); ok {
			return casted
		}
		if casted, ok := typ.(S7PayloadUserDataItem); ok {
			return CastS7PayloadUserDataItemCpuFunctionReadSzlResponse(casted.Child)
		}
		if casted, ok := typ.(*S7PayloadUserDataItem); ok {
			return CastS7PayloadUserDataItemCpuFunctionReadSzlResponse(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *S7PayloadUserDataItemCpuFunctionReadSzlResponse) GetTypeName() string {
	return "S7PayloadUserDataItemCpuFunctionReadSzlResponse"
}

func (m *S7PayloadUserDataItemCpuFunctionReadSzlResponse) GetTypeName() string {
    return "S7PayloadUserDataItemCpuFunctionReadSzlResponse"
}

func (m *S7PayloadUserDataItemCpuFunctionReadSzlResponse) LengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Const Field (szlItemLength)
	lengthInBits += 16

	// Implicit Field (szlItemCount)
	lengthInBits += 16

	// Array field
	if len(m.Items) > 0 {
		for _, element := range m.Items {
			lengthInBits += element.LengthInBits()
		}
	}

	return lengthInBits
}

func (m *S7PayloadUserDataItemCpuFunctionReadSzlResponse) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func S7PayloadUserDataItemCpuFunctionReadSzlResponseParse(io *utils.ReadBuffer) (*S7PayloadUserDataItem, error) {

	// Const Field (szlItemLength)
	szlItemLength, _szlItemLengthErr := io.ReadUint16(16)
	if _szlItemLengthErr != nil {
		return nil, errors.Wrap(_szlItemLengthErr, "Error parsing 'szlItemLength' field")
	}
	if szlItemLength != S7PayloadUserDataItemCpuFunctionReadSzlResponse_SZLITEMLENGTH {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", S7PayloadUserDataItemCpuFunctionReadSzlResponse_SZLITEMLENGTH) + " but got " + fmt.Sprintf("%d", szlItemLength))
	}

	// Implicit Field (szlItemCount) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	szlItemCount, _szlItemCountErr := io.ReadUint16(16)
	_ = szlItemCount
	if _szlItemCountErr != nil {
		return nil, errors.Wrap(_szlItemCountErr, "Error parsing 'szlItemCount' field")
	}

	// Array field (items)
	// Count array
	items := make([]*SzlDataTreeItem, szlItemCount)
	for curItem := uint16(0); curItem < uint16(szlItemCount); curItem++ {
		_item, _err := SzlDataTreeItemParse(io)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'items' field")
		}
		items[curItem] = _item
	}

	// Create a partially initialized instance
	_child := &S7PayloadUserDataItemCpuFunctionReadSzlResponse{
		Items:  items,
		Parent: &S7PayloadUserDataItem{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *S7PayloadUserDataItemCpuFunctionReadSzlResponse) Serialize(io utils.WriteBuffer) error {
	ser := func() error {

		// Const Field (szlItemLength)
		_szlItemLengthErr := io.WriteUint16(16, 28)
		if _szlItemLengthErr != nil {
			return errors.Wrap(_szlItemLengthErr, "Error serializing 'szlItemLength' field")
		}

		// Implicit Field (szlItemCount) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		szlItemCount := uint16(uint16(len(m.Items)))
		_szlItemCountErr := io.WriteUint16(16, (szlItemCount))
		if _szlItemCountErr != nil {
			return errors.Wrap(_szlItemCountErr, "Error serializing 'szlItemCount' field")
		}

		// Array Field (items)
		if m.Items != nil {
			for _, _element := range m.Items {
				_elementErr := _element.Serialize(io)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'items' field")
				}
			}
		}

		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *S7PayloadUserDataItemCpuFunctionReadSzlResponse) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	token = start
	for {
		switch token.(type) {
		case xml.StartElement:
			tok := token.(xml.StartElement)
			switch tok.Name.Local {
			case "items":
				var data []*SzlDataTreeItem
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.Items = data
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

func (m *S7PayloadUserDataItemCpuFunctionReadSzlResponse) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, arrayElement := range m.Items {
		if err := e.EncodeToken(xml.StartElement{Name: xml.Name{Local: "items"}}); err != nil {
			return err
		}
		if err := e.EncodeElement(arrayElement, xml.StartElement{Name: xml.Name{Local: "items"}}); err != nil {
			return err
		}
		if err := e.EncodeToken(xml.EndElement{Name: xml.Name{Local: "items"}}); err != nil {
			return err
		}
	}
	return nil
}

func (m S7PayloadUserDataItemCpuFunctionReadSzlResponse) String() string {
	return string(m.Box("S7PayloadUserDataItemCpuFunctionReadSzlResponse", utils.DefaultWidth*2))
}

func (m S7PayloadUserDataItemCpuFunctionReadSzlResponse) Box(name string, width int) utils.AsciiBox {
	if name == "" {
		name = "S7PayloadUserDataItemCpuFunctionReadSzlResponse"
	}
	boxes := make([]utils.AsciiBox, 0)
	boxes = append(boxes, utils.BoxAnything("Items", m.Items, width-2))
	return utils.BoxBox(name, utils.AlignBoxes(boxes, width-2), 0)
}
