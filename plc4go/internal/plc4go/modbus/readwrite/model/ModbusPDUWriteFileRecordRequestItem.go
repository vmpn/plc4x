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
type ModbusPDUWriteFileRecordRequestItem struct {
	ReferenceType uint8
	FileNumber    uint16
	RecordNumber  uint16
	RecordData    []byte
}

// The corresponding interface
type IModbusPDUWriteFileRecordRequestItem interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

func NewModbusPDUWriteFileRecordRequestItem(referenceType uint8, fileNumber uint16, recordNumber uint16, recordData []byte) *ModbusPDUWriteFileRecordRequestItem {
	return &ModbusPDUWriteFileRecordRequestItem{ReferenceType: referenceType, FileNumber: fileNumber, RecordNumber: recordNumber, RecordData: recordData}
}

func CastModbusPDUWriteFileRecordRequestItem(structType interface{}) *ModbusPDUWriteFileRecordRequestItem {
	castFunc := func(typ interface{}) *ModbusPDUWriteFileRecordRequestItem {
		if casted, ok := typ.(ModbusPDUWriteFileRecordRequestItem); ok {
			return &casted
		}
		if casted, ok := typ.(*ModbusPDUWriteFileRecordRequestItem); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ModbusPDUWriteFileRecordRequestItem) GetTypeName() string {
	return "ModbusPDUWriteFileRecordRequestItem"
}

func (m *ModbusPDUWriteFileRecordRequestItem) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *ModbusPDUWriteFileRecordRequestItem) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (referenceType)
	lengthInBits += 8

	// Simple field (fileNumber)
	lengthInBits += 16

	// Simple field (recordNumber)
	lengthInBits += 16

	// Implicit Field (recordLength)
	lengthInBits += 16

	// Array field
	if len(m.RecordData) > 0 {
		lengthInBits += 8 * uint16(len(m.RecordData))
	}

	return lengthInBits
}

func (m *ModbusPDUWriteFileRecordRequestItem) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ModbusPDUWriteFileRecordRequestItemParse(readBuffer utils.ReadBuffer) (*ModbusPDUWriteFileRecordRequestItem, error) {
	if pullErr := readBuffer.PullContext("ModbusPDUWriteFileRecordRequestItem"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (referenceType)
	referenceType, _referenceTypeErr := readBuffer.ReadUint8("referenceType", 8)
	if _referenceTypeErr != nil {
		return nil, errors.Wrap(_referenceTypeErr, "Error parsing 'referenceType' field")
	}

	// Simple Field (fileNumber)
	fileNumber, _fileNumberErr := readBuffer.ReadUint16("fileNumber", 16)
	if _fileNumberErr != nil {
		return nil, errors.Wrap(_fileNumberErr, "Error parsing 'fileNumber' field")
	}

	// Simple Field (recordNumber)
	recordNumber, _recordNumberErr := readBuffer.ReadUint16("recordNumber", 16)
	if _recordNumberErr != nil {
		return nil, errors.Wrap(_recordNumberErr, "Error parsing 'recordNumber' field")
	}

	// Implicit Field (recordLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	recordLength, _recordLengthErr := readBuffer.ReadUint16("recordLength", 16)
	_ = recordLength
	if _recordLengthErr != nil {
		return nil, errors.Wrap(_recordLengthErr, "Error parsing 'recordLength' field")
	}
	// Byte Array field (recordData)
	numberOfBytesrecordData := int(uint16(recordLength) * uint16(uint16(2)))
	recordData, _readArrayErr := readBuffer.ReadByteArray("recordData", numberOfBytesrecordData)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'recordData' field")
	}

	if closeErr := readBuffer.CloseContext("ModbusPDUWriteFileRecordRequestItem"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewModbusPDUWriteFileRecordRequestItem(referenceType, fileNumber, recordNumber, recordData), nil
}

func (m *ModbusPDUWriteFileRecordRequestItem) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("ModbusPDUWriteFileRecordRequestItem"); pushErr != nil {
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

	// Implicit Field (recordLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	recordLength := uint16(uint16(uint16(len(m.RecordData))) / uint16(uint16(2)))
	_recordLengthErr := writeBuffer.WriteUint16("recordLength", 16, (recordLength))
	if _recordLengthErr != nil {
		return errors.Wrap(_recordLengthErr, "Error serializing 'recordLength' field")
	}

	// Array Field (recordData)
	if m.RecordData != nil {
		// Byte Array field (recordData)
		_writeArrayErr := writeBuffer.WriteByteArray("recordData", m.RecordData)
		if _writeArrayErr != nil {
			return errors.Wrap(_writeArrayErr, "Error serializing 'recordData' field")
		}
	}

	if popErr := writeBuffer.PopContext("ModbusPDUWriteFileRecordRequestItem"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *ModbusPDUWriteFileRecordRequestItem) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
