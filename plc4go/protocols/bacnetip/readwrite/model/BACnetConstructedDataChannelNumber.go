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

// BACnetConstructedDataChannelNumber is the corresponding interface of BACnetConstructedDataChannelNumber
type BACnetConstructedDataChannelNumber interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetChannelNumber returns ChannelNumber (property field)
	GetChannelNumber() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataChannelNumberExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataChannelNumber.
// This is useful for switch cases.
type BACnetConstructedDataChannelNumberExactly interface {
	BACnetConstructedDataChannelNumber
	isBACnetConstructedDataChannelNumber() bool
}

// _BACnetConstructedDataChannelNumber is the data-structure of this message
type _BACnetConstructedDataChannelNumber struct {
	*_BACnetConstructedData
	ChannelNumber BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataChannelNumber) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataChannelNumber) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_CHANNEL_NUMBER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataChannelNumber) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataChannelNumber) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataChannelNumber) GetChannelNumber() BACnetApplicationTagUnsignedInteger {
	return m.ChannelNumber
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataChannelNumber) GetActualValue() BACnetApplicationTagUnsignedInteger {
	return CastBACnetApplicationTagUnsignedInteger(m.GetChannelNumber())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataChannelNumber factory function for _BACnetConstructedDataChannelNumber
func NewBACnetConstructedDataChannelNumber(channelNumber BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataChannelNumber {
	_result := &_BACnetConstructedDataChannelNumber{
		ChannelNumber:          channelNumber,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataChannelNumber(structType interface{}) BACnetConstructedDataChannelNumber {
	if casted, ok := structType.(BACnetConstructedDataChannelNumber); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataChannelNumber); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataChannelNumber) GetTypeName() string {
	return "BACnetConstructedDataChannelNumber"
}

func (m *_BACnetConstructedDataChannelNumber) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataChannelNumber) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (channelNumber)
	lengthInBits += m.ChannelNumber.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataChannelNumber) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataChannelNumberParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataChannelNumber, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataChannelNumber"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataChannelNumber")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (channelNumber)
	if pullErr := readBuffer.PullContext("channelNumber"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for channelNumber")
	}
	_channelNumber, _channelNumberErr := BACnetApplicationTagParse(readBuffer)
	if _channelNumberErr != nil {
		return nil, errors.Wrap(_channelNumberErr, "Error parsing 'channelNumber' field of BACnetConstructedDataChannelNumber")
	}
	channelNumber := _channelNumber.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("channelNumber"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for channelNumber")
	}

	// Virtual field
	_actualValue := channelNumber
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataChannelNumber"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataChannelNumber")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataChannelNumber{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		ChannelNumber: channelNumber,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataChannelNumber) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataChannelNumber) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataChannelNumber"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataChannelNumber")
		}

		// Simple Field (channelNumber)
		if pushErr := writeBuffer.PushContext("channelNumber"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for channelNumber")
		}
		_channelNumberErr := writeBuffer.WriteSerializable(m.GetChannelNumber())
		if popErr := writeBuffer.PopContext("channelNumber"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for channelNumber")
		}
		if _channelNumberErr != nil {
			return errors.Wrap(_channelNumberErr, "Error serializing 'channelNumber' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataChannelNumber"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataChannelNumber")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataChannelNumber) isBACnetConstructedDataChannelNumber() bool {
	return true
}

func (m *_BACnetConstructedDataChannelNumber) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
