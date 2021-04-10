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

type KnxLayer uint8

const (
	KnxLayer_TUNNEL_LINK_LAYER KnxLayer = 0x02
	KnxLayer_TUNNEL_RAW        KnxLayer = 0x04
	KnxLayer_TUNNEL_BUSMONITOR KnxLayer = 0x80
)

func KnxLayerParse(io spi.ReadBuffer) (KnxLayer, error) {
	// TODO: Implement ...
	return 0, nil
}

func (e KnxLayer) Serialize(io spi.WriteBuffer) {
	// TODO: Implement ...
}
