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
type AdsDeviceNotificationResponse struct {
    Parent *AdsData
    IAdsDeviceNotificationResponse
}

// The corresponding interface
type IAdsDeviceNotificationResponse interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *AdsDeviceNotificationResponse) CommandId() CommandId {
    return CommandId_ADS_DEVICE_NOTIFICATION
}

func (m *AdsDeviceNotificationResponse) Response() bool {
    return true
}


func (m *AdsDeviceNotificationResponse) InitializeParent(parent *AdsData) {
}

func NewAdsDeviceNotificationResponse() *AdsData {
    child := &AdsDeviceNotificationResponse{
        Parent: NewAdsData(),
    }
    child.Parent.Child = child
    return child.Parent
}

func CastAdsDeviceNotificationResponse(structType interface{}) *AdsDeviceNotificationResponse {
    castFunc := func(typ interface{}) *AdsDeviceNotificationResponse {
        if casted, ok := typ.(AdsDeviceNotificationResponse); ok {
            return &casted
        }
        if casted, ok := typ.(*AdsDeviceNotificationResponse); ok {
            return casted
        }
        if casted, ok := typ.(AdsData); ok {
            return CastAdsDeviceNotificationResponse(casted.Child)
        }
        if casted, ok := typ.(*AdsData); ok {
            return CastAdsDeviceNotificationResponse(casted.Child)
        }
        return nil
    }
    return castFunc(structType)
}

func (m *AdsDeviceNotificationResponse) GetTypeName() string {
    return "AdsDeviceNotificationResponse"
}

func (m *AdsDeviceNotificationResponse) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    return lengthInBits
}

func (m *AdsDeviceNotificationResponse) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func AdsDeviceNotificationResponseParse(io *utils.ReadBuffer) (*AdsData, error) {

    // Create a partially initialized instance
    _child := &AdsDeviceNotificationResponse{
        Parent: &AdsData{},
    }
    _child.Parent.Child = _child
    return _child.Parent, nil
}

func (m *AdsDeviceNotificationResponse) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

        return nil
    }
    return m.Parent.SerializeParent(io, m, ser)
}

func (m *AdsDeviceNotificationResponse) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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

func (m *AdsDeviceNotificationResponse) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    return nil
}

