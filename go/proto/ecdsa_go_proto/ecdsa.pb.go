// Copyright 2020 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
////////////////////////////////////////////////////////////////////////////////

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: third_party/tink/proto/ecdsa.proto

package ecdsa_go_proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	common_go_proto "github.com/nandawangsa/tink/go/proto/common_go_proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type EcdsaSignatureEncoding int32

const (
	EcdsaSignatureEncoding_UNKNOWN_ENCODING EcdsaSignatureEncoding = 0
	// The signature's format is r || s, where r and s are zero-padded and have the same size in
	// bytes as the order of the curve. For example, for NIST P-256 curve, r and s are zero-padded to
	// 32 bytes.
	EcdsaSignatureEncoding_IEEE_P1363 EcdsaSignatureEncoding = 1
	// The signature is encoded using ASN.1
	// (https://tools.ietf.org/html/rfc5480#appendix-A):
	// ECDSA-Sig-Value :: = SEQUENCE {
	//  r INTEGER,
	//  s INTEGER
	// }
	EcdsaSignatureEncoding_DER EcdsaSignatureEncoding = 2
)

var EcdsaSignatureEncoding_name = map[int32]string{
	0: "UNKNOWN_ENCODING",
	1: "IEEE_P1363",
	2: "DER",
}

var EcdsaSignatureEncoding_value = map[string]int32{
	"UNKNOWN_ENCODING": 0,
	"IEEE_P1363":       1,
	"DER":              2,
}

func (x EcdsaSignatureEncoding) String() string {
	return proto.EnumName(EcdsaSignatureEncoding_name, int32(x))
}

func (EcdsaSignatureEncoding) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8eef580f8138be98, []int{0}
}

// Protos for Ecdsa.
type EcdsaParams struct {
	// Required.
	HashType common_go_proto.HashType `protobuf:"varint,1,opt,name=hash_type,json=hashType,proto3,enum=google.crypto.tink.HashType" json:"hash_type,omitempty"`
	// Required.
	Curve common_go_proto.EllipticCurveType `protobuf:"varint,2,opt,name=curve,proto3,enum=google.crypto.tink.EllipticCurveType" json:"curve,omitempty"`
	// Required.
	Encoding             EcdsaSignatureEncoding `protobuf:"varint,3,opt,name=encoding,proto3,enum=google.crypto.tink.EcdsaSignatureEncoding" json:"encoding,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *EcdsaParams) Reset()         { *m = EcdsaParams{} }
func (m *EcdsaParams) String() string { return proto.CompactTextString(m) }
func (*EcdsaParams) ProtoMessage()    {}
func (*EcdsaParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_8eef580f8138be98, []int{0}
}

func (m *EcdsaParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EcdsaParams.Unmarshal(m, b)
}
func (m *EcdsaParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EcdsaParams.Marshal(b, m, deterministic)
}
func (m *EcdsaParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EcdsaParams.Merge(m, src)
}
func (m *EcdsaParams) XXX_Size() int {
	return xxx_messageInfo_EcdsaParams.Size(m)
}
func (m *EcdsaParams) XXX_DiscardUnknown() {
	xxx_messageInfo_EcdsaParams.DiscardUnknown(m)
}

var xxx_messageInfo_EcdsaParams proto.InternalMessageInfo

func (m *EcdsaParams) GetHashType() common_go_proto.HashType {
	if m != nil {
		return m.HashType
	}
	return common_go_proto.HashType_UNKNOWN_HASH
}

func (m *EcdsaParams) GetCurve() common_go_proto.EllipticCurveType {
	if m != nil {
		return m.Curve
	}
	return common_go_proto.EllipticCurveType_UNKNOWN_CURVE
}

func (m *EcdsaParams) GetEncoding() EcdsaSignatureEncoding {
	if m != nil {
		return m.Encoding
	}
	return EcdsaSignatureEncoding_UNKNOWN_ENCODING
}

// key_type: type.googleapis.com/google.crypto.tink.EcdsaPublicKey
type EcdsaPublicKey struct {
	// Required.
	Version uint32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	// Required.
	Params *EcdsaParams `protobuf:"bytes,2,opt,name=params,proto3" json:"params,omitempty"`
	// Affine coordinates of the public key in bigendian representation. The
	// public key is a point (x, y) on the curve defined by params.curve. For
	// ECDH, it is crucial to verify whether the public key point (x, y) is on the
	// private's key curve. For ECDSA, such verification is a defense in depth.
	// Required.
	X []byte `protobuf:"bytes,3,opt,name=x,proto3" json:"x,omitempty"`
	// Required.
	Y                    []byte   `protobuf:"bytes,4,opt,name=y,proto3" json:"y,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EcdsaPublicKey) Reset()         { *m = EcdsaPublicKey{} }
func (m *EcdsaPublicKey) String() string { return proto.CompactTextString(m) }
func (*EcdsaPublicKey) ProtoMessage()    {}
func (*EcdsaPublicKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_8eef580f8138be98, []int{1}
}

func (m *EcdsaPublicKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EcdsaPublicKey.Unmarshal(m, b)
}
func (m *EcdsaPublicKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EcdsaPublicKey.Marshal(b, m, deterministic)
}
func (m *EcdsaPublicKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EcdsaPublicKey.Merge(m, src)
}
func (m *EcdsaPublicKey) XXX_Size() int {
	return xxx_messageInfo_EcdsaPublicKey.Size(m)
}
func (m *EcdsaPublicKey) XXX_DiscardUnknown() {
	xxx_messageInfo_EcdsaPublicKey.DiscardUnknown(m)
}

var xxx_messageInfo_EcdsaPublicKey proto.InternalMessageInfo

func (m *EcdsaPublicKey) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *EcdsaPublicKey) GetParams() *EcdsaParams {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *EcdsaPublicKey) GetX() []byte {
	if m != nil {
		return m.X
	}
	return nil
}

