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

// BACnetApplicationTagSignedInteger is the corresponding interface of BACnetApplicationTagSignedInteger
type BACnetApplicationTagSignedInteger interface {
	utils.LengthAware
	utils.Serializable
	BACnetApplicationTag
	// GetPayload returns Payload (property field)
	GetPayload() BACnetTagPayloadSignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() uint64
}

// BACnetApplicationTagSignedIntegerExactly can be used when we want exactly this type and not a type which fulfills BACnetApplicationTagSignedInteger.
// This is useful for switch cases.
type BACnetApplicationTagSignedIntegerExactly interface {
	BACnetApplicationTagSignedInteger
	isBACnetApplicationTagSignedInteger() bool
}

// _BACnetApplicationTagSignedInteger is the data-structure of this message
type _BACnetApplicationTagSignedInteger struct {
	*_BACnetApplicationTag
	Payload BACnetTagPayloadSignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetApplicationTagSignedInteger) InitializeParent(parent BACnetApplicationTag, header BACnetTagHeader) {
	m.Header = header
}

func (m *_BACnetApplicationTagSignedInteger) GetParent() BACnetApplicationTag {
	return m._BACnetApplicationTag
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetApplicationTagSignedInteger) GetPayload() BACnetTagPayloadSignedInteger {
	return m.Payload
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetApplicationTagSignedInteger) GetActualValue() uint64 {
	return uint64(m.GetPayload().GetActualValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetApplicationTagSignedInteger factory function for _BACnetApplicationTagSignedInteger
func NewBACnetApplicationTagSignedInteger(payload BACnetTagPayloadSignedInteger, header BACnetTagHeader) *_BACnetApplicationTagSignedInteger {
	_result := &_BACnetApplicationTagSignedInteger{
		Payload:               payload,
		_BACnetApplicationTag: NewBACnetApplicationTag(header),
	}
	_result._BACnetApplicationTag._BACnetApplicationTagChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetApplicationTagSignedInteger(structType interface{}) BACnetApplicationTagSignedInteger {
	if casted, ok := structType.(BACnetApplicationTagSignedInteger); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetApplicationTagSignedInteger); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetApplicationTagSignedInteger) GetTypeName() string {
	return "BACnetApplicationTagSignedInteger"
}

func (m *_BACnetApplicationTagSignedInteger) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetApplicationTagSignedInteger) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetApplicationTagSignedInteger) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetApplicationTagSignedIntegerParse(readBuffer utils.ReadBuffer, header BACnetTagHeader) (BACnetApplicationTagSignedInteger, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetApplicationTagSignedInteger"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetApplicationTagSignedInteger")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (payload)
	if pullErr := readBuffer.PullContext("payload"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for payload")
	}
	_payload, _payloadErr := BACnetTagPayloadSignedIntegerParse(readBuffer, uint32(header.GetActualLength()))
	if _payloadErr != nil {
		return nil, errors.Wrap(_payloadErr, "Error parsing 'payload' field of BACnetApplicationTagSignedInteger")
	}
	payload := _payload.(BACnetTagPayloadSignedInteger)
	if closeErr := readBuffer.CloseContext("payload"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for payload")
	}

	// Virtual field
	_actualValue := payload.GetActualValue()
	actualValue := uint64(_actualValue)
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetApplicationTagSignedInteger"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetApplicationTagSignedInteger")
	}

	// Create a partially initialized instance
	_child := &_BACnetApplicationTagSignedInteger{
		_BACnetApplicationTag: &_BACnetApplicationTag{},
		Payload:               payload,
	}
	_child._BACnetApplicationTag._BACnetApplicationTagChildRequirements = _child
	return _child, nil
}

func (m *_BACnetApplicationTagSignedInteger) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetApplicationTagSignedInteger) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetApplicationTagSignedInteger"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetApplicationTagSignedInteger")
		}

		// Simple Field (payload)
		if pushErr := writeBuffer.PushContext("payload"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for payload")
		}
		_payloadErr := writeBuffer.WriteSerializable(m.GetPayload())
		if popErr := writeBuffer.PopContext("payload"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for payload")
		}
		if _payloadErr != nil {
			return errors.Wrap(_payloadErr, "Error serializing 'payload' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetApplicationTagSignedInteger"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetApplicationTagSignedInteger")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetApplicationTagSignedInteger) isBACnetApplicationTagSignedInteger() bool {
	return true
}

func (m *_BACnetApplicationTagSignedInteger) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
