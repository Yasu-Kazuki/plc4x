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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// TunnelingRequest is the corresponding interface of TunnelingRequest
type TunnelingRequest interface {
	utils.LengthAware
	utils.Serializable
	KnxNetIpMessage
	// GetTunnelingRequestDataBlock returns TunnelingRequestDataBlock (property field)
	GetTunnelingRequestDataBlock() TunnelingRequestDataBlock
	// GetCemi returns Cemi (property field)
	GetCemi() CEMI
}

// TunnelingRequestExactly can be used when we want exactly this type and not a type which fulfills TunnelingRequest.
// This is useful for switch cases.
type TunnelingRequestExactly interface {
	TunnelingRequest
	isTunnelingRequest() bool
}

// _TunnelingRequest is the data-structure of this message
type _TunnelingRequest struct {
	*_KnxNetIpMessage
	TunnelingRequestDataBlock TunnelingRequestDataBlock
	Cemi                      CEMI

	// Arguments.
	TotalLength uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_TunnelingRequest) GetMsgType() uint16 {
	return 0x0420
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_TunnelingRequest) InitializeParent(parent KnxNetIpMessage) {}

func (m *_TunnelingRequest) GetParent() KnxNetIpMessage {
	return m._KnxNetIpMessage
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_TunnelingRequest) GetTunnelingRequestDataBlock() TunnelingRequestDataBlock {
	return m.TunnelingRequestDataBlock
}

func (m *_TunnelingRequest) GetCemi() CEMI {
	return m.Cemi
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewTunnelingRequest factory function for _TunnelingRequest
func NewTunnelingRequest(tunnelingRequestDataBlock TunnelingRequestDataBlock, cemi CEMI, totalLength uint16) *_TunnelingRequest {
	_result := &_TunnelingRequest{
		TunnelingRequestDataBlock: tunnelingRequestDataBlock,
		Cemi:                      cemi,
		_KnxNetIpMessage:          NewKnxNetIpMessage(),
	}
	_result._KnxNetIpMessage._KnxNetIpMessageChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastTunnelingRequest(structType interface{}) TunnelingRequest {
	if casted, ok := structType.(TunnelingRequest); ok {
		return casted
	}
	if casted, ok := structType.(*TunnelingRequest); ok {
		return *casted
	}
	return nil
}

func (m *_TunnelingRequest) GetTypeName() string {
	return "TunnelingRequest"
}

func (m *_TunnelingRequest) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_TunnelingRequest) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (tunnelingRequestDataBlock)
	lengthInBits += m.TunnelingRequestDataBlock.GetLengthInBits()

	// Simple field (cemi)
	lengthInBits += m.Cemi.GetLengthInBits()

	return lengthInBits
}

func (m *_TunnelingRequest) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func TunnelingRequestParse(readBuffer utils.ReadBuffer, totalLength uint16) (TunnelingRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("TunnelingRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TunnelingRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (tunnelingRequestDataBlock)
	if pullErr := readBuffer.PullContext("tunnelingRequestDataBlock"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for tunnelingRequestDataBlock")
	}
	_tunnelingRequestDataBlock, _tunnelingRequestDataBlockErr := TunnelingRequestDataBlockParse(readBuffer)
	if _tunnelingRequestDataBlockErr != nil {
		return nil, errors.Wrap(_tunnelingRequestDataBlockErr, "Error parsing 'tunnelingRequestDataBlock' field of TunnelingRequest")
	}
	tunnelingRequestDataBlock := _tunnelingRequestDataBlock.(TunnelingRequestDataBlock)
	if closeErr := readBuffer.CloseContext("tunnelingRequestDataBlock"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for tunnelingRequestDataBlock")
	}

	// Simple Field (cemi)
	if pullErr := readBuffer.PullContext("cemi"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for cemi")
	}
	_cemi, _cemiErr := CEMIParse(readBuffer, uint16(uint16(totalLength)-uint16((uint16(uint16(6))+uint16(tunnelingRequestDataBlock.GetLengthInBytes())))))
	if _cemiErr != nil {
		return nil, errors.Wrap(_cemiErr, "Error parsing 'cemi' field of TunnelingRequest")
	}
	cemi := _cemi.(CEMI)
	if closeErr := readBuffer.CloseContext("cemi"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for cemi")
	}

	if closeErr := readBuffer.CloseContext("TunnelingRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TunnelingRequest")
	}

	// Create a partially initialized instance
	_child := &_TunnelingRequest{
		TunnelingRequestDataBlock: tunnelingRequestDataBlock,
		Cemi:                      cemi,
		_KnxNetIpMessage:          &_KnxNetIpMessage{},
	}
	_child._KnxNetIpMessage._KnxNetIpMessageChildRequirements = _child
	return _child, nil
}

func (m *_TunnelingRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("TunnelingRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for TunnelingRequest")
		}

		// Simple Field (tunnelingRequestDataBlock)
		if pushErr := writeBuffer.PushContext("tunnelingRequestDataBlock"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for tunnelingRequestDataBlock")
		}
		_tunnelingRequestDataBlockErr := writeBuffer.WriteSerializable(m.GetTunnelingRequestDataBlock())
		if popErr := writeBuffer.PopContext("tunnelingRequestDataBlock"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for tunnelingRequestDataBlock")
		}
		if _tunnelingRequestDataBlockErr != nil {
			return errors.Wrap(_tunnelingRequestDataBlockErr, "Error serializing 'tunnelingRequestDataBlock' field")
		}

		// Simple Field (cemi)
		if pushErr := writeBuffer.PushContext("cemi"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for cemi")
		}
		_cemiErr := writeBuffer.WriteSerializable(m.GetCemi())
		if popErr := writeBuffer.PopContext("cemi"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for cemi")
		}
		if _cemiErr != nil {
			return errors.Wrap(_cemiErr, "Error serializing 'cemi' field")
		}

		if popErr := writeBuffer.PopContext("TunnelingRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for TunnelingRequest")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

////
// Arguments Getter

func (m *_TunnelingRequest) GetTotalLength() uint16 {
	return m.TotalLength
}

//
////

func (m *_TunnelingRequest) isTunnelingRequest() bool {
	return true
}

func (m *_TunnelingRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
