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

// The data-structure of this message
type BACnetConfirmedServiceACKAtomicReadFile struct {
	Parent *BACnetConfirmedServiceACK
	IBACnetConfirmedServiceACKAtomicReadFile
}

// The corresponding interface
type IBACnetConfirmedServiceACKAtomicReadFile interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetConfirmedServiceACKAtomicReadFile) ServiceChoice() uint8 {
	return 0x06
}

func (m *BACnetConfirmedServiceACKAtomicReadFile) InitializeParent(parent *BACnetConfirmedServiceACK) {
}

func NewBACnetConfirmedServiceACKAtomicReadFile() *BACnetConfirmedServiceACK {
	child := &BACnetConfirmedServiceACKAtomicReadFile{
		Parent: NewBACnetConfirmedServiceACK(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastBACnetConfirmedServiceACKAtomicReadFile(structType interface{}) *BACnetConfirmedServiceACKAtomicReadFile {
	castFunc := func(typ interface{}) *BACnetConfirmedServiceACKAtomicReadFile {
		if casted, ok := typ.(BACnetConfirmedServiceACKAtomicReadFile); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetConfirmedServiceACKAtomicReadFile); ok {
			return casted
		}
		if casted, ok := typ.(BACnetConfirmedServiceACK); ok {
			return CastBACnetConfirmedServiceACKAtomicReadFile(casted.Child)
		}
		if casted, ok := typ.(*BACnetConfirmedServiceACK); ok {
			return CastBACnetConfirmedServiceACKAtomicReadFile(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetConfirmedServiceACKAtomicReadFile) GetTypeName() string {
	return "BACnetConfirmedServiceACKAtomicReadFile"
}

func (m *BACnetConfirmedServiceACKAtomicReadFile) LengthInBits() uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *BACnetConfirmedServiceACKAtomicReadFile) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetConfirmedServiceACKAtomicReadFileParse(io *utils.ReadBuffer) (*BACnetConfirmedServiceACK, error) {

	// Create a partially initialized instance
	_child := &BACnetConfirmedServiceACKAtomicReadFile{
		Parent: &BACnetConfirmedServiceACK{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *BACnetConfirmedServiceACKAtomicReadFile) Serialize(io utils.WriteBuffer) error {
	ser := func() error {

		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *BACnetConfirmedServiceACKAtomicReadFile) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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

func (m *BACnetConfirmedServiceACKAtomicReadFile) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return nil
}
