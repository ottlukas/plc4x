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
package spi

import "time"

// If this function returns true, the message is forwarded to the message handler
type AcceptsMessage func(message interface{}) bool
// Function for handling the message, returns an error if anything goes wrong
type HandleMessage func(message interface{}) error

type MessageCodec interface {

    Connect() error
    Disconnect() error

    // Sends a given message
    Send(message interface{}) error
    // Wait for a given timespan for a message to come in, which returns 'true' for 'acceptMessage'
    // and is then forwarded to the 'handleMessage' function
    Expect(acceptsMessage AcceptsMessage, handleMessage HandleMessage, ttl time.Duration) error
    // A combination that sends a message first and then waits for a response
    SendRequest(message interface{}, acceptsMessage AcceptsMessage, handleMessage HandleMessage, ttl time.Duration) error

    GetDefaultIncomingMessageChannel() chan interface{}

}
