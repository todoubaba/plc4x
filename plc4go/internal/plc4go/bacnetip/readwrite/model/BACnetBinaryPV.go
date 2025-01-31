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
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type BACnetBinaryPV struct {
	RawData    *BACnetContextTagEnumerated
	IsInactive bool
	IsActive   bool
}

// The corresponding interface
type IBACnetBinaryPV interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

func NewBACnetBinaryPV(rawData *BACnetContextTagEnumerated, isInactive bool, isActive bool) *BACnetBinaryPV {
	return &BACnetBinaryPV{RawData: rawData, IsInactive: isInactive, IsActive: isActive}
}

func CastBACnetBinaryPV(structType interface{}) *BACnetBinaryPV {
	castFunc := func(typ interface{}) *BACnetBinaryPV {
		if casted, ok := typ.(BACnetBinaryPV); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetBinaryPV); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetBinaryPV) GetTypeName() string {
	return "BACnetBinaryPV"
}

func (m *BACnetBinaryPV) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetBinaryPV) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Optional Field (rawData)
	if m.RawData != nil {
		lengthInBits += (*m.RawData).LengthInBits()
	}

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetBinaryPV) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetBinaryPVParse(readBuffer utils.ReadBuffer, tagNumber uint8) (*BACnetBinaryPV, error) {
	if pullErr := readBuffer.PullContext("BACnetBinaryPV"); pullErr != nil {
		return nil, pullErr
	}

	// Optional Field (rawData) (Can be skipped, if a given expression evaluates to false)
	var rawData *BACnetContextTagEnumerated = nil
	{
		currentPos := readBuffer.GetPos()
		if pullErr := readBuffer.PullContext("rawData"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := BACnetContextTagParse(readBuffer, tagNumber, BACnetDataType_ENUMERATED)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'rawData' field")
		default:
			rawData = CastBACnetContextTagEnumerated(_val)
			if closeErr := readBuffer.CloseContext("rawData"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	// Virtual field
	_isInactive := bool(bool((rawData) != (nil))) && bool(bool(((*rawData).ActualValue) == (0)))
	isInactive := bool(_isInactive)

	// Virtual field
	_isActive := bool(bool((rawData) != (nil))) && bool(bool(((*rawData).ActualValue) == (1)))
	isActive := bool(_isActive)

	if closeErr := readBuffer.CloseContext("BACnetBinaryPV"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewBACnetBinaryPV(rawData, isInactive, isActive), nil
}

func (m *BACnetBinaryPV) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("BACnetBinaryPV"); pushErr != nil {
		return pushErr
	}

	// Optional Field (rawData) (Can be skipped, if the value is null)
	var rawData *BACnetContextTagEnumerated = nil
	if m.RawData != nil {
		if pushErr := writeBuffer.PushContext("rawData"); pushErr != nil {
			return pushErr
		}
		rawData = m.RawData
		_rawDataErr := rawData.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("rawData"); popErr != nil {
			return popErr
		}
		if _rawDataErr != nil {
			return errors.Wrap(_rawDataErr, "Error serializing 'rawData' field")
		}
	}
	// Virtual field
	if _isInactiveErr := writeBuffer.WriteVirtual("isInactive", m.IsInactive); _isInactiveErr != nil {
		return errors.Wrap(_isInactiveErr, "Error serializing 'isInactive' field")
	}
	// Virtual field
	if _isActiveErr := writeBuffer.WriteVirtual("isActive", m.IsActive); _isActiveErr != nil {
		return errors.Wrap(_isActiveErr, "Error serializing 'isActive' field")
	}

	if popErr := writeBuffer.PopContext("BACnetBinaryPV"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetBinaryPV) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
