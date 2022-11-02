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

// BACnetConstructedDataInactiveText is the corresponding interface of BACnetConstructedDataInactiveText
type BACnetConstructedDataInactiveText interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetInactiveText returns InactiveText (property field)
	GetInactiveText() BACnetApplicationTagCharacterString
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagCharacterString
}

// BACnetConstructedDataInactiveTextExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataInactiveText.
// This is useful for switch cases.
type BACnetConstructedDataInactiveTextExactly interface {
	BACnetConstructedDataInactiveText
	isBACnetConstructedDataInactiveText() bool
}

// _BACnetConstructedDataInactiveText is the data-structure of this message
type _BACnetConstructedDataInactiveText struct {
	*_BACnetConstructedData
	InactiveText BACnetApplicationTagCharacterString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataInactiveText) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataInactiveText) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_INACTIVE_TEXT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataInactiveText) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataInactiveText) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataInactiveText) GetInactiveText() BACnetApplicationTagCharacterString {
	return m.InactiveText
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataInactiveText) GetActualValue() BACnetApplicationTagCharacterString {
	return CastBACnetApplicationTagCharacterString(m.GetInactiveText())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataInactiveText factory function for _BACnetConstructedDataInactiveText
func NewBACnetConstructedDataInactiveText(inactiveText BACnetApplicationTagCharacterString, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataInactiveText {
	_result := &_BACnetConstructedDataInactiveText{
		InactiveText:           inactiveText,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataInactiveText(structType interface{}) BACnetConstructedDataInactiveText {
	if casted, ok := structType.(BACnetConstructedDataInactiveText); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataInactiveText); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataInactiveText) GetTypeName() string {
	return "BACnetConstructedDataInactiveText"
}

func (m *_BACnetConstructedDataInactiveText) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataInactiveText) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (inactiveText)
	lengthInBits += m.InactiveText.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataInactiveText) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataInactiveTextParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataInactiveText, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataInactiveText"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataInactiveText")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (inactiveText)
	if pullErr := readBuffer.PullContext("inactiveText"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for inactiveText")
	}
	_inactiveText, _inactiveTextErr := BACnetApplicationTagParse(readBuffer)
	if _inactiveTextErr != nil {
		return nil, errors.Wrap(_inactiveTextErr, "Error parsing 'inactiveText' field of BACnetConstructedDataInactiveText")
	}
	inactiveText := _inactiveText.(BACnetApplicationTagCharacterString)
	if closeErr := readBuffer.CloseContext("inactiveText"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for inactiveText")
	}

	// Virtual field
	_actualValue := inactiveText
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataInactiveText"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataInactiveText")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataInactiveText{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		InactiveText: inactiveText,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataInactiveText) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataInactiveText) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataInactiveText"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataInactiveText")
		}

		// Simple Field (inactiveText)
		if pushErr := writeBuffer.PushContext("inactiveText"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for inactiveText")
		}
		_inactiveTextErr := writeBuffer.WriteSerializable(m.GetInactiveText())
		if popErr := writeBuffer.PopContext("inactiveText"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for inactiveText")
		}
		if _inactiveTextErr != nil {
			return errors.Wrap(_inactiveTextErr, "Error serializing 'inactiveText' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataInactiveText"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataInactiveText")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataInactiveText) isBACnetConstructedDataInactiveText() bool {
	return true
}

func (m *_BACnetConstructedDataInactiveText) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
