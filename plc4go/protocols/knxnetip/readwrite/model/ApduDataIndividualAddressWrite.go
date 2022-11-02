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

// ApduDataIndividualAddressWrite is the corresponding interface of ApduDataIndividualAddressWrite
type ApduDataIndividualAddressWrite interface {
	utils.LengthAware
	utils.Serializable
	ApduData
}

// ApduDataIndividualAddressWriteExactly can be used when we want exactly this type and not a type which fulfills ApduDataIndividualAddressWrite.
// This is useful for switch cases.
type ApduDataIndividualAddressWriteExactly interface {
	ApduDataIndividualAddressWrite
	isApduDataIndividualAddressWrite() bool
}

// _ApduDataIndividualAddressWrite is the data-structure of this message
type _ApduDataIndividualAddressWrite struct {
	*_ApduData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataIndividualAddressWrite) GetApciType() uint8 {
	return 0x3
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataIndividualAddressWrite) InitializeParent(parent ApduData) {}

func (m *_ApduDataIndividualAddressWrite) GetParent() ApduData {
	return m._ApduData
}

// NewApduDataIndividualAddressWrite factory function for _ApduDataIndividualAddressWrite
func NewApduDataIndividualAddressWrite(dataLength uint8) *_ApduDataIndividualAddressWrite {
	_result := &_ApduDataIndividualAddressWrite{
		_ApduData: NewApduData(dataLength),
	}
	_result._ApduData._ApduDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastApduDataIndividualAddressWrite(structType interface{}) ApduDataIndividualAddressWrite {
	if casted, ok := structType.(ApduDataIndividualAddressWrite); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataIndividualAddressWrite); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataIndividualAddressWrite) GetTypeName() string {
	return "ApduDataIndividualAddressWrite"
}

func (m *_ApduDataIndividualAddressWrite) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ApduDataIndividualAddressWrite) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *_ApduDataIndividualAddressWrite) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApduDataIndividualAddressWriteParse(readBuffer utils.ReadBuffer, dataLength uint8) (ApduDataIndividualAddressWrite, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataIndividualAddressWrite"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataIndividualAddressWrite")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataIndividualAddressWrite"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataIndividualAddressWrite")
	}

	// Create a partially initialized instance
	_child := &_ApduDataIndividualAddressWrite{
		_ApduData: &_ApduData{
			DataLength: dataLength,
		},
	}
	_child._ApduData._ApduDataChildRequirements = _child
	return _child, nil
}

func (m *_ApduDataIndividualAddressWrite) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataIndividualAddressWrite) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataIndividualAddressWrite"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataIndividualAddressWrite")
		}

		if popErr := writeBuffer.PopContext("ApduDataIndividualAddressWrite"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataIndividualAddressWrite")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_ApduDataIndividualAddressWrite) isApduDataIndividualAddressWrite() bool {
	return true
}

func (m *_ApduDataIndividualAddressWrite) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
