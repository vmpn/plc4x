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

// ModbusPDUReadFileRecordRequestItem is the data-structure of this message
type ModbusPDUReadFileRecordRequestItem struct {
	ReferenceType uint8
	FileNumber    uint16
	RecordNumber  uint16
	RecordLength  uint16
}

// IModbusPDUReadFileRecordRequestItem is the corresponding interface of ModbusPDUReadFileRecordRequestItem
type IModbusPDUReadFileRecordRequestItem interface {
	// GetReferenceType returns ReferenceType (property field)
	GetReferenceType() uint8
	// GetFileNumber returns FileNumber (property field)
	GetFileNumber() uint16
	// GetRecordNumber returns RecordNumber (property field)
	GetRecordNumber() uint16
	// GetRecordLength returns RecordLength (property field)
	GetRecordLength() uint16
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

func (m *ModbusPDUReadFileRecordRequestItem) GetReferenceType() uint8 {
	return m.ReferenceType
}

func (m *ModbusPDUReadFileRecordRequestItem) GetFileNumber() uint16 {
	return m.FileNumber
}

func (m *ModbusPDUReadFileRecordRequestItem) GetRecordNumber() uint16 {
	return m.RecordNumber
}

func (m *ModbusPDUReadFileRecordRequestItem) GetRecordLength() uint16 {
	return m.RecordLength
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewModbusPDUReadFileRecordRequestItem factory function for ModbusPDUReadFileRecordRequestItem
func NewModbusPDUReadFileRecordRequestItem(referenceType uint8, fileNumber uint16, recordNumber uint16, recordLength uint16) *ModbusPDUReadFileRecordRequestItem {
	return &ModbusPDUReadFileRecordRequestItem{ReferenceType: referenceType, FileNumber: fileNumber, RecordNumber: recordNumber, RecordLength: recordLength}
}

func CastModbusPDUReadFileRecordRequestItem(structType interface{}) *ModbusPDUReadFileRecordRequestItem {
	if casted, ok := structType.(ModbusPDUReadFileRecordRequestItem); ok {
		return &casted
	}
	if casted, ok := structType.(*ModbusPDUReadFileRecordRequestItem); ok {
		return casted
	}
	return nil
}

func (m *ModbusPDUReadFileRecordRequestItem) GetTypeName() string {
	return "ModbusPDUReadFileRecordRequestItem"
}

func (m *ModbusPDUReadFileRecordRequestItem) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ModbusPDUReadFileRecordRequestItem) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (referenceType)
	lengthInBits += 8

	// Simple field (fileNumber)
	lengthInBits += 16

	// Simple field (recordNumber)
	lengthInBits += 16

	// Simple field (recordLength)
	lengthInBits += 16

	return lengthInBits
}

func (m *ModbusPDUReadFileRecordRequestItem) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ModbusPDUReadFileRecordRequestItemParse(readBuffer utils.ReadBuffer) (*ModbusPDUReadFileRecordRequestItem, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUReadFileRecordRequestItem"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (referenceType)
	_referenceType, _referenceTypeErr := readBuffer.ReadUint8("referenceType", 8)
	if _referenceTypeErr != nil {
		return nil, errors.Wrap(_referenceTypeErr, "Error parsing 'referenceType' field")
	}
	referenceType := _referenceType

	// Simple Field (fileNumber)
	_fileNumber, _fileNumberErr := readBuffer.ReadUint16("fileNumber", 16)
	if _fileNumberErr != nil {
		return nil, errors.Wrap(_fileNumberErr, "Error parsing 'fileNumber' field")
	}
	fileNumber := _fileNumber

	// Simple Field (recordNumber)
	_recordNumber, _recordNumberErr := readBuffer.ReadUint16("recordNumber", 16)
	if _recordNumberErr != nil {
		return nil, errors.Wrap(_recordNumberErr, "Error parsing 'recordNumber' field")
	}
	recordNumber := _recordNumber

	// Simple Field (recordLength)
	_recordLength, _recordLengthErr := readBuffer.ReadUint16("recordLength", 16)
	if _recordLengthErr != nil {
		return nil, errors.Wrap(_recordLengthErr, "Error parsing 'recordLength' field")
	}
	recordLength := _recordLength

	if closeErr := readBuffer.CloseContext("ModbusPDUReadFileRecordRequestItem"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewModbusPDUReadFileRecordRequestItem(referenceType, fileNumber, recordNumber, recordLength), nil
}

func (m *ModbusPDUReadFileRecordRequestItem) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("ModbusPDUReadFileRecordRequestItem"); pushErr != nil {
		return pushErr
	}

	// Simple Field (referenceType)
	referenceType := uint8(m.ReferenceType)
	_referenceTypeErr := writeBuffer.WriteUint8("referenceType", 8, (referenceType))
	if _referenceTypeErr != nil {
		return errors.Wrap(_referenceTypeErr, "Error serializing 'referenceType' field")
	}

	// Simple Field (fileNumber)
	fileNumber := uint16(m.FileNumber)
	_fileNumberErr := writeBuffer.WriteUint16("fileNumber", 16, (fileNumber))
	if _fileNumberErr != nil {
		return errors.Wrap(_fileNumberErr, "Error serializing 'fileNumber' field")
	}

	// Simple Field (recordNumber)
	recordNumber := uint16(m.RecordNumber)
	_recordNumberErr := writeBuffer.WriteUint16("recordNumber", 16, (recordNumber))
	if _recordNumberErr != nil {
		return errors.Wrap(_recordNumberErr, "Error serializing 'recordNumber' field")
	}

	// Simple Field (recordLength)
	recordLength := uint16(m.RecordLength)
	_recordLengthErr := writeBuffer.WriteUint16("recordLength", 16, (recordLength))
	if _recordLengthErr != nil {
		return errors.Wrap(_recordLengthErr, "Error serializing 'recordLength' field")
	}

	if popErr := writeBuffer.PopContext("ModbusPDUReadFileRecordRequestItem"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *ModbusPDUReadFileRecordRequestItem) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}