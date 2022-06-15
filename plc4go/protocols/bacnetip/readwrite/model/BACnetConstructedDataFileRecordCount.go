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

// BACnetConstructedDataFileRecordCount is the data-structure of this message
type BACnetConstructedDataFileRecordCount struct {
	*BACnetConstructedData
	RecordCount *BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataFileRecordCount is the corresponding interface of BACnetConstructedDataFileRecordCount
type IBACnetConstructedDataFileRecordCount interface {
	IBACnetConstructedData
	// GetRecordCount returns RecordCount (property field)
	GetRecordCount() *BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() *BACnetApplicationTagUnsignedInteger
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

func (m *BACnetConstructedDataFileRecordCount) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_FILE
}

func (m *BACnetConstructedDataFileRecordCount) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_RECORD_COUNT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataFileRecordCount) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataFileRecordCount) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataFileRecordCount) GetRecordCount() *BACnetApplicationTagUnsignedInteger {
	return m.RecordCount
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *BACnetConstructedDataFileRecordCount) GetActualValue() *BACnetApplicationTagUnsignedInteger {
	return CastBACnetApplicationTagUnsignedInteger(m.GetRecordCount())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataFileRecordCount factory function for BACnetConstructedDataFileRecordCount
func NewBACnetConstructedDataFileRecordCount(recordCount *BACnetApplicationTagUnsignedInteger, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataFileRecordCount {
	_result := &BACnetConstructedDataFileRecordCount{
		RecordCount:           recordCount,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataFileRecordCount(structType interface{}) *BACnetConstructedDataFileRecordCount {
	if casted, ok := structType.(BACnetConstructedDataFileRecordCount); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataFileRecordCount); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataFileRecordCount(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataFileRecordCount(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataFileRecordCount) GetTypeName() string {
	return "BACnetConstructedDataFileRecordCount"
}

func (m *BACnetConstructedDataFileRecordCount) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataFileRecordCount) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (recordCount)
	lengthInBits += m.RecordCount.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetConstructedDataFileRecordCount) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataFileRecordCountParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataFileRecordCount, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataFileRecordCount"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataFileRecordCount")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (recordCount)
	if pullErr := readBuffer.PullContext("recordCount"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for recordCount")
	}
	_recordCount, _recordCountErr := BACnetApplicationTagParse(readBuffer)
	if _recordCountErr != nil {
		return nil, errors.Wrap(_recordCountErr, "Error parsing 'recordCount' field")
	}
	recordCount := CastBACnetApplicationTagUnsignedInteger(_recordCount)
	if closeErr := readBuffer.CloseContext("recordCount"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for recordCount")
	}

	// Virtual field
	_actualValue := recordCount
	actualValue := CastBACnetApplicationTagUnsignedInteger(_actualValue)
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataFileRecordCount"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataFileRecordCount")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataFileRecordCount{
		RecordCount:           CastBACnetApplicationTagUnsignedInteger(recordCount),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataFileRecordCount) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataFileRecordCount"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataFileRecordCount")
		}

		// Simple Field (recordCount)
		if pushErr := writeBuffer.PushContext("recordCount"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for recordCount")
		}
		_recordCountErr := writeBuffer.WriteSerializable(m.RecordCount)
		if popErr := writeBuffer.PopContext("recordCount"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for recordCount")
		}
		if _recordCountErr != nil {
			return errors.Wrap(_recordCountErr, "Error serializing 'recordCount' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataFileRecordCount"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataFileRecordCount")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataFileRecordCount) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}