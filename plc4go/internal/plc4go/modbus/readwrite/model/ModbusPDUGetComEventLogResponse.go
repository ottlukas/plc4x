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
type ModbusPDUGetComEventLogResponse struct {
	Status       uint16
	EventCount   uint16
	MessageCount uint16
	Events       []int8
	Parent       *ModbusPDU
}

// The corresponding interface
type IModbusPDUGetComEventLogResponse interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ModbusPDUGetComEventLogResponse) ErrorFlag() bool {
	return false
}

func (m *ModbusPDUGetComEventLogResponse) FunctionFlag() uint8 {
	return 0x0C
}

func (m *ModbusPDUGetComEventLogResponse) Response() bool {
	return true
}

func (m *ModbusPDUGetComEventLogResponse) InitializeParent(parent *ModbusPDU) {
}

func NewModbusPDUGetComEventLogResponse(status uint16, eventCount uint16, messageCount uint16, events []int8) *ModbusPDU {
	child := &ModbusPDUGetComEventLogResponse{
		Status:       status,
		EventCount:   eventCount,
		MessageCount: messageCount,
		Events:       events,
		Parent:       NewModbusPDU(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastModbusPDUGetComEventLogResponse(structType interface{}) *ModbusPDUGetComEventLogResponse {
	castFunc := func(typ interface{}) *ModbusPDUGetComEventLogResponse {
		if casted, ok := typ.(ModbusPDUGetComEventLogResponse); ok {
			return &casted
		}
		if casted, ok := typ.(*ModbusPDUGetComEventLogResponse); ok {
			return casted
		}
		if casted, ok := typ.(ModbusPDU); ok {
			return CastModbusPDUGetComEventLogResponse(casted.Child)
		}
		if casted, ok := typ.(*ModbusPDU); ok {
			return CastModbusPDUGetComEventLogResponse(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ModbusPDUGetComEventLogResponse) GetTypeName() string {
	return "ModbusPDUGetComEventLogResponse"
}

func (m *ModbusPDUGetComEventLogResponse) LengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Implicit Field (byteCount)
	lengthInBits += 8

	// Simple field (status)
	lengthInBits += 16

	// Simple field (eventCount)
	lengthInBits += 16

	// Simple field (messageCount)
	lengthInBits += 16

	// Array field
	if len(m.Events) > 0 {
		lengthInBits += 8 * uint16(len(m.Events))
	}

	return lengthInBits
}

func (m *ModbusPDUGetComEventLogResponse) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ModbusPDUGetComEventLogResponseParse(io *utils.ReadBuffer) (*ModbusPDU, error) {

	// Implicit Field (byteCount) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	byteCount, _byteCountErr := io.ReadUint8(8)
	_ = byteCount
	if _byteCountErr != nil {
		return nil, errors.Wrap(_byteCountErr, "Error parsing 'byteCount' field")
	}

	// Simple Field (status)
	status, _statusErr := io.ReadUint16(16)
	if _statusErr != nil {
		return nil, errors.Wrap(_statusErr, "Error parsing 'status' field")
	}

	// Simple Field (eventCount)
	eventCount, _eventCountErr := io.ReadUint16(16)
	if _eventCountErr != nil {
		return nil, errors.Wrap(_eventCountErr, "Error parsing 'eventCount' field")
	}

	// Simple Field (messageCount)
	messageCount, _messageCountErr := io.ReadUint16(16)
	if _messageCountErr != nil {
		return nil, errors.Wrap(_messageCountErr, "Error parsing 'messageCount' field")
	}

	// Array field (events)
	// Count array
	events := make([]int8, uint16(byteCount)-uint16(uint16(6)))
	for curItem := uint16(0); curItem < uint16(uint16(byteCount)-uint16(uint16(6))); curItem++ {

		_item, _err := io.ReadInt8(8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'events' field")
		}
		events[curItem] = _item
	}

	// Create a partially initialized instance
	_child := &ModbusPDUGetComEventLogResponse{
		Status:       status,
		EventCount:   eventCount,
		MessageCount: messageCount,
		Events:       events,
		Parent:       &ModbusPDU{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *ModbusPDUGetComEventLogResponse) Serialize(io utils.WriteBuffer) error {
	ser := func() error {

		// Implicit Field (byteCount) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		byteCount := uint8(uint8(uint8(len(m.Events))) + uint8(uint8(6)))
		_byteCountErr := io.WriteUint8(8, (byteCount))
		if _byteCountErr != nil {
			return errors.Wrap(_byteCountErr, "Error serializing 'byteCount' field")
		}

		// Simple Field (status)
		status := uint16(m.Status)
		_statusErr := io.WriteUint16(16, (status))
		if _statusErr != nil {
			return errors.Wrap(_statusErr, "Error serializing 'status' field")
		}

		// Simple Field (eventCount)
		eventCount := uint16(m.EventCount)
		_eventCountErr := io.WriteUint16(16, (eventCount))
		if _eventCountErr != nil {
			return errors.Wrap(_eventCountErr, "Error serializing 'eventCount' field")
		}

		// Simple Field (messageCount)
		messageCount := uint16(m.MessageCount)
		_messageCountErr := io.WriteUint16(16, (messageCount))
		if _messageCountErr != nil {
			return errors.Wrap(_messageCountErr, "Error serializing 'messageCount' field")
		}

		// Array Field (events)
		if m.Events != nil {
			for _, _element := range m.Events {
				_elementErr := io.WriteInt8(8, _element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'events' field")
				}
			}
		}

		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *ModbusPDUGetComEventLogResponse) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	token = start
	for {
		switch token.(type) {
		case xml.StartElement:
			tok := token.(xml.StartElement)
			switch tok.Name.Local {
			case "status":
				var data uint16
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.Status = data
			case "eventCount":
				var data uint16
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.EventCount = data
			case "messageCount":
				var data uint16
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.MessageCount = data
			case "events":
				var _encoded string
				if err := d.DecodeElement(&_encoded, &tok); err != nil {
					return err
				}
				_decoded, err := hex.DecodeString(_encoded)
				_len := len(_decoded)
				if err != nil {
					return err
				}
				m.Events = utils.ByteArrayToInt8Array(_decoded[0:_len])
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

func (m *ModbusPDUGetComEventLogResponse) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.Status, xml.StartElement{Name: xml.Name{Local: "status"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.EventCount, xml.StartElement{Name: xml.Name{Local: "eventCount"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.MessageCount, xml.StartElement{Name: xml.Name{Local: "messageCount"}}); err != nil {
		return err
	}
	_encodedEvents := hex.EncodeToString(utils.Int8ArrayToByteArray(m.Events))
	_encodedEvents = strings.ToUpper(_encodedEvents)
	if err := e.EncodeElement(_encodedEvents, xml.StartElement{Name: xml.Name{Local: "events"}}); err != nil {
		return err
	}
	return nil
}

func (m ModbusPDUGetComEventLogResponse) String() string {
	return string(m.Box("ModbusPDUGetComEventLogResponse", utils.DefaultWidth*2))
}

func (m ModbusPDUGetComEventLogResponse) Box(name string, width int) utils.AsciiBox {
	if name == "" {
		name = "ModbusPDUGetComEventLogResponse"
	}
	boxes := make([]utils.AsciiBox, 0)
	boxes = append(boxes, utils.BoxAnything("Status", m.Status, width-2))
	boxes = append(boxes, utils.BoxAnything("EventCount", m.EventCount, width-2))
	boxes = append(boxes, utils.BoxAnything("MessageCount", m.MessageCount, width-2))
	boxes = append(boxes, utils.BoxAnything("Events", m.Events, width-2))
	return utils.BoxBox(name, utils.AlignBoxes(boxes, width-2), 0)
}
