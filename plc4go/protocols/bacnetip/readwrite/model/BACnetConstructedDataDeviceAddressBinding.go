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

// BACnetConstructedDataDeviceAddressBinding is the corresponding interface of BACnetConstructedDataDeviceAddressBinding
type BACnetConstructedDataDeviceAddressBinding interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetDeviceAddressBinding returns DeviceAddressBinding (property field)
	GetDeviceAddressBinding() []BACnetAddressBinding
}

// BACnetConstructedDataDeviceAddressBindingExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataDeviceAddressBinding.
// This is useful for switch cases.
type BACnetConstructedDataDeviceAddressBindingExactly interface {
	BACnetConstructedDataDeviceAddressBinding
	isBACnetConstructedDataDeviceAddressBinding() bool
}

// _BACnetConstructedDataDeviceAddressBinding is the data-structure of this message
type _BACnetConstructedDataDeviceAddressBinding struct {
	*_BACnetConstructedData
	DeviceAddressBinding []BACnetAddressBinding
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataDeviceAddressBinding) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataDeviceAddressBinding) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_DEVICE_ADDRESS_BINDING
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataDeviceAddressBinding) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataDeviceAddressBinding) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataDeviceAddressBinding) GetDeviceAddressBinding() []BACnetAddressBinding {
	return m.DeviceAddressBinding
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataDeviceAddressBinding factory function for _BACnetConstructedDataDeviceAddressBinding
func NewBACnetConstructedDataDeviceAddressBinding(deviceAddressBinding []BACnetAddressBinding, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataDeviceAddressBinding {
	_result := &_BACnetConstructedDataDeviceAddressBinding{
		DeviceAddressBinding:   deviceAddressBinding,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataDeviceAddressBinding(structType interface{}) BACnetConstructedDataDeviceAddressBinding {
	if casted, ok := structType.(BACnetConstructedDataDeviceAddressBinding); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDeviceAddressBinding); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataDeviceAddressBinding) GetTypeName() string {
	return "BACnetConstructedDataDeviceAddressBinding"
}

func (m *_BACnetConstructedDataDeviceAddressBinding) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataDeviceAddressBinding) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.DeviceAddressBinding) > 0 {
		for _, element := range m.DeviceAddressBinding {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataDeviceAddressBinding) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataDeviceAddressBindingParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataDeviceAddressBinding, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDeviceAddressBinding"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataDeviceAddressBinding")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (deviceAddressBinding)
	if pullErr := readBuffer.PullContext("deviceAddressBinding", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for deviceAddressBinding")
	}
	// Terminated array
	var deviceAddressBinding []BACnetAddressBinding
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetAddressBindingParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'deviceAddressBinding' field of BACnetConstructedDataDeviceAddressBinding")
			}
			deviceAddressBinding = append(deviceAddressBinding, _item.(BACnetAddressBinding))

		}
	}
	if closeErr := readBuffer.CloseContext("deviceAddressBinding", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for deviceAddressBinding")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDeviceAddressBinding"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDeviceAddressBinding")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataDeviceAddressBinding{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		DeviceAddressBinding: deviceAddressBinding,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataDeviceAddressBinding) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataDeviceAddressBinding) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDeviceAddressBinding"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataDeviceAddressBinding")
		}

		// Array Field (deviceAddressBinding)
		if pushErr := writeBuffer.PushContext("deviceAddressBinding", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for deviceAddressBinding")
		}
		for _, _element := range m.GetDeviceAddressBinding() {
			_elementErr := writeBuffer.WriteSerializable(_element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'deviceAddressBinding' field")
			}
		}
		if popErr := writeBuffer.PopContext("deviceAddressBinding", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for deviceAddressBinding")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDeviceAddressBinding"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataDeviceAddressBinding")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataDeviceAddressBinding) isBACnetConstructedDataDeviceAddressBinding() bool {
	return true
}

func (m *_BACnetConstructedDataDeviceAddressBinding) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
