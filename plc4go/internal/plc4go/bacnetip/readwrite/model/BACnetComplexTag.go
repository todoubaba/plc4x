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
type BACnetComplexTag struct {
	TagNumber                uint8
	TagClass                 TagClass
	LengthValueType          uint8
	ExtTagNumber             *uint8
	ExtLength                *uint8
	ExtExtLength             *uint16
	ExtExtExtLength          *uint32
	ActualTagNumber          uint8
	IsPrimitiveAndNotBoolean bool
	ActualLength             uint32
	Child                    IBACnetComplexTagChild
}

// The corresponding interface
type IBACnetComplexTag interface {
	DataType() BACnetDataType
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

type IBACnetComplexTagParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child IBACnetComplexTag, serializeChildFunction func() error) error
	GetTypeName() string
}

type IBACnetComplexTagChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *BACnetComplexTag, tagNumber uint8, tagClass TagClass, lengthValueType uint8, extTagNumber *uint8, extLength *uint8, extExtLength *uint16, extExtExtLength *uint32, actualTagNumber uint8, isPrimitiveAndNotBoolean bool, actualLength uint32)
	GetTypeName() string
	IBACnetComplexTag
}

func NewBACnetComplexTag(tagNumber uint8, tagClass TagClass, lengthValueType uint8, extTagNumber *uint8, extLength *uint8, extExtLength *uint16, extExtExtLength *uint32) *BACnetComplexTag {
	return &BACnetComplexTag{TagNumber: tagNumber, TagClass: tagClass, LengthValueType: lengthValueType, ExtTagNumber: extTagNumber, ExtLength: extLength, ExtExtLength: extExtLength, ExtExtExtLength: extExtExtLength}
}

func CastBACnetComplexTag(structType interface{}) *BACnetComplexTag {
	castFunc := func(typ interface{}) *BACnetComplexTag {
		if casted, ok := typ.(BACnetComplexTag); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetComplexTag); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetComplexTag) GetTypeName() string {
	return "BACnetComplexTag"
}

func (m *BACnetComplexTag) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetComplexTag) LengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.LengthInBits()
}

