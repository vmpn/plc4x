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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type FirmataCommandSystemReset struct {
	*FirmataCommand

	// Arguments.
	Response bool
}

// The corresponding interface
type IFirmataCommandSystemReset interface {
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *FirmataCommandSystemReset) CommandCode() uint8 {
	return 0xF
}

func (m *FirmataCommandSystemReset) GetCommandCode() uint8 {
	return 0xF
}

func (m *FirmataCommandSystemReset) InitializeParent(parent *FirmataCommand) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewFirmataCommandSystemReset factory function for FirmataCommandSystemReset
func NewFirmataCommandSystemReset(response bool) *FirmataCommand {
	child := &FirmataCommandSystemReset{
		FirmataCommand: NewFirmataCommand(response),
	}
	child.Child = child
	return child.FirmataCommand
}

func CastFirmataCommandSystemReset(structType interface{}) *FirmataCommandSystemReset {
	castFunc := func(typ interface{}) *FirmataCommandSystemReset {
		if casted, ok := typ.(FirmataCommandSystemReset); ok {
			return &casted
		}
		if casted, ok := typ.(*FirmataCommandSystemReset); ok {
			return casted
		}
		if casted, ok := typ.(FirmataCommand); ok {
			return CastFirmataCommandSystemReset(casted.Child)
		}
		if casted, ok := typ.(*FirmataCommand); ok {
			return CastFirmataCommandSystemReset(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *FirmataCommandSystemReset) GetTypeName() string {
	return "FirmataCommandSystemReset"
}

func (m *FirmataCommandSystemReset) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *FirmataCommandSystemReset) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *FirmataCommandSystemReset) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func FirmataCommandSystemResetParse(readBuffer utils.ReadBuffer, response bool) (*FirmataCommand, error) {
	if pullErr := readBuffer.PullContext("FirmataCommandSystemReset"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("FirmataCommandSystemReset"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &FirmataCommandSystemReset{
		FirmataCommand: &FirmataCommand{},
	}
	_child.FirmataCommand.Child = _child
	return _child.FirmataCommand, nil
}

func (m *FirmataCommandSystemReset) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("FirmataCommandSystemReset"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("FirmataCommandSystemReset"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *FirmataCommandSystemReset) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
