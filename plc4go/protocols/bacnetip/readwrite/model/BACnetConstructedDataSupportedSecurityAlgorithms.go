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

// BACnetConstructedDataSupportedSecurityAlgorithms is the data-structure of this message
type BACnetConstructedDataSupportedSecurityAlgorithms struct {
	*BACnetConstructedData
	SupportedSecurityAlgorithms []*BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataSupportedSecurityAlgorithms is the corresponding interface of BACnetConstructedDataSupportedSecurityAlgorithms
type IBACnetConstructedDataSupportedSecurityAlgorithms interface {
	IBACnetConstructedData
	// GetSupportedSecurityAlgorithms returns SupportedSecurityAlgorithms (property field)
	GetSupportedSecurityAlgorithms() []*BACnetApplicationTagUnsignedInteger
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

func (m *BACnetConstructedDataSupportedSecurityAlgorithms) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataSupportedSecurityAlgorithms) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_SUPPORTED_SECURITY_ALGORITHMS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataSupportedSecurityAlgorithms) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataSupportedSecurityAlgorithms) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataSupportedSecurityAlgorithms) GetSupportedSecurityAlgorithms() []*BACnetApplicationTagUnsignedInteger {
	return m.SupportedSecurityAlgorithms
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataSupportedSecurityAlgorithms factory function for BACnetConstructedDataSupportedSecurityAlgorithms
func NewBACnetConstructedDataSupportedSecurityAlgorithms(supportedSecurityAlgorithms []*BACnetApplicationTagUnsignedInteger, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataSupportedSecurityAlgorithms {
	_result := &BACnetConstructedDataSupportedSecurityAlgorithms{
		SupportedSecurityAlgorithms: supportedSecurityAlgorithms,
		BACnetConstructedData:       NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataSupportedSecurityAlgorithms(structType interface{}) *BACnetConstructedDataSupportedSecurityAlgorithms {
	if casted, ok := structType.(BACnetConstructedDataSupportedSecurityAlgorithms); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataSupportedSecurityAlgorithms); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataSupportedSecurityAlgorithms(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataSupportedSecurityAlgorithms(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataSupportedSecurityAlgorithms) GetTypeName() string {
	return "BACnetConstructedDataSupportedSecurityAlgorithms"
}

func (m *BACnetConstructedDataSupportedSecurityAlgorithms) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataSupportedSecurityAlgorithms) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.SupportedSecurityAlgorithms) > 0 {
		for _, element := range m.SupportedSecurityAlgorithms {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *BACnetConstructedDataSupportedSecurityAlgorithms) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataSupportedSecurityAlgorithmsParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataSupportedSecurityAlgorithms, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataSupportedSecurityAlgorithms"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataSupportedSecurityAlgorithms")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (supportedSecurityAlgorithms)
	if pullErr := readBuffer.PullContext("supportedSecurityAlgorithms", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for supportedSecurityAlgorithms")
	}
	// Terminated array
	supportedSecurityAlgorithms := make([]*BACnetApplicationTagUnsignedInteger, 0)
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetApplicationTagParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'supportedSecurityAlgorithms' field")
			}
			supportedSecurityAlgorithms = append(supportedSecurityAlgorithms, CastBACnetApplicationTagUnsignedInteger(_item))

		}
	}
	if closeErr := readBuffer.CloseContext("supportedSecurityAlgorithms", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for supportedSecurityAlgorithms")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataSupportedSecurityAlgorithms"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataSupportedSecurityAlgorithms")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataSupportedSecurityAlgorithms{
		SupportedSecurityAlgorithms: supportedSecurityAlgorithms,
		BACnetConstructedData:       &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataSupportedSecurityAlgorithms) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataSupportedSecurityAlgorithms"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataSupportedSecurityAlgorithms")
		}

		// Array Field (supportedSecurityAlgorithms)
		if m.SupportedSecurityAlgorithms != nil {
			if pushErr := writeBuffer.PushContext("supportedSecurityAlgorithms", utils.WithRenderAsList(true)); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for supportedSecurityAlgorithms")
			}
			for _, _element := range m.SupportedSecurityAlgorithms {
				_elementErr := writeBuffer.WriteSerializable(_element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'supportedSecurityAlgorithms' field")
				}
			}
			if popErr := writeBuffer.PopContext("supportedSecurityAlgorithms", utils.WithRenderAsList(true)); popErr != nil {
				return errors.Wrap(popErr, "Error popping for supportedSecurityAlgorithms")
			}
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataSupportedSecurityAlgorithms"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataSupportedSecurityAlgorithms")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataSupportedSecurityAlgorithms) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}