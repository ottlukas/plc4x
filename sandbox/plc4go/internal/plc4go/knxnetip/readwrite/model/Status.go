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

import "plc4x.apache.org/plc4go-modbus-driver/0.8.0/internal/plc4go/spi"

type Status uint8

const (
	Status_NO_ERROR                        Status = 0x00
	Status_PROTOCOL_TYPE_NOT_SUPPORTED     Status = 0x01
	Status_UNSUPPORTED_PROTOCOL_VERSION    Status = 0x02
	Status_OUT_OF_ORDER_SEQUENCE_NUMBER    Status = 0x04
	Status_INVALID_CONNECTION_ID           Status = 0x21
	Status_CONNECTION_TYPE_NOT_SUPPORTED   Status = 0x22
	Status_CONNECTION_OPTION_NOT_SUPPORTED Status = 0x23
	Status_NO_MORE_CONNECTIONS             Status = 0x24
	Status_NO_MORE_UNIQUE_CONNECTIONS      Status = 0x25
	Status_DATA_CONNECTION                 Status = 0x26
	Status_KNX_CONNECTION                  Status = 0x27
	Status_TUNNELLING_LAYER_NOT_SUPPORTED  Status = 0x29
)

func StatusParse(io spi.ReadBuffer) (Status, error) {
	// TODO: Implement ...
	return 0, nil
}

func (e Status) Serialize(io spi.WriteBuffer) {
	// TODO: Implement ...
}
