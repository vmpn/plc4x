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
type BACnetConfirmedServiceRequestRemovedAuthenticate struct {
	*BACnetConfirmedServiceRequest

	// Arguments.
	Len uint16
}

// The corresponding interface
type IBACnetConfirmedServiceRequestRemovedAuthenticate interface {
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
func (m *BACnetConfirmedServiceRequestRemovedAuthenticate) ServiceChoice() uint8 {
	return 0x18
}

func (m *BACnetConfirmedServiceRequestRemovedAuthenticate) GetServiceChoice() uint8 {
	return 0x18
}

func (m *BACnetConfirmedServiceRequestRemovedAuthenticate) InitializeParent(parent *BACnetConfirmedServiceRequest) {
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestRemovedAuthenticate factory function for BACnetConfirmedServiceRequestRemovedAuthenticate
func NewBACnetConfirmedServiceRequestRemovedAuthenticate(len uint16) *BACnetConfirmedServiceRequest {
	child := &BACnetConfirmedServiceRequestRemovedAuthenticate{
		BACnetConfirmedServiceRequest: NewBACnetConfirmedServiceRequest(len),
	}
	child.Child = child
	return child.BACnetConfirmedServiceRequest
}

func CastBACnetConfirmedServiceRequestRemovedAuthenticate(structType interface{}) *BACnetConfirmedServiceRequestRemovedAuthenticate {
	castFunc := func(typ interface{}) *BACnetConfirmedServiceRequestRemovedAuthenticate {
		if casted, ok := typ.(BACnetConfirmedServiceRequestRemovedAuthenticate); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetConfirmedServiceRequestRemovedAuthenticate); ok {
			return casted
		}
		if casted, ok := typ.(BACnetConfirmedServiceRequest); ok {
			return CastBACnetConfirmedServiceRequestRemovedAuthenticate(casted.Child)
		}
		if casted, ok := typ.(*BACnetConfirmedServiceRequest); ok {
			return CastBACnetConfirmedServiceRequestRemovedAuthenticate(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetConfirmedServiceRequestRemovedAuthenticate) GetTypeName() string {
	return "BACnetConfirmedServiceRequestRemovedAuthenticate"
}

func (m *BACnetConfirmedServiceRequestRemovedAuthenticate) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConfirmedServiceRequestRemovedAuthenticate) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *BACnetConfirmedServiceRequestRemovedAuthenticate) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConfirmedServiceRequestRemovedAuthenticateParse(readBuffer utils.ReadBuffer, len uint16) (*BACnetConfirmedServiceRequest, error) {
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestRemovedAuthenticate"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestRemovedAuthenticate"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConfirmedServiceRequestRemovedAuthenticate{
		BACnetConfirmedServiceRequest: &BACnetConfirmedServiceRequest{},
	}
	_child.BACnetConfirmedServiceRequest.Child = _child
	return _child.BACnetConfirmedServiceRequest, nil
}

func (m *BACnetConfirmedServiceRequestRemovedAuthenticate) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestRemovedAuthenticate"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestRemovedAuthenticate"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConfirmedServiceRequestRemovedAuthenticate) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
