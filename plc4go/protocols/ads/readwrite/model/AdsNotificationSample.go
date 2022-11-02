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

// AdsNotificationSample is the corresponding interface of AdsNotificationSample
type AdsNotificationSample interface {
	utils.LengthAware
	utils.Serializable
	// GetNotificationHandle returns NotificationHandle (property field)
	GetNotificationHandle() uint32
	// GetSampleSize returns SampleSize (property field)
	GetSampleSize() uint32
	// GetData returns Data (property field)
	GetData() []byte
}

// AdsNotificationSampleExactly can be used when we want exactly this type and not a type which fulfills AdsNotificationSample.
// This is useful for switch cases.
type AdsNotificationSampleExactly interface {
	AdsNotificationSample
	isAdsNotificationSample() bool
}

// _AdsNotificationSample is the data-structure of this message
type _AdsNotificationSample struct {
	NotificationHandle uint32
	SampleSize         uint32
	Data               []byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AdsNotificationSample) GetNotificationHandle() uint32 {
	return m.NotificationHandle
}

func (m *_AdsNotificationSample) GetSampleSize() uint32 {
	return m.SampleSize
}

func (m *_AdsNotificationSample) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAdsNotificationSample factory function for _AdsNotificationSample
func NewAdsNotificationSample(notificationHandle uint32, sampleSize uint32, data []byte) *_AdsNotificationSample {
	return &_AdsNotificationSample{NotificationHandle: notificationHandle, SampleSize: sampleSize, Data: data}
}

// Deprecated: use the interface for direct cast
func CastAdsNotificationSample(structType interface{}) AdsNotificationSample {
	if casted, ok := structType.(AdsNotificationSample); ok {
		return casted
	}
	if casted, ok := structType.(*AdsNotificationSample); ok {
		return *casted
	}
	return nil
}

func (m *_AdsNotificationSample) GetTypeName() string {
	return "AdsNotificationSample"
}

func (m *_AdsNotificationSample) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_AdsNotificationSample) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (notificationHandle)
	lengthInBits += 32

	// Simple field (sampleSize)
	lengthInBits += 32

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_AdsNotificationSample) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func AdsNotificationSampleParse(readBuffer utils.ReadBuffer) (AdsNotificationSample, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsNotificationSample"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsNotificationSample")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (notificationHandle)
	_notificationHandle, _notificationHandleErr := readBuffer.ReadUint32("notificationHandle", 32)
	if _notificationHandleErr != nil {
		return nil, errors.Wrap(_notificationHandleErr, "Error parsing 'notificationHandle' field of AdsNotificationSample")
	}
	notificationHandle := _notificationHandle

	// Simple Field (sampleSize)
	_sampleSize, _sampleSizeErr := readBuffer.ReadUint32("sampleSize", 32)
	if _sampleSizeErr != nil {
		return nil, errors.Wrap(_sampleSizeErr, "Error parsing 'sampleSize' field of AdsNotificationSample")
	}
	sampleSize := _sampleSize
	// Byte Array field (data)
	numberOfBytesdata := int(sampleSize)
	data, _readArrayErr := readBuffer.ReadByteArray("data", numberOfBytesdata)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'data' field of AdsNotificationSample")
	}

	if closeErr := readBuffer.CloseContext("AdsNotificationSample"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsNotificationSample")
	}

	// Create the instance
	return &_AdsNotificationSample{
		NotificationHandle: notificationHandle,
		SampleSize:         sampleSize,
		Data:               data,
	}, nil
}

func (m *_AdsNotificationSample) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AdsNotificationSample) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("AdsNotificationSample"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for AdsNotificationSample")
	}

	// Simple Field (notificationHandle)
	notificationHandle := uint32(m.GetNotificationHandle())
	_notificationHandleErr := writeBuffer.WriteUint32("notificationHandle", 32, (notificationHandle))
	if _notificationHandleErr != nil {
		return errors.Wrap(_notificationHandleErr, "Error serializing 'notificationHandle' field")
	}

	// Simple Field (sampleSize)
	sampleSize := uint32(m.GetSampleSize())
	_sampleSizeErr := writeBuffer.WriteUint32("sampleSize", 32, (sampleSize))
	if _sampleSizeErr != nil {
		return errors.Wrap(_sampleSizeErr, "Error serializing 'sampleSize' field")
	}

	// Array Field (data)
	// Byte Array field (data)
	if err := writeBuffer.WriteByteArray("data", m.GetData()); err != nil {
		return errors.Wrap(err, "Error serializing 'data' field")
	}

	if popErr := writeBuffer.PopContext("AdsNotificationSample"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for AdsNotificationSample")
	}
	return nil
}

func (m *_AdsNotificationSample) isAdsNotificationSample() bool {
	return true
}

func (m *_AdsNotificationSample) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
