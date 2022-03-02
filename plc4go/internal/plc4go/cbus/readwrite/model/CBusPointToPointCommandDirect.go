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
	"github.com/rs/zerolog/log"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type CBusPointToPointCommandDirect struct {
	*CBusPointToPointCommand
	UnitAddress *UnitAddress

	// Arguments.
	Srchk bool
}

// The corresponding interface
type ICBusPointToPointCommandDirect interface {
	// GetUnitAddress returns UnitAddress
	GetUnitAddress() *UnitAddress
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
func (m *CBusPointToPointCommandDirect) IsDirect() bool {
	return bool(true)
}

func (m *CBusPointToPointCommandDirect) GetIsDirect() bool {
	return bool(true)
}

func (m *CBusPointToPointCommandDirect) InitializeParent(parent *CBusPointToPointCommand, bridgeAddressCountPeek uint16, calData *CALData, crc *Checksum, peekAlpha byte, alpha *Alpha) {
	m.CBusPointToPointCommand.BridgeAddressCountPeek = bridgeAddressCountPeek
	m.CBusPointToPointCommand.CalData = calData
	m.CBusPointToPointCommand.Crc = crc
	m.CBusPointToPointCommand.PeekAlpha = peekAlpha
	m.CBusPointToPointCommand.Alpha = alpha
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *CBusPointToPointCommandDirect) GetUnitAddress() *UnitAddress {
	return m.UnitAddress
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewCBusPointToPointCommandDirect factory function for CBusPointToPointCommandDirect
func NewCBusPointToPointCommandDirect(unitAddress *UnitAddress, bridgeAddressCountPeek uint16, calData *CALData, crc *Checksum, peekAlpha byte, alpha *Alpha, srchk bool) *CBusPointToPointCommand {
	child := &CBusPointToPointCommandDirect{
		UnitAddress:             unitAddress,
		CBusPointToPointCommand: NewCBusPointToPointCommand(bridgeAddressCountPeek, calData, crc, peekAlpha, alpha, srchk),
	}
	child.Child = child
	return child.CBusPointToPointCommand
}

func CastCBusPointToPointCommandDirect(structType interface{}) *CBusPointToPointCommandDirect {
	castFunc := func(typ interface{}) *CBusPointToPointCommandDirect {
		if casted, ok := typ.(CBusPointToPointCommandDirect); ok {
			return &casted
		}
		if casted, ok := typ.(*CBusPointToPointCommandDirect); ok {
			return casted
		}
		if casted, ok := typ.(CBusPointToPointCommand); ok {
			return CastCBusPointToPointCommandDirect(casted.Child)
		}
		if casted, ok := typ.(*CBusPointToPointCommand); ok {
			return CastCBusPointToPointCommandDirect(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *CBusPointToPointCommandDirect) GetTypeName() string {
	return "CBusPointToPointCommandDirect"
}

func (m *CBusPointToPointCommandDirect) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *CBusPointToPointCommandDirect) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (unitAddress)
	lengthInBits += m.UnitAddress.GetLengthInBits()

	// Reserved Field (reserved)
	lengthInBits += 8

	return lengthInBits
}

func (m *CBusPointToPointCommandDirect) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CBusPointToPointCommandDirectParse(readBuffer utils.ReadBuffer, srchk bool) (*CBusPointToPointCommand, error) {
	if pullErr := readBuffer.PullContext("CBusPointToPointCommandDirect"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (unitAddress)
	if pullErr := readBuffer.PullContext("unitAddress"); pullErr != nil {
		return nil, pullErr
	}
	_unitAddress, _unitAddressErr := UnitAddressParse(readBuffer)
	if _unitAddressErr != nil {
		return nil, errors.Wrap(_unitAddressErr, "Error parsing 'unitAddress' field")
	}
	unitAddress := CastUnitAddress(_unitAddress)
	if closeErr := readBuffer.CloseContext("unitAddress"); closeErr != nil {
		return nil, closeErr
	}

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != uint8(0x00) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint8(0x00),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	if closeErr := readBuffer.CloseContext("CBusPointToPointCommandDirect"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &CBusPointToPointCommandDirect{
		UnitAddress:             CastUnitAddress(unitAddress),
		CBusPointToPointCommand: &CBusPointToPointCommand{},
	}
	_child.CBusPointToPointCommand.Child = _child
	return _child.CBusPointToPointCommand, nil
}

func (m *CBusPointToPointCommandDirect) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CBusPointToPointCommandDirect"); pushErr != nil {
			return pushErr
		}

		// Simple Field (unitAddress)
		if pushErr := writeBuffer.PushContext("unitAddress"); pushErr != nil {
			return pushErr
		}
		_unitAddressErr := m.UnitAddress.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("unitAddress"); popErr != nil {
			return popErr
		}
		if _unitAddressErr != nil {
			return errors.Wrap(_unitAddressErr, "Error serializing 'unitAddress' field")
		}

		// Reserved Field (reserved)
		{
			_err := writeBuffer.WriteUint8("reserved", 8, uint8(0x00))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		if popErr := writeBuffer.PopContext("CBusPointToPointCommandDirect"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *CBusPointToPointCommandDirect) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
