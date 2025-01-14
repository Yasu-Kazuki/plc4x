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
package org.apache.plc4x.java.bacnetip.readwrite;

import java.util.HashMap;
import java.util.Map;

// Code generated by code-generation. DO NOT EDIT.

public enum BACnetObjectType {
  ACCESS_CREDENTIAL((int) 32),
  ACCESS_DOOR((int) 30),
  ACCESS_POINT((int) 33),
  ACCESS_RIGHTS((int) 34),
  ACCESS_USER((int) 35),
  ACCESS_ZONE((int) 36),
  ACCUMULATOR((int) 23),
  ALERT_ENROLLMENT((int) 52),
  ANALOG_INPUT((int) 0),
  ANALOG_OUTPUT((int) 1),
  ANALOG_VALUE((int) 2),
  AVERAGING((int) 18),
  BINARY_INPUT((int) 3),
  BINARY_LIGHTING_OUTPUT((int) 55),
  BINARY_OUTPUT((int) 4),
  BINARY_VALUE((int) 5),
  BITSTRING_VALUE((int) 39),
  CALENDAR((int) 6),
  CHANNEL((int) 53),
  CHARACTERSTRING_VALUE((int) 40),
  COMMAND((int) 7),
  CREDENTIAL_DATA_INPUT((int) 37),
  DATEPATTERN_VALUE((int) 41),
  DATE_VALUE((int) 42),
  DATETIMEPATTERN_VALUE((int) 43),
  DATETIME_VALUE((int) 44),
  DEVICE((int) 8),
  ELEVATOR_GROUP((int) 57),
  ESCALATOR((int) 58),
  EVENT_ENROLLMENT((int) 9),
  EVENT_LOG((int) 25),
  FILE((int) 10),
  GLOBAL_GROUP((int) 26),
  GROUP((int) 11),
  INTEGER_VALUE((int) 45),
  LARGE_ANALOG_VALUE((int) 46),
  LIFE_SAFETY_POINT((int) 21),
  LIFE_SAFETY_ZONE((int) 22),
  LIFT((int) 59),
  LIGHTING_OUTPUT((int) 54),
  LOAD_CONTROL((int) 28),
  LOOP((int) 12),
  MULTI_STATE_INPUT((int) 13),
  MULTI_STATE_OUTPUT((int) 14),
  MULTI_STATE_VALUE((int) 19),
  NETWORK_PORT((int) 56),
  NETWORK_SECURITY((int) 38),
  NOTIFICATION_CLASS((int) 15),
  NOTIFICATION_FORWARDER((int) 51),
  OCTETSTRING_VALUE((int) 47),
  POSITIVE_INTEGER_VALUE((int) 48),
  PROGRAM((int) 16),
  PULSE_CONVERTER((int) 24),
  SCHEDULE((int) 17),
  STRUCTURED_VIEW((int) 29),
  TIMEPATTERN_VALUE((int) 49),
  TIME_VALUE((int) 50),
  TIMER((int) 31),
  TREND_LOG((int) 20),
  TREND_LOG_MULTIPLE((int) 27),
  VENDOR_PROPRIETARY_VALUE((int) 0x3FF);
  private static final Map<Integer, BACnetObjectType> map;

  static {
    map = new HashMap<>();
    for (BACnetObjectType value : BACnetObjectType.values()) {
      map.put((int) value.getValue(), value);
    }
  }

  private final int value;

  BACnetObjectType(int value) {
    this.value = value;
  }

  public int getValue() {
    return value;
  }

  public static BACnetObjectType enumForValue(int value) {
    return map.get(value);
  }

  public static Boolean isDefined(int value) {
    return map.containsKey(value);
  }
}
