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
package iec61131

import "plc4x.apache.org/plc4go-modbus-driver/0.8.0/internal/plc4go/spi/values"

type PlcBYTE struct {
	value uint8
	values.PlcSimpleValueAdapter
}

func NewPlcBYTE(value uint8) PlcBYTE {
	return PlcBYTE{
		value: value,
	}
}

func (m PlcBYTE) IsBoolean() bool {
	return true
}

func (m PlcBYTE) GetBooleanLength() uint32 {
	return 8
}

func (m PlcBYTE) GetBoolean() bool {
	return m.value&1 == 1
}

func (m PlcBYTE) GetBooleanAt(index uint32) bool {
	if index > 7 {
		return false
	}
	return m.value>>index&1 == 1
}

func (m PlcBYTE) GetBooleanArray() []bool {
	return []bool{m.value&1 == 1, m.value>>1&1 == 1,
		m.value>>2&1 == 1, m.value>>3&1 == 1,
		m.value>>4&1 == 1, m.value>>5&1 == 1,
		m.value>>6&1 == 1, m.value>>7&1 == 1}
}
