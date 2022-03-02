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
type ReplyNetwork struct {
	RouteType                 RouteType
	AdditionalBridgeAddresses []*BridgeAddress
	UnitAddress               *UnitAddress
}

// The corresponding interface
type IReplyNetwork interface {
	// GetRouteType returns RouteType
	GetRouteType() RouteType
	// GetAdditionalBridgeAddresses returns AdditionalBridgeAddresses
	GetAdditionalBridgeAddresses() []*BridgeAddress
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
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *ReplyNetwork) GetRouteType() RouteType {
	return m.RouteType
}

func (m *ReplyNetwork) GetAdditionalBridgeAddresses() []*BridgeAddress {
	return m.AdditionalBridgeAddresses
}

func (m *ReplyNetwork) GetUnitAddress() *UnitAddress {
	return m.UnitAddress
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewReplyNetwork factory function for ReplyNetwork
func NewReplyNetwork(routeType RouteType, additionalBridgeAddresses []*BridgeAddress, unitAddress *UnitAddress) *ReplyNetwork {
	return &ReplyNetwork{RouteType: routeType, AdditionalBridgeAddresses: additionalBridgeAddresses, UnitAddress: unitAddress}
}

func CastReplyNetwork(structType interface{}) *ReplyNetwork {
	castFunc := func(typ interface{}) *ReplyNetwork {
		if casted, ok := typ.(ReplyNetwork); ok {
			return &casted
		}
		if casted, ok := typ.(*ReplyNetwork); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ReplyNetwork) GetTypeName() string {
	return "ReplyNetwork"
}

func (m *ReplyNetwork) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ReplyNetwork) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (routeType)
	lengthInBits += 8

	// Array field
	if len(m.AdditionalBridgeAddresses) > 0 {
		for i, element := range m.AdditionalBridgeAddresses {
			last := i == len(m.AdditionalBridgeAddresses)-1
			lengthInBits += element.GetLengthInBitsConditional(last)
		}
	}

	// Simple field (unitAddress)
	lengthInBits += m.UnitAddress.GetLengthInBits()

	return lengthInBits
}

func (m *ReplyNetwork) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ReplyNetworkParse(readBuffer utils.ReadBuffer) (*ReplyNetwork, error) {
	if pullErr := readBuffer.PullContext("ReplyNetwork"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (routeType)
	if pullErr := readBuffer.PullContext("routeType"); pullErr != nil {
		return nil, pullErr
	}
	_routeType, _routeTypeErr := RouteTypeParse(readBuffer)
	if _routeTypeErr != nil {
		return nil, errors.Wrap(_routeTypeErr, "Error parsing 'routeType' field")
	}
	routeType := _routeType
	if closeErr := readBuffer.CloseContext("routeType"); closeErr != nil {
		return nil, closeErr
	}

	// Array field (additionalBridgeAddresses)
	if pullErr := readBuffer.PullContext("additionalBridgeAddresses", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Count array
	additionalBridgeAddresses := make([]*BridgeAddress, routeType.AdditionalBridges())
	{
		for curItem := uint16(0); curItem < uint16(routeType.AdditionalBridges()); curItem++ {
			_item, _err := BridgeAddressParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'additionalBridgeAddresses' field")
			}
			additionalBridgeAddresses[curItem] = CastBridgeAddress(_item)
		}
	}
	if closeErr := readBuffer.CloseContext("additionalBridgeAddresses", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

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

	if closeErr := readBuffer.CloseContext("ReplyNetwork"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewReplyNetwork(routeType, additionalBridgeAddresses, unitAddress), nil
}

func (m *ReplyNetwork) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("ReplyNetwork"); pushErr != nil {
		return pushErr
	}

	// Simple Field (routeType)
	if pushErr := writeBuffer.PushContext("routeType"); pushErr != nil {
		return pushErr
	}
	_routeTypeErr := m.RouteType.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("routeType"); popErr != nil {
		return popErr
	}
	if _routeTypeErr != nil {
		return errors.Wrap(_routeTypeErr, "Error serializing 'routeType' field")
	}

	// Array Field (additionalBridgeAddresses)
	if m.AdditionalBridgeAddresses != nil {
		if pushErr := writeBuffer.PushContext("additionalBridgeAddresses", utils.WithRenderAsList(true)); pushErr != nil {
			return pushErr
		}
		for _, _element := range m.AdditionalBridgeAddresses {
			_elementErr := _element.Serialize(writeBuffer)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'additionalBridgeAddresses' field")
			}
		}
		if popErr := writeBuffer.PopContext("additionalBridgeAddresses", utils.WithRenderAsList(true)); popErr != nil {
			return popErr
		}
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

	if popErr := writeBuffer.PopContext("ReplyNetwork"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *ReplyNetwork) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
