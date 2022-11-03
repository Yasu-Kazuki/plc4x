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

// BACnetLogRecordLogDatum is the corresponding interface of BACnetLogRecordLogDatum
type BACnetLogRecordLogDatum interface {
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
}

// BACnetLogRecordLogDatumExactly can be used when we want exactly this type and not a type which fulfills BACnetLogRecordLogDatum.
// This is useful for switch cases.
type BACnetLogRecordLogDatumExactly interface {
	BACnetLogRecordLogDatum
	isBACnetLogRecordLogDatum() bool
}

// _BACnetLogRecordLogDatum is the data-structure of this message
type _BACnetLogRecordLogDatum struct {
	_BACnetLogRecordLogDatumChildRequirements
	OpeningTag      BACnetOpeningTag
	PeekedTagHeader BACnetTagHeader
	ClosingTag      BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

type _BACnetLogRecordLogDatumChildRequirements interface {
	utils.Serializable
	GetLengthInBits() uint16
	GetLengthInBitsConditional(lastItem bool) uint16
}

type BACnetLogRecordLogDatumParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child BACnetLogRecordLogDatum, serializeChildFunction func() error) error
	GetTypeName() string
}

type BACnetLogRecordLogDatumChild interface {
	utils.Serializable
	InitializeParent(parent BACnetLogRecordLogDatum, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag)
	GetParent() *BACnetLogRecordLogDatum

	GetTypeName() string
	BACnetLogRecordLogDatum
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLogRecordLogDatum) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetLogRecordLogDatum) GetPeekedTagHeader() BACnetTagHeader {
	return m.PeekedTagHeader
}

func (m *_BACnetLogRecordLogDatum) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetLogRecordLogDatum) GetPeekedTagNumber() uint8 {
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetLogRecordLogDatum factory function for _BACnetLogRecordLogDatum
func NewBACnetLogRecordLogDatum(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetLogRecordLogDatum {
	return &_BACnetLogRecordLogDatum{OpeningTag: openingTag, PeekedTagHeader: peekedTagHeader, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastBACnetLogRecordLogDatum(structType interface{}) BACnetLogRecordLogDatum {
	if casted, ok := structType.(BACnetLogRecordLogDatum); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLogRecordLogDatum); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLogRecordLogDatum) GetTypeName() string {
	return "BACnetLogRecordLogDatum"
}

func (m *_BACnetLogRecordLogDatum) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetLogRecordLogDatum) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetLogRecordLogDatumParse(theBytes []byte, tagNumber uint8) (BACnetLogRecordLogDatum, error) {
	return BACnetLogRecordLogDatumParseWithBuffer(utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)), tagNumber) // TODO: get endianness from mspec
}

func BACnetLogRecordLogDatumParseWithBuffer(readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetLogRecordLogDatum, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLogRecordLogDatum"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLogRecordLogDatum")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParseWithBuffer(readBuffer, uint8(tagNumber))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field of BACnetLogRecordLogDatum")
	}
	openingTag := _openingTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for openingTag")
	}

	// Peek Field (peekedTagHeader)
	currentPos = positionAware.GetPos()
	if pullErr := readBuffer.PullContext("peekedTagHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for peekedTagHeader")
	}
	peekedTagHeader, _ := BACnetTagHeaderParseWithBuffer(readBuffer)
	readBuffer.Reset(currentPos)

	// Virtual field
	_peekedTagNumber := peekedTagHeader.GetActualTagNumber()
	peekedTagNumber := uint8(_peekedTagNumber)
	_ = peekedTagNumber

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type BACnetLogRecordLogDatumChildSerializeRequirement interface {
		BACnetLogRecordLogDatum
		InitializeParent(BACnetLogRecordLogDatum, BACnetOpeningTag, BACnetTagHeader, BACnetClosingTag)
		GetParent() BACnetLogRecordLogDatum
	}
	var _childTemp interface{}
	var _child BACnetLogRecordLogDatumChildSerializeRequirement
	var typeSwitchError error
	switch {
	case peekedTagNumber == uint8(0): // BACnetLogRecordLogDatumLogStatus
		_childTemp, typeSwitchError = BACnetLogRecordLogDatumLogStatusParseWithBuffer(readBuffer, tagNumber)
	case peekedTagNumber == uint8(1): // BACnetLogRecordLogDatumBooleanValue
		_childTemp, typeSwitchError = BACnetLogRecordLogDatumBooleanValueParseWithBuffer(readBuffer, tagNumber)
	case peekedTagNumber == uint8(2): // BACnetLogRecordLogDatumRealValue
		_childTemp, typeSwitchError = BACnetLogRecordLogDatumRealValueParseWithBuffer(readBuffer, tagNumber)
	case peekedTagNumber == uint8(3): // BACnetLogRecordLogDatumEnumeratedValue
		_childTemp, typeSwitchError = BACnetLogRecordLogDatumEnumeratedValueParseWithBuffer(readBuffer, tagNumber)
	case peekedTagNumber == uint8(4): // BACnetLogRecordLogDatumUnsignedValue
		_childTemp, typeSwitchError = BACnetLogRecordLogDatumUnsignedValueParseWithBuffer(readBuffer, tagNumber)
	case peekedTagNumber == uint8(5): // BACnetLogRecordLogDatumIntegerValue
		_childTemp, typeSwitchError = BACnetLogRecordLogDatumIntegerValueParseWithBuffer(readBuffer, tagNumber)
	case peekedTagNumber == uint8(6): // BACnetLogRecordLogDatumBitStringValue
		_childTemp, typeSwitchError = BACnetLogRecordLogDatumBitStringValueParseWithBuffer(readBuffer, tagNumber)
	case peekedTagNumber == uint8(7): // BACnetLogRecordLogDatumNullValue
		_childTemp, typeSwitchError = BACnetLogRecordLogDatumNullValueParseWithBuffer(readBuffer, tagNumber)
	case peekedTagNumber == uint8(8): // BACnetLogRecordLogDatumFailure
		_childTemp, typeSwitchError = BACnetLogRecordLogDatumFailureParseWithBuffer(readBuffer, tagNumber)
	case peekedTagNumber == uint8(9): // BACnetLogRecordLogDatumTimeChange
		_childTemp, typeSwitchError = BACnetLogRecordLogDatumTimeChangeParseWithBuffer(readBuffer, tagNumber)
	case peekedTagNumber == uint8(10): // BACnetLogRecordLogDatumAnyValue
		_childTemp, typeSwitchError = BACnetLogRecordLogDatumAnyValueParseWithBuffer(readBuffer, tagNumber)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [peekedTagNumber=%v]", peekedTagNumber)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of BACnetLogRecordLogDatum")
	}
	_child = _childTemp.(BACnetLogRecordLogDatumChildSerializeRequirement)

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
	_closingTag, _closingTagErr := BACnetClosingTagParseWithBuffer(readBuffer, uint8(tagNumber))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field of BACnetLogRecordLogDatum")
	}
	closingTag := _closingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetLogRecordLogDatum"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLogRecordLogDatum")
	}

	// Finish initializing
	_child.InitializeParent(_child, openingTag, peekedTagHeader, closingTag)
	return _child, nil
}

func (pm *_BACnetLogRecordLogDatum) SerializeParent(writeBuffer utils.WriteBuffer, child BACnetLogRecordLogDatum, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetLogRecordLogDatum"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetLogRecordLogDatum")
	}

	// Simple Field (openingTag)
	if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for openingTag")
	}
	_openingTagErr := writeBuffer.WriteSerializable(m.GetOpeningTag())
	if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for openingTag")
	}
	if _openingTagErr != nil {
		return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
	}
	// Virtual field
	if _peekedTagNumberErr := writeBuffer.WriteVirtual("peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	// Simple Field (closingTag)
	if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for closingTag")
	}
	_closingTagErr := writeBuffer.WriteSerializable(m.GetClosingTag())
	if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for closingTag")
	}
	if _closingTagErr != nil {
		return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetLogRecordLogDatum"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetLogRecordLogDatum")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetLogRecordLogDatum) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetLogRecordLogDatum) isBACnetLogRecordLogDatum() bool {
	return true
}

func (m *_BACnetLogRecordLogDatum) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}