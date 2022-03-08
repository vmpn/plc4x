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
	"fmt"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const S7DataAlarmMessage_FUNCTIONID uint8 = 0x00
const S7DataAlarmMessage_NUMBERMESSAGEOBJ uint8 = 0x01

// The data-structure of this message
type S7DataAlarmMessage struct {
	Child IS7DataAlarmMessageChild
}

// The corresponding interface
type IS7DataAlarmMessage interface {
	// GetCpuFunctionType returns CpuFunctionType (discriminator field)
	GetCpuFunctionType() uint8
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

type IS7DataAlarmMessageParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child IS7DataAlarmMessage, serializeChildFunction func() error) error
	GetTypeName() string
}

type IS7DataAlarmMessageChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *S7DataAlarmMessage)
	GetParent() *S7DataAlarmMessage

	GetTypeName() string
	IS7DataAlarmMessage
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for conts fields.
///////////////////////
func (m *S7DataAlarmMessage) GetFunctionId() uint8 {
	return S7DataAlarmMessage_FUNCTIONID
}

func (m *S7DataAlarmMessage) GetNumberMessageObj() uint8 {
	return S7DataAlarmMessage_NUMBERMESSAGEOBJ
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewS7DataAlarmMessage factory function for S7DataAlarmMessage
func NewS7DataAlarmMessage() *S7DataAlarmMessage {
	return &S7DataAlarmMessage{}
}

func CastS7DataAlarmMessage(structType interface{}) *S7DataAlarmMessage {
	if casted, ok := structType.(S7DataAlarmMessage); ok {
		return &casted
	}
	if casted, ok := structType.(*S7DataAlarmMessage); ok {
		return casted
	}
	if casted, ok := structType.(IS7DataAlarmMessageChild); ok {
		return casted.GetParent()
	}
	return nil
}

func (m *S7DataAlarmMessage) GetTypeName() string {
	return "S7DataAlarmMessage"
}

func (m *S7DataAlarmMessage) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *S7DataAlarmMessage) GetLengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.GetLengthInBits()
}

func (m *S7DataAlarmMessage) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Const Field (functionId)
	lengthInBits += 8

	// Const Field (numberMessageObj)
	lengthInBits += 8

	return lengthInBits
}

func (m *S7DataAlarmMessage) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func S7DataAlarmMessageParse(readBuffer utils.ReadBuffer, cpuFunctionType uint8) (*S7DataAlarmMessage, error) {
	if pullErr := readBuffer.PullContext("S7DataAlarmMessage"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Const Field (functionId)
	functionId, _functionIdErr := readBuffer.ReadUint8("functionId", 8)
	if _functionIdErr != nil {
		return nil, errors.Wrap(_functionIdErr, "Error parsing 'functionId' field")
	}
	if functionId != S7DataAlarmMessage_FUNCTIONID {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", S7DataAlarmMessage_FUNCTIONID) + " but got " + fmt.Sprintf("%d", functionId))
	}

	// Const Field (numberMessageObj)
	numberMessageObj, _numberMessageObjErr := readBuffer.ReadUint8("numberMessageObj", 8)
	if _numberMessageObjErr != nil {
		return nil, errors.Wrap(_numberMessageObjErr, "Error parsing 'numberMessageObj' field")
	}
	if numberMessageObj != S7DataAlarmMessage_NUMBERMESSAGEOBJ {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", S7DataAlarmMessage_NUMBERMESSAGEOBJ) + " but got " + fmt.Sprintf("%d", numberMessageObj))
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type S7DataAlarmMessageChild interface {
		InitializeParent(*S7DataAlarmMessage)
		GetParent() *S7DataAlarmMessage
	}
	var _child S7DataAlarmMessageChild
	var typeSwitchError error
	switch {
	case cpuFunctionType == 0x04: // S7MessageObjectRequest
		_child, typeSwitchError = S7MessageObjectRequestParse(readBuffer, cpuFunctionType)
	case cpuFunctionType == 0x08: // S7MessageObjectResponse
		_child, typeSwitchError = S7MessageObjectResponseParse(readBuffer, cpuFunctionType)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("S7DataAlarmMessage"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_child.InitializeParent(_child.GetParent())
	return _child.GetParent(), nil
}

func (m *S7DataAlarmMessage) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *S7DataAlarmMessage) SerializeParent(writeBuffer utils.WriteBuffer, child IS7DataAlarmMessage, serializeChildFunction func() error) error {
	if pushErr := writeBuffer.PushContext("S7DataAlarmMessage"); pushErr != nil {
		return pushErr
	}

	// Const Field (functionId)
	_functionIdErr := writeBuffer.WriteUint8("functionId", 8, 0x00)
	if _functionIdErr != nil {
		return errors.Wrap(_functionIdErr, "Error serializing 'functionId' field")
	}

	// Const Field (numberMessageObj)
	_numberMessageObjErr := writeBuffer.WriteUint8("numberMessageObj", 8, 0x01)
	if _numberMessageObjErr != nil {
		return errors.Wrap(_numberMessageObjErr, "Error serializing 'numberMessageObj' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("S7DataAlarmMessage"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *S7DataAlarmMessage) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
