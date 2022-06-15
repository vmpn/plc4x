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

// BACnetConstructedDataAccompaniment is the data-structure of this message
type BACnetConstructedDataAccompaniment struct {
	*BACnetConstructedData
	Accompaniment *BACnetDeviceObjectReference

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataAccompaniment is the corresponding interface of BACnetConstructedDataAccompaniment
type IBACnetConstructedDataAccompaniment interface {
	IBACnetConstructedData
	// GetAccompaniment returns Accompaniment (property field)
	GetAccompaniment() *BACnetDeviceObjectReference
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() *BACnetDeviceObjectReference
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

func (m *BACnetConstructedDataAccompaniment) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataAccompaniment) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ACCOMPANIMENT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataAccompaniment) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataAccompaniment) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataAccompaniment) GetAccompaniment() *BACnetDeviceObjectReference {
	return m.Accompaniment
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *BACnetConstructedDataAccompaniment) GetActualValue() *BACnetDeviceObjectReference {
	return CastBACnetDeviceObjectReference(m.GetAccompaniment())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAccompaniment factory function for BACnetConstructedDataAccompaniment
func NewBACnetConstructedDataAccompaniment(accompaniment *BACnetDeviceObjectReference, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataAccompaniment {
	_result := &BACnetConstructedDataAccompaniment{
		Accompaniment:         accompaniment,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataAccompaniment(structType interface{}) *BACnetConstructedDataAccompaniment {
	if casted, ok := structType.(BACnetConstructedDataAccompaniment); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAccompaniment); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataAccompaniment(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataAccompaniment(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataAccompaniment) GetTypeName() string {
	return "BACnetConstructedDataAccompaniment"
}

func (m *BACnetConstructedDataAccompaniment) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataAccompaniment) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (accompaniment)
	lengthInBits += m.Accompaniment.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetConstructedDataAccompaniment) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataAccompanimentParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataAccompaniment, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAccompaniment"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAccompaniment")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (accompaniment)
	if pullErr := readBuffer.PullContext("accompaniment"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for accompaniment")
	}
	_accompaniment, _accompanimentErr := BACnetDeviceObjectReferenceParse(readBuffer)
	if _accompanimentErr != nil {
		return nil, errors.Wrap(_accompanimentErr, "Error parsing 'accompaniment' field")
	}
	accompaniment := CastBACnetDeviceObjectReference(_accompaniment)
	if closeErr := readBuffer.CloseContext("accompaniment"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for accompaniment")
	}

	// Virtual field
	_actualValue := accompaniment
	actualValue := CastBACnetDeviceObjectReference(_actualValue)
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAccompaniment"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAccompaniment")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataAccompaniment{
		Accompaniment:         CastBACnetDeviceObjectReference(accompaniment),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataAccompaniment) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAccompaniment"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAccompaniment")
		}

		// Simple Field (accompaniment)
		if pushErr := writeBuffer.PushContext("accompaniment"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for accompaniment")
		}
		_accompanimentErr := writeBuffer.WriteSerializable(m.Accompaniment)
		if popErr := writeBuffer.PopContext("accompaniment"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for accompaniment")
		}
		if _accompanimentErr != nil {
			return errors.Wrap(_accompanimentErr, "Error serializing 'accompaniment' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAccompaniment"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAccompaniment")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataAccompaniment) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}