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

// NLM is the corresponding interface of NLM
type NLM interface {
	utils.LengthAware
	utils.Serializable
	// GetMessageType returns MessageType (discriminator field)
	GetMessageType() uint8
	// GetVendorId returns VendorId (property field)
	GetVendorId() *BACnetVendorId
}

// NLMExactly can be used when we want exactly this type and not a type which fulfills NLM.
// This is useful for switch cases.
type NLMExactly interface {
	NLM
	isNLM() bool
}

// _NLM is the data-structure of this message
type _NLM struct {
	_NLMChildRequirements
	VendorId *BACnetVendorId

	// Arguments.
	ApduLength uint16
}

type _NLMChildRequirements interface {
	utils.Serializable
	GetLengthInBits() uint16
	GetLengthInBitsConditional(lastItem bool) uint16
	GetMessageType() uint8
}

type NLMParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child NLM, serializeChildFunction func() error) error
	GetTypeName() string
}

type NLMChild interface {
	utils.Serializable
	InitializeParent(parent NLM, vendorId *BACnetVendorId)
	GetParent() *NLM

	GetTypeName() string
	NLM
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NLM) GetVendorId() *BACnetVendorId {
	return m.VendorId
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewNLM factory function for _NLM
func NewNLM(vendorId *BACnetVendorId, apduLength uint16) *_NLM {
	return &_NLM{VendorId: vendorId, ApduLength: apduLength}
}

// Deprecated: use the interface for direct cast
func CastNLM(structType interface{}) NLM {
	if casted, ok := structType.(NLM); ok {
		return casted
	}
	if casted, ok := structType.(*NLM); ok {
		return *casted
	}
	return nil
}

func (m *_NLM) GetTypeName() string {
	return "NLM"
}

func (m *_NLM) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (messageType)
	lengthInBits += 8

	// Optional Field (vendorId)
	if m.VendorId != nil {
		lengthInBits += 16
	}

	return lengthInBits
}

func (m *_NLM) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func NLMParse(readBuffer utils.ReadBuffer, apduLength uint16) (NLM, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NLM"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NLM")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Discriminator Field (messageType) (Used as input to a switch field)
	messageType, _messageTypeErr := readBuffer.ReadUint8("messageType", 8)
	if _messageTypeErr != nil {
		return nil, errors.Wrap(_messageTypeErr, "Error parsing 'messageType' field of NLM")
	}

	// Optional Field (vendorId) (Can be skipped, if a given expression evaluates to false)
	var vendorId *BACnetVendorId = nil
	if bool((bool((messageType) >= (128)))) && bool((bool((messageType) <= (255)))) {
		if pullErr := readBuffer.PullContext("vendorId"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for vendorId")
		}
		_val, _err := BACnetVendorIdParse(readBuffer)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'vendorId' field of NLM")
		}
		vendorId = &_val
		if closeErr := readBuffer.CloseContext("vendorId"); closeErr != nil {
			return nil, errors.Wrap(closeErr, "Error closing for vendorId")
		}
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type NLMChildSerializeRequirement interface {
		NLM
		InitializeParent(NLM, *BACnetVendorId)
		GetParent() NLM
	}
	var _childTemp interface{}
	var _child NLMChildSerializeRequirement
	var typeSwitchError error
	switch {
	case messageType == 0x00: // NLMWhoIsRouterToNetwork
		_childTemp, typeSwitchError = NLMWhoIsRouterToNetworkParse(readBuffer, apduLength, messageType)
	case messageType == 0x01: // NLMIAmRouterToNetwork
		_childTemp, typeSwitchError = NLMIAmRouterToNetworkParse(readBuffer, apduLength, messageType)
	case messageType == 0x02: // NLMICouldBeRouterToNetwork
		_childTemp, typeSwitchError = NLMICouldBeRouterToNetworkParse(readBuffer, apduLength, messageType)
	case messageType == 0x03: // NLMRejectRouterToNetwork
		_childTemp, typeSwitchError = NLMRejectRouterToNetworkParse(readBuffer, apduLength, messageType)
	case messageType == 0x04: // NLMRouterBusyToNetwork
		_childTemp, typeSwitchError = NLMRouterBusyToNetworkParse(readBuffer, apduLength, messageType)
	case messageType == 0x05: // NLMRouterAvailableToNetwork
		_childTemp, typeSwitchError = NLMRouterAvailableToNetworkParse(readBuffer, apduLength, messageType)
	case messageType == 0x06: // NLMInitalizeRoutingTable
		_childTemp, typeSwitchError = NLMInitalizeRoutingTableParse(readBuffer, apduLength, messageType)
	case messageType == 0x07: // NLMInitalizeRoutingTableAck
		_childTemp, typeSwitchError = NLMInitalizeRoutingTableAckParse(readBuffer, apduLength, messageType)
	case messageType == 0x08: // NLMEstablishConnectionToNetwork
		_childTemp, typeSwitchError = NLMEstablishConnectionToNetworkParse(readBuffer, apduLength, messageType)
	case messageType == 0x09: // NLMDisconnectConnectionToNetwork
		_childTemp, typeSwitchError = NLMDisconnectConnectionToNetworkParse(readBuffer, apduLength, messageType)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [messageType=%v]", messageType)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of NLM")
	}
	_child = _childTemp.(NLMChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("NLM"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NLM")
	}

	// Finish initializing
	_child.InitializeParent(_child, vendorId)
	return _child, nil
}

func (pm *_NLM) SerializeParent(writeBuffer utils.WriteBuffer, child NLM, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("NLM"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for NLM")
	}

	// Discriminator Field (messageType) (Used as input to a switch field)
	messageType := uint8(child.GetMessageType())
	_messageTypeErr := writeBuffer.WriteUint8("messageType", 8, (messageType))

	if _messageTypeErr != nil {
		return errors.Wrap(_messageTypeErr, "Error serializing 'messageType' field")
	}

	// Optional Field (vendorId) (Can be skipped, if the value is null)
	var vendorId *BACnetVendorId = nil
	if m.GetVendorId() != nil {
		if pushErr := writeBuffer.PushContext("vendorId"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for vendorId")
		}
		vendorId = m.GetVendorId()
		_vendorIdErr := writeBuffer.WriteSerializable(vendorId)
		if popErr := writeBuffer.PopContext("vendorId"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for vendorId")
		}
		if _vendorIdErr != nil {
			return errors.Wrap(_vendorIdErr, "Error serializing 'vendorId' field")
		}
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("NLM"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for NLM")
	}
	return nil
}

////
// Arguments Getter

func (m *_NLM) GetApduLength() uint16 {
	return m.ApduLength
}

//
////

func (m *_NLM) isNLM() bool {
	return true
}

func (m *_NLM) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
