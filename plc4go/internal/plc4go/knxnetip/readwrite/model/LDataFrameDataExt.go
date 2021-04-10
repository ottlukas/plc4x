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
    "encoding/base64"
    "encoding/xml"
    "errors"
    "io"
    "github.com/apache/plc4x/plc4go/internal/plc4go/utils"
)

// The data-structure of this message
type LDataFrameDataExt struct {
    GroupAddress bool
    HopCount uint8
    ExtendedFrameFormat uint8
    SourceAddress *KnxAddress
    DestinationAddress []int8
    Apdu *Apdu
    Parent *LDataFrame
    ILDataFrameDataExt
}

// The corresponding interface
type ILDataFrameDataExt interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *LDataFrameDataExt) ExtendedFrame() bool {
    return true
}

func (m *LDataFrameDataExt) Polling() bool {
    return false
}


func (m *LDataFrameDataExt) InitializeParent(parent *LDataFrame, repeated bool, notAckFrame bool, priority CEMIPriority, acknowledgeRequested bool, errorFlag bool) {
    m.Parent.Repeated = repeated
    m.Parent.NotAckFrame = notAckFrame
    m.Parent.Priority = priority
    m.Parent.AcknowledgeRequested = acknowledgeRequested
    m.Parent.ErrorFlag = errorFlag
}

func NewLDataFrameDataExt(groupAddress bool, hopCount uint8, extendedFrameFormat uint8, sourceAddress *KnxAddress, destinationAddress []int8, apdu *Apdu, repeated bool, notAckFrame bool, priority CEMIPriority, acknowledgeRequested bool, errorFlag bool) *LDataFrame {
    child := &LDataFrameDataExt{
        GroupAddress: groupAddress,
        HopCount: hopCount,
        ExtendedFrameFormat: extendedFrameFormat,
        SourceAddress: sourceAddress,
        DestinationAddress: destinationAddress,
        Apdu: apdu,
        Parent: NewLDataFrame(repeated, notAckFrame, priority, acknowledgeRequested, errorFlag),
    }
    child.Parent.Child = child
    return child.Parent
}

func CastLDataFrameDataExt(structType interface{}) *LDataFrameDataExt {
    castFunc := func(typ interface{}) *LDataFrameDataExt {
        if casted, ok := typ.(LDataFrameDataExt); ok {
            return &casted
        }
        if casted, ok := typ.(*LDataFrameDataExt); ok {
            return casted
        }
        if casted, ok := typ.(LDataFrame); ok {
            return CastLDataFrameDataExt(casted.Child)
        }
        if casted, ok := typ.(*LDataFrame); ok {
            return CastLDataFrameDataExt(casted.Child)
        }
        return nil
    }
    return castFunc(structType)
}

func (m *LDataFrameDataExt) GetTypeName() string {
    return "LDataFrameDataExt"
}

func (m *LDataFrameDataExt) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    // Simple field (groupAddress)
    lengthInBits += 1

    // Simple field (hopCount)
    lengthInBits += 3

    // Simple field (extendedFrameFormat)
    lengthInBits += 4

    // Simple field (sourceAddress)
    lengthInBits += m.SourceAddress.LengthInBits()

    // Array field
    if len(m.DestinationAddress) > 0 {
        lengthInBits += 8 * uint16(len(m.DestinationAddress))
    }

    // Simple field (apdu)
    lengthInBits += m.Apdu.LengthInBits()

    return lengthInBits
}

func (m *LDataFrameDataExt) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func LDataFrameDataExtParse(io *utils.ReadBuffer) (*LDataFrame, error) {

    // Simple Field (groupAddress)
    groupAddress, _groupAddressErr := io.ReadBit()
    if _groupAddressErr != nil {
        return nil, errors.New("Error parsing 'groupAddress' field " + _groupAddressErr.Error())
    }

    // Simple Field (hopCount)
    hopCount, _hopCountErr := io.ReadUint8(3)
    if _hopCountErr != nil {
        return nil, errors.New("Error parsing 'hopCount' field " + _hopCountErr.Error())
    }

    // Simple Field (extendedFrameFormat)
    extendedFrameFormat, _extendedFrameFormatErr := io.ReadUint8(4)
    if _extendedFrameFormatErr != nil {
        return nil, errors.New("Error parsing 'extendedFrameFormat' field " + _extendedFrameFormatErr.Error())
    }

    // Simple Field (sourceAddress)
    sourceAddress, _sourceAddressErr := KnxAddressParse(io)
    if _sourceAddressErr != nil {
        return nil, errors.New("Error parsing 'sourceAddress' field " + _sourceAddressErr.Error())
    }

    // Array field (destinationAddress)
    // Count array
    destinationAddress := make([]int8, uint16(2))
    for curItem := uint16(0); curItem < uint16(uint16(2)); curItem++ {
        _item, _err := io.ReadInt8(8)
        if _err != nil {
            return nil, errors.New("Error parsing 'destinationAddress' field " + _err.Error())
        }
        destinationAddress[curItem] = _item
    }

    // Simple Field (apdu)
    apdu, _apduErr := ApduParse(io)
    if _apduErr != nil {
        return nil, errors.New("Error parsing 'apdu' field " + _apduErr.Error())
    }

    // Create a partially initialized instance
    _child := &LDataFrameDataExt{
        GroupAddress: groupAddress,
        HopCount: hopCount,
        ExtendedFrameFormat: extendedFrameFormat,
        SourceAddress: sourceAddress,
        DestinationAddress: destinationAddress,
        Apdu: apdu,
        Parent: &LDataFrame{},
    }
    _child.Parent.Child = _child
    return _child.Parent, nil
}

