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

// BACnetCalendarEntryDateRange is the corresponding interface of BACnetCalendarEntryDateRange
type BACnetCalendarEntryDateRange interface {
	utils.LengthAware
	utils.Serializable
	BACnetCalendarEntry
	// GetDateRange returns DateRange (property field)
	GetDateRange() BACnetDateRangeEnclosed
}

// BACnetCalendarEntryDateRangeExactly can be used when we want exactly this type and not a type which fulfills BACnetCalendarEntryDateRange.
// This is useful for switch cases.
type BACnetCalendarEntryDateRangeExactly interface {
	BACnetCalendarEntryDateRange
	isBACnetCalendarEntryDateRange() bool
}

// _BACnetCalendarEntryDateRange is the data-structure of this message
type _BACnetCalendarEntryDateRange struct {
	*_BACnetCalendarEntry
	DateRange BACnetDateRangeEnclosed
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetCalendarEntryDateRange) InitializeParent(parent BACnetCalendarEntry, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetCalendarEntryDateRange) GetParent() BACnetCalendarEntry {
	return m._BACnetCalendarEntry
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetCalendarEntryDateRange) GetDateRange() BACnetDateRangeEnclosed {
	return m.DateRange
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetCalendarEntryDateRange factory function for _BACnetCalendarEntryDateRange
func NewBACnetCalendarEntryDateRange(dateRange BACnetDateRangeEnclosed, peekedTagHeader BACnetTagHeader) *_BACnetCalendarEntryDateRange {
	_result := &_BACnetCalendarEntryDateRange{
		DateRange:            dateRange,
		_BACnetCalendarEntry: NewBACnetCalendarEntry(peekedTagHeader),
	}
	_result._BACnetCalendarEntry._BACnetCalendarEntryChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetCalendarEntryDateRange(structType interface{}) BACnetCalendarEntryDateRange {
	if casted, ok := structType.(BACnetCalendarEntryDateRange); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetCalendarEntryDateRange); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetCalendarEntryDateRange) GetTypeName() string {
	return "BACnetCalendarEntryDateRange"
}

func (m *_BACnetCalendarEntryDateRange) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetCalendarEntryDateRange) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (dateRange)
	lengthInBits += m.DateRange.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetCalendarEntryDateRange) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetCalendarEntryDateRangeParse(readBuffer utils.ReadBuffer) (BACnetCalendarEntryDateRange, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetCalendarEntryDateRange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetCalendarEntryDateRange")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (dateRange)
	if pullErr := readBuffer.PullContext("dateRange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for dateRange")
	}
	_dateRange, _dateRangeErr := BACnetDateRangeEnclosedParse(readBuffer, uint8(uint8(1)))
	if _dateRangeErr != nil {
		return nil, errors.Wrap(_dateRangeErr, "Error parsing 'dateRange' field of BACnetCalendarEntryDateRange")
	}
	dateRange := _dateRange.(BACnetDateRangeEnclosed)
	if closeErr := readBuffer.CloseContext("dateRange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for dateRange")
	}

	if closeErr := readBuffer.CloseContext("BACnetCalendarEntryDateRange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetCalendarEntryDateRange")
	}

	// Create a partially initialized instance
	_child := &_BACnetCalendarEntryDateRange{
		_BACnetCalendarEntry: &_BACnetCalendarEntry{},
		DateRange:            dateRange,
	}
	_child._BACnetCalendarEntry._BACnetCalendarEntryChildRequirements = _child
	return _child, nil
}

func (m *_BACnetCalendarEntryDateRange) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetCalendarEntryDateRange) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetCalendarEntryDateRange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetCalendarEntryDateRange")
		}

		// Simple Field (dateRange)
		if pushErr := writeBuffer.PushContext("dateRange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for dateRange")
		}
		_dateRangeErr := writeBuffer.WriteSerializable(m.GetDateRange())
		if popErr := writeBuffer.PopContext("dateRange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for dateRange")
		}
		if _dateRangeErr != nil {
			return errors.Wrap(_dateRangeErr, "Error serializing 'dateRange' field")
		}

		if popErr := writeBuffer.PopContext("BACnetCalendarEntryDateRange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetCalendarEntryDateRange")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetCalendarEntryDateRange) isBACnetCalendarEntryDateRange() bool {
	return true
}

func (m *_BACnetCalendarEntryDateRange) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
