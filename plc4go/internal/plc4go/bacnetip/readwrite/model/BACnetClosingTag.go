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
type BACnetClosingTag struct {
	*BACnetContextTag

	// Arguments.
	TagNumberArgument uint8
	ActualLength      uint32
}

// The corresponding interface
type IBACnetClosingTag interface {
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
func (m *BACnetClosingTag) DataType() BACnetDataType {
	return BACnetDataType_CLOSING_TAG
}

func (m *BACnetClosingTag) GetDataType() BACnetDataType {
	return BACnetDataType_CLOSING_TAG
}

func (m *BACnetClosingTag) InitializeParent(parent *BACnetContextTag, header *BACnetTagHeader) {
	m.BACnetContextTag.Header = header
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewBACnetClosingTag factory function for BACnetClosingTag
func NewBACnetClosingTag(header *BACnetTagHeader, tagNumberArgument uint8, actualLength uint32) *BACnetContextTag {
	child := &BACnetClosingTag{
		BACnetContextTag: NewBACnetContextTag(header, tagNumberArgument),
	}
	child.Child = child
	return child.BACnetContextTag
}

func CastBACnetClosingTag(structType interface{}) *BACnetClosingTag {
	castFunc := func(typ interface{}) *BACnetClosingTag {
		if casted, ok := typ.(BACnetClosingTag); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetClosingTag); ok {
			return casted
		}
		if casted, ok := typ.(BACnetContextTag); ok {
			return CastBACnetClosingTag(casted.Child)
		}
		if casted, ok := typ.(*BACnetContextTag); ok {
			return CastBACnetClosingTag(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetClosingTag) GetTypeName() string {
	return "BACnetClosingTag"
}

func (m *BACnetClosingTag) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetClosingTag) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *BACnetClosingTag) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetClosingTagParse(readBuffer utils.ReadBuffer, tagNumberArgument uint8, dataType BACnetDataType, actualLength uint32) (*BACnetContextTag, error) {
	if pullErr := readBuffer.PullContext("BACnetClosingTag"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Validation
	if !(bool((actualLength) == (7))) {
		return nil, utils.ParseAssertError{"closing tag should habe a value of 7"}
	}

	if closeErr := readBuffer.CloseContext("BACnetClosingTag"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetClosingTag{
		BACnetContextTag: &BACnetContextTag{},
	}
	_child.BACnetContextTag.Child = _child
	return _child.BACnetContextTag, nil
}

func (m *BACnetClosingTag) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetClosingTag"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("BACnetClosingTag"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetClosingTag) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
