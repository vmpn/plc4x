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

// COTPParameterDisconnectAdditionalInformation is the data-structure of this message
type COTPParameterDisconnectAdditionalInformation struct {
	*COTPParameter
	Data []byte

	// Arguments.
	Rest uint8
}

// ICOTPParameterDisconnectAdditionalInformation is the corresponding interface of COTPParameterDisconnectAdditionalInformation
type ICOTPParameterDisconnectAdditionalInformation interface {
	ICOTPParameter
	// GetData returns Data (property field)
	GetData() []byte
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

func (m *COTPParameterDisconnectAdditionalInformation) GetParameterType() uint8 {
	return 0xE0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *COTPParameterDisconnectAdditionalInformation) InitializeParent(parent *COTPParameter) {}

func (m *COTPParameterDisconnectAdditionalInformation) GetParent() *COTPParameter {
	return m.COTPParameter
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *COTPParameterDisconnectAdditionalInformation) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCOTPParameterDisconnectAdditionalInformation factory function for COTPParameterDisconnectAdditionalInformation
func NewCOTPParameterDisconnectAdditionalInformation(data []byte, rest uint8) *COTPParameterDisconnectAdditionalInformation {
	_result := &COTPParameterDisconnectAdditionalInformation{
		Data:          data,
		COTPParameter: NewCOTPParameter(rest),
	}
	_result.Child = _result
	return _result
}

func CastCOTPParameterDisconnectAdditionalInformation(structType interface{}) *COTPParameterDisconnectAdditionalInformation {
	if casted, ok := structType.(COTPParameterDisconnectAdditionalInformation); ok {
		return &casted
	}
	if casted, ok := structType.(*COTPParameterDisconnectAdditionalInformation); ok {
		return casted
	}
	if casted, ok := structType.(COTPParameter); ok {
		return CastCOTPParameterDisconnectAdditionalInformation(casted.Child)
	}
	if casted, ok := structType.(*COTPParameter); ok {
		return CastCOTPParameterDisconnectAdditionalInformation(casted.Child)
	}
	return nil
}

func (m *COTPParameterDisconnectAdditionalInformation) GetTypeName() string {
	return "COTPParameterDisconnectAdditionalInformation"
}

func (m *COTPParameterDisconnectAdditionalInformation) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *COTPParameterDisconnectAdditionalInformation) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *COTPParameterDisconnectAdditionalInformation) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func COTPParameterDisconnectAdditionalInformationParse(readBuffer utils.ReadBuffer, rest uint8) (*COTPParameterDisconnectAdditionalInformation, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("COTPParameterDisconnectAdditionalInformation"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos
	// Byte Array field (data)
	numberOfBytesdata := int(rest)
	data, _readArrayErr := readBuffer.ReadByteArray("data", numberOfBytesdata)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'data' field")
	}

	if closeErr := readBuffer.CloseContext("COTPParameterDisconnectAdditionalInformation"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &COTPParameterDisconnectAdditionalInformation{
		Data:          data,
		COTPParameter: &COTPParameter{},
	}
	_child.COTPParameter.Child = _child
	return _child, nil
}

func (m *COTPParameterDisconnectAdditionalInformation) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("COTPParameterDisconnectAdditionalInformation"); pushErr != nil {
			return pushErr
		}

		// Array Field (data)
		if m.Data != nil {
			// Byte Array field (data)
			_writeArrayErr := writeBuffer.WriteByteArray("data", m.Data)
			if _writeArrayErr != nil {
				return errors.Wrap(_writeArrayErr, "Error serializing 'data' field")
			}
		}

		if popErr := writeBuffer.PopContext("COTPParameterDisconnectAdditionalInformation"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *COTPParameterDisconnectAdditionalInformation) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}