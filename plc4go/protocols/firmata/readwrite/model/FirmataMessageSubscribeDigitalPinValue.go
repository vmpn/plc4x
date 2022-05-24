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
	"github.com/rs/zerolog/log"
)

// Code generated by code-generation. DO NOT EDIT.

// FirmataMessageSubscribeDigitalPinValue is the data-structure of this message
type FirmataMessageSubscribeDigitalPinValue struct {
	*FirmataMessage
	Pin    uint8
	Enable bool

	// Arguments.
	Response bool
}

// IFirmataMessageSubscribeDigitalPinValue is the corresponding interface of FirmataMessageSubscribeDigitalPinValue
type IFirmataMessageSubscribeDigitalPinValue interface {
	IFirmataMessage
	// GetPin returns Pin (property field)
	GetPin() uint8
	// GetEnable returns Enable (property field)
	GetEnable() bool
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

func (m *FirmataMessageSubscribeDigitalPinValue) GetMessageType() uint8 {
	return 0xD
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *FirmataMessageSubscribeDigitalPinValue) InitializeParent(parent *FirmataMessage) {}

func (m *FirmataMessageSubscribeDigitalPinValue) GetParent() *FirmataMessage {
	return m.FirmataMessage
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *FirmataMessageSubscribeDigitalPinValue) GetPin() uint8 {
	return m.Pin
}

func (m *FirmataMessageSubscribeDigitalPinValue) GetEnable() bool {
	return m.Enable
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewFirmataMessageSubscribeDigitalPinValue factory function for FirmataMessageSubscribeDigitalPinValue
func NewFirmataMessageSubscribeDigitalPinValue(pin uint8, enable bool, response bool) *FirmataMessageSubscribeDigitalPinValue {
	_result := &FirmataMessageSubscribeDigitalPinValue{
		Pin:            pin,
		Enable:         enable,
		FirmataMessage: NewFirmataMessage(response),
	}
	_result.Child = _result
	return _result
}

func CastFirmataMessageSubscribeDigitalPinValue(structType interface{}) *FirmataMessageSubscribeDigitalPinValue {
	if casted, ok := structType.(FirmataMessageSubscribeDigitalPinValue); ok {
		return &casted
	}
	if casted, ok := structType.(*FirmataMessageSubscribeDigitalPinValue); ok {
		return casted
	}
	if casted, ok := structType.(FirmataMessage); ok {
		return CastFirmataMessageSubscribeDigitalPinValue(casted.Child)
	}
	if casted, ok := structType.(*FirmataMessage); ok {
		return CastFirmataMessageSubscribeDigitalPinValue(casted.Child)
	}
	return nil
}

func (m *FirmataMessageSubscribeDigitalPinValue) GetTypeName() string {
	return "FirmataMessageSubscribeDigitalPinValue"
}

func (m *FirmataMessageSubscribeDigitalPinValue) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *FirmataMessageSubscribeDigitalPinValue) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (pin)
	lengthInBits += 4

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (enable)
	lengthInBits += 1

	return lengthInBits
}

func (m *FirmataMessageSubscribeDigitalPinValue) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func FirmataMessageSubscribeDigitalPinValueParse(readBuffer utils.ReadBuffer, response bool) (*FirmataMessageSubscribeDigitalPinValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("FirmataMessageSubscribeDigitalPinValue"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (pin)
	_pin, _pinErr := readBuffer.ReadUint8("pin", 4)
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

	// Simple Field (enable)
	_enable, _enableErr := readBuffer.ReadBit("enable")
	if _enableErr != nil {
		return nil, errors.Wrap(_enableErr, "Error parsing 'enable' field")
	}
	enable := _enable

	if closeErr := readBuffer.CloseContext("FirmataMessageSubscribeDigitalPinValue"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &FirmataMessageSubscribeDigitalPinValue{
		Pin:            pin,
		Enable:         enable,
		FirmataMessage: &FirmataMessage{},
	}
	_child.FirmataMessage.Child = _child
	return _child, nil
}

func (m *FirmataMessageSubscribeDigitalPinValue) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("FirmataMessageSubscribeDigitalPinValue"); pushErr != nil {
			return pushErr
		}

		// Simple Field (pin)
		pin := uint8(m.Pin)
		_pinErr := writeBuffer.WriteUint8("pin", 4, (pin))
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

		// Simple Field (enable)
		enable := bool(m.Enable)
		_enableErr := writeBuffer.WriteBit("enable", (enable))
		if _enableErr != nil {
			return errors.Wrap(_enableErr, "Error serializing 'enable' field")
		}

		if popErr := writeBuffer.PopContext("FirmataMessageSubscribeDigitalPinValue"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *FirmataMessageSubscribeDigitalPinValue) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}