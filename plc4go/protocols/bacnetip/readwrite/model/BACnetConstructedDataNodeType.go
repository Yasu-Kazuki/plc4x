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

// BACnetConstructedDataNodeType is the corresponding interface of BACnetConstructedDataNodeType
type BACnetConstructedDataNodeType interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetNodeType returns NodeType (property field)
	GetNodeType() BACnetNodeTypeTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetNodeTypeTagged
}

// BACnetConstructedDataNodeTypeExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataNodeType.
// This is useful for switch cases.
type BACnetConstructedDataNodeTypeExactly interface {
	BACnetConstructedDataNodeType
	isBACnetConstructedDataNodeType() bool
}

// _BACnetConstructedDataNodeType is the data-structure of this message
type _BACnetConstructedDataNodeType struct {
	*_BACnetConstructedData
	NodeType BACnetNodeTypeTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataNodeType) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataNodeType) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_NODE_TYPE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataNodeType) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataNodeType) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataNodeType) GetNodeType() BACnetNodeTypeTagged {
	return m.NodeType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataNodeType) GetActualValue() BACnetNodeTypeTagged {
	return CastBACnetNodeTypeTagged(m.GetNodeType())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataNodeType factory function for _BACnetConstructedDataNodeType
func NewBACnetConstructedDataNodeType(nodeType BACnetNodeTypeTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataNodeType {
	_result := &_BACnetConstructedDataNodeType{
		NodeType:               nodeType,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataNodeType(structType interface{}) BACnetConstructedDataNodeType {
	if casted, ok := structType.(BACnetConstructedDataNodeType); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataNodeType); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataNodeType) GetTypeName() string {
	return "BACnetConstructedDataNodeType"
}

func (m *_BACnetConstructedDataNodeType) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataNodeType) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (nodeType)
	lengthInBits += m.NodeType.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataNodeType) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataNodeTypeParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataNodeType, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataNodeType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataNodeType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (nodeType)
	if pullErr := readBuffer.PullContext("nodeType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for nodeType")
	}
	_nodeType, _nodeTypeErr := BACnetNodeTypeTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _nodeTypeErr != nil {
		return nil, errors.Wrap(_nodeTypeErr, "Error parsing 'nodeType' field of BACnetConstructedDataNodeType")
	}
	nodeType := _nodeType.(BACnetNodeTypeTagged)
	if closeErr := readBuffer.CloseContext("nodeType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for nodeType")
	}

	// Virtual field
	_actualValue := nodeType
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataNodeType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataNodeType")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataNodeType{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		NodeType: nodeType,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataNodeType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataNodeType) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataNodeType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataNodeType")
		}

		// Simple Field (nodeType)
		if pushErr := writeBuffer.PushContext("nodeType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for nodeType")
		}
		_nodeTypeErr := writeBuffer.WriteSerializable(m.GetNodeType())
		if popErr := writeBuffer.PopContext("nodeType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for nodeType")
		}
		if _nodeTypeErr != nil {
			return errors.Wrap(_nodeTypeErr, "Error serializing 'nodeType' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataNodeType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataNodeType")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataNodeType) isBACnetConstructedDataNodeType() bool {
	return true
}

func (m *_BACnetConstructedDataNodeType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
