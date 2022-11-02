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

// ModbusPDUWriteSingleRegisterResponse is the corresponding interface of ModbusPDUWriteSingleRegisterResponse
type ModbusPDUWriteSingleRegisterResponse interface {
	utils.LengthAware
	utils.Serializable
	ModbusPDU
	// GetAddress returns Address (property field)
	GetAddress() uint16
	// GetValue returns Value (property field)
	GetValue() uint16
}

// ModbusPDUWriteSingleRegisterResponseExactly can be used when we want exactly this type and not a type which fulfills ModbusPDUWriteSingleRegisterResponse.
// This is useful for switch cases.
type ModbusPDUWriteSingleRegisterResponseExactly interface {
	ModbusPDUWriteSingleRegisterResponse
	isModbusPDUWriteSingleRegisterResponse() bool
}

// _ModbusPDUWriteSingleRegisterResponse is the data-structure of this message
type _ModbusPDUWriteSingleRegisterResponse struct {
	*_ModbusPDU
	Address uint16
	Value   uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusPDUWriteSingleRegisterResponse) GetErrorFlag() bool {
	return bool(false)
}

func (m *_ModbusPDUWriteSingleRegisterResponse) GetFunctionFlag() uint8 {
	return 0x06
}

func (m *_ModbusPDUWriteSingleRegisterResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusPDUWriteSingleRegisterResponse) InitializeParent(parent ModbusPDU) {}

func (m *_ModbusPDUWriteSingleRegisterResponse) GetParent() ModbusPDU {
	return m._ModbusPDU
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusPDUWriteSingleRegisterResponse) GetAddress() uint16 {
	return m.Address
}

func (m *_ModbusPDUWriteSingleRegisterResponse) GetValue() uint16 {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewModbusPDUWriteSingleRegisterResponse factory function for _ModbusPDUWriteSingleRegisterResponse
func NewModbusPDUWriteSingleRegisterResponse(address uint16, value uint16) *_ModbusPDUWriteSingleRegisterResponse {
	_result := &_ModbusPDUWriteSingleRegisterResponse{
		Address:    address,
		Value:      value,
		_ModbusPDU: NewModbusPDU(),
	}
	_result._ModbusPDU._ModbusPDUChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastModbusPDUWriteSingleRegisterResponse(structType interface{}) ModbusPDUWriteSingleRegisterResponse {
	if casted, ok := structType.(ModbusPDUWriteSingleRegisterResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUWriteSingleRegisterResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUWriteSingleRegisterResponse) GetTypeName() string {
	return "ModbusPDUWriteSingleRegisterResponse"
}

func (m *_ModbusPDUWriteSingleRegisterResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ModbusPDUWriteSingleRegisterResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (address)
	lengthInBits += 16

	// Simple field (value)
	lengthInBits += 16

	return lengthInBits
}

func (m *_ModbusPDUWriteSingleRegisterResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ModbusPDUWriteSingleRegisterResponseParse(readBuffer utils.ReadBuffer, response bool) (ModbusPDUWriteSingleRegisterResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUWriteSingleRegisterResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUWriteSingleRegisterResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (address)
	_address, _addressErr := readBuffer.ReadUint16("address", 16)
	if _addressErr != nil {
		return nil, errors.Wrap(_addressErr, "Error parsing 'address' field of ModbusPDUWriteSingleRegisterResponse")
	}
	address := _address

	// Simple Field (value)
	_value, _valueErr := readBuffer.ReadUint16("value", 16)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field of ModbusPDUWriteSingleRegisterResponse")
	}
	value := _value

	if closeErr := readBuffer.CloseContext("ModbusPDUWriteSingleRegisterResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUWriteSingleRegisterResponse")
	}

	// Create a partially initialized instance
	_child := &_ModbusPDUWriteSingleRegisterResponse{
		_ModbusPDU: &_ModbusPDU{},
		Address:    address,
		Value:      value,
	}
	_child._ModbusPDU._ModbusPDUChildRequirements = _child
	return _child, nil
}

func (m *_ModbusPDUWriteSingleRegisterResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusPDUWriteSingleRegisterResponse) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUWriteSingleRegisterResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUWriteSingleRegisterResponse")
		}

		// Simple Field (address)
		address := uint16(m.GetAddress())
		_addressErr := writeBuffer.WriteUint16("address", 16, (address))
		if _addressErr != nil {
			return errors.Wrap(_addressErr, "Error serializing 'address' field")
		}

		// Simple Field (value)
		value := uint16(m.GetValue())
		_valueErr := writeBuffer.WriteUint16("value", 16, (value))
		if _valueErr != nil {
			return errors.Wrap(_valueErr, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUWriteSingleRegisterResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUWriteSingleRegisterResponse")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_ModbusPDUWriteSingleRegisterResponse) isModbusPDUWriteSingleRegisterResponse() bool {
	return true
}

func (m *_ModbusPDUWriteSingleRegisterResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
