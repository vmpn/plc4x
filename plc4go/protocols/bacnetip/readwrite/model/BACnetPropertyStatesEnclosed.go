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

// BACnetPropertyStatesEnclosed is the data-structure of this message
type BACnetPropertyStatesEnclosed struct {
	OpeningTag    *BACnetOpeningTag
	PropertyState *BACnetPropertyStates
	ClosingTag    *BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

// IBACnetPropertyStatesEnclosed is the corresponding interface of BACnetPropertyStatesEnclosed
type IBACnetPropertyStatesEnclosed interface {
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() *BACnetOpeningTag
	// GetPropertyState returns PropertyState (property field)
	GetPropertyState() *BACnetPropertyStates
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() *BACnetClosingTag
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

func (m *BACnetPropertyStatesEnclosed) GetOpeningTag() *BACnetOpeningTag {
	return m.OpeningTag
}

func (m *BACnetPropertyStatesEnclosed) GetPropertyState() *BACnetPropertyStates {
	return m.PropertyState
}

func (m *BACnetPropertyStatesEnclosed) GetClosingTag() *BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesEnclosed factory function for BACnetPropertyStatesEnclosed
func NewBACnetPropertyStatesEnclosed(openingTag *BACnetOpeningTag, propertyState *BACnetPropertyStates, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetPropertyStatesEnclosed {
	return &BACnetPropertyStatesEnclosed{OpeningTag: openingTag, PropertyState: propertyState, ClosingTag: closingTag, TagNumber: tagNumber}
}

func CastBACnetPropertyStatesEnclosed(structType interface{}) *BACnetPropertyStatesEnclosed {
	if casted, ok := structType.(BACnetPropertyStatesEnclosed); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesEnclosed); ok {
		return casted
	}
	return nil
}

func (m *BACnetPropertyStatesEnclosed) GetTypeName() string {
	return "BACnetPropertyStatesEnclosed"
}

func (m *BACnetPropertyStatesEnclosed) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetPropertyStatesEnclosed) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits()

	// Simple field (propertyState)
	lengthInBits += m.PropertyState.GetLengthInBits()

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetPropertyStatesEnclosed) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPropertyStatesEnclosedParse(readBuffer utils.ReadBuffer, tagNumber uint8) (*BACnetPropertyStatesEnclosed, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesEnclosed"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, pullErr
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParse(readBuffer, uint8(tagNumber))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field")
	}
	openingTag := CastBACnetOpeningTag(_openingTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (propertyState)
	if pullErr := readBuffer.PullContext("propertyState"); pullErr != nil {
		return nil, pullErr
	}
	_propertyState, _propertyStateErr := BACnetPropertyStatesParse(readBuffer)
	if _propertyStateErr != nil {
		return nil, errors.Wrap(_propertyStateErr, "Error parsing 'propertyState' field")
	}
	propertyState := CastBACnetPropertyStates(_propertyState)
	if closeErr := readBuffer.CloseContext("propertyState"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, pullErr
	}
	_closingTag, _closingTagErr := BACnetClosingTagParse(readBuffer, uint8(tagNumber))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field")
	}
	closingTag := CastBACnetClosingTag(_closingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesEnclosed"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewBACnetPropertyStatesEnclosed(openingTag, propertyState, closingTag, tagNumber), nil
}

func (m *BACnetPropertyStatesEnclosed) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetPropertyStatesEnclosed"); pushErr != nil {
		return pushErr
	}

	// Simple Field (openingTag)
	if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
		return pushErr
	}
	_openingTagErr := m.OpeningTag.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
		return popErr
	}
	if _openingTagErr != nil {
		return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
	}

	// Simple Field (propertyState)
	if pushErr := writeBuffer.PushContext("propertyState"); pushErr != nil {
		return pushErr
	}
	_propertyStateErr := m.PropertyState.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("propertyState"); popErr != nil {
		return popErr
	}
	if _propertyStateErr != nil {
		return errors.Wrap(_propertyStateErr, "Error serializing 'propertyState' field")
	}

	// Simple Field (closingTag)
	if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
		return pushErr
	}
	_closingTagErr := m.ClosingTag.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
		return popErr
	}
	if _closingTagErr != nil {
		return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetPropertyStatesEnclosed"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetPropertyStatesEnclosed) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}