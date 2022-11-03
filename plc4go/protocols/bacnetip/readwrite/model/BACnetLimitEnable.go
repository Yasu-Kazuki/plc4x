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

// BACnetLimitEnable is an enum
type BACnetLimitEnable uint8

type IBACnetLimitEnable interface {
	utils.Serializable
}

const (
	BACnetLimitEnable_LOW_LIMIT_ENABLE  BACnetLimitEnable = 0
	BACnetLimitEnable_HIGH_LIMIT_ENABLE BACnetLimitEnable = 1
)

var BACnetLimitEnableValues []BACnetLimitEnable

func init() {
	_ = errors.New
	BACnetLimitEnableValues = []BACnetLimitEnable{
		BACnetLimitEnable_LOW_LIMIT_ENABLE,
		BACnetLimitEnable_HIGH_LIMIT_ENABLE,
	}
}

func BACnetLimitEnableByValue(value uint8) (enum BACnetLimitEnable, ok bool) {
	switch value {
	case 0:
		return BACnetLimitEnable_LOW_LIMIT_ENABLE, true
	case 1:
		return BACnetLimitEnable_HIGH_LIMIT_ENABLE, true
	}
	return 0, false
}

func BACnetLimitEnableByName(value string) (enum BACnetLimitEnable, ok bool) {
	switch value {
	case "LOW_LIMIT_ENABLE":
		return BACnetLimitEnable_LOW_LIMIT_ENABLE, true
	case "HIGH_LIMIT_ENABLE":
		return BACnetLimitEnable_HIGH_LIMIT_ENABLE, true
	}
	return 0, false
}

func BACnetLimitEnableKnows(value uint8) bool {
	for _, typeValue := range BACnetLimitEnableValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetLimitEnable(structType interface{}) BACnetLimitEnable {
	castFunc := func(typ interface{}) BACnetLimitEnable {
		if sBACnetLimitEnable, ok := typ.(BACnetLimitEnable); ok {
			return sBACnetLimitEnable
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetLimitEnable) GetLengthInBits() uint16 {
	return 8
}

func (m BACnetLimitEnable) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetLimitEnableParse(theBytes []byte) (BACnetLimitEnable, error) {
	return BACnetLimitEnableParseWithBuffer(utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian))) // TODO: get endianness from mspec
}

func BACnetLimitEnableParseWithBuffer(readBuffer utils.ReadBuffer) (BACnetLimitEnable, error) {
	val, err := readBuffer.ReadUint8("BACnetLimitEnable", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetLimitEnable")
	}
	if enum, ok := BACnetLimitEnableByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return BACnetLimitEnable(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetLimitEnable) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := e.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetLimitEnable) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("BACnetLimitEnable", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetLimitEnable) PLC4XEnumName() string {
	switch e {
	case BACnetLimitEnable_LOW_LIMIT_ENABLE:
		return "LOW_LIMIT_ENABLE"
	case BACnetLimitEnable_HIGH_LIMIT_ENABLE:
		return "HIGH_LIMIT_ENABLE"
	}
	return ""
}

func (e BACnetLimitEnable) String() string {
	return e.PLC4XEnumName()
}