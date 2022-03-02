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
type BACnetConfirmedServiceACKConfirmedPrivateTransfer struct {
	*BACnetConfirmedServiceACK
}

// The corresponding interface
type IBACnetConfirmedServiceACKConfirmedPrivateTransfer interface {
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
func (m *BACnetConfirmedServiceACKConfirmedPrivateTransfer) ServiceChoice() uint8 {
	return 0x12
}

func (m *BACnetConfirmedServiceACKConfirmedPrivateTransfer) GetServiceChoice() uint8 {
	return 0x12
}

func (m *BACnetConfirmedServiceACKConfirmedPrivateTransfer) InitializeParent(parent *BACnetConfirmedServiceACK) {
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceACKConfirmedPrivateTransfer factory function for BACnetConfirmedServiceACKConfirmedPrivateTransfer
func NewBACnetConfirmedServiceACKConfirmedPrivateTransfer() *BACnetConfirmedServiceACK {
	child := &BACnetConfirmedServiceACKConfirmedPrivateTransfer{
		BACnetConfirmedServiceACK: NewBACnetConfirmedServiceACK(),
	}
	child.Child = child
	return child.BACnetConfirmedServiceACK
}

func CastBACnetConfirmedServiceACKConfirmedPrivateTransfer(structType interface{}) *BACnetConfirmedServiceACKConfirmedPrivateTransfer {
	castFunc := func(typ interface{}) *BACnetConfirmedServiceACKConfirmedPrivateTransfer {
		if casted, ok := typ.(BACnetConfirmedServiceACKConfirmedPrivateTransfer); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetConfirmedServiceACKConfirmedPrivateTransfer); ok {
			return casted
		}
		if casted, ok := typ.(BACnetConfirmedServiceACK); ok {
			return CastBACnetConfirmedServiceACKConfirmedPrivateTransfer(casted.Child)
		}
		if casted, ok := typ.(*BACnetConfirmedServiceACK); ok {
			return CastBACnetConfirmedServiceACKConfirmedPrivateTransfer(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetConfirmedServiceACKConfirmedPrivateTransfer) GetTypeName() string {
	return "BACnetConfirmedServiceACKConfirmedPrivateTransfer"
}

func (m *BACnetConfirmedServiceACKConfirmedPrivateTransfer) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConfirmedServiceACKConfirmedPrivateTransfer) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *BACnetConfirmedServiceACKConfirmedPrivateTransfer) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConfirmedServiceACKConfirmedPrivateTransferParse(readBuffer utils.ReadBuffer) (*BACnetConfirmedServiceACK, error) {
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceACKConfirmedPrivateTransfer"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceACKConfirmedPrivateTransfer"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConfirmedServiceACKConfirmedPrivateTransfer{
		BACnetConfirmedServiceACK: &BACnetConfirmedServiceACK{},
	}
	_child.BACnetConfirmedServiceACK.Child = _child
	return _child.BACnetConfirmedServiceACK, nil
}

func (m *BACnetConfirmedServiceACKConfirmedPrivateTransfer) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceACKConfirmedPrivateTransfer"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceACKConfirmedPrivateTransfer"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConfirmedServiceACKConfirmedPrivateTransfer) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
