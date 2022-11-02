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

// BACnetLogRecordLogDatumEnumeratedValue is the corresponding interface of BACnetLogRecordLogDatumEnumeratedValue
type BACnetLogRecordLogDatumEnumeratedValue interface {
	utils.LengthAware
	utils.Serializable
	BACnetLogRecordLogDatum
	// GetEnumeratedValue returns EnumeratedValue (property field)
	GetEnumeratedValue() BACnetContextTagEnumerated
}

// BACnetLogRecordLogDatumEnumeratedValueExactly can be used when we want exactly this type and not a type which fulfills BACnetLogRecordLogDatumEnumeratedValue.
// This is useful for switch cases.
type BACnetLogRecordLogDatumEnumeratedValueExactly interface {
	BACnetLogRecordLogDatumEnumeratedValue
	isBACnetLogRecordLogDatumEnumeratedValue() bool
}

// _BACnetLogRecordLogDatumEnumeratedValue is the data-structure of this message
type _BACnetLogRecordLogDatumEnumeratedValue struct {
	*_BACnetLogRecordLogDatum
	EnumeratedValue BACnetContextTagEnumerated
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetLogRecordLogDatumEnumeratedValue) InitializeParent(parent BACnetLogRecordLogDatum, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetLogRecordLogDatumEnumeratedValue) GetParent() BACnetLogRecordLogDatum {
	return m._BACnetLogRecordLogDatum
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLogRecordLogDatumEnumeratedValue) GetEnumeratedValue() BACnetContextTagEnumerated {
	return m.EnumeratedValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetLogRecordLogDatumEnumeratedValue factory function for _BACnetLogRecordLogDatumEnumeratedValue
func NewBACnetLogRecordLogDatumEnumeratedValue(enumeratedValue BACnetContextTagEnumerated, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetLogRecordLogDatumEnumeratedValue {
	_result := &_BACnetLogRecordLogDatumEnumeratedValue{
		EnumeratedValue:          enumeratedValue,
		_BACnetLogRecordLogDatum: NewBACnetLogRecordLogDatum(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result._BACnetLogRecordLogDatum._BACnetLogRecordLogDatumChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetLogRecordLogDatumEnumeratedValue(structType interface{}) BACnetLogRecordLogDatumEnumeratedValue {
	if casted, ok := structType.(BACnetLogRecordLogDatumEnumeratedValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLogRecordLogDatumEnumeratedValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLogRecordLogDatumEnumeratedValue) GetTypeName() string {
	return "BACnetLogRecordLogDatumEnumeratedValue"
}

func (m *_BACnetLogRecordLogDatumEnumeratedValue) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetLogRecordLogDatumEnumeratedValue) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (enumeratedValue)
	lengthInBits += m.EnumeratedValue.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetLogRecordLogDatumEnumeratedValue) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetLogRecordLogDatumEnumeratedValueParse(readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetLogRecordLogDatumEnumeratedValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLogRecordLogDatumEnumeratedValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLogRecordLogDatumEnumeratedValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (enumeratedValue)
	if pullErr := readBuffer.PullContext("enumeratedValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for enumeratedValue")
	}
	_enumeratedValue, _enumeratedValueErr := BACnetContextTagParse(readBuffer, uint8(uint8(3)), BACnetDataType(BACnetDataType_ENUMERATED))
	if _enumeratedValueErr != nil {
		return nil, errors.Wrap(_enumeratedValueErr, "Error parsing 'enumeratedValue' field of BACnetLogRecordLogDatumEnumeratedValue")
	}
	enumeratedValue := _enumeratedValue.(BACnetContextTagEnumerated)
	if closeErr := readBuffer.CloseContext("enumeratedValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for enumeratedValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetLogRecordLogDatumEnumeratedValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLogRecordLogDatumEnumeratedValue")
	}

	// Create a partially initialized instance
	_child := &_BACnetLogRecordLogDatumEnumeratedValue{
		_BACnetLogRecordLogDatum: &_BACnetLogRecordLogDatum{
			TagNumber: tagNumber,
		},
		EnumeratedValue: enumeratedValue,
	}
	_child._BACnetLogRecordLogDatum._BACnetLogRecordLogDatumChildRequirements = _child
	return _child, nil
}

func (m *_BACnetLogRecordLogDatumEnumeratedValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLogRecordLogDatumEnumeratedValue) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetLogRecordLogDatumEnumeratedValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetLogRecordLogDatumEnumeratedValue")
		}

		// Simple Field (enumeratedValue)
		if pushErr := writeBuffer.PushContext("enumeratedValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for enumeratedValue")
		}
		_enumeratedValueErr := writeBuffer.WriteSerializable(m.GetEnumeratedValue())
		if popErr := writeBuffer.PopContext("enumeratedValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for enumeratedValue")
		}
		if _enumeratedValueErr != nil {
			return errors.Wrap(_enumeratedValueErr, "Error serializing 'enumeratedValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetLogRecordLogDatumEnumeratedValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetLogRecordLogDatumEnumeratedValue")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetLogRecordLogDatumEnumeratedValue) isBACnetLogRecordLogDatumEnumeratedValue() bool {
	return true
}

func (m *_BACnetLogRecordLogDatumEnumeratedValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
