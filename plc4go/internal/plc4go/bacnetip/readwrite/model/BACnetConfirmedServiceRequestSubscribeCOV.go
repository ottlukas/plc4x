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
	"strconv"
	"strings"
)

// Code generated by build-utils. DO NOT EDIT.

// Constant values.
const BACnetConfirmedServiceRequestSubscribeCOV_SUBSCRIBERPROCESSIDENTIFIERHEADER uint8 = 0x09
const BACnetConfirmedServiceRequestSubscribeCOV_MONITOREDOBJECTIDENTIFIERHEADER uint8 = 0x1C
const BACnetConfirmedServiceRequestSubscribeCOV_ISSUECONFIRMEDNOTIFICATIONSHEADER uint8 = 0x29
const BACnetConfirmedServiceRequestSubscribeCOV_ISSUECONFIRMEDNOTIFICATIONSSKIPBITS uint8 = 0x00
const BACnetConfirmedServiceRequestSubscribeCOV_LIFETIMEHEADER uint8 = 0x07

// The data-structure of this message
type BACnetConfirmedServiceRequestSubscribeCOV struct {
	SubscriberProcessIdentifier   uint8
	MonitoredObjectType           uint16
	MonitoredObjectInstanceNumber uint32
	IssueConfirmedNotifications   bool
	LifetimeLength                uint8
	LifetimeSeconds               []int8
	Parent                        *BACnetConfirmedServiceRequest
}

// The corresponding interface
type IBACnetConfirmedServiceRequestSubscribeCOV interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetConfirmedServiceRequestSubscribeCOV) ServiceChoice() uint8 {
	return 0x05
}

func (m *BACnetConfirmedServiceRequestSubscribeCOV) InitializeParent(parent *BACnetConfirmedServiceRequest) {
}

