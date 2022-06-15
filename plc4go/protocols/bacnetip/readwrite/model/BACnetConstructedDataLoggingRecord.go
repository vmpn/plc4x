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

// BACnetConstructedDataLoggingRecord is the data-structure of this message
type BACnetConstructedDataLoggingRecord struct {
	*BACnetConstructedData
	LoggingRecord *BACnetAccumulatorRecord

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataLoggingRecord is the corresponding interface of BACnetConstructedDataLoggingRecord
type IBACnetConstructedDataLoggingRecord interface {
	IBACnetConstructedData
	// GetLoggingRecord returns LoggingRecord (property field)
	GetLoggingRecord() *BACnetAccumulatorRecord
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() *BACnetAccumulatorRecord
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

func (m *BACnetConstructedDataLoggingRecord) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataLoggingRecord) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LOGGING_RECORD
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataLoggingRecord) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataLoggingRecord) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataLoggingRecord) GetLoggingRecord() *BACnetAccumulatorRecord {
	return m.LoggingRecord
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *BACnetConstructedDataLoggingRecord) GetActualValue() *BACnetAccumulatorRecord {
	return CastBACnetAccumulatorRecord(m.GetLoggingRecord())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataLoggingRecord factory function for BACnetConstructedDataLoggingRecord
func NewBACnetConstructedDataLoggingRecord(loggingRecord *BACnetAccumulatorRecord, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataLoggingRecord {
	_result := &BACnetConstructedDataLoggingRecord{
		LoggingRecord:         loggingRecord,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataLoggingRecord(structType interface{}) *BACnetConstructedDataLoggingRecord {
	if casted, ok := structType.(BACnetConstructedDataLoggingRecord); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLoggingRecord); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataLoggingRecord(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataLoggingRecord(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataLoggingRecord) GetTypeName() string {
	return "BACnetConstructedDataLoggingRecord"
}

func (m *BACnetConstructedDataLoggingRecord) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataLoggingRecord) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (loggingRecord)
	lengthInBits += m.LoggingRecord.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetConstructedDataLoggingRecord) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataLoggingRecordParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataLoggingRecord, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLoggingRecord"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLoggingRecord")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (loggingRecord)
	if pullErr := readBuffer.PullContext("loggingRecord"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for loggingRecord")
	}
	_loggingRecord, _loggingRecordErr := BACnetAccumulatorRecordParse(readBuffer)
	if _loggingRecordErr != nil {
		return nil, errors.Wrap(_loggingRecordErr, "Error parsing 'loggingRecord' field")
	}
	loggingRecord := CastBACnetAccumulatorRecord(_loggingRecord)
	if closeErr := readBuffer.CloseContext("loggingRecord"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for loggingRecord")
	}

	// Virtual field
	_actualValue := loggingRecord
	actualValue := CastBACnetAccumulatorRecord(_actualValue)
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLoggingRecord"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLoggingRecord")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataLoggingRecord{
		LoggingRecord:         CastBACnetAccumulatorRecord(loggingRecord),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataLoggingRecord) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLoggingRecord"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLoggingRecord")
		}

		// Simple Field (loggingRecord)
		if pushErr := writeBuffer.PushContext("loggingRecord"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for loggingRecord")
		}
		_loggingRecordErr := writeBuffer.WriteSerializable(m.LoggingRecord)
		if popErr := writeBuffer.PopContext("loggingRecord"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for loggingRecord")
		}
		if _loggingRecordErr != nil {
			return errors.Wrap(_loggingRecordErr, "Error serializing 'loggingRecord' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLoggingRecord"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLoggingRecord")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataLoggingRecord) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}