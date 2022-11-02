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

// BACnetConstructedDataEventState is the corresponding interface of BACnetConstructedDataEventState
type BACnetConstructedDataEventState interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetEventState returns EventState (property field)
	GetEventState() BACnetEventStateTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetEventStateTagged
}

// BACnetConstructedDataEventStateExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataEventState.
// This is useful for switch cases.
type BACnetConstructedDataEventStateExactly interface {
	BACnetConstructedDataEventState
	isBACnetConstructedDataEventState() bool
}

// _BACnetConstructedDataEventState is the data-structure of this message
type _BACnetConstructedDataEventState struct {
	*_BACnetConstructedData
	EventState BACnetEventStateTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataEventState) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataEventState) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_EVENT_STATE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataEventState) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataEventState) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataEventState) GetEventState() BACnetEventStateTagged {
	return m.EventState
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataEventState) GetActualValue() BACnetEventStateTagged {
	return CastBACnetEventStateTagged(m.GetEventState())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataEventState factory function for _BACnetConstructedDataEventState
func NewBACnetConstructedDataEventState(eventState BACnetEventStateTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataEventState {
	_result := &_BACnetConstructedDataEventState{
		EventState:             eventState,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataEventState(structType interface{}) BACnetConstructedDataEventState {
	if casted, ok := structType.(BACnetConstructedDataEventState); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataEventState); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataEventState) GetTypeName() string {
	return "BACnetConstructedDataEventState"
}

func (m *_BACnetConstructedDataEventState) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataEventState) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (eventState)
	lengthInBits += m.EventState.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataEventState) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataEventStateParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataEventState, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataEventState"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataEventState")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (eventState)
	if pullErr := readBuffer.PullContext("eventState"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for eventState")
	}
	_eventState, _eventStateErr := BACnetEventStateTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _eventStateErr != nil {
		return nil, errors.Wrap(_eventStateErr, "Error parsing 'eventState' field of BACnetConstructedDataEventState")
	}
	eventState := _eventState.(BACnetEventStateTagged)
	if closeErr := readBuffer.CloseContext("eventState"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for eventState")
	}

	// Virtual field
	_actualValue := eventState
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataEventState"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataEventState")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataEventState{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		EventState: eventState,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataEventState) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataEventState) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataEventState"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataEventState")
		}

		// Simple Field (eventState)
		if pushErr := writeBuffer.PushContext("eventState"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for eventState")
		}
		_eventStateErr := writeBuffer.WriteSerializable(m.GetEventState())
		if popErr := writeBuffer.PopContext("eventState"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for eventState")
		}
		if _eventStateErr != nil {
			return errors.Wrap(_eventStateErr, "Error serializing 'eventState' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataEventState"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataEventState")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataEventState) isBACnetConstructedDataEventState() bool {
	return true
}

func (m *_BACnetConstructedDataEventState) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