func (m *EcdsaPublicKey) GetY() []byte {
	if m != nil {
		return m.Y
	}
	return nil
}

// key_type: type.googleapis.com/google.crypto.tink.EcdsaPrivateKey
type EcdsaPrivateKey struct {
	// Required.
	Version uint32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	// Required.
	PublicKey *EcdsaPublicKey `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	// Unsigned big integer in bigendian representation.
	// Required.
	KeyValue             []byte   `protobuf:"bytes,3,opt,name=key_value,json=keyValue,proto3" json:"key_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EcdsaPrivateKey) Reset()         { *m = EcdsaPrivateKey{} }
func (m *EcdsaPrivateKey) String() string { return proto.CompactTextString(m) }
func (*EcdsaPrivateKey) ProtoMessage()    {}
func (*EcdsaPrivateKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_8eef580f8138be98, []int{2}
}

func (m *EcdsaPrivateKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EcdsaPrivateKey.Unmarshal(m, b)
}
func (m *EcdsaPrivateKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EcdsaPrivateKey.Marshal(b, m, deterministic)
}
func (m *EcdsaPrivateKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EcdsaPrivateKey.Merge(m, src)
}
func (m *EcdsaPrivateKey) XXX_Size() int {
	return xxx_messageInfo_EcdsaPrivateKey.Size(m)
}
func (m *EcdsaPrivateKey) XXX_DiscardUnknown() {
	xxx_messageInfo_EcdsaPrivateKey.DiscardUnknown(m)
}

var xxx_messageInfo_EcdsaPrivateKey proto.InternalMessageInfo

func (m *EcdsaPrivateKey) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *EcdsaPrivateKey) GetPublicKey() *EcdsaPublicKey {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *EcdsaPrivateKey) GetKeyValue() []byte {
	if m != nil {
		return m.KeyValue
	}
	return nil
}

