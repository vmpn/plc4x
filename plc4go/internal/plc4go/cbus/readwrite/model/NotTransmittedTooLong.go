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
type NotTransmittedTooLong struct {
	*Confirmation
}

// The corresponding interface
type INotTransmittedTooLong interface {
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
func (m *NotTransmittedTooLong) ConfirmationType() byte {
	return 0x27
}

func (m *NotTransmittedTooLong) GetConfirmationType() byte {
	return 0x27
}

func (m *NotTransmittedTooLong) InitializeParent(parent *Confirmation, alpha *Alpha) {
	m.Confirmation.Alpha = alpha
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewNotTransmittedTooLong factory function for NotTransmittedTooLong
func NewNotTransmittedTooLong(alpha *Alpha) *Confirmation {
	child := &NotTransmittedTooLong{
		Confirmation: NewConfirmation(alpha),
	}
	child.Child = child
	return child.Confirmation
}

func CastNotTransmittedTooLong(structType interface{}) *NotTransmittedTooLong {
	castFunc := func(typ interface{}) *NotTransmittedTooLong {
		if casted, ok := typ.(NotTransmittedTooLong); ok {
			return &casted
		}
		if casted, ok := typ.(*NotTransmittedTooLong); ok {
			return casted
		}
		if casted, ok := typ.(Confirmation); ok {
			return CastNotTransmittedTooLong(casted.Child)
		}
		if casted, ok := typ.(*Confirmation); ok {
			return CastNotTransmittedTooLong(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *NotTransmittedTooLong) GetTypeName() string {
	return "NotTransmittedTooLong"
}

func (m *NotTransmittedTooLong) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *NotTransmittedTooLong) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *NotTransmittedTooLong) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func NotTransmittedTooLongParse(readBuffer utils.ReadBuffer) (*Confirmation, error) {
	if pullErr := readBuffer.PullContext("NotTransmittedTooLong"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("NotTransmittedTooLong"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &NotTransmittedTooLong{
		Confirmation: &Confirmation{},
	}
	_child.Confirmation.Child = _child
	return _child.Confirmation, nil
}

func (m *NotTransmittedTooLong) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NotTransmittedTooLong"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("NotTransmittedTooLong"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *NotTransmittedTooLong) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
