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

// BACnetConstructedDataIPDHCPLeaseTimeRemaining is the data-structure of this message
type BACnetConstructedDataIPDHCPLeaseTimeRemaining struct {
	*BACnetConstructedData
	IpDhcpLeaseTimeRemaining *BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataIPDHCPLeaseTimeRemaining is the corresponding interface of BACnetConstructedDataIPDHCPLeaseTimeRemaining
type IBACnetConstructedDataIPDHCPLeaseTimeRemaining interface {
	IBACnetConstructedData
	// GetIpDhcpLeaseTimeRemaining returns IpDhcpLeaseTimeRemaining (property field)
	GetIpDhcpLeaseTimeRemaining() *BACnetApplicationTagUnsignedInteger
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

func (m *BACnetConstructedDataIPDHCPLeaseTimeRemaining) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataIPDHCPLeaseTimeRemaining) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_IP_DHCP_LEASE_TIME_REMAINING
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataIPDHCPLeaseTimeRemaining) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataIPDHCPLeaseTimeRemaining) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataIPDHCPLeaseTimeRemaining) GetIpDhcpLeaseTimeRemaining() *BACnetApplicationTagUnsignedInteger {
	return m.IpDhcpLeaseTimeRemaining
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *BACnetConstructedDataIPDHCPLeaseTimeRemaining) GetActualValue() *BACnetApplicationTagUnsignedInteger {
	return CastBACnetApplicationTagUnsignedInteger(m.GetIpDhcpLeaseTimeRemaining())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataIPDHCPLeaseTimeRemaining factory function for BACnetConstructedDataIPDHCPLeaseTimeRemaining
func NewBACnetConstructedDataIPDHCPLeaseTimeRemaining(ipDhcpLeaseTimeRemaining *BACnetApplicationTagUnsignedInteger, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataIPDHCPLeaseTimeRemaining {
	_result := &BACnetConstructedDataIPDHCPLeaseTimeRemaining{
		IpDhcpLeaseTimeRemaining: ipDhcpLeaseTimeRemaining,
		BACnetConstructedData:    NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataIPDHCPLeaseTimeRemaining(structType interface{}) *BACnetConstructedDataIPDHCPLeaseTimeRemaining {
	if casted, ok := structType.(BACnetConstructedDataIPDHCPLeaseTimeRemaining); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataIPDHCPLeaseTimeRemaining); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataIPDHCPLeaseTimeRemaining(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataIPDHCPLeaseTimeRemaining(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataIPDHCPLeaseTimeRemaining) GetTypeName() string {
	return "BACnetConstructedDataIPDHCPLeaseTimeRemaining"
}

func (m *BACnetConstructedDataIPDHCPLeaseTimeRemaining) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataIPDHCPLeaseTimeRemaining) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (ipDhcpLeaseTimeRemaining)
	lengthInBits += m.IpDhcpLeaseTimeRemaining.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetConstructedDataIPDHCPLeaseTimeRemaining) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataIPDHCPLeaseTimeRemainingParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataIPDHCPLeaseTimeRemaining, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataIPDHCPLeaseTimeRemaining"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataIPDHCPLeaseTimeRemaining")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (ipDhcpLeaseTimeRemaining)
	if pullErr := readBuffer.PullContext("ipDhcpLeaseTimeRemaining"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ipDhcpLeaseTimeRemaining")
	}
	_ipDhcpLeaseTimeRemaining, _ipDhcpLeaseTimeRemainingErr := BACnetApplicationTagParse(readBuffer)
	if _ipDhcpLeaseTimeRemainingErr != nil {
		return nil, errors.Wrap(_ipDhcpLeaseTimeRemainingErr, "Error parsing 'ipDhcpLeaseTimeRemaining' field")
	}
	ipDhcpLeaseTimeRemaining := CastBACnetApplicationTagUnsignedInteger(_ipDhcpLeaseTimeRemaining)
	if closeErr := readBuffer.CloseContext("ipDhcpLeaseTimeRemaining"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ipDhcpLeaseTimeRemaining")
	}

	// Virtual field
	_actualValue := ipDhcpLeaseTimeRemaining
	actualValue := CastBACnetApplicationTagUnsignedInteger(_actualValue)
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataIPDHCPLeaseTimeRemaining"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataIPDHCPLeaseTimeRemaining")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataIPDHCPLeaseTimeRemaining{
		IpDhcpLeaseTimeRemaining: CastBACnetApplicationTagUnsignedInteger(ipDhcpLeaseTimeRemaining),
		BACnetConstructedData:    &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataIPDHCPLeaseTimeRemaining) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataIPDHCPLeaseTimeRemaining"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataIPDHCPLeaseTimeRemaining")
		}

		// Simple Field (ipDhcpLeaseTimeRemaining)
		if pushErr := writeBuffer.PushContext("ipDhcpLeaseTimeRemaining"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ipDhcpLeaseTimeRemaining")
		}
		_ipDhcpLeaseTimeRemainingErr := writeBuffer.WriteSerializable(m.IpDhcpLeaseTimeRemaining)
		if popErr := writeBuffer.PopContext("ipDhcpLeaseTimeRemaining"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ipDhcpLeaseTimeRemaining")
		}
		if _ipDhcpLeaseTimeRemainingErr != nil {
			return errors.Wrap(_ipDhcpLeaseTimeRemainingErr, "Error serializing 'ipDhcpLeaseTimeRemaining' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataIPDHCPLeaseTimeRemaining"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataIPDHCPLeaseTimeRemaining")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataIPDHCPLeaseTimeRemaining) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}