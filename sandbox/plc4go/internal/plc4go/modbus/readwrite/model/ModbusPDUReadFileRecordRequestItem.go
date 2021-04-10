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
)

// The data-structure of this message
type ModbusPDUReadFileRecordRequestItem struct {
	referenceType uint8
	fileNumber    uint16
	recordNumber  uint16
	recordLength  uint16
}

// The corresponding interface
type IModbusPDUReadFileRecordRequestItem interface {
	spi.Message
	Serialize(io spi.WriteBuffer)
}

func NewModbusPDUReadFileRecordRequestItem(referenceType uint8, fileNumber uint16, recordNumber uint16, recordLength uint16) spi.Message {
	return &ModbusPDUReadFileRecordRequestItem{referenceType: referenceType, fileNumber: fileNumber, recordNumber: recordNumber, recordLength: recordLength}
}

func CastIModbusPDUReadFileRecordRequestItem(structType interface{}) IModbusPDUReadFileRecordRequestItem {
	castFunc := func(typ interface{}) IModbusPDUReadFileRecordRequestItem {
		if iModbusPDUReadFileRecordRequestItem, ok := typ.(IModbusPDUReadFileRecordRequestItem); ok {
			return iModbusPDUReadFileRecordRequestItem
		}
		return nil
	}
	return castFunc(structType)
}

func CastModbusPDUReadFileRecordRequestItem(structType interface{}) ModbusPDUReadFileRecordRequestItem {
	castFunc := func(typ interface{}) ModbusPDUReadFileRecordRequestItem {
		if sModbusPDUReadFileRecordRequestItem, ok := typ.(ModbusPDUReadFileRecordRequestItem); ok {
			return sModbusPDUReadFileRecordRequestItem
		}
		return ModbusPDUReadFileRecordRequestItem{}
	}
	return castFunc(structType)
}

func (m ModbusPDUReadFileRecordRequestItem) LengthInBits() uint16 {
	var lengthInBits uint16 = 0

	// Simple field (referenceType)
	lengthInBits += 8

	// Simple field (fileNumber)
	lengthInBits += 16

	// Simple field (recordNumber)
	lengthInBits += 16

	// Simple field (recordLength)
	lengthInBits += 16

	return lengthInBits
}

func (m ModbusPDUReadFileRecordRequestItem) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ModbusPDUReadFileRecordRequestItemParse(io spi.ReadBuffer) (spi.Message, error) {

	// Simple Field (referenceType)
	referenceType, _referenceTypeErr := io.ReadUint8(8)
	if _referenceTypeErr != nil {
		return nil, errors.New("Error parsing 'referenceType' field " + _referenceTypeErr.Error())
	}

	// Simple Field (fileNumber)
	fileNumber, _fileNumberErr := io.ReadUint16(16)
	if _fileNumberErr != nil {
		return nil, errors.New("Error parsing 'fileNumber' field " + _fileNumberErr.Error())
	}

	// Simple Field (recordNumber)
	recordNumber, _recordNumberErr := io.ReadUint16(16)
	if _recordNumberErr != nil {
		return nil, errors.New("Error parsing 'recordNumber' field " + _recordNumberErr.Error())
	}

	// Simple Field (recordLength)
	recordLength, _recordLengthErr := io.ReadUint16(16)
	if _recordLengthErr != nil {
		return nil, errors.New("Error parsing 'recordLength' field " + _recordLengthErr.Error())
	}

	// Create the instance
	return NewModbusPDUReadFileRecordRequestItem(referenceType, fileNumber, recordNumber, recordLength), nil
}

func (m ModbusPDUReadFileRecordRequestItem) Serialize(io spi.WriteBuffer) {

	// Simple Field (referenceType)
	referenceType := uint8(m.referenceType)
	io.WriteUint8(8, (referenceType))

	// Simple Field (fileNumber)
	fileNumber := uint16(m.fileNumber)
	io.WriteUint16(16, (fileNumber))

	// Simple Field (recordNumber)
	recordNumber := uint16(m.recordNumber)
	io.WriteUint16(16, (recordNumber))

	// Simple Field (recordLength)
	recordLength := uint16(m.recordLength)
	io.WriteUint16(16, (recordLength))

}
