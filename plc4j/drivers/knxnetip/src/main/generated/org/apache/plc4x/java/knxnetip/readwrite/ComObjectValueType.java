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
package org.apache.plc4x.java.knxnetip.readwrite;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

// Code generated by code-generation. DO NOT EDIT.

public enum ComObjectValueType {
  BIT1((short) 0x00, (short) 1),
  BIT2((short) 0x01, (short) 1),
  BIT3((short) 0x02, (short) 1),
  BIT4((short) 0x03, (short) 1),
  BIT5((short) 0x04, (short) 1),
  BIT6((short) 0x05, (short) 1),
  BIT7((short) 0x06, (short) 1),
  BYTE1((short) 0x07, (short) 1),
  BYTE2((short) 0x08, (short) 2),
  BYTE3((short) 0x09, (short) 3),
  BYTE4((short) 0x0A, (short) 4),
  BYTE6((short) 0x0B, (short) 6),
  BYTE8((short) 0x0C, (short) 8),
  BYTE10((short) 0x0D, (short) 10),
  BYTE14((short) 0x0E, (short) 14);
  private static final Map<Short, ComObjectValueType> map;

  static {
    map = new HashMap<>();
    for (ComObjectValueType value : ComObjectValueType.values()) {
      map.put((short) value.getValue(), value);
    }
  }

  private final short value;
  private short sizeInBytes;

  ComObjectValueType(short value, short sizeInBytes) {
    this.value = value;
    this.sizeInBytes = sizeInBytes;
  }

  public short getValue() {
    return value;
  }

  public short getSizeInBytes() {
    return sizeInBytes;
  }

  public static ComObjectValueType firstEnumForFieldSizeInBytes(short fieldValue) {
    for (ComObjectValueType _val : ComObjectValueType.values()) {
      if (_val.getSizeInBytes() == fieldValue) {
        return _val;
      }
    }
    return null;
  }

  public static List<ComObjectValueType> enumsForFieldSizeInBytes(short fieldValue) {
    List<ComObjectValueType> _values = new ArrayList<>();
    for (ComObjectValueType _val : ComObjectValueType.values()) {
      if (_val.getSizeInBytes() == fieldValue) {
        _values.add(_val);
      }
    }
    return _values;
  }

  public static ComObjectValueType enumForValue(short value) {
    return map.get(value);
  }

  public static Boolean isDefined(short value) {
    return map.containsKey(value);
  }
}