func (m *LDataFrameDataExt) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

    // Simple Field (groupAddress)
    groupAddress := bool(m.GroupAddress)
    _groupAddressErr := io.WriteBit((groupAddress))
    if _groupAddressErr != nil {
        return errors.New("Error serializing 'groupAddress' field " + _groupAddressErr.Error())
    }

    // Simple Field (hopCount)
    hopCount := uint8(m.HopCount)
    _hopCountErr := io.WriteUint8(3, (hopCount))
    if _hopCountErr != nil {
        return errors.New("Error serializing 'hopCount' field " + _hopCountErr.Error())
    }

    // Simple Field (extendedFrameFormat)
    extendedFrameFormat := uint8(m.ExtendedFrameFormat)
    _extendedFrameFormatErr := io.WriteUint8(4, (extendedFrameFormat))
    if _extendedFrameFormatErr != nil {
        return errors.New("Error serializing 'extendedFrameFormat' field " + _extendedFrameFormatErr.Error())
    }

    // Simple Field (sourceAddress)
    _sourceAddressErr := m.SourceAddress.Serialize(io)
    if _sourceAddressErr != nil {
        return errors.New("Error serializing 'sourceAddress' field " + _sourceAddressErr.Error())
    }

    // Array Field (destinationAddress)
    if m.DestinationAddress != nil {
        for _, _element := range m.DestinationAddress {
            _elementErr := io.WriteInt8(8, _element)
            if _elementErr != nil {
                return errors.New("Error serializing 'destinationAddress' field " + _elementErr.Error())
            }
        }
    }

    // Simple Field (apdu)
    _apduErr := m.Apdu.Serialize(io)
    if _apduErr != nil {
        return errors.New("Error serializing 'apdu' field " + _apduErr.Error())
    }

        return nil
    }
    return m.Parent.SerializeParent(io, m, ser)
}

func (m *LDataFrameDataExt) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
    var token xml.Token
    var err error
    token = start
    for {
        switch token.(type) {
        case xml.StartElement:
            tok := token.(xml.StartElement)
            switch tok.Name.Local {
            case "groupAddress":
                var data bool
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.GroupAddress = data
            case "hopCount":
                var data uint8
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.HopCount = data
            case "extendedFrameFormat":
                var data uint8
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.ExtendedFrameFormat = data
            case "sourceAddress":
                var data *KnxAddress
                if err := d.DecodeElement(data, &tok); err != nil {
                    return err
                }
                m.SourceAddress = data
            case "destinationAddress":
                var _encoded string
                if err := d.DecodeElement(&_encoded, &tok); err != nil {
                    return err
                }
                _decoded := make([]byte, base64.StdEncoding.DecodedLen(len(_encoded)))
                _len, err := base64.StdEncoding.Decode(_decoded, []byte(_encoded))
                if err != nil {
                    return err
                }
                m.DestinationAddress = utils.ByteArrayToInt8Array(_decoded[0:_len])
            case "apdu":
                var dt *Apdu
                if err := d.DecodeElement(&dt, &tok); err != nil {
                    return err
                }
                m.Apdu = dt
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

func (m *LDataFrameDataExt) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    if err := e.EncodeElement(m.GroupAddress, xml.StartElement{Name: xml.Name{Local: "groupAddress"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.HopCount, xml.StartElement{Name: xml.Name{Local: "hopCount"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.ExtendedFrameFormat, xml.StartElement{Name: xml.Name{Local: "extendedFrameFormat"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.SourceAddress, xml.StartElement{Name: xml.Name{Local: "sourceAddress"}}); err != nil {
        return err
    }
    _encodedDestinationAddress := make([]byte, base64.StdEncoding.EncodedLen(len(m.DestinationAddress)))
    base64.StdEncoding.Encode(_encodedDestinationAddress, utils.Int8ArrayToByteArray(m.DestinationAddress))
    if err := e.EncodeElement(_encodedDestinationAddress, xml.StartElement{Name: xml.Name{Local: "destinationAddress"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.Apdu, xml.StartElement{Name: xml.Name{Local: "apdu"}}); err != nil {
        return err
    }
    return nil
}

