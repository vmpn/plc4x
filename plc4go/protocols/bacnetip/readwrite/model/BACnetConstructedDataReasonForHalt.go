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

// BACnetConstructedDataReasonForHalt is the data-structure of this message
type BACnetConstructedDataReasonForHalt struct {
	*BACnetConstructedData
	ProgramError *BACnetProgramErrorTagged

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataReasonForHalt is the corresponding interface of BACnetConstructedDataReasonForHalt
type IBACnetConstructedDataReasonForHalt interface {
	IBACnetConstructedData
	// GetProgramError returns ProgramError (property field)
	GetProgramError() *BACnetProgramErrorTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() *BACnetProgramErrorTagged
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

func (m *BACnetConstructedDataReasonForHalt) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataReasonForHalt) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_REASON_FOR_HALT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataReasonForHalt) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataReasonForHalt) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataReasonForHalt) GetProgramError() *BACnetProgramErrorTagged {
	return m.ProgramError
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *BACnetConstructedDataReasonForHalt) GetActualValue() *BACnetProgramErrorTagged {
	return CastBACnetProgramErrorTagged(m.GetProgramError())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataReasonForHalt factory function for BACnetConstructedDataReasonForHalt
func NewBACnetConstructedDataReasonForHalt(programError *BACnetProgramErrorTagged, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataReasonForHalt {
	_result := &BACnetConstructedDataReasonForHalt{
		ProgramError:          programError,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataReasonForHalt(structType interface{}) *BACnetConstructedDataReasonForHalt {
	if casted, ok := structType.(BACnetConstructedDataReasonForHalt); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataReasonForHalt); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataReasonForHalt(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataReasonForHalt(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataReasonForHalt) GetTypeName() string {
	return "BACnetConstructedDataReasonForHalt"
}

func (m *BACnetConstructedDataReasonForHalt) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataReasonForHalt) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (programError)
	lengthInBits += m.ProgramError.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetConstructedDataReasonForHalt) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataReasonForHaltParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataReasonForHalt, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataReasonForHalt"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataReasonForHalt")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (programError)
	if pullErr := readBuffer.PullContext("programError"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for programError")
	}
	_programError, _programErrorErr := BACnetProgramErrorTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _programErrorErr != nil {
		return nil, errors.Wrap(_programErrorErr, "Error parsing 'programError' field")
	}
	programError := CastBACnetProgramErrorTagged(_programError)
	if closeErr := readBuffer.CloseContext("programError"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for programError")
	}

	// Virtual field
	_actualValue := programError
	actualValue := CastBACnetProgramErrorTagged(_actualValue)
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataReasonForHalt"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataReasonForHalt")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataReasonForHalt{
		ProgramError:          CastBACnetProgramErrorTagged(programError),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataReasonForHalt) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataReasonForHalt"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataReasonForHalt")
		}

		// Simple Field (programError)
		if pushErr := writeBuffer.PushContext("programError"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for programError")
		}
		_programErrorErr := writeBuffer.WriteSerializable(m.ProgramError)
		if popErr := writeBuffer.PopContext("programError"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for programError")
		}
		if _programErrorErr != nil {
			return errors.Wrap(_programErrorErr, "Error serializing 'programError' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataReasonForHalt"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataReasonForHalt")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataReasonForHalt) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}