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
	"encoding/binary"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// S7VarPayloadStatusItem is the corresponding interface of S7VarPayloadStatusItem
type S7VarPayloadStatusItem interface {
	utils.LengthAware
	utils.Serializable
	// GetReturnCode returns ReturnCode (property field)
	GetReturnCode() DataTransportErrorCode
}

// S7VarPayloadStatusItemExactly can be used when we want exactly this type and not a type which fulfills S7VarPayloadStatusItem.
// This is useful for switch cases.
type S7VarPayloadStatusItemExactly interface {
	S7VarPayloadStatusItem
	isS7VarPayloadStatusItem() bool
}

// _S7VarPayloadStatusItem is the data-structure of this message
type _S7VarPayloadStatusItem struct {
	ReturnCode DataTransportErrorCode
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7VarPayloadStatusItem) GetReturnCode() DataTransportErrorCode {
	return m.ReturnCode
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewS7VarPayloadStatusItem factory function for _S7VarPayloadStatusItem
func NewS7VarPayloadStatusItem(returnCode DataTransportErrorCode) *_S7VarPayloadStatusItem {
	return &_S7VarPayloadStatusItem{ReturnCode: returnCode}
}

// Deprecated: use the interface for direct cast
func CastS7VarPayloadStatusItem(structType interface{}) S7VarPayloadStatusItem {
	if casted, ok := structType.(S7VarPayloadStatusItem); ok {
		return casted
	}
	if casted, ok := structType.(*S7VarPayloadStatusItem); ok {
		return *casted
	}
	return nil
}

func (m *_S7VarPayloadStatusItem) GetTypeName() string {
	return "S7VarPayloadStatusItem"
}

func (m *_S7VarPayloadStatusItem) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_S7VarPayloadStatusItem) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (returnCode)
	lengthInBits += 8

	return lengthInBits
}

func (m *_S7VarPayloadStatusItem) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func S7VarPayloadStatusItemParse(readBuffer utils.ReadBuffer) (S7VarPayloadStatusItem, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7VarPayloadStatusItem"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7VarPayloadStatusItem")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (returnCode)
	if pullErr := readBuffer.PullContext("returnCode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for returnCode")
	}
	_returnCode, _returnCodeErr := DataTransportErrorCodeParse(readBuffer)
	if _returnCodeErr != nil {
		return nil, errors.Wrap(_returnCodeErr, "Error parsing 'returnCode' field of S7VarPayloadStatusItem")
	}
	returnCode := _returnCode
	if closeErr := readBuffer.CloseContext("returnCode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for returnCode")
	}

	if closeErr := readBuffer.CloseContext("S7VarPayloadStatusItem"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7VarPayloadStatusItem")
	}

	// Create the instance
	return &_S7VarPayloadStatusItem{
		ReturnCode: returnCode,
	}, nil
}

func (m *_S7VarPayloadStatusItem) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7VarPayloadStatusItem) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("S7VarPayloadStatusItem"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for S7VarPayloadStatusItem")
	}

	// Simple Field (returnCode)
	if pushErr := writeBuffer.PushContext("returnCode"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for returnCode")
	}
	_returnCodeErr := writeBuffer.WriteSerializable(m.GetReturnCode())
	if popErr := writeBuffer.PopContext("returnCode"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for returnCode")
	}
	if _returnCodeErr != nil {
		return errors.Wrap(_returnCodeErr, "Error serializing 'returnCode' field")
	}

	if popErr := writeBuffer.PopContext("S7VarPayloadStatusItem"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for S7VarPayloadStatusItem")
	}
	return nil
}

func (m *_S7VarPayloadStatusItem) isS7VarPayloadStatusItem() bool {
	return true
}

func (m *_S7VarPayloadStatusItem) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
