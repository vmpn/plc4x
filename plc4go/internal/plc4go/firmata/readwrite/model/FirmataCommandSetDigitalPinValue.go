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
type FirmataCommandSetDigitalPinValue struct {
	*FirmataCommand
	Pin uint8
	On  bool

	// Arguments.
	Response bool
}

// The corresponding interface
type IFirmataCommandSetDigitalPinValue interface {
	// GetPin returns Pin
	GetPin() uint8
	// GetOn returns On
	GetOn() bool
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
func (m *FirmataCommandSetDigitalPinValue) CommandCode() uint8 {
	return 0x5
}

func (m *FirmataCommandSetDigitalPinValue) GetCommandCode() uint8 {
	return 0x5
}

func (m *FirmataCommandSetDigitalPinValue) InitializeParent(parent *FirmataCommand) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *FirmataCommandSetDigitalPinValue) GetPin() uint8 {
	return m.Pin
}

func (m *FirmataCommandSetDigitalPinValue) GetOn() bool {
	return m.On
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewFirmataCommandSetDigitalPinValue factory function for FirmataCommandSetDigitalPinValue
func NewFirmataCommandSetDigitalPinValue(pin uint8, on bool, response bool) *FirmataCommand {
	child := &FirmataCommandSetDigitalPinValue{
		Pin:            pin,
		On:             on,
		FirmataCommand: NewFirmataCommand(response),
	}
	child.Child = child
	return child.FirmataCommand
}

func CastFirmataCommandSetDigitalPinValue(structType interface{}) *FirmataCommandSetDigitalPinValue {
	castFunc := func(typ interface{}) *FirmataCommandSetDigitalPinValue {
		if casted, ok := typ.(FirmataCommandSetDigitalPinValue); ok {
			return &casted
		}
		if casted, ok := typ.(*FirmataCommandSetDigitalPinValue); ok {
			return casted
		}
		if casted, ok := typ.(FirmataCommand); ok {
			return CastFirmataCommandSetDigitalPinValue(casted.Child)
		}
		if casted, ok := typ.(*FirmataCommand); ok {
			return CastFirmataCommandSetDigitalPinValue(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *FirmataCommandSetDigitalPinValue) GetTypeName() string {
	return "FirmataCommandSetDigitalPinValue"
}

func (m *FirmataCommandSetDigitalPinValue) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *FirmataCommandSetDigitalPinValue) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (pin)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (on)
	lengthInBits += 1

	return lengthInBits
}

func (m *FirmataCommandSetDigitalPinValue) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func FirmataCommandSetDigitalPinValueParse(readBuffer utils.ReadBuffer, response bool) (*FirmataCommand, error) {
	if pullErr := readBuffer.PullContext("FirmataCommandSetDigitalPinValue"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (pin)
	_pin, _pinErr := readBuffer.ReadUint8("pin", 8)
	if _pinErr != nil {
		return nil, errors.Wrap(_pinErr, "Error parsing 'pin' field")
	}
	pin := _pin

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 7)
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

	// Simple Field (on)
	_on, _onErr := readBuffer.ReadBit("on")
	if _onErr != nil {
		return nil, errors.Wrap(_onErr, "Error parsing 'on' field")
	}
	on := _on

	if closeErr := readBuffer.CloseContext("FirmataCommandSetDigitalPinValue"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &FirmataCommandSetDigitalPinValue{
		Pin:            pin,
		On:             on,
		FirmataCommand: &FirmataCommand{},
	}
	_child.FirmataCommand.Child = _child
	return _child.FirmataCommand, nil
}

func (m *FirmataCommandSetDigitalPinValue) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("FirmataCommandSetDigitalPinValue"); pushErr != nil {
			return pushErr
		}

		// Simple Field (pin)
		pin := uint8(m.Pin)
		_pinErr := writeBuffer.WriteUint8("pin", 8, (pin))
		if _pinErr != nil {
			return errors.Wrap(_pinErr, "Error serializing 'pin' field")
		}

		// Reserved Field (reserved)
		{
			_err := writeBuffer.WriteUint8("reserved", 7, uint8(0x00))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (on)
		on := bool(m.On)
		_onErr := writeBuffer.WriteBit("on", (on))
		if _onErr != nil {
			return errors.Wrap(_onErr, "Error serializing 'on' field")
		}

		if popErr := writeBuffer.PopContext("FirmataCommandSetDigitalPinValue"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *FirmataCommandSetDigitalPinValue) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
