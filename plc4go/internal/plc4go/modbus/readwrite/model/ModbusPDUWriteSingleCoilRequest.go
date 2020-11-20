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
	"github.com/pkg/errors"
	"io"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type ModbusPDUWriteSingleCoilRequest struct {
	Address uint16
	Value   uint16
	Parent  *ModbusPDU
}

// The corresponding interface
type IModbusPDUWriteSingleCoilRequest interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ModbusPDUWriteSingleCoilRequest) ErrorFlag() bool {
	return false
}

func (m *ModbusPDUWriteSingleCoilRequest) FunctionFlag() uint8 {
	return 0x05
}

func (m *ModbusPDUWriteSingleCoilRequest) Response() bool {
	return false
}

func (m *ModbusPDUWriteSingleCoilRequest) InitializeParent(parent *ModbusPDU) {
}

func NewModbusPDUWriteSingleCoilRequest(address uint16, value uint16) *ModbusPDU {
	child := &ModbusPDUWriteSingleCoilRequest{
		Address: address,
		Value:   value,
		Parent:  NewModbusPDU(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastModbusPDUWriteSingleCoilRequest(structType interface{}) *ModbusPDUWriteSingleCoilRequest {
	castFunc := func(typ interface{}) *ModbusPDUWriteSingleCoilRequest {
		if casted, ok := typ.(ModbusPDUWriteSingleCoilRequest); ok {
			return &casted
		}
		if casted, ok := typ.(*ModbusPDUWriteSingleCoilRequest); ok {
			return casted
		}
		if casted, ok := typ.(ModbusPDU); ok {
			return CastModbusPDUWriteSingleCoilRequest(casted.Child)
		}
		if casted, ok := typ.(*ModbusPDU); ok {
			return CastModbusPDUWriteSingleCoilRequest(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ModbusPDUWriteSingleCoilRequest) GetTypeName() string {
	return "ModbusPDUWriteSingleCoilRequest"
}

func (m *ModbusPDUWriteSingleCoilRequest) GetTypeName() string {
    return "ModbusPDUWriteSingleCoilRequest"
}

func (m *ModbusPDUWriteSingleCoilRequest) LengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Simple field (address)
	lengthInBits += 16

	// Simple field (value)
	lengthInBits += 16

	return lengthInBits
}

func (m *ModbusPDUWriteSingleCoilRequest) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ModbusPDUWriteSingleCoilRequestParse(io *utils.ReadBuffer) (*ModbusPDU, error) {

	// Simple Field (address)
	address, _addressErr := io.ReadUint16(16)
	if _addressErr != nil {
		return nil, errors.Wrap(_addressErr, "Error parsing 'address' field")
	}

	// Simple Field (value)
	value, _valueErr := io.ReadUint16(16)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
	}

	// Create a partially initialized instance
	_child := &ModbusPDUWriteSingleCoilRequest{
		Address: address,
		Value:   value,
		Parent:  &ModbusPDU{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *ModbusPDUWriteSingleCoilRequest) Serialize(io utils.WriteBuffer) error {
	ser := func() error {

		// Simple Field (address)
		address := uint16(m.Address)
		_addressErr := io.WriteUint16(16, (address))
		if _addressErr != nil {
			return errors.Wrap(_addressErr, "Error serializing 'address' field")
		}

		// Simple Field (value)
		value := uint16(m.Value)
		_valueErr := io.WriteUint16(16, (value))
		if _valueErr != nil {
			return errors.Wrap(_valueErr, "Error serializing 'value' field")
		}

		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *ModbusPDUWriteSingleCoilRequest) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	token = start
	for {
		switch token.(type) {
		case xml.StartElement:
			tok := token.(xml.StartElement)
			switch tok.Name.Local {
			case "address":
				var data uint16
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.Address = data
			case "value":
				var data uint16
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.Value = data
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

func (m *ModbusPDUWriteSingleCoilRequest) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.Address, xml.StartElement{Name: xml.Name{Local: "address"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.Value, xml.StartElement{Name: xml.Name{Local: "value"}}); err != nil {
		return err
	}
	return nil
}

func (m ModbusPDUWriteSingleCoilRequest) String() string {
	return string(m.Box("ModbusPDUWriteSingleCoilRequest", utils.DefaultWidth*2))
}

func (m ModbusPDUWriteSingleCoilRequest) Box(name string, width int) utils.AsciiBox {
	if name == "" {
		name = "ModbusPDUWriteSingleCoilRequest"
	}
	boxes := make([]utils.AsciiBox, 0)
	boxes = append(boxes, utils.BoxAnything("Address", m.Address, width-2))
	boxes = append(boxes, utils.BoxAnything("Value", m.Value, width-2))
	return utils.BoxBox(name, utils.AlignBoxes(boxes, width-2), 0)
}
