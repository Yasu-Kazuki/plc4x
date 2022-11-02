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

// BACnetConstructedDataApplicationSoftwareVersion is the corresponding interface of BACnetConstructedDataApplicationSoftwareVersion
type BACnetConstructedDataApplicationSoftwareVersion interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetApplicationSoftwareVersion returns ApplicationSoftwareVersion (property field)
	GetApplicationSoftwareVersion() BACnetApplicationTagCharacterString
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagCharacterString
}

// BACnetConstructedDataApplicationSoftwareVersionExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataApplicationSoftwareVersion.
// This is useful for switch cases.
type BACnetConstructedDataApplicationSoftwareVersionExactly interface {
	BACnetConstructedDataApplicationSoftwareVersion
	isBACnetConstructedDataApplicationSoftwareVersion() bool
}

// _BACnetConstructedDataApplicationSoftwareVersion is the data-structure of this message
type _BACnetConstructedDataApplicationSoftwareVersion struct {
	*_BACnetConstructedData
	ApplicationSoftwareVersion BACnetApplicationTagCharacterString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataApplicationSoftwareVersion) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataApplicationSoftwareVersion) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_APPLICATION_SOFTWARE_VERSION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataApplicationSoftwareVersion) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataApplicationSoftwareVersion) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataApplicationSoftwareVersion) GetApplicationSoftwareVersion() BACnetApplicationTagCharacterString {
	return m.ApplicationSoftwareVersion
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataApplicationSoftwareVersion) GetActualValue() BACnetApplicationTagCharacterString {
	return CastBACnetApplicationTagCharacterString(m.GetApplicationSoftwareVersion())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataApplicationSoftwareVersion factory function for _BACnetConstructedDataApplicationSoftwareVersion
func NewBACnetConstructedDataApplicationSoftwareVersion(applicationSoftwareVersion BACnetApplicationTagCharacterString, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataApplicationSoftwareVersion {
	_result := &_BACnetConstructedDataApplicationSoftwareVersion{
		ApplicationSoftwareVersion: applicationSoftwareVersion,
		_BACnetConstructedData:     NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataApplicationSoftwareVersion(structType interface{}) BACnetConstructedDataApplicationSoftwareVersion {
	if casted, ok := structType.(BACnetConstructedDataApplicationSoftwareVersion); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataApplicationSoftwareVersion); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataApplicationSoftwareVersion) GetTypeName() string {
	return "BACnetConstructedDataApplicationSoftwareVersion"
}

func (m *_BACnetConstructedDataApplicationSoftwareVersion) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataApplicationSoftwareVersion) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (applicationSoftwareVersion)
	lengthInBits += m.ApplicationSoftwareVersion.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataApplicationSoftwareVersion) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataApplicationSoftwareVersionParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataApplicationSoftwareVersion, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataApplicationSoftwareVersion"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataApplicationSoftwareVersion")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (applicationSoftwareVersion)
	if pullErr := readBuffer.PullContext("applicationSoftwareVersion"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for applicationSoftwareVersion")
	}
	_applicationSoftwareVersion, _applicationSoftwareVersionErr := BACnetApplicationTagParse(readBuffer)
	if _applicationSoftwareVersionErr != nil {
		return nil, errors.Wrap(_applicationSoftwareVersionErr, "Error parsing 'applicationSoftwareVersion' field of BACnetConstructedDataApplicationSoftwareVersion")
	}
	applicationSoftwareVersion := _applicationSoftwareVersion.(BACnetApplicationTagCharacterString)
	if closeErr := readBuffer.CloseContext("applicationSoftwareVersion"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for applicationSoftwareVersion")
	}

	// Virtual field
	_actualValue := applicationSoftwareVersion
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataApplicationSoftwareVersion"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataApplicationSoftwareVersion")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataApplicationSoftwareVersion{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		ApplicationSoftwareVersion: applicationSoftwareVersion,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataApplicationSoftwareVersion) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataApplicationSoftwareVersion) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataApplicationSoftwareVersion"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataApplicationSoftwareVersion")
		}

		// Simple Field (applicationSoftwareVersion)
		if pushErr := writeBuffer.PushContext("applicationSoftwareVersion"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for applicationSoftwareVersion")
		}
		_applicationSoftwareVersionErr := writeBuffer.WriteSerializable(m.GetApplicationSoftwareVersion())
		if popErr := writeBuffer.PopContext("applicationSoftwareVersion"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for applicationSoftwareVersion")
		}
		if _applicationSoftwareVersionErr != nil {
			return errors.Wrap(_applicationSoftwareVersionErr, "Error serializing 'applicationSoftwareVersion' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataApplicationSoftwareVersion"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataApplicationSoftwareVersion")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataApplicationSoftwareVersion) isBACnetConstructedDataApplicationSoftwareVersion() bool {
	return true
}

func (m *_BACnetConstructedDataApplicationSoftwareVersion) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
