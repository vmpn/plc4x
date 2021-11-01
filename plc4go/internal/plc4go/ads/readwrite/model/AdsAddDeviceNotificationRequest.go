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
	"math/big"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type AdsAddDeviceNotificationRequest struct {
	IndexGroup       uint32
	IndexOffset      uint32
	Length           uint32
	TransmissionMode uint32
	MaxDelay         uint32
	CycleTime        uint32
	Parent           *AdsData
}

// The corresponding interface
type IAdsAddDeviceNotificationRequest interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *AdsAddDeviceNotificationRequest) CommandId() CommandId {
	return CommandId_ADS_ADD_DEVICE_NOTIFICATION
}

func (m *AdsAddDeviceNotificationRequest) Response() bool {
	return false
}

func (m *AdsAddDeviceNotificationRequest) InitializeParent(parent *AdsData) {
}

func NewAdsAddDeviceNotificationRequest(indexGroup uint32, indexOffset uint32, length uint32, transmissionMode uint32, maxDelay uint32, cycleTime uint32) *AdsData {
	child := &AdsAddDeviceNotificationRequest{
		IndexGroup:       indexGroup,
		IndexOffset:      indexOffset,
		Length:           length,
		TransmissionMode: transmissionMode,
		MaxDelay:         maxDelay,
		CycleTime:        cycleTime,
		Parent:           NewAdsData(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastAdsAddDeviceNotificationRequest(structType interface{}) *AdsAddDeviceNotificationRequest {
	castFunc := func(typ interface{}) *AdsAddDeviceNotificationRequest {
		if casted, ok := typ.(AdsAddDeviceNotificationRequest); ok {
			return &casted
		}
		if casted, ok := typ.(*AdsAddDeviceNotificationRequest); ok {
			return casted
		}
		if casted, ok := typ.(AdsData); ok {
			return CastAdsAddDeviceNotificationRequest(casted.Child)
		}
		if casted, ok := typ.(*AdsData); ok {
			return CastAdsAddDeviceNotificationRequest(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *AdsAddDeviceNotificationRequest) GetTypeName() string {
	return "AdsAddDeviceNotificationRequest"
}

func (m *AdsAddDeviceNotificationRequest) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *AdsAddDeviceNotificationRequest) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (indexGroup)
	lengthInBits += 32

	// Simple field (indexOffset)
	lengthInBits += 32

	// Simple field (length)
	lengthInBits += 32

	// Simple field (transmissionMode)
	lengthInBits += 32

	// Simple field (maxDelay)
	lengthInBits += 32

	// Simple field (cycleTime)
	lengthInBits += 32

	// Reserved Field (reserved)
	lengthInBits += 128

	return lengthInBits
}

func (m *AdsAddDeviceNotificationRequest) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func AdsAddDeviceNotificationRequestParse(readBuffer utils.ReadBuffer, commandId CommandId, response bool) (*AdsData, error) {
	if pullErr := readBuffer.PullContext("AdsAddDeviceNotificationRequest"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (indexGroup)
	indexGroup, _indexGroupErr := readBuffer.ReadUint32("indexGroup", 32)
	if _indexGroupErr != nil {
		return nil, errors.Wrap(_indexGroupErr, "Error parsing 'indexGroup' field")
	}

	// Simple Field (indexOffset)
	indexOffset, _indexOffsetErr := readBuffer.ReadUint32("indexOffset", 32)
	if _indexOffsetErr != nil {
		return nil, errors.Wrap(_indexOffsetErr, "Error parsing 'indexOffset' field")
	}

	// Simple Field (length)
	length, _lengthErr := readBuffer.ReadUint32("length", 32)
	if _lengthErr != nil {
		return nil, errors.Wrap(_lengthErr, "Error parsing 'length' field")
	}

	// Simple Field (transmissionMode)
	transmissionMode, _transmissionModeErr := readBuffer.ReadUint32("transmissionMode", 32)
	if _transmissionModeErr != nil {
		return nil, errors.Wrap(_transmissionModeErr, "Error parsing 'transmissionMode' field")
	}

	// Simple Field (maxDelay)
	maxDelay, _maxDelayErr := readBuffer.ReadUint32("maxDelay", 32)
	if _maxDelayErr != nil {
		return nil, errors.Wrap(_maxDelayErr, "Error parsing 'maxDelay' field")
	}

	// Simple Field (cycleTime)
	cycleTime, _cycleTimeErr := readBuffer.ReadUint32("cycleTime", 32)
	if _cycleTimeErr != nil {
		return nil, errors.Wrap(_cycleTimeErr, "Error parsing 'cycleTime' field")
	}

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadBigInt("reserved", 128)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved.Cmp(big.NewInt(0x0000)) != 0 {
			log.Info().Fields(map[string]interface{}{
				"expected value": big.NewInt(0x0000),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	if closeErr := readBuffer.CloseContext("AdsAddDeviceNotificationRequest"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &AdsAddDeviceNotificationRequest{
		IndexGroup:       indexGroup,
		IndexOffset:      indexOffset,
		Length:           length,
		TransmissionMode: transmissionMode,
		MaxDelay:         maxDelay,
		CycleTime:        cycleTime,
		Parent:           &AdsData{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *AdsAddDeviceNotificationRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsAddDeviceNotificationRequest"); pushErr != nil {
			return pushErr
		}

		// Simple Field (indexGroup)
		indexGroup := uint32(m.IndexGroup)
		_indexGroupErr := writeBuffer.WriteUint32("indexGroup", 32, (indexGroup))
		if _indexGroupErr != nil {
			return errors.Wrap(_indexGroupErr, "Error serializing 'indexGroup' field")
		}

		// Simple Field (indexOffset)
		indexOffset := uint32(m.IndexOffset)
		_indexOffsetErr := writeBuffer.WriteUint32("indexOffset", 32, (indexOffset))
		if _indexOffsetErr != nil {
			return errors.Wrap(_indexOffsetErr, "Error serializing 'indexOffset' field")
		}

		// Simple Field (length)
		length := uint32(m.Length)
		_lengthErr := writeBuffer.WriteUint32("length", 32, (length))
		if _lengthErr != nil {
			return errors.Wrap(_lengthErr, "Error serializing 'length' field")
		}

		// Simple Field (transmissionMode)
		transmissionMode := uint32(m.TransmissionMode)
		_transmissionModeErr := writeBuffer.WriteUint32("transmissionMode", 32, (transmissionMode))
		if _transmissionModeErr != nil {
			return errors.Wrap(_transmissionModeErr, "Error serializing 'transmissionMode' field")
		}

		// Simple Field (maxDelay)
		maxDelay := uint32(m.MaxDelay)
		_maxDelayErr := writeBuffer.WriteUint32("maxDelay", 32, (maxDelay))
		if _maxDelayErr != nil {
			return errors.Wrap(_maxDelayErr, "Error serializing 'maxDelay' field")
		}

		// Simple Field (cycleTime)
		cycleTime := uint32(m.CycleTime)
		_cycleTimeErr := writeBuffer.WriteUint32("cycleTime", 32, (cycleTime))
		if _cycleTimeErr != nil {
			return errors.Wrap(_cycleTimeErr, "Error serializing 'cycleTime' field")
		}

		// Reserved Field (reserved)
		{
			_err := writeBuffer.WriteBigInt("reserved", 128, big.NewInt(0x0000))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		if popErr := writeBuffer.PopContext("AdsAddDeviceNotificationRequest"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *AdsAddDeviceNotificationRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
