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

// BACnetConstructedDataZoneTo is the data-structure of this message
type BACnetConstructedDataZoneTo struct {
	*BACnetConstructedData
	ZoneTo *BACnetDeviceObjectReference

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataZoneTo is the corresponding interface of BACnetConstructedDataZoneTo
type IBACnetConstructedDataZoneTo interface {
	IBACnetConstructedData
	// GetZoneTo returns ZoneTo (property field)
	GetZoneTo() *BACnetDeviceObjectReference
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

func (m *BACnetConstructedDataZoneTo) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataZoneTo) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ZONE_TO
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataZoneTo) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataZoneTo) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataZoneTo) GetZoneTo() *BACnetDeviceObjectReference {
	return m.ZoneTo
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *BACnetConstructedDataZoneTo) GetActualValue() *BACnetDeviceObjectReference {
	return CastBACnetDeviceObjectReference(m.GetZoneTo())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataZoneTo factory function for BACnetConstructedDataZoneTo
func NewBACnetConstructedDataZoneTo(zoneTo *BACnetDeviceObjectReference, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataZoneTo {
	_result := &BACnetConstructedDataZoneTo{
		ZoneTo:                zoneTo,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataZoneTo(structType interface{}) *BACnetConstructedDataZoneTo {
	if casted, ok := structType.(BACnetConstructedDataZoneTo); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataZoneTo); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataZoneTo(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataZoneTo(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataZoneTo) GetTypeName() string {
	return "BACnetConstructedDataZoneTo"
}

func (m *BACnetConstructedDataZoneTo) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataZoneTo) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (zoneTo)
	lengthInBits += m.ZoneTo.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetConstructedDataZoneTo) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataZoneToParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataZoneTo, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataZoneTo"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataZoneTo")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (zoneTo)
	if pullErr := readBuffer.PullContext("zoneTo"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for zoneTo")
	}
	_zoneTo, _zoneToErr := BACnetDeviceObjectReferenceParse(readBuffer)
	if _zoneToErr != nil {
		return nil, errors.Wrap(_zoneToErr, "Error parsing 'zoneTo' field")
	}
	zoneTo := CastBACnetDeviceObjectReference(_zoneTo)
	if closeErr := readBuffer.CloseContext("zoneTo"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for zoneTo")
	}

	// Virtual field
	_actualValue := zoneTo
	actualValue := CastBACnetDeviceObjectReference(_actualValue)
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataZoneTo"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataZoneTo")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataZoneTo{
		ZoneTo:                CastBACnetDeviceObjectReference(zoneTo),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataZoneTo) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataZoneTo"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataZoneTo")
		}

		// Simple Field (zoneTo)
		if pushErr := writeBuffer.PushContext("zoneTo"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for zoneTo")
		}
		_zoneToErr := writeBuffer.WriteSerializable(m.ZoneTo)
		if popErr := writeBuffer.PopContext("zoneTo"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for zoneTo")
		}
		if _zoneToErr != nil {
			return errors.Wrap(_zoneToErr, "Error serializing 'zoneTo' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataZoneTo"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataZoneTo")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataZoneTo) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}