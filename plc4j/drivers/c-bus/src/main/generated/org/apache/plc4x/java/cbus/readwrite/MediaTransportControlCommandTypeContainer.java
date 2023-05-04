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
package org.apache.plc4x.java.cbus.readwrite;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

// Code generated by code-generation. DO NOT EDIT.

public enum MediaTransportControlCommandTypeContainer {
  MediaTransportControlCommandStop((short) 0x01, (short) 1, MediaTransportControlCommandType.STOP),
  MediaTransportControlCommandPlay((short) 0x79, (short) 1, MediaTransportControlCommandType.PLAY),
  MediaTransportControlCommandPauseResume(
      (short) 0x0A, (short) 2, MediaTransportControlCommandType.PAUSE_RESUME),
  MediaTransportControlCommandSelectCategory(
      (short) 0x12, (short) 2, MediaTransportControlCommandType.SELECT_CATEGORY),
  MediaTransportControlCommandSelectSelection(
      (short) 0x1B, (short) 3, MediaTransportControlCommandType.SELECT_SELECTION),
  MediaTransportControlCommandSelectTrack(
      (short) 0x25, (short) 5, MediaTransportControlCommandType.SELECT_TRACK),
  MediaTransportControlCommandShuffleOnOff(
      (short) 0x2A, (short) 2, MediaTransportControlCommandType.SHUFFLE_ON_OFF),
  MediaTransportControlCommandRepeatOnOff(
      (short) 0x32, (short) 2, MediaTransportControlCommandType.REPEAT_ON_OFF),
  MediaTransportControlCommandNextPreviousCategory(
      (short) 0x3A, (short) 2, MediaTransportControlCommandType.NEXT_PREVIOUS_CATEGORY),
  MediaTransportControlCommandNextPreviousSelection(
      (short) 0x42, (short) 2, MediaTransportControlCommandType.NEXT_PREVIOUS_SELECTION),
  MediaTransportControlCommandNextPreviousTrack(
      (short) 0x4A, (short) 2, MediaTransportControlCommandType.NEXT_PREVIOUS_TRACK),
  MediaTransportControlCommandFastForward(
      (short) 0x52, (short) 2, MediaTransportControlCommandType.FAST_FORWARD),
  MediaTransportControlCommandRewind(
      (short) 0x5A, (short) 2, MediaTransportControlCommandType.REWIND),
  MediaTransportControlCommandSourcePowerControl(
      (short) 0x62, (short) 2, MediaTransportControlCommandType.SOURCE_POWER_CONTROL),
  MediaTransportControlCommandTotalTracks(
      (short) 0x6D, (short) 5, MediaTransportControlCommandType.TOTAL_TRACKS),
  MediaTransportControlCommandStatusRequest(
      (short) 0x71, (short) 1, MediaTransportControlCommandType.STATUS_REQUEST),
  MediaTransportControlCommandEnumerateCategoriesSelectionsTracks(
      (short) 0x73,
      (short) 3,
      MediaTransportControlCommandType.ENUMERATE_CATEGORIES_SELECTIONS_TRACKS),
  MediaTransportControlCommandEnumerationSize(
      (short) 0x74, (short) 4, MediaTransportControlCommandType.ENUMERATION_SIZE),
  MediaTransportControlCommandTrackName_0Bytes(
      (short) 0x80, (short) 0, MediaTransportControlCommandType.TRACK_NAME),
  MediaTransportControlCommandTrackName_1Bytes(
      (short) 0x81, (short) 1, MediaTransportControlCommandType.TRACK_NAME),
  MediaTransportControlCommandTrackName_2Bytes(
      (short) 0x82, (short) 2, MediaTransportControlCommandType.TRACK_NAME),
  MediaTransportControlCommandTrackName_3Bytes(
      (short) 0x83, (short) 3, MediaTransportControlCommandType.TRACK_NAME),
  MediaTransportControlCommandTrackName_4Bytes(
      (short) 0x84, (short) 4, MediaTransportControlCommandType.TRACK_NAME),
  MediaTransportControlCommandTrackName_5Bytes(
      (short) 0x85, (short) 5, MediaTransportControlCommandType.TRACK_NAME),
  MediaTransportControlCommandTrackName_6Bytes(
      (short) 0x86, (short) 6, MediaTransportControlCommandType.TRACK_NAME),
  MediaTransportControlCommandTrackName_7Bytes(
      (short) 0x87, (short) 7, MediaTransportControlCommandType.TRACK_NAME),
  MediaTransportControlCommandTrackName_8Bytes(
      (short) 0x88, (short) 8, MediaTransportControlCommandType.TRACK_NAME),
  MediaTransportControlCommandTrackName_9Bytes(
      (short) 0x89, (short) 9, MediaTransportControlCommandType.TRACK_NAME),
  MediaTransportControlCommandTrackName_10Bytes(
      (short) 0x8A, (short) 10, MediaTransportControlCommandType.TRACK_NAME),
  MediaTransportControlCommandTrackName_11Bytes(
      (short) 0x8B, (short) 11, MediaTransportControlCommandType.TRACK_NAME),
  MediaTransportControlCommandTrackName_12Bytes(
      (short) 0x8C, (short) 12, MediaTransportControlCommandType.TRACK_NAME),
  MediaTransportControlCommandTrackName_13Bytes(
      (short) 0x8D, (short) 13, MediaTransportControlCommandType.TRACK_NAME),
  MediaTransportControlCommandTrackName_14Bytes(
      (short) 0x8E, (short) 14, MediaTransportControlCommandType.TRACK_NAME),
  MediaTransportControlCommandTrackName_15Bytes(
      (short) 0x8F, (short) 15, MediaTransportControlCommandType.TRACK_NAME),
  MediaTransportControlCommandTrackName_16Bytes(
      (short) 0x90, (short) 16, MediaTransportControlCommandType.TRACK_NAME),
  MediaTransportControlCommandTrackName_17Bytes(
      (short) 0x91, (short) 17, MediaTransportControlCommandType.TRACK_NAME),
  MediaTransportControlCommandTrackName_18Bytes(
      (short) 0x92, (short) 18, MediaTransportControlCommandType.TRACK_NAME),
  MediaTransportControlCommandTrackName_19Bytes(
      (short) 0x93, (short) 19, MediaTransportControlCommandType.TRACK_NAME),
  MediaTransportControlCommandTrackName_20Bytes(
      (short) 0x94, (short) 20, MediaTransportControlCommandType.TRACK_NAME),
  MediaTransportControlCommandTrackName_21Bytes(
      (short) 0x95, (short) 21, MediaTransportControlCommandType.TRACK_NAME),
  MediaTransportControlCommandTrackName_22Bytes(
      (short) 0x96, (short) 22, MediaTransportControlCommandType.TRACK_NAME),
  MediaTransportControlCommandTrackName_23Bytes(
      (short) 0x97, (short) 23, MediaTransportControlCommandType.TRACK_NAME),
  MediaTransportControlCommandTrackName_24Bytes(
      (short) 0x98, (short) 24, MediaTransportControlCommandType.TRACK_NAME),
  MediaTransportControlCommandTrackName_25Bytes(
      (short) 0x99, (short) 25, MediaTransportControlCommandType.TRACK_NAME),
  MediaTransportControlCommandTrackName_26Bytes(
      (short) 0x9A, (short) 26, MediaTransportControlCommandType.TRACK_NAME),
  MediaTransportControlCommandTrackName_27Bytes(
      (short) 0x9B, (short) 27, MediaTransportControlCommandType.TRACK_NAME),
  MediaTransportControlCommandTrackName_28Bytes(
      (short) 0x9C, (short) 28, MediaTransportControlCommandType.TRACK_NAME),
  MediaTransportControlCommandTrackName_29Bytes(
      (short) 0x9D, (short) 29, MediaTransportControlCommandType.TRACK_NAME),
  MediaTransportControlCommandTrackName_30Bytes(
      (short) 0x9E, (short) 30, MediaTransportControlCommandType.TRACK_NAME),
  MediaTransportControlCommandTrackName_31Bytes(
      (short) 0x9F, (short) 31, MediaTransportControlCommandType.TRACK_NAME),
  MediaTransportControlCommandSelectionName_0Bytes(
      (short) 0xA0, (short) 0, MediaTransportControlCommandType.SELECTION_NAME),
  MediaTransportControlCommandSelectionName_1Bytes(
      (short) 0xA1, (short) 1, MediaTransportControlCommandType.SELECTION_NAME),
  MediaTransportControlCommandSelectionName_2Bytes(
      (short) 0xA2, (short) 2, MediaTransportControlCommandType.SELECTION_NAME),
  MediaTransportControlCommandSelectionName_3Bytes(
      (short) 0xA3, (short) 3, MediaTransportControlCommandType.SELECTION_NAME),
  MediaTransportControlCommandSelectionName_4Bytes(
      (short) 0xA4, (short) 4, MediaTransportControlCommandType.SELECTION_NAME),
  MediaTransportControlCommandSelectionName_5Bytes(
      (short) 0xA5, (short) 5, MediaTransportControlCommandType.SELECTION_NAME),
  MediaTransportControlCommandSelectionName_6Bytes(
      (short) 0xA6, (short) 6, MediaTransportControlCommandType.SELECTION_NAME),
  MediaTransportControlCommandSelectionName_7Bytes(
      (short) 0xA7, (short) 7, MediaTransportControlCommandType.SELECTION_NAME),
  MediaTransportControlCommandSelectionName_8Bytes(
      (short) 0xA8, (short) 8, MediaTransportControlCommandType.SELECTION_NAME),
  MediaTransportControlCommandSelectionName_9Bytes(
      (short) 0xA9, (short) 9, MediaTransportControlCommandType.SELECTION_NAME),
  MediaTransportControlCommandSelectionName_10Bytes(
      (short) 0xAA, (short) 10, MediaTransportControlCommandType.SELECTION_NAME),
  MediaTransportControlCommandSelectionName_11Bytes(
      (short) 0xAB, (short) 11, MediaTransportControlCommandType.SELECTION_NAME),
  MediaTransportControlCommandSelectionName_12Bytes(
      (short) 0xAC, (short) 12, MediaTransportControlCommandType.SELECTION_NAME),
  MediaTransportControlCommandSelectionName_13Bytes(
      (short) 0xAD, (short) 13, MediaTransportControlCommandType.SELECTION_NAME),
  MediaTransportControlCommandSelectionName_14Bytes(
      (short) 0xAE, (short) 14, MediaTransportControlCommandType.SELECTION_NAME),
  MediaTransportControlCommandSelectionName_15Bytes(
      (short) 0xAF, (short) 15, MediaTransportControlCommandType.SELECTION_NAME),
  MediaTransportControlCommandSelectionName_16Bytes(
      (short) 0xB0, (short) 16, MediaTransportControlCommandType.SELECTION_NAME),
  MediaTransportControlCommandSelectionName_17Bytes(
      (short) 0xB1, (short) 17, MediaTransportControlCommandType.SELECTION_NAME),
  MediaTransportControlCommandSelectionName_18Bytes(
      (short) 0xB2, (short) 18, MediaTransportControlCommandType.SELECTION_NAME),
  MediaTransportControlCommandSelectionName_19Bytes(
      (short) 0xB3, (short) 19, MediaTransportControlCommandType.SELECTION_NAME),
  MediaTransportControlCommandSelectionName_20Bytes(
      (short) 0xB4, (short) 20, MediaTransportControlCommandType.SELECTION_NAME),
  MediaTransportControlCommandSelectionName_21Bytes(
      (short) 0xB5, (short) 21, MediaTransportControlCommandType.SELECTION_NAME),
  MediaTransportControlCommandSelectionName_22Bytes(
      (short) 0xB6, (short) 22, MediaTransportControlCommandType.SELECTION_NAME),
  MediaTransportControlCommandSelectionName_23Bytes(
      (short) 0xB7, (short) 23, MediaTransportControlCommandType.SELECTION_NAME),
  MediaTransportControlCommandSelectionName_24Bytes(
      (short) 0xB8, (short) 24, MediaTransportControlCommandType.SELECTION_NAME),
  MediaTransportControlCommandSelectionName_25Bytes(
      (short) 0xB9, (short) 25, MediaTransportControlCommandType.SELECTION_NAME),
  MediaTransportControlCommandSelectionName_26Bytes(
      (short) 0xBA, (short) 26, MediaTransportControlCommandType.SELECTION_NAME),
  MediaTransportControlCommandSelectionName_27Bytes(
      (short) 0xBB, (short) 27, MediaTransportControlCommandType.SELECTION_NAME),
  MediaTransportControlCommandSelectionName_28Bytes(
      (short) 0xBC, (short) 28, MediaTransportControlCommandType.SELECTION_NAME),
  MediaTransportControlCommandSelectionName_29Bytes(
      (short) 0xBD, (short) 29, MediaTransportControlCommandType.SELECTION_NAME),
  MediaTransportControlCommandSelectionName_30Bytes(
      (short) 0xBE, (short) 30, MediaTransportControlCommandType.SELECTION_NAME),
  MediaTransportControlCommandSelectionName_31Bytes(
      (short) 0xBF, (short) 31, MediaTransportControlCommandType.SELECTION_NAME),
  MediaTransportControlCommandCategoryName_0Bytes(
      (short) 0xC0, (short) 0, MediaTransportControlCommandType.CATEGORY_NAME),
  MediaTransportControlCommandCategoryName_1Bytes(
      (short) 0xC1, (short) 1, MediaTransportControlCommandType.CATEGORY_NAME),
  MediaTransportControlCommandCategoryName_2Bytes(
      (short) 0xC2, (short) 2, MediaTransportControlCommandType.CATEGORY_NAME),
  MediaTransportControlCommandCategoryName_3Bytes(
      (short) 0xC3, (short) 3, MediaTransportControlCommandType.CATEGORY_NAME),
  MediaTransportControlCommandCategoryName_4Bytes(
      (short) 0xC4, (short) 4, MediaTransportControlCommandType.CATEGORY_NAME),
  MediaTransportControlCommandCategoryName_5Bytes(
      (short) 0xC5, (short) 5, MediaTransportControlCommandType.CATEGORY_NAME),
  MediaTransportControlCommandCategoryName_6Bytes(
      (short) 0xC6, (short) 6, MediaTransportControlCommandType.CATEGORY_NAME),
  MediaTransportControlCommandCategoryName_7Bytes(
      (short) 0xC7, (short) 7, MediaTransportControlCommandType.CATEGORY_NAME),
  MediaTransportControlCommandCategoryName_8Bytes(
      (short) 0xC8, (short) 8, MediaTransportControlCommandType.CATEGORY_NAME),
  MediaTransportControlCommandCategoryName_9Bytes(
      (short) 0xC9, (short) 9, MediaTransportControlCommandType.CATEGORY_NAME),
  MediaTransportControlCommandCategoryName_10Bytes(
      (short) 0xCA, (short) 10, MediaTransportControlCommandType.CATEGORY_NAME),
  MediaTransportControlCommandCategoryName_11Bytes(
      (short) 0xCB, (short) 11, MediaTransportControlCommandType.CATEGORY_NAME),
  MediaTransportControlCommandCategoryName_12Bytes(
      (short) 0xCC, (short) 12, MediaTransportControlCommandType.CATEGORY_NAME),
  MediaTransportControlCommandCategoryName_13Bytes(
      (short) 0xCD, (short) 13, MediaTransportControlCommandType.CATEGORY_NAME),
  MediaTransportControlCommandCategoryName_14Bytes(
      (short) 0xCE, (short) 14, MediaTransportControlCommandType.CATEGORY_NAME),
  MediaTransportControlCommandCategoryName_15Bytes(
      (short) 0xCF, (short) 15, MediaTransportControlCommandType.CATEGORY_NAME),
  MediaTransportControlCommandCategoryName_16Bytes(
      (short) 0xD0, (short) 16, MediaTransportControlCommandType.CATEGORY_NAME),
  MediaTransportControlCommandCategoryName_17Bytes(
      (short) 0xD1, (short) 17, MediaTransportControlCommandType.CATEGORY_NAME),
  MediaTransportControlCommandCategoryName_18Bytes(
      (short) 0xD2, (short) 18, MediaTransportControlCommandType.CATEGORY_NAME),
  MediaTransportControlCommandCategoryName_19Bytes(
      (short) 0xD3, (short) 19, MediaTransportControlCommandType.CATEGORY_NAME),
  MediaTransportControlCommandCategoryName_20Bytes(
      (short) 0xD4, (short) 20, MediaTransportControlCommandType.CATEGORY_NAME),
  MediaTransportControlCommandCategoryName_21Bytes(
      (short) 0xD5, (short) 21, MediaTransportControlCommandType.CATEGORY_NAME),
  MediaTransportControlCommandCategoryName_22Bytes(
      (short) 0xD6, (short) 22, MediaTransportControlCommandType.CATEGORY_NAME),
  MediaTransportControlCommandCategoryName_23Bytes(
      (short) 0xD7, (short) 23, MediaTransportControlCommandType.CATEGORY_NAME),
  MediaTransportControlCommandCategoryName_24Bytes(
      (short) 0xD8, (short) 24, MediaTransportControlCommandType.CATEGORY_NAME),
  MediaTransportControlCommandCategoryName_25Bytes(
      (short) 0xD9, (short) 25, MediaTransportControlCommandType.CATEGORY_NAME),
  MediaTransportControlCommandCategoryName_26Bytes(
      (short) 0xDA, (short) 26, MediaTransportControlCommandType.CATEGORY_NAME),
  MediaTransportControlCommandCategoryName_27Bytes(
      (short) 0xDB, (short) 27, MediaTransportControlCommandType.CATEGORY_NAME),
  MediaTransportControlCommandCategoryName_28Bytes(
      (short) 0xDC, (short) 28, MediaTransportControlCommandType.CATEGORY_NAME),
  MediaTransportControlCommandCategoryName_29Bytes(
      (short) 0xDD, (short) 29, MediaTransportControlCommandType.CATEGORY_NAME),
  MediaTransportControlCommandCategoryName_30Bytes(
      (short) 0xDE, (short) 30, MediaTransportControlCommandType.CATEGORY_NAME),
  MediaTransportControlCommandCategoryName_31Bytes(
      (short) 0xDF, (short) 31, MediaTransportControlCommandType.CATEGORY_NAME);
  private static final Map<Short, MediaTransportControlCommandTypeContainer> map;

  static {
    map = new HashMap<>();
    for (MediaTransportControlCommandTypeContainer value :
        MediaTransportControlCommandTypeContainer.values()) {
      map.put((short) value.getValue(), value);
    }
  }

  private final short value;
  private short numBytes;
  private MediaTransportControlCommandType commandType;

  MediaTransportControlCommandTypeContainer(
      short value, short numBytes, MediaTransportControlCommandType commandType) {
    this.value = value;
    this.numBytes = numBytes;
    this.commandType = commandType;
  }

  public short getValue() {
    return value;
  }

  public short getNumBytes() {
    return numBytes;
  }

  public static MediaTransportControlCommandTypeContainer firstEnumForFieldNumBytes(
      short fieldValue) {
    for (MediaTransportControlCommandTypeContainer _val :
        MediaTransportControlCommandTypeContainer.values()) {
      if (_val.getNumBytes() == fieldValue) {
        return _val;
      }
    }
    return null;
  }

  public static List<MediaTransportControlCommandTypeContainer> enumsForFieldNumBytes(
      short fieldValue) {
    List<MediaTransportControlCommandTypeContainer> _values = new ArrayList<>();
    for (MediaTransportControlCommandTypeContainer _val :
        MediaTransportControlCommandTypeContainer.values()) {
      if (_val.getNumBytes() == fieldValue) {
        _values.add(_val);
      }
    }
    return _values;
  }

  public MediaTransportControlCommandType getCommandType() {
    return commandType;
  }

  public static MediaTransportControlCommandTypeContainer firstEnumForFieldCommandType(
      MediaTransportControlCommandType fieldValue) {
    for (MediaTransportControlCommandTypeContainer _val :
        MediaTransportControlCommandTypeContainer.values()) {
      if (_val.getCommandType() == fieldValue) {
        return _val;
      }
    }
    return null;
  }

  public static List<MediaTransportControlCommandTypeContainer> enumsForFieldCommandType(
      MediaTransportControlCommandType fieldValue) {
    List<MediaTransportControlCommandTypeContainer> _values = new ArrayList<>();
    for (MediaTransportControlCommandTypeContainer _val :
        MediaTransportControlCommandTypeContainer.values()) {
      if (_val.getCommandType() == fieldValue) {
        _values.add(_val);
      }
    }
    return _values;
  }

  public static MediaTransportControlCommandTypeContainer enumForValue(short value) {
    return map.get(value);
  }

  public static Boolean isDefined(short value) {
    return map.containsKey(value);
  }
}
