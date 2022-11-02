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

// ApduDataExtPropertyValueWrite is the corresponding interface of ApduDataExtPropertyValueWrite
type ApduDataExtPropertyValueWrite interface {
	utils.LengthAware
	utils.Serializable
	ApduDataExt
	// GetObjectIndex returns ObjectIndex (property field)
	GetObjectIndex() uint8
	// GetPropertyId returns PropertyId (property field)
	GetPropertyId() uint8
	// GetCount returns Count (property field)
	GetCount() uint8
	// GetIndex returns Index (property field)
	GetIndex() uint16
	// GetData returns Data (property field)
	GetData() []byte
}

// ApduDataExtPropertyValueWriteExactly can be used when we want exactly this type and not a type which fulfills ApduDataExtPropertyValueWrite.
// This is useful for switch cases.
type ApduDataExtPropertyValueWriteExactly interface {
	ApduDataExtPropertyValueWrite
	isApduDataExtPropertyValueWrite() bool
}

// _ApduDataExtPropertyValueWrite is the data-structure of this message
type _ApduDataExtPropertyValueWrite struct {
	*_ApduDataExt
	ObjectIndex uint8
	PropertyId  uint8
	Count       uint8
	Index       uint16
	Data        []byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataExtPropertyValueWrite) GetExtApciType() uint8 {
	return 0x17
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataExtPropertyValueWrite) InitializeParent(parent ApduDataExt) {}

func (m *_ApduDataExtPropertyValueWrite) GetParent() ApduDataExt {
	return m._ApduDataExt
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ApduDataExtPropertyValueWrite) GetObjectIndex() uint8 {
	return m.ObjectIndex
}

func (m *_ApduDataExtPropertyValueWrite) GetPropertyId() uint8 {
	return m.PropertyId
}

func (m *_ApduDataExtPropertyValueWrite) GetCount() uint8 {
	return m.Count
}

func (m *_ApduDataExtPropertyValueWrite) GetIndex() uint16 {
	return m.Index
}

func (m *_ApduDataExtPropertyValueWrite) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewApduDataExtPropertyValueWrite factory function for _ApduDataExtPropertyValueWrite
func NewApduDataExtPropertyValueWrite(objectIndex uint8, propertyId uint8, count uint8, index uint16, data []byte, length uint8) *_ApduDataExtPropertyValueWrite {
	_result := &_ApduDataExtPropertyValueWrite{
		ObjectIndex:  objectIndex,
		PropertyId:   propertyId,
		Count:        count,
		Index:        index,
		Data:         data,
		_ApduDataExt: NewApduDataExt(length),
	}
	_result._ApduDataExt._ApduDataExtChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastApduDataExtPropertyValueWrite(structType interface{}) ApduDataExtPropertyValueWrite {
	if casted, ok := structType.(ApduDataExtPropertyValueWrite); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataExtPropertyValueWrite); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataExtPropertyValueWrite) GetTypeName() string {
	return "ApduDataExtPropertyValueWrite"
}

func (m *_ApduDataExtPropertyValueWrite) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ApduDataExtPropertyValueWrite) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (objectIndex)
	lengthInBits += 8

	// Simple field (propertyId)
	lengthInBits += 8

	// Simple field (count)
	lengthInBits += 4

	// Simple field (index)
	lengthInBits += 12

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_ApduDataExtPropertyValueWrite) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApduDataExtPropertyValueWriteParse(readBuffer utils.ReadBuffer, length uint8) (ApduDataExtPropertyValueWrite, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtPropertyValueWrite"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtPropertyValueWrite")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (objectIndex)
	_objectIndex, _objectIndexErr := readBuffer.ReadUint8("objectIndex", 8)
	if _objectIndexErr != nil {
		return nil, errors.Wrap(_objectIndexErr, "Error parsing 'objectIndex' field of ApduDataExtPropertyValueWrite")
	}
	objectIndex := _objectIndex

	// Simple Field (propertyId)
	_propertyId, _propertyIdErr := readBuffer.ReadUint8("propertyId", 8)
	if _propertyIdErr != nil {
		return nil, errors.Wrap(_propertyIdErr, "Error parsing 'propertyId' field of ApduDataExtPropertyValueWrite")
	}
	propertyId := _propertyId

	// Simple Field (count)
	_count, _countErr := readBuffer.ReadUint8("count", 4)
	if _countErr != nil {
		return nil, errors.Wrap(_countErr, "Error parsing 'count' field of ApduDataExtPropertyValueWrite")
	}
	count := _count

	// Simple Field (index)
	_index, _indexErr := readBuffer.ReadUint16("index", 12)
	if _indexErr != nil {
		return nil, errors.Wrap(_indexErr, "Error parsing 'index' field of ApduDataExtPropertyValueWrite")
	}
	index := _index
	// Byte Array field (data)
	numberOfBytesdata := int(uint16(length) - uint16(uint16(5)))
	data, _readArrayErr := readBuffer.ReadByteArray("data", numberOfBytesdata)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'data' field of ApduDataExtPropertyValueWrite")
	}

	if closeErr := readBuffer.CloseContext("ApduDataExtPropertyValueWrite"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtPropertyValueWrite")
	}

	// Create a partially initialized instance
	_child := &_ApduDataExtPropertyValueWrite{
		_ApduDataExt: &_ApduDataExt{
			Length: length,
		},
		ObjectIndex: objectIndex,
		PropertyId:  propertyId,
		Count:       count,
		Index:       index,
		Data:        data,
	}
	_child._ApduDataExt._ApduDataExtChildRequirements = _child
	return _child, nil
}

func (m *_ApduDataExtPropertyValueWrite) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataExtPropertyValueWrite) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtPropertyValueWrite"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtPropertyValueWrite")
		}

		// Simple Field (objectIndex)
		objectIndex := uint8(m.GetObjectIndex())
		_objectIndexErr := writeBuffer.WriteUint8("objectIndex", 8, (objectIndex))
		if _objectIndexErr != nil {
			return errors.Wrap(_objectIndexErr, "Error serializing 'objectIndex' field")
		}

		// Simple Field (propertyId)
		propertyId := uint8(m.GetPropertyId())
		_propertyIdErr := writeBuffer.WriteUint8("propertyId", 8, (propertyId))
		if _propertyIdErr != nil {
			return errors.Wrap(_propertyIdErr, "Error serializing 'propertyId' field")
		}

		// Simple Field (count)
		count := uint8(m.GetCount())
		_countErr := writeBuffer.WriteUint8("count", 4, (count))
		if _countErr != nil {
			return errors.Wrap(_countErr, "Error serializing 'count' field")
		}

		// Simple Field (index)
		index := uint16(m.GetIndex())
		_indexErr := writeBuffer.WriteUint16("index", 12, (index))
		if _indexErr != nil {
			return errors.Wrap(_indexErr, "Error serializing 'index' field")
		}

		// Array Field (data)
		// Byte Array field (data)
		if err := writeBuffer.WriteByteArray("data", m.GetData()); err != nil {
			return errors.Wrap(err, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtPropertyValueWrite"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtPropertyValueWrite")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_ApduDataExtPropertyValueWrite) isApduDataExtPropertyValueWrite() bool {
	return true
}

func (m *_ApduDataExtPropertyValueWrite) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
