/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type ApduDataExtAuthorizeRequest struct {
	*ApduDataExt
	Level uint8
	Data  []byte
}

// The corresponding interface
type IApduDataExtAuthorizeRequest interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ApduDataExtAuthorizeRequest) ExtApciType() uint8 {
	return 0x11
}

func (m *ApduDataExtAuthorizeRequest) InitializeParent(parent *ApduDataExt) {}

func NewApduDataExtAuthorizeRequest(level uint8, data []byte) *ApduDataExt {
	child := &ApduDataExtAuthorizeRequest{
		Level:       level,
		Data:        data,
		ApduDataExt: NewApduDataExt(),
	}
	child.Child = child
	return child.ApduDataExt
}

func CastApduDataExtAuthorizeRequest(structType interface{}) *ApduDataExtAuthorizeRequest {
	castFunc := func(typ interface{}) *ApduDataExtAuthorizeRequest {
		if casted, ok := typ.(ApduDataExtAuthorizeRequest); ok {
			return &casted
		}
		if casted, ok := typ.(*ApduDataExtAuthorizeRequest); ok {
			return casted
		}
		if casted, ok := typ.(ApduDataExt); ok {
			return CastApduDataExtAuthorizeRequest(casted.Child)
		}
		if casted, ok := typ.(*ApduDataExt); ok {
			return CastApduDataExtAuthorizeRequest(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ApduDataExtAuthorizeRequest) GetTypeName() string {
	return "ApduDataExtAuthorizeRequest"
}

func (m *ApduDataExtAuthorizeRequest) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *ApduDataExtAuthorizeRequest) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	// Simple field (level)
	lengthInBits += 8

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *ApduDataExtAuthorizeRequest) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ApduDataExtAuthorizeRequestParse(readBuffer utils.ReadBuffer, length uint8) (*ApduDataExt, error) {
	if pullErr := readBuffer.PullContext("ApduDataExtAuthorizeRequest"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (level)
	_level, _levelErr := readBuffer.ReadUint8("level", 8)
	if _levelErr != nil {
		return nil, errors.Wrap(_levelErr, "Error parsing 'level' field")
	}
	level := _level
	// Byte Array field (data)
	numberOfBytesdata := int(uint16(4))
	data, _readArrayErr := readBuffer.ReadByteArray("data", numberOfBytesdata)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'data' field")
	}

	if closeErr := readBuffer.CloseContext("ApduDataExtAuthorizeRequest"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ApduDataExtAuthorizeRequest{
		Level:       level,
		Data:        data,
		ApduDataExt: &ApduDataExt{},
	}
	_child.ApduDataExt.Child = _child
	return _child.ApduDataExt, nil
}

func (m *ApduDataExtAuthorizeRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtAuthorizeRequest"); pushErr != nil {
			return pushErr
		}

		// Simple Field (level)
		level := uint8(m.Level)
		_levelErr := writeBuffer.WriteUint8("level", 8, (level))
		if _levelErr != nil {
			return errors.Wrap(_levelErr, "Error serializing 'level' field")
		}

		// Array Field (data)
		if m.Data != nil {
			// Byte Array field (data)
			_writeArrayErr := writeBuffer.WriteByteArray("data", m.Data)
			if _writeArrayErr != nil {
				return errors.Wrap(_writeArrayErr, "Error serializing 'data' field")
			}
		}

		if popErr := writeBuffer.PopContext("ApduDataExtAuthorizeRequest"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ApduDataExtAuthorizeRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
