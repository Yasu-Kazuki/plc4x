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

// BACnetConstructedDataCountChangeTime is the corresponding interface of BACnetConstructedDataCountChangeTime
type BACnetConstructedDataCountChangeTime interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetCountChangeTime returns CountChangeTime (property field)
	GetCountChangeTime() BACnetDateTime
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetDateTime
}

// BACnetConstructedDataCountChangeTimeExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataCountChangeTime.
// This is useful for switch cases.
type BACnetConstructedDataCountChangeTimeExactly interface {
	BACnetConstructedDataCountChangeTime
	isBACnetConstructedDataCountChangeTime() bool
}

// _BACnetConstructedDataCountChangeTime is the data-structure of this message
type _BACnetConstructedDataCountChangeTime struct {
	*_BACnetConstructedData
	CountChangeTime BACnetDateTime
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataCountChangeTime) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataCountChangeTime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_COUNT_CHANGE_TIME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataCountChangeTime) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataCountChangeTime) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataCountChangeTime) GetCountChangeTime() BACnetDateTime {
	return m.CountChangeTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataCountChangeTime) GetActualValue() BACnetDateTime {
	return CastBACnetDateTime(m.GetCountChangeTime())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataCountChangeTime factory function for _BACnetConstructedDataCountChangeTime
func NewBACnetConstructedDataCountChangeTime(countChangeTime BACnetDateTime, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataCountChangeTime {
	_result := &_BACnetConstructedDataCountChangeTime{
		CountChangeTime:        countChangeTime,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataCountChangeTime(structType interface{}) BACnetConstructedDataCountChangeTime {
	if casted, ok := structType.(BACnetConstructedDataCountChangeTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataCountChangeTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataCountChangeTime) GetTypeName() string {
	return "BACnetConstructedDataCountChangeTime"
}

func (m *_BACnetConstructedDataCountChangeTime) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataCountChangeTime) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (countChangeTime)
	lengthInBits += m.CountChangeTime.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataCountChangeTime) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataCountChangeTimeParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataCountChangeTime, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataCountChangeTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataCountChangeTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (countChangeTime)
	if pullErr := readBuffer.PullContext("countChangeTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for countChangeTime")
	}
	_countChangeTime, _countChangeTimeErr := BACnetDateTimeParse(readBuffer)
	if _countChangeTimeErr != nil {
		return nil, errors.Wrap(_countChangeTimeErr, "Error parsing 'countChangeTime' field of BACnetConstructedDataCountChangeTime")
	}
	countChangeTime := _countChangeTime.(BACnetDateTime)
	if closeErr := readBuffer.CloseContext("countChangeTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for countChangeTime")
	}

	// Virtual field
	_actualValue := countChangeTime
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataCountChangeTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataCountChangeTime")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataCountChangeTime{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		CountChangeTime: countChangeTime,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataCountChangeTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataCountChangeTime) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataCountChangeTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataCountChangeTime")
		}

		// Simple Field (countChangeTime)
		if pushErr := writeBuffer.PushContext("countChangeTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for countChangeTime")
		}
		_countChangeTimeErr := writeBuffer.WriteSerializable(m.GetCountChangeTime())
		if popErr := writeBuffer.PopContext("countChangeTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for countChangeTime")
		}
		if _countChangeTimeErr != nil {
			return errors.Wrap(_countChangeTimeErr, "Error serializing 'countChangeTime' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataCountChangeTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataCountChangeTime")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataCountChangeTime) isBACnetConstructedDataCountChangeTime() bool {
	return true
}

func (m *_BACnetConstructedDataCountChangeTime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
