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
type S7ParameterWriteVarRequest struct {
	Items  []*S7VarRequestParameterItem
	Parent *S7Parameter
}

// The corresponding interface
type IS7ParameterWriteVarRequest interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *S7ParameterWriteVarRequest) ParameterType() uint8 {
	return 0x05
}

func (m *S7ParameterWriteVarRequest) MessageType() uint8 {
	return 0x01
}

func (m *S7ParameterWriteVarRequest) InitializeParent(parent *S7Parameter) {
}

func NewS7ParameterWriteVarRequest(items []*S7VarRequestParameterItem) *S7Parameter {
	child := &S7ParameterWriteVarRequest{
		Items:  items,
		Parent: NewS7Parameter(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastS7ParameterWriteVarRequest(structType interface{}) *S7ParameterWriteVarRequest {
	castFunc := func(typ interface{}) *S7ParameterWriteVarRequest {
		if casted, ok := typ.(S7ParameterWriteVarRequest); ok {
			return &casted
		}
		if casted, ok := typ.(*S7ParameterWriteVarRequest); ok {
			return casted
		}
		if casted, ok := typ.(S7Parameter); ok {
			return CastS7ParameterWriteVarRequest(casted.Child)
		}
		if casted, ok := typ.(*S7Parameter); ok {
			return CastS7ParameterWriteVarRequest(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *S7ParameterWriteVarRequest) GetTypeName() string {
	return "S7ParameterWriteVarRequest"
}

func (m *S7ParameterWriteVarRequest) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *S7ParameterWriteVarRequest) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Implicit Field (numItems)
	lengthInBits += 8

	// Array field
	if len(m.Items) > 0 {
		for i, element := range m.Items {
			last := i == len(m.Items)-1
			lengthInBits += element.LengthInBitsConditional(last)
		}
	}

	return lengthInBits
}

func (m *S7ParameterWriteVarRequest) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func S7ParameterWriteVarRequestParse(readBuffer utils.ReadBuffer, messageType uint8) (*S7Parameter, error) {
	if pullErr := readBuffer.PullContext("S7ParameterWriteVarRequest"); pullErr != nil {
		return nil, pullErr
	}

	// Implicit Field (numItems) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	numItems, _numItemsErr := readBuffer.ReadUint8("numItems", 8)
	_ = numItems
	if _numItemsErr != nil {
		return nil, errors.Wrap(_numItemsErr, "Error parsing 'numItems' field")
	}

	// Array field (items)
	if pullErr := readBuffer.PullContext("items", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Count array
	items := make([]*S7VarRequestParameterItem, numItems)
	for curItem := uint16(0); curItem < uint16(numItems); curItem++ {
		_item, _err := S7VarRequestParameterItemParse(readBuffer)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'items' field")
		}
		items[curItem] = _item
	}
	if closeErr := readBuffer.CloseContext("items", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("S7ParameterWriteVarRequest"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &S7ParameterWriteVarRequest{
		Items:  items,
		Parent: &S7Parameter{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *S7ParameterWriteVarRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7ParameterWriteVarRequest"); pushErr != nil {
			return pushErr
		}

		// Implicit Field (numItems) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		numItems := uint8(uint8(len(m.Items)))
		_numItemsErr := writeBuffer.WriteUint8("numItems", 8, (numItems))
		if _numItemsErr != nil {
			return errors.Wrap(_numItemsErr, "Error serializing 'numItems' field")
		}

		// Array Field (items)
		if m.Items != nil {
			if pushErr := writeBuffer.PushContext("items", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, _element := range m.Items {
				_elementErr := _element.Serialize(writeBuffer)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'items' field")
				}
			}
			if popErr := writeBuffer.PopContext("items", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		if popErr := writeBuffer.PopContext("S7ParameterWriteVarRequest"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *S7ParameterWriteVarRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
