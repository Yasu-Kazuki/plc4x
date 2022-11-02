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

// BACnetConstructedDataChangeOfStateCount is the corresponding interface of BACnetConstructedDataChangeOfStateCount
type BACnetConstructedDataChangeOfStateCount interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetChangeIfStateCount returns ChangeIfStateCount (property field)
	GetChangeIfStateCount() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataChangeOfStateCountExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataChangeOfStateCount.
// This is useful for switch cases.
type BACnetConstructedDataChangeOfStateCountExactly interface {
	BACnetConstructedDataChangeOfStateCount
	isBACnetConstructedDataChangeOfStateCount() bool
}

// _BACnetConstructedDataChangeOfStateCount is the data-structure of this message
type _BACnetConstructedDataChangeOfStateCount struct {
	*_BACnetConstructedData
	ChangeIfStateCount BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataChangeOfStateCount) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataChangeOfStateCount) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_CHANGE_OF_STATE_COUNT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataChangeOfStateCount) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataChangeOfStateCount) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataChangeOfStateCount) GetChangeIfStateCount() BACnetApplicationTagUnsignedInteger {
	return m.ChangeIfStateCount
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataChangeOfStateCount) GetActualValue() BACnetApplicationTagUnsignedInteger {
	return CastBACnetApplicationTagUnsignedInteger(m.GetChangeIfStateCount())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataChangeOfStateCount factory function for _BACnetConstructedDataChangeOfStateCount
func NewBACnetConstructedDataChangeOfStateCount(changeIfStateCount BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataChangeOfStateCount {
	_result := &_BACnetConstructedDataChangeOfStateCount{
		ChangeIfStateCount:     changeIfStateCount,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataChangeOfStateCount(structType interface{}) BACnetConstructedDataChangeOfStateCount {
	if casted, ok := structType.(BACnetConstructedDataChangeOfStateCount); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataChangeOfStateCount); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataChangeOfStateCount) GetTypeName() string {
	return "BACnetConstructedDataChangeOfStateCount"
}

func (m *_BACnetConstructedDataChangeOfStateCount) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataChangeOfStateCount) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (changeIfStateCount)
	lengthInBits += m.ChangeIfStateCount.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataChangeOfStateCount) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataChangeOfStateCountParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataChangeOfStateCount, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataChangeOfStateCount"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataChangeOfStateCount")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (changeIfStateCount)
	if pullErr := readBuffer.PullContext("changeIfStateCount"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for changeIfStateCount")
	}
	_changeIfStateCount, _changeIfStateCountErr := BACnetApplicationTagParse(readBuffer)
	if _changeIfStateCountErr != nil {
		return nil, errors.Wrap(_changeIfStateCountErr, "Error parsing 'changeIfStateCount' field of BACnetConstructedDataChangeOfStateCount")
	}
	changeIfStateCount := _changeIfStateCount.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("changeIfStateCount"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for changeIfStateCount")
	}

	// Virtual field
	_actualValue := changeIfStateCount
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataChangeOfStateCount"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataChangeOfStateCount")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataChangeOfStateCount{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		ChangeIfStateCount: changeIfStateCount,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataChangeOfStateCount) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataChangeOfStateCount) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataChangeOfStateCount"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataChangeOfStateCount")
		}

		// Simple Field (changeIfStateCount)
		if pushErr := writeBuffer.PushContext("changeIfStateCount"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for changeIfStateCount")
		}
		_changeIfStateCountErr := writeBuffer.WriteSerializable(m.GetChangeIfStateCount())
		if popErr := writeBuffer.PopContext("changeIfStateCount"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for changeIfStateCount")
		}
		if _changeIfStateCountErr != nil {
			return errors.Wrap(_changeIfStateCountErr, "Error serializing 'changeIfStateCount' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataChangeOfStateCount"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataChangeOfStateCount")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataChangeOfStateCount) isBACnetConstructedDataChangeOfStateCount() bool {
	return true
}

func (m *_BACnetConstructedDataChangeOfStateCount) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
