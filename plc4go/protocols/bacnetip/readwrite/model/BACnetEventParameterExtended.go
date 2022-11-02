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

// BACnetEventParameterExtended is the corresponding interface of BACnetEventParameterExtended
type BACnetEventParameterExtended interface {
	utils.LengthAware
	utils.Serializable
	BACnetEventParameter
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetVendorId returns VendorId (property field)
	GetVendorId() BACnetVendorIdTagged
	// GetExtendedEventType returns ExtendedEventType (property field)
	GetExtendedEventType() BACnetContextTagUnsignedInteger
	// GetParameters returns Parameters (property field)
	GetParameters() BACnetEventParameterExtendedParameters
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetEventParameterExtendedExactly can be used when we want exactly this type and not a type which fulfills BACnetEventParameterExtended.
// This is useful for switch cases.
type BACnetEventParameterExtendedExactly interface {
	BACnetEventParameterExtended
	isBACnetEventParameterExtended() bool
}

// _BACnetEventParameterExtended is the data-structure of this message
type _BACnetEventParameterExtended struct {
	*_BACnetEventParameter
	OpeningTag        BACnetOpeningTag
	VendorId          BACnetVendorIdTagged
	ExtendedEventType BACnetContextTagUnsignedInteger
	Parameters        BACnetEventParameterExtendedParameters
	ClosingTag        BACnetClosingTag
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetEventParameterExtended) InitializeParent(parent BACnetEventParameter, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetEventParameterExtended) GetParent() BACnetEventParameter {
	return m._BACnetEventParameter
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventParameterExtended) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetEventParameterExtended) GetVendorId() BACnetVendorIdTagged {
	return m.VendorId
}

func (m *_BACnetEventParameterExtended) GetExtendedEventType() BACnetContextTagUnsignedInteger {
	return m.ExtendedEventType
}

func (m *_BACnetEventParameterExtended) GetParameters() BACnetEventParameterExtendedParameters {
	return m.Parameters
}

func (m *_BACnetEventParameterExtended) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetEventParameterExtended factory function for _BACnetEventParameterExtended
func NewBACnetEventParameterExtended(openingTag BACnetOpeningTag, vendorId BACnetVendorIdTagged, extendedEventType BACnetContextTagUnsignedInteger, parameters BACnetEventParameterExtendedParameters, closingTag BACnetClosingTag, peekedTagHeader BACnetTagHeader) *_BACnetEventParameterExtended {
	_result := &_BACnetEventParameterExtended{
		OpeningTag:            openingTag,
		VendorId:              vendorId,
		ExtendedEventType:     extendedEventType,
		Parameters:            parameters,
		ClosingTag:            closingTag,
		_BACnetEventParameter: NewBACnetEventParameter(peekedTagHeader),
	}
	_result._BACnetEventParameter._BACnetEventParameterChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetEventParameterExtended(structType interface{}) BACnetEventParameterExtended {
	if casted, ok := structType.(BACnetEventParameterExtended); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventParameterExtended); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventParameterExtended) GetTypeName() string {
	return "BACnetEventParameterExtended"
}

func (m *_BACnetEventParameterExtended) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetEventParameterExtended) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits()

	// Simple field (vendorId)
	lengthInBits += m.VendorId.GetLengthInBits()

	// Simple field (extendedEventType)
	lengthInBits += m.ExtendedEventType.GetLengthInBits()

	// Simple field (parameters)
	lengthInBits += m.Parameters.GetLengthInBits()

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetEventParameterExtended) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetEventParameterExtendedParse(readBuffer utils.ReadBuffer) (BACnetEventParameterExtended, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventParameterExtended"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventParameterExtended")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParse(readBuffer, uint8(uint8(9)))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field of BACnetEventParameterExtended")
	}
	openingTag := _openingTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for openingTag")
	}

	// Simple Field (vendorId)
	if pullErr := readBuffer.PullContext("vendorId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for vendorId")
	}
	_vendorId, _vendorIdErr := BACnetVendorIdTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _vendorIdErr != nil {
		return nil, errors.Wrap(_vendorIdErr, "Error parsing 'vendorId' field of BACnetEventParameterExtended")
	}
	vendorId := _vendorId.(BACnetVendorIdTagged)
	if closeErr := readBuffer.CloseContext("vendorId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for vendorId")
	}

	// Simple Field (extendedEventType)
	if pullErr := readBuffer.PullContext("extendedEventType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for extendedEventType")
	}
	_extendedEventType, _extendedEventTypeErr := BACnetContextTagParse(readBuffer, uint8(uint8(1)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _extendedEventTypeErr != nil {
		return nil, errors.Wrap(_extendedEventTypeErr, "Error parsing 'extendedEventType' field of BACnetEventParameterExtended")
	}
	extendedEventType := _extendedEventType.(BACnetContextTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("extendedEventType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for extendedEventType")
	}

	// Simple Field (parameters)
	if pullErr := readBuffer.PullContext("parameters"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for parameters")
	}
	_parameters, _parametersErr := BACnetEventParameterExtendedParametersParse(readBuffer, uint8(uint8(2)))
	if _parametersErr != nil {
		return nil, errors.Wrap(_parametersErr, "Error parsing 'parameters' field of BACnetEventParameterExtended")
	}
	parameters := _parameters.(BACnetEventParameterExtendedParameters)
	if closeErr := readBuffer.CloseContext("parameters"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for parameters")
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
	_closingTag, _closingTagErr := BACnetClosingTagParse(readBuffer, uint8(uint8(9)))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field of BACnetEventParameterExtended")
	}
	closingTag := _closingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetEventParameterExtended"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventParameterExtended")
	}

	// Create a partially initialized instance
	_child := &_BACnetEventParameterExtended{
		_BACnetEventParameter: &_BACnetEventParameter{},
		OpeningTag:            openingTag,
		VendorId:              vendorId,
		ExtendedEventType:     extendedEventType,
		Parameters:            parameters,
		ClosingTag:            closingTag,
	}
	_child._BACnetEventParameter._BACnetEventParameterChildRequirements = _child
	return _child, nil
}

func (m *_BACnetEventParameterExtended) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetEventParameterExtended) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetEventParameterExtended"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetEventParameterExtended")
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

		// Simple Field (vendorId)
		if pushErr := writeBuffer.PushContext("vendorId"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for vendorId")
		}
		_vendorIdErr := writeBuffer.WriteSerializable(m.GetVendorId())
		if popErr := writeBuffer.PopContext("vendorId"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for vendorId")
		}
		if _vendorIdErr != nil {
			return errors.Wrap(_vendorIdErr, "Error serializing 'vendorId' field")
		}

		// Simple Field (extendedEventType)
		if pushErr := writeBuffer.PushContext("extendedEventType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for extendedEventType")
		}
		_extendedEventTypeErr := writeBuffer.WriteSerializable(m.GetExtendedEventType())
		if popErr := writeBuffer.PopContext("extendedEventType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for extendedEventType")
		}
		if _extendedEventTypeErr != nil {
			return errors.Wrap(_extendedEventTypeErr, "Error serializing 'extendedEventType' field")
		}

		// Simple Field (parameters)
		if pushErr := writeBuffer.PushContext("parameters"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for parameters")
		}
		_parametersErr := writeBuffer.WriteSerializable(m.GetParameters())
		if popErr := writeBuffer.PopContext("parameters"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for parameters")
		}
		if _parametersErr != nil {
			return errors.Wrap(_parametersErr, "Error serializing 'parameters' field")
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

		if popErr := writeBuffer.PopContext("BACnetEventParameterExtended"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetEventParameterExtended")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetEventParameterExtended) isBACnetEventParameterExtended() bool {
	return true
}

func (m *_BACnetEventParameterExtended) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