func NewBACnetConfirmedServiceRequestSubscribeCOV(subscriberProcessIdentifier uint8, monitoredObjectType uint16, monitoredObjectInstanceNumber uint32, issueConfirmedNotifications bool, lifetimeLength uint8, lifetimeSeconds []int8) *BACnetConfirmedServiceRequest {
	child := &BACnetConfirmedServiceRequestSubscribeCOV{
		SubscriberProcessIdentifier:   subscriberProcessIdentifier,
		MonitoredObjectType:           monitoredObjectType,
		MonitoredObjectInstanceNumber: monitoredObjectInstanceNumber,
		IssueConfirmedNotifications:   issueConfirmedNotifications,
		LifetimeLength:                lifetimeLength,
		LifetimeSeconds:               lifetimeSeconds,
		Parent:                        NewBACnetConfirmedServiceRequest(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastBACnetConfirmedServiceRequestSubscribeCOV(structType interface{}) *BACnetConfirmedServiceRequestSubscribeCOV {
	castFunc := func(typ interface{}) *BACnetConfirmedServiceRequestSubscribeCOV {
		if casted, ok := typ.(BACnetConfirmedServiceRequestSubscribeCOV); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetConfirmedServiceRequestSubscribeCOV); ok {
			return casted
		}
		if casted, ok := typ.(BACnetConfirmedServiceRequest); ok {
			return CastBACnetConfirmedServiceRequestSubscribeCOV(casted.Child)
		}
		if casted, ok := typ.(*BACnetConfirmedServiceRequest); ok {
			return CastBACnetConfirmedServiceRequestSubscribeCOV(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetConfirmedServiceRequestSubscribeCOV) GetTypeName() string {
	return "BACnetConfirmedServiceRequestSubscribeCOV"
}

func (m *BACnetConfirmedServiceRequestSubscribeCOV) LengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Const Field (subscriberProcessIdentifierHeader)
	lengthInBits += 8

	// Simple field (subscriberProcessIdentifier)
	lengthInBits += 8

	// Const Field (monitoredObjectIdentifierHeader)
	lengthInBits += 8

	// Simple field (monitoredObjectType)
	lengthInBits += 10

	// Simple field (monitoredObjectInstanceNumber)
	lengthInBits += 22

	// Const Field (issueConfirmedNotificationsHeader)
	lengthInBits += 8

	// Const Field (issueConfirmedNotificationsSkipBits)
	lengthInBits += 7

	// Simple field (issueConfirmedNotifications)
	lengthInBits += 1

	// Const Field (lifetimeHeader)
	lengthInBits += 5

	// Simple field (lifetimeLength)
	lengthInBits += 3

	// Array field
	if len(m.LifetimeSeconds) > 0 {
		lengthInBits += 8 * uint16(len(m.LifetimeSeconds))
	}

	return lengthInBits
}

func (m *BACnetConfirmedServiceRequestSubscribeCOV) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetConfirmedServiceRequestSubscribeCOVParse(io *utils.ReadBuffer) (*BACnetConfirmedServiceRequest, error) {

	// Const Field (subscriberProcessIdentifierHeader)
	subscriberProcessIdentifierHeader, _subscriberProcessIdentifierHeaderErr := io.ReadUint8(8)
	if _subscriberProcessIdentifierHeaderErr != nil {
		return nil, errors.Wrap(_subscriberProcessIdentifierHeaderErr, "Error parsing 'subscriberProcessIdentifierHeader' field")
	}
	if subscriberProcessIdentifierHeader != BACnetConfirmedServiceRequestSubscribeCOV_SUBSCRIBERPROCESSIDENTIFIERHEADER {
		return nil, errors.New("Expected constant value " + strconv.Itoa(int(BACnetConfirmedServiceRequestSubscribeCOV_SUBSCRIBERPROCESSIDENTIFIERHEADER)) + " but got " + strconv.Itoa(int(subscriberProcessIdentifierHeader)))
	}

	// Simple Field (subscriberProcessIdentifier)
	subscriberProcessIdentifier, _subscriberProcessIdentifierErr := io.ReadUint8(8)
	if _subscriberProcessIdentifierErr != nil {
		return nil, errors.Wrap(_subscriberProcessIdentifierErr, "Error parsing 'subscriberProcessIdentifier' field")
	}

	// Const Field (monitoredObjectIdentifierHeader)
	monitoredObjectIdentifierHeader, _monitoredObjectIdentifierHeaderErr := io.ReadUint8(8)
	if _monitoredObjectIdentifierHeaderErr != nil {
		return nil, errors.Wrap(_monitoredObjectIdentifierHeaderErr, "Error parsing 'monitoredObjectIdentifierHeader' field")
	}
	if monitoredObjectIdentifierHeader != BACnetConfirmedServiceRequestSubscribeCOV_MONITOREDOBJECTIDENTIFIERHEADER {
		return nil, errors.New("Expected constant value " + strconv.Itoa(int(BACnetConfirmedServiceRequestSubscribeCOV_MONITOREDOBJECTIDENTIFIERHEADER)) + " but got " + strconv.Itoa(int(monitoredObjectIdentifierHeader)))
	}

	// Simple Field (monitoredObjectType)
	monitoredObjectType, _monitoredObjectTypeErr := io.ReadUint16(10)
	if _monitoredObjectTypeErr != nil {
		return nil, errors.Wrap(_monitoredObjectTypeErr, "Error parsing 'monitoredObjectType' field")
	}

	// Simple Field (monitoredObjectInstanceNumber)
	monitoredObjectInstanceNumber, _monitoredObjectInstanceNumberErr := io.ReadUint32(22)
	if _monitoredObjectInstanceNumberErr != nil {
		return nil, errors.Wrap(_monitoredObjectInstanceNumberErr, "Error parsing 'monitoredObjectInstanceNumber' field")
	}

	// Const Field (issueConfirmedNotificationsHeader)
	issueConfirmedNotificationsHeader, _issueConfirmedNotificationsHeaderErr := io.ReadUint8(8)
	if _issueConfirmedNotificationsHeaderErr != nil {
		return nil, errors.Wrap(_issueConfirmedNotificationsHeaderErr, "Error parsing 'issueConfirmedNotificationsHeader' field")
	}
	if issueConfirmedNotificationsHeader != BACnetConfirmedServiceRequestSubscribeCOV_ISSUECONFIRMEDNOTIFICATIONSHEADER {
		return nil, errors.New("Expected constant value " + strconv.Itoa(int(BACnetConfirmedServiceRequestSubscribeCOV_ISSUECONFIRMEDNOTIFICATIONSHEADER)) + " but got " + strconv.Itoa(int(issueConfirmedNotificationsHeader)))
	}

	// Const Field (issueConfirmedNotificationsSkipBits)
	issueConfirmedNotificationsSkipBits, _issueConfirmedNotificationsSkipBitsErr := io.ReadUint8(7)
	if _issueConfirmedNotificationsSkipBitsErr != nil {
		return nil, errors.Wrap(_issueConfirmedNotificationsSkipBitsErr, "Error parsing 'issueConfirmedNotificationsSkipBits' field")
	}
	if issueConfirmedNotificationsSkipBits != BACnetConfirmedServiceRequestSubscribeCOV_ISSUECONFIRMEDNOTIFICATIONSSKIPBITS {
		return nil, errors.New("Expected constant value " + strconv.Itoa(int(BACnetConfirmedServiceRequestSubscribeCOV_ISSUECONFIRMEDNOTIFICATIONSSKIPBITS)) + " but got " + strconv.Itoa(int(issueConfirmedNotificationsSkipBits)))
	}

	// Simple Field (issueConfirmedNotifications)
	issueConfirmedNotifications, _issueConfirmedNotificationsErr := io.ReadBit()
	if _issueConfirmedNotificationsErr != nil {
		return nil, errors.Wrap(_issueConfirmedNotificationsErr, "Error parsing 'issueConfirmedNotifications' field")
	}

	// Const Field (lifetimeHeader)
	lifetimeHeader, _lifetimeHeaderErr := io.ReadUint8(5)
	if _lifetimeHeaderErr != nil {
		return nil, errors.Wrap(_lifetimeHeaderErr, "Error parsing 'lifetimeHeader' field")
	}
	if lifetimeHeader != BACnetConfirmedServiceRequestSubscribeCOV_LIFETIMEHEADER {
		return nil, errors.New("Expected constant value " + strconv.Itoa(int(BACnetConfirmedServiceRequestSubscribeCOV_LIFETIMEHEADER)) + " but got " + strconv.Itoa(int(lifetimeHeader)))
	}

	// Simple Field (lifetimeLength)
	lifetimeLength, _lifetimeLengthErr := io.ReadUint8(3)
	if _lifetimeLengthErr != nil {
		return nil, errors.Wrap(_lifetimeLengthErr, "Error parsing 'lifetimeLength' field")
	}

	// Array field (lifetimeSeconds)
	// Count array
	lifetimeSeconds := make([]int8, lifetimeLength)
	for curItem := uint16(0); curItem < uint16(lifetimeLength); curItem++ {
		_item, _err := io.ReadInt8(8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'lifetimeSeconds' field")
		}
		lifetimeSeconds[curItem] = _item
	}

	// Create a partially initialized instance
	_child := &BACnetConfirmedServiceRequestSubscribeCOV{
		SubscriberProcessIdentifier:   subscriberProcessIdentifier,
		MonitoredObjectType:           monitoredObjectType,
		MonitoredObjectInstanceNumber: monitoredObjectInstanceNumber,
		IssueConfirmedNotifications:   issueConfirmedNotifications,
		LifetimeLength:                lifetimeLength,
		LifetimeSeconds:               lifetimeSeconds,
		Parent:                        &BACnetConfirmedServiceRequest{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *BACnetConfirmedServiceRequestSubscribeCOV) Serialize(io utils.WriteBuffer) error {
	ser := func() error {

		// Const Field (subscriberProcessIdentifierHeader)
		_subscriberProcessIdentifierHeaderErr := io.WriteUint8(8, 0x09)
		if _subscriberProcessIdentifierHeaderErr != nil {
			return errors.Wrap(_subscriberProcessIdentifierHeaderErr, "Error serializing 'subscriberProcessIdentifierHeader' field")
		}

		// Simple Field (subscriberProcessIdentifier)
		subscriberProcessIdentifier := uint8(m.SubscriberProcessIdentifier)
		_subscriberProcessIdentifierErr := io.WriteUint8(8, (subscriberProcessIdentifier))
		if _subscriberProcessIdentifierErr != nil {
			return errors.Wrap(_subscriberProcessIdentifierErr, "Error serializing 'subscriberProcessIdentifier' field")
		}

		// Const Field (monitoredObjectIdentifierHeader)
		_monitoredObjectIdentifierHeaderErr := io.WriteUint8(8, 0x1C)
		if _monitoredObjectIdentifierHeaderErr != nil {
			return errors.Wrap(_monitoredObjectIdentifierHeaderErr, "Error serializing 'monitoredObjectIdentifierHeader' field")
		}

		// Simple Field (monitoredObjectType)
		monitoredObjectType := uint16(m.MonitoredObjectType)
		_monitoredObjectTypeErr := io.WriteUint16(10, (monitoredObjectType))
		if _monitoredObjectTypeErr != nil {
			return errors.Wrap(_monitoredObjectTypeErr, "Error serializing 'monitoredObjectType' field")
		}

		// Simple Field (monitoredObjectInstanceNumber)
		monitoredObjectInstanceNumber := uint32(m.MonitoredObjectInstanceNumber)
		_monitoredObjectInstanceNumberErr := io.WriteUint32(22, (monitoredObjectInstanceNumber))
		if _monitoredObjectInstanceNumberErr != nil {
			return errors.Wrap(_monitoredObjectInstanceNumberErr, "Error serializing 'monitoredObjectInstanceNumber' field")
		}

		// Const Field (issueConfirmedNotificationsHeader)
		_issueConfirmedNotificationsHeaderErr := io.WriteUint8(8, 0x29)
		if _issueConfirmedNotificationsHeaderErr != nil {
			return errors.Wrap(_issueConfirmedNotificationsHeaderErr, "Error serializing 'issueConfirmedNotificationsHeader' field")
		}

		// Const Field (issueConfirmedNotificationsSkipBits)
		_issueConfirmedNotificationsSkipBitsErr := io.WriteUint8(7, 0x00)
		if _issueConfirmedNotificationsSkipBitsErr != nil {
			return errors.Wrap(_issueConfirmedNotificationsSkipBitsErr, "Error serializing 'issueConfirmedNotificationsSkipBits' field")
		}

		// Simple Field (issueConfirmedNotifications)
		issueConfirmedNotifications := bool(m.IssueConfirmedNotifications)
		_issueConfirmedNotificationsErr := io.WriteBit((issueConfirmedNotifications))
		if _issueConfirmedNotificationsErr != nil {
			return errors.Wrap(_issueConfirmedNotificationsErr, "Error serializing 'issueConfirmedNotifications' field")
		}

		// Const Field (lifetimeHeader)
		_lifetimeHeaderErr := io.WriteUint8(5, 0x07)
		if _lifetimeHeaderErr != nil {
			return errors.Wrap(_lifetimeHeaderErr, "Error serializing 'lifetimeHeader' field")
		}

		// Simple Field (lifetimeLength)
		lifetimeLength := uint8(m.LifetimeLength)
		_lifetimeLengthErr := io.WriteUint8(3, (lifetimeLength))
		if _lifetimeLengthErr != nil {
			return errors.Wrap(_lifetimeLengthErr, "Error serializing 'lifetimeLength' field")
		}

		// Array Field (lifetimeSeconds)
		if m.LifetimeSeconds != nil {
			for _, _element := range m.LifetimeSeconds {
				_elementErr := io.WriteInt8(8, _element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'lifetimeSeconds' field")
				}
			}
		}

		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *BACnetConfirmedServiceRequestSubscribeCOV) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	token = start
	for {
		switch token.(type) {
		case xml.StartElement:
			tok := token.(xml.StartElement)
			switch tok.Name.Local {
			case "subscriberProcessIdentifier":
				var data uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.SubscriberProcessIdentifier = data
			case "monitoredObjectType":
				var data uint16
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.MonitoredObjectType = data
			case "monitoredObjectInstanceNumber":
				var data uint32
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.MonitoredObjectInstanceNumber = data
			case "issueConfirmedNotifications":
				var data bool
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.IssueConfirmedNotifications = data
			case "lifetimeLength":
				var data uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.LifetimeLength = data
			case "lifetimeSeconds":
				var _encoded string
				if err := d.DecodeElement(&_encoded, &tok); err != nil {
					return err
				}
				_decoded, err := hex.DecodeString(_encoded)
				_len := len(_decoded)
				if err != nil {
					return err
				}
				m.LifetimeSeconds = utils.ByteArrayToInt8Array(_decoded[0:_len])
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

func (m *BACnetConfirmedServiceRequestSubscribeCOV) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.SubscriberProcessIdentifier, xml.StartElement{Name: xml.Name{Local: "subscriberProcessIdentifier"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.MonitoredObjectType, xml.StartElement{Name: xml.Name{Local: "monitoredObjectType"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.MonitoredObjectInstanceNumber, xml.StartElement{Name: xml.Name{Local: "monitoredObjectInstanceNumber"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.IssueConfirmedNotifications, xml.StartElement{Name: xml.Name{Local: "issueConfirmedNotifications"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.LifetimeLength, xml.StartElement{Name: xml.Name{Local: "lifetimeLength"}}); err != nil {
		return err
	}
	_encodedLifetimeSeconds := hex.EncodeToString(utils.Int8ArrayToByteArray(m.LifetimeSeconds))
	_encodedLifetimeSeconds = strings.ToUpper(_encodedLifetimeSeconds)
	if err := e.EncodeElement(_encodedLifetimeSeconds, xml.StartElement{Name: xml.Name{Local: "lifetimeSeconds"}}); err != nil {
		return err
	}
	return nil
}

func (m BACnetConfirmedServiceRequestSubscribeCOV) String() string {
	return string(m.Box("BACnetConfirmedServiceRequestSubscribeCOV", utils.DefaultWidth*2))
}

func (m BACnetConfirmedServiceRequestSubscribeCOV) Box(name string, width int) utils.AsciiBox {
	if name == "" {
		name = "BACnetConfirmedServiceRequestSubscribeCOV"
	}
	boxes := make([]utils.AsciiBox, 0)
	boxes = append(boxes, utils.BoxAnything("SubscriberProcessIdentifier", m.SubscriberProcessIdentifier, width-2))
	boxes = append(boxes, utils.BoxAnything("MonitoredObjectType", m.MonitoredObjectType, width-2))
	boxes = append(boxes, utils.BoxAnything("MonitoredObjectInstanceNumber", m.MonitoredObjectInstanceNumber, width-2))
	boxes = append(boxes, utils.BoxAnything("IssueConfirmedNotifications", m.IssueConfirmedNotifications, width-2))
	boxes = append(boxes, utils.BoxAnything("LifetimeLength", m.LifetimeLength, width-2))
	boxes = append(boxes, utils.BoxAnything("LifetimeSeconds", m.LifetimeSeconds, width-2))
	return utils.BoxBox(name, utils.AlignBoxes(boxes, width-2), 0)
}
