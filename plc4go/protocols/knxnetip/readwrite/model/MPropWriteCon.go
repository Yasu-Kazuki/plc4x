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

// MPropWriteCon is the corresponding interface of MPropWriteCon
type MPropWriteCon interface {
	utils.LengthAware
	utils.Serializable
	CEMI
}

// MPropWriteConExactly can be used when we want exactly this type and not a type which fulfills MPropWriteCon.
// This is useful for switch cases.
type MPropWriteConExactly interface {
	MPropWriteCon
	isMPropWriteCon() bool
}

// _MPropWriteCon is the data-structure of this message
type _MPropWriteCon struct {
	*_CEMI
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_MPropWriteCon) GetMessageCode() uint8 {
	return 0xF5
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MPropWriteCon) InitializeParent(parent CEMI) {}

func (m *_MPropWriteCon) GetParent() CEMI {
	return m._CEMI
}

// NewMPropWriteCon factory function for _MPropWriteCon
func NewMPropWriteCon(size uint16) *_MPropWriteCon {
	_result := &_MPropWriteCon{
		_CEMI: NewCEMI(size),
	}
	_result._CEMI._CEMIChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastMPropWriteCon(structType interface{}) MPropWriteCon {
	if casted, ok := structType.(MPropWriteCon); ok {
		return casted
	}
	if casted, ok := structType.(*MPropWriteCon); ok {
		return *casted
	}
	return nil
}

func (m *_MPropWriteCon) GetTypeName() string {
	return "MPropWriteCon"
}

func (m *_MPropWriteCon) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_MPropWriteCon) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *_MPropWriteCon) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func MPropWriteConParse(readBuffer utils.ReadBuffer, size uint16) (MPropWriteCon, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MPropWriteCon"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MPropWriteCon")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("MPropWriteCon"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MPropWriteCon")
	}

	// Create a partially initialized instance
	_child := &_MPropWriteCon{
		_CEMI: &_CEMI{
			Size: size,
		},
	}
	_child._CEMI._CEMIChildRequirements = _child
	return _child, nil
}

func (m *_MPropWriteCon) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MPropWriteCon) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MPropWriteCon"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MPropWriteCon")
		}

		if popErr := writeBuffer.PopContext("MPropWriteCon"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MPropWriteCon")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_MPropWriteCon) isMPropWriteCon() bool {
	return true
}

func (m *_MPropWriteCon) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
