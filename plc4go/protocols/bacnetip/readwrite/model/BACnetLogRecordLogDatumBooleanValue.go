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

// BACnetLogRecordLogDatumBooleanValue is the data-structure of this message
type BACnetLogRecordLogDatumBooleanValue struct {
	*BACnetLogRecordLogDatum
	BooleanValue *BACnetContextTagBoolean

	// Arguments.
	TagNumber uint8
}

// IBACnetLogRecordLogDatumBooleanValue is the corresponding interface of BACnetLogRecordLogDatumBooleanValue
type IBACnetLogRecordLogDatumBooleanValue interface {
	IBACnetLogRecordLogDatum
	// GetBooleanValue returns BooleanValue (property field)
	GetBooleanValue() *BACnetContextTagBoolean
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

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetLogRecordLogDatumBooleanValue) InitializeParent(parent *BACnetLogRecordLogDatum, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetLogRecordLogDatum.OpeningTag = openingTag
	m.BACnetLogRecordLogDatum.PeekedTagHeader = peekedTagHeader
	m.BACnetLogRecordLogDatum.ClosingTag = closingTag
}

func (m *BACnetLogRecordLogDatumBooleanValue) GetParent() *BACnetLogRecordLogDatum {
	return m.BACnetLogRecordLogDatum
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetLogRecordLogDatumBooleanValue) GetBooleanValue() *BACnetContextTagBoolean {
	return m.BooleanValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetLogRecordLogDatumBooleanValue factory function for BACnetLogRecordLogDatumBooleanValue
func NewBACnetLogRecordLogDatumBooleanValue(booleanValue *BACnetContextTagBoolean, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetLogRecordLogDatumBooleanValue {
	_result := &BACnetLogRecordLogDatumBooleanValue{
		BooleanValue:            booleanValue,
		BACnetLogRecordLogDatum: NewBACnetLogRecordLogDatum(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetLogRecordLogDatumBooleanValue(structType interface{}) *BACnetLogRecordLogDatumBooleanValue {
	if casted, ok := structType.(BACnetLogRecordLogDatumBooleanValue); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetLogRecordLogDatumBooleanValue); ok {
		return casted
	}
	if casted, ok := structType.(BACnetLogRecordLogDatum); ok {
		return CastBACnetLogRecordLogDatumBooleanValue(casted.Child)
	}
	if casted, ok := structType.(*BACnetLogRecordLogDatum); ok {
		return CastBACnetLogRecordLogDatumBooleanValue(casted.Child)
	}
	return nil
}

func (m *BACnetLogRecordLogDatumBooleanValue) GetTypeName() string {
	return "BACnetLogRecordLogDatumBooleanValue"
}

func (m *BACnetLogRecordLogDatumBooleanValue) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetLogRecordLogDatumBooleanValue) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (booleanValue)
	lengthInBits += m.BooleanValue.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetLogRecordLogDatumBooleanValue) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetLogRecordLogDatumBooleanValueParse(readBuffer utils.ReadBuffer, tagNumber uint8) (*BACnetLogRecordLogDatumBooleanValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLogRecordLogDatumBooleanValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLogRecordLogDatumBooleanValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (booleanValue)
	if pullErr := readBuffer.PullContext("booleanValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for booleanValue")
	}
	_booleanValue, _booleanValueErr := BACnetContextTagParse(readBuffer, uint8(uint8(1)), BACnetDataType(BACnetDataType_BOOLEAN))
	if _booleanValueErr != nil {
		return nil, errors.Wrap(_booleanValueErr, "Error parsing 'booleanValue' field")
	}
	booleanValue := CastBACnetContextTagBoolean(_booleanValue)
	if closeErr := readBuffer.CloseContext("booleanValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for booleanValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetLogRecordLogDatumBooleanValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLogRecordLogDatumBooleanValue")
	}

	// Create a partially initialized instance
	_child := &BACnetLogRecordLogDatumBooleanValue{
		BooleanValue:            CastBACnetContextTagBoolean(booleanValue),
		BACnetLogRecordLogDatum: &BACnetLogRecordLogDatum{},
	}
	_child.BACnetLogRecordLogDatum.Child = _child
	return _child, nil
}

func (m *BACnetLogRecordLogDatumBooleanValue) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetLogRecordLogDatumBooleanValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetLogRecordLogDatumBooleanValue")
		}

		// Simple Field (booleanValue)
		if pushErr := writeBuffer.PushContext("booleanValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for booleanValue")
		}
		_booleanValueErr := writeBuffer.WriteSerializable(m.BooleanValue)
		if popErr := writeBuffer.PopContext("booleanValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for booleanValue")
		}
		if _booleanValueErr != nil {
			return errors.Wrap(_booleanValueErr, "Error serializing 'booleanValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetLogRecordLogDatumBooleanValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetLogRecordLogDatumBooleanValue")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetLogRecordLogDatumBooleanValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}