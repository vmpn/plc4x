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
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type BACnetTagPayloadCharacterString struct {
	Encoding BACnetCharacterEncoding
	Value    string

	// Arguments.
	ActualLength uint32
}

// The corresponding interface
type IBACnetTagPayloadCharacterString interface {
	// GetEncoding returns Encoding
	GetEncoding() BACnetCharacterEncoding
	// GetValue returns Value
	GetValue() string
	// GetActualLengthInBit returns ActualLengthInBit
	GetActualLengthInBit() uint16
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
func (m *BACnetTagPayloadCharacterString) GetEncoding() BACnetCharacterEncoding {
	return m.Encoding
}

func (m *BACnetTagPayloadCharacterString) GetValue() string {
	return m.Value
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////
func (m *BACnetTagPayloadCharacterString) GetActualLengthInBit() uint16 {
	return uint16(uint16(m.ActualLength)*uint16(uint16(8))) - uint16(uint16(8))
}

// NewBACnetTagPayloadCharacterString factory function for BACnetTagPayloadCharacterString
func NewBACnetTagPayloadCharacterString(encoding BACnetCharacterEncoding, value string, actualLength uint32) *BACnetTagPayloadCharacterString {
	return &BACnetTagPayloadCharacterString{Encoding: encoding, Value: value, ActualLength: actualLength}
}

func CastBACnetTagPayloadCharacterString(structType interface{}) *BACnetTagPayloadCharacterString {
	castFunc := func(typ interface{}) *BACnetTagPayloadCharacterString {
		if casted, ok := typ.(BACnetTagPayloadCharacterString); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetTagPayloadCharacterString); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetTagPayloadCharacterString) GetTypeName() string {
	return "BACnetTagPayloadCharacterString"
}

func (m *BACnetTagPayloadCharacterString) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetTagPayloadCharacterString) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (encoding)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// Simple field (value)
	lengthInBits += uint16(m.GetActualLengthInBit())

	return lengthInBits
}

func (m *BACnetTagPayloadCharacterString) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetTagPayloadCharacterStringParse(readBuffer utils.ReadBuffer, actualLength uint32) (*BACnetTagPayloadCharacterString, error) {
	if pullErr := readBuffer.PullContext("BACnetTagPayloadCharacterString"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (encoding)
	if pullErr := readBuffer.PullContext("encoding"); pullErr != nil {
		return nil, pullErr
	}
	_encoding, _encodingErr := BACnetCharacterEncodingParse(readBuffer)
	if _encodingErr != nil {
		return nil, errors.Wrap(_encodingErr, "Error parsing 'encoding' field")
	}
	encoding := _encoding
	if closeErr := readBuffer.CloseContext("encoding"); closeErr != nil {
		return nil, closeErr
	}

	// Virtual field
	_actualLengthInBit := uint16(uint16(actualLength)*uint16(uint16(8))) - uint16(uint16(8))
	actualLengthInBit := uint16(_actualLengthInBit)
	_ = actualLengthInBit

	// Simple Field (value)
	_value, _valueErr := readBuffer.ReadString("value", uint32(actualLengthInBit))
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
	}
	value := _value

	if closeErr := readBuffer.CloseContext("BACnetTagPayloadCharacterString"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewBACnetTagPayloadCharacterString(encoding, value, actualLength), nil
}

func (m *BACnetTagPayloadCharacterString) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("BACnetTagPayloadCharacterString"); pushErr != nil {
		return pushErr
	}

	// Simple Field (encoding)
	if pushErr := writeBuffer.PushContext("encoding"); pushErr != nil {
		return pushErr
	}
	_encodingErr := m.Encoding.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("encoding"); popErr != nil {
		return popErr
	}
	if _encodingErr != nil {
		return errors.Wrap(_encodingErr, "Error serializing 'encoding' field")
	}
	// Virtual field
	if _actualLengthInBitErr := writeBuffer.WriteVirtual("actualLengthInBit", m.GetActualLengthInBit()); _actualLengthInBitErr != nil {
		return errors.Wrap(_actualLengthInBitErr, "Error serializing 'actualLengthInBit' field")
	}

	// Simple Field (value)
	value := string(m.Value)
	_valueErr := writeBuffer.WriteString("value", uint32(m.GetActualLengthInBit()), "UTF-8", (value))
	if _valueErr != nil {
		return errors.Wrap(_valueErr, "Error serializing 'value' field")
	}

	if popErr := writeBuffer.PopContext("BACnetTagPayloadCharacterString"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetTagPayloadCharacterString) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
