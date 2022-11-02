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

// BACnetFaultParameterFaultExtendedParametersEntryDate is the corresponding interface of BACnetFaultParameterFaultExtendedParametersEntryDate
type BACnetFaultParameterFaultExtendedParametersEntryDate interface {
	utils.LengthAware
	utils.Serializable
	BACnetFaultParameterFaultExtendedParametersEntry
	// GetDateValue returns DateValue (property field)
	GetDateValue() BACnetApplicationTagDate
}

// BACnetFaultParameterFaultExtendedParametersEntryDateExactly can be used when we want exactly this type and not a type which fulfills BACnetFaultParameterFaultExtendedParametersEntryDate.
// This is useful for switch cases.
type BACnetFaultParameterFaultExtendedParametersEntryDateExactly interface {
	BACnetFaultParameterFaultExtendedParametersEntryDate
	isBACnetFaultParameterFaultExtendedParametersEntryDate() bool
}

// _BACnetFaultParameterFaultExtendedParametersEntryDate is the data-structure of this message
type _BACnetFaultParameterFaultExtendedParametersEntryDate struct {
	*_BACnetFaultParameterFaultExtendedParametersEntry
	DateValue BACnetApplicationTagDate
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetFaultParameterFaultExtendedParametersEntryDate) InitializeParent(parent BACnetFaultParameterFaultExtendedParametersEntry, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryDate) GetParent() BACnetFaultParameterFaultExtendedParametersEntry {
	return m._BACnetFaultParameterFaultExtendedParametersEntry
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFaultParameterFaultExtendedParametersEntryDate) GetDateValue() BACnetApplicationTagDate {
	return m.DateValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetFaultParameterFaultExtendedParametersEntryDate factory function for _BACnetFaultParameterFaultExtendedParametersEntryDate
func NewBACnetFaultParameterFaultExtendedParametersEntryDate(dateValue BACnetApplicationTagDate, peekedTagHeader BACnetTagHeader) *_BACnetFaultParameterFaultExtendedParametersEntryDate {
	_result := &_BACnetFaultParameterFaultExtendedParametersEntryDate{
		DateValue: dateValue,
		_BACnetFaultParameterFaultExtendedParametersEntry: NewBACnetFaultParameterFaultExtendedParametersEntry(peekedTagHeader),
	}
	_result._BACnetFaultParameterFaultExtendedParametersEntry._BACnetFaultParameterFaultExtendedParametersEntryChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetFaultParameterFaultExtendedParametersEntryDate(structType interface{}) BACnetFaultParameterFaultExtendedParametersEntryDate {
	if casted, ok := structType.(BACnetFaultParameterFaultExtendedParametersEntryDate); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultExtendedParametersEntryDate); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryDate) GetTypeName() string {
	return "BACnetFaultParameterFaultExtendedParametersEntryDate"
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryDate) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryDate) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (dateValue)
	lengthInBits += m.DateValue.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryDate) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetFaultParameterFaultExtendedParametersEntryDateParse(readBuffer utils.ReadBuffer) (BACnetFaultParameterFaultExtendedParametersEntryDate, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultExtendedParametersEntryDate"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultExtendedParametersEntryDate")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (dateValue)
	if pullErr := readBuffer.PullContext("dateValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for dateValue")
	}
	_dateValue, _dateValueErr := BACnetApplicationTagParse(readBuffer)
	if _dateValueErr != nil {
		return nil, errors.Wrap(_dateValueErr, "Error parsing 'dateValue' field of BACnetFaultParameterFaultExtendedParametersEntryDate")
	}
	dateValue := _dateValue.(BACnetApplicationTagDate)
	if closeErr := readBuffer.CloseContext("dateValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for dateValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultExtendedParametersEntryDate"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultExtendedParametersEntryDate")
	}

	// Create a partially initialized instance
	_child := &_BACnetFaultParameterFaultExtendedParametersEntryDate{
		_BACnetFaultParameterFaultExtendedParametersEntry: &_BACnetFaultParameterFaultExtendedParametersEntry{},
		DateValue: dateValue,
	}
	_child._BACnetFaultParameterFaultExtendedParametersEntry._BACnetFaultParameterFaultExtendedParametersEntryChildRequirements = _child
	return _child, nil
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryDate) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryDate) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultExtendedParametersEntryDate"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultExtendedParametersEntryDate")
		}

		// Simple Field (dateValue)
		if pushErr := writeBuffer.PushContext("dateValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for dateValue")
		}
		_dateValueErr := writeBuffer.WriteSerializable(m.GetDateValue())
		if popErr := writeBuffer.PopContext("dateValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for dateValue")
		}
		if _dateValueErr != nil {
			return errors.Wrap(_dateValueErr, "Error serializing 'dateValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultExtendedParametersEntryDate"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultExtendedParametersEntryDate")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryDate) isBACnetFaultParameterFaultExtendedParametersEntryDate() bool {
	return true
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryDate) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
