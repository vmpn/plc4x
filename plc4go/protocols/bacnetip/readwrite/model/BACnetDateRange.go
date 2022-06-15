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

// BACnetDateRange is the data-structure of this message
type BACnetDateRange struct {
	StartDate *BACnetApplicationTagDate
	EndDate   *BACnetApplicationTagDate
}

// IBACnetDateRange is the corresponding interface of BACnetDateRange
type IBACnetDateRange interface {
	// GetStartDate returns StartDate (property field)
	GetStartDate() *BACnetApplicationTagDate
	// GetEndDate returns EndDate (property field)
	GetEndDate() *BACnetApplicationTagDate
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetDateRange) GetStartDate() *BACnetApplicationTagDate {
	return m.StartDate
}

func (m *BACnetDateRange) GetEndDate() *BACnetApplicationTagDate {
	return m.EndDate
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetDateRange factory function for BACnetDateRange
func NewBACnetDateRange(startDate *BACnetApplicationTagDate, endDate *BACnetApplicationTagDate) *BACnetDateRange {
	return &BACnetDateRange{StartDate: startDate, EndDate: endDate}
}

func CastBACnetDateRange(structType interface{}) *BACnetDateRange {
	if casted, ok := structType.(BACnetDateRange); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetDateRange); ok {
		return casted
	}
	return nil
}

func (m *BACnetDateRange) GetTypeName() string {
	return "BACnetDateRange"
}

func (m *BACnetDateRange) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetDateRange) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (startDate)
	lengthInBits += m.StartDate.GetLengthInBits()

	// Simple field (endDate)
	lengthInBits += m.EndDate.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetDateRange) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetDateRangeParse(readBuffer utils.ReadBuffer) (*BACnetDateRange, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetDateRange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetDateRange")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (startDate)
	if pullErr := readBuffer.PullContext("startDate"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for startDate")
	}
	_startDate, _startDateErr := BACnetApplicationTagParse(readBuffer)
	if _startDateErr != nil {
		return nil, errors.Wrap(_startDateErr, "Error parsing 'startDate' field")
	}
	startDate := CastBACnetApplicationTagDate(_startDate)
	if closeErr := readBuffer.CloseContext("startDate"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for startDate")
	}

	// Simple Field (endDate)
	if pullErr := readBuffer.PullContext("endDate"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for endDate")
	}
	_endDate, _endDateErr := BACnetApplicationTagParse(readBuffer)
	if _endDateErr != nil {
		return nil, errors.Wrap(_endDateErr, "Error parsing 'endDate' field")
	}
	endDate := CastBACnetApplicationTagDate(_endDate)
	if closeErr := readBuffer.CloseContext("endDate"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for endDate")
	}

	if closeErr := readBuffer.CloseContext("BACnetDateRange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetDateRange")
	}

	// Create the instance
	return NewBACnetDateRange(startDate, endDate), nil
}

func (m *BACnetDateRange) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetDateRange"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetDateRange")
	}

	// Simple Field (startDate)
	if pushErr := writeBuffer.PushContext("startDate"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for startDate")
	}
	_startDateErr := writeBuffer.WriteSerializable(m.StartDate)
	if popErr := writeBuffer.PopContext("startDate"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for startDate")
	}
	if _startDateErr != nil {
		return errors.Wrap(_startDateErr, "Error serializing 'startDate' field")
	}

	// Simple Field (endDate)
	if pushErr := writeBuffer.PushContext("endDate"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for endDate")
	}
	_endDateErr := writeBuffer.WriteSerializable(m.EndDate)
	if popErr := writeBuffer.PopContext("endDate"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for endDate")
	}
	if _endDateErr != nil {
		return errors.Wrap(_endDateErr, "Error serializing 'endDate' field")
	}

	if popErr := writeBuffer.PopContext("BACnetDateRange"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetDateRange")
	}
	return nil
}

func (m *BACnetDateRange) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}