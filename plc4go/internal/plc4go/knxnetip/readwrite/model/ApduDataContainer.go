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
type ApduDataContainer struct {
	*Apdu
	DataApdu *ApduData

	// Arguments.
	DataLength uint8
}

// The corresponding interface
type IApduDataContainer interface {
	// GetDataApdu returns DataApdu
	GetDataApdu() *ApduData
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
func (m *ApduDataContainer) Control() uint8 {
	return uint8(0)
}

func (m *ApduDataContainer) GetControl() uint8 {
	return uint8(0)
}

func (m *ApduDataContainer) InitializeParent(parent *Apdu, numbered bool, counter uint8) {
	m.Apdu.Numbered = numbered
	m.Apdu.Counter = counter
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *ApduDataContainer) GetDataApdu() *ApduData {
	return m.DataApdu
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewApduDataContainer factory function for ApduDataContainer
func NewApduDataContainer(dataApdu *ApduData, numbered bool, counter uint8, dataLength uint8) *Apdu {
	child := &ApduDataContainer{
		DataApdu: dataApdu,
		Apdu:     NewApdu(numbered, counter, dataLength),
	}
	child.Child = child
	return child.Apdu
}

func CastApduDataContainer(structType interface{}) *ApduDataContainer {
	castFunc := func(typ interface{}) *ApduDataContainer {
		if casted, ok := typ.(ApduDataContainer); ok {
			return &casted
		}
		if casted, ok := typ.(*ApduDataContainer); ok {
			return casted
		}
		if casted, ok := typ.(Apdu); ok {
			return CastApduDataContainer(casted.Child)
		}
		if casted, ok := typ.(*Apdu); ok {
			return CastApduDataContainer(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ApduDataContainer) GetTypeName() string {
	return "ApduDataContainer"
}

func (m *ApduDataContainer) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ApduDataContainer) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (dataApdu)
	lengthInBits += m.DataApdu.GetLengthInBits()

	return lengthInBits
}

func (m *ApduDataContainer) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApduDataContainerParse(readBuffer utils.ReadBuffer, dataLength uint8) (*Apdu, error) {
	if pullErr := readBuffer.PullContext("ApduDataContainer"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (dataApdu)
	if pullErr := readBuffer.PullContext("dataApdu"); pullErr != nil {
		return nil, pullErr
	}
	_dataApdu, _dataApduErr := ApduDataParse(readBuffer, uint8(dataLength))
	if _dataApduErr != nil {
		return nil, errors.Wrap(_dataApduErr, "Error parsing 'dataApdu' field")
	}
	dataApdu := CastApduData(_dataApdu)
	if closeErr := readBuffer.CloseContext("dataApdu"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("ApduDataContainer"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ApduDataContainer{
		DataApdu: CastApduData(dataApdu),
		Apdu:     &Apdu{},
	}
	_child.Apdu.Child = _child
	return _child.Apdu, nil
}

func (m *ApduDataContainer) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataContainer"); pushErr != nil {
			return pushErr
		}

		// Simple Field (dataApdu)
		if pushErr := writeBuffer.PushContext("dataApdu"); pushErr != nil {
			return pushErr
		}
		_dataApduErr := m.DataApdu.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("dataApdu"); popErr != nil {
			return popErr
		}
		if _dataApduErr != nil {
			return errors.Wrap(_dataApduErr, "Error serializing 'dataApdu' field")
		}

		if popErr := writeBuffer.PopContext("ApduDataContainer"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ApduDataContainer) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
