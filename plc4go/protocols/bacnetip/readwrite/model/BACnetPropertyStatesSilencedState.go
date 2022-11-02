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

// BACnetPropertyStatesSilencedState is the corresponding interface of BACnetPropertyStatesSilencedState
type BACnetPropertyStatesSilencedState interface {
	utils.LengthAware
	utils.Serializable
	BACnetPropertyStates
	// GetSilencedState returns SilencedState (property field)
	GetSilencedState() BACnetSilencedStateTagged
}

// BACnetPropertyStatesSilencedStateExactly can be used when we want exactly this type and not a type which fulfills BACnetPropertyStatesSilencedState.
// This is useful for switch cases.
type BACnetPropertyStatesSilencedStateExactly interface {
	BACnetPropertyStatesSilencedState
	isBACnetPropertyStatesSilencedState() bool
}

// _BACnetPropertyStatesSilencedState is the data-structure of this message
type _BACnetPropertyStatesSilencedState struct {
	*_BACnetPropertyStates
	SilencedState BACnetSilencedStateTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesSilencedState) InitializeParent(parent BACnetPropertyStates, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetPropertyStatesSilencedState) GetParent() BACnetPropertyStates {
	return m._BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesSilencedState) GetSilencedState() BACnetSilencedStateTagged {
	return m.SilencedState
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesSilencedState factory function for _BACnetPropertyStatesSilencedState
func NewBACnetPropertyStatesSilencedState(silencedState BACnetSilencedStateTagged, peekedTagHeader BACnetTagHeader) *_BACnetPropertyStatesSilencedState {
	_result := &_BACnetPropertyStatesSilencedState{
		SilencedState:         silencedState,
		_BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesSilencedState(structType interface{}) BACnetPropertyStatesSilencedState {
	if casted, ok := structType.(BACnetPropertyStatesSilencedState); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesSilencedState); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesSilencedState) GetTypeName() string {
	return "BACnetPropertyStatesSilencedState"
}

func (m *_BACnetPropertyStatesSilencedState) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetPropertyStatesSilencedState) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (silencedState)
	lengthInBits += m.SilencedState.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetPropertyStatesSilencedState) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPropertyStatesSilencedStateParse(readBuffer utils.ReadBuffer, peekedTagNumber uint8) (BACnetPropertyStatesSilencedState, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesSilencedState"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesSilencedState")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (silencedState)
	if pullErr := readBuffer.PullContext("silencedState"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for silencedState")
	}
	_silencedState, _silencedStateErr := BACnetSilencedStateTaggedParse(readBuffer, uint8(peekedTagNumber), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _silencedStateErr != nil {
		return nil, errors.Wrap(_silencedStateErr, "Error parsing 'silencedState' field of BACnetPropertyStatesSilencedState")
	}
	silencedState := _silencedState.(BACnetSilencedStateTagged)
	if closeErr := readBuffer.CloseContext("silencedState"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for silencedState")
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesSilencedState"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesSilencedState")
	}

	// Create a partially initialized instance
	_child := &_BACnetPropertyStatesSilencedState{
		_BACnetPropertyStates: &_BACnetPropertyStates{},
		SilencedState:         silencedState,
	}
	_child._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _child
	return _child, nil
}

func (m *_BACnetPropertyStatesSilencedState) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesSilencedState) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesSilencedState"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesSilencedState")
		}

		// Simple Field (silencedState)
		if pushErr := writeBuffer.PushContext("silencedState"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for silencedState")
		}
		_silencedStateErr := writeBuffer.WriteSerializable(m.GetSilencedState())
		if popErr := writeBuffer.PopContext("silencedState"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for silencedState")
		}
		if _silencedStateErr != nil {
			return errors.Wrap(_silencedStateErr, "Error serializing 'silencedState' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesSilencedState"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesSilencedState")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesSilencedState) isBACnetPropertyStatesSilencedState() bool {
	return true
}

func (m *_BACnetPropertyStatesSilencedState) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
