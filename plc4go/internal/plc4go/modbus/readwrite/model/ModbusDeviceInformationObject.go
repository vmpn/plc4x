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
type ModbusDeviceInformationObject struct {
	ObjectId uint8
	Data     []byte
}

// The corresponding interface
type IModbusDeviceInformationObject interface {
	// GetObjectId returns ObjectId
	GetObjectId() uint8
	// GetData returns Data
	GetData() []byte
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
func (m *ModbusDeviceInformationObject) GetObjectId() uint8 {
	return m.ObjectId
}

func (m *ModbusDeviceInformationObject) GetData() []byte {
	return m.Data
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewModbusDeviceInformationObject factory function for ModbusDeviceInformationObject
func NewModbusDeviceInformationObject(objectId uint8, data []byte) *ModbusDeviceInformationObject {
	return &ModbusDeviceInformationObject{ObjectId: objectId, Data: data}
}

func CastModbusDeviceInformationObject(structType interface{}) *ModbusDeviceInformationObject {
	castFunc := func(typ interface{}) *ModbusDeviceInformationObject {
		if casted, ok := typ.(ModbusDeviceInformationObject); ok {
			return &casted
		}
		if casted, ok := typ.(*ModbusDeviceInformationObject); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ModbusDeviceInformationObject) GetTypeName() string {
	return "ModbusDeviceInformationObject"
}

func (m *ModbusDeviceInformationObject) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ModbusDeviceInformationObject) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (objectId)
	lengthInBits += 8

	// Implicit Field (objectLength)
	lengthInBits += 8

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *ModbusDeviceInformationObject) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ModbusDeviceInformationObjectParse(readBuffer utils.ReadBuffer) (*ModbusDeviceInformationObject, error) {
	if pullErr := readBuffer.PullContext("ModbusDeviceInformationObject"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (objectId)
	_objectId, _objectIdErr := readBuffer.ReadUint8("objectId", 8)
	if _objectIdErr != nil {
		return nil, errors.Wrap(_objectIdErr, "Error parsing 'objectId' field")
	}
	objectId := _objectId

	// Implicit Field (objectLength) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	objectLength, _objectLengthErr := readBuffer.ReadUint8("objectLength", 8)
	_ = objectLength
	if _objectLengthErr != nil {
		return nil, errors.Wrap(_objectLengthErr, "Error parsing 'objectLength' field")
	}
	// Byte Array field (data)
	numberOfBytesdata := int(objectLength)
	data, _readArrayErr := readBuffer.ReadByteArray("data", numberOfBytesdata)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'data' field")
	}

	if closeErr := readBuffer.CloseContext("ModbusDeviceInformationObject"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewModbusDeviceInformationObject(objectId, data), nil
}

func (m *ModbusDeviceInformationObject) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("ModbusDeviceInformationObject"); pushErr != nil {
		return pushErr
	}

	// Simple Field (objectId)
	objectId := uint8(m.ObjectId)
	_objectIdErr := writeBuffer.WriteUint8("objectId", 8, (objectId))
	if _objectIdErr != nil {
		return errors.Wrap(_objectIdErr, "Error serializing 'objectId' field")
	}

	// Implicit Field (objectLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	objectLength := uint8(uint8(len(m.GetData())))
	_objectLengthErr := writeBuffer.WriteUint8("objectLength", 8, (objectLength))
	if _objectLengthErr != nil {
		return errors.Wrap(_objectLengthErr, "Error serializing 'objectLength' field")
	}

	// Array Field (data)
	if m.Data != nil {
		// Byte Array field (data)
		_writeArrayErr := writeBuffer.WriteByteArray("data", m.Data)
		if _writeArrayErr != nil {
			return errors.Wrap(_writeArrayErr, "Error serializing 'data' field")
		}
	}

	if popErr := writeBuffer.PopContext("ModbusDeviceInformationObject"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *ModbusDeviceInformationObject) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
