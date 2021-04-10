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
package testutils

import (
    "encoding/hex"
    "errors"
    "fmt"
    "github.com/subchen/go-xmldom"
    "os"
    "plc4x.apache.org/plc4go-modbus-driver/v0/internal/plc4go/modbus/readwrite/model"
    "plc4x.apache.org/plc4go-modbus-driver/v0/internal/plc4go/spi"
    "plc4x.apache.org/plc4go-modbus-driver/v0/internal/plc4go/transports/test"
    "plc4x.apache.org/plc4go-modbus-driver/v0/internal/plc4go/utils"
    "plc4x.apache.org/plc4go-modbus-driver/v0/pkg/plc4go"
    api "plc4x.apache.org/plc4go-modbus-driver/v0/pkg/plc4go/model"
    "strconv"
    "testing"
    "time"
)

type DriverTestsuite struct {
    name             string
    driverName       string
    driverParameters map[string]string
    setupSteps       []TestStep
    teardownSteps    []TestStep
    testcases        []Testcase
}

func (m DriverTestsuite) Run(driverManager plc4go.PlcDriverManager, testcase Testcase) error {
    // Get a connection
    connectionChan := driverManager.GetConnection(m.driverName + ":test://hurz")
    connectionResult := <-connectionChan

    if connectionResult.Err != nil {
        return errors.New("error getting a connection: " + connectionResult.Err.Error())
    }

    // Run the setup steps
    for _, testStep := range m.setupSteps {
        m.ExecuteStep(connectionResult.Connection, &testcase, testStep)
    }

    // Run the actual scenario steps
    for _, testStep := range testcase.steps {
        m.ExecuteStep(connectionResult.Connection, &testcase, testStep)
    }

    // Run the teardown steps
    for _, testStep := range m.teardownSteps {
        m.ExecuteStep(connectionResult.Connection, &testcase, testStep)
    }

    return nil
}

