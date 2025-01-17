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
package org.apache.plc4x.java.eip.readwrite;

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

public class CipConnectionManagerResponse extends CipService implements Message {

  // Accessors for discriminator values.
  public Short getService() {
    return (short) 0x5B;
  }

  public Boolean getResponse() {
    return (boolean) true;
  }

  public Boolean getConnected() {
    return false;
  }

  // Properties.
  protected final long otConnectionId;
  protected final long toConnectionId;
  protected final int connectionSerialNumber;
  protected final int originatorVendorId;
  protected final long originatorSerialNumber;
  protected final long otApi;
  protected final long toApi;

  public CipConnectionManagerResponse(
      long otConnectionId,
      long toConnectionId,
      int connectionSerialNumber,
      int originatorVendorId,
      long originatorSerialNumber,
      long otApi,
      long toApi) {
    super();
    this.otConnectionId = otConnectionId;
    this.toConnectionId = toConnectionId;
    this.connectionSerialNumber = connectionSerialNumber;
    this.originatorVendorId = originatorVendorId;
    this.originatorSerialNumber = originatorSerialNumber;
    this.otApi = otApi;
    this.toApi = toApi;
  }

  public long getOtConnectionId() {
    return otConnectionId;
  }

  public long getToConnectionId() {
    return toConnectionId;
  }

  public int getConnectionSerialNumber() {
    return connectionSerialNumber;
  }

  public int getOriginatorVendorId() {
    return originatorVendorId;
  }

  public long getOriginatorSerialNumber() {
    return originatorSerialNumber;
  }

  public long getOtApi() {
    return otApi;
  }

  public long getToApi() {
    return toApi;
  }

  @Override
  protected void serializeCipServiceChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("CipConnectionManagerResponse");

    // Reserved Field (reserved)
    writeReservedField("reserved", (long) 0x000000, writeUnsignedLong(writeBuffer, 24));

    // Simple Field (otConnectionId)
    writeSimpleField("otConnectionId", otConnectionId, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (toConnectionId)
    writeSimpleField("toConnectionId", toConnectionId, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (connectionSerialNumber)
    writeSimpleField(
        "connectionSerialNumber", connectionSerialNumber, writeUnsignedInt(writeBuffer, 16));

    // Simple Field (originatorVendorId)
    writeSimpleField("originatorVendorId", originatorVendorId, writeUnsignedInt(writeBuffer, 16));

    // Simple Field (originatorSerialNumber)
    writeSimpleField(
        "originatorSerialNumber", originatorSerialNumber, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (otApi)
    writeSimpleField("otApi", otApi, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (toApi)
    writeSimpleField("toApi", toApi, writeUnsignedLong(writeBuffer, 32));

    // Implicit Field (replySize) (Used for parsing, but its value is not stored as it's implicitly
    // given by the objects content)
    short replySize = (short) ((getLengthInBytes()) - (30));
    writeImplicitField("replySize", replySize, writeUnsignedShort(writeBuffer, 8));

    // Reserved Field (reserved)
    writeReservedField("reserved", (short) 0x00, writeUnsignedShort(writeBuffer, 8));

    writeBuffer.popContext("CipConnectionManagerResponse");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    CipConnectionManagerResponse _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Reserved Field (reserved)
    lengthInBits += 24;

    // Simple field (otConnectionId)
    lengthInBits += 32;

    // Simple field (toConnectionId)
    lengthInBits += 32;

    // Simple field (connectionSerialNumber)
    lengthInBits += 16;

    // Simple field (originatorVendorId)
    lengthInBits += 16;

    // Simple field (originatorSerialNumber)
    lengthInBits += 32;

    // Simple field (otApi)
    lengthInBits += 32;

    // Simple field (toApi)
    lengthInBits += 32;

    // Implicit Field (replySize)
    lengthInBits += 8;

    // Reserved Field (reserved)
    lengthInBits += 8;

    return lengthInBits;
  }

  public static CipServiceBuilder staticParseCipServiceBuilder(
      ReadBuffer readBuffer, Boolean connected, Integer serviceLen) throws ParseException {
    readBuffer.pullContext("CipConnectionManagerResponse");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    Long reservedField0 =
        readReservedField("reserved", readUnsignedLong(readBuffer, 24), (long) 0x000000);

    long otConnectionId = readSimpleField("otConnectionId", readUnsignedLong(readBuffer, 32));

    long toConnectionId = readSimpleField("toConnectionId", readUnsignedLong(readBuffer, 32));

    int connectionSerialNumber =
        readSimpleField("connectionSerialNumber", readUnsignedInt(readBuffer, 16));

    int originatorVendorId = readSimpleField("originatorVendorId", readUnsignedInt(readBuffer, 16));

    long originatorSerialNumber =
        readSimpleField("originatorSerialNumber", readUnsignedLong(readBuffer, 32));

    long otApi = readSimpleField("otApi", readUnsignedLong(readBuffer, 32));

    long toApi = readSimpleField("toApi", readUnsignedLong(readBuffer, 32));

    short replySize = readImplicitField("replySize", readUnsignedShort(readBuffer, 8));

    Short reservedField1 =
        readReservedField("reserved", readUnsignedShort(readBuffer, 8), (short) 0x00);

    readBuffer.closeContext("CipConnectionManagerResponse");
    // Create the instance
    return new CipConnectionManagerResponseBuilderImpl(
        otConnectionId,
        toConnectionId,
        connectionSerialNumber,
        originatorVendorId,
        originatorSerialNumber,
        otApi,
        toApi);
  }

  public static class CipConnectionManagerResponseBuilderImpl
      implements CipService.CipServiceBuilder {
    private final long otConnectionId;
    private final long toConnectionId;
    private final int connectionSerialNumber;
    private final int originatorVendorId;
    private final long originatorSerialNumber;
    private final long otApi;
    private final long toApi;

    public CipConnectionManagerResponseBuilderImpl(
        long otConnectionId,
        long toConnectionId,
        int connectionSerialNumber,
        int originatorVendorId,
        long originatorSerialNumber,
        long otApi,
        long toApi) {
      this.otConnectionId = otConnectionId;
      this.toConnectionId = toConnectionId;
      this.connectionSerialNumber = connectionSerialNumber;
      this.originatorVendorId = originatorVendorId;
      this.originatorSerialNumber = originatorSerialNumber;
      this.otApi = otApi;
      this.toApi = toApi;
    }

    public CipConnectionManagerResponse build() {
      CipConnectionManagerResponse cipConnectionManagerResponse =
          new CipConnectionManagerResponse(
              otConnectionId,
              toConnectionId,
              connectionSerialNumber,
              originatorVendorId,
              originatorSerialNumber,
              otApi,
              toApi);
      return cipConnectionManagerResponse;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof CipConnectionManagerResponse)) {
      return false;
    }
    CipConnectionManagerResponse that = (CipConnectionManagerResponse) o;
    return (getOtConnectionId() == that.getOtConnectionId())
        && (getToConnectionId() == that.getToConnectionId())
        && (getConnectionSerialNumber() == that.getConnectionSerialNumber())
        && (getOriginatorVendorId() == that.getOriginatorVendorId())
        && (getOriginatorSerialNumber() == that.getOriginatorSerialNumber())
        && (getOtApi() == that.getOtApi())
        && (getToApi() == that.getToApi())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getOtConnectionId(),
        getToConnectionId(),
        getConnectionSerialNumber(),
        getOriginatorVendorId(),
        getOriginatorSerialNumber(),
        getOtApi(),
        getToApi());
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
