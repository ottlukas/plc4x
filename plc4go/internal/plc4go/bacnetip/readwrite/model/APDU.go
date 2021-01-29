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
    "io"
    "github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
    "reflect"
    "strings"
)

// The data-structure of this message
type APDU struct {
    Child IAPDUChild
    IAPDU
    IAPDUParent
}

// The corresponding interface
type IAPDU interface {
    ApduType() uint8
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

type IAPDUParent interface {
    SerializeParent(io utils.WriteBuffer, child IAPDU, serializeChildFunction func() error) error
    GetTypeName() string
}

type IAPDUChild interface {
    Serialize(io utils.WriteBuffer) error
    InitializeParent(parent *APDU)
    GetTypeName() string
    IAPDU
}

func NewAPDU() *APDU {
    return &APDU{}
}

func CastAPDU(structType interface{}) *APDU {
    castFunc := func(typ interface{}) *APDU {
        if casted, ok := typ.(APDU); ok {
            return &casted
        }
        if casted, ok := typ.(*APDU); ok {
            return casted
        }
        return nil
    }
    return castFunc(structType)
}

func (m *APDU) GetTypeName() string {
    return "APDU"
}

func (m *APDU) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    // Discriminator Field (apduType)
    lengthInBits += 4

    // Length of sub-type elements will be added by sub-type...
    lengthInBits += m.Child.LengthInBits()

    return lengthInBits
}

func (m *APDU) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func APDUParse(io *utils.ReadBuffer, apduLength uint16) (*APDU, error) {

    // Discriminator Field (apduType) (Used as input to a switch field)
    apduType, _apduTypeErr := io.ReadUint8(4)
    if _apduTypeErr != nil {
        return nil, errors.New("Error parsing 'apduType' field " + _apduTypeErr.Error())
    }

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
    var _parent *APDU
    var typeSwitchError error
    switch {
    case apduType == 0x0:
        _parent, typeSwitchError = APDUConfirmedRequestParse(io, apduLength)
    case apduType == 0x1:
        _parent, typeSwitchError = APDUUnconfirmedRequestParse(io, apduLength)
    case apduType == 0x2:
        _parent, typeSwitchError = APDUSimpleAckParse(io)
    case apduType == 0x3:
        _parent, typeSwitchError = APDUComplexAckParse(io)
    case apduType == 0x4:
        _parent, typeSwitchError = APDUSegmentAckParse(io)
    case apduType == 0x5:
        _parent, typeSwitchError = APDUErrorParse(io)
    case apduType == 0x6:
        _parent, typeSwitchError = APDURejectParse(io)
    case apduType == 0x7:
        _parent, typeSwitchError = APDUAbortParse(io)
    }
    if typeSwitchError != nil {
        return nil, errors.New("Error parsing sub-type for type-switch. " + typeSwitchError.Error())
    }

    // Finish initializing
    _parent.Child.InitializeParent(_parent)
    return _parent, nil
}

func (m *APDU) Serialize(io utils.WriteBuffer) error {
    return m.Child.Serialize(io)
}

func (m *APDU) SerializeParent(io utils.WriteBuffer, child IAPDU, serializeChildFunction func() error) error {

    // Discriminator Field (apduType) (Used as input to a switch field)
    apduType := uint8(child.ApduType())
    _apduTypeErr := io.WriteUint8(4, (apduType))
    if _apduTypeErr != nil {
        return errors.New("Error serializing 'apduType' field " + _apduTypeErr.Error())
    }

    // Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
    _typeSwitchErr := serializeChildFunction()
    if _typeSwitchErr != nil {
        return errors.New("Error serializing sub-type field " + _typeSwitchErr.Error())
    }

    return nil
}

func (m *APDU) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
            default:
                switch start.Attr[0].Value {
                    case "org.apache.plc4x.java.bacnetip.readwrite.APDUConfirmedRequest":
                        var dt *APDUConfirmedRequest
                        if m.Child != nil {
                            dt = m.Child.(*APDUConfirmedRequest)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.bacnetip.readwrite.APDUUnconfirmedRequest":
                        var dt *APDUUnconfirmedRequest
                        if m.Child != nil {
                            dt = m.Child.(*APDUUnconfirmedRequest)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.bacnetip.readwrite.APDUSimpleAck":
                        var dt *APDUSimpleAck
                        if m.Child != nil {
                            dt = m.Child.(*APDUSimpleAck)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.bacnetip.readwrite.APDUComplexAck":
                        var dt *APDUComplexAck
                        if m.Child != nil {
                            dt = m.Child.(*APDUComplexAck)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.bacnetip.readwrite.APDUSegmentAck":
                        var dt *APDUSegmentAck
                        if m.Child != nil {
                            dt = m.Child.(*APDUSegmentAck)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.bacnetip.readwrite.APDUError":
                        var dt *APDUError
                        if m.Child != nil {
                            dt = m.Child.(*APDUError)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.bacnetip.readwrite.APDUReject":
                        var dt *APDUReject
                        if m.Child != nil {
                            dt = m.Child.(*APDUReject)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.bacnetip.readwrite.APDUAbort":
                        var dt *APDUAbort
                        if m.Child != nil {
                            dt = m.Child.(*APDUAbort)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                }
            }
        }
    }
}

func (m *APDU) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    className := reflect.TypeOf(m.Child).String()
    className = "org.apache.plc4x.java.bacnetip.readwrite." + className[strings.LastIndex(className, ".") + 1:]
    if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
            {Name: xml.Name{Local: "className"}, Value: className},
        }}); err != nil {
        return err
    }
    marshaller, ok := m.Child.(xml.Marshaler)
    if !ok {
        return errors.New("child is not castable to Marshaler")
    }
    marshaller.MarshalXML(e, start)
    if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
        return err
    }
    return nil
}

