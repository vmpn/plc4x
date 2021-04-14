//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//
package model

import (
	"encoding/xml"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
	"io"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type ModbusPDUWriteFileRecordResponse struct {
	Items  []*ModbusPDUWriteFileRecordResponseItem
	Parent *ModbusPDU
}

// The corresponding interface
type IModbusPDUWriteFileRecordResponse interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ModbusPDUWriteFileRecordResponse) ErrorFlag() bool {
	return false
}

func (m *ModbusPDUWriteFileRecordResponse) FunctionFlag() uint8 {
	return 0x15
}

func (m *ModbusPDUWriteFileRecordResponse) Response() bool {
	return true
}

func (m *ModbusPDUWriteFileRecordResponse) InitializeParent(parent *ModbusPDU) {
}

func NewModbusPDUWriteFileRecordResponse(items []*ModbusPDUWriteFileRecordResponseItem) *ModbusPDU {
	child := &ModbusPDUWriteFileRecordResponse{
		Items:  items,
		Parent: NewModbusPDU(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastModbusPDUWriteFileRecordResponse(structType interface{}) *ModbusPDUWriteFileRecordResponse {
	castFunc := func(typ interface{}) *ModbusPDUWriteFileRecordResponse {
		if casted, ok := typ.(ModbusPDUWriteFileRecordResponse); ok {
			return &casted
		}
		if casted, ok := typ.(*ModbusPDUWriteFileRecordResponse); ok {
			return casted
		}
		if casted, ok := typ.(ModbusPDU); ok {
			return CastModbusPDUWriteFileRecordResponse(casted.Child)
		}
		if casted, ok := typ.(*ModbusPDU); ok {
			return CastModbusPDUWriteFileRecordResponse(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ModbusPDUWriteFileRecordResponse) GetTypeName() string {
	return "ModbusPDUWriteFileRecordResponse"
}

func (m *ModbusPDUWriteFileRecordResponse) LengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Implicit Field (byteCount)
	lengthInBits += 8

	// Array field
	if len(m.Items) > 0 {
		for _, element := range m.Items {
			lengthInBits += element.LengthInBits()
		}
	}

	return lengthInBits
}

func (m *ModbusPDUWriteFileRecordResponse) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ModbusPDUWriteFileRecordResponseParse(io *utils.ReadBuffer) (*ModbusPDU, error) {

	// Implicit Field (byteCount) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	byteCount, _byteCountErr := io.ReadUint8(8)
	_ = byteCount
	if _byteCountErr != nil {
		return nil, errors.Wrap(_byteCountErr, "Error parsing 'byteCount' field")
	}

	// Array field (items)
	// Length array
	items := make([]*ModbusPDUWriteFileRecordResponseItem, 0)
	_itemsLength := byteCount
	_itemsEndPos := io.GetPos() + uint16(_itemsLength)
	for io.GetPos() < _itemsEndPos {
		_item, _err := ModbusPDUWriteFileRecordResponseItemParse(io)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'items' field")
		}
		items = append(items, _item)
	}

	// Create a partially initialized instance
	_child := &ModbusPDUWriteFileRecordResponse{
		Items:  items,
		Parent: &ModbusPDU{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *ModbusPDUWriteFileRecordResponse) Serialize(io utils.WriteBuffer) error {
	itemsArraySizeInBytes := func(items []*ModbusPDUWriteFileRecordResponseItem) uint32 {
		var sizeInBytes uint32 = 0
		for _, v := range items {
			sizeInBytes += uint32(v.LengthInBytes())
		}
		return sizeInBytes
	}
	ser := func() error {

		// Implicit Field (byteCount) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		byteCount := uint8(uint8(itemsArraySizeInBytes(m.Items)))
		_byteCountErr := io.WriteUint8(8, (byteCount))
		if _byteCountErr != nil {
			return errors.Wrap(_byteCountErr, "Error serializing 'byteCount' field")
		}

		// Array Field (items)
		if m.Items != nil {
			for _, _element := range m.Items {
				_elementErr := _element.Serialize(io)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'items' field")
				}
			}
		}

		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *ModbusPDUWriteFileRecordResponse) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	foundContent := false
	token = start
	for {
		switch token.(type) {
		case xml.StartElement:
			foundContent = true
			tok := token.(xml.StartElement)
			switch tok.Name.Local {
			case "items":
				var data []*ModbusPDUWriteFileRecordResponseItem
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.Items = data
			}
		}
		token, err = d.Token()
		if err != nil {
			if err == io.EOF && foundContent {
				return nil
			}
			return err
		}
	}
}

func (m *ModbusPDUWriteFileRecordResponse) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeToken(xml.StartElement{Name: xml.Name{Local: "items"}}); err != nil {
		return err
	}
	for _, arrayElement := range m.Items {
		if err := e.EncodeElement(arrayElement, xml.StartElement{Name: xml.Name{Local: "items"}}); err != nil {
			return err
		}
	}
	if err := e.EncodeToken(xml.EndElement{Name: xml.Name{Local: "items"}}); err != nil {
		return err
	}
	return nil
}

func (m ModbusPDUWriteFileRecordResponse) String() string {
	return string(m.Box("ModbusPDUWriteFileRecordResponse", utils.DefaultWidth*2))
}

func (m ModbusPDUWriteFileRecordResponse) Box(name string, width int) utils.AsciiBox {
	if name == "" {
		name = "ModbusPDUWriteFileRecordResponse"
	}
	boxes := make([]utils.AsciiBox, 0)
	boxes = append(boxes, utils.BoxAnything("Items", m.Items, width-2))
	return utils.BoxBox(name, utils.AlignBoxes(boxes, width-2), 0)
}
