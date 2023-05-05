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
	"context"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// S7PayloadUserDataItemCyclicServicesChangeDrivenPush is the corresponding interface of S7PayloadUserDataItemCyclicServicesChangeDrivenPush
type S7PayloadUserDataItemCyclicServicesChangeDrivenPush interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	S7PayloadUserDataItem
	// GetItemsCount returns ItemsCount (property field)
	GetItemsCount() uint16
	// GetItems returns Items (property field)
	GetItems() []AssociatedQueryValueType
}

// S7PayloadUserDataItemCyclicServicesChangeDrivenPushExactly can be used when we want exactly this type and not a type which fulfills S7PayloadUserDataItemCyclicServicesChangeDrivenPush.
// This is useful for switch cases.
type S7PayloadUserDataItemCyclicServicesChangeDrivenPushExactly interface {
	S7PayloadUserDataItemCyclicServicesChangeDrivenPush
	isS7PayloadUserDataItemCyclicServicesChangeDrivenPush() bool
}

// _S7PayloadUserDataItemCyclicServicesChangeDrivenPush is the data-structure of this message
type _S7PayloadUserDataItemCyclicServicesChangeDrivenPush struct {
	*_S7PayloadUserDataItem
	ItemsCount uint16
	Items      []AssociatedQueryValueType
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7PayloadUserDataItemCyclicServicesChangeDrivenPush) GetCpuFunctionGroup() uint8 {
	return 0x02
}

func (m *_S7PayloadUserDataItemCyclicServicesChangeDrivenPush) GetCpuFunctionType() uint8 {
	return 0x00
}

func (m *_S7PayloadUserDataItemCyclicServicesChangeDrivenPush) GetCpuSubfunction() uint8 {
	return 0x05
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7PayloadUserDataItemCyclicServicesChangeDrivenPush) InitializeParent(parent S7PayloadUserDataItem, returnCode DataTransportErrorCode, transportSize DataTransportSize, dataLength uint16) {
	m.ReturnCode = returnCode
	m.TransportSize = transportSize
	m.DataLength = dataLength
}

func (m *_S7PayloadUserDataItemCyclicServicesChangeDrivenPush) GetParent() S7PayloadUserDataItem {
	return m._S7PayloadUserDataItem
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7PayloadUserDataItemCyclicServicesChangeDrivenPush) GetItemsCount() uint16 {
	return m.ItemsCount
}

func (m *_S7PayloadUserDataItemCyclicServicesChangeDrivenPush) GetItems() []AssociatedQueryValueType {
	return m.Items
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewS7PayloadUserDataItemCyclicServicesChangeDrivenPush factory function for _S7PayloadUserDataItemCyclicServicesChangeDrivenPush
func NewS7PayloadUserDataItemCyclicServicesChangeDrivenPush(itemsCount uint16, items []AssociatedQueryValueType, returnCode DataTransportErrorCode, transportSize DataTransportSize, dataLength uint16) *_S7PayloadUserDataItemCyclicServicesChangeDrivenPush {
	_result := &_S7PayloadUserDataItemCyclicServicesChangeDrivenPush{
		ItemsCount:             itemsCount,
		Items:                  items,
		_S7PayloadUserDataItem: NewS7PayloadUserDataItem(returnCode, transportSize, dataLength),
	}
	_result._S7PayloadUserDataItem._S7PayloadUserDataItemChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastS7PayloadUserDataItemCyclicServicesChangeDrivenPush(structType any) S7PayloadUserDataItemCyclicServicesChangeDrivenPush {
	if casted, ok := structType.(S7PayloadUserDataItemCyclicServicesChangeDrivenPush); ok {
		return casted
	}
	if casted, ok := structType.(*S7PayloadUserDataItemCyclicServicesChangeDrivenPush); ok {
		return *casted
	}
	return nil
}

func (m *_S7PayloadUserDataItemCyclicServicesChangeDrivenPush) GetTypeName() string {
	return "S7PayloadUserDataItemCyclicServicesChangeDrivenPush"
}

func (m *_S7PayloadUserDataItemCyclicServicesChangeDrivenPush) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (itemsCount)
	lengthInBits += 16

	// Array field
	if len(m.Items) > 0 {
		for _curItem, element := range m.Items {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Items), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_S7PayloadUserDataItemCyclicServicesChangeDrivenPush) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func S7PayloadUserDataItemCyclicServicesChangeDrivenPushParse(theBytes []byte, cpuFunctionGroup uint8, cpuFunctionType uint8, cpuSubfunction uint8) (S7PayloadUserDataItemCyclicServicesChangeDrivenPush, error) {
	return S7PayloadUserDataItemCyclicServicesChangeDrivenPushParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), cpuFunctionGroup, cpuFunctionType, cpuSubfunction)
}

func S7PayloadUserDataItemCyclicServicesChangeDrivenPushParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, cpuFunctionGroup uint8, cpuFunctionType uint8, cpuSubfunction uint8) (S7PayloadUserDataItemCyclicServicesChangeDrivenPush, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7PayloadUserDataItemCyclicServicesChangeDrivenPush"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7PayloadUserDataItemCyclicServicesChangeDrivenPush")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (itemsCount)
	_itemsCount, _itemsCountErr := readBuffer.ReadUint16("itemsCount", 16)
	if _itemsCountErr != nil {
		return nil, errors.Wrap(_itemsCountErr, "Error parsing 'itemsCount' field of S7PayloadUserDataItemCyclicServicesChangeDrivenPush")
	}
	itemsCount := _itemsCount

	// Array field (items)
	if pullErr := readBuffer.PullContext("items", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for items")
	}
	// Count array
	items := make([]AssociatedQueryValueType, itemsCount)
	// This happens when the size is set conditional to 0
	if len(items) == 0 {
		items = nil
	}
	{
		_numItems := uint16(itemsCount)
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := AssociatedQueryValueTypeParseWithBuffer(arrayCtx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'items' field of S7PayloadUserDataItemCyclicServicesChangeDrivenPush")
			}
			items[_curItem] = _item.(AssociatedQueryValueType)
		}
	}
	if closeErr := readBuffer.CloseContext("items", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for items")
	}

	if closeErr := readBuffer.CloseContext("S7PayloadUserDataItemCyclicServicesChangeDrivenPush"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7PayloadUserDataItemCyclicServicesChangeDrivenPush")
	}

	// Create a partially initialized instance
	_child := &_S7PayloadUserDataItemCyclicServicesChangeDrivenPush{
		_S7PayloadUserDataItem: &_S7PayloadUserDataItem{},
		ItemsCount:             itemsCount,
		Items:                  items,
	}
	_child._S7PayloadUserDataItem._S7PayloadUserDataItemChildRequirements = _child
	return _child, nil
}

func (m *_S7PayloadUserDataItemCyclicServicesChangeDrivenPush) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7PayloadUserDataItemCyclicServicesChangeDrivenPush) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadUserDataItemCyclicServicesChangeDrivenPush"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7PayloadUserDataItemCyclicServicesChangeDrivenPush")
		}

		// Simple Field (itemsCount)
		itemsCount := uint16(m.GetItemsCount())
		_itemsCountErr := writeBuffer.WriteUint16("itemsCount", 16, (itemsCount))
		if _itemsCountErr != nil {
			return errors.Wrap(_itemsCountErr, "Error serializing 'itemsCount' field")
		}

		// Array Field (items)
		if pushErr := writeBuffer.PushContext("items", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for items")
		}
		for _curItem, _element := range m.GetItems() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetItems()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'items' field")
			}
		}
		if popErr := writeBuffer.PopContext("items", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for items")
		}

		if popErr := writeBuffer.PopContext("S7PayloadUserDataItemCyclicServicesChangeDrivenPush"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7PayloadUserDataItemCyclicServicesChangeDrivenPush")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_S7PayloadUserDataItemCyclicServicesChangeDrivenPush) isS7PayloadUserDataItemCyclicServicesChangeDrivenPush() bool {
	return true
}

func (m *_S7PayloadUserDataItemCyclicServicesChangeDrivenPush) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}