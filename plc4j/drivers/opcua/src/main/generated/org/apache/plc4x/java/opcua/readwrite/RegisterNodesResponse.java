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
package org.apache.plc4x.java.opcua.readwrite;

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

public class RegisterNodesResponse extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public String getIdentifier() {
    return (String) "563";
  }

  // Properties.
  protected final ExtensionObjectDefinition responseHeader;
  protected final int noOfRegisteredNodeIds;
  protected final List<NodeId> registeredNodeIds;

  public RegisterNodesResponse(
      ExtensionObjectDefinition responseHeader,
      int noOfRegisteredNodeIds,
      List<NodeId> registeredNodeIds) {
    super();
    this.responseHeader = responseHeader;
    this.noOfRegisteredNodeIds = noOfRegisteredNodeIds;
    this.registeredNodeIds = registeredNodeIds;
  }

  public ExtensionObjectDefinition getResponseHeader() {
    return responseHeader;
  }

  public int getNoOfRegisteredNodeIds() {
    return noOfRegisteredNodeIds;
  }

  public List<NodeId> getRegisteredNodeIds() {
    return registeredNodeIds;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("RegisterNodesResponse");

    // Simple Field (responseHeader)
    writeSimpleField("responseHeader", responseHeader, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (noOfRegisteredNodeIds)
    writeSimpleField(
        "noOfRegisteredNodeIds", noOfRegisteredNodeIds, writeSignedInt(writeBuffer, 32));

    // Array Field (registeredNodeIds)
    writeComplexTypeArrayField("registeredNodeIds", registeredNodeIds, writeBuffer);

    writeBuffer.popContext("RegisterNodesResponse");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    RegisterNodesResponse _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (responseHeader)
    lengthInBits += responseHeader.getLengthInBits();

    // Simple field (noOfRegisteredNodeIds)
    lengthInBits += 32;

    // Array field
    if (registeredNodeIds != null) {
      int i = 0;
      for (NodeId element : registeredNodeIds) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= registeredNodeIds.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, String identifier) throws ParseException {
    readBuffer.pullContext("RegisterNodesResponse");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    ExtensionObjectDefinition responseHeader =
        readSimpleField(
            "responseHeader",
            new DataReaderComplexDefault<>(
                () -> ExtensionObjectDefinition.staticParse(readBuffer, (String) ("394")),
                readBuffer));

    int noOfRegisteredNodeIds =
        readSimpleField("noOfRegisteredNodeIds", readSignedInt(readBuffer, 32));

    List<NodeId> registeredNodeIds =
        readCountArrayField(
            "registeredNodeIds",
            new DataReaderComplexDefault<>(() -> NodeId.staticParse(readBuffer), readBuffer),
            noOfRegisteredNodeIds);

    readBuffer.closeContext("RegisterNodesResponse");
    // Create the instance
    return new RegisterNodesResponseBuilderImpl(
        responseHeader, noOfRegisteredNodeIds, registeredNodeIds);
  }

  public static class RegisterNodesResponseBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final ExtensionObjectDefinition responseHeader;
    private final int noOfRegisteredNodeIds;
    private final List<NodeId> registeredNodeIds;

    public RegisterNodesResponseBuilderImpl(
        ExtensionObjectDefinition responseHeader,
        int noOfRegisteredNodeIds,
        List<NodeId> registeredNodeIds) {
      this.responseHeader = responseHeader;
      this.noOfRegisteredNodeIds = noOfRegisteredNodeIds;
      this.registeredNodeIds = registeredNodeIds;
    }

    public RegisterNodesResponse build() {
      RegisterNodesResponse registerNodesResponse =
          new RegisterNodesResponse(responseHeader, noOfRegisteredNodeIds, registeredNodeIds);
      return registerNodesResponse;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof RegisterNodesResponse)) {
      return false;
    }
    RegisterNodesResponse that = (RegisterNodesResponse) o;
    return (getResponseHeader() == that.getResponseHeader())
        && (getNoOfRegisteredNodeIds() == that.getNoOfRegisteredNodeIds())
        && (getRegisteredNodeIds() == that.getRegisteredNodeIds())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(), getResponseHeader(), getNoOfRegisteredNodeIds(), getRegisteredNodeIds());
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
