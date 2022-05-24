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

// S7PayloadNotify is the data-structure of this message
type S7PayloadNotify struct {
	*S7PayloadUserDataItem
	AlarmMessage *AlarmMessagePushType
}

// IS7PayloadNotify is the corresponding interface of S7PayloadNotify
type IS7PayloadNotify interface {
	IS7PayloadUserDataItem
	// GetAlarmMessage returns AlarmMessage (property field)
	GetAlarmMessage() *AlarmMessagePushType
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

func (m *S7PayloadNotify) GetCpuFunctionType() uint8 {
	return 0x00
}

func (m *S7PayloadNotify) GetCpuSubfunction() uint8 {
	return 0x06
}

func (m *S7PayloadNotify) GetDataLength() uint16 {
	return 0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *S7PayloadNotify) InitializeParent(parent *S7PayloadUserDataItem, returnCode DataTransportErrorCode, transportSize DataTransportSize) {
	m.S7PayloadUserDataItem.ReturnCode = returnCode
	m.S7PayloadUserDataItem.TransportSize = transportSize
}

func (m *S7PayloadNotify) GetParent() *S7PayloadUserDataItem {
	return m.S7PayloadUserDataItem
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *S7PayloadNotify) GetAlarmMessage() *AlarmMessagePushType {
	return m.AlarmMessage
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewS7PayloadNotify factory function for S7PayloadNotify
func NewS7PayloadNotify(alarmMessage *AlarmMessagePushType, returnCode DataTransportErrorCode, transportSize DataTransportSize) *S7PayloadNotify {
	_result := &S7PayloadNotify{
		AlarmMessage:          alarmMessage,
		S7PayloadUserDataItem: NewS7PayloadUserDataItem(returnCode, transportSize),
	}
	_result.Child = _result
	return _result
}

func CastS7PayloadNotify(structType interface{}) *S7PayloadNotify {
	if casted, ok := structType.(S7PayloadNotify); ok {
		return &casted
	}
	if casted, ok := structType.(*S7PayloadNotify); ok {
		return casted
	}
	if casted, ok := structType.(S7PayloadUserDataItem); ok {
		return CastS7PayloadNotify(casted.Child)
	}
	if casted, ok := structType.(*S7PayloadUserDataItem); ok {
		return CastS7PayloadNotify(casted.Child)
	}
	return nil
}

func (m *S7PayloadNotify) GetTypeName() string {
	return "S7PayloadNotify"
}

func (m *S7PayloadNotify) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *S7PayloadNotify) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (alarmMessage)
	lengthInBits += m.AlarmMessage.GetLengthInBits()

	return lengthInBits
}

func (m *S7PayloadNotify) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func S7PayloadNotifyParse(readBuffer utils.ReadBuffer, cpuFunctionType uint8, cpuSubfunction uint8) (*S7PayloadNotify, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7PayloadNotify"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (alarmMessage)
	if pullErr := readBuffer.PullContext("alarmMessage"); pullErr != nil {
		return nil, pullErr
	}
	_alarmMessage, _alarmMessageErr := AlarmMessagePushTypeParse(readBuffer)
	if _alarmMessageErr != nil {
		return nil, errors.Wrap(_alarmMessageErr, "Error parsing 'alarmMessage' field")
	}
	alarmMessage := CastAlarmMessagePushType(_alarmMessage)
	if closeErr := readBuffer.CloseContext("alarmMessage"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("S7PayloadNotify"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &S7PayloadNotify{
		AlarmMessage:          CastAlarmMessagePushType(alarmMessage),
		S7PayloadUserDataItem: &S7PayloadUserDataItem{},
	}
	_child.S7PayloadUserDataItem.Child = _child
	return _child, nil
}

func (m *S7PayloadNotify) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadNotify"); pushErr != nil {
			return pushErr
		}

		// Simple Field (alarmMessage)
		if pushErr := writeBuffer.PushContext("alarmMessage"); pushErr != nil {
			return pushErr
		}
		_alarmMessageErr := m.AlarmMessage.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("alarmMessage"); popErr != nil {
			return popErr
		}
		if _alarmMessageErr != nil {
			return errors.Wrap(_alarmMessageErr, "Error serializing 'alarmMessage' field")
		}

		if popErr := writeBuffer.PopContext("S7PayloadNotify"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *S7PayloadNotify) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}