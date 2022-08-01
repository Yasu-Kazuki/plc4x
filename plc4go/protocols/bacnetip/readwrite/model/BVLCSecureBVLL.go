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

// BVLCSecureBVLL is the corresponding interface of BVLCSecureBVLL
type BVLCSecureBVLL interface {
	utils.LengthAware
	utils.Serializable
	BVLC
	// GetSecurityWrapper returns SecurityWrapper (property field)
	GetSecurityWrapper() []byte
}

// BVLCSecureBVLLExactly can be used when we want exactly this type and not a type which fulfills BVLCSecureBVLL.
// This is useful for switch cases.
type BVLCSecureBVLLExactly interface {
	BVLCSecureBVLL
	isBVLCSecureBVLL() bool
}

// _BVLCSecureBVLL is the data-structure of this message
type _BVLCSecureBVLL struct {
	*_BVLC
	SecurityWrapper []byte

	// Arguments.
	BvlcPayloadLength uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BVLCSecureBVLL) GetBvlcFunction() uint8 {
	return 0x0C
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BVLCSecureBVLL) InitializeParent(parent BVLC) {}

func (m *_BVLCSecureBVLL) GetParent() BVLC {
	return m._BVLC
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BVLCSecureBVLL) GetSecurityWrapper() []byte {
	return m.SecurityWrapper
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBVLCSecureBVLL factory function for _BVLCSecureBVLL
func NewBVLCSecureBVLL(securityWrapper []byte, bvlcPayloadLength uint16) *_BVLCSecureBVLL {
	_result := &_BVLCSecureBVLL{
		SecurityWrapper: securityWrapper,
		_BVLC:           NewBVLC(),
	}
	_result._BVLC._BVLCChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBVLCSecureBVLL(structType interface{}) BVLCSecureBVLL {
	if casted, ok := structType.(BVLCSecureBVLL); ok {
		return casted
	}
	if casted, ok := structType.(*BVLCSecureBVLL); ok {
		return *casted
	}
	return nil
}

func (m *_BVLCSecureBVLL) GetTypeName() string {
	return "BVLCSecureBVLL"
}

func (m *_BVLCSecureBVLL) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BVLCSecureBVLL) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.SecurityWrapper) > 0 {
		lengthInBits += 8 * uint16(len(m.SecurityWrapper))
	}

	return lengthInBits
}

func (m *_BVLCSecureBVLL) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BVLCSecureBVLLParse(readBuffer utils.ReadBuffer, bvlcPayloadLength uint16) (BVLCSecureBVLL, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BVLCSecureBVLL"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BVLCSecureBVLL")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos
	// Byte Array field (securityWrapper)
	numberOfBytessecurityWrapper := int(bvlcPayloadLength)
	securityWrapper, _readArrayErr := readBuffer.ReadByteArray("securityWrapper", numberOfBytessecurityWrapper)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'securityWrapper' field of BVLCSecureBVLL")
	}

	if closeErr := readBuffer.CloseContext("BVLCSecureBVLL"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BVLCSecureBVLL")
	}

	// Create a partially initialized instance
	_child := &_BVLCSecureBVLL{
		SecurityWrapper: securityWrapper,
		_BVLC:           &_BVLC{},
	}
	_child._BVLC._BVLCChildRequirements = _child
	return _child, nil
}

func (m *_BVLCSecureBVLL) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BVLCSecureBVLL"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BVLCSecureBVLL")
		}

		// Array Field (securityWrapper)
		// Byte Array field (securityWrapper)
		if err := writeBuffer.WriteByteArray("securityWrapper", m.GetSecurityWrapper()); err != nil {
			return errors.Wrap(err, "Error serializing 'securityWrapper' field")
		}

		if popErr := writeBuffer.PopContext("BVLCSecureBVLL"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BVLCSecureBVLL")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

////
// Arguments Getter

func (m *_BVLCSecureBVLL) GetBvlcPayloadLength() uint16 {
	return m.BvlcPayloadLength
}

//
////

func (m *_BVLCSecureBVLL) isBVLCSecureBVLL() bool {
	return true
}

func (m *_BVLCSecureBVLL) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
