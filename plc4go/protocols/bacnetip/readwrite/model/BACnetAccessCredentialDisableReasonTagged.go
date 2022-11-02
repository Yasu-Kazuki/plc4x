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

// BACnetAccessCredentialDisableReasonTagged is the corresponding interface of BACnetAccessCredentialDisableReasonTagged
type BACnetAccessCredentialDisableReasonTagged interface {
	utils.LengthAware
	utils.Serializable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetAccessCredentialDisableReason
	// GetProprietaryValue returns ProprietaryValue (property field)
	GetProprietaryValue() uint32
	// GetIsProprietary returns IsProprietary (virtual field)
	GetIsProprietary() bool
}

// BACnetAccessCredentialDisableReasonTaggedExactly can be used when we want exactly this type and not a type which fulfills BACnetAccessCredentialDisableReasonTagged.
// This is useful for switch cases.
type BACnetAccessCredentialDisableReasonTaggedExactly interface {
	BACnetAccessCredentialDisableReasonTagged
	isBACnetAccessCredentialDisableReasonTagged() bool
}

// _BACnetAccessCredentialDisableReasonTagged is the data-structure of this message
type _BACnetAccessCredentialDisableReasonTagged struct {
	Header           BACnetTagHeader
	Value            BACnetAccessCredentialDisableReason
	ProprietaryValue uint32

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetAccessCredentialDisableReasonTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetAccessCredentialDisableReasonTagged) GetValue() BACnetAccessCredentialDisableReason {
	return m.Value
}

func (m *_BACnetAccessCredentialDisableReasonTagged) GetProprietaryValue() uint32 {
	return m.ProprietaryValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetAccessCredentialDisableReasonTagged) GetIsProprietary() bool {
	return bool(bool((m.GetValue()) == (BACnetAccessCredentialDisableReason_VENDOR_PROPRIETARY_VALUE)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetAccessCredentialDisableReasonTagged factory function for _BACnetAccessCredentialDisableReasonTagged
func NewBACnetAccessCredentialDisableReasonTagged(header BACnetTagHeader, value BACnetAccessCredentialDisableReason, proprietaryValue uint32, tagNumber uint8, tagClass TagClass) *_BACnetAccessCredentialDisableReasonTagged {
	return &_BACnetAccessCredentialDisableReasonTagged{Header: header, Value: value, ProprietaryValue: proprietaryValue, TagNumber: tagNumber, TagClass: tagClass}
}

// Deprecated: use the interface for direct cast
func CastBACnetAccessCredentialDisableReasonTagged(structType interface{}) BACnetAccessCredentialDisableReasonTagged {
	if casted, ok := structType.(BACnetAccessCredentialDisableReasonTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetAccessCredentialDisableReasonTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetAccessCredentialDisableReasonTagged) GetTypeName() string {
	return "BACnetAccessCredentialDisableReasonTagged"
}

func (m *_BACnetAccessCredentialDisableReasonTagged) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetAccessCredentialDisableReasonTagged) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits()

	// Manual Field (value)
	lengthInBits += uint16(utils.InlineIf(m.GetIsProprietary(), func() interface{} { return int32(int32(0)) }, func() interface{} { return int32((int32(m.GetHeader().GetActualLength()) * int32(int32(8)))) }).(int32))

	// A virtual field doesn't have any in- or output.

	// Manual Field (proprietaryValue)
	lengthInBits += uint16(utils.InlineIf(m.GetIsProprietary(), func() interface{} { return int32((int32(m.GetHeader().GetActualLength()) * int32(int32(8)))) }, func() interface{} { return int32(int32(0)) }).(int32))

	return lengthInBits
}

func (m *_BACnetAccessCredentialDisableReasonTagged) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetAccessCredentialDisableReasonTaggedParse(readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetAccessCredentialDisableReasonTagged, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetAccessCredentialDisableReasonTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetAccessCredentialDisableReasonTagged")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (header)
	if pullErr := readBuffer.PullContext("header"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for header")
	}
	_header, _headerErr := BACnetTagHeaderParse(readBuffer)
	if _headerErr != nil {
		return nil, errors.Wrap(_headerErr, "Error parsing 'header' field of BACnetAccessCredentialDisableReasonTagged")
	}
	header := _header.(BACnetTagHeader)
	if closeErr := readBuffer.CloseContext("header"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for header")
	}

	// Validation
	if !(bool((header.GetTagClass()) == (tagClass))) {
		return nil, errors.WithStack(utils.ParseValidationError{"tag class doesn't match"})
	}

	// Validation
	if !(bool((bool((header.GetTagClass()) == (TagClass_APPLICATION_TAGS)))) || bool((bool((header.GetActualTagNumber()) == (tagNumber))))) {
		return nil, errors.WithStack(utils.ParseAssertError{"tagnumber doesn't match"})
	}

	// Manual Field (value)
	_value, _valueErr := ReadEnumGeneric(readBuffer, header.GetActualLength(), BACnetAccessCredentialDisableReason_VENDOR_PROPRIETARY_VALUE)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field of BACnetAccessCredentialDisableReasonTagged")
	}
	var value BACnetAccessCredentialDisableReason
	if _value != nil {
		value = _value.(BACnetAccessCredentialDisableReason)
	}

	// Virtual field
	_isProprietary := bool((value) == (BACnetAccessCredentialDisableReason_VENDOR_PROPRIETARY_VALUE))
	isProprietary := bool(_isProprietary)
	_ = isProprietary

	// Manual Field (proprietaryValue)
	_proprietaryValue, _proprietaryValueErr := ReadProprietaryEnumGeneric(readBuffer, header.GetActualLength(), isProprietary)
	if _proprietaryValueErr != nil {
		return nil, errors.Wrap(_proprietaryValueErr, "Error parsing 'proprietaryValue' field of BACnetAccessCredentialDisableReasonTagged")
	}
	var proprietaryValue uint32
	if _proprietaryValue != nil {
		proprietaryValue = _proprietaryValue.(uint32)
	}

	if closeErr := readBuffer.CloseContext("BACnetAccessCredentialDisableReasonTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetAccessCredentialDisableReasonTagged")
	}

	// Create the instance
	return &_BACnetAccessCredentialDisableReasonTagged{
		TagNumber:        tagNumber,
		TagClass:         tagClass,
		Header:           header,
		Value:            value,
		ProprietaryValue: proprietaryValue,
	}, nil
}

func (m *_BACnetAccessCredentialDisableReasonTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetAccessCredentialDisableReasonTagged) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetAccessCredentialDisableReasonTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetAccessCredentialDisableReasonTagged")
	}

	// Simple Field (header)
	if pushErr := writeBuffer.PushContext("header"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for header")
	}
	_headerErr := writeBuffer.WriteSerializable(m.GetHeader())
	if popErr := writeBuffer.PopContext("header"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for header")
	}
	if _headerErr != nil {
		return errors.Wrap(_headerErr, "Error serializing 'header' field")
	}

	// Manual Field (value)
	_valueErr := WriteEnumGeneric(writeBuffer, m.GetValue())
	if _valueErr != nil {
		return errors.Wrap(_valueErr, "Error serializing 'value' field")
	}
	// Virtual field
	if _isProprietaryErr := writeBuffer.WriteVirtual("isProprietary", m.GetIsProprietary()); _isProprietaryErr != nil {
		return errors.Wrap(_isProprietaryErr, "Error serializing 'isProprietary' field")
	}

	// Manual Field (proprietaryValue)
	_proprietaryValueErr := WriteProprietaryEnumGeneric(writeBuffer, m.GetProprietaryValue(), m.GetIsProprietary())
	if _proprietaryValueErr != nil {
		return errors.Wrap(_proprietaryValueErr, "Error serializing 'proprietaryValue' field")
	}

	if popErr := writeBuffer.PopContext("BACnetAccessCredentialDisableReasonTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetAccessCredentialDisableReasonTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetAccessCredentialDisableReasonTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetAccessCredentialDisableReasonTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetAccessCredentialDisableReasonTagged) isBACnetAccessCredentialDisableReasonTagged() bool {
	return true
}

func (m *_BACnetAccessCredentialDisableReasonTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
