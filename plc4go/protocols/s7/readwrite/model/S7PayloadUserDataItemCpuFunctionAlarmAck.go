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

// S7PayloadUserDataItemCpuFunctionAlarmAck is the corresponding interface of S7PayloadUserDataItemCpuFunctionAlarmAck
type S7PayloadUserDataItemCpuFunctionAlarmAck interface {
	utils.LengthAware
	utils.Serializable
	S7PayloadUserDataItem
	// GetFunctionId returns FunctionId (property field)
	GetFunctionId() uint8
	// GetMessageObjects returns MessageObjects (property field)
	GetMessageObjects() []AlarmMessageObjectAckType
}

// S7PayloadUserDataItemCpuFunctionAlarmAckExactly can be used when we want exactly this type and not a type which fulfills S7PayloadUserDataItemCpuFunctionAlarmAck.
// This is useful for switch cases.
type S7PayloadUserDataItemCpuFunctionAlarmAckExactly interface {
	S7PayloadUserDataItemCpuFunctionAlarmAck
	isS7PayloadUserDataItemCpuFunctionAlarmAck() bool
}

// _S7PayloadUserDataItemCpuFunctionAlarmAck is the data-structure of this message
type _S7PayloadUserDataItemCpuFunctionAlarmAck struct {
	*_S7PayloadUserDataItem
	FunctionId     uint8
	MessageObjects []AlarmMessageObjectAckType
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7PayloadUserDataItemCpuFunctionAlarmAck) GetCpuFunctionType() uint8 {
	return 0x04
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmAck) GetCpuSubfunction() uint8 {
	return 0x0b
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmAck) GetDataLength() uint16 {
	return 0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7PayloadUserDataItemCpuFunctionAlarmAck) InitializeParent(parent S7PayloadUserDataItem, returnCode DataTransportErrorCode, transportSize DataTransportSize) {
	m.ReturnCode = returnCode
	m.TransportSize = transportSize
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmAck) GetParent() S7PayloadUserDataItem {
	return m._S7PayloadUserDataItem
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7PayloadUserDataItemCpuFunctionAlarmAck) GetFunctionId() uint8 {
	return m.FunctionId
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmAck) GetMessageObjects() []AlarmMessageObjectAckType {
	return m.MessageObjects
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewS7PayloadUserDataItemCpuFunctionAlarmAck factory function for _S7PayloadUserDataItemCpuFunctionAlarmAck
func NewS7PayloadUserDataItemCpuFunctionAlarmAck(functionId uint8, messageObjects []AlarmMessageObjectAckType, returnCode DataTransportErrorCode, transportSize DataTransportSize) *_S7PayloadUserDataItemCpuFunctionAlarmAck {
	_result := &_S7PayloadUserDataItemCpuFunctionAlarmAck{
		FunctionId:             functionId,
		MessageObjects:         messageObjects,
		_S7PayloadUserDataItem: NewS7PayloadUserDataItem(returnCode, transportSize),
	}
	_result._S7PayloadUserDataItem._S7PayloadUserDataItemChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastS7PayloadUserDataItemCpuFunctionAlarmAck(structType interface{}) S7PayloadUserDataItemCpuFunctionAlarmAck {
	if casted, ok := structType.(S7PayloadUserDataItemCpuFunctionAlarmAck); ok {
		return casted
	}
	if casted, ok := structType.(*S7PayloadUserDataItemCpuFunctionAlarmAck); ok {
		return *casted
	}
	return nil
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmAck) GetTypeName() string {
	return "S7PayloadUserDataItemCpuFunctionAlarmAck"
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmAck) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmAck) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (functionId)
	lengthInBits += 8

	// Implicit Field (numberOfObjects)
	lengthInBits += 8

	// Array field
	if len(m.MessageObjects) > 0 {
		for i, element := range m.MessageObjects {
			last := i == len(m.MessageObjects)-1
			lengthInBits += element.(interface{ GetLengthInBitsConditional(bool) uint16 }).GetLengthInBitsConditional(last)
		}
	}

	return lengthInBits
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmAck) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func S7PayloadUserDataItemCpuFunctionAlarmAckParse(readBuffer utils.ReadBuffer, cpuFunctionType uint8, cpuSubfunction uint8) (S7PayloadUserDataItemCpuFunctionAlarmAck, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7PayloadUserDataItemCpuFunctionAlarmAck"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7PayloadUserDataItemCpuFunctionAlarmAck")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (functionId)
	_functionId, _functionIdErr := readBuffer.ReadUint8("functionId", 8)
	if _functionIdErr != nil {
		return nil, errors.Wrap(_functionIdErr, "Error parsing 'functionId' field of S7PayloadUserDataItemCpuFunctionAlarmAck")
	}
	functionId := _functionId

	// Implicit Field (numberOfObjects) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	numberOfObjects, _numberOfObjectsErr := readBuffer.ReadUint8("numberOfObjects", 8)
	_ = numberOfObjects
	if _numberOfObjectsErr != nil {
		return nil, errors.Wrap(_numberOfObjectsErr, "Error parsing 'numberOfObjects' field of S7PayloadUserDataItemCpuFunctionAlarmAck")
	}

	// Array field (messageObjects)
	if pullErr := readBuffer.PullContext("messageObjects", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for messageObjects")
	}
	// Count array
	messageObjects := make([]AlarmMessageObjectAckType, numberOfObjects)
	// This happens when the size is set conditional to 0
	if len(messageObjects) == 0 {
		messageObjects = nil
	}
	{
		for curItem := uint16(0); curItem < uint16(numberOfObjects); curItem++ {
			_item, _err := AlarmMessageObjectAckTypeParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'messageObjects' field of S7PayloadUserDataItemCpuFunctionAlarmAck")
			}
			messageObjects[curItem] = _item.(AlarmMessageObjectAckType)
		}
	}
	if closeErr := readBuffer.CloseContext("messageObjects", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for messageObjects")
	}

	if closeErr := readBuffer.CloseContext("S7PayloadUserDataItemCpuFunctionAlarmAck"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7PayloadUserDataItemCpuFunctionAlarmAck")
	}

	// Create a partially initialized instance
	_child := &_S7PayloadUserDataItemCpuFunctionAlarmAck{
		_S7PayloadUserDataItem: &_S7PayloadUserDataItem{},
		FunctionId:             functionId,
		MessageObjects:         messageObjects,
	}
	_child._S7PayloadUserDataItem._S7PayloadUserDataItemChildRequirements = _child
	return _child, nil
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmAck) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmAck) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadUserDataItemCpuFunctionAlarmAck"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7PayloadUserDataItemCpuFunctionAlarmAck")
		}

		// Simple Field (functionId)
		functionId := uint8(m.GetFunctionId())
		_functionIdErr := writeBuffer.WriteUint8("functionId", 8, (functionId))
		if _functionIdErr != nil {
			return errors.Wrap(_functionIdErr, "Error serializing 'functionId' field")
		}

		// Implicit Field (numberOfObjects) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		numberOfObjects := uint8(uint8(len(m.GetMessageObjects())))
		_numberOfObjectsErr := writeBuffer.WriteUint8("numberOfObjects", 8, (numberOfObjects))
		if _numberOfObjectsErr != nil {
			return errors.Wrap(_numberOfObjectsErr, "Error serializing 'numberOfObjects' field")
		}

		// Array Field (messageObjects)
		if pushErr := writeBuffer.PushContext("messageObjects", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for messageObjects")
		}
		for _, _element := range m.GetMessageObjects() {
			_elementErr := writeBuffer.WriteSerializable(_element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'messageObjects' field")
			}
		}
		if popErr := writeBuffer.PopContext("messageObjects", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for messageObjects")
		}

		if popErr := writeBuffer.PopContext("S7PayloadUserDataItemCpuFunctionAlarmAck"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7PayloadUserDataItemCpuFunctionAlarmAck")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmAck) isS7PayloadUserDataItemCpuFunctionAlarmAck() bool {
	return true
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmAck) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
