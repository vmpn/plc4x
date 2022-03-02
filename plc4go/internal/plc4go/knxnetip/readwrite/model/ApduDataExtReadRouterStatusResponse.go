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
type ApduDataExtReadRouterStatusResponse struct {
	*ApduDataExt

	// Arguments.
	Length uint8
}

// The corresponding interface
type IApduDataExtReadRouterStatusResponse interface {
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
func (m *ApduDataExtReadRouterStatusResponse) ExtApciType() uint8 {
	return 0x0E
}

func (m *ApduDataExtReadRouterStatusResponse) GetExtApciType() uint8 {
	return 0x0E
}

func (m *ApduDataExtReadRouterStatusResponse) InitializeParent(parent *ApduDataExt) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewApduDataExtReadRouterStatusResponse factory function for ApduDataExtReadRouterStatusResponse
func NewApduDataExtReadRouterStatusResponse(length uint8) *ApduDataExt {
	child := &ApduDataExtReadRouterStatusResponse{
		ApduDataExt: NewApduDataExt(length),
	}
	child.Child = child
	return child.ApduDataExt
}

func CastApduDataExtReadRouterStatusResponse(structType interface{}) *ApduDataExtReadRouterStatusResponse {
	castFunc := func(typ interface{}) *ApduDataExtReadRouterStatusResponse {
		if casted, ok := typ.(ApduDataExtReadRouterStatusResponse); ok {
			return &casted
		}
		if casted, ok := typ.(*ApduDataExtReadRouterStatusResponse); ok {
			return casted
		}
		if casted, ok := typ.(ApduDataExt); ok {
			return CastApduDataExtReadRouterStatusResponse(casted.Child)
		}
		if casted, ok := typ.(*ApduDataExt); ok {
			return CastApduDataExtReadRouterStatusResponse(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ApduDataExtReadRouterStatusResponse) GetTypeName() string {
	return "ApduDataExtReadRouterStatusResponse"
}

func (m *ApduDataExtReadRouterStatusResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ApduDataExtReadRouterStatusResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *ApduDataExtReadRouterStatusResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApduDataExtReadRouterStatusResponseParse(readBuffer utils.ReadBuffer, length uint8) (*ApduDataExt, error) {
	if pullErr := readBuffer.PullContext("ApduDataExtReadRouterStatusResponse"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataExtReadRouterStatusResponse"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ApduDataExtReadRouterStatusResponse{
		ApduDataExt: &ApduDataExt{},
	}
	_child.ApduDataExt.Child = _child
	return _child.ApduDataExt, nil
}

func (m *ApduDataExtReadRouterStatusResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtReadRouterStatusResponse"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("ApduDataExtReadRouterStatusResponse"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ApduDataExtReadRouterStatusResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
