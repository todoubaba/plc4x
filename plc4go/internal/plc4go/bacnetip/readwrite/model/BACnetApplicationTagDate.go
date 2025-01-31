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
type BACnetApplicationTagDate struct {
	*BACnetApplicationTag
	YearMinus1900          int8
	Month                  int8
	DayOfMonth             int8
	DayOfWeek              int8
	Wildcard               int8
	YearIsWildcard         bool
	Year                   int16
	MonthIsWildcard        bool
	OddMonthWildcard       bool
	EvenMonthWildcard      bool
	DayOfMonthIsWildcard   bool
	LastDayOfMonthWildcard bool
	OddDayOfMonthWildcard  bool
	EvenDayOfMonthWildcard bool
	DayOfWeekIsWildcard    bool
}

// The corresponding interface
type IBACnetApplicationTagDate interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetApplicationTagDate) ActualTagNumber() uint8 {
	return 0xA
}

func (m *BACnetApplicationTagDate) InitializeParent(parent *BACnetApplicationTag, header *BACnetTagHeader, actualTagNumber uint8, actualLength uint32) {
	m.BACnetApplicationTag.Header = header
	m.BACnetApplicationTag.ActualTagNumber = actualTagNumber
	m.BACnetApplicationTag.ActualLength = actualLength
}

func NewBACnetApplicationTagDate(yearMinus1900 int8, month int8, dayOfMonth int8, dayOfWeek int8, wildcard int8, yearIsWildcard bool, year int16, monthIsWildcard bool, oddMonthWildcard bool, evenMonthWildcard bool, dayOfMonthIsWildcard bool, lastDayOfMonthWildcard bool, oddDayOfMonthWildcard bool, evenDayOfMonthWildcard bool, dayOfWeekIsWildcard bool, header *BACnetTagHeader, actualTagNumber uint8, actualLength uint32) *BACnetApplicationTag {
	child := &BACnetApplicationTagDate{
		YearMinus1900:          yearMinus1900,
		Month:                  month,
		DayOfMonth:             dayOfMonth,
		DayOfWeek:              dayOfWeek,
		Wildcard:               wildcard,
		YearIsWildcard:         yearIsWildcard,
		Year:                   year,
		MonthIsWildcard:        monthIsWildcard,
		OddMonthWildcard:       oddMonthWildcard,
		EvenMonthWildcard:      evenMonthWildcard,
		DayOfMonthIsWildcard:   dayOfMonthIsWildcard,
		LastDayOfMonthWildcard: lastDayOfMonthWildcard,
		OddDayOfMonthWildcard:  oddDayOfMonthWildcard,
		EvenDayOfMonthWildcard: evenDayOfMonthWildcard,
		DayOfWeekIsWildcard:    dayOfWeekIsWildcard,
		BACnetApplicationTag:   NewBACnetApplicationTag(header, actualTagNumber, actualLength),
	}
	child.Child = child
	return child.BACnetApplicationTag
}