func (m DriverTestsuite) ExecuteStep(connection plc4go.PlcConnection, testcase *Testcase, step TestStep) error {
    mc, ok := connection.(spi.TransportInstanceExposer)
    if !ok {
        return errors.New("couldn't access connections transport instance")
    }
    transportInstance := mc.GetTransportInstance()

    switch step.stepType {
    case StepType_API_REQUEST:
        switch step.payload.Name {
        case "TestReadRequest":
            // Assemble a read-request according to the information in the test xml
            rrb := connection.ReadRequestBuilder()
            for _, fieldNode := range step.payload.GetChild("fields").GetChildren("field") {
                fieldName := fieldNode.GetChild("name").Text
                fieldAddress := fieldNode.GetChild("address").Text
                rrb.AddItem(fieldName, fieldAddress)
            }
            readRequest, err := rrb.Build()
            if err != nil {
                return errors.New("Error creating read-request: " + err.Error())
            }

            // Execute the read-request and store the response-channel in the testcase.
            if testcase.readRequestResultChannel != nil {
                return errors.New("testcase read-request result channel already occupied")
            }
            testcase.readRequestResultChannel = readRequest.Execute()
        case "TestWriteRequest":
            wrb := connection.WriteRequestBuilder()
            for _, fieldNode := range step.payload.GetChild("fields").GetChildren("field") {
                fieldName := fieldNode.GetChild("name").Text
                fieldAddress := fieldNode.GetChild("address").Text
                fieldValue := fieldNode.GetChild("value").Text
                wrb.AddItem(fieldName, fieldAddress, fieldValue)
            }
            writeRequest, err := wrb.Build()
            if err != nil {
                return errors.New("Error creating write-request: " + err.Error())
            }
            if testcase.writeRequestResultChannel != nil {
                return errors.New("testcase write-request result channel already occupied")
            }
            testcase.writeRequestResultChannel = writeRequest.Execute()
        }
    case StepType_API_RESPONSE:
        switch step.payload.Name {
        case "DefaultPlcReadResponse":
            if testcase.readRequestResultChannel == nil {
                return errors.New("no read response expected")
            }
            readResponse := <- testcase.readRequestResultChannel
            // TODO: Implement this ...
            fmt.Println(readResponse)
        case "DefaultPlcWriteResponse":
            if testcase.writeRequestResultChannel == nil {
                return errors.New("no write response expected")
            }
            writeResponse := <- testcase.writeRequestResultChannel
            // TODO: Implement this ...
            fmt.Println(writeResponse)
        }
    case StepType_OUTGOING_PLC_MESSAGE:
        typeName := step.payload.Name
        payloadString := step.payload.XML()

        // Parse the xml into a real model
        message, err := model.ModbusXmlParserHelper{}.Parse(typeName, payloadString)
        if err != nil {
            return errors.New("error parsing xml: " + err.Error())
        }

        // Serialize the model into bytes
        ser, ok := message.(utils.Serializable)
        if !ok {
            return errors.New("error converting type into Serializable type: " + err.Error())
        }
        wb := utils.WriteBufferNew()
        err = ser.Serialize(*wb)
        if err != nil {
            return errors.New("error serializing message: " + err.Error())
        }
        expectedRawOutput := wb.GetBytes()

        // Read exactly this amount of bytes from the transport
        rawOutput, err := transportInstance.Read(uint32(len(expectedRawOutput)))
        if err != nil {
            return errors.New("error getting bytes from transport: " + err.Error())
        }

        // Compare the bytes read with the ones we expect
        for i := range expectedRawOutput {
            if expectedRawOutput[i] != rawOutput[i] {
                return errors.New("actual output doesn't match expected output")
            }
        }
        // If there's a difference, parse the input and display it to simplify debugging
    case StepType_OUTGOING_PLC_BYTES:
        // Read exactly this amount of bytes from the transport
        expectedRawInput, err := hex.DecodeString(step.payload.Text)
        if err != nil {
            return errors.New("error decoding hex-encoded byte data: " + err.Error())
        }
        rawInput, err := transportInstance.Read(uint32(len(expectedRawInput)))
        if err != nil {
            return errors.New("error getting bytes from transport: " + err.Error())
        }

        // Compare the bytes read with the ones we expect
        for i := range expectedRawInput {
            if expectedRawInput[i] != rawInput[i] {
                return errors.New("actual output doesn't match expected output")
            }
        }
        // If there's a difference, parse the input and display it to simplify debugging
    case StepType_INCOMING_PLC_MESSAGE:
        typeName := step.payload.Name
        payloadString := step.payload.XML()

        // Parse the xml into a real model
        message, err := model.ModbusXmlParserHelper{}.Parse(typeName, payloadString)
        if err != nil {
            return errors.New("error parsing xml: " + err.Error())
        }

        // Serialize the model into bytes
        ser, ok := message.(utils.Serializable)
        if !ok {
            return errors.New("error converting type into Serializable type: " + err.Error())
        }
        wb := utils.WriteBufferNew()
        err = ser.Serialize(*wb)
        if err != nil {
            return errors.New("error serializing message: " + err.Error())
        }

        // Send these bytes to the transport
        err = transportInstance.Write(wb.GetBytes())
        if err != nil {
            return errors.New("error writing data to transport: " + err.Error())
        }
    case StepType_INCOMING_PLC_BYTES:
        // Get the raw hex-data.
        rawInput, err := hex.DecodeString(step.payload.Text)
        if err != nil {
            return errors.New("error decoding hex-encoded byte data: " + err.Error())
        }

        // Send these bytes to the transport
        err = transportInstance.Write(rawInput)
        if err != nil {
            return errors.New("error writing data to transport: " + err.Error())
        }
    case StepType_DELAY:
        // Get the number of milliseconds
        delay, err := strconv.Atoi(step.payload.Text)
        if err != nil {
            return errors.New("invalid delay format: " + err.Error())
        }
        // Sleep for that long
        time.Sleep(time.Duration(delay))
    case StepType_TERMINATE:
        // Simply close the transport connection
        err := transportInstance.Close()
        if err != nil {
            return errors.New("error closing transport: " + err.Error())
        }
    }
    return nil
}

func (m DriverTestsuite) ParseXml(referenceXml *xmldom.Node, parserArguments []string)  {
    normalizeXml(referenceXml)
    //referenceSerialized := referenceXml.FirstChild().XML()
}

type Testcase struct {
    name  string
    steps                    []TestStep
    readRequestResultChannel <-chan api.PlcReadRequestResult
    writeRequestResultChannel <-chan api.PlcWriteRequestResult
}

type TestStep struct {
    name            string
    stepType        StepType
    parserArguments []string
    payload         xmldom.Node
}

type StepType uint8

const (
    StepType_OUTGOING_PLC_MESSAGE StepType = 0x01
    StepType_OUTGOING_PLC_BYTES   StepType = 0x02
    StepType_INCOMING_PLC_MESSAGE StepType = 0x03
    StepType_INCOMING_PLC_BYTES   StepType = 0x04
    StepType_API_REQUEST          StepType = 0x05
    StepType_API_RESPONSE         StepType = 0x06
    StepType_DELAY                StepType = 0x07
    StepType_TERMINATE            StepType = 0x08
)

