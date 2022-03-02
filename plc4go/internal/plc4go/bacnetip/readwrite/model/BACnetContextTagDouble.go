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
type BACnetContextTagDouble struct {
	*BACnetContextTag
	Payload *BACnetTagPayloadDouble

	// Arguments.
	TagNumberArgument        uint8
	IsNotOpeningOrClosingTag bool
}

// The corresponding interface
type IBACnetContextTagDouble interface {
	// GetPayload returns Payload
	GetPayload() *BACnetTagPayloadDouble
	// GetActualValue returns ActualValue
	GetActualValue() float64
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetContextTagDouble) DataType() BACnetDataType {
	return BACnetDataType_DOUBLE
}

func (m *BACnetContextTagDouble) GetDataType() BACnetDataType {
	return BACnetDataType_DOUBLE
}

func (m *BACnetContextTagDouble) InitializeParent(parent *BACnetContextTag, header *BACnetTagHeader) {
	m.BACnetContextTag.Header = header
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *BACnetContextTagDouble) GetPayload() *BACnetTagPayloadDouble {
	return m.Payload
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////
func (m *BACnetContextTagDouble) GetActualValue() float64 {
	return m.GetPayload().GetValue()
}

// NewBACnetContextTagDouble factory function for BACnetContextTagDouble
func NewBACnetContextTagDouble(payload *BACnetTagPayloadDouble, header *BACnetTagHeader, tagNumberArgument uint8, isNotOpeningOrClosingTag bool) *BACnetContextTag {
	child := &BACnetContextTagDouble{
		Payload:          payload,
		BACnetContextTag: NewBACnetContextTag(header, tagNumberArgument),
	}
	child.Child = child
	return child.BACnetContextTag
}

func CastBACnetContextTagDouble(structType interface{}) *BACnetContextTagDouble {
	castFunc := func(typ interface{}) *BACnetContextTagDouble {
		if casted, ok := typ.(BACnetContextTagDouble); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetContextTagDouble); ok {
			return casted
		}
		if casted, ok := typ.(BACnetContextTag); ok {
			return CastBACnetContextTagDouble(casted.Child)
		}
		if casted, ok := typ.(*BACnetContextTag); ok {
			return CastBACnetContextTagDouble(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetContextTagDouble) GetTypeName() string {
	return "BACnetContextTagDouble"
}

func (m *BACnetContextTagDouble) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetContextTagDouble) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetContextTagDouble) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetContextTagDoubleParse(readBuffer utils.ReadBuffer, tagNumberArgument uint8, dataType BACnetDataType, isNotOpeningOrClosingTag bool) (*BACnetContextTag, error) {
	if pullErr := readBuffer.PullContext("BACnetContextTagDouble"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Validation
	if !(isNotOpeningOrClosingTag) {
		return nil, utils.ParseAssertError{"length 6 and 7 reserved for opening and closing tag"}
	}

	// Simple Field (payload)
	if pullErr := readBuffer.PullContext("payload"); pullErr != nil {
		return nil, pullErr
	}
	_payload, _payloadErr := BACnetTagPayloadDoubleParse(readBuffer)
	if _payloadErr != nil {
		return nil, errors.Wrap(_payloadErr, "Error parsing 'payload' field")
	}
	payload := CastBACnetTagPayloadDouble(_payload)
	if closeErr := readBuffer.CloseContext("payload"); closeErr != nil {
		return nil, closeErr
	}

	// Virtual field
	_actualValue := payload.GetValue()
	actualValue := float64(_actualValue)
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetContextTagDouble"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetContextTagDouble{
		Payload:          CastBACnetTagPayloadDouble(payload),
		BACnetContextTag: &BACnetContextTag{},
	}
	_child.BACnetContextTag.Child = _child
	return _child.BACnetContextTag, nil
}

func (m *BACnetContextTagDouble) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetContextTagDouble"); pushErr != nil {
			return pushErr
		}

		// Simple Field (payload)
		if pushErr := writeBuffer.PushContext("payload"); pushErr != nil {
			return pushErr
		}
		_payloadErr := m.Payload.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("payload"); popErr != nil {
			return popErr
		}
		if _payloadErr != nil {
			return errors.Wrap(_payloadErr, "Error serializing 'payload' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetContextTagDouble"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetContextTagDouble) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
