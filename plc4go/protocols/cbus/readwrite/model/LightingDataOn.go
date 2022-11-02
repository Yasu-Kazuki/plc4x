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

// LightingDataOn is the corresponding interface of LightingDataOn
type LightingDataOn interface {
	utils.LengthAware
	utils.Serializable
	LightingData
	// GetGroup returns Group (property field)
	GetGroup() byte
}

// LightingDataOnExactly can be used when we want exactly this type and not a type which fulfills LightingDataOn.
// This is useful for switch cases.
type LightingDataOnExactly interface {
	LightingDataOn
	isLightingDataOn() bool
}

// _LightingDataOn is the data-structure of this message
type _LightingDataOn struct {
	*_LightingData
	Group byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_LightingDataOn) InitializeParent(parent LightingData, commandTypeContainer LightingCommandTypeContainer) {
	m.CommandTypeContainer = commandTypeContainer
}

func (m *_LightingDataOn) GetParent() LightingData {
	return m._LightingData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_LightingDataOn) GetGroup() byte {
	return m.Group
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewLightingDataOn factory function for _LightingDataOn
func NewLightingDataOn(group byte, commandTypeContainer LightingCommandTypeContainer) *_LightingDataOn {
	_result := &_LightingDataOn{
		Group:         group,
		_LightingData: NewLightingData(commandTypeContainer),
	}
	_result._LightingData._LightingDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastLightingDataOn(structType interface{}) LightingDataOn {
	if casted, ok := structType.(LightingDataOn); ok {
		return casted
	}
	if casted, ok := structType.(*LightingDataOn); ok {
		return *casted
	}
	return nil
}

func (m *_LightingDataOn) GetTypeName() string {
	return "LightingDataOn"
}

func (m *_LightingDataOn) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_LightingDataOn) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (group)
	lengthInBits += 8

	return lengthInBits
}

func (m *_LightingDataOn) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func LightingDataOnParse(readBuffer utils.ReadBuffer) (LightingDataOn, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("LightingDataOn"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for LightingDataOn")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (group)
	_group, _groupErr := readBuffer.ReadByte("group")
	if _groupErr != nil {
		return nil, errors.Wrap(_groupErr, "Error parsing 'group' field of LightingDataOn")
	}
	group := _group

	if closeErr := readBuffer.CloseContext("LightingDataOn"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for LightingDataOn")
	}

	// Create a partially initialized instance
	_child := &_LightingDataOn{
		_LightingData: &_LightingData{},
		Group:         group,
	}
	_child._LightingData._LightingDataChildRequirements = _child
	return _child, nil
}

func (m *_LightingDataOn) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_LightingDataOn) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("LightingDataOn"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for LightingDataOn")
		}

		// Simple Field (group)
		group := byte(m.GetGroup())
		_groupErr := writeBuffer.WriteByte("group", (group))
		if _groupErr != nil {
			return errors.Wrap(_groupErr, "Error serializing 'group' field")
		}

		if popErr := writeBuffer.PopContext("LightingDataOn"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for LightingDataOn")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_LightingDataOn) isLightingDataOn() bool {
	return true
}

func (m *_LightingDataOn) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
