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
type DeviceConfigurationAckDataBlock struct {
	CommunicationChannelId uint8
	SequenceCounter        uint8
	Status                 Status
}

// The corresponding interface
type IDeviceConfigurationAckDataBlock interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

func NewDeviceConfigurationAckDataBlock(communicationChannelId uint8, sequenceCounter uint8, status Status) *DeviceConfigurationAckDataBlock {
	return &DeviceConfigurationAckDataBlock{CommunicationChannelId: communicationChannelId, SequenceCounter: sequenceCounter, Status: status}
}

func CastDeviceConfigurationAckDataBlock(structType interface{}) *DeviceConfigurationAckDataBlock {
	castFunc := func(typ interface{}) *DeviceConfigurationAckDataBlock {
		if casted, ok := typ.(DeviceConfigurationAckDataBlock); ok {
			return &casted
		}
		if casted, ok := typ.(*DeviceConfigurationAckDataBlock); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *DeviceConfigurationAckDataBlock) GetTypeName() string {
	return "DeviceConfigurationAckDataBlock"
}

func (m *DeviceConfigurationAckDataBlock) GetTypeName() string {
    return "DeviceConfigurationAckDataBlock"
}

func (m *DeviceConfigurationAckDataBlock) LengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Implicit Field (structureLength)
	lengthInBits += 8

	// Simple field (communicationChannelId)
	lengthInBits += 8

	// Simple field (sequenceCounter)
	lengthInBits += 8

	// Simple field (status)
	lengthInBits += 8

	return lengthInBits
}

func (m *DeviceConfigurationAckDataBlock) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func DeviceConfigurationAckDataBlockParse(io *utils.ReadBuffer) (*DeviceConfigurationAckDataBlock, error) {

	// Implicit Field (structureLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	structureLength, _structureLengthErr := io.ReadUint8(8)
	_ = structureLength
	if _structureLengthErr != nil {
		return nil, errors.Wrap(_structureLengthErr, "Error parsing 'structureLength' field")
	}

	// Simple Field (communicationChannelId)
	communicationChannelId, _communicationChannelIdErr := io.ReadUint8(8)
	if _communicationChannelIdErr != nil {
		return nil, errors.Wrap(_communicationChannelIdErr, "Error parsing 'communicationChannelId' field")
	}

	// Simple Field (sequenceCounter)
	sequenceCounter, _sequenceCounterErr := io.ReadUint8(8)
	if _sequenceCounterErr != nil {
		return nil, errors.Wrap(_sequenceCounterErr, "Error parsing 'sequenceCounter' field")
	}

	// Simple Field (status)
	status, _statusErr := StatusParse(io)
	if _statusErr != nil {
		return nil, errors.Wrap(_statusErr, "Error parsing 'status' field")
	}

	// Create the instance
	return NewDeviceConfigurationAckDataBlock(communicationChannelId, sequenceCounter, status), nil
}

func (m *DeviceConfigurationAckDataBlock) Serialize(io utils.WriteBuffer) error {

	// Implicit Field (structureLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	structureLength := uint8(uint8(m.LengthInBytes()))
	_structureLengthErr := io.WriteUint8(8, (structureLength))
	if _structureLengthErr != nil {
		return errors.Wrap(_structureLengthErr, "Error serializing 'structureLength' field")
	}

	// Simple Field (communicationChannelId)
	communicationChannelId := uint8(m.CommunicationChannelId)
	_communicationChannelIdErr := io.WriteUint8(8, (communicationChannelId))
	if _communicationChannelIdErr != nil {
		return errors.Wrap(_communicationChannelIdErr, "Error serializing 'communicationChannelId' field")
	}

	// Simple Field (sequenceCounter)
	sequenceCounter := uint8(m.SequenceCounter)
	_sequenceCounterErr := io.WriteUint8(8, (sequenceCounter))
	if _sequenceCounterErr != nil {
		return errors.Wrap(_sequenceCounterErr, "Error serializing 'sequenceCounter' field")
	}

	// Simple Field (status)
	_statusErr := m.Status.Serialize(io)
	if _statusErr != nil {
		return errors.Wrap(_statusErr, "Error serializing 'status' field")
	}

	return nil
}

func (m *DeviceConfigurationAckDataBlock) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "communicationChannelId":
				var data uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.CommunicationChannelId = data
			case "sequenceCounter":
				var data uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.SequenceCounter = data
			case "status":
				var data Status
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.Status = data
			}
		}
	}
}

func (m *DeviceConfigurationAckDataBlock) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	className := "org.apache.plc4x.java.knxnetip.readwrite.DeviceConfigurationAckDataBlock"
	if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
		{Name: xml.Name{Local: "className"}, Value: className},
	}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.CommunicationChannelId, xml.StartElement{Name: xml.Name{Local: "communicationChannelId"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.SequenceCounter, xml.StartElement{Name: xml.Name{Local: "sequenceCounter"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.Status, xml.StartElement{Name: xml.Name{Local: "status"}}); err != nil {
		return err
	}
	if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
		return err
	}
	return nil
}

func (m DeviceConfigurationAckDataBlock) String() string {
	return string(m.Box("DeviceConfigurationAckDataBlock", utils.DefaultWidth*2))
}

func (m DeviceConfigurationAckDataBlock) Box(name string, width int) utils.AsciiBox {
	if name == "" {
		name = "DeviceConfigurationAckDataBlock"
	}
	boxes := make([]utils.AsciiBox, 0)
	boxes = append(boxes, utils.BoxAnything("CommunicationChannelId", m.CommunicationChannelId, width-2))
	boxes = append(boxes, utils.BoxAnything("SequenceCounter", m.SequenceCounter, width-2))
	boxes = append(boxes, utils.BoxAnything("Status", m.Status, width-2))
	return utils.BoxBox(name, utils.AlignBoxes(boxes, width-2), 0)
}
