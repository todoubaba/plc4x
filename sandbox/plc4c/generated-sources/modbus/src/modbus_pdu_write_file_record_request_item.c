/*
  Licensed to the Apache Software Foundation (ASF) under one
  or more contributor license agreements.  See the NOTICE file
  distributed with this work for additional information
  regarding copyright ownership.  The ASF licenses this file
  to you under the Apache License, Version 2.0 (the
  "License"); you may not use this file except in compliance
  with the License.  You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing,
  software distributed under the License is distributed on an
  "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
  KIND, either express or implied.  See the License for the
  specific language governing permissions and limitations
  under the License.
*/

#include <stdio.h>
#include <plc4c/spi/evaluation_helper.h>
#include "modbus_pdu_write_file_record_request_item.h"


// Parse function.
plc4c_return_code plc4c_modbus_read_write_modbus_pdu_write_file_record_request_item_parse(plc4c_spi_read_buffer* io, plc4c_modbus_read_write_modbus_pdu_write_file_record_request_item** _message) {
  uint16_t startPos = plc4c_spi_read_get_pos(io);
  uint16_t curPos;
  plc4c_return_code _res = OK;

  // Allocate enough memory to contain this data structure.
  (*_message) = malloc(sizeof(plc4c_modbus_read_write_modbus_pdu_write_file_record_request_item));
  if(*_message == NULL) {
    return NO_MEMORY;
  }

  // Simple Field (referenceType)
  uint8_t referenceType = 0;
  _res = plc4c_spi_read_unsigned_byte(io, 8, (uint8_t*) &referenceType);
  if(_res != OK) {
    return _res;
  }
  (*_message)->reference_type = referenceType;

  // Simple Field (fileNumber)
  uint16_t fileNumber = 0;
  _res = plc4c_spi_read_unsigned_short(io, 16, (uint16_t*) &fileNumber);
  if(_res != OK) {
    return _res;
  }
  (*_message)->file_number = fileNumber;

  // Simple Field (recordNumber)
  uint16_t recordNumber = 0;
  _res = plc4c_spi_read_unsigned_short(io, 16, (uint16_t*) &recordNumber);
  if(_res != OK) {
    return _res;
  }
  (*_message)->record_number = recordNumber;

  // Implicit Field (recordLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
  uint16_t recordLength = 0;
  _res = plc4c_spi_read_unsigned_short(io, 16, (uint16_t*) &recordLength);
  if(_res != OK) {
    return _res;
  }

  // Array field (recordData)
  plc4c_list* recordData = NULL;
  plc4c_utils_list_create(&recordData);
  if(recordData == NULL) {
    return NO_MEMORY;
  }
  {
    // Length array
    uint8_t _recordDataLength = (recordLength) * (2);
    uint8_t recordDataEndPos = plc4c_spi_read_get_pos(io) + _recordDataLength;
    while(plc4c_spi_read_get_pos(io) < recordDataEndPos) {
      int8_t _value = 0;
      _res = plc4c_spi_read_signed_byte(io, 8, (int8_t*) &_value);
      if(_res != OK) {
        return _res;
      }
      plc4c_utils_list_insert_head_value(recordData, &_value);
    }
  }
  (*_message)->record_data = recordData;

  return OK;
}

plc4c_return_code plc4c_modbus_read_write_modbus_pdu_write_file_record_request_item_serialize(plc4c_spi_write_buffer* io, plc4c_modbus_read_write_modbus_pdu_write_file_record_request_item* _message) {
  plc4c_return_code _res = OK;

  // Simple Field (referenceType)
  _res = plc4c_spi_write_unsigned_byte(io, 8, _message->reference_type);
  if(_res != OK) {
    return _res;
  }

  // Simple Field (fileNumber)
  _res = plc4c_spi_write_unsigned_short(io, 16, _message->file_number);
  if(_res != OK) {
    return _res;
  }

  // Simple Field (recordNumber)
  _res = plc4c_spi_write_unsigned_short(io, 16, _message->record_number);
  if(_res != OK) {
    return _res;
  }

  // Implicit Field (recordLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
  _res = plc4c_spi_write_unsigned_short(io, 16, (plc4c_spi_evaluation_helper_count(_message->record_data)) / (2));
  if(_res != OK) {
    return _res;
  }

  // Array field (recordData)
  {
    uint8_t itemCount = plc4c_utils_list_size(_message->record_data);
    for(int curItem = 0; curItem < itemCount; curItem++) {

      int8_t* _value = (int8_t*) plc4c_utils_list_get_value(_message->record_data, curItem);
      plc4c_spi_write_signed_byte(io, 8, *_value);
    }
  }

  return OK;
}

uint16_t plc4c_modbus_read_write_modbus_pdu_write_file_record_request_item_length_in_bytes(plc4c_modbus_read_write_modbus_pdu_write_file_record_request_item* _message) {
  return plc4c_modbus_read_write_modbus_pdu_write_file_record_request_item_length_in_bits(_message) / 8;
}

uint16_t plc4c_modbus_read_write_modbus_pdu_write_file_record_request_item_length_in_bits(plc4c_modbus_read_write_modbus_pdu_write_file_record_request_item* _message) {
  uint16_t lengthInBits = 0;

  // Simple field (referenceType)
  lengthInBits += 8;

  // Simple field (fileNumber)
  lengthInBits += 16;

  // Simple field (recordNumber)
  lengthInBits += 16;

  // Implicit Field (recordLength)
  lengthInBits += 16;

  // Array field
  lengthInBits += 8 * plc4c_utils_list_size(_message->record_data);

  return lengthInBits;
}

