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

// BACnetConstructedDataBias is the data-structure of this message
type BACnetConstructedDataBias struct {
	*BACnetConstructedData
	Bias *BACnetApplicationTagReal

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataBias is the corresponding interface of BACnetConstructedDataBias
type IBACnetConstructedDataBias interface {
	IBACnetConstructedData
	// GetBias returns Bias (property field)
	GetBias() *BACnetApplicationTagReal
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() *BACnetApplicationTagReal
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

func (m *BACnetConstructedDataBias) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataBias) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_BIAS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataBias) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataBias) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataBias) GetBias() *BACnetApplicationTagReal {
	return m.Bias
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *BACnetConstructedDataBias) GetActualValue() *BACnetApplicationTagReal {
	return CastBACnetApplicationTagReal(m.GetBias())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataBias factory function for BACnetConstructedDataBias
func NewBACnetConstructedDataBias(bias *BACnetApplicationTagReal, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataBias {
	_result := &BACnetConstructedDataBias{
		Bias:                  bias,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataBias(structType interface{}) *BACnetConstructedDataBias {
	if casted, ok := structType.(BACnetConstructedDataBias); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataBias); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataBias(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataBias(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataBias) GetTypeName() string {
	return "BACnetConstructedDataBias"
}

func (m *BACnetConstructedDataBias) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataBias) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (bias)
	lengthInBits += m.Bias.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetConstructedDataBias) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataBiasParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataBias, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataBias"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataBias")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (bias)
	if pullErr := readBuffer.PullContext("bias"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for bias")
	}
	_bias, _biasErr := BACnetApplicationTagParse(readBuffer)
	if _biasErr != nil {
		return nil, errors.Wrap(_biasErr, "Error parsing 'bias' field")
	}
	bias := CastBACnetApplicationTagReal(_bias)
	if closeErr := readBuffer.CloseContext("bias"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for bias")
	}

	// Virtual field
	_actualValue := bias
	actualValue := CastBACnetApplicationTagReal(_actualValue)
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataBias"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataBias")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataBias{
		Bias:                  CastBACnetApplicationTagReal(bias),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataBias) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataBias"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataBias")
		}

		// Simple Field (bias)
		if pushErr := writeBuffer.PushContext("bias"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for bias")
		}
		_biasErr := writeBuffer.WriteSerializable(m.Bias)
		if popErr := writeBuffer.PopContext("bias"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for bias")
		}
		if _biasErr != nil {
			return errors.Wrap(_biasErr, "Error serializing 'bias' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataBias"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataBias")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataBias) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}