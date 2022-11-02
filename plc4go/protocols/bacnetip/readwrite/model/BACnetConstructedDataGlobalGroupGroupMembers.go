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
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataGlobalGroupGroupMembers is the corresponding interface of BACnetConstructedDataGlobalGroupGroupMembers
type BACnetConstructedDataGlobalGroupGroupMembers interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetNumberOfDataElements returns NumberOfDataElements (property field)
	GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger
	// GetGroupMembers returns GroupMembers (property field)
	GetGroupMembers() []BACnetDeviceObjectPropertyReference
	// GetZero returns Zero (virtual field)
	GetZero() uint64
}

// BACnetConstructedDataGlobalGroupGroupMembersExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataGlobalGroupGroupMembers.
// This is useful for switch cases.
type BACnetConstructedDataGlobalGroupGroupMembersExactly interface {
	BACnetConstructedDataGlobalGroupGroupMembers
	isBACnetConstructedDataGlobalGroupGroupMembers() bool
}

// _BACnetConstructedDataGlobalGroupGroupMembers is the data-structure of this message
type _BACnetConstructedDataGlobalGroupGroupMembers struct {
	*_BACnetConstructedData
	NumberOfDataElements BACnetApplicationTagUnsignedInteger
	GroupMembers         []BACnetDeviceObjectPropertyReference
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataGlobalGroupGroupMembers) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_GLOBAL_GROUP
}

func (m *_BACnetConstructedDataGlobalGroupGroupMembers) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_GROUP_MEMBERS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataGlobalGroupGroupMembers) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataGlobalGroupGroupMembers) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataGlobalGroupGroupMembers) GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger {
	return m.NumberOfDataElements
}

func (m *_BACnetConstructedDataGlobalGroupGroupMembers) GetGroupMembers() []BACnetDeviceObjectPropertyReference {
	return m.GroupMembers
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataGlobalGroupGroupMembers) GetZero() uint64 {
	numberOfDataElements := m.NumberOfDataElements
	_ = numberOfDataElements
	return uint64(uint64(0))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataGlobalGroupGroupMembers factory function for _BACnetConstructedDataGlobalGroupGroupMembers
func NewBACnetConstructedDataGlobalGroupGroupMembers(numberOfDataElements BACnetApplicationTagUnsignedInteger, groupMembers []BACnetDeviceObjectPropertyReference, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataGlobalGroupGroupMembers {
	_result := &_BACnetConstructedDataGlobalGroupGroupMembers{
		NumberOfDataElements:   numberOfDataElements,
		GroupMembers:           groupMembers,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataGlobalGroupGroupMembers(structType interface{}) BACnetConstructedDataGlobalGroupGroupMembers {
	if casted, ok := structType.(BACnetConstructedDataGlobalGroupGroupMembers); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataGlobalGroupGroupMembers); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataGlobalGroupGroupMembers) GetTypeName() string {
	return "BACnetConstructedDataGlobalGroupGroupMembers"
}

func (m *_BACnetConstructedDataGlobalGroupGroupMembers) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataGlobalGroupGroupMembers) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// A virtual field doesn't have any in- or output.

	// Optional Field (numberOfDataElements)
	if m.NumberOfDataElements != nil {
		lengthInBits += m.NumberOfDataElements.GetLengthInBits()
	}

	// Array field
	if len(m.GroupMembers) > 0 {
		for _, element := range m.GroupMembers {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataGlobalGroupGroupMembers) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataGlobalGroupGroupMembersParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataGlobalGroupGroupMembers, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataGlobalGroupGroupMembers"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataGlobalGroupGroupMembers")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Virtual field
	_zero := uint64(0)
	zero := uint64(_zero)
	_ = zero

	// Optional Field (numberOfDataElements) (Can be skipped, if a given expression evaluates to false)
	var numberOfDataElements BACnetApplicationTagUnsignedInteger = nil
	if bool(bool((arrayIndexArgument) != (nil))) && bool(bool((arrayIndexArgument.GetActualValue()) == (zero))) {
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("numberOfDataElements"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for numberOfDataElements")
		}
		_val, _err := BACnetApplicationTagParse(readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			Plc4xModelLog.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'numberOfDataElements' field of BACnetConstructedDataGlobalGroupGroupMembers")
		default:
			numberOfDataElements = _val.(BACnetApplicationTagUnsignedInteger)
			if closeErr := readBuffer.CloseContext("numberOfDataElements"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for numberOfDataElements")
			}
		}
	}

	// Array field (groupMembers)
	if pullErr := readBuffer.PullContext("groupMembers", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for groupMembers")
	}
	// Terminated array
	var groupMembers []BACnetDeviceObjectPropertyReference
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetDeviceObjectPropertyReferenceParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'groupMembers' field of BACnetConstructedDataGlobalGroupGroupMembers")
			}
			groupMembers = append(groupMembers, _item.(BACnetDeviceObjectPropertyReference))

		}
	}
	if closeErr := readBuffer.CloseContext("groupMembers", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for groupMembers")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataGlobalGroupGroupMembers"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataGlobalGroupGroupMembers")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataGlobalGroupGroupMembers{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		NumberOfDataElements: numberOfDataElements,
		GroupMembers:         groupMembers,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataGlobalGroupGroupMembers) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataGlobalGroupGroupMembers) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataGlobalGroupGroupMembers"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataGlobalGroupGroupMembers")
		}
		// Virtual field
		if _zeroErr := writeBuffer.WriteVirtual("zero", m.GetZero()); _zeroErr != nil {
			return errors.Wrap(_zeroErr, "Error serializing 'zero' field")
		}

		// Optional Field (numberOfDataElements) (Can be skipped, if the value is null)
		var numberOfDataElements BACnetApplicationTagUnsignedInteger = nil
		if m.GetNumberOfDataElements() != nil {
			if pushErr := writeBuffer.PushContext("numberOfDataElements"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for numberOfDataElements")
			}
			numberOfDataElements = m.GetNumberOfDataElements()
			_numberOfDataElementsErr := writeBuffer.WriteSerializable(numberOfDataElements)
			if popErr := writeBuffer.PopContext("numberOfDataElements"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for numberOfDataElements")
			}
			if _numberOfDataElementsErr != nil {
				return errors.Wrap(_numberOfDataElementsErr, "Error serializing 'numberOfDataElements' field")
			}
		}

		// Array Field (groupMembers)
		if pushErr := writeBuffer.PushContext("groupMembers", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for groupMembers")
		}
		for _, _element := range m.GetGroupMembers() {
			_elementErr := writeBuffer.WriteSerializable(_element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'groupMembers' field")
			}
		}
		if popErr := writeBuffer.PopContext("groupMembers", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for groupMembers")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataGlobalGroupGroupMembers"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataGlobalGroupGroupMembers")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataGlobalGroupGroupMembers) isBACnetConstructedDataGlobalGroupGroupMembers() bool {
	return true
}

func (m *_BACnetConstructedDataGlobalGroupGroupMembers) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
