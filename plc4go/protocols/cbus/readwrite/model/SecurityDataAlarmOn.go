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

// SecurityDataAlarmOn is the corresponding interface of SecurityDataAlarmOn
type SecurityDataAlarmOn interface {
	utils.LengthAware
	utils.Serializable
	SecurityData
}

// SecurityDataAlarmOnExactly can be used when we want exactly this type and not a type which fulfills SecurityDataAlarmOn.
// This is useful for switch cases.
type SecurityDataAlarmOnExactly interface {
	SecurityDataAlarmOn
	isSecurityDataAlarmOn() bool
}

// _SecurityDataAlarmOn is the data-structure of this message
type _SecurityDataAlarmOn struct {
	*_SecurityData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SecurityDataAlarmOn) InitializeParent(parent SecurityData, commandTypeContainer SecurityCommandTypeContainer, argument byte) {
	m.CommandTypeContainer = commandTypeContainer
	m.Argument = argument
}

func (m *_SecurityDataAlarmOn) GetParent() SecurityData {
	return m._SecurityData
}

// NewSecurityDataAlarmOn factory function for _SecurityDataAlarmOn
func NewSecurityDataAlarmOn(commandTypeContainer SecurityCommandTypeContainer, argument byte) *_SecurityDataAlarmOn {
	_result := &_SecurityDataAlarmOn{
		_SecurityData: NewSecurityData(commandTypeContainer, argument),
	}
	_result._SecurityData._SecurityDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSecurityDataAlarmOn(structType interface{}) SecurityDataAlarmOn {
	if casted, ok := structType.(SecurityDataAlarmOn); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataAlarmOn); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataAlarmOn) GetTypeName() string {
	return "SecurityDataAlarmOn"
}

func (m *_SecurityDataAlarmOn) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_SecurityDataAlarmOn) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *_SecurityDataAlarmOn) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func SecurityDataAlarmOnParse(readBuffer utils.ReadBuffer) (SecurityDataAlarmOn, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataAlarmOn"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataAlarmOn")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SecurityDataAlarmOn"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataAlarmOn")
	}

	// Create a partially initialized instance
	_child := &_SecurityDataAlarmOn{
		_SecurityData: &_SecurityData{},
	}
	_child._SecurityData._SecurityDataChildRequirements = _child
	return _child, nil
}

func (m *_SecurityDataAlarmOn) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataAlarmOn) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataAlarmOn"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataAlarmOn")
		}

		if popErr := writeBuffer.PopContext("SecurityDataAlarmOn"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataAlarmOn")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_SecurityDataAlarmOn) isSecurityDataAlarmOn() bool {
	return true
}

func (m *_SecurityDataAlarmOn) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
