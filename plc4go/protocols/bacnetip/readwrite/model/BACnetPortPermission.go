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

// BACnetPortPermission is the corresponding interface of BACnetPortPermission
type BACnetPortPermission interface {
	utils.LengthAware
	utils.Serializable
	// GetPort returns Port (property field)
	GetPort() BACnetContextTagUnsignedInteger
	// GetEnable returns Enable (property field)
	GetEnable() BACnetContextTagBoolean
}

// BACnetPortPermissionExactly can be used when we want exactly this type and not a type which fulfills BACnetPortPermission.
// This is useful for switch cases.
type BACnetPortPermissionExactly interface {
	BACnetPortPermission
	isBACnetPortPermission() bool
}

// _BACnetPortPermission is the data-structure of this message
type _BACnetPortPermission struct {
	Port   BACnetContextTagUnsignedInteger
	Enable BACnetContextTagBoolean
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPortPermission) GetPort() BACnetContextTagUnsignedInteger {
	return m.Port
}

func (m *_BACnetPortPermission) GetEnable() BACnetContextTagBoolean {
	return m.Enable
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPortPermission factory function for _BACnetPortPermission
func NewBACnetPortPermission(port BACnetContextTagUnsignedInteger, enable BACnetContextTagBoolean) *_BACnetPortPermission {
	return &_BACnetPortPermission{Port: port, Enable: enable}
}

// Deprecated: use the interface for direct cast
func CastBACnetPortPermission(structType interface{}) BACnetPortPermission {
	if casted, ok := structType.(BACnetPortPermission); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPortPermission); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPortPermission) GetTypeName() string {
	return "BACnetPortPermission"
}

func (m *_BACnetPortPermission) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetPortPermission) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (port)
	lengthInBits += m.Port.GetLengthInBits()

	// Optional Field (enable)
	if m.Enable != nil {
		lengthInBits += m.Enable.GetLengthInBits()
	}

	return lengthInBits
}

func (m *_BACnetPortPermission) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPortPermissionParse(readBuffer utils.ReadBuffer) (BACnetPortPermission, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPortPermission"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPortPermission")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (port)
	if pullErr := readBuffer.PullContext("port"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for port")
	}
	_port, _portErr := BACnetContextTagParse(readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _portErr != nil {
		return nil, errors.Wrap(_portErr, "Error parsing 'port' field of BACnetPortPermission")
	}
	port := _port.(BACnetContextTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("port"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for port")
	}

	// Optional Field (enable) (Can be skipped, if a given expression evaluates to false)
	var enable BACnetContextTagBoolean = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("enable"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for enable")
		}
		_val, _err := BACnetContextTagParse(readBuffer, uint8(1), BACnetDataType_BOOLEAN)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			Plc4xModelLog.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'enable' field of BACnetPortPermission")
		default:
			enable = _val.(BACnetContextTagBoolean)
			if closeErr := readBuffer.CloseContext("enable"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for enable")
			}
		}
	}

	if closeErr := readBuffer.CloseContext("BACnetPortPermission"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPortPermission")
	}

	// Create the instance
	return &_BACnetPortPermission{
		Port:   port,
		Enable: enable,
	}, nil
}

func (m *_BACnetPortPermission) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPortPermission) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetPortPermission"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetPortPermission")
	}

	// Simple Field (port)
	if pushErr := writeBuffer.PushContext("port"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for port")
	}
	_portErr := writeBuffer.WriteSerializable(m.GetPort())
	if popErr := writeBuffer.PopContext("port"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for port")
	}
	if _portErr != nil {
		return errors.Wrap(_portErr, "Error serializing 'port' field")
	}

	// Optional Field (enable) (Can be skipped, if the value is null)
	var enable BACnetContextTagBoolean = nil
	if m.GetEnable() != nil {
		if pushErr := writeBuffer.PushContext("enable"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for enable")
		}
		enable = m.GetEnable()
		_enableErr := writeBuffer.WriteSerializable(enable)
		if popErr := writeBuffer.PopContext("enable"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for enable")
		}
		if _enableErr != nil {
			return errors.Wrap(_enableErr, "Error serializing 'enable' field")
		}
	}

	if popErr := writeBuffer.PopContext("BACnetPortPermission"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetPortPermission")
	}
	return nil
}

func (m *_BACnetPortPermission) isBACnetPortPermission() bool {
	return true
}

func (m *_BACnetPortPermission) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
