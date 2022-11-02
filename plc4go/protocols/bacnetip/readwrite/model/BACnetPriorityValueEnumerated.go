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

// BACnetPriorityValueEnumerated is the corresponding interface of BACnetPriorityValueEnumerated
type BACnetPriorityValueEnumerated interface {
	utils.LengthAware
	utils.Serializable
	BACnetPriorityValue
	// GetEnumeratedValue returns EnumeratedValue (property field)
	GetEnumeratedValue() BACnetApplicationTagEnumerated
}

// BACnetPriorityValueEnumeratedExactly can be used when we want exactly this type and not a type which fulfills BACnetPriorityValueEnumerated.
// This is useful for switch cases.
type BACnetPriorityValueEnumeratedExactly interface {
	BACnetPriorityValueEnumerated
	isBACnetPriorityValueEnumerated() bool
}

// _BACnetPriorityValueEnumerated is the data-structure of this message
type _BACnetPriorityValueEnumerated struct {
	*_BACnetPriorityValue
	EnumeratedValue BACnetApplicationTagEnumerated
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPriorityValueEnumerated) InitializeParent(parent BACnetPriorityValue, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetPriorityValueEnumerated) GetParent() BACnetPriorityValue {
	return m._BACnetPriorityValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPriorityValueEnumerated) GetEnumeratedValue() BACnetApplicationTagEnumerated {
	return m.EnumeratedValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPriorityValueEnumerated factory function for _BACnetPriorityValueEnumerated
func NewBACnetPriorityValueEnumerated(enumeratedValue BACnetApplicationTagEnumerated, peekedTagHeader BACnetTagHeader, objectTypeArgument BACnetObjectType) *_BACnetPriorityValueEnumerated {
	_result := &_BACnetPriorityValueEnumerated{
		EnumeratedValue:      enumeratedValue,
		_BACnetPriorityValue: NewBACnetPriorityValue(peekedTagHeader, objectTypeArgument),
	}
	_result._BACnetPriorityValue._BACnetPriorityValueChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPriorityValueEnumerated(structType interface{}) BACnetPriorityValueEnumerated {
	if casted, ok := structType.(BACnetPriorityValueEnumerated); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPriorityValueEnumerated); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPriorityValueEnumerated) GetTypeName() string {
	return "BACnetPriorityValueEnumerated"
}

func (m *_BACnetPriorityValueEnumerated) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetPriorityValueEnumerated) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (enumeratedValue)
	lengthInBits += m.EnumeratedValue.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetPriorityValueEnumerated) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPriorityValueEnumeratedParse(readBuffer utils.ReadBuffer, objectTypeArgument BACnetObjectType) (BACnetPriorityValueEnumerated, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPriorityValueEnumerated"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPriorityValueEnumerated")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (enumeratedValue)
	if pullErr := readBuffer.PullContext("enumeratedValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for enumeratedValue")
	}
	_enumeratedValue, _enumeratedValueErr := BACnetApplicationTagParse(readBuffer)
	if _enumeratedValueErr != nil {
		return nil, errors.Wrap(_enumeratedValueErr, "Error parsing 'enumeratedValue' field of BACnetPriorityValueEnumerated")
	}
	enumeratedValue := _enumeratedValue.(BACnetApplicationTagEnumerated)
	if closeErr := readBuffer.CloseContext("enumeratedValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for enumeratedValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetPriorityValueEnumerated"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPriorityValueEnumerated")
	}

	// Create a partially initialized instance
	_child := &_BACnetPriorityValueEnumerated{
		_BACnetPriorityValue: &_BACnetPriorityValue{
			ObjectTypeArgument: objectTypeArgument,
		},
		EnumeratedValue: enumeratedValue,
	}
	_child._BACnetPriorityValue._BACnetPriorityValueChildRequirements = _child
	return _child, nil
}

func (m *_BACnetPriorityValueEnumerated) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPriorityValueEnumerated) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPriorityValueEnumerated"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPriorityValueEnumerated")
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

		if popErr := writeBuffer.PopContext("BACnetPriorityValueEnumerated"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPriorityValueEnumerated")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetPriorityValueEnumerated) isBACnetPriorityValueEnumerated() bool {
	return true
}

func (m *_BACnetPriorityValueEnumerated) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
