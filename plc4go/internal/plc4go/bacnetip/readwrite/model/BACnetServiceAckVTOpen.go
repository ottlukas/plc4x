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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"io"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type BACnetServiceAckVTOpen struct {
	Parent *BACnetServiceAck
}

// The corresponding interface
type IBACnetServiceAckVTOpen interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetServiceAckVTOpen) ServiceChoice() uint8 {
	return 0x15
}

func (m *BACnetServiceAckVTOpen) InitializeParent(parent *BACnetServiceAck) {
}

func NewBACnetServiceAckVTOpen() *BACnetServiceAck {
	child := &BACnetServiceAckVTOpen{
		Parent: NewBACnetServiceAck(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastBACnetServiceAckVTOpen(structType interface{}) *BACnetServiceAckVTOpen {
	castFunc := func(typ interface{}) *BACnetServiceAckVTOpen {
		if casted, ok := typ.(BACnetServiceAckVTOpen); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetServiceAckVTOpen); ok {
			return casted
		}
		if casted, ok := typ.(BACnetServiceAck); ok {
			return CastBACnetServiceAckVTOpen(casted.Child)
		}
		if casted, ok := typ.(*BACnetServiceAck); ok {
			return CastBACnetServiceAckVTOpen(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetServiceAckVTOpen) GetTypeName() string {
	return "BACnetServiceAckVTOpen"
}

func (m *BACnetServiceAckVTOpen) GetTypeName() string {
    return "BACnetServiceAckVTOpen"
}

func (m *BACnetServiceAckVTOpen) LengthInBits() uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *BACnetServiceAckVTOpen) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetServiceAckVTOpenParse(io *utils.ReadBuffer) (*BACnetServiceAck, error) {

	// Create a partially initialized instance
	_child := &BACnetServiceAckVTOpen{
		Parent: &BACnetServiceAck{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *BACnetServiceAckVTOpen) Serialize(io utils.WriteBuffer) error {
	ser := func() error {

		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *BACnetServiceAckVTOpen) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	token = start
	for {
		switch token.(type) {
		case xml.StartElement:
			tok := token.(xml.StartElement)
			switch tok.Name.Local {
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

func (m *BACnetServiceAckVTOpen) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return nil
}

func (m BACnetServiceAckVTOpen) String() string {
	return string(m.Box("BACnetServiceAckVTOpen", utils.DefaultWidth*2))
}

func (m BACnetServiceAckVTOpen) Box(name string, width int) utils.AsciiBox {
	if name == "" {
		name = "BACnetServiceAckVTOpen"
	}
	boxes := make([]utils.AsciiBox, 0)
	return utils.BoxBox(name, utils.AlignBoxes(boxes, width-2), 0)
}
