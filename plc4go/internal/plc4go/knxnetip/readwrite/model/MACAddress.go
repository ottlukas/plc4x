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
	"encoding/hex"
	"encoding/xml"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
	"io"
	"strings"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type MACAddress struct {
	Addr []int8
}

// The corresponding interface
type IMACAddress interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

func NewMACAddress(addr []int8) *MACAddress {
	return &MACAddress{Addr: addr}
}

func CastMACAddress(structType interface{}) *MACAddress {
	castFunc := func(typ interface{}) *MACAddress {
		if casted, ok := typ.(MACAddress); ok {
			return &casted
		}
		if casted, ok := typ.(*MACAddress); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *MACAddress) GetTypeName() string {
	return "MACAddress"
}

func (m *MACAddress) LengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Array field
	if len(m.Addr) > 0 {
		lengthInBits += 8 * uint16(len(m.Addr))
	}

	return lengthInBits
}

func (m *MACAddress) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func MACAddressParse(io *utils.ReadBuffer) (*MACAddress, error) {

	// Array field (addr)
	// Count array
	addr := make([]int8, uint16(6))
	for curItem := uint16(0); curItem < uint16(uint16(6)); curItem++ {

		_item, _err := io.ReadInt8(8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'addr' field")
		}
		addr[curItem] = _item
	}

	// Create the instance
	return NewMACAddress(addr), nil
}

func (m *MACAddress) Serialize(io utils.WriteBuffer) error {

	// Array Field (addr)
	if m.Addr != nil {
		for _, _element := range m.Addr {
			_elementErr := io.WriteInt8(8, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'addr' field")
			}
		}
	}

	return nil
}

func (m *MACAddress) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	for {
		token, err = d.Token()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		switch token.(type) {
		case xml.StartElement:
			tok := token.(xml.StartElement)
			switch tok.Name.Local {
			case "addr":
				var _encoded string
				if err := d.DecodeElement(&_encoded, &tok); err != nil {
					return err
				}
				_decoded, err := hex.DecodeString(_encoded)
				_len := len(_decoded)
				if err != nil {
					return err
				}
				m.Addr = utils.ByteArrayToInt8Array(_decoded[0:_len])
			}
		}
	}
}

func (m *MACAddress) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	className := "org.apache.plc4x.java.knxnetip.readwrite.MACAddress"
	if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
		{Name: xml.Name{Local: "className"}, Value: className},
	}}); err != nil {
		return err
	}
	_encodedAddr := hex.EncodeToString(utils.Int8ArrayToByteArray(m.Addr))
	_encodedAddr = strings.ToUpper(_encodedAddr)
	if err := e.EncodeElement(_encodedAddr, xml.StartElement{Name: xml.Name{Local: "addr"}}); err != nil {
		return err
	}
	if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
		return err
	}
	return nil
}

func (m MACAddress) String() string {
	return string(m.Box("MACAddress", utils.DefaultWidth*2))
}

func (m MACAddress) Box(name string, width int) utils.AsciiBox {
	if name == "" {
		name = "MACAddress"
	}
	boxes := make([]utils.AsciiBox, 0)
	boxes = append(boxes, utils.BoxAnything("Addr", m.Addr, width-2))
	return utils.BoxBox(name, utils.AlignBoxes(boxes, width-2), 0)
}
