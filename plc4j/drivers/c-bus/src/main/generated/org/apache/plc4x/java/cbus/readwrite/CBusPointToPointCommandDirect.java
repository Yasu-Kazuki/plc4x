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

import static org.apache.plc4x.java.spi.codegen.fields.FieldReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.fields.FieldWriterFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataWriterFactory.*;
import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;

// Code generated by code-generation. DO NOT EDIT.

public class CBusPointToPointCommandDirect extends CBusPointToPointCommand implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final UnitAddress unitAddress;

  // Arguments.
  protected final CBusOptions cBusOptions;
  // Reserved Fields
  private Short reservedField0;

  public CBusPointToPointCommandDirect(
      int bridgeAddressCountPeek,
      CALData calData,
      UnitAddress unitAddress,
      CBusOptions cBusOptions) {
    super(bridgeAddressCountPeek, calData, cBusOptions);
    this.unitAddress = unitAddress;
    this.cBusOptions = cBusOptions;
  }

  public UnitAddress getUnitAddress() {
    return unitAddress;
  }

  @Override
  protected void serializeCBusPointToPointCommandChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("CBusPointToPointCommandDirect");

    // Simple Field (unitAddress)
    writeSimpleField("unitAddress", unitAddress, new DataWriterComplexDefault<>(writeBuffer));

    // Reserved Field (reserved)
    writeReservedField(
        "reserved",
        reservedField0 != null ? reservedField0 : (short) 0x00,
        writeUnsignedShort(writeBuffer, 8));

    writeBuffer.popContext("CBusPointToPointCommandDirect");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    CBusPointToPointCommandDirect _value = this;

    // Simple field (unitAddress)
    lengthInBits += unitAddress.getLengthInBits();

    // Reserved Field (reserved)
    lengthInBits += 8;

    return lengthInBits;
  }

  public static CBusPointToPointCommandDirectBuilder staticParseBuilder(
      ReadBuffer readBuffer, CBusOptions cBusOptions) throws ParseException {
    readBuffer.pullContext("CBusPointToPointCommandDirect");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    UnitAddress unitAddress =
        readSimpleField(
            "unitAddress",
            new DataReaderComplexDefault<>(() -> UnitAddress.staticParse(readBuffer), readBuffer));

    Short reservedField0 =
        readReservedField("reserved", readUnsignedShort(readBuffer, 8), (short) 0x00);

    readBuffer.closeContext("CBusPointToPointCommandDirect");
    // Create the instance
    return new CBusPointToPointCommandDirectBuilder(unitAddress, cBusOptions, reservedField0);
  }

  public static class CBusPointToPointCommandDirectBuilder
      implements CBusPointToPointCommand.CBusPointToPointCommandBuilder {
    private final UnitAddress unitAddress;
    private final CBusOptions cBusOptions;
    private final Short reservedField0;

    public CBusPointToPointCommandDirectBuilder(
        UnitAddress unitAddress, CBusOptions cBusOptions, Short reservedField0) {
      this.unitAddress = unitAddress;
      this.cBusOptions = cBusOptions;
      this.reservedField0 = reservedField0;
    }

    public CBusPointToPointCommandDirect build(
        int bridgeAddressCountPeek, CALData calData, CBusOptions cBusOptions) {
      CBusPointToPointCommandDirect cBusPointToPointCommandDirect =
          new CBusPointToPointCommandDirect(
              bridgeAddressCountPeek, calData, unitAddress, cBusOptions);
      cBusPointToPointCommandDirect.reservedField0 = reservedField0;
      return cBusPointToPointCommandDirect;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof CBusPointToPointCommandDirect)) {
      return false;
    }
    CBusPointToPointCommandDirect that = (CBusPointToPointCommandDirect) o;
    return (getUnitAddress() == that.getUnitAddress()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getUnitAddress());
  }

  @Override
  public String toString() {
    WriteBufferBoxBased writeBufferBoxBased = new WriteBufferBoxBased(true, true);
    try {
      writeBufferBoxBased.writeSerializable(this);
    } catch (SerializationException e) {
      throw new RuntimeException(e);
    }
    return "\n" + writeBufferBoxBased.getBox().toString() + "\n";
  }
}