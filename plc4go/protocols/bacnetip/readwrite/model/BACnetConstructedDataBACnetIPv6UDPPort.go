/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
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

// BACnetConstructedDataBACnetIPv6UDPPort is the data-structure of this message
type BACnetConstructedDataBACnetIPv6UDPPort struct {
	*BACnetConstructedData
	Ipv6UdpPort *BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataBACnetIPv6UDPPort is the corresponding interface of BACnetConstructedDataBACnetIPv6UDPPort
type IBACnetConstructedDataBACnetIPv6UDPPort interface {
	IBACnetConstructedData
	// GetIpv6UdpPort returns Ipv6UdpPort (property field)
	GetIpv6UdpPort() *BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() *BACnetApplicationTagUnsignedInteger
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

func (m *BACnetConstructedDataBACnetIPv6UDPPort) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataBACnetIPv6UDPPort) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_BACNET_IPV6_UDP_PORT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataBACnetIPv6UDPPort) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataBACnetIPv6UDPPort) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataBACnetIPv6UDPPort) GetIpv6UdpPort() *BACnetApplicationTagUnsignedInteger {
	return m.Ipv6UdpPort
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *BACnetConstructedDataBACnetIPv6UDPPort) GetActualValue() *BACnetApplicationTagUnsignedInteger {
	return CastBACnetApplicationTagUnsignedInteger(m.GetIpv6UdpPort())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataBACnetIPv6UDPPort factory function for BACnetConstructedDataBACnetIPv6UDPPort
func NewBACnetConstructedDataBACnetIPv6UDPPort(ipv6UdpPort *BACnetApplicationTagUnsignedInteger, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataBACnetIPv6UDPPort {
	_result := &BACnetConstructedDataBACnetIPv6UDPPort{
		Ipv6UdpPort:           ipv6UdpPort,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataBACnetIPv6UDPPort(structType interface{}) *BACnetConstructedDataBACnetIPv6UDPPort {
	if casted, ok := structType.(BACnetConstructedDataBACnetIPv6UDPPort); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataBACnetIPv6UDPPort); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataBACnetIPv6UDPPort(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataBACnetIPv6UDPPort(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataBACnetIPv6UDPPort) GetTypeName() string {
	return "BACnetConstructedDataBACnetIPv6UDPPort"
}

func (m *BACnetConstructedDataBACnetIPv6UDPPort) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataBACnetIPv6UDPPort) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (ipv6UdpPort)
	lengthInBits += m.Ipv6UdpPort.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetConstructedDataBACnetIPv6UDPPort) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataBACnetIPv6UDPPortParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataBACnetIPv6UDPPort, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataBACnetIPv6UDPPort"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataBACnetIPv6UDPPort")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (ipv6UdpPort)
	if pullErr := readBuffer.PullContext("ipv6UdpPort"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ipv6UdpPort")
	}
	_ipv6UdpPort, _ipv6UdpPortErr := BACnetApplicationTagParse(readBuffer)
	if _ipv6UdpPortErr != nil {
		return nil, errors.Wrap(_ipv6UdpPortErr, "Error parsing 'ipv6UdpPort' field")
	}
	ipv6UdpPort := CastBACnetApplicationTagUnsignedInteger(_ipv6UdpPort)
	if closeErr := readBuffer.CloseContext("ipv6UdpPort"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ipv6UdpPort")
	}

	// Virtual field
	_actualValue := ipv6UdpPort
	actualValue := CastBACnetApplicationTagUnsignedInteger(_actualValue)
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataBACnetIPv6UDPPort"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataBACnetIPv6UDPPort")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataBACnetIPv6UDPPort{
		Ipv6UdpPort:           CastBACnetApplicationTagUnsignedInteger(ipv6UdpPort),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataBACnetIPv6UDPPort) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataBACnetIPv6UDPPort"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataBACnetIPv6UDPPort")
		}

		// Simple Field (ipv6UdpPort)
		if pushErr := writeBuffer.PushContext("ipv6UdpPort"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ipv6UdpPort")
		}
		_ipv6UdpPortErr := writeBuffer.WriteSerializable(m.Ipv6UdpPort)
		if popErr := writeBuffer.PopContext("ipv6UdpPort"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ipv6UdpPort")
		}
		if _ipv6UdpPortErr != nil {
			return errors.Wrap(_ipv6UdpPortErr, "Error serializing 'ipv6UdpPort' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataBACnetIPv6UDPPort"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataBACnetIPv6UDPPort")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataBACnetIPv6UDPPort) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}