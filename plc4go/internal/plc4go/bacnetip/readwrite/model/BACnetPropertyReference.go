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
	"github.com/pkg/errors"
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type BACnetPropertyReference struct {
	PropertyIdentifier *BACnetContextTagPropertyIdentifier
	ArrayIndex         *BACnetContextTagUnsignedInteger
}

// The corresponding interface
type IBACnetPropertyReference interface {
	// GetPropertyIdentifier returns PropertyIdentifier
	GetPropertyIdentifier() *BACnetContextTagPropertyIdentifier
	// GetArrayIndex returns ArrayIndex
	GetArrayIndex() *BACnetContextTagUnsignedInteger
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *BACnetPropertyReference) GetPropertyIdentifier() *BACnetContextTagPropertyIdentifier {
	return m.PropertyIdentifier
}

func (m *BACnetPropertyReference) GetArrayIndex() *BACnetContextTagUnsignedInteger {
	return m.ArrayIndex
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewBACnetPropertyReference factory function for BACnetPropertyReference
func NewBACnetPropertyReference(propertyIdentifier *BACnetContextTagPropertyIdentifier, arrayIndex *BACnetContextTagUnsignedInteger) *BACnetPropertyReference {
	return &BACnetPropertyReference{PropertyIdentifier: propertyIdentifier, ArrayIndex: arrayIndex}
}

func CastBACnetPropertyReference(structType interface{}) *BACnetPropertyReference {
	castFunc := func(typ interface{}) *BACnetPropertyReference {
		if casted, ok := typ.(BACnetPropertyReference); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetPropertyReference); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetPropertyReference) GetTypeName() string {
	return "BACnetPropertyReference"
}

func (m *BACnetPropertyReference) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetPropertyReference) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (propertyIdentifier)
	lengthInBits += m.PropertyIdentifier.GetLengthInBits()

	// Optional Field (arrayIndex)
	if m.ArrayIndex != nil {
		lengthInBits += (*m.ArrayIndex).GetLengthInBits()
	}

	return lengthInBits
}

func (m *BACnetPropertyReference) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPropertyReferenceParse(readBuffer utils.ReadBuffer) (*BACnetPropertyReference, error) {
	if pullErr := readBuffer.PullContext("BACnetPropertyReference"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (propertyIdentifier)
	if pullErr := readBuffer.PullContext("propertyIdentifier"); pullErr != nil {
		return nil, pullErr
	}
	_propertyIdentifier, _propertyIdentifierErr := BACnetContextTagParse(readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_BACNET_PROPERTY_IDENTIFIER))
	if _propertyIdentifierErr != nil {
		return nil, errors.Wrap(_propertyIdentifierErr, "Error parsing 'propertyIdentifier' field")
	}
	propertyIdentifier := CastBACnetContextTagPropertyIdentifier(_propertyIdentifier)
	if closeErr := readBuffer.CloseContext("propertyIdentifier"); closeErr != nil {
		return nil, closeErr
	}

	// Optional Field (arrayIndex) (Can be skipped, if a given expression evaluates to false)
	var arrayIndex *BACnetContextTagUnsignedInteger = nil
	{
		currentPos = readBuffer.GetPos()
		if pullErr := readBuffer.PullContext("arrayIndex"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := BACnetContextTagParse(readBuffer, uint8(1), BACnetDataType_UNSIGNED_INTEGER)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'arrayIndex' field")
		default:
			arrayIndex = CastBACnetContextTagUnsignedInteger(_val)
			if closeErr := readBuffer.CloseContext("arrayIndex"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyReference"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewBACnetPropertyReference(propertyIdentifier, arrayIndex), nil
}

func (m *BACnetPropertyReference) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("BACnetPropertyReference"); pushErr != nil {
		return pushErr
	}

	// Simple Field (propertyIdentifier)
	if pushErr := writeBuffer.PushContext("propertyIdentifier"); pushErr != nil {
		return pushErr
	}
	_propertyIdentifierErr := m.PropertyIdentifier.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("propertyIdentifier"); popErr != nil {
		return popErr
	}
	if _propertyIdentifierErr != nil {
		return errors.Wrap(_propertyIdentifierErr, "Error serializing 'propertyIdentifier' field")
	}

	// Optional Field (arrayIndex) (Can be skipped, if the value is null)
	var arrayIndex *BACnetContextTagUnsignedInteger = nil
	if m.ArrayIndex != nil {
		if pushErr := writeBuffer.PushContext("arrayIndex"); pushErr != nil {
			return pushErr
		}
		arrayIndex = m.ArrayIndex
		_arrayIndexErr := arrayIndex.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("arrayIndex"); popErr != nil {
			return popErr
		}
		if _arrayIndexErr != nil {
			return errors.Wrap(_arrayIndexErr, "Error serializing 'arrayIndex' field")
		}
	}

	if popErr := writeBuffer.PopContext("BACnetPropertyReference"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetPropertyReference) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
