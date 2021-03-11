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
type AmsPacket struct {
	TargetAmsNetId *AmsNetId
	TargetAmsPort  uint16
	SourceAmsNetId *AmsNetId
	SourceAmsPort  uint16
	CommandId      CommandId
	State          *State
	ErrorCode      uint32
	InvokeId       uint32
	Data           *AdsData
	IAmsPacket
}

// The corresponding interface
type IAmsPacket interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
}

func NewAmsPacket(targetAmsNetId *AmsNetId, targetAmsPort uint16, sourceAmsNetId *AmsNetId, sourceAmsPort uint16, commandId CommandId, state *State, errorCode uint32, invokeId uint32, data *AdsData) *AmsPacket {
	return &AmsPacket{TargetAmsNetId: targetAmsNetId, TargetAmsPort: targetAmsPort, SourceAmsNetId: sourceAmsNetId, SourceAmsPort: sourceAmsPort, CommandId: commandId, State: state, ErrorCode: errorCode, InvokeId: invokeId, Data: data}
}

func CastAmsPacket(structType interface{}) *AmsPacket {
	castFunc := func(typ interface{}) *AmsPacket {
		if casted, ok := typ.(AmsPacket); ok {
			return &casted
		}
		if casted, ok := typ.(*AmsPacket); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *AmsPacket) GetTypeName() string {
	return "AmsPacket"
}

func (m *AmsPacket) LengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Simple field (targetAmsNetId)
	lengthInBits += m.TargetAmsNetId.LengthInBits()

	// Simple field (targetAmsPort)
	lengthInBits += 16

	// Simple field (sourceAmsNetId)
	lengthInBits += m.SourceAmsNetId.LengthInBits()

	// Simple field (sourceAmsPort)
	lengthInBits += 16

	// Enum Field (commandId)
	lengthInBits += 16

	// Simple field (state)
	lengthInBits += m.State.LengthInBits()

	// Implicit Field (length)
	lengthInBits += 32

	// Simple field (errorCode)
	lengthInBits += 32

	// Simple field (invokeId)
	lengthInBits += 32

	// Simple field (data)
	lengthInBits += m.Data.LengthInBits()

	return lengthInBits
}

func (m *AmsPacket) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func AmsPacketParse(io *utils.ReadBuffer) (*AmsPacket, error) {

	// Simple Field (targetAmsNetId)
	targetAmsNetId, _targetAmsNetIdErr := AmsNetIdParse(io)
	if _targetAmsNetIdErr != nil {
		return nil, errors.New("Error parsing 'targetAmsNetId' field " + _targetAmsNetIdErr.Error())
	}

	// Simple Field (targetAmsPort)
	targetAmsPort, _targetAmsPortErr := io.ReadUint16(16)
	if _targetAmsPortErr != nil {
		return nil, errors.New("Error parsing 'targetAmsPort' field " + _targetAmsPortErr.Error())
	}

	// Simple Field (sourceAmsNetId)
	sourceAmsNetId, _sourceAmsNetIdErr := AmsNetIdParse(io)
	if _sourceAmsNetIdErr != nil {
		return nil, errors.New("Error parsing 'sourceAmsNetId' field " + _sourceAmsNetIdErr.Error())
	}

	// Simple Field (sourceAmsPort)
	sourceAmsPort, _sourceAmsPortErr := io.ReadUint16(16)
	if _sourceAmsPortErr != nil {
		return nil, errors.New("Error parsing 'sourceAmsPort' field " + _sourceAmsPortErr.Error())
	}

	// Enum field (commandId)
	commandId, _commandIdErr := CommandIdParse(io)
	if _commandIdErr != nil {
		return nil, errors.New("Error parsing 'commandId' field " + _commandIdErr.Error())
	}

	// Simple Field (state)
	state, _stateErr := StateParse(io)
	if _stateErr != nil {
		return nil, errors.New("Error parsing 'state' field " + _stateErr.Error())
	}

	// Implicit Field (length) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	_, _lengthErr := io.ReadUint32(32)
	if _lengthErr != nil {
		return nil, errors.New("Error parsing 'length' field " + _lengthErr.Error())
	}

	// Simple Field (errorCode)
	errorCode, _errorCodeErr := io.ReadUint32(32)
	if _errorCodeErr != nil {
		return nil, errors.New("Error parsing 'errorCode' field " + _errorCodeErr.Error())
	}

	// Simple Field (invokeId)
	invokeId, _invokeIdErr := io.ReadUint32(32)
	if _invokeIdErr != nil {
		return nil, errors.New("Error parsing 'invokeId' field " + _invokeIdErr.Error())
	}

	// Simple Field (data)
	data, _dataErr := AdsDataParse(io, commandId, state.Response)
	if _dataErr != nil {
		return nil, errors.New("Error parsing 'data' field " + _dataErr.Error())
	}

	// Create the instance
	return NewAmsPacket(targetAmsNetId, targetAmsPort, sourceAmsNetId, sourceAmsPort, commandId, state, errorCode, invokeId, data), nil
}

