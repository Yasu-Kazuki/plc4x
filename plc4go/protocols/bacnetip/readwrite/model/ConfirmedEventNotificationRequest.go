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

// ConfirmedEventNotificationRequest is the corresponding interface of ConfirmedEventNotificationRequest
type ConfirmedEventNotificationRequest interface {
	utils.LengthAware
	utils.Serializable
	// GetProcessIdentifier returns ProcessIdentifier (property field)
	GetProcessIdentifier() BACnetContextTagUnsignedInteger
	// GetInitiatingDeviceIdentifier returns InitiatingDeviceIdentifier (property field)
	GetInitiatingDeviceIdentifier() BACnetContextTagObjectIdentifier
	// GetEventObjectIdentifier returns EventObjectIdentifier (property field)
	GetEventObjectIdentifier() BACnetContextTagObjectIdentifier
	// GetTimestamp returns Timestamp (property field)
	GetTimestamp() BACnetTimeStampEnclosed
	// GetNotificationClass returns NotificationClass (property field)
	GetNotificationClass() BACnetContextTagUnsignedInteger
	// GetPriority returns Priority (property field)
	GetPriority() BACnetContextTagUnsignedInteger
	// GetEventType returns EventType (property field)
	GetEventType() BACnetEventTypeTagged
	// GetMessageText returns MessageText (property field)
	GetMessageText() BACnetContextTagCharacterString
	// GetNotifyType returns NotifyType (property field)
	GetNotifyType() BACnetNotifyTypeTagged
	// GetAckRequired returns AckRequired (property field)
	GetAckRequired() BACnetContextTagBoolean
	// GetFromState returns FromState (property field)
	GetFromState() BACnetEventStateTagged
	// GetToState returns ToState (property field)
	GetToState() BACnetEventStateTagged
	// GetEventValues returns EventValues (property field)
	GetEventValues() BACnetNotificationParameters
}

// ConfirmedEventNotificationRequestExactly can be used when we want exactly this type and not a type which fulfills ConfirmedEventNotificationRequest.
// This is useful for switch cases.
type ConfirmedEventNotificationRequestExactly interface {
	ConfirmedEventNotificationRequest
	isConfirmedEventNotificationRequest() bool
}

// _ConfirmedEventNotificationRequest is the data-structure of this message
type _ConfirmedEventNotificationRequest struct {
	ProcessIdentifier          BACnetContextTagUnsignedInteger
	InitiatingDeviceIdentifier BACnetContextTagObjectIdentifier
	EventObjectIdentifier      BACnetContextTagObjectIdentifier
	Timestamp                  BACnetTimeStampEnclosed
	NotificationClass          BACnetContextTagUnsignedInteger
	Priority                   BACnetContextTagUnsignedInteger
	EventType                  BACnetEventTypeTagged
	MessageText                BACnetContextTagCharacterString
	NotifyType                 BACnetNotifyTypeTagged
	AckRequired                BACnetContextTagBoolean
	FromState                  BACnetEventStateTagged
	ToState                    BACnetEventStateTagged
	EventValues                BACnetNotificationParameters
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ConfirmedEventNotificationRequest) GetProcessIdentifier() BACnetContextTagUnsignedInteger {
	return m.ProcessIdentifier
}

func (m *_ConfirmedEventNotificationRequest) GetInitiatingDeviceIdentifier() BACnetContextTagObjectIdentifier {
	return m.InitiatingDeviceIdentifier
}

func (m *_ConfirmedEventNotificationRequest) GetEventObjectIdentifier() BACnetContextTagObjectIdentifier {
	return m.EventObjectIdentifier
}

func (m *_ConfirmedEventNotificationRequest) GetTimestamp() BACnetTimeStampEnclosed {
	return m.Timestamp
}

func (m *_ConfirmedEventNotificationRequest) GetNotificationClass() BACnetContextTagUnsignedInteger {
	return m.NotificationClass
}

func (m *_ConfirmedEventNotificationRequest) GetPriority() BACnetContextTagUnsignedInteger {
	return m.Priority
}

func (m *_ConfirmedEventNotificationRequest) GetEventType() BACnetEventTypeTagged {
	return m.EventType
}

func (m *_ConfirmedEventNotificationRequest) GetMessageText() BACnetContextTagCharacterString {
	return m.MessageText
}

func (m *_ConfirmedEventNotificationRequest) GetNotifyType() BACnetNotifyTypeTagged {
	return m.NotifyType
}

func (m *_ConfirmedEventNotificationRequest) GetAckRequired() BACnetContextTagBoolean {
	return m.AckRequired
}

func (m *_ConfirmedEventNotificationRequest) GetFromState() BACnetEventStateTagged {
	return m.FromState
}

func (m *_ConfirmedEventNotificationRequest) GetToState() BACnetEventStateTagged {
	return m.ToState
}

func (m *_ConfirmedEventNotificationRequest) GetEventValues() BACnetNotificationParameters {
	return m.EventValues
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewConfirmedEventNotificationRequest factory function for _ConfirmedEventNotificationRequest
func NewConfirmedEventNotificationRequest(processIdentifier BACnetContextTagUnsignedInteger, initiatingDeviceIdentifier BACnetContextTagObjectIdentifier, eventObjectIdentifier BACnetContextTagObjectIdentifier, timestamp BACnetTimeStampEnclosed, notificationClass BACnetContextTagUnsignedInteger, priority BACnetContextTagUnsignedInteger, eventType BACnetEventTypeTagged, messageText BACnetContextTagCharacterString, notifyType BACnetNotifyTypeTagged, ackRequired BACnetContextTagBoolean, fromState BACnetEventStateTagged, toState BACnetEventStateTagged, eventValues BACnetNotificationParameters) *_ConfirmedEventNotificationRequest {
	return &_ConfirmedEventNotificationRequest{ProcessIdentifier: processIdentifier, InitiatingDeviceIdentifier: initiatingDeviceIdentifier, EventObjectIdentifier: eventObjectIdentifier, Timestamp: timestamp, NotificationClass: notificationClass, Priority: priority, EventType: eventType, MessageText: messageText, NotifyType: notifyType, AckRequired: ackRequired, FromState: fromState, ToState: toState, EventValues: eventValues}
}

// Deprecated: use the interface for direct cast
func CastConfirmedEventNotificationRequest(structType interface{}) ConfirmedEventNotificationRequest {
	if casted, ok := structType.(ConfirmedEventNotificationRequest); ok {
		return casted
	}
	if casted, ok := structType.(*ConfirmedEventNotificationRequest); ok {
		return *casted
	}
	return nil
}

func (m *_ConfirmedEventNotificationRequest) GetTypeName() string {
	return "ConfirmedEventNotificationRequest"
}

func (m *_ConfirmedEventNotificationRequest) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ConfirmedEventNotificationRequest) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (processIdentifier)
	lengthInBits += m.ProcessIdentifier.GetLengthInBits()

	// Simple field (initiatingDeviceIdentifier)
	lengthInBits += m.InitiatingDeviceIdentifier.GetLengthInBits()

	// Simple field (eventObjectIdentifier)
	lengthInBits += m.EventObjectIdentifier.GetLengthInBits()

	// Simple field (timestamp)
	lengthInBits += m.Timestamp.GetLengthInBits()

	// Simple field (notificationClass)
	lengthInBits += m.NotificationClass.GetLengthInBits()

	// Simple field (priority)
	lengthInBits += m.Priority.GetLengthInBits()

	// Simple field (eventType)
	lengthInBits += m.EventType.GetLengthInBits()

	// Optional Field (messageText)
	if m.MessageText != nil {
		lengthInBits += m.MessageText.GetLengthInBits()
	}

	// Simple field (notifyType)
	lengthInBits += m.NotifyType.GetLengthInBits()

	// Optional Field (ackRequired)
	if m.AckRequired != nil {
		lengthInBits += m.AckRequired.GetLengthInBits()
	}

	// Optional Field (fromState)
	if m.FromState != nil {
		lengthInBits += m.FromState.GetLengthInBits()
	}

	// Simple field (toState)
	lengthInBits += m.ToState.GetLengthInBits()

	// Optional Field (eventValues)
	if m.EventValues != nil {
		lengthInBits += m.EventValues.GetLengthInBits()
	}

	return lengthInBits
}

func (m *_ConfirmedEventNotificationRequest) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ConfirmedEventNotificationRequestParse(readBuffer utils.ReadBuffer) (ConfirmedEventNotificationRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ConfirmedEventNotificationRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ConfirmedEventNotificationRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (processIdentifier)
	if pullErr := readBuffer.PullContext("processIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for processIdentifier")
	}
	_processIdentifier, _processIdentifierErr := BACnetContextTagParse(readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _processIdentifierErr != nil {
		return nil, errors.Wrap(_processIdentifierErr, "Error parsing 'processIdentifier' field of ConfirmedEventNotificationRequest")
	}
	processIdentifier := _processIdentifier.(BACnetContextTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("processIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for processIdentifier")
	}

	// Simple Field (initiatingDeviceIdentifier)
	if pullErr := readBuffer.PullContext("initiatingDeviceIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for initiatingDeviceIdentifier")
	}
	_initiatingDeviceIdentifier, _initiatingDeviceIdentifierErr := BACnetContextTagParse(readBuffer, uint8(uint8(1)), BACnetDataType(BACnetDataType_BACNET_OBJECT_IDENTIFIER))
	if _initiatingDeviceIdentifierErr != nil {
		return nil, errors.Wrap(_initiatingDeviceIdentifierErr, "Error parsing 'initiatingDeviceIdentifier' field of ConfirmedEventNotificationRequest")
	}
	initiatingDeviceIdentifier := _initiatingDeviceIdentifier.(BACnetContextTagObjectIdentifier)
	if closeErr := readBuffer.CloseContext("initiatingDeviceIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for initiatingDeviceIdentifier")
	}

	// Simple Field (eventObjectIdentifier)
	if pullErr := readBuffer.PullContext("eventObjectIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for eventObjectIdentifier")
	}
	_eventObjectIdentifier, _eventObjectIdentifierErr := BACnetContextTagParse(readBuffer, uint8(uint8(2)), BACnetDataType(BACnetDataType_BACNET_OBJECT_IDENTIFIER))
	if _eventObjectIdentifierErr != nil {
		return nil, errors.Wrap(_eventObjectIdentifierErr, "Error parsing 'eventObjectIdentifier' field of ConfirmedEventNotificationRequest")
	}
	eventObjectIdentifier := _eventObjectIdentifier.(BACnetContextTagObjectIdentifier)
	if closeErr := readBuffer.CloseContext("eventObjectIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for eventObjectIdentifier")
	}

	// Simple Field (timestamp)
	if pullErr := readBuffer.PullContext("timestamp"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for timestamp")
	}
	_timestamp, _timestampErr := BACnetTimeStampEnclosedParse(readBuffer, uint8(uint8(3)))
	if _timestampErr != nil {
		return nil, errors.Wrap(_timestampErr, "Error parsing 'timestamp' field of ConfirmedEventNotificationRequest")
	}
	timestamp := _timestamp.(BACnetTimeStampEnclosed)
	if closeErr := readBuffer.CloseContext("timestamp"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for timestamp")
	}

	// Simple Field (notificationClass)
	if pullErr := readBuffer.PullContext("notificationClass"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for notificationClass")
	}
	_notificationClass, _notificationClassErr := BACnetContextTagParse(readBuffer, uint8(uint8(4)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _notificationClassErr != nil {
		return nil, errors.Wrap(_notificationClassErr, "Error parsing 'notificationClass' field of ConfirmedEventNotificationRequest")
	}
	notificationClass := _notificationClass.(BACnetContextTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("notificationClass"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for notificationClass")
	}

	// Simple Field (priority)
	if pullErr := readBuffer.PullContext("priority"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for priority")
	}
	_priority, _priorityErr := BACnetContextTagParse(readBuffer, uint8(uint8(5)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _priorityErr != nil {
		return nil, errors.Wrap(_priorityErr, "Error parsing 'priority' field of ConfirmedEventNotificationRequest")
	}
	priority := _priority.(BACnetContextTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("priority"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for priority")
	}

	// Simple Field (eventType)
	if pullErr := readBuffer.PullContext("eventType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for eventType")
	}
	_eventType, _eventTypeErr := BACnetEventTypeTaggedParse(readBuffer, uint8(uint8(6)), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _eventTypeErr != nil {
		return nil, errors.Wrap(_eventTypeErr, "Error parsing 'eventType' field of ConfirmedEventNotificationRequest")
	}
	eventType := _eventType.(BACnetEventTypeTagged)
	if closeErr := readBuffer.CloseContext("eventType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for eventType")
	}

	// Optional Field (messageText) (Can be skipped, if a given expression evaluates to false)
	var messageText BACnetContextTagCharacterString = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("messageText"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for messageText")
		}
		_val, _err := BACnetContextTagParse(readBuffer, uint8(7), BACnetDataType_CHARACTER_STRING)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			Plc4xModelLog.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'messageText' field of ConfirmedEventNotificationRequest")
		default:
			messageText = _val.(BACnetContextTagCharacterString)
			if closeErr := readBuffer.CloseContext("messageText"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for messageText")
			}
		}
	}

	// Simple Field (notifyType)
	if pullErr := readBuffer.PullContext("notifyType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for notifyType")
	}
	_notifyType, _notifyTypeErr := BACnetNotifyTypeTaggedParse(readBuffer, uint8(uint8(8)), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _notifyTypeErr != nil {
		return nil, errors.Wrap(_notifyTypeErr, "Error parsing 'notifyType' field of ConfirmedEventNotificationRequest")
	}
	notifyType := _notifyType.(BACnetNotifyTypeTagged)
	if closeErr := readBuffer.CloseContext("notifyType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for notifyType")
	}

	// Optional Field (ackRequired) (Can be skipped, if a given expression evaluates to false)
	var ackRequired BACnetContextTagBoolean = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("ackRequired"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for ackRequired")
		}
		_val, _err := BACnetContextTagParse(readBuffer, uint8(9), BACnetDataType_BOOLEAN)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			Plc4xModelLog.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'ackRequired' field of ConfirmedEventNotificationRequest")
		default:
			ackRequired = _val.(BACnetContextTagBoolean)
			if closeErr := readBuffer.CloseContext("ackRequired"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for ackRequired")
			}
		}
	}

	// Optional Field (fromState) (Can be skipped, if a given expression evaluates to false)
	var fromState BACnetEventStateTagged = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("fromState"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for fromState")
		}
		_val, _err := BACnetEventStateTaggedParse(readBuffer, uint8(10), TagClass_CONTEXT_SPECIFIC_TAGS)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			Plc4xModelLog.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'fromState' field of ConfirmedEventNotificationRequest")
		default:
			fromState = _val.(BACnetEventStateTagged)
			if closeErr := readBuffer.CloseContext("fromState"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for fromState")
			}
		}
	}

	// Simple Field (toState)
	if pullErr := readBuffer.PullContext("toState"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for toState")
	}
	_toState, _toStateErr := BACnetEventStateTaggedParse(readBuffer, uint8(uint8(11)), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _toStateErr != nil {
		return nil, errors.Wrap(_toStateErr, "Error parsing 'toState' field of ConfirmedEventNotificationRequest")
	}
	toState := _toState.(BACnetEventStateTagged)
	if closeErr := readBuffer.CloseContext("toState"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for toState")
	}

	// Optional Field (eventValues) (Can be skipped, if a given expression evaluates to false)
	var eventValues BACnetNotificationParameters = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("eventValues"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for eventValues")
		}
		_val, _err := BACnetNotificationParametersParse(readBuffer, uint8(12), eventObjectIdentifier.GetObjectType())
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			Plc4xModelLog.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'eventValues' field of ConfirmedEventNotificationRequest")
		default:
			eventValues = _val.(BACnetNotificationParameters)
			if closeErr := readBuffer.CloseContext("eventValues"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for eventValues")
			}
		}
	}

	if closeErr := readBuffer.CloseContext("ConfirmedEventNotificationRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ConfirmedEventNotificationRequest")
	}

	// Create the instance
	return &_ConfirmedEventNotificationRequest{
		ProcessIdentifier:          processIdentifier,
		InitiatingDeviceIdentifier: initiatingDeviceIdentifier,
		EventObjectIdentifier:      eventObjectIdentifier,
		Timestamp:                  timestamp,
		NotificationClass:          notificationClass,
		Priority:                   priority,
		EventType:                  eventType,
		MessageText:                messageText,
		NotifyType:                 notifyType,
		AckRequired:                ackRequired,
		FromState:                  fromState,
		ToState:                    toState,
		EventValues:                eventValues,
	}, nil
}

func (m *_ConfirmedEventNotificationRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ConfirmedEventNotificationRequest) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("ConfirmedEventNotificationRequest"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ConfirmedEventNotificationRequest")
	}

	// Simple Field (processIdentifier)
	if pushErr := writeBuffer.PushContext("processIdentifier"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for processIdentifier")
	}
	_processIdentifierErr := writeBuffer.WriteSerializable(m.GetProcessIdentifier())
	if popErr := writeBuffer.PopContext("processIdentifier"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for processIdentifier")
	}
	if _processIdentifierErr != nil {
		return errors.Wrap(_processIdentifierErr, "Error serializing 'processIdentifier' field")
	}

	// Simple Field (initiatingDeviceIdentifier)
	if pushErr := writeBuffer.PushContext("initiatingDeviceIdentifier"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for initiatingDeviceIdentifier")
	}
	_initiatingDeviceIdentifierErr := writeBuffer.WriteSerializable(m.GetInitiatingDeviceIdentifier())
	if popErr := writeBuffer.PopContext("initiatingDeviceIdentifier"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for initiatingDeviceIdentifier")
	}
	if _initiatingDeviceIdentifierErr != nil {
		return errors.Wrap(_initiatingDeviceIdentifierErr, "Error serializing 'initiatingDeviceIdentifier' field")
	}

	// Simple Field (eventObjectIdentifier)
	if pushErr := writeBuffer.PushContext("eventObjectIdentifier"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for eventObjectIdentifier")
	}
	_eventObjectIdentifierErr := writeBuffer.WriteSerializable(m.GetEventObjectIdentifier())
	if popErr := writeBuffer.PopContext("eventObjectIdentifier"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for eventObjectIdentifier")
	}
	if _eventObjectIdentifierErr != nil {
		return errors.Wrap(_eventObjectIdentifierErr, "Error serializing 'eventObjectIdentifier' field")
	}

	// Simple Field (timestamp)
	if pushErr := writeBuffer.PushContext("timestamp"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for timestamp")
	}
	_timestampErr := writeBuffer.WriteSerializable(m.GetTimestamp())
	if popErr := writeBuffer.PopContext("timestamp"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for timestamp")
	}
	if _timestampErr != nil {
		return errors.Wrap(_timestampErr, "Error serializing 'timestamp' field")
	}

	// Simple Field (notificationClass)
	if pushErr := writeBuffer.PushContext("notificationClass"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for notificationClass")
	}
	_notificationClassErr := writeBuffer.WriteSerializable(m.GetNotificationClass())
	if popErr := writeBuffer.PopContext("notificationClass"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for notificationClass")
	}
	if _notificationClassErr != nil {
		return errors.Wrap(_notificationClassErr, "Error serializing 'notificationClass' field")
	}

	// Simple Field (priority)
	if pushErr := writeBuffer.PushContext("priority"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for priority")
	}
	_priorityErr := writeBuffer.WriteSerializable(m.GetPriority())
	if popErr := writeBuffer.PopContext("priority"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for priority")
	}
	if _priorityErr != nil {
		return errors.Wrap(_priorityErr, "Error serializing 'priority' field")
	}

	// Simple Field (eventType)
	if pushErr := writeBuffer.PushContext("eventType"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for eventType")
	}
	_eventTypeErr := writeBuffer.WriteSerializable(m.GetEventType())
	if popErr := writeBuffer.PopContext("eventType"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for eventType")
	}
	if _eventTypeErr != nil {
		return errors.Wrap(_eventTypeErr, "Error serializing 'eventType' field")
	}

	// Optional Field (messageText) (Can be skipped, if the value is null)
	var messageText BACnetContextTagCharacterString = nil
	if m.GetMessageText() != nil {
		if pushErr := writeBuffer.PushContext("messageText"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for messageText")
		}
		messageText = m.GetMessageText()
		_messageTextErr := writeBuffer.WriteSerializable(messageText)
		if popErr := writeBuffer.PopContext("messageText"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for messageText")
		}
		if _messageTextErr != nil {
			return errors.Wrap(_messageTextErr, "Error serializing 'messageText' field")
		}
	}

	// Simple Field (notifyType)
	if pushErr := writeBuffer.PushContext("notifyType"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for notifyType")
	}
	_notifyTypeErr := writeBuffer.WriteSerializable(m.GetNotifyType())
	if popErr := writeBuffer.PopContext("notifyType"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for notifyType")
	}
	if _notifyTypeErr != nil {
		return errors.Wrap(_notifyTypeErr, "Error serializing 'notifyType' field")
	}

	// Optional Field (ackRequired) (Can be skipped, if the value is null)
	var ackRequired BACnetContextTagBoolean = nil
	if m.GetAckRequired() != nil {
		if pushErr := writeBuffer.PushContext("ackRequired"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ackRequired")
		}
		ackRequired = m.GetAckRequired()
		_ackRequiredErr := writeBuffer.WriteSerializable(ackRequired)
		if popErr := writeBuffer.PopContext("ackRequired"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ackRequired")
		}
		if _ackRequiredErr != nil {
			return errors.Wrap(_ackRequiredErr, "Error serializing 'ackRequired' field")
		}
	}

	// Optional Field (fromState) (Can be skipped, if the value is null)
	var fromState BACnetEventStateTagged = nil
	if m.GetFromState() != nil {
		if pushErr := writeBuffer.PushContext("fromState"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for fromState")
		}
		fromState = m.GetFromState()
		_fromStateErr := writeBuffer.WriteSerializable(fromState)
		if popErr := writeBuffer.PopContext("fromState"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for fromState")
		}
		if _fromStateErr != nil {
			return errors.Wrap(_fromStateErr, "Error serializing 'fromState' field")
		}
	}

	// Simple Field (toState)
	if pushErr := writeBuffer.PushContext("toState"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for toState")
	}
	_toStateErr := writeBuffer.WriteSerializable(m.GetToState())
	if popErr := writeBuffer.PopContext("toState"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for toState")
	}
	if _toStateErr != nil {
		return errors.Wrap(_toStateErr, "Error serializing 'toState' field")
	}

	// Optional Field (eventValues) (Can be skipped, if the value is null)
	var eventValues BACnetNotificationParameters = nil
	if m.GetEventValues() != nil {
		if pushErr := writeBuffer.PushContext("eventValues"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for eventValues")
		}
		eventValues = m.GetEventValues()
		_eventValuesErr := writeBuffer.WriteSerializable(eventValues)
		if popErr := writeBuffer.PopContext("eventValues"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for eventValues")
		}
		if _eventValuesErr != nil {
			return errors.Wrap(_eventValuesErr, "Error serializing 'eventValues' field")
		}
	}

	if popErr := writeBuffer.PopContext("ConfirmedEventNotificationRequest"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ConfirmedEventNotificationRequest")
	}
	return nil
}

func (m *_ConfirmedEventNotificationRequest) isConfirmedEventNotificationRequest() bool {
	return true
}

func (m *_ConfirmedEventNotificationRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
