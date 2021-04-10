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
package knxnetip

import (
    "bytes"
    "errors"
    "fmt"
    driverModel "github.com/apache/plc4x/plc4go/internal/plc4go/knxnetip/readwrite/model"
    "github.com/apache/plc4x/plc4go/internal/plc4go/spi"
    "github.com/apache/plc4x/plc4go/internal/plc4go/transports"
    "github.com/apache/plc4x/plc4go/internal/plc4go/transports/udp"
    "github.com/apache/plc4x/plc4go/internal/plc4go/utils"
    "github.com/apache/plc4x/plc4go/pkg/plc4go/model"
    "net"
    "net/url"
    "time"
)

type KnxNetIpDiscoverer struct {
    messageCodec spi.MessageCodec
    spi.Discoverer
}

func NewKnxNetIpDiscoverer() *KnxNetIpDiscoverer {
    return &KnxNetIpDiscoverer{
    }
}

func (d *KnxNetIpDiscoverer) Discover(callback func(event model.PlcDiscoveryEvent)) error {
    udpTransport := udp.NewUdpTransport()

    // Create a connection string for the KNX broadcast discovery address.
    connectionUrl, err := url.Parse("udp://224.0.23.12:3671")
    if err != nil {
        return err
    }

    interfaces, err := net.Interfaces()
    if err != nil {
        return err
    }

    var tranportInstances []transports.TransportInstance
    // Iterate over all network devices of this system.
    for _, interf := range interfaces {
        addrs, err := interf.Addrs()
        if err != nil {
            return err
        }
        // Iterate over all addresses the current interface has configured
        // For KNX we're only interested in IPv4 addresses, as it doesn't
        // seem to work with IPv6.
        for _, addr := range addrs {
            var ipv4Addr net.IP
            switch addr.(type) {
            // If the device is configured to communicate with a subnet
            case *net.IPNet:
                ipv4Addr = addr.(*net.IPNet).IP.To4()

            // If the device is configured for a point-to-point connection
            case *net.IPAddr:
                ipv4Addr = addr.(*net.IPAddr).IP.To4()
            }

            // If we found an IPv4 address and this is not a loopback address,
            // add it to the list of devices we will open ports and send discovery
            // messages from.
            if ipv4Addr != nil && !ipv4Addr.IsLoopback() {
                // Create a new "connection" (Actually open a local udp socket and target outgoing packets to that address)
                transportInstance, err :=
                    udpTransport.CreateTransportInstanceForLocalAddress(*connectionUrl, nil,
                        &net.UDPAddr{IP: ipv4Addr, Port: 0})
                if err != nil {
                    return err
                }
                err = transportInstance.Connect()
                if err != nil {
                    continue
                }

                tranportInstances = append(tranportInstances, transportInstance)
            }
        }
    }

    if len(tranportInstances) > 0 {
        for _, transportInstance := range tranportInstances {
            // Create a codec for sending and receiving messages.
            codec := NewKnxNetIpMessageCodec(transportInstance, nil)

            // Cast to the UDP transport instance so we can access information on the local port.
            udpTransportInstance, ok := transportInstance.(*udp.UdpTransportInstance)
            if !ok {
                return errors.New("couldn't cast transport instance to UDP transport instance")
            }
            localAddress := udpTransportInstance.LocalAddress

            // Prepare the discovery packet data
            searchRequestMessage := driverModel.NewSearchRequest(driverModel.NewHPAIDiscoveryEndpoint(
                driverModel.HostProtocolCode_IPV4_UDP,
                driverModel.NewIPAddress(utils.ByteArrayToInt8Array(localAddress.IP.To4())),
                uint16(localAddress.Port)))

            err = codec.Send(searchRequestMessage)
            if err != nil {
                return err
            }

            // Register an expected response
            check := func(response interface{}) (bool, bool) {
                searchResponseMessage := driverModel.CastSearchResponse(response)
                // As we can expect multiple responses, we always tell the codec to keep this selector active
                return searchResponseMessage != nil, true
            }
            // Register a callback to handle the response
            responseChan := codec.Expect(check)
            go func() {
                // As we can receive multiple responses, stay in a loop
                for {
                    select {
                    case response := <-responseChan:
                        searchResponse := driverModel.CastSearchResponse(response)

                        addr := searchResponse.HpaiControlEndpoint.IpAddress.Addr
                        remoteUrl, err := url.Parse(fmt.Sprintf("udp://%d.%d.%d.%d:%d",
                            uint8(addr[0]), uint8(addr[1]), uint8(addr[2]), uint8(addr[3]), searchResponse.HpaiControlEndpoint.IpPort))
                        if err != nil {
                            fmt.Print(err.Error())
                        }
                        deviceName := string(bytes.Trim(utils.Int8ArrayToByteArray(
                            searchResponse.DibDeviceInfo.DeviceFriendlyName), "\x00"))
                        discoveryEvent := model.NewPlcDiscoveryEvent(
                            "knxnet-ip", "udp", *remoteUrl, nil, deviceName)
                        // Pass the event back to the callback
                        callback(discoveryEvent)
                    case <-time.After(time.Second * 5):
                        // After timeout of 5 seconds, stop waiting for search-responses
                        _ = codec.RemoveExpectation(check)
                    }
                }
            }()
        }
    }
    return nil
}