func (m *AmsPacket) Serialize(io utils.WriteBuffer) error {

	// Simple Field (targetAmsNetId)
	_targetAmsNetIdErr := m.TargetAmsNetId.Serialize(io)
	if _targetAmsNetIdErr != nil {
		return errors.New("Error serializing 'targetAmsNetId' field " + _targetAmsNetIdErr.Error())
	}

	// Simple Field (targetAmsPort)
	targetAmsPort := uint16(m.TargetAmsPort)
	_targetAmsPortErr := io.WriteUint16(16, (targetAmsPort))
	if _targetAmsPortErr != nil {
		return errors.New("Error serializing 'targetAmsPort' field " + _targetAmsPortErr.Error())
	}

	// Simple Field (sourceAmsNetId)
	_sourceAmsNetIdErr := m.SourceAmsNetId.Serialize(io)
	if _sourceAmsNetIdErr != nil {
		return errors.New("Error serializing 'sourceAmsNetId' field " + _sourceAmsNetIdErr.Error())
	}

	// Simple Field (sourceAmsPort)
	sourceAmsPort := uint16(m.SourceAmsPort)
	_sourceAmsPortErr := io.WriteUint16(16, (sourceAmsPort))
	if _sourceAmsPortErr != nil {
		return errors.New("Error serializing 'sourceAmsPort' field " + _sourceAmsPortErr.Error())
	}

	// Enum field (commandId)
	commandId := CastCommandId(m.CommandId)
	_commandIdErr := commandId.Serialize(io)
	if _commandIdErr != nil {
		return errors.New("Error serializing 'commandId' field " + _commandIdErr.Error())
	}

	// Simple Field (state)
	_stateErr := m.State.Serialize(io)
	if _stateErr != nil {
		return errors.New("Error serializing 'state' field " + _stateErr.Error())
	}

	// Implicit Field (length) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	length := uint32(m.Data.LengthInBytes())
	_lengthErr := io.WriteUint32(32, (length))
	if _lengthErr != nil {
		return errors.New("Error serializing 'length' field " + _lengthErr.Error())
	}

	// Simple Field (errorCode)
	errorCode := uint32(m.ErrorCode)
	_errorCodeErr := io.WriteUint32(32, (errorCode))
	if _errorCodeErr != nil {
		return errors.New("Error serializing 'errorCode' field " + _errorCodeErr.Error())
	}

	// Simple Field (invokeId)
	invokeId := uint32(m.InvokeId)
	_invokeIdErr := io.WriteUint32(32, (invokeId))
	if _invokeIdErr != nil {
		return errors.New("Error serializing 'invokeId' field " + _invokeIdErr.Error())
	}

	// Simple Field (data)
	_dataErr := m.Data.Serialize(io)
	if _dataErr != nil {
		return errors.New("Error serializing 'data' field " + _dataErr.Error())
	}

	return nil
}

func (m *AmsPacket) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "targetAmsNetId":
				var data *AmsNetId
				if err := d.DecodeElement(data, &tok); err != nil {
					return err
				}
				m.TargetAmsNetId = data
			case "targetAmsPort":
				var data uint16
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.TargetAmsPort = data
			case "sourceAmsNetId":
				var data *AmsNetId
				if err := d.DecodeElement(data, &tok); err != nil {
					return err
				}
				m.SourceAmsNetId = data
			case "sourceAmsPort":
				var data uint16
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.SourceAmsPort = data
			case "commandId":
				var data CommandId
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.CommandId = data
			case "state":
				var data *State
				if err := d.DecodeElement(data, &tok); err != nil {
					return err
				}
				m.State = data
			case "errorCode":
				var data uint32
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.ErrorCode = data
			case "invokeId":
				var data uint32
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.InvokeId = data
			case "data":
				var dt *AdsData
				if err := d.DecodeElement(&dt, &tok); err != nil {
					return err
				}
				m.Data = dt
			}
		}
	}
}

func (m *AmsPacket) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	className := "org.apache.plc4x.java.ads.readwrite.AmsPacket"
	if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
		{Name: xml.Name{Local: "className"}, Value: className},
	}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.TargetAmsNetId, xml.StartElement{Name: xml.Name{Local: "targetAmsNetId"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.TargetAmsPort, xml.StartElement{Name: xml.Name{Local: "targetAmsPort"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.SourceAmsNetId, xml.StartElement{Name: xml.Name{Local: "sourceAmsNetId"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.SourceAmsPort, xml.StartElement{Name: xml.Name{Local: "sourceAmsPort"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.CommandId, xml.StartElement{Name: xml.Name{Local: "commandId"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.State, xml.StartElement{Name: xml.Name{Local: "state"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.ErrorCode, xml.StartElement{Name: xml.Name{Local: "errorCode"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.InvokeId, xml.StartElement{Name: xml.Name{Local: "invokeId"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.Data, xml.StartElement{Name: xml.Name{Local: "data"}}); err != nil {
		return err
	}
	if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
		return err
	}
	return nil
}
