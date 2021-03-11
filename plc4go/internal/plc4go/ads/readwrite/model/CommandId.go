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
    "github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
)

type CommandId uint16

type ICommandId interface {
    Serialize(io utils.WriteBuffer) error
}

const(
    CommandId_INVALID CommandId = 0x0000
    CommandId_ADS_READ_DEVICE_INFO CommandId = 0x0001
    CommandId_ADS_READ CommandId = 0x0002
    CommandId_ADS_WRITE CommandId = 0x0003
    CommandId_ADS_READ_STATE CommandId = 0x0004
    CommandId_ADS_WRITE_CONTROL CommandId = 0x0005
    CommandId_ADS_ADD_DEVICE_NOTIFICATION CommandId = 0x0006
    CommandId_ADS_DELETE_DEVICE_NOTIFICATION CommandId = 0x0007
    CommandId_ADS_DEVICE_NOTIFICATION CommandId = 0x0008
    CommandId_ADS_READ_WRITE CommandId = 0x0009
)

func CommandIdByValue(value uint16) CommandId {
    switch value {
        case 0x0000:
            return CommandId_INVALID
        case 0x0001:
            return CommandId_ADS_READ_DEVICE_INFO
        case 0x0002:
            return CommandId_ADS_READ
        case 0x0003:
            return CommandId_ADS_WRITE
        case 0x0004:
            return CommandId_ADS_READ_STATE
        case 0x0005:
            return CommandId_ADS_WRITE_CONTROL
        case 0x0006:
            return CommandId_ADS_ADD_DEVICE_NOTIFICATION
        case 0x0007:
            return CommandId_ADS_DELETE_DEVICE_NOTIFICATION
        case 0x0008:
            return CommandId_ADS_DEVICE_NOTIFICATION
        case 0x0009:
            return CommandId_ADS_READ_WRITE
    }
    return 0
}

func CommandIdByName(value string) CommandId {
    switch value {
    case "INVALID":
        return CommandId_INVALID
    case "ADS_READ_DEVICE_INFO":
        return CommandId_ADS_READ_DEVICE_INFO
    case "ADS_READ":
        return CommandId_ADS_READ
    case "ADS_WRITE":
        return CommandId_ADS_WRITE
    case "ADS_READ_STATE":
        return CommandId_ADS_READ_STATE
    case "ADS_WRITE_CONTROL":
        return CommandId_ADS_WRITE_CONTROL
    case "ADS_ADD_DEVICE_NOTIFICATION":
        return CommandId_ADS_ADD_DEVICE_NOTIFICATION
    case "ADS_DELETE_DEVICE_NOTIFICATION":
        return CommandId_ADS_DELETE_DEVICE_NOTIFICATION
    case "ADS_DEVICE_NOTIFICATION":
        return CommandId_ADS_DEVICE_NOTIFICATION
    case "ADS_READ_WRITE":
        return CommandId_ADS_READ_WRITE
    }
    return 0
}

func CastCommandId(structType interface{}) CommandId {
    castFunc := func(typ interface{}) CommandId {
        if sCommandId, ok := typ.(CommandId); ok {
            return sCommandId
        }
        return 0
    }
    return castFunc(structType)
}

func (m CommandId) LengthInBits() uint16 {
    return 16
}

func (m CommandId) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func CommandIdParse(io *utils.ReadBuffer) (CommandId, error) {
    val, err := io.ReadUint16(16)
    if err != nil {
        return 0, nil
    }
    return CommandIdByValue(val), nil
}

func (e CommandId) Serialize(io utils.WriteBuffer) error {
    err := io.WriteUint16(16, uint16(e))
    return err
}

func (e CommandId) String() string {
    switch e {
    case CommandId_INVALID:
        return "INVALID"
    case CommandId_ADS_READ_DEVICE_INFO:
        return "ADS_READ_DEVICE_INFO"
    case CommandId_ADS_READ:
        return "ADS_READ"
    case CommandId_ADS_WRITE:
        return "ADS_WRITE"
    case CommandId_ADS_READ_STATE:
        return "ADS_READ_STATE"
    case CommandId_ADS_WRITE_CONTROL:
        return "ADS_WRITE_CONTROL"
    case CommandId_ADS_ADD_DEVICE_NOTIFICATION:
        return "ADS_ADD_DEVICE_NOTIFICATION"
    case CommandId_ADS_DELETE_DEVICE_NOTIFICATION:
        return "ADS_DELETE_DEVICE_NOTIFICATION"
    case CommandId_ADS_DEVICE_NOTIFICATION:
        return "ADS_DEVICE_NOTIFICATION"
    case CommandId_ADS_READ_WRITE:
        return "ADS_READ_WRITE"
    }
    return ""
}
