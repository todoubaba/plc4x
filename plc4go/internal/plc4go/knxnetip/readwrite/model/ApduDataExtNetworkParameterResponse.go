//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//
package model

import (
    "encoding/xml"
    "github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
    "io"
)

// The data-structure of this message
type ApduDataExtNetworkParameterResponse struct {
    Parent *ApduDataExt
    IApduDataExtNetworkParameterResponse
}

// The corresponding interface
type IApduDataExtNetworkParameterResponse interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ApduDataExtNetworkParameterResponse) ExtApciType() uint8 {
    return 0x1B
}


func (m *ApduDataExtNetworkParameterResponse) InitializeParent(parent *ApduDataExt) {
}

func NewApduDataExtNetworkParameterResponse() *ApduDataExt {
    child := &ApduDataExtNetworkParameterResponse{
        Parent: NewApduDataExt(),
    }
    child.Parent.Child = child
    return child.Parent
}

func CastApduDataExtNetworkParameterResponse(structType interface{}) *ApduDataExtNetworkParameterResponse {
    castFunc := func(typ interface{}) *ApduDataExtNetworkParameterResponse {
        if casted, ok := typ.(ApduDataExtNetworkParameterResponse); ok {
            return &casted
        }
        if casted, ok := typ.(*ApduDataExtNetworkParameterResponse); ok {
            return casted
        }
        if casted, ok := typ.(ApduDataExt); ok {
            return CastApduDataExtNetworkParameterResponse(casted.Child)
        }
        if casted, ok := typ.(*ApduDataExt); ok {
            return CastApduDataExtNetworkParameterResponse(casted.Child)
        }
        return nil
    }
    return castFunc(structType)
}

func (m *ApduDataExtNetworkParameterResponse) GetTypeName() string {
    return "ApduDataExtNetworkParameterResponse"
}

func (m *ApduDataExtNetworkParameterResponse) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    return lengthInBits
}

func (m *ApduDataExtNetworkParameterResponse) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func ApduDataExtNetworkParameterResponseParse(io *utils.ReadBuffer) (*ApduDataExt, error) {

    // Create a partially initialized instance
    _child := &ApduDataExtNetworkParameterResponse{
        Parent: &ApduDataExt{},
    }
    _child.Parent.Child = _child
    return _child.Parent, nil
}

func (m *ApduDataExtNetworkParameterResponse) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

        return nil
    }
    return m.Parent.SerializeParent(io, m, ser)
}

func (m *ApduDataExtNetworkParameterResponse) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
    var token xml.Token
    var err error
    token = start
    for {
        switch token.(type) {
        case xml.StartElement:
            tok := token.(xml.StartElement)
            switch tok.Name.Local {
            }
        }
        token, err = d.Token()
        if err != nil {
            if err == io.EOF {
                return nil
            }
            return err
        }
    }
}

func (m *ApduDataExtNetworkParameterResponse) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    return nil
}

