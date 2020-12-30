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
    "io"
    "github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
)

// The data-structure of this message
type MFuncPropStateReadReq struct {
    Parent *CEMI
    IMFuncPropStateReadReq
}

// The corresponding interface
type IMFuncPropStateReadReq interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *MFuncPropStateReadReq) MessageCode() uint8 {
    return 0xF9
}


func (m *MFuncPropStateReadReq) InitializeParent(parent *CEMI) {
}

func NewMFuncPropStateReadReq() *CEMI {
    child := &MFuncPropStateReadReq{
        Parent: NewCEMI(),
    }
    child.Parent.Child = child
    return child.Parent
}

func CastMFuncPropStateReadReq(structType interface{}) *MFuncPropStateReadReq {
    castFunc := func(typ interface{}) *MFuncPropStateReadReq {
        if casted, ok := typ.(MFuncPropStateReadReq); ok {
            return &casted
        }
        if casted, ok := typ.(*MFuncPropStateReadReq); ok {
            return casted
        }
        if casted, ok := typ.(CEMI); ok {
            return CastMFuncPropStateReadReq(casted.Child)
        }
        if casted, ok := typ.(*CEMI); ok {
            return CastMFuncPropStateReadReq(casted.Child)
        }
        return nil
    }
    return castFunc(structType)
}

func (m *MFuncPropStateReadReq) GetTypeName() string {
    return "MFuncPropStateReadReq"
}

func (m *MFuncPropStateReadReq) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    return lengthInBits
}

func (m *MFuncPropStateReadReq) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func MFuncPropStateReadReqParse(io *utils.ReadBuffer) (*CEMI, error) {

    // Create a partially initialized instance
    _child := &MFuncPropStateReadReq{
        Parent: &CEMI{},
    }
    _child.Parent.Child = _child
    return _child.Parent, nil
}

func (m *MFuncPropStateReadReq) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

        return nil
    }
    return m.Parent.SerializeParent(io, m, ser)
}

func (m *MFuncPropStateReadReq) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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

func (m *MFuncPropStateReadReq) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    return nil
}

