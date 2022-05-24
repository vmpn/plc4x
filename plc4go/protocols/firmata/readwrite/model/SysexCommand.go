/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// SysexCommand is the data-structure of this message
type SysexCommand struct {
	Child ISysexCommandChild
}

// ISysexCommand is the corresponding interface of SysexCommand
type ISysexCommand interface {
	// GetCommandType returns CommandType (discriminator field)
	GetCommandType() uint8
	// GetResponse returns Response (discriminator field)
	GetResponse() bool
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

type ISysexCommandParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child ISysexCommand, serializeChildFunction func() error) error
	GetTypeName() string
}

type ISysexCommandChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *SysexCommand)
	GetParent() *SysexCommand

	GetTypeName() string
	ISysexCommand
}

// NewSysexCommand factory function for SysexCommand
func NewSysexCommand() *SysexCommand {
	return &SysexCommand{}
}

func CastSysexCommand(structType interface{}) *SysexCommand {
	if casted, ok := structType.(SysexCommand); ok {
		return &casted
	}
	if casted, ok := structType.(*SysexCommand); ok {
		return casted
	}
	if casted, ok := structType.(ISysexCommandChild); ok {
		return casted.GetParent()
	}
	return nil
}

func (m *SysexCommand) GetTypeName() string {
	return "SysexCommand"
}

func (m *SysexCommand) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *SysexCommand) GetLengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.GetLengthInBits()
}

func (m *SysexCommand) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (commandType)
	lengthInBits += 8

	return lengthInBits
}

func (m *SysexCommand) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func SysexCommandParse(readBuffer utils.ReadBuffer, response bool) (*SysexCommand, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SysexCommand"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Discriminator Field (commandType) (Used as input to a switch field)
	commandType, _commandTypeErr := readBuffer.ReadUint8("commandType", 8)
	if _commandTypeErr != nil {
		return nil, errors.Wrap(_commandTypeErr, "Error parsing 'commandType' field")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type SysexCommandChild interface {
		InitializeParent(*SysexCommand)
		GetParent() *SysexCommand
	}
	var _child SysexCommandChild
	var typeSwitchError error
	switch {
	case commandType == 0x00: // SysexCommandExtendedId
		_child, typeSwitchError = SysexCommandExtendedIdParse(readBuffer, response)
	case commandType == 0x69 && response == bool(false): // SysexCommandAnalogMappingQueryRequest
		_child, typeSwitchError = SysexCommandAnalogMappingQueryRequestParse(readBuffer, response)
	case commandType == 0x69 && response == bool(true): // SysexCommandAnalogMappingQueryResponse
		_child, typeSwitchError = SysexCommandAnalogMappingQueryResponseParse(readBuffer, response)
	case commandType == 0x6A: // SysexCommandAnalogMappingResponse
		_child, typeSwitchError = SysexCommandAnalogMappingResponseParse(readBuffer, response)
	case commandType == 0x6B: // SysexCommandCapabilityQuery
		_child, typeSwitchError = SysexCommandCapabilityQueryParse(readBuffer, response)
	case commandType == 0x6C: // SysexCommandCapabilityResponse
		_child, typeSwitchError = SysexCommandCapabilityResponseParse(readBuffer, response)
	case commandType == 0x6D: // SysexCommandPinStateQuery
		_child, typeSwitchError = SysexCommandPinStateQueryParse(readBuffer, response)
	case commandType == 0x6E: // SysexCommandPinStateResponse
		_child, typeSwitchError = SysexCommandPinStateResponseParse(readBuffer, response)
	case commandType == 0x6F: // SysexCommandExtendedAnalog
		_child, typeSwitchError = SysexCommandExtendedAnalogParse(readBuffer, response)
	case commandType == 0x71: // SysexCommandStringData
		_child, typeSwitchError = SysexCommandStringDataParse(readBuffer, response)
	case commandType == 0x79 && response == bool(false): // SysexCommandReportFirmwareRequest
		_child, typeSwitchError = SysexCommandReportFirmwareRequestParse(readBuffer, response)
	case commandType == 0x79 && response == bool(true): // SysexCommandReportFirmwareResponse
		_child, typeSwitchError = SysexCommandReportFirmwareResponseParse(readBuffer, response)
	case commandType == 0x7A: // SysexCommandSamplingInterval
		_child, typeSwitchError = SysexCommandSamplingIntervalParse(readBuffer, response)
	case commandType == 0x7E: // SysexCommandSysexNonRealtime
		_child, typeSwitchError = SysexCommandSysexNonRealtimeParse(readBuffer, response)
	case commandType == 0x7F: // SysexCommandSysexRealtime
		_child, typeSwitchError = SysexCommandSysexRealtimeParse(readBuffer, response)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("SysexCommand"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_child.InitializeParent(_child.GetParent())
	return _child.GetParent(), nil
}

func (m *SysexCommand) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *SysexCommand) SerializeParent(writeBuffer utils.WriteBuffer, child ISysexCommand, serializeChildFunction func() error) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("SysexCommand"); pushErr != nil {
		return pushErr
	}

	// Discriminator Field (commandType) (Used as input to a switch field)
	commandType := uint8(child.GetCommandType())
	_commandTypeErr := writeBuffer.WriteUint8("commandType", 8, (commandType))

	if _commandTypeErr != nil {
		return errors.Wrap(_commandTypeErr, "Error serializing 'commandType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("SysexCommand"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *SysexCommand) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}