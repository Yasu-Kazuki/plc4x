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

// CBusPointToPointCommand is the corresponding interface of CBusPointToPointCommand
type CBusPointToPointCommand interface {
	utils.LengthAware
	utils.Serializable
	// GetBridgeAddressCountPeek returns BridgeAddressCountPeek (property field)
	GetBridgeAddressCountPeek() uint16
	// GetCalData returns CalData (property field)
	GetCalData() CALData
	// GetIsDirect returns IsDirect (virtual field)
	GetIsDirect() bool
}

// CBusPointToPointCommandExactly can be used when we want exactly this type and not a type which fulfills CBusPointToPointCommand.
// This is useful for switch cases.
type CBusPointToPointCommandExactly interface {
	CBusPointToPointCommand
	isCBusPointToPointCommand() bool
}

// _CBusPointToPointCommand is the data-structure of this message
type _CBusPointToPointCommand struct {
	_CBusPointToPointCommandChildRequirements
	BridgeAddressCountPeek uint16
	CalData                CALData

	// Arguments.
	CBusOptions CBusOptions
}

type _CBusPointToPointCommandChildRequirements interface {
	utils.Serializable
	GetLengthInBits() uint16
	GetLengthInBitsConditional(lastItem bool) uint16
}

type CBusPointToPointCommandParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child CBusPointToPointCommand, serializeChildFunction func() error) error
	GetTypeName() string
}

type CBusPointToPointCommandChild interface {
	utils.Serializable
	InitializeParent(parent CBusPointToPointCommand, bridgeAddressCountPeek uint16, calData CALData)
	GetParent() *CBusPointToPointCommand

	GetTypeName() string
	CBusPointToPointCommand
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CBusPointToPointCommand) GetBridgeAddressCountPeek() uint16 {
	return m.BridgeAddressCountPeek
}

func (m *_CBusPointToPointCommand) GetCalData() CALData {
	return m.CalData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_CBusPointToPointCommand) GetIsDirect() bool {
	return bool(bool((m.GetBridgeAddressCountPeek() & 0x00FF) == (0x0000)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCBusPointToPointCommand factory function for _CBusPointToPointCommand
func NewCBusPointToPointCommand(bridgeAddressCountPeek uint16, calData CALData, cBusOptions CBusOptions) *_CBusPointToPointCommand {
	return &_CBusPointToPointCommand{BridgeAddressCountPeek: bridgeAddressCountPeek, CalData: calData, CBusOptions: cBusOptions}
}

// Deprecated: use the interface for direct cast
func CastCBusPointToPointCommand(structType interface{}) CBusPointToPointCommand {
	if casted, ok := structType.(CBusPointToPointCommand); ok {
		return casted
	}
	if casted, ok := structType.(*CBusPointToPointCommand); ok {
		return *casted
	}
	return nil
}

func (m *_CBusPointToPointCommand) GetTypeName() string {
	return "CBusPointToPointCommand"
}

func (m *_CBusPointToPointCommand) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	// Simple field (calData)
	lengthInBits += m.CalData.GetLengthInBits()

	return lengthInBits
}

func (m *_CBusPointToPointCommand) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CBusPointToPointCommandParse(readBuffer utils.ReadBuffer, cBusOptions CBusOptions) (CBusPointToPointCommand, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CBusPointToPointCommand"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CBusPointToPointCommand")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Peek Field (bridgeAddressCountPeek)
	currentPos = positionAware.GetPos()
	bridgeAddressCountPeek, _err := readBuffer.ReadUint16("bridgeAddressCountPeek", 16)
	if _err != nil {
		return nil, errors.Wrap(_err, "Error parsing 'bridgeAddressCountPeek' field of CBusPointToPointCommand")
	}

	readBuffer.Reset(currentPos)

	// Virtual field
	_isDirect := bool((bridgeAddressCountPeek & 0x00FF) == (0x0000))
	isDirect := bool(_isDirect)
	_ = isDirect

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type CBusPointToPointCommandChildSerializeRequirement interface {
		CBusPointToPointCommand
		InitializeParent(CBusPointToPointCommand, uint16, CALData)
		GetParent() CBusPointToPointCommand
	}
	var _childTemp interface{}
	var _child CBusPointToPointCommandChildSerializeRequirement
	var typeSwitchError error
	switch {
	case isDirect == bool(true): // CBusPointToPointCommandDirect
		_childTemp, typeSwitchError = CBusPointToPointCommandDirectParse(readBuffer, cBusOptions)
	case isDirect == bool(false): // CBusPointToPointCommandIndirect
		_childTemp, typeSwitchError = CBusPointToPointCommandIndirectParse(readBuffer, cBusOptions)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [isDirect=%v]", isDirect)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of CBusPointToPointCommand")
	}
	_child = _childTemp.(CBusPointToPointCommandChildSerializeRequirement)

	// Simple Field (calData)
	if pullErr := readBuffer.PullContext("calData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for calData")
	}
	_calData, _calDataErr := CALDataParse(readBuffer, nil)
	if _calDataErr != nil {
		return nil, errors.Wrap(_calDataErr, "Error parsing 'calData' field of CBusPointToPointCommand")
	}
	calData := _calData.(CALData)
	if closeErr := readBuffer.CloseContext("calData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for calData")
	}

	if closeErr := readBuffer.CloseContext("CBusPointToPointCommand"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CBusPointToPointCommand")
	}

	// Finish initializing
	_child.InitializeParent(_child, bridgeAddressCountPeek, calData)
	return _child, nil
}

func (pm *_CBusPointToPointCommand) SerializeParent(writeBuffer utils.WriteBuffer, child CBusPointToPointCommand, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("CBusPointToPointCommand"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for CBusPointToPointCommand")
	}
	// Virtual field
	if _isDirectErr := writeBuffer.WriteVirtual("isDirect", m.GetIsDirect()); _isDirectErr != nil {
		return errors.Wrap(_isDirectErr, "Error serializing 'isDirect' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	// Simple Field (calData)
	if pushErr := writeBuffer.PushContext("calData"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for calData")
	}
	_calDataErr := writeBuffer.WriteSerializable(m.GetCalData())
	if popErr := writeBuffer.PopContext("calData"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for calData")
	}
	if _calDataErr != nil {
		return errors.Wrap(_calDataErr, "Error serializing 'calData' field")
	}

	if popErr := writeBuffer.PopContext("CBusPointToPointCommand"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for CBusPointToPointCommand")
	}
	return nil
}

////
// Arguments Getter

func (m *_CBusPointToPointCommand) GetCBusOptions() CBusOptions {
	return m.CBusOptions
}

//
////

func (m *_CBusPointToPointCommand) isCBusPointToPointCommand() bool {
	return true
}

func (m *_CBusPointToPointCommand) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
