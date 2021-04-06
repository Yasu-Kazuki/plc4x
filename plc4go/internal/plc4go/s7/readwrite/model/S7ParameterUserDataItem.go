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
	"reflect"
	"strings"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type S7ParameterUserDataItem struct {
	Child IS7ParameterUserDataItemChild
}

// The corresponding interface
type IS7ParameterUserDataItem interface {
	ItemType() uint8
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

type IS7ParameterUserDataItemParent interface {
	SerializeParent(io utils.WriteBuffer, child IS7ParameterUserDataItem, serializeChildFunction func() error) error
	GetTypeName() string
}

type IS7ParameterUserDataItemChild interface {
	Serialize(io utils.WriteBuffer) error
	InitializeParent(parent *S7ParameterUserDataItem)
	GetTypeName() string
	IS7ParameterUserDataItem
}

func NewS7ParameterUserDataItem() *S7ParameterUserDataItem {
	return &S7ParameterUserDataItem{}
}

func CastS7ParameterUserDataItem(structType interface{}) *S7ParameterUserDataItem {
	castFunc := func(typ interface{}) *S7ParameterUserDataItem {
		if casted, ok := typ.(S7ParameterUserDataItem); ok {
			return &casted
		}
		if casted, ok := typ.(*S7ParameterUserDataItem); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *S7ParameterUserDataItem) GetTypeName() string {
	return "S7ParameterUserDataItem"
}

func (m *S7ParameterUserDataItem) LengthInBits() uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (itemType)
	lengthInBits += 8

	// Length of sub-type elements will be added by sub-type...
	lengthInBits += m.Child.LengthInBits()

	return lengthInBits
}

func (m *S7ParameterUserDataItem) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func S7ParameterUserDataItemParse(io *utils.ReadBuffer) (*S7ParameterUserDataItem, error) {

	// Discriminator Field (itemType) (Used as input to a switch field)
	itemType, _itemTypeErr := io.ReadUint8(8)
	if _itemTypeErr != nil {
		return nil, errors.Wrap(_itemTypeErr, "Error parsing 'itemType' field")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _parent *S7ParameterUserDataItem
	var typeSwitchError error
	switch {
	case itemType == 0x12: // S7ParameterUserDataItemCPUFunctions
		_parent, typeSwitchError = S7ParameterUserDataItemCPUFunctionsParse(io)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	// Finish initializing
	_parent.Child.InitializeParent(_parent)
	return _parent, nil
}

func (m *S7ParameterUserDataItem) Serialize(io utils.WriteBuffer) error {
	return m.Child.Serialize(io)
}

func (m *S7ParameterUserDataItem) SerializeParent(io utils.WriteBuffer, child IS7ParameterUserDataItem, serializeChildFunction func() error) error {

	// Discriminator Field (itemType) (Used as input to a switch field)
	itemType := uint8(child.ItemType())
	_itemTypeErr := io.WriteUint8(8, (itemType))

	if _itemTypeErr != nil {
		return errors.Wrap(_itemTypeErr, "Error serializing 'itemType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	_typeSwitchErr := serializeChildFunction()
	if _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	return nil
}

func (m *S7ParameterUserDataItem) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	for {
		token, err = d.Token()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		switch token.(type) {
		case xml.StartElement:
			tok := token.(xml.StartElement)
			switch tok.Name.Local {
			default:
				attr := start.Attr
				if attr == nil || len(attr) <= 0 {
					// TODO: workaround for bug with nested lists
					attr = tok.Attr
				}
				if attr == nil || len(attr) <= 0 {
					panic("Couldn't determine class type for childs of S7ParameterUserDataItem")
				}
				switch attr[0].Value {
				case "org.apache.plc4x.java.s7.readwrite.S7ParameterUserDataItemCPUFunctions":
					var dt *S7ParameterUserDataItemCPUFunctions
					if m.Child != nil {
						dt = m.Child.(*S7ParameterUserDataItemCPUFunctions)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				}
			}
		}
	}
}

func (m *S7ParameterUserDataItem) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	className := reflect.TypeOf(m.Child).String()
	className = "org.apache.plc4x.java.s7.readwrite." + className[strings.LastIndex(className, ".")+1:]
	if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
		{Name: xml.Name{Local: "className"}, Value: className},
	}}); err != nil {
		return err
	}
	marshaller, ok := m.Child.(xml.Marshaler)
	if !ok {
		return errors.Errorf("child is not castable to Marshaler. Actual type %T", m.Child)
	}
	if err := marshaller.MarshalXML(e, start); err != nil {
		return err
	}
	if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
		return err
	}
	return nil
}

func (m S7ParameterUserDataItem) String() string {
	return string(m.Box("S7ParameterUserDataItem", utils.DefaultWidth*2))
}

func (m S7ParameterUserDataItem) Box(name string, width int) utils.AsciiBox {
	if name == "" {
		name = "S7ParameterUserDataItem"
	}
	boxes := make([]utils.AsciiBox, 0)
	boxes = append(boxes, utils.BoxAnything("", m.Child, width-2))
	return utils.BoxBox(name, utils.AlignBoxes(boxes, width-2), 0)
}
