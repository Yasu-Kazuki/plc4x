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

// BACnetChannelValueObjectidentifier is the corresponding interface of BACnetChannelValueObjectidentifier
type BACnetChannelValueObjectidentifier interface {
	utils.LengthAware
	utils.Serializable
	BACnetChannelValue
	// GetObjectidentifierValue returns ObjectidentifierValue (property field)
	GetObjectidentifierValue() BACnetApplicationTagObjectIdentifier
}

// BACnetChannelValueObjectidentifierExactly can be used when we want exactly this type and not a type which fulfills BACnetChannelValueObjectidentifier.
// This is useful for switch cases.
type BACnetChannelValueObjectidentifierExactly interface {
	BACnetChannelValueObjectidentifier
	isBACnetChannelValueObjectidentifier() bool
}

// _BACnetChannelValueObjectidentifier is the data-structure of this message
type _BACnetChannelValueObjectidentifier struct {
	*_BACnetChannelValue
	ObjectidentifierValue BACnetApplicationTagObjectIdentifier
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetChannelValueObjectidentifier) InitializeParent(parent BACnetChannelValue, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetChannelValueObjectidentifier) GetParent() BACnetChannelValue {
	return m._BACnetChannelValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetChannelValueObjectidentifier) GetObjectidentifierValue() BACnetApplicationTagObjectIdentifier {
	return m.ObjectidentifierValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetChannelValueObjectidentifier factory function for _BACnetChannelValueObjectidentifier
func NewBACnetChannelValueObjectidentifier(objectidentifierValue BACnetApplicationTagObjectIdentifier, peekedTagHeader BACnetTagHeader) *_BACnetChannelValueObjectidentifier {
	_result := &_BACnetChannelValueObjectidentifier{
		ObjectidentifierValue: objectidentifierValue,
		_BACnetChannelValue:   NewBACnetChannelValue(peekedTagHeader),
	}
	_result._BACnetChannelValue._BACnetChannelValueChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetChannelValueObjectidentifier(structType interface{}) BACnetChannelValueObjectidentifier {
	if casted, ok := structType.(BACnetChannelValueObjectidentifier); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetChannelValueObjectidentifier); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetChannelValueObjectidentifier) GetTypeName() string {
	return "BACnetChannelValueObjectidentifier"
}

func (m *_BACnetChannelValueObjectidentifier) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetChannelValueObjectidentifier) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (objectidentifierValue)
	lengthInBits += m.ObjectidentifierValue.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetChannelValueObjectidentifier) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetChannelValueObjectidentifierParse(theBytes []byte) (BACnetChannelValueObjectidentifier, error) {
	return BACnetChannelValueObjectidentifierParseWithBuffer(utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian))) // TODO: get endianness from mspec
}

func BACnetChannelValueObjectidentifierParseWithBuffer(readBuffer utils.ReadBuffer) (BACnetChannelValueObjectidentifier, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetChannelValueObjectidentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetChannelValueObjectidentifier")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (objectidentifierValue)
	if pullErr := readBuffer.PullContext("objectidentifierValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for objectidentifierValue")
	}
	_objectidentifierValue, _objectidentifierValueErr := BACnetApplicationTagParseWithBuffer(readBuffer)
	if _objectidentifierValueErr != nil {
		return nil, errors.Wrap(_objectidentifierValueErr, "Error parsing 'objectidentifierValue' field of BACnetChannelValueObjectidentifier")
	}
	objectidentifierValue := _objectidentifierValue.(BACnetApplicationTagObjectIdentifier)
	if closeErr := readBuffer.CloseContext("objectidentifierValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for objectidentifierValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetChannelValueObjectidentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetChannelValueObjectidentifier")
	}

	// Create a partially initialized instance
	_child := &_BACnetChannelValueObjectidentifier{
		_BACnetChannelValue:   &_BACnetChannelValue{},
		ObjectidentifierValue: objectidentifierValue,
	}
	_child._BACnetChannelValue._BACnetChannelValueChildRequirements = _child
	return _child, nil
}

func (m *_BACnetChannelValueObjectidentifier) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetChannelValueObjectidentifier) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetChannelValueObjectidentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetChannelValueObjectidentifier")
		}

		// Simple Field (objectidentifierValue)
		if pushErr := writeBuffer.PushContext("objectidentifierValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for objectidentifierValue")
		}
		_objectidentifierValueErr := writeBuffer.WriteSerializable(m.GetObjectidentifierValue())
		if popErr := writeBuffer.PopContext("objectidentifierValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for objectidentifierValue")
		}
		if _objectidentifierValueErr != nil {
			return errors.Wrap(_objectidentifierValueErr, "Error serializing 'objectidentifierValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetChannelValueObjectidentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetChannelValueObjectidentifier")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetChannelValueObjectidentifier) isBACnetChannelValueObjectidentifier() bool {
	return true
}

func (m *_BACnetChannelValueObjectidentifier) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}