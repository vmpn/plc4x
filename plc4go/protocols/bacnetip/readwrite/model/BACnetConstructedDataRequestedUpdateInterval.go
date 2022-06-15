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

// BACnetConstructedDataRequestedUpdateInterval is the data-structure of this message
type BACnetConstructedDataRequestedUpdateInterval struct {
	*BACnetConstructedData
	RequestedUpdateInterval *BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataRequestedUpdateInterval is the corresponding interface of BACnetConstructedDataRequestedUpdateInterval
type IBACnetConstructedDataRequestedUpdateInterval interface {
	IBACnetConstructedData
	// GetRequestedUpdateInterval returns RequestedUpdateInterval (property field)
	GetRequestedUpdateInterval() *BACnetApplicationTagUnsignedInteger
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

func (m *BACnetConstructedDataRequestedUpdateInterval) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataRequestedUpdateInterval) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_REQUESTED_UPDATE_INTERVAL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataRequestedUpdateInterval) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataRequestedUpdateInterval) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataRequestedUpdateInterval) GetRequestedUpdateInterval() *BACnetApplicationTagUnsignedInteger {
	return m.RequestedUpdateInterval
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *BACnetConstructedDataRequestedUpdateInterval) GetActualValue() *BACnetApplicationTagUnsignedInteger {
	return CastBACnetApplicationTagUnsignedInteger(m.GetRequestedUpdateInterval())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataRequestedUpdateInterval factory function for BACnetConstructedDataRequestedUpdateInterval
func NewBACnetConstructedDataRequestedUpdateInterval(requestedUpdateInterval *BACnetApplicationTagUnsignedInteger, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataRequestedUpdateInterval {
	_result := &BACnetConstructedDataRequestedUpdateInterval{
		RequestedUpdateInterval: requestedUpdateInterval,
		BACnetConstructedData:   NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataRequestedUpdateInterval(structType interface{}) *BACnetConstructedDataRequestedUpdateInterval {
	if casted, ok := structType.(BACnetConstructedDataRequestedUpdateInterval); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataRequestedUpdateInterval); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataRequestedUpdateInterval(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataRequestedUpdateInterval(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataRequestedUpdateInterval) GetTypeName() string {
	return "BACnetConstructedDataRequestedUpdateInterval"
}

func (m *BACnetConstructedDataRequestedUpdateInterval) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataRequestedUpdateInterval) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (requestedUpdateInterval)
	lengthInBits += m.RequestedUpdateInterval.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetConstructedDataRequestedUpdateInterval) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataRequestedUpdateIntervalParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataRequestedUpdateInterval, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataRequestedUpdateInterval"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataRequestedUpdateInterval")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (requestedUpdateInterval)
	if pullErr := readBuffer.PullContext("requestedUpdateInterval"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for requestedUpdateInterval")
	}
	_requestedUpdateInterval, _requestedUpdateIntervalErr := BACnetApplicationTagParse(readBuffer)
	if _requestedUpdateIntervalErr != nil {
		return nil, errors.Wrap(_requestedUpdateIntervalErr, "Error parsing 'requestedUpdateInterval' field")
	}
	requestedUpdateInterval := CastBACnetApplicationTagUnsignedInteger(_requestedUpdateInterval)
	if closeErr := readBuffer.CloseContext("requestedUpdateInterval"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for requestedUpdateInterval")
	}

	// Virtual field
	_actualValue := requestedUpdateInterval
	actualValue := CastBACnetApplicationTagUnsignedInteger(_actualValue)
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataRequestedUpdateInterval"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataRequestedUpdateInterval")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataRequestedUpdateInterval{
		RequestedUpdateInterval: CastBACnetApplicationTagUnsignedInteger(requestedUpdateInterval),
		BACnetConstructedData:   &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataRequestedUpdateInterval) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataRequestedUpdateInterval"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataRequestedUpdateInterval")
		}

		// Simple Field (requestedUpdateInterval)
		if pushErr := writeBuffer.PushContext("requestedUpdateInterval"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for requestedUpdateInterval")
		}
		_requestedUpdateIntervalErr := writeBuffer.WriteSerializable(m.RequestedUpdateInterval)
		if popErr := writeBuffer.PopContext("requestedUpdateInterval"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for requestedUpdateInterval")
		}
		if _requestedUpdateIntervalErr != nil {
			return errors.Wrap(_requestedUpdateIntervalErr, "Error serializing 'requestedUpdateInterval' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataRequestedUpdateInterval"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataRequestedUpdateInterval")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataRequestedUpdateInterval) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}