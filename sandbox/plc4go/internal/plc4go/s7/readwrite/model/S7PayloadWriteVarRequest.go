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
	"errors"
	"plc4x.apache.org/plc4go-modbus-driver/0.8.0/internal/plc4go/spi"
	"reflect"
)

// The data-structure of this message
type S7PayloadWriteVarRequest struct {
	items []S7VarPayloadDataItem
	S7Payload
}

// The corresponding interface
type IS7PayloadWriteVarRequest interface {
	IS7Payload
	Serialize(io spi.WriteBuffer)
}

// Accessors for discriminator values.
func (m S7PayloadWriteVarRequest) ParameterParameterType() uint8 {
	return 0x05
}

func (m S7PayloadWriteVarRequest) MessageType() uint8 {
	return 0x01
}

func (m S7PayloadWriteVarRequest) initialize() spi.Message {
	return m
}

func NewS7PayloadWriteVarRequest(items []S7VarPayloadDataItem) S7PayloadInitializer {
	return &S7PayloadWriteVarRequest{items: items}
}

func (m S7PayloadWriteVarRequest) LengthInBits() uint16 {
	var lengthInBits uint16 = m.S7Payload.LengthInBits()

	// Array field
	if len(m.items) > 0 {
		for _, element := range m.items {
			lengthInBits += element.LengthInBits()
		}
	}

	return lengthInBits
}

func (m S7PayloadWriteVarRequest) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func S7PayloadWriteVarRequestParse(io spi.ReadBuffer, parameter S7Parameter) (S7PayloadInitializer, error) {

	// Array field (items)
	var items []S7VarPayloadDataItem
	// Count array
	{
		items := make([]S7VarPayloadDataItem, uint8(len(COUNT)))
		for curItem := uint16(0); curItem < uint16(uint8(len(COUNT))); curItem++ {
			lastItem := curItem == uint16(uint8(len(COUNT))-1)
			_message, _err := S7VarPayloadDataItemParse(io, bool(lastItem))
			if _err != nil {
				return nil, errors.New("Error parsing 'items' field " + _err.Error())
			}
			var _item S7VarPayloadDataItem
			_item, _ok := _message.(S7VarPayloadDataItem)
			if !_ok {
				return nil, errors.New("Couldn't cast message of type " + reflect.TypeOf(_item).Name() + " to S7VarPayloadDataItem")
			}
			items = append(items, _item)
		}
	}

	// Create the instance
	return NewS7PayloadWriteVarRequest(items), nil
}

func (m S7PayloadWriteVarRequest) Serialize(io spi.WriteBuffer) {
	serializeFunc := func(typ interface{}) {
		if _, ok := typ.(IS7PayloadWriteVarRequest); ok {

			// Array Field (items)
			if m.items != nil {
				itemCount := uint16(len(m.items))
				var curItem uint16 = 0
				for _, _element := range m.items {
					var lastItem bool = curItem == (itemCount - 1)
					_element.Serialize(io, lastItem)
					curItem++
				}
			}
		}
	}
	serializeFunc(m)
}