func (m *BACnetComplexTag) ParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Simple field (lengthValueType)
	lengthInBits += 3

	// Optional Field (extTagNumber)
	if m.ExtTagNumber != nil {
		lengthInBits += 8
	}

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// Optional Field (extLength)
	if m.ExtLength != nil {
		lengthInBits += 8
	}

	// Optional Field (extExtLength)
	if m.ExtExtLength != nil {
		lengthInBits += 16
	}

	// Optional Field (extExtExtLength)
	if m.ExtExtExtLength != nil {
		lengthInBits += 32
	}

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetComplexTag) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetComplexTagParse(readBuffer utils.ReadBuffer, tagNumberArgument uint8, dataType BACnetDataType) (*BACnetComplexTag, error) {
	if pullErr := readBuffer.PullContext("BACnetComplexTag"); pullErr != nil {
		return nil, pullErr
	}

	// Assert Field (tagNumber) (Can be skipped, if a given expression evaluates to false)
	tagNumber, _err := readBuffer.ReadUint8("tagNumber", 4)
	if _err != nil {
		return nil, errors.Wrap(_err, "Error parsing 'tagNumber' field")
	}
	if tagNumber != tagNumberArgument {
		return nil, utils.ParseAssertError
	}

	// Assert Field (tagClass) (Can be skipped, if a given expression evaluates to false)
	tagClass, _err := TagClassParse(readBuffer)
	if _err != nil {
		return nil, errors.Wrap(_err, "Error parsing 'tagClass' field")
	}
	if tagClass != TagClass_CONTEXT_SPECIFIC_TAGS {
		return nil, utils.ParseAssertError
	}

	// Simple Field (lengthValueType)
	lengthValueType, _lengthValueTypeErr := readBuffer.ReadUint8("lengthValueType", 3)
	if _lengthValueTypeErr != nil {
		return nil, errors.Wrap(_lengthValueTypeErr, "Error parsing 'lengthValueType' field")
	}

	// Optional Field (extTagNumber) (Can be skipped, if a given expression evaluates to false)
	var extTagNumber *uint8 = nil
	if bool((tagNumber) == (15)) {
		_val, _err := readBuffer.ReadUint8("extTagNumber", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'extTagNumber' field")
		}
		extTagNumber = &_val
	}

	// Virtual field
	actualTagNumber := utils.InlineIf(bool((tagNumber) < (15)), func() interface{} { return uint8(tagNumber) }, func() interface{} { return uint8((*extTagNumber)) }).(uint8)

	// Virtual field
	isPrimitiveAndNotBoolean := bool(!(bool(bool((lengthValueType) == (6))))) && bool(bool((tagNumber) != (1)))

	// Optional Field (extLength) (Can be skipped, if a given expression evaluates to false)
	var extLength *uint8 = nil
	if bool(isPrimitiveAndNotBoolean) && bool(bool((lengthValueType) == (5))) {
		_val, _err := readBuffer.ReadUint8("extLength", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'extLength' field")
		}
		extLength = &_val
	}

	// Optional Field (extExtLength) (Can be skipped, if a given expression evaluates to false)
	var extExtLength *uint16 = nil
	if bool(bool(isPrimitiveAndNotBoolean) && bool(bool((lengthValueType) == (5)))) && bool(bool((*extLength) == (254))) {
		_val, _err := readBuffer.ReadUint16("extExtLength", 16)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'extExtLength' field")
		}
		extExtLength = &_val
	}

	// Optional Field (extExtExtLength) (Can be skipped, if a given expression evaluates to false)
	var extExtExtLength *uint32 = nil
	if bool(bool(isPrimitiveAndNotBoolean) && bool(bool((lengthValueType) == (5)))) && bool(bool((*extLength) == (255))) {
		_val, _err := readBuffer.ReadUint32("extExtExtLength", 32)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'extExtExtLength' field")
		}
		extExtExtLength = &_val
	}

	// Virtual field
	actualLength := utils.InlineIf(bool(bool((lengthValueType) == (5))) && bool(bool((*extLength) == (255))), func() interface{} { return uint32((*extExtExtLength)) }, func() interface{} {
		return uint32(uint32(utils.InlineIf(bool(bool((lengthValueType) == (5))) && bool(bool((*extLength) == (254))), func() interface{} { return uint32((*extExtLength)) }, func() interface{} {
			return uint32(uint32(utils.InlineIf(bool((lengthValueType) == (5)), func() interface{} { return uint32((*extLength)) }, func() interface{} {
				return uint32(uint32(utils.InlineIf(isPrimitiveAndNotBoolean, func() interface{} { return uint32(lengthValueType) }, func() interface{} { return uint32(uint32(0)) }).(uint32)))
			}).(uint32)))
		}).(uint32)))
	}).(uint32)

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _parent *BACnetComplexTag
	var typeSwitchError error
	switch {
	case dataType == BACnetDataType_NULL: // BACnetComplexTagNull
		_parent, typeSwitchError = BACnetComplexTagNullParse(readBuffer, tagNumberArgument, dataType)
	case dataType == BACnetDataType_BOOLEAN: // BACnetComplexTagBoolean
		_parent, typeSwitchError = BACnetComplexTagBooleanParse(readBuffer, tagNumberArgument, dataType)
	case dataType == BACnetDataType_UNSIGNED_INTEGER: // BACnetComplexTagUnsignedInteger
		_parent, typeSwitchError = BACnetComplexTagUnsignedIntegerParse(readBuffer, tagNumberArgument, dataType, actualLength)
	case dataType == BACnetDataType_SIGNED_INTEGER: // BACnetComplexTagSignedInteger
		_parent, typeSwitchError = BACnetComplexTagSignedIntegerParse(readBuffer, tagNumberArgument, dataType, actualLength)
	case dataType == BACnetDataType_REAL: // BACnetComplexTagReal
		_parent, typeSwitchError = BACnetComplexTagRealParse(readBuffer, tagNumberArgument, dataType, actualLength)
	case dataType == BACnetDataType_DOUBLE: // BACnetComplexTagDouble
		_parent, typeSwitchError = BACnetComplexTagDoubleParse(readBuffer, tagNumberArgument, dataType, actualLength)
	case dataType == BACnetDataType_OCTET_STRING: // BACnetComplexTagOctetString
		_parent, typeSwitchError = BACnetComplexTagOctetStringParse(readBuffer, tagNumberArgument, dataType, actualLength)
	case dataType == BACnetDataType_CHARACTER_STRING: // BACnetComplexTagCharacterString
		_parent, typeSwitchError = BACnetComplexTagCharacterStringParse(readBuffer, tagNumberArgument, dataType, actualLength)
	case dataType == BACnetDataType_BIT_STRING: // BACnetComplexTagBitString
		_parent, typeSwitchError = BACnetComplexTagBitStringParse(readBuffer, tagNumberArgument, dataType, actualLength)
	case dataType == BACnetDataType_ENUMERATED: // BACnetComplexTagEnumerated
		_parent, typeSwitchError = BACnetComplexTagEnumeratedParse(readBuffer, tagNumberArgument, dataType, actualLength)
	case dataType == BACnetDataType_DATE: // BACnetComplexTagDate
		_parent, typeSwitchError = BACnetComplexTagDateParse(readBuffer, tagNumberArgument, dataType)
	case dataType == BACnetDataType_TIME: // BACnetComplexTagTime
		_parent, typeSwitchError = BACnetComplexTagTimeParse(readBuffer, tagNumberArgument, dataType)
	case dataType == BACnetDataType_BACNET_OBJECT_IDENTIFIER: // BACnetComplexTagObjectIdentifier
		_parent, typeSwitchError = BACnetComplexTagObjectIdentifierParse(readBuffer, tagNumberArgument, dataType)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("BACnetComplexTag"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_parent.Child.InitializeParent(_parent, tagNumber, tagClass, lengthValueType, extTagNumber, extLength, extExtLength, extExtExtLength, actualTagNumber, isPrimitiveAndNotBoolean, actualLength)
	return _parent, nil
}

func (m *BACnetComplexTag) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *BACnetComplexTag) SerializeParent(writeBuffer utils.WriteBuffer, child IBACnetComplexTag, serializeChildFunction func() error) error {
	if pushErr := writeBuffer.PushContext("BACnetComplexTag"); pushErr != nil {
		return pushErr
	}

	// Simple Field (lengthValueType)
	lengthValueType := uint8(m.LengthValueType)
	_lengthValueTypeErr := writeBuffer.WriteUint8("lengthValueType", 3, (lengthValueType))
	if _lengthValueTypeErr != nil {
		return errors.Wrap(_lengthValueTypeErr, "Error serializing 'lengthValueType' field")
	}

	// Optional Field (extTagNumber) (Can be skipped, if the value is null)
	var extTagNumber *uint8 = nil
	if m.ExtTagNumber != nil {
		extTagNumber = m.ExtTagNumber
		_extTagNumberErr := writeBuffer.WriteUint8("extTagNumber", 8, *(extTagNumber))
		if _extTagNumberErr != nil {
			return errors.Wrap(_extTagNumberErr, "Error serializing 'extTagNumber' field")
		}
	}

	// Optional Field (extLength) (Can be skipped, if the value is null)
	var extLength *uint8 = nil
	if m.ExtLength != nil {
		extLength = m.ExtLength
		_extLengthErr := writeBuffer.WriteUint8("extLength", 8, *(extLength))
		if _extLengthErr != nil {
			return errors.Wrap(_extLengthErr, "Error serializing 'extLength' field")
		}
	}

	// Optional Field (extExtLength) (Can be skipped, if the value is null)
	var extExtLength *uint16 = nil
	if m.ExtExtLength != nil {
		extExtLength = m.ExtExtLength
		_extExtLengthErr := writeBuffer.WriteUint16("extExtLength", 16, *(extExtLength))
		if _extExtLengthErr != nil {
			return errors.Wrap(_extExtLengthErr, "Error serializing 'extExtLength' field")
		}
	}

	// Optional Field (extExtExtLength) (Can be skipped, if the value is null)
	var extExtExtLength *uint32 = nil
	if m.ExtExtExtLength != nil {
		extExtExtLength = m.ExtExtExtLength
		_extExtExtLengthErr := writeBuffer.WriteUint32("extExtExtLength", 32, *(extExtExtLength))
		if _extExtExtLengthErr != nil {
			return errors.Wrap(_extExtExtLengthErr, "Error serializing 'extExtExtLength' field")
		}
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	_typeSwitchErr := serializeChildFunction()
	if _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("BACnetComplexTag"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetComplexTag) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
