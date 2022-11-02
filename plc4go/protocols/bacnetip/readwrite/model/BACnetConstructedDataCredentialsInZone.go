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

// BACnetConstructedDataCredentialsInZone is the corresponding interface of BACnetConstructedDataCredentialsInZone
type BACnetConstructedDataCredentialsInZone interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetCredentialsInZone returns CredentialsInZone (property field)
	GetCredentialsInZone() []BACnetDeviceObjectReference
}

// BACnetConstructedDataCredentialsInZoneExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataCredentialsInZone.
// This is useful for switch cases.
type BACnetConstructedDataCredentialsInZoneExactly interface {
	BACnetConstructedDataCredentialsInZone
	isBACnetConstructedDataCredentialsInZone() bool
}

// _BACnetConstructedDataCredentialsInZone is the data-structure of this message
type _BACnetConstructedDataCredentialsInZone struct {
	*_BACnetConstructedData
	CredentialsInZone []BACnetDeviceObjectReference
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataCredentialsInZone) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataCredentialsInZone) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_CREDENTIALS_IN_ZONE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataCredentialsInZone) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataCredentialsInZone) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataCredentialsInZone) GetCredentialsInZone() []BACnetDeviceObjectReference {
	return m.CredentialsInZone
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataCredentialsInZone factory function for _BACnetConstructedDataCredentialsInZone
func NewBACnetConstructedDataCredentialsInZone(credentialsInZone []BACnetDeviceObjectReference, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataCredentialsInZone {
	_result := &_BACnetConstructedDataCredentialsInZone{
		CredentialsInZone:      credentialsInZone,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataCredentialsInZone(structType interface{}) BACnetConstructedDataCredentialsInZone {
	if casted, ok := structType.(BACnetConstructedDataCredentialsInZone); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataCredentialsInZone); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataCredentialsInZone) GetTypeName() string {
	return "BACnetConstructedDataCredentialsInZone"
}

func (m *_BACnetConstructedDataCredentialsInZone) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataCredentialsInZone) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.CredentialsInZone) > 0 {
		for _, element := range m.CredentialsInZone {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataCredentialsInZone) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataCredentialsInZoneParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataCredentialsInZone, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataCredentialsInZone"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataCredentialsInZone")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (credentialsInZone)
	if pullErr := readBuffer.PullContext("credentialsInZone", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for credentialsInZone")
	}
	// Terminated array
	var credentialsInZone []BACnetDeviceObjectReference
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetDeviceObjectReferenceParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'credentialsInZone' field of BACnetConstructedDataCredentialsInZone")
			}
			credentialsInZone = append(credentialsInZone, _item.(BACnetDeviceObjectReference))

		}
	}
	if closeErr := readBuffer.CloseContext("credentialsInZone", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for credentialsInZone")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataCredentialsInZone"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataCredentialsInZone")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataCredentialsInZone{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		CredentialsInZone: credentialsInZone,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataCredentialsInZone) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataCredentialsInZone) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataCredentialsInZone"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataCredentialsInZone")
		}

		// Array Field (credentialsInZone)
		if pushErr := writeBuffer.PushContext("credentialsInZone", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for credentialsInZone")
		}
		for _, _element := range m.GetCredentialsInZone() {
			_elementErr := writeBuffer.WriteSerializable(_element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'credentialsInZone' field")
			}
		}
		if popErr := writeBuffer.PopContext("credentialsInZone", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for credentialsInZone")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataCredentialsInZone"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataCredentialsInZone")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataCredentialsInZone) isBACnetConstructedDataCredentialsInZone() bool {
	return true
}

func (m *_BACnetConstructedDataCredentialsInZone) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
