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

// BVLCDeleteForeignDeviceTableEntry is the corresponding interface of BVLCDeleteForeignDeviceTableEntry
type BVLCDeleteForeignDeviceTableEntry interface {
	utils.LengthAware
	utils.Serializable
	BVLC
	// GetIp returns Ip (property field)
	GetIp() []uint8
	// GetPort returns Port (property field)
	GetPort() uint16
}

// BVLCDeleteForeignDeviceTableEntryExactly can be used when we want exactly this type and not a type which fulfills BVLCDeleteForeignDeviceTableEntry.
// This is useful for switch cases.
type BVLCDeleteForeignDeviceTableEntryExactly interface {
	BVLCDeleteForeignDeviceTableEntry
	isBVLCDeleteForeignDeviceTableEntry() bool
}

// _BVLCDeleteForeignDeviceTableEntry is the data-structure of this message
type _BVLCDeleteForeignDeviceTableEntry struct {
	*_BVLC
	Ip   []uint8
	Port uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BVLCDeleteForeignDeviceTableEntry) GetBvlcFunction() uint8 {
	return 0x08
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BVLCDeleteForeignDeviceTableEntry) InitializeParent(parent BVLC) {}

func (m *_BVLCDeleteForeignDeviceTableEntry) GetParent() BVLC {
	return m._BVLC
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BVLCDeleteForeignDeviceTableEntry) GetIp() []uint8 {
	return m.Ip
}

func (m *_BVLCDeleteForeignDeviceTableEntry) GetPort() uint16 {
	return m.Port
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBVLCDeleteForeignDeviceTableEntry factory function for _BVLCDeleteForeignDeviceTableEntry
func NewBVLCDeleteForeignDeviceTableEntry(ip []uint8, port uint16) *_BVLCDeleteForeignDeviceTableEntry {
	_result := &_BVLCDeleteForeignDeviceTableEntry{
		Ip:    ip,
		Port:  port,
		_BVLC: NewBVLC(),
	}
	_result._BVLC._BVLCChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBVLCDeleteForeignDeviceTableEntry(structType interface{}) BVLCDeleteForeignDeviceTableEntry {
	if casted, ok := structType.(BVLCDeleteForeignDeviceTableEntry); ok {
		return casted
	}
	if casted, ok := structType.(*BVLCDeleteForeignDeviceTableEntry); ok {
		return *casted
	}
	return nil
}

func (m *_BVLCDeleteForeignDeviceTableEntry) GetTypeName() string {
	return "BVLCDeleteForeignDeviceTableEntry"
}

func (m *_BVLCDeleteForeignDeviceTableEntry) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BVLCDeleteForeignDeviceTableEntry) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.Ip) > 0 {
		lengthInBits += 8 * uint16(len(m.Ip))
	}

	// Simple field (port)
	lengthInBits += 16

	return lengthInBits
}

func (m *_BVLCDeleteForeignDeviceTableEntry) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BVLCDeleteForeignDeviceTableEntryParse(readBuffer utils.ReadBuffer) (BVLCDeleteForeignDeviceTableEntry, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BVLCDeleteForeignDeviceTableEntry"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BVLCDeleteForeignDeviceTableEntry")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (ip)
	if pullErr := readBuffer.PullContext("ip", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ip")
	}
	// Count array
	ip := make([]uint8, uint16(4))
	// This happens when the size is set conditional to 0
	if len(ip) == 0 {
		ip = nil
	}
	{
		for curItem := uint16(0); curItem < uint16(uint16(4)); curItem++ {
			_item, _err := readBuffer.ReadUint8("", 8)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'ip' field of BVLCDeleteForeignDeviceTableEntry")
			}
			ip[curItem] = _item
		}
	}
	if closeErr := readBuffer.CloseContext("ip", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ip")
	}

	// Simple Field (port)
	_port, _portErr := readBuffer.ReadUint16("port", 16)
	if _portErr != nil {
		return nil, errors.Wrap(_portErr, "Error parsing 'port' field of BVLCDeleteForeignDeviceTableEntry")
	}
	port := _port

	if closeErr := readBuffer.CloseContext("BVLCDeleteForeignDeviceTableEntry"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BVLCDeleteForeignDeviceTableEntry")
	}

	// Create a partially initialized instance
	_child := &_BVLCDeleteForeignDeviceTableEntry{
		_BVLC: &_BVLC{},
		Ip:    ip,
		Port:  port,
	}
	_child._BVLC._BVLCChildRequirements = _child
	return _child, nil
}

func (m *_BVLCDeleteForeignDeviceTableEntry) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BVLCDeleteForeignDeviceTableEntry) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BVLCDeleteForeignDeviceTableEntry"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BVLCDeleteForeignDeviceTableEntry")
		}

		// Array Field (ip)
		if pushErr := writeBuffer.PushContext("ip", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ip")
		}
		for _, _element := range m.GetIp() {
			_elementErr := writeBuffer.WriteUint8("", 8, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'ip' field")
			}
		}
		if popErr := writeBuffer.PopContext("ip", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ip")
		}

		// Simple Field (port)
		port := uint16(m.GetPort())
		_portErr := writeBuffer.WriteUint16("port", 16, (port))
		if _portErr != nil {
			return errors.Wrap(_portErr, "Error serializing 'port' field")
		}

		if popErr := writeBuffer.PopContext("BVLCDeleteForeignDeviceTableEntry"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BVLCDeleteForeignDeviceTableEntry")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BVLCDeleteForeignDeviceTableEntry) isBVLCDeleteForeignDeviceTableEntry() bool {
	return true
}

func (m *_BVLCDeleteForeignDeviceTableEntry) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
