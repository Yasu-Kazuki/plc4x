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
type AssociatedValueType struct {
	ReturnCode    DataTransportErrorCode
	TransportSize DataTransportSize
	ValueLength   uint16
	Data          []uint8
}

// The corresponding interface
type IAssociatedValueType interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

func NewAssociatedValueType(returnCode DataTransportErrorCode, transportSize DataTransportSize, valueLength uint16, data []uint8) *AssociatedValueType {
	return &AssociatedValueType{ReturnCode: returnCode, TransportSize: transportSize, ValueLength: valueLength, Data: data}
}

func CastAssociatedValueType(structType interface{}) *AssociatedValueType {
	castFunc := func(typ interface{}) *AssociatedValueType {
		if casted, ok := typ.(AssociatedValueType); ok {
			return &casted
		}
		if casted, ok := typ.(*AssociatedValueType); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *AssociatedValueType) GetTypeName() string {
	return "AssociatedValueType"
}

func (m *AssociatedValueType) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *AssociatedValueType) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (returnCode)
	lengthInBits += 8

	// Simple field (transportSize)
	lengthInBits += 8

	// Manual Field (valueLength)
	lengthInBits += uint16(uint16(2) * 8)

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *AssociatedValueType) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func AssociatedValueTypeParse(readBuffer utils.ReadBuffer) (*AssociatedValueType, error) {
	if pullErr := readBuffer.PullContext("AssociatedValueType"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (returnCode)
	if pullErr := readBuffer.PullContext("returnCode"); pullErr != nil {
		return nil, pullErr
	}
	returnCode, _returnCodeErr := DataTransportErrorCodeParse(readBuffer)
	if _returnCodeErr != nil {
		return nil, errors.Wrap(_returnCodeErr, "Error parsing 'returnCode' field")
	}
	if closeErr := readBuffer.CloseContext("returnCode"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (transportSize)
	if pullErr := readBuffer.PullContext("transportSize"); pullErr != nil {
		return nil, pullErr
	}
	transportSize, _transportSizeErr := DataTransportSizeParse(readBuffer)
	if _transportSizeErr != nil {
		return nil, errors.Wrap(_transportSizeErr, "Error parsing 'transportSize' field")
	}
	if closeErr := readBuffer.CloseContext("transportSize"); closeErr != nil {
		return nil, closeErr
	}

	// Manual Field (valueLength)
	if pullErr := readBuffer.PullContext("valueLength"); pullErr != nil {
		return nil, pullErr
	}
	valueLength, _valueLengthErr := S7EventHelperRightShift3(readBuffer)
	if _valueLengthErr != nil {
		return nil, errors.Wrap(_valueLengthErr, "Error parsing 'valueLength' field")
	}
	if closeErr := readBuffer.CloseContext("valueLength"); closeErr != nil {
		return nil, closeErr
	}

	// Array field (data)
	if pullErr := readBuffer.PullContext("data", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Count array
	data := make([]uint8, S7EventHelperEventItemLength(readBuffer, valueLength))
	for curItem := uint16(0); curItem < uint16(S7EventHelperEventItemLength(readBuffer, valueLength)); curItem++ {
		_item, _err := readBuffer.ReadUint8("", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'data' field")
		}
		data[curItem] = _item
	}
	if closeErr := readBuffer.CloseContext("data", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("AssociatedValueType"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewAssociatedValueType(returnCode, transportSize, valueLength, data), nil
}

func (m *AssociatedValueType) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("AssociatedValueType"); pushErr != nil {
		return pushErr
	}

	// Simple Field (returnCode)
	if pushErr := writeBuffer.PushContext("returnCode"); pushErr != nil {
		return pushErr
	}
	_returnCodeErr := m.ReturnCode.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("returnCode"); popErr != nil {
		return popErr
	}
	if _returnCodeErr != nil {
		return errors.Wrap(_returnCodeErr, "Error serializing 'returnCode' field")
	}

	// Simple Field (transportSize)
	if pushErr := writeBuffer.PushContext("transportSize"); pushErr != nil {
		return pushErr
	}
	_transportSizeErr := m.TransportSize.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("transportSize"); popErr != nil {
		return popErr
	}
	if _transportSizeErr != nil {
		return errors.Wrap(_transportSizeErr, "Error serializing 'transportSize' field")
	}

	// Manual Field (valueLength)
	if pushErr := writeBuffer.PushContext("transportSize"); pushErr != nil {
		return pushErr
	}
	_valueLengthErr := S7EventHelperLeftShift3(writeBuffer, m.ValueLength)
	if popErr := writeBuffer.PopContext("valueLength"); popErr != nil {
		return popErr
	}
	if _valueLengthErr != nil {
		return errors.Wrap(_valueLengthErr, "Error serializing 'valueLength' field")
	}

	// Array Field (data)
	if m.Data != nil {
		if pushErr := writeBuffer.PushContext("data", utils.WithRenderAsList(true)); pushErr != nil {
			return pushErr
		}
		for _, _element := range m.Data {
			_elementErr := writeBuffer.WriteUint8("", 8, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'data' field")
			}
		}
		if popErr := writeBuffer.PopContext("data", utils.WithRenderAsList(true)); popErr != nil {
			return popErr
		}
	}

	if popErr := writeBuffer.PopContext("AssociatedValueType"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *AssociatedValueType) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}