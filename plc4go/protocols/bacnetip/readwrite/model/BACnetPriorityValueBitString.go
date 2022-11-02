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

// BACnetPriorityValueBitString is the corresponding interface of BACnetPriorityValueBitString
type BACnetPriorityValueBitString interface {
	utils.LengthAware
	utils.Serializable
	BACnetPriorityValue
	// GetBitStringValue returns BitStringValue (property field)
	GetBitStringValue() BACnetApplicationTagBitString
}

// BACnetPriorityValueBitStringExactly can be used when we want exactly this type and not a type which fulfills BACnetPriorityValueBitString.
// This is useful for switch cases.
type BACnetPriorityValueBitStringExactly interface {
	BACnetPriorityValueBitString
	isBACnetPriorityValueBitString() bool
}

// _BACnetPriorityValueBitString is the data-structure of this message
type _BACnetPriorityValueBitString struct {
	*_BACnetPriorityValue
	BitStringValue BACnetApplicationTagBitString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPriorityValueBitString) InitializeParent(parent BACnetPriorityValue, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetPriorityValueBitString) GetParent() BACnetPriorityValue {
	return m._BACnetPriorityValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPriorityValueBitString) GetBitStringValue() BACnetApplicationTagBitString {
	return m.BitStringValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPriorityValueBitString factory function for _BACnetPriorityValueBitString
func NewBACnetPriorityValueBitString(bitStringValue BACnetApplicationTagBitString, peekedTagHeader BACnetTagHeader, objectTypeArgument BACnetObjectType) *_BACnetPriorityValueBitString {
	_result := &_BACnetPriorityValueBitString{
		BitStringValue:       bitStringValue,
		_BACnetPriorityValue: NewBACnetPriorityValue(peekedTagHeader, objectTypeArgument),
	}
	_result._BACnetPriorityValue._BACnetPriorityValueChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPriorityValueBitString(structType interface{}) BACnetPriorityValueBitString {
	if casted, ok := structType.(BACnetPriorityValueBitString); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPriorityValueBitString); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPriorityValueBitString) GetTypeName() string {
	return "BACnetPriorityValueBitString"
}

func (m *_BACnetPriorityValueBitString) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetPriorityValueBitString) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (bitStringValue)
	lengthInBits += m.BitStringValue.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetPriorityValueBitString) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPriorityValueBitStringParse(readBuffer utils.ReadBuffer, objectTypeArgument BACnetObjectType) (BACnetPriorityValueBitString, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPriorityValueBitString"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPriorityValueBitString")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (bitStringValue)
	if pullErr := readBuffer.PullContext("bitStringValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for bitStringValue")
	}
	_bitStringValue, _bitStringValueErr := BACnetApplicationTagParse(readBuffer)
	if _bitStringValueErr != nil {
		return nil, errors.Wrap(_bitStringValueErr, "Error parsing 'bitStringValue' field of BACnetPriorityValueBitString")
	}
	bitStringValue := _bitStringValue.(BACnetApplicationTagBitString)
	if closeErr := readBuffer.CloseContext("bitStringValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for bitStringValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetPriorityValueBitString"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPriorityValueBitString")
	}

	// Create a partially initialized instance
	_child := &_BACnetPriorityValueBitString{
		_BACnetPriorityValue: &_BACnetPriorityValue{
			ObjectTypeArgument: objectTypeArgument,
		},
		BitStringValue: bitStringValue,
	}
	_child._BACnetPriorityValue._BACnetPriorityValueChildRequirements = _child
	return _child, nil
}

func (m *_BACnetPriorityValueBitString) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPriorityValueBitString) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPriorityValueBitString"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPriorityValueBitString")
		}

		// Simple Field (bitStringValue)
		if pushErr := writeBuffer.PushContext("bitStringValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for bitStringValue")
		}
		_bitStringValueErr := writeBuffer.WriteSerializable(m.GetBitStringValue())
		if popErr := writeBuffer.PopContext("bitStringValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for bitStringValue")
		}
		if _bitStringValueErr != nil {
			return errors.Wrap(_bitStringValueErr, "Error serializing 'bitStringValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPriorityValueBitString"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPriorityValueBitString")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetPriorityValueBitString) isBACnetPriorityValueBitString() bool {
	return true
}

func (m *_BACnetPriorityValueBitString) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