func RunDriverTestsuite(t *testing.T, driver plc4go.PlcDriver, testPath string) {
    // Read the test-specification as XML file
    rootNode, err := ParseDriverTestsuiteXml(testPath)
    if err != nil {
        t.Error(err.Error())
        t.Fail()
        return
    }

    // Parse the contents of the test-specification
    testsuite, err := ParseDriverTestsuite(*rootNode)
    if err != nil {
        t.Error(err.Error())
        t.Fail()
        return
    }

    // Initialize the driver manager
    driverManager := plc4go.NewPlcDriverManager()
    driverManager.RegisterTransport(test.NewTestTransport())
    driverManager.RegisterDriver(driver)

    for _, testcase := range testsuite.testcases {
        testsuite.Run(driverManager, testcase)
    }
    // Execute the tests in the testsuite
    fmt.Printf(testsuite.name)
}

func ParseDriverTestsuiteXml(testPath string) (*xmldom.Node, error) {
    // Get the current working directory
    path, err := os.Getwd()
    if err != nil {
        return nil, err
    }

    // Check if the test-file is available
    info, err := os.Stat(path + "/../../../../" + testPath)
    if os.IsNotExist(err) {
        return nil, errors.New("test-File doesn't exist")
    }
    if info.IsDir() {
        return nil, errors.New("test-file refers to a directory")
    }

    // Open a reader for this file
    dat, err := os.Open(path + "/../../../../" + testPath)
    if err != nil {
        return nil, errors.New("error opening file")
    }

    // Read the xml
    node := xmldom.Must(xmldom.Parse(dat)).Root
    return node, nil
}

func ParseDriverTestsuite(node xmldom.Node) (*DriverTestsuite, error) {
    if node.Name != "driver-testsuite" {
        return nil, errors.New("invalid document structure")
    }
    var testsuiteName string
    var driverName string
    var setupSteps []TestStep
    var teardownSteps []TestStep
    var testcases []Testcase
    for _, childPtr := range node.Children {
        child := *childPtr
        if child.Name == "name" {
            testsuiteName = child.Text
        } else if child.Name == "driver-name" {
            driverName = child.Text
        } else if child.Name == "setup" {
            steps, err := ParseDriverTestsuiteSteps(child)
            if err != nil {
                return nil, errors.New("error parsing setup steps")
            }
            setupSteps = steps
        } else if child.Name == "teardown" {
            steps, err := ParseDriverTestsuiteSteps(child)
            if err != nil {
                return nil, errors.New("error teardown setup steps")
            }
            teardownSteps = steps
        } else if child.Name == "testcase" {
            testcaseName := child.FindOneByName("name").Text
            stepsNode := child.FindOneByName("steps")
            steps, err := ParseDriverTestsuiteSteps(*stepsNode)
            if err != nil {
                return nil, errors.New("error parsing testcase " + testcaseName + ": " + err.Error())
            }
            testcase := Testcase{
                name:  testcaseName,
                steps: steps,
            }
            testcases = append(testcases, testcase)
        } else {
            return nil, errors.New("invalid document structure")
        }
    }

    return &DriverTestsuite{
        name:          testsuiteName,
        driverName:    driverName,
        setupSteps:    setupSteps,
        teardownSteps: teardownSteps,
        testcases:     testcases,
    }, nil
}

func ParseDriverTestsuiteSteps(node xmldom.Node) ([]TestStep, error) {
    var testSteps []TestStep
    for _, step := range node.Children {
        name := step.GetAttributeValue("name")
        var stepType StepType
        switch step.Name {
        case "api-request":
            stepType = StepType_API_REQUEST
        case "api-response":
            stepType = StepType_API_RESPONSE
        case "outgoing-plc-message":
            stepType = StepType_OUTGOING_PLC_MESSAGE
        case "incoming-plc-message":
            stepType = StepType_INCOMING_PLC_MESSAGE
        case "outgoing-plc-bytes":
            stepType = StepType_OUTGOING_PLC_BYTES
        case "incoming-plc-bytes":
            stepType = StepType_INCOMING_PLC_BYTES
        case "delay":
            stepType = StepType_DELAY
        case "terminate":
            stepType = StepType_TERMINATE
        }
        var parserArguments []string
        var payload *xmldom.Node
        for _, childNode := range step.Children {
            if childNode.Name == "parser-arguments" {
                for _, parserArgumentNode := range childNode.Children {
                    parserArguments = append(parserArguments, parserArgumentNode.Text)
                }
            } else if payload == nil {
                payload = childNode
            } else {
                return nil, errors.New("test step can only contain a single payload element")
            }
        }
        if payload == nil {
            return nil, errors.New("missing payload element")
        }
        testSteps = append(testSteps, TestStep{
            name:            name,
            stepType:        stepType,
            parserArguments: parserArguments,
            payload:         *payload,
        })
    }
    return testSteps, nil
}
