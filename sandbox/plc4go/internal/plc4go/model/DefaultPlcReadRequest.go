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
	"plc4x.apache.org/plc4go-modbus-driver/0.8.0/pkg/plc4go/model"
)

type DefaultPlcReadRequestBuilder struct {
	reader       PlcReader
	fieldHandler spi.PlcFieldHandler
	queries      map[string]string
	model.PlcReadRequestBuilder
}

func NewDefaultPlcReadRequestBuilder(fieldHandler spi.PlcFieldHandler, reader PlcReader) *DefaultPlcReadRequestBuilder {
	return &DefaultPlcReadRequestBuilder{
		reader:       reader,
		fieldHandler: fieldHandler,
		queries:      map[string]string{},
	}
}

func (m *DefaultPlcReadRequestBuilder) AddItem(name string, query string) {
	m.queries[name] = query
}

func (m *DefaultPlcReadRequestBuilder) Build() (model.PlcReadRequest, error) {
	fields := make(map[string]model.PlcField)
	for _, name := range m.queries {
		query := m.queries[name]
		field, err := m.fieldHandler.ParseQuery(query)
		if err != nil {
			return nil, errors.New("Error parsing query: " + query + ". Got error: " + err.Error())
		}
		fields[name] = field
	}
	return DefaultPlcReadRequest{
		fields: fields,
		reader: m.reader,
	}, nil
}

type DefaultPlcReadRequest struct {
	fields map[string]model.PlcField
	reader PlcReader
	model.PlcReadRequest
}

func (m DefaultPlcReadRequest) Execute() <-chan model.PlcReadRequestResult {
	return m.reader.Read()
}
