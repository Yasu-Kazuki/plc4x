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

// ApduDataDeviceDescriptorResponse is the corresponding interface of ApduDataDeviceDescriptorResponse
type ApduDataDeviceDescriptorResponse interface {
	utils.LengthAware
	utils.Serializable
	ApduData
	// GetDescriptorType returns DescriptorType (property field)
	GetDescriptorType() uint8
	// GetData returns Data (property field)
	GetData() []byte
}

// ApduDataDeviceDescriptorResponseExactly can be used when we want exactly this type and not a type which fulfills ApduDataDeviceDescriptorResponse.
// This is useful for switch cases.
type ApduDataDeviceDescriptorResponseExactly interface {
	ApduDataDeviceDescriptorResponse
	isApduDataDeviceDescriptorResponse() bool
}

// _ApduDataDeviceDescriptorResponse is the data-structure of this message
type _ApduDataDeviceDescriptorResponse struct {
	*_ApduData
	DescriptorType uint8
	Data           []byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataDeviceDescriptorResponse) GetApciType() uint8 {
	return 0xD
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataDeviceDescriptorResponse) InitializeParent(parent ApduData) {}

func (m *_ApduDataDeviceDescriptorResponse) GetParent() ApduData {
	return m._ApduData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ApduDataDeviceDescriptorResponse) GetDescriptorType() uint8 {
	return m.DescriptorType
}

func (m *_ApduDataDeviceDescriptorResponse) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewApduDataDeviceDescriptorResponse factory function for _ApduDataDeviceDescriptorResponse
func NewApduDataDeviceDescriptorResponse(descriptorType uint8, data []byte, dataLength uint8) *_ApduDataDeviceDescriptorResponse {
	_result := &_ApduDataDeviceDescriptorResponse{
		DescriptorType: descriptorType,
		Data:           data,
		_ApduData:      NewApduData(dataLength),
	}
	_result._ApduData._ApduDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastApduDataDeviceDescriptorResponse(structType interface{}) ApduDataDeviceDescriptorResponse {
	if casted, ok := structType.(ApduDataDeviceDescriptorResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataDeviceDescriptorResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataDeviceDescriptorResponse) GetTypeName() string {
	return "ApduDataDeviceDescriptorResponse"
}

func (m *_ApduDataDeviceDescriptorResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ApduDataDeviceDescriptorResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (descriptorType)
	lengthInBits += 6

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_ApduDataDeviceDescriptorResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApduDataDeviceDescriptorResponseParse(readBuffer utils.ReadBuffer, dataLength uint8) (ApduDataDeviceDescriptorResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataDeviceDescriptorResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataDeviceDescriptorResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (descriptorType)
	_descriptorType, _descriptorTypeErr := readBuffer.ReadUint8("descriptorType", 6)
	if _descriptorTypeErr != nil {
		return nil, errors.Wrap(_descriptorTypeErr, "Error parsing 'descriptorType' field of ApduDataDeviceDescriptorResponse")
	}
	descriptorType := _descriptorType
	// Byte Array field (data)
	numberOfBytesdata := int(utils.InlineIf((bool((dataLength) < (1))), func() interface{} { return uint16(uint16(0)) }, func() interface{} { return uint16(uint16(dataLength) - uint16(uint16(1))) }).(uint16))
	data, _readArrayErr := readBuffer.ReadByteArray("data", numberOfBytesdata)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'data' field of ApduDataDeviceDescriptorResponse")
	}

	if closeErr := readBuffer.CloseContext("ApduDataDeviceDescriptorResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataDeviceDescriptorResponse")
	}

	// Create a partially initialized instance
	_child := &_ApduDataDeviceDescriptorResponse{
		_ApduData: &_ApduData{
			DataLength: dataLength,
		},
		DescriptorType: descriptorType,
		Data:           data,
	}
	_child._ApduData._ApduDataChildRequirements = _child
	return _child, nil
}

func (m *_ApduDataDeviceDescriptorResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataDeviceDescriptorResponse) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataDeviceDescriptorResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataDeviceDescriptorResponse")
		}

		// Simple Field (descriptorType)
		descriptorType := uint8(m.GetDescriptorType())
		_descriptorTypeErr := writeBuffer.WriteUint8("descriptorType", 6, (descriptorType))
		if _descriptorTypeErr != nil {
			return errors.Wrap(_descriptorTypeErr, "Error serializing 'descriptorType' field")
		}

		// Array Field (data)
		// Byte Array field (data)
		if err := writeBuffer.WriteByteArray("data", m.GetData()); err != nil {
			return errors.Wrap(err, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("ApduDataDeviceDescriptorResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataDeviceDescriptorResponse")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_ApduDataDeviceDescriptorResponse) isApduDataDeviceDescriptorResponse() bool {
	return true
}

func (m *_ApduDataDeviceDescriptorResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
