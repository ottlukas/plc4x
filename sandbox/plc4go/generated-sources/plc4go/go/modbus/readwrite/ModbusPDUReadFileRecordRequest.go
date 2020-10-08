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
package readwrite

import (
	"errors"
	"plc4x.apache.org/plc4go-modbus-driver/0.8.0/src/plc4go/spi"
	"reflect"
)

type ModbusPDUReadFileRecordRequest struct {
	items []ModbusPDUReadFileRecordRequestItem
	ModbusPDU
}

func (m ModbusPDUReadFileRecordRequest) initialize() spi.Message {
	return spi.Message(m)
}

func NewModbusPDUReadFileRecordRequest(items []ModbusPDUReadFileRecordRequestItem) ModbusPDUInitializer {
	return &ModbusPDUReadFileRecordRequest{items: items}
}

func (m ModbusPDUReadFileRecordRequest) LengthInBits() uint16 {
	var lengthInBits uint16 = m.ModbusPDU.LengthInBits()

	// Implicit Field (byteCount)
	lengthInBits += 8

	// Array field
	if len(m.items) > 0 {
		for _, element := range m.items {
			lengthInBits += element.LengthInBits()
		}
	}

	return lengthInBits
}

func (m ModbusPDUReadFileRecordRequest) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ModbusPDUReadFileRecordRequestParse(io spi.ReadBuffer) (ModbusPDUInitializer, error) {

	// Implicit Field (byteCount) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	var byteCount uint8 = io.ReadUint8(8)

	// Array field (items)
	var items []ModbusPDUReadFileRecordRequestItem
	// Length array
	_itemsLength := uint16(byteCount)
	_itemsEndPos := io.GetPos() + _itemsLength
	for io.GetPos() < _itemsEndPos {
		_message, _err := ModbusPDUReadFileRecordRequestItemParse(io)
		if _err != nil {
			return nil, errors.New("Error parsing 'items' field " + _err.Error())
		}
		var _item ModbusPDUReadFileRecordRequestItem
		_item, _ok := _message.(ModbusPDUReadFileRecordRequestItem)
		if !_ok {
			return nil, errors.New("Couldn't cast message of type " + reflect.TypeOf(_item).Name() + " to ModbusPDUReadFileRecordRequestItem")
		}
		items = append(items, _item)
	}

	// Create the instance
	return NewModbusPDUReadFileRecordRequest(items), nil
}
