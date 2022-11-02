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

// BACnetConstructedDataUserName is the corresponding interface of BACnetConstructedDataUserName
type BACnetConstructedDataUserName interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetUserName returns UserName (property field)
	GetUserName() BACnetApplicationTagCharacterString
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagCharacterString
}

// BACnetConstructedDataUserNameExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataUserName.
// This is useful for switch cases.
type BACnetConstructedDataUserNameExactly interface {
	BACnetConstructedDataUserName
	isBACnetConstructedDataUserName() bool
}

// _BACnetConstructedDataUserName is the data-structure of this message
type _BACnetConstructedDataUserName struct {
	*_BACnetConstructedData
	UserName BACnetApplicationTagCharacterString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataUserName) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataUserName) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_USER_NAME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataUserName) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataUserName) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataUserName) GetUserName() BACnetApplicationTagCharacterString {
	return m.UserName
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataUserName) GetActualValue() BACnetApplicationTagCharacterString {
	return CastBACnetApplicationTagCharacterString(m.GetUserName())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataUserName factory function for _BACnetConstructedDataUserName
func NewBACnetConstructedDataUserName(userName BACnetApplicationTagCharacterString, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataUserName {
	_result := &_BACnetConstructedDataUserName{
		UserName:               userName,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataUserName(structType interface{}) BACnetConstructedDataUserName {
	if casted, ok := structType.(BACnetConstructedDataUserName); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataUserName); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataUserName) GetTypeName() string {
	return "BACnetConstructedDataUserName"
}

func (m *_BACnetConstructedDataUserName) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataUserName) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (userName)
	lengthInBits += m.UserName.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataUserName) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataUserNameParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataUserName, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataUserName"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataUserName")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (userName)
	if pullErr := readBuffer.PullContext("userName"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for userName")
	}
	_userName, _userNameErr := BACnetApplicationTagParse(readBuffer)
	if _userNameErr != nil {
		return nil, errors.Wrap(_userNameErr, "Error parsing 'userName' field of BACnetConstructedDataUserName")
	}
	userName := _userName.(BACnetApplicationTagCharacterString)
	if closeErr := readBuffer.CloseContext("userName"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for userName")
	}

	// Virtual field
	_actualValue := userName
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataUserName"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataUserName")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataUserName{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		UserName: userName,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataUserName) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataUserName) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataUserName"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataUserName")
		}

		// Simple Field (userName)
		if pushErr := writeBuffer.PushContext("userName"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for userName")
		}
		_userNameErr := writeBuffer.WriteSerializable(m.GetUserName())
		if popErr := writeBuffer.PopContext("userName"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for userName")
		}
		if _userNameErr != nil {
			return errors.Wrap(_userNameErr, "Error serializing 'userName' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataUserName"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataUserName")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataUserName) isBACnetConstructedDataUserName() bool {
	return true
}

func (m *_BACnetConstructedDataUserName) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
