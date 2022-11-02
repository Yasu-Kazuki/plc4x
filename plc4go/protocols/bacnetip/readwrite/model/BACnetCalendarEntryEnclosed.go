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

// BACnetCalendarEntryEnclosed is the corresponding interface of BACnetCalendarEntryEnclosed
type BACnetCalendarEntryEnclosed interface {
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetCalendarEntry returns CalendarEntry (property field)
	GetCalendarEntry() BACnetCalendarEntry
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetCalendarEntryEnclosedExactly can be used when we want exactly this type and not a type which fulfills BACnetCalendarEntryEnclosed.
// This is useful for switch cases.
type BACnetCalendarEntryEnclosedExactly interface {
	BACnetCalendarEntryEnclosed
	isBACnetCalendarEntryEnclosed() bool
}

// _BACnetCalendarEntryEnclosed is the data-structure of this message
type _BACnetCalendarEntryEnclosed struct {
	OpeningTag    BACnetOpeningTag
	CalendarEntry BACnetCalendarEntry
	ClosingTag    BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetCalendarEntryEnclosed) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetCalendarEntryEnclosed) GetCalendarEntry() BACnetCalendarEntry {
	return m.CalendarEntry
}

func (m *_BACnetCalendarEntryEnclosed) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetCalendarEntryEnclosed factory function for _BACnetCalendarEntryEnclosed
func NewBACnetCalendarEntryEnclosed(openingTag BACnetOpeningTag, calendarEntry BACnetCalendarEntry, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetCalendarEntryEnclosed {
	return &_BACnetCalendarEntryEnclosed{OpeningTag: openingTag, CalendarEntry: calendarEntry, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastBACnetCalendarEntryEnclosed(structType interface{}) BACnetCalendarEntryEnclosed {
	if casted, ok := structType.(BACnetCalendarEntryEnclosed); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetCalendarEntryEnclosed); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetCalendarEntryEnclosed) GetTypeName() string {
	return "BACnetCalendarEntryEnclosed"
}

func (m *_BACnetCalendarEntryEnclosed) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetCalendarEntryEnclosed) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits()

	// Simple field (calendarEntry)
	lengthInBits += m.CalendarEntry.GetLengthInBits()

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetCalendarEntryEnclosed) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetCalendarEntryEnclosedParse(readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetCalendarEntryEnclosed, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetCalendarEntryEnclosed"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetCalendarEntryEnclosed")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParse(readBuffer, uint8(tagNumber))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field of BACnetCalendarEntryEnclosed")
	}
	openingTag := _openingTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for openingTag")
	}

	// Simple Field (calendarEntry)
	if pullErr := readBuffer.PullContext("calendarEntry"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for calendarEntry")
	}
	_calendarEntry, _calendarEntryErr := BACnetCalendarEntryParse(readBuffer)
	if _calendarEntryErr != nil {
		return nil, errors.Wrap(_calendarEntryErr, "Error parsing 'calendarEntry' field of BACnetCalendarEntryEnclosed")
	}
	calendarEntry := _calendarEntry.(BACnetCalendarEntry)
	if closeErr := readBuffer.CloseContext("calendarEntry"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for calendarEntry")
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
	_closingTag, _closingTagErr := BACnetClosingTagParse(readBuffer, uint8(tagNumber))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field of BACnetCalendarEntryEnclosed")
	}
	closingTag := _closingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetCalendarEntryEnclosed"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetCalendarEntryEnclosed")
	}

	// Create the instance
	return &_BACnetCalendarEntryEnclosed{
		TagNumber:     tagNumber,
		OpeningTag:    openingTag,
		CalendarEntry: calendarEntry,
		ClosingTag:    closingTag,
	}, nil
}

func (m *_BACnetCalendarEntryEnclosed) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetCalendarEntryEnclosed) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetCalendarEntryEnclosed"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetCalendarEntryEnclosed")
	}

	// Simple Field (openingTag)
	if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for openingTag")
	}
	_openingTagErr := writeBuffer.WriteSerializable(m.GetOpeningTag())
	if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for openingTag")
	}
	if _openingTagErr != nil {
		return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
	}

	// Simple Field (calendarEntry)
	if pushErr := writeBuffer.PushContext("calendarEntry"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for calendarEntry")
	}
	_calendarEntryErr := writeBuffer.WriteSerializable(m.GetCalendarEntry())
	if popErr := writeBuffer.PopContext("calendarEntry"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for calendarEntry")
	}
	if _calendarEntryErr != nil {
		return errors.Wrap(_calendarEntryErr, "Error serializing 'calendarEntry' field")
	}

	// Simple Field (closingTag)
	if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for closingTag")
	}
	_closingTagErr := writeBuffer.WriteSerializable(m.GetClosingTag())
	if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for closingTag")
	}
	if _closingTagErr != nil {
		return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetCalendarEntryEnclosed"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetCalendarEntryEnclosed")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetCalendarEntryEnclosed) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetCalendarEntryEnclosed) isBACnetCalendarEntryEnclosed() bool {
	return true
}

func (m *_BACnetCalendarEntryEnclosed) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
