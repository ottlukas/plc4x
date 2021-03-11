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
type AdsDeviceNotificationRequest struct {
    Length uint32
    Stamps uint32
    AdsStampHeaders []*AdsStampHeader
    Parent *AdsData
    IAdsDeviceNotificationRequest
}

// The corresponding interface
type IAdsDeviceNotificationRequest interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *AdsDeviceNotificationRequest) CommandId() CommandId {
    return CommandId_ADS_DEVICE_NOTIFICATION
}

func (m *AdsDeviceNotificationRequest) Response() bool {
    return false
}


func (m *AdsDeviceNotificationRequest) InitializeParent(parent *AdsData) {
}

func NewAdsDeviceNotificationRequest(length uint32, stamps uint32, adsStampHeaders []*AdsStampHeader) *AdsData {
    child := &AdsDeviceNotificationRequest{
        Length: length,
        Stamps: stamps,
        AdsStampHeaders: adsStampHeaders,
        Parent: NewAdsData(),
    }
    child.Parent.Child = child
    return child.Parent
}

func CastAdsDeviceNotificationRequest(structType interface{}) *AdsDeviceNotificationRequest {
    castFunc := func(typ interface{}) *AdsDeviceNotificationRequest {
        if casted, ok := typ.(AdsDeviceNotificationRequest); ok {
            return &casted
        }
        if casted, ok := typ.(*AdsDeviceNotificationRequest); ok {
            return casted
        }
        if casted, ok := typ.(AdsData); ok {
            return CastAdsDeviceNotificationRequest(casted.Child)
        }
        if casted, ok := typ.(*AdsData); ok {
            return CastAdsDeviceNotificationRequest(casted.Child)
        }
        return nil
    }
    return castFunc(structType)
}

func (m *AdsDeviceNotificationRequest) GetTypeName() string {
    return "AdsDeviceNotificationRequest"
}

func (m *AdsDeviceNotificationRequest) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    // Simple field (length)
    lengthInBits += 32

    // Simple field (stamps)
    lengthInBits += 32

    // Array field
    if len(m.AdsStampHeaders) > 0 {
        for _, element := range m.AdsStampHeaders {
            lengthInBits += element.LengthInBits()
        }
    }

    return lengthInBits
}

func (m *AdsDeviceNotificationRequest) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func AdsDeviceNotificationRequestParse(io *utils.ReadBuffer) (*AdsData, error) {

    // Simple Field (length)
    length, _lengthErr := io.ReadUint32(32)
    if _lengthErr != nil {
        return nil, errors.New("Error parsing 'length' field " + _lengthErr.Error())
    }

    // Simple Field (stamps)
    stamps, _stampsErr := io.ReadUint32(32)
    if _stampsErr != nil {
        return nil, errors.New("Error parsing 'stamps' field " + _stampsErr.Error())
    }

    // Array field (adsStampHeaders)
    // Count array
    adsStampHeaders := make([]*AdsStampHeader, stamps)
    for curItem := uint16(0); curItem < uint16(stamps); curItem++ {
        _item, _err := AdsStampHeaderParse(io)
        if _err != nil {
            return nil, errors.New("Error parsing 'adsStampHeaders' field " + _err.Error())
        }
        adsStampHeaders[curItem] = _item
    }

    // Create a partially initialized instance
    _child := &AdsDeviceNotificationRequest{
        Length: length,
        Stamps: stamps,
        AdsStampHeaders: adsStampHeaders,
        Parent: &AdsData{},
    }
    _child.Parent.Child = _child
    return _child.Parent, nil
}

func (m *AdsDeviceNotificationRequest) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

    // Simple Field (length)
    length := uint32(m.Length)
    _lengthErr := io.WriteUint32(32, (length))
    if _lengthErr != nil {
        return errors.New("Error serializing 'length' field " + _lengthErr.Error())
    }

    // Simple Field (stamps)
    stamps := uint32(m.Stamps)
    _stampsErr := io.WriteUint32(32, (stamps))
    if _stampsErr != nil {
        return errors.New("Error serializing 'stamps' field " + _stampsErr.Error())
    }

    // Array Field (adsStampHeaders)
    if m.AdsStampHeaders != nil {
        for _, _element := range m.AdsStampHeaders {
            _elementErr := _element.Serialize(io)
            if _elementErr != nil {
                return errors.New("Error serializing 'adsStampHeaders' field " + _elementErr.Error())
            }
        }
    }

        return nil
    }
    return m.Parent.SerializeParent(io, m, ser)
}

func (m *AdsDeviceNotificationRequest) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
    var token xml.Token
    var err error
    token = start
    for {
        switch token.(type) {
        case xml.StartElement:
            tok := token.(xml.StartElement)
            switch tok.Name.Local {
            case "length":
                var data uint32
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.Length = data
            case "stamps":
                var data uint32
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.Stamps = data
            case "adsStampHeaders":
                var data []*AdsStampHeader
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.AdsStampHeaders = data
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

func (m *AdsDeviceNotificationRequest) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    if err := e.EncodeElement(m.Length, xml.StartElement{Name: xml.Name{Local: "length"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.Stamps, xml.StartElement{Name: xml.Name{Local: "stamps"}}); err != nil {
        return err
    }
    if err := e.EncodeToken(xml.StartElement{Name: xml.Name{Local: "adsStampHeaders"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.AdsStampHeaders, xml.StartElement{Name: xml.Name{Local: "adsStampHeaders"}}); err != nil {
        return err
    }
    if err := e.EncodeToken(xml.EndElement{Name: xml.Name{Local: "adsStampHeaders"}}); err != nil {
        return err
    }
    return nil
}

