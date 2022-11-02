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

// BACnetConstructedDataStopWhenFull is the corresponding interface of BACnetConstructedDataStopWhenFull
type BACnetConstructedDataStopWhenFull interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetStopWhenFull returns StopWhenFull (property field)
	GetStopWhenFull() BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagBoolean
}

// BACnetConstructedDataStopWhenFullExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataStopWhenFull.
// This is useful for switch cases.
type BACnetConstructedDataStopWhenFullExactly interface {
	BACnetConstructedDataStopWhenFull
	isBACnetConstructedDataStopWhenFull() bool
}

// _BACnetConstructedDataStopWhenFull is the data-structure of this message
type _BACnetConstructedDataStopWhenFull struct {
	*_BACnetConstructedData
	StopWhenFull BACnetApplicationTagBoolean
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataStopWhenFull) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataStopWhenFull) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_STOP_WHEN_FULL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataStopWhenFull) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataStopWhenFull) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataStopWhenFull) GetStopWhenFull() BACnetApplicationTagBoolean {
	return m.StopWhenFull
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataStopWhenFull) GetActualValue() BACnetApplicationTagBoolean {
	return CastBACnetApplicationTagBoolean(m.GetStopWhenFull())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataStopWhenFull factory function for _BACnetConstructedDataStopWhenFull
func NewBACnetConstructedDataStopWhenFull(stopWhenFull BACnetApplicationTagBoolean, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataStopWhenFull {
	_result := &_BACnetConstructedDataStopWhenFull{
		StopWhenFull:           stopWhenFull,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataStopWhenFull(structType interface{}) BACnetConstructedDataStopWhenFull {
	if casted, ok := structType.(BACnetConstructedDataStopWhenFull); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataStopWhenFull); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataStopWhenFull) GetTypeName() string {
	return "BACnetConstructedDataStopWhenFull"
}

func (m *_BACnetConstructedDataStopWhenFull) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataStopWhenFull) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (stopWhenFull)
	lengthInBits += m.StopWhenFull.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataStopWhenFull) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataStopWhenFullParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataStopWhenFull, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataStopWhenFull"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataStopWhenFull")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (stopWhenFull)
	if pullErr := readBuffer.PullContext("stopWhenFull"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for stopWhenFull")
	}
	_stopWhenFull, _stopWhenFullErr := BACnetApplicationTagParse(readBuffer)
	if _stopWhenFullErr != nil {
		return nil, errors.Wrap(_stopWhenFullErr, "Error parsing 'stopWhenFull' field of BACnetConstructedDataStopWhenFull")
	}
	stopWhenFull := _stopWhenFull.(BACnetApplicationTagBoolean)
	if closeErr := readBuffer.CloseContext("stopWhenFull"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for stopWhenFull")
	}

	// Virtual field
	_actualValue := stopWhenFull
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataStopWhenFull"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataStopWhenFull")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataStopWhenFull{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		StopWhenFull: stopWhenFull,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataStopWhenFull) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataStopWhenFull) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataStopWhenFull"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataStopWhenFull")
		}

		// Simple Field (stopWhenFull)
		if pushErr := writeBuffer.PushContext("stopWhenFull"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for stopWhenFull")
		}
		_stopWhenFullErr := writeBuffer.WriteSerializable(m.GetStopWhenFull())
		if popErr := writeBuffer.PopContext("stopWhenFull"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for stopWhenFull")
		}
		if _stopWhenFullErr != nil {
			return errors.Wrap(_stopWhenFullErr, "Error serializing 'stopWhenFull' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataStopWhenFull"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataStopWhenFull")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataStopWhenFull) isBACnetConstructedDataStopWhenFull() bool {
	return true
}

func (m *_BACnetConstructedDataStopWhenFull) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
