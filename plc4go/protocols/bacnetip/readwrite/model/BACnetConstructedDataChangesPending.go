/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
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

// BACnetConstructedDataChangesPending is the data-structure of this message
type BACnetConstructedDataChangesPending struct {
	*BACnetConstructedData
	ChangesPending *BACnetApplicationTagBoolean

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataChangesPending is the corresponding interface of BACnetConstructedDataChangesPending
type IBACnetConstructedDataChangesPending interface {
	IBACnetConstructedData
	// GetChangesPending returns ChangesPending (property field)
	GetChangesPending() *BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() *BACnetApplicationTagBoolean
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

func (m *BACnetConstructedDataChangesPending) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataChangesPending) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_CHANGES_PENDING
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataChangesPending) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataChangesPending) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataChangesPending) GetChangesPending() *BACnetApplicationTagBoolean {
	return m.ChangesPending
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *BACnetConstructedDataChangesPending) GetActualValue() *BACnetApplicationTagBoolean {
	return CastBACnetApplicationTagBoolean(m.GetChangesPending())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataChangesPending factory function for BACnetConstructedDataChangesPending
func NewBACnetConstructedDataChangesPending(changesPending *BACnetApplicationTagBoolean, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataChangesPending {
	_result := &BACnetConstructedDataChangesPending{
		ChangesPending:        changesPending,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataChangesPending(structType interface{}) *BACnetConstructedDataChangesPending {
	if casted, ok := structType.(BACnetConstructedDataChangesPending); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataChangesPending); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataChangesPending(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataChangesPending(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataChangesPending) GetTypeName() string {
	return "BACnetConstructedDataChangesPending"
}

func (m *BACnetConstructedDataChangesPending) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataChangesPending) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (changesPending)
	lengthInBits += m.ChangesPending.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetConstructedDataChangesPending) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataChangesPendingParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataChangesPending, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataChangesPending"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataChangesPending")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (changesPending)
	if pullErr := readBuffer.PullContext("changesPending"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for changesPending")
	}
	_changesPending, _changesPendingErr := BACnetApplicationTagParse(readBuffer)
	if _changesPendingErr != nil {
		return nil, errors.Wrap(_changesPendingErr, "Error parsing 'changesPending' field")
	}
	changesPending := CastBACnetApplicationTagBoolean(_changesPending)
	if closeErr := readBuffer.CloseContext("changesPending"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for changesPending")
	}

	// Virtual field
	_actualValue := changesPending
	actualValue := CastBACnetApplicationTagBoolean(_actualValue)
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataChangesPending"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataChangesPending")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataChangesPending{
		ChangesPending:        CastBACnetApplicationTagBoolean(changesPending),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataChangesPending) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataChangesPending"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataChangesPending")
		}

		// Simple Field (changesPending)
		if pushErr := writeBuffer.PushContext("changesPending"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for changesPending")
		}
		_changesPendingErr := writeBuffer.WriteSerializable(m.ChangesPending)
		if popErr := writeBuffer.PopContext("changesPending"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for changesPending")
		}
		if _changesPendingErr != nil {
			return errors.Wrap(_changesPendingErr, "Error serializing 'changesPending' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataChangesPending"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataChangesPending")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataChangesPending) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}