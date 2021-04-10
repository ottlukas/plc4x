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
    "plc4x.apache.org/plc4go-modbus-driver/v0/internal/plc4go/spi"
)

type ModbusXmlParserHelper struct {
}

func (m ModbusXmlParserHelper) Parse(typeName string, xmlString string) (spi.Message, error) {
    switch typeName {
    case "ModbusPDUWriteFileRecordRequestItem":
        var obj ModbusPDUWriteFileRecordRequestItem
        err := xml.Unmarshal([]byte(xmlString), &obj)
        if err != nil {
            return nil, errors.New("error unmarshalling xml: " + err.Error())
        }
        return obj, nil
    case "ModbusPDUReadFileRecordResponseItem":
        var obj ModbusPDUReadFileRecordResponseItem
        err := xml.Unmarshal([]byte(xmlString), &obj)
        if err != nil {
            return nil, errors.New("error unmarshalling xml: " + err.Error())
        }
        return obj, nil
    case "ModbusConstants":
        var obj ModbusConstants
        err := xml.Unmarshal([]byte(xmlString), &obj)
        if err != nil {
            return nil, errors.New("error unmarshalling xml: " + err.Error())
        }
        return obj, nil
    case "ModbusTcpADU":
        var obj ModbusTcpADU
        err := xml.Unmarshal([]byte(xmlString), &obj)
        if err != nil {
            return nil, errors.New("error unmarshalling xml: " + err.Error())
        }
        return obj, nil
    case "ModbusPDUWriteFileRecordResponseItem":
        var obj ModbusPDUWriteFileRecordResponseItem
        err := xml.Unmarshal([]byte(xmlString), &obj)
        if err != nil {
            return nil, errors.New("error unmarshalling xml: " + err.Error())
        }
        return obj, nil
    case "ModbusPDU":
        var obj ModbusPDU
        err := xml.Unmarshal([]byte(xmlString), &obj)
        if err != nil {
            return nil, errors.New("error unmarshalling xml: " + err.Error())
        }
        return obj, nil
    case "ModbusPDUReadFileRecordRequestItem":
        var obj ModbusPDUReadFileRecordRequestItem
        err := xml.Unmarshal([]byte(xmlString), &obj)
        if err != nil {
            return nil, errors.New("error unmarshalling xml: " + err.Error())
        }
        return obj, nil
    case "ModbusSerialADU":
        var obj ModbusSerialADU
        err := xml.Unmarshal([]byte(xmlString), &obj)
        if err != nil {
            return nil, errors.New("error unmarshalling xml: " + err.Error())
        }
        return obj, nil
    }
    return nil, errors.New("Unsupported type " + typeName)
}
