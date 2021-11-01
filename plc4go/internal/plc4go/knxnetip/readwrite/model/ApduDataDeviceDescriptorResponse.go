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
type ApduDataDeviceDescriptorResponse struct {
	DescriptorType uint8
	Data           []byte
	Parent         *ApduData
}

// The corresponding interface
type IApduDataDeviceDescriptorResponse interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ApduDataDeviceDescriptorResponse) ApciType() uint8 {
	return 0xD
}

func (m *ApduDataDeviceDescriptorResponse) InitializeParent(parent *ApduData) {
}

func NewApduDataDeviceDescriptorResponse(descriptorType uint8, data []byte) *ApduData {
	child := &ApduDataDeviceDescriptorResponse{
		DescriptorType: descriptorType,
		Data:           data,
		Parent:         NewApduData(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastApduDataDeviceDescriptorResponse(structType interface{}) *ApduDataDeviceDescriptorResponse {
	castFunc := func(typ interface{}) *ApduDataDeviceDescriptorResponse {
		if casted, ok := typ.(ApduDataDeviceDescriptorResponse); ok {
			return &casted
		}
		if casted, ok := typ.(*ApduDataDeviceDescriptorResponse); ok {
			return casted
		}
		if casted, ok := typ.(ApduData); ok {
			return CastApduDataDeviceDescriptorResponse(casted.Child)
		}
		if casted, ok := typ.(*ApduData); ok {
			return CastApduDataDeviceDescriptorResponse(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ApduDataDeviceDescriptorResponse) GetTypeName() string {
	return "ApduDataDeviceDescriptorResponse"
}

func (m *ApduDataDeviceDescriptorResponse) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *ApduDataDeviceDescriptorResponse) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (descriptorType)
	lengthInBits += 6

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *ApduDataDeviceDescriptorResponse) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ApduDataDeviceDescriptorResponseParse(readBuffer utils.ReadBuffer, dataLength uint8) (*ApduData, error) {
	if pullErr := readBuffer.PullContext("ApduDataDeviceDescriptorResponse"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (descriptorType)
	descriptorType, _descriptorTypeErr := readBuffer.ReadUint8("descriptorType", 6)
	if _descriptorTypeErr != nil {
		return nil, errors.Wrap(_descriptorTypeErr, "Error parsing 'descriptorType' field")
	}
	// Byte Array field (data)
	numberOfBytesdata := int(utils.InlineIf(bool(bool((dataLength) < (1))), func() interface{} { return uint16(uint16(0)) }, func() interface{} { return uint16(uint16(dataLength) - uint16(uint16(1))) }).(uint16))
	data, _readArrayErr := readBuffer.ReadByteArray("data", numberOfBytesdata)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'data' field")
	}

	if closeErr := readBuffer.CloseContext("ApduDataDeviceDescriptorResponse"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ApduDataDeviceDescriptorResponse{
		DescriptorType: descriptorType,
		Data:           data,
		Parent:         &ApduData{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *ApduDataDeviceDescriptorResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataDeviceDescriptorResponse"); pushErr != nil {
			return pushErr
		}

		// Simple Field (descriptorType)
		descriptorType := uint8(m.DescriptorType)
		_descriptorTypeErr := writeBuffer.WriteUint8("descriptorType", 6, (descriptorType))
		if _descriptorTypeErr != nil {
			return errors.Wrap(_descriptorTypeErr, "Error serializing 'descriptorType' field")
		}

		// Array Field (data)
		if m.Data != nil {
			// Byte Array field (data)
			_writeArrayErr := writeBuffer.WriteByteArray("data", m.Data)
			if _writeArrayErr != nil {
				return errors.Wrap(_writeArrayErr, "Error serializing 'data' field")
			}
		}

		if popErr := writeBuffer.PopContext("ApduDataDeviceDescriptorResponse"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *ApduDataDeviceDescriptorResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
