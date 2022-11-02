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

// BACnetFaultParameterFaultStateListOfFaultValues is the corresponding interface of BACnetFaultParameterFaultStateListOfFaultValues
type BACnetFaultParameterFaultStateListOfFaultValues interface {
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetListIfFaultValues returns ListIfFaultValues (property field)
	GetListIfFaultValues() []BACnetPropertyStates
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetFaultParameterFaultStateListOfFaultValuesExactly can be used when we want exactly this type and not a type which fulfills BACnetFaultParameterFaultStateListOfFaultValues.
// This is useful for switch cases.
type BACnetFaultParameterFaultStateListOfFaultValuesExactly interface {
	BACnetFaultParameterFaultStateListOfFaultValues
	isBACnetFaultParameterFaultStateListOfFaultValues() bool
}

// _BACnetFaultParameterFaultStateListOfFaultValues is the data-structure of this message
type _BACnetFaultParameterFaultStateListOfFaultValues struct {
	OpeningTag        BACnetOpeningTag
	ListIfFaultValues []BACnetPropertyStates
	ClosingTag        BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFaultParameterFaultStateListOfFaultValues) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetFaultParameterFaultStateListOfFaultValues) GetListIfFaultValues() []BACnetPropertyStates {
	return m.ListIfFaultValues
}

func (m *_BACnetFaultParameterFaultStateListOfFaultValues) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetFaultParameterFaultStateListOfFaultValues factory function for _BACnetFaultParameterFaultStateListOfFaultValues
func NewBACnetFaultParameterFaultStateListOfFaultValues(openingTag BACnetOpeningTag, listIfFaultValues []BACnetPropertyStates, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetFaultParameterFaultStateListOfFaultValues {
	return &_BACnetFaultParameterFaultStateListOfFaultValues{OpeningTag: openingTag, ListIfFaultValues: listIfFaultValues, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastBACnetFaultParameterFaultStateListOfFaultValues(structType interface{}) BACnetFaultParameterFaultStateListOfFaultValues {
	if casted, ok := structType.(BACnetFaultParameterFaultStateListOfFaultValues); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultStateListOfFaultValues); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFaultParameterFaultStateListOfFaultValues) GetTypeName() string {
	return "BACnetFaultParameterFaultStateListOfFaultValues"
}

func (m *_BACnetFaultParameterFaultStateListOfFaultValues) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetFaultParameterFaultStateListOfFaultValues) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits()

	// Array field
	if len(m.ListIfFaultValues) > 0 {
		for _, element := range m.ListIfFaultValues {
			lengthInBits += element.GetLengthInBits()
		}
	}

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetFaultParameterFaultStateListOfFaultValues) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetFaultParameterFaultStateListOfFaultValuesParse(readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetFaultParameterFaultStateListOfFaultValues, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultStateListOfFaultValues"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultStateListOfFaultValues")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParse(readBuffer, uint8(tagNumber))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field of BACnetFaultParameterFaultStateListOfFaultValues")
	}
	openingTag := _openingTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for openingTag")
	}

	// Array field (listIfFaultValues)
	if pullErr := readBuffer.PullContext("listIfFaultValues", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for listIfFaultValues")
	}
	// Terminated array
	var listIfFaultValues []BACnetPropertyStates
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetPropertyStatesParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'listIfFaultValues' field of BACnetFaultParameterFaultStateListOfFaultValues")
			}
			listIfFaultValues = append(listIfFaultValues, _item.(BACnetPropertyStates))

		}
	}
	if closeErr := readBuffer.CloseContext("listIfFaultValues", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for listIfFaultValues")
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
	_closingTag, _closingTagErr := BACnetClosingTagParse(readBuffer, uint8(tagNumber))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field of BACnetFaultParameterFaultStateListOfFaultValues")
	}
	closingTag := _closingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultStateListOfFaultValues"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultStateListOfFaultValues")
	}

	// Create the instance
	return &_BACnetFaultParameterFaultStateListOfFaultValues{
		TagNumber:         tagNumber,
		OpeningTag:        openingTag,
		ListIfFaultValues: listIfFaultValues,
		ClosingTag:        closingTag,
	}, nil
}

func (m *_BACnetFaultParameterFaultStateListOfFaultValues) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetFaultParameterFaultStateListOfFaultValues) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultStateListOfFaultValues"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultStateListOfFaultValues")
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

	// Array Field (listIfFaultValues)
	if pushErr := writeBuffer.PushContext("listIfFaultValues", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for listIfFaultValues")
	}
	for _, _element := range m.GetListIfFaultValues() {
		_elementErr := writeBuffer.WriteSerializable(_element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'listIfFaultValues' field")
		}
	}
	if popErr := writeBuffer.PopContext("listIfFaultValues", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for listIfFaultValues")
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

	if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultStateListOfFaultValues"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultStateListOfFaultValues")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetFaultParameterFaultStateListOfFaultValues) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetFaultParameterFaultStateListOfFaultValues) isBACnetFaultParameterFaultStateListOfFaultValues() bool {
	return true
}

func (m *_BACnetFaultParameterFaultStateListOfFaultValues) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
