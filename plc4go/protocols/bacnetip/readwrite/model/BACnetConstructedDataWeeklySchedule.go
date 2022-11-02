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
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataWeeklySchedule is the corresponding interface of BACnetConstructedDataWeeklySchedule
type BACnetConstructedDataWeeklySchedule interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetNumberOfDataElements returns NumberOfDataElements (property field)
	GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger
	// GetWeeklySchedule returns WeeklySchedule (property field)
	GetWeeklySchedule() []BACnetDailySchedule
	// GetZero returns Zero (virtual field)
	GetZero() uint64
}

// BACnetConstructedDataWeeklyScheduleExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataWeeklySchedule.
// This is useful for switch cases.
type BACnetConstructedDataWeeklyScheduleExactly interface {
	BACnetConstructedDataWeeklySchedule
	isBACnetConstructedDataWeeklySchedule() bool
}

// _BACnetConstructedDataWeeklySchedule is the data-structure of this message
type _BACnetConstructedDataWeeklySchedule struct {
	*_BACnetConstructedData
	NumberOfDataElements BACnetApplicationTagUnsignedInteger
	WeeklySchedule       []BACnetDailySchedule
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataWeeklySchedule) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataWeeklySchedule) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_WEEKLY_SCHEDULE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataWeeklySchedule) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataWeeklySchedule) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataWeeklySchedule) GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger {
	return m.NumberOfDataElements
}

func (m *_BACnetConstructedDataWeeklySchedule) GetWeeklySchedule() []BACnetDailySchedule {
	return m.WeeklySchedule
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataWeeklySchedule) GetZero() uint64 {
	numberOfDataElements := m.NumberOfDataElements
	_ = numberOfDataElements
	return uint64(uint64(0))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataWeeklySchedule factory function for _BACnetConstructedDataWeeklySchedule
func NewBACnetConstructedDataWeeklySchedule(numberOfDataElements BACnetApplicationTagUnsignedInteger, weeklySchedule []BACnetDailySchedule, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataWeeklySchedule {
	_result := &_BACnetConstructedDataWeeklySchedule{
		NumberOfDataElements:   numberOfDataElements,
		WeeklySchedule:         weeklySchedule,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataWeeklySchedule(structType interface{}) BACnetConstructedDataWeeklySchedule {
	if casted, ok := structType.(BACnetConstructedDataWeeklySchedule); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataWeeklySchedule); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataWeeklySchedule) GetTypeName() string {
	return "BACnetConstructedDataWeeklySchedule"
}

func (m *_BACnetConstructedDataWeeklySchedule) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataWeeklySchedule) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// A virtual field doesn't have any in- or output.

	// Optional Field (numberOfDataElements)
	if m.NumberOfDataElements != nil {
		lengthInBits += m.NumberOfDataElements.GetLengthInBits()
	}

	// Array field
	if len(m.WeeklySchedule) > 0 {
		for _, element := range m.WeeklySchedule {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataWeeklySchedule) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataWeeklyScheduleParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataWeeklySchedule, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataWeeklySchedule"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataWeeklySchedule")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Virtual field
	_zero := uint64(0)
	zero := uint64(_zero)
	_ = zero

	// Optional Field (numberOfDataElements) (Can be skipped, if a given expression evaluates to false)
	var numberOfDataElements BACnetApplicationTagUnsignedInteger = nil
	if bool(bool((arrayIndexArgument) != (nil))) && bool(bool((arrayIndexArgument.GetActualValue()) == (zero))) {
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("numberOfDataElements"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for numberOfDataElements")
		}
		_val, _err := BACnetApplicationTagParse(readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			Plc4xModelLog.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'numberOfDataElements' field of BACnetConstructedDataWeeklySchedule")
		default:
			numberOfDataElements = _val.(BACnetApplicationTagUnsignedInteger)
			if closeErr := readBuffer.CloseContext("numberOfDataElements"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for numberOfDataElements")
			}
		}
	}

	// Array field (weeklySchedule)
	if pullErr := readBuffer.PullContext("weeklySchedule", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for weeklySchedule")
	}
	// Terminated array
	var weeklySchedule []BACnetDailySchedule
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetDailyScheduleParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'weeklySchedule' field of BACnetConstructedDataWeeklySchedule")
			}
			weeklySchedule = append(weeklySchedule, _item.(BACnetDailySchedule))

		}
	}
	if closeErr := readBuffer.CloseContext("weeklySchedule", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for weeklySchedule")
	}

	// Validation
	if !(bool(bool((arrayIndexArgument) != (nil))) || bool(bool((len(weeklySchedule)) == (7)))) {
		return nil, errors.WithStack(utils.ParseValidationError{"weeklySchedule should have exactly 7 values"})
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataWeeklySchedule"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataWeeklySchedule")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataWeeklySchedule{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		NumberOfDataElements: numberOfDataElements,
		WeeklySchedule:       weeklySchedule,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataWeeklySchedule) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataWeeklySchedule) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataWeeklySchedule"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataWeeklySchedule")
		}
		// Virtual field
		if _zeroErr := writeBuffer.WriteVirtual("zero", m.GetZero()); _zeroErr != nil {
			return errors.Wrap(_zeroErr, "Error serializing 'zero' field")
		}

		// Optional Field (numberOfDataElements) (Can be skipped, if the value is null)
		var numberOfDataElements BACnetApplicationTagUnsignedInteger = nil
		if m.GetNumberOfDataElements() != nil {
			if pushErr := writeBuffer.PushContext("numberOfDataElements"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for numberOfDataElements")
			}
			numberOfDataElements = m.GetNumberOfDataElements()
			_numberOfDataElementsErr := writeBuffer.WriteSerializable(numberOfDataElements)
			if popErr := writeBuffer.PopContext("numberOfDataElements"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for numberOfDataElements")
			}
			if _numberOfDataElementsErr != nil {
				return errors.Wrap(_numberOfDataElementsErr, "Error serializing 'numberOfDataElements' field")
			}
		}

		// Array Field (weeklySchedule)
		if pushErr := writeBuffer.PushContext("weeklySchedule", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for weeklySchedule")
		}
		for _, _element := range m.GetWeeklySchedule() {
			_elementErr := writeBuffer.WriteSerializable(_element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'weeklySchedule' field")
			}
		}
		if popErr := writeBuffer.PopContext("weeklySchedule", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for weeklySchedule")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataWeeklySchedule"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataWeeklySchedule")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataWeeklySchedule) isBACnetConstructedDataWeeklySchedule() bool {
	return true
}

func (m *_BACnetConstructedDataWeeklySchedule) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