type EcdsaKeyFormat struct {
	// Required.
	Params               *EcdsaParams `protobuf:"bytes,2,opt,name=params,proto3" json:"params,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *EcdsaKeyFormat) Reset()         { *m = EcdsaKeyFormat{} }
func (m *EcdsaKeyFormat) String() string { return proto.CompactTextString(m) }
func (*EcdsaKeyFormat) ProtoMessage()    {}
func (*EcdsaKeyFormat) Descriptor() ([]byte, []int) {
	return fileDescriptor_8eef580f8138be98, []int{3}
}

func (m *EcdsaKeyFormat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EcdsaKeyFormat.Unmarshal(m, b)
}
func (m *EcdsaKeyFormat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EcdsaKeyFormat.Marshal(b, m, deterministic)
}
func (m *EcdsaKeyFormat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EcdsaKeyFormat.Merge(m, src)
}
func (m *EcdsaKeyFormat) XXX_Size() int {
	return xxx_messageInfo_EcdsaKeyFormat.Size(m)
}
func (m *EcdsaKeyFormat) XXX_DiscardUnknown() {
	xxx_messageInfo_EcdsaKeyFormat.DiscardUnknown(m)
}

var xxx_messageInfo_EcdsaKeyFormat proto.InternalMessageInfo

func (m *EcdsaKeyFormat) GetParams() *EcdsaParams {
	if m != nil {
		return m.Params
	}
	return nil
}

func init() {
	proto.RegisterEnum("google.crypto.tink.EcdsaSignatureEncoding", EcdsaSignatureEncoding_name, EcdsaSignatureEncoding_value)
	proto.RegisterType((*EcdsaParams)(nil), "google.crypto.tink.EcdsaParams")
	proto.RegisterType((*EcdsaPublicKey)(nil), "google.crypto.tink.EcdsaPublicKey")
	proto.RegisterType((*EcdsaPrivateKey)(nil), "google.crypto.tink.EcdsaPrivateKey")
	proto.RegisterType((*EcdsaKeyFormat)(nil), "google.crypto.tink.EcdsaKeyFormat")
}

func init() {
	proto.RegisterFile("proto/ecdsa.proto", fileDescriptor_8eef580f8138be98)
}

var fileDescriptor_8eef580f8138be98 = []byte{
	// 435 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xd1, 0x6a, 0xdb, 0x30,
	0x14, 0x86, 0xe7, 0x74, 0x4b, 0x93, 0xd3, 0x2e, 0x0b, 0x62, 0x0c, 0xb3, 0x15, 0x36, 0x3c, 0x06,
	0xa3, 0x03, 0x87, 0x35, 0xb0, 0x31, 0x76, 0xb5, 0xb6, 0x6a, 0x17, 0x02, 0x6e, 0xf0, 0xba, 0x0e,
	0x76, 0x23, 0x14, 0x45, 0xd8, 0x22, 0xb6, 0x25, 0x64, 0x39, 0x54, 0x57, 0x7b, 0x80, 0xbd, 0xc5,
	0xde, 0x67, 0xef, 0x34, 0x2c, 0xbb, 0xbd, 0xa9, 0xb3, 0x8b, 0xde, 0xe9, 0x87, 0xf3, 0x49, 0xdf,
	0x7f, 0x10, 0x04, 0x26, 0x15, 0x7a, 0x45, 0x14, 0xd5, 0xc6, 0x4e, 0x8c, 0x28, 0xd6, 0x13, 0xa5,
	0xa5, 0x91, 0x13, 0xce, 0x56, 0x25, 0x0d, 0xdd, 0x19, 0xa1, 0x44, 0xca, 0x24, 0xe3, 0x21, 0xd3,
	0x56, 0x19, 0x19, 0xd6, 0x53, 0xcf, 0x5f, 0x6f, 0xe1, 0x98, 0xcc, 0x73, 0x59, 0x34, 0x60, 0xf0,
	0xd7, 0x83, 0x3d, 0x5c, 0x5f, 0xb4, 0xa0, 0x9a, 0xe6, 0x25, 0xfa, 0x04, 0xc3, 0x94, 0x96, 0x29,
	0x31, 0x56, 0x71, 0xdf, 0x7b, 0xe5, 0xbd, 0x1d, 0x1d, 0x1d, 0x84, 0x77, 0x2f, 0x0f, 0xbf, 0xd2,
	0x32, 0xbd, 0xb4, 0x8a, 0xc7, 0x83, 0xb4, 0x3d, 0xa1, 0xcf, 0xf0, 0x88, 0x55, 0x7a, 0xc3, 0xfd,
	0x9e, 0xc3, 0xde, 0x74, 0x61, 0x38, 0xcb, 0x84, 0x32, 0x82, 0x9d, 0xd4, 0x83, 0x8e, 0x6f, 0x18,
	0x74, 0x06, 0x03, 0x5e, 0x30, 0xb9, 0x12, 0x45, 0xe2, 0xef, 0x38, 0xfe, 0xb0, 0x93, 0xaf, 0x55,
	0xbf, 0x89, 0xa4, 0xa0, 0xa6, 0xd2, 0x1c, 0xb7, 0x44, 0x7c, 0xcb, 0x06, 0xbf, 0x60, 0xd4, 0xd4,
	0xa9, 0x96, 0x99, 0x60, 0x73, 0x6e, 0x91, 0x0f, 0xbb, 0x1b, 0xae, 0x4b, 0x21, 0x0b, 0xd7, 0xe7,
	0x71, 0x7c, 0x13, 0xd1, 0x47, 0xe8, 0x2b, 0xd7, 0xda, 0x19, 0xef, 0x1d, 0xbd, 0xdc, 0xfa, 0x62,
	0xb3, 0x9c, 0xb8, 0x1d, 0x47, 0xfb, 0xe0, 0x5d, 0x3b, 0xcb, 0xfd, 0xd8, 0xbb, 0xae, 0x93, 0xf5,
	0x1f, 0x36, 0xc9, 0x06, 0xbf, 0x3d, 0x78, 0xd2, 0x30, 0x5a, 0x6c, 0xa8, 0xe1, 0xff, 0x57, 0xf8,
	0x02, 0xa0, 0x9c, 0x29, 0x59, 0x73, 0xdb, 0x6a, 0x04, 0xdb, 0x35, 0x6e, 0x4a, 0xc5, 0x43, 0x75,
	0xdb, 0xef, 0x05, 0x0c, 0xd7, 0xdc, 0x92, 0x0d, 0xcd, 0x2a, 0xde, 0x4a, 0x0d, 0xd6, 0xdc, 0x5e,
	0xd5, 0x39, 0x98, 0xb5, 0xeb, 0x98, 0x73, 0x7b, 0x26, 0x75, 0x4e, 0xcd, 0xbd, 0x4b, 0x1f, 0x9e,
	0xc3, 0xb3, 0xee, 0xed, 0xa3, 0xa7, 0x30, 0xfe, 0x1e, 0xcd, 0xa3, 0x8b, 0x1f, 0x11, 0xc1, 0xd1,
	0xc9, 0xc5, 0xe9, 0x2c, 0x3a, 0x1f, 0x3f, 0x40, 0x23, 0x80, 0x19, 0xc6, 0x98, 0x2c, 0xde, 0x4f,
	0x3f, 0x4c, 0xc7, 0x1e, 0xda, 0x85, 0x9d, 0x53, 0x1c, 0x8f, 0x7b, 0xc7, 0x57, 0x70, 0xc0, 0x64,
	0xde, 0xf5, 0xac, 0xfb, 0x92, 0x0b, 0xef, 0xe7, 0xbb, 0x44, 0x98, 0xb4, 0x5a, 0x86, 0x4c, 0xe6,
	0x93, 0x66, 0xec, 0xce, 0xbf, 0x27, 0x89, 0x24, 0x2e, 0xfe, 0xe9, 0xf5, 0x2f, 0x67, 0xd1, 0x7c,
	0x71, 0xbc, 0xec, 0xbb, 0x3c, 0xfd, 0x17, 0x00, 0x00, 0xff, 0xff, 0x40, 0x46, 0x00, 0xdd, 0x30,
	0x03, 0x00, 0x00,
}
