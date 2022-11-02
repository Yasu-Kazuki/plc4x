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

// SecurityDataGasAlarmCleared is the corresponding interface of SecurityDataGasAlarmCleared
type SecurityDataGasAlarmCleared interface {
	utils.LengthAware
	utils.Serializable
	SecurityData
}

// SecurityDataGasAlarmClearedExactly can be used when we want exactly this type and not a type which fulfills SecurityDataGasAlarmCleared.
// This is useful for switch cases.
type SecurityDataGasAlarmClearedExactly interface {
	SecurityDataGasAlarmCleared
	isSecurityDataGasAlarmCleared() bool
}

// _SecurityDataGasAlarmCleared is the data-structure of this message
type _SecurityDataGasAlarmCleared struct {
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

func (m *_SecurityDataGasAlarmCleared) InitializeParent(parent SecurityData, commandTypeContainer SecurityCommandTypeContainer, argument byte) {
	m.CommandTypeContainer = commandTypeContainer
	m.Argument = argument
}

func (m *_SecurityDataGasAlarmCleared) GetParent() SecurityData {
	return m._SecurityData
}

// NewSecurityDataGasAlarmCleared factory function for _SecurityDataGasAlarmCleared
func NewSecurityDataGasAlarmCleared(commandTypeContainer SecurityCommandTypeContainer, argument byte) *_SecurityDataGasAlarmCleared {
	_result := &_SecurityDataGasAlarmCleared{
		_SecurityData: NewSecurityData(commandTypeContainer, argument),
	}
	_result._SecurityData._SecurityDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSecurityDataGasAlarmCleared(structType interface{}) SecurityDataGasAlarmCleared {
	if casted, ok := structType.(SecurityDataGasAlarmCleared); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataGasAlarmCleared); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataGasAlarmCleared) GetTypeName() string {
	return "SecurityDataGasAlarmCleared"
}

func (m *_SecurityDataGasAlarmCleared) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_SecurityDataGasAlarmCleared) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *_SecurityDataGasAlarmCleared) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func SecurityDataGasAlarmClearedParse(readBuffer utils.ReadBuffer) (SecurityDataGasAlarmCleared, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataGasAlarmCleared"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataGasAlarmCleared")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SecurityDataGasAlarmCleared"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataGasAlarmCleared")
	}

	// Create a partially initialized instance
	_child := &_SecurityDataGasAlarmCleared{
		_SecurityData: &_SecurityData{},
	}
	_child._SecurityData._SecurityDataChildRequirements = _child
	return _child, nil
}

func (m *_SecurityDataGasAlarmCleared) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataGasAlarmCleared) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataGasAlarmCleared"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataGasAlarmCleared")
		}

		if popErr := writeBuffer.PopContext("SecurityDataGasAlarmCleared"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataGasAlarmCleared")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_SecurityDataGasAlarmCleared) isSecurityDataGasAlarmCleared() bool {
	return true
}

func (m *_SecurityDataGasAlarmCleared) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
