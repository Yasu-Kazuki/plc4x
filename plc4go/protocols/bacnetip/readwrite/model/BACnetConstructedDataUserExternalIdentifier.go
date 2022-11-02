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

// BACnetConstructedDataUserExternalIdentifier is the corresponding interface of BACnetConstructedDataUserExternalIdentifier
type BACnetConstructedDataUserExternalIdentifier interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetUserExternalIdentifier returns UserExternalIdentifier (property field)
	GetUserExternalIdentifier() BACnetApplicationTagCharacterString
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagCharacterString
}

// BACnetConstructedDataUserExternalIdentifierExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataUserExternalIdentifier.
// This is useful for switch cases.
type BACnetConstructedDataUserExternalIdentifierExactly interface {
	BACnetConstructedDataUserExternalIdentifier
	isBACnetConstructedDataUserExternalIdentifier() bool
}

// _BACnetConstructedDataUserExternalIdentifier is the data-structure of this message
type _BACnetConstructedDataUserExternalIdentifier struct {
	*_BACnetConstructedData
	UserExternalIdentifier BACnetApplicationTagCharacterString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataUserExternalIdentifier) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataUserExternalIdentifier) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_USER_EXTERNAL_IDENTIFIER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataUserExternalIdentifier) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataUserExternalIdentifier) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataUserExternalIdentifier) GetUserExternalIdentifier() BACnetApplicationTagCharacterString {
	return m.UserExternalIdentifier
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataUserExternalIdentifier) GetActualValue() BACnetApplicationTagCharacterString {
	return CastBACnetApplicationTagCharacterString(m.GetUserExternalIdentifier())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataUserExternalIdentifier factory function for _BACnetConstructedDataUserExternalIdentifier
func NewBACnetConstructedDataUserExternalIdentifier(userExternalIdentifier BACnetApplicationTagCharacterString, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataUserExternalIdentifier {
	_result := &_BACnetConstructedDataUserExternalIdentifier{
		UserExternalIdentifier: userExternalIdentifier,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataUserExternalIdentifier(structType interface{}) BACnetConstructedDataUserExternalIdentifier {
	if casted, ok := structType.(BACnetConstructedDataUserExternalIdentifier); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataUserExternalIdentifier); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataUserExternalIdentifier) GetTypeName() string {
	return "BACnetConstructedDataUserExternalIdentifier"
}

func (m *_BACnetConstructedDataUserExternalIdentifier) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataUserExternalIdentifier) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (userExternalIdentifier)
	lengthInBits += m.UserExternalIdentifier.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataUserExternalIdentifier) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataUserExternalIdentifierParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataUserExternalIdentifier, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataUserExternalIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataUserExternalIdentifier")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (userExternalIdentifier)
	if pullErr := readBuffer.PullContext("userExternalIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for userExternalIdentifier")
	}
	_userExternalIdentifier, _userExternalIdentifierErr := BACnetApplicationTagParse(readBuffer)
	if _userExternalIdentifierErr != nil {
		return nil, errors.Wrap(_userExternalIdentifierErr, "Error parsing 'userExternalIdentifier' field of BACnetConstructedDataUserExternalIdentifier")
	}
	userExternalIdentifier := _userExternalIdentifier.(BACnetApplicationTagCharacterString)
	if closeErr := readBuffer.CloseContext("userExternalIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for userExternalIdentifier")
	}

	// Virtual field
	_actualValue := userExternalIdentifier
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataUserExternalIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataUserExternalIdentifier")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataUserExternalIdentifier{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		UserExternalIdentifier: userExternalIdentifier,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataUserExternalIdentifier) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataUserExternalIdentifier) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataUserExternalIdentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataUserExternalIdentifier")
		}

		// Simple Field (userExternalIdentifier)
		if pushErr := writeBuffer.PushContext("userExternalIdentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for userExternalIdentifier")
		}
		_userExternalIdentifierErr := writeBuffer.WriteSerializable(m.GetUserExternalIdentifier())
		if popErr := writeBuffer.PopContext("userExternalIdentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for userExternalIdentifier")
		}
		if _userExternalIdentifierErr != nil {
			return errors.Wrap(_userExternalIdentifierErr, "Error serializing 'userExternalIdentifier' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataUserExternalIdentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataUserExternalIdentifier")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataUserExternalIdentifier) isBACnetConstructedDataUserExternalIdentifier() bool {
	return true
}

func (m *_BACnetConstructedDataUserExternalIdentifier) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
