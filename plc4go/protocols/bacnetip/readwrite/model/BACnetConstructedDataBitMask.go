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

// BACnetConstructedDataBitMask is the data-structure of this message
type BACnetConstructedDataBitMask struct {
	*BACnetConstructedData
	BitString *BACnetApplicationTagBitString

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataBitMask is the corresponding interface of BACnetConstructedDataBitMask
type IBACnetConstructedDataBitMask interface {
	IBACnetConstructedData
	// GetBitString returns BitString (property field)
	GetBitString() *BACnetApplicationTagBitString
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *BACnetConstructedDataBitMask) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataBitMask) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_BIT_MASK
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataBitMask) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataBitMask) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataBitMask) GetBitString() *BACnetApplicationTagBitString {
	return m.BitString
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataBitMask factory function for BACnetConstructedDataBitMask
func NewBACnetConstructedDataBitMask(bitString *BACnetApplicationTagBitString, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataBitMask {
	_result := &BACnetConstructedDataBitMask{
		BitString:             bitString,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataBitMask(structType interface{}) *BACnetConstructedDataBitMask {
	if casted, ok := structType.(BACnetConstructedDataBitMask); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataBitMask); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataBitMask(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataBitMask(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataBitMask) GetTypeName() string {
	return "BACnetConstructedDataBitMask"
}

func (m *BACnetConstructedDataBitMask) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataBitMask) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (bitString)
	lengthInBits += m.BitString.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataBitMask) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataBitMaskParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataBitMask, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataBitMask"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (bitString)
	if pullErr := readBuffer.PullContext("bitString"); pullErr != nil {
		return nil, pullErr
	}
	_bitString, _bitStringErr := BACnetApplicationTagParse(readBuffer)
	if _bitStringErr != nil {
		return nil, errors.Wrap(_bitStringErr, "Error parsing 'bitString' field")
	}
	bitString := CastBACnetApplicationTagBitString(_bitString)
	if closeErr := readBuffer.CloseContext("bitString"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataBitMask"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataBitMask{
		BitString:             CastBACnetApplicationTagBitString(bitString),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataBitMask) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataBitMask"); pushErr != nil {
			return pushErr
		}

		// Simple Field (bitString)
		if pushErr := writeBuffer.PushContext("bitString"); pushErr != nil {
			return pushErr
		}
		_bitStringErr := m.BitString.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("bitString"); popErr != nil {
			return popErr
		}
		if _bitStringErr != nil {
			return errors.Wrap(_bitStringErr, "Error serializing 'bitString' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataBitMask"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataBitMask) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