func CastBACnetApplicationTagDate(structType interface{}) *BACnetApplicationTagDate {
	castFunc := func(typ interface{}) *BACnetApplicationTagDate {
		if casted, ok := typ.(BACnetApplicationTagDate); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetApplicationTagDate); ok {
			return casted
		}
		if casted, ok := typ.(BACnetApplicationTag); ok {
			return CastBACnetApplicationTagDate(casted.Child)
		}
		if casted, ok := typ.(*BACnetApplicationTag); ok {
			return CastBACnetApplicationTagDate(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetApplicationTagDate) GetTypeName() string {
	return "BACnetApplicationTagDate"
}

func (m *BACnetApplicationTagDate) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetApplicationTagDate) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	// A virtual field doesn't have any in- or output.

	// Simple field (yearMinus1900)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// Simple field (month)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// Simple field (dayOfMonth)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// Simple field (dayOfWeek)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetApplicationTagDate) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetApplicationTagDateParse(readBuffer utils.ReadBuffer) (*BACnetApplicationTag, error) {
	if pullErr := readBuffer.PullContext("BACnetApplicationTagDate"); pullErr != nil {
		return nil, pullErr
	}

	// Virtual field
	_wildcard := 0xFF
	wildcard := int8(_wildcard)

	// Simple Field (yearMinus1900)
	_yearMinus1900, _yearMinus1900Err := readBuffer.ReadInt8("yearMinus1900", 8)
	if _yearMinus1900Err != nil {
		return nil, errors.Wrap(_yearMinus1900Err, "Error parsing 'yearMinus1900' field")
	}
	yearMinus1900 := _yearMinus1900

	// Virtual field
	_yearIsWildcard := bool((yearMinus1900) == (wildcard))
	yearIsWildcard := bool(_yearIsWildcard)

	// Virtual field
	_year := int16(yearMinus1900) + int16(int16(1900))
	year := int16(_year)

	// Simple Field (month)
	_month, _monthErr := readBuffer.ReadInt8("month", 8)
	if _monthErr != nil {
		return nil, errors.Wrap(_monthErr, "Error parsing 'month' field")
	}
	month := _month

	// Virtual field
	_monthIsWildcard := bool((month) == (wildcard))
	monthIsWildcard := bool(_monthIsWildcard)

	// Virtual field
	_oddMonthWildcard := bool((month) == (13))
	oddMonthWildcard := bool(_oddMonthWildcard)

	// Virtual field
	_evenMonthWildcard := bool((month) == (14))
	evenMonthWildcard := bool(_evenMonthWildcard)

	// Simple Field (dayOfMonth)
	_dayOfMonth, _dayOfMonthErr := readBuffer.ReadInt8("dayOfMonth", 8)
	if _dayOfMonthErr != nil {
		return nil, errors.Wrap(_dayOfMonthErr, "Error parsing 'dayOfMonth' field")
	}
	dayOfMonth := _dayOfMonth

	// Virtual field
	_dayOfMonthIsWildcard := bool((dayOfMonth) == (wildcard))
	dayOfMonthIsWildcard := bool(_dayOfMonthIsWildcard)

	// Virtual field
	_lastDayOfMonthWildcard := bool((dayOfMonth) == (32))
	lastDayOfMonthWildcard := bool(_lastDayOfMonthWildcard)

	// Virtual field
	_oddDayOfMonthWildcard := bool((dayOfMonth) == (33))
	oddDayOfMonthWildcard := bool(_oddDayOfMonthWildcard)

	// Virtual field
	_evenDayOfMonthWildcard := bool((dayOfMonth) == (34))
	evenDayOfMonthWildcard := bool(_evenDayOfMonthWildcard)

	// Simple Field (dayOfWeek)
	_dayOfWeek, _dayOfWeekErr := readBuffer.ReadInt8("dayOfWeek", 8)
	if _dayOfWeekErr != nil {
		return nil, errors.Wrap(_dayOfWeekErr, "Error parsing 'dayOfWeek' field")
	}
	dayOfWeek := _dayOfWeek

	// Virtual field
	_dayOfWeekIsWildcard := bool((dayOfWeek) == (wildcard))
	dayOfWeekIsWildcard := bool(_dayOfWeekIsWildcard)

	if closeErr := readBuffer.CloseContext("BACnetApplicationTagDate"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetApplicationTagDate{
		YearMinus1900:          yearMinus1900,
		Month:                  month,
		DayOfMonth:             dayOfMonth,
		DayOfWeek:              dayOfWeek,
		Wildcard:               wildcard,
		YearIsWildcard:         yearIsWildcard,
		Year:                   year,
		MonthIsWildcard:        monthIsWildcard,
		OddMonthWildcard:       oddMonthWildcard,
		EvenMonthWildcard:      evenMonthWildcard,
		DayOfMonthIsWildcard:   dayOfMonthIsWildcard,
		LastDayOfMonthWildcard: lastDayOfMonthWildcard,
		OddDayOfMonthWildcard:  oddDayOfMonthWildcard,
		EvenDayOfMonthWildcard: evenDayOfMonthWildcard,
		DayOfWeekIsWildcard:    dayOfWeekIsWildcard,
		BACnetApplicationTag:   &BACnetApplicationTag{},
	}
	_child.BACnetApplicationTag.Child = _child
	return _child.BACnetApplicationTag, nil
}

func (m *BACnetApplicationTagDate) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetApplicationTagDate"); pushErr != nil {
			return pushErr
		}
		// Virtual field
		if _wildcardErr := writeBuffer.WriteVirtual("wildcard", m.Wildcard); _wildcardErr != nil {
			return errors.Wrap(_wildcardErr, "Error serializing 'wildcard' field")
		}

		// Simple Field (yearMinus1900)
		yearMinus1900 := int8(m.YearMinus1900)
		_yearMinus1900Err := writeBuffer.WriteInt8("yearMinus1900", 8, (yearMinus1900))
		if _yearMinus1900Err != nil {
			return errors.Wrap(_yearMinus1900Err, "Error serializing 'yearMinus1900' field")
		}
		// Virtual field
		if _yearIsWildcardErr := writeBuffer.WriteVirtual("yearIsWildcard", m.YearIsWildcard); _yearIsWildcardErr != nil {
			return errors.Wrap(_yearIsWildcardErr, "Error serializing 'yearIsWildcard' field")
		}
		// Virtual field
		if _yearErr := writeBuffer.WriteVirtual("year", m.Year); _yearErr != nil {
			return errors.Wrap(_yearErr, "Error serializing 'year' field")
		}

		// Simple Field (month)
		month := int8(m.Month)
		_monthErr := writeBuffer.WriteInt8("month", 8, (month))
		if _monthErr != nil {
			return errors.Wrap(_monthErr, "Error serializing 'month' field")
		}
		// Virtual field
		if _monthIsWildcardErr := writeBuffer.WriteVirtual("monthIsWildcard", m.MonthIsWildcard); _monthIsWildcardErr != nil {
			return errors.Wrap(_monthIsWildcardErr, "Error serializing 'monthIsWildcard' field")
		}
		// Virtual field
		if _oddMonthWildcardErr := writeBuffer.WriteVirtual("oddMonthWildcard", m.OddMonthWildcard); _oddMonthWildcardErr != nil {
			return errors.Wrap(_oddMonthWildcardErr, "Error serializing 'oddMonthWildcard' field")
		}
		// Virtual field
		if _evenMonthWildcardErr := writeBuffer.WriteVirtual("evenMonthWildcard", m.EvenMonthWildcard); _evenMonthWildcardErr != nil {
			return errors.Wrap(_evenMonthWildcardErr, "Error serializing 'evenMonthWildcard' field")
		}

		// Simple Field (dayOfMonth)
		dayOfMonth := int8(m.DayOfMonth)
		_dayOfMonthErr := writeBuffer.WriteInt8("dayOfMonth", 8, (dayOfMonth))
		if _dayOfMonthErr != nil {
			return errors.Wrap(_dayOfMonthErr, "Error serializing 'dayOfMonth' field")
		}
		// Virtual field
		if _dayOfMonthIsWildcardErr := writeBuffer.WriteVirtual("dayOfMonthIsWildcard", m.DayOfMonthIsWildcard); _dayOfMonthIsWildcardErr != nil {
			return errors.Wrap(_dayOfMonthIsWildcardErr, "Error serializing 'dayOfMonthIsWildcard' field")
		}
		// Virtual field
		if _lastDayOfMonthWildcardErr := writeBuffer.WriteVirtual("lastDayOfMonthWildcard", m.LastDayOfMonthWildcard); _lastDayOfMonthWildcardErr != nil {
			return errors.Wrap(_lastDayOfMonthWildcardErr, "Error serializing 'lastDayOfMonthWildcard' field")
		}
		// Virtual field
		if _oddDayOfMonthWildcardErr := writeBuffer.WriteVirtual("oddDayOfMonthWildcard", m.OddDayOfMonthWildcard); _oddDayOfMonthWildcardErr != nil {
			return errors.Wrap(_oddDayOfMonthWildcardErr, "Error serializing 'oddDayOfMonthWildcard' field")
		}
		// Virtual field
		if _evenDayOfMonthWildcardErr := writeBuffer.WriteVirtual("evenDayOfMonthWildcard", m.EvenDayOfMonthWildcard); _evenDayOfMonthWildcardErr != nil {
			return errors.Wrap(_evenDayOfMonthWildcardErr, "Error serializing 'evenDayOfMonthWildcard' field")
		}

		// Simple Field (dayOfWeek)
		dayOfWeek := int8(m.DayOfWeek)
		_dayOfWeekErr := writeBuffer.WriteInt8("dayOfWeek", 8, (dayOfWeek))
		if _dayOfWeekErr != nil {
			return errors.Wrap(_dayOfWeekErr, "Error serializing 'dayOfWeek' field")
		}
		// Virtual field
		if _dayOfWeekIsWildcardErr := writeBuffer.WriteVirtual("dayOfWeekIsWildcard", m.DayOfWeekIsWildcard); _dayOfWeekIsWildcardErr != nil {
			return errors.Wrap(_dayOfWeekIsWildcardErr, "Error serializing 'dayOfWeekIsWildcard' field")
		}

		if popErr := writeBuffer.PopContext("BACnetApplicationTagDate"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetApplicationTagDate) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
