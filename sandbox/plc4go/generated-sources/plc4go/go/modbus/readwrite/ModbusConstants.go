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
    "math"
    "plc4x.apache.org/plc4go-modbus-driver/0.8.0/src/plc4go/spi"
    log "github.com/sirupsen/logrus"
)

// Constant values.
const MODBUSTCPDEFAULTPORT uint16 = 502;

type ModbusConstants struct {

}


func NewModbusConstants() spi.Message {
    return &ModbusConstants{}
}

func (m ModbusConstants) LengthInBits() uint16 {
    var lengthInBits uint16 = 0

    // Const Field (modbusTcpDefaultPort)
    lengthInBits += 16

    return lengthInBits
}

func (m ModbusConstants) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func ModbusConstantsParse(io spi.ReadBuffer) spi.Message {
    var startPos = io.GetPos()
    var curPos uint16

    // Const Field (modbusTcpDefaultPort)
    var modbusTcpDefaultPort uint16 = io.ReadUint16(16)
    if modbusTcpDefaultPort != MODBUSTCPDEFAULTPORT {
        throw new ParseException("Expected constant value " + ModbusConstants.MODBUSTCPDEFAULTPORT + " but got " + modbusTcpDefaultPort)
    }

    // Create the instance
    return NewModbusConstants()
}

