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
// source: third_party/tink/proto/rsa_ssa_pkcs1.proto

package rsa_ssa_pkcs1_go_proto

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

type RsaSsaPkcs1Params struct {
	// Hash function used in computing hash of the signing message
	// (see https://tools.ietf.org/html/rfc8017#section-9.2).
	// Required.
	HashType             common_go_proto.HashType `protobuf:"varint,1,opt,name=hash_type,json=hashType,proto3,enum=google.crypto.tink.HashType" json:"hash_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *RsaSsaPkcs1Params) Reset()         { *m = RsaSsaPkcs1Params{} }
func (m *RsaSsaPkcs1Params) String() string { return proto.CompactTextString(m) }
func (*RsaSsaPkcs1Params) ProtoMessage()    {}
func (*RsaSsaPkcs1Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_294a4bbabf091722, []int{0}
}

func (m *RsaSsaPkcs1Params) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsaSsaPkcs1Params.Unmarshal(m, b)
}
func (m *RsaSsaPkcs1Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsaSsaPkcs1Params.Marshal(b, m, deterministic)
}
func (m *RsaSsaPkcs1Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsaSsaPkcs1Params.Merge(m, src)
}
func (m *RsaSsaPkcs1Params) XXX_Size() int {
	return xxx_messageInfo_RsaSsaPkcs1Params.Size(m)
}
func (m *RsaSsaPkcs1Params) XXX_DiscardUnknown() {
	xxx_messageInfo_RsaSsaPkcs1Params.DiscardUnknown(m)
}

var xxx_messageInfo_RsaSsaPkcs1Params proto.InternalMessageInfo

func (m *RsaSsaPkcs1Params) GetHashType() common_go_proto.HashType {
	if m != nil {
		return m.HashType
	}
	return common_go_proto.HashType_UNKNOWN_HASH
}

// key_type: type.googleapis.com/google.crypto.tink.RsaSsaPkcs1PublicKey
type RsaSsaPkcs1PublicKey struct {
	// Required.
	Version uint32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	// Required.
	Params *RsaSsaPkcs1Params `protobuf:"bytes,2,opt,name=params,proto3" json:"params,omitempty"`
	// Modulus.
	// Unsigned big integer in bigendian representation.
	N []byte `protobuf:"bytes,3,opt,name=n,proto3" json:"n,omitempty"`
	// Public exponent.
	// Unsigned big integer in bigendian representation.
	E                    []byte   `protobuf:"bytes,4,opt,name=e,proto3" json:"e,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsaSsaPkcs1PublicKey) Reset()         { *m = RsaSsaPkcs1PublicKey{} }
func (m *RsaSsaPkcs1PublicKey) String() string { return proto.CompactTextString(m) }
func (*RsaSsaPkcs1PublicKey) ProtoMessage()    {}
func (*RsaSsaPkcs1PublicKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_294a4bbabf091722, []int{1}
}

func (m *RsaSsaPkcs1PublicKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsaSsaPkcs1PublicKey.Unmarshal(m, b)
}
func (m *RsaSsaPkcs1PublicKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsaSsaPkcs1PublicKey.Marshal(b, m, deterministic)
}
func (m *RsaSsaPkcs1PublicKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsaSsaPkcs1PublicKey.Merge(m, src)
}
func (m *RsaSsaPkcs1PublicKey) XXX_Size() int {
	return xxx_messageInfo_RsaSsaPkcs1PublicKey.Size(m)
}
func (m *RsaSsaPkcs1PublicKey) XXX_DiscardUnknown() {
	xxx_messageInfo_RsaSsaPkcs1PublicKey.DiscardUnknown(m)
}

var xxx_messageInfo_RsaSsaPkcs1PublicKey proto.InternalMessageInfo

func (m *RsaSsaPkcs1PublicKey) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *RsaSsaPkcs1PublicKey) GetParams() *RsaSsaPkcs1Params {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *RsaSsaPkcs1PublicKey) GetN() []byte {
	if m != nil {
		return m.N
	}
	return nil
}

func (m *RsaSsaPkcs1PublicKey) GetE() []byte {
	if m != nil {
		return m.E
	}
	return nil
}

// key_type: type.googleapis.com/google.crypto.tink.RsaSsaPkcs1PrivateKey
type RsaSsaPkcs1PrivateKey struct {
	// Required.
	Version uint32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	// Required.
	PublicKey *RsaSsaPkcs1PublicKey `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	// Private exponent.
	// Unsigned big integer in bigendian representation.
	// Required.
	D []byte `protobuf:"bytes,3,opt,name=d,proto3" json:"d,omitempty"`
	// The following parameters are used to optimize RSA signature computation.
	// The prime factor p of n.
	// Unsigned big integer in bigendian representation.
	// Required.
	P []byte `protobuf:"bytes,4,opt,name=p,proto3" json:"p,omitempty"`
	// The prime factor q of n.
	// Unsigned big integer in bigendian representation.
	// Required.
	Q []byte `protobuf:"bytes,5,opt,name=q,proto3" json:"q,omitempty"`
	// d mod (p - 1).
	// Unsigned big integer in bigendian representation.
	// Required.
	Dp []byte `protobuf:"bytes,6,opt,name=dp,proto3" json:"dp,omitempty"`
	// d mod (q - 1).
	// Unsigned big integer in bigendian representation.
	// Required.
	Dq []byte `protobuf:"bytes,7,opt,name=dq,proto3" json:"dq,omitempty"`
	// Chinese Remainder Theorem coefficient q^(-1) mod p.
	// Unsigned big integer in bigendian representation.
	// Required.
	Crt                  []byte   `protobuf:"bytes,8,opt,name=crt,proto3" json:"crt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsaSsaPkcs1PrivateKey) Reset()         { *m = RsaSsaPkcs1PrivateKey{} }
func (m *RsaSsaPkcs1PrivateKey) String() string { return proto.CompactTextString(m) }
func (*RsaSsaPkcs1PrivateKey) ProtoMessage()    {}
func (*RsaSsaPkcs1PrivateKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_294a4bbabf091722, []int{2}
}

func (m *RsaSsaPkcs1PrivateKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsaSsaPkcs1PrivateKey.Unmarshal(m, b)
}
func (m *RsaSsaPkcs1PrivateKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsaSsaPkcs1PrivateKey.Marshal(b, m, deterministic)
}
func (m *RsaSsaPkcs1PrivateKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsaSsaPkcs1PrivateKey.Merge(m, src)
}
func (m *RsaSsaPkcs1PrivateKey) XXX_Size() int {
	return xxx_messageInfo_RsaSsaPkcs1PrivateKey.Size(m)
}
func (m *RsaSsaPkcs1PrivateKey) XXX_DiscardUnknown() {
	xxx_messageInfo_RsaSsaPkcs1PrivateKey.DiscardUnknown(m)
}

var xxx_messageInfo_RsaSsaPkcs1PrivateKey proto.InternalMessageInfo

func (m *RsaSsaPkcs1PrivateKey) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *RsaSsaPkcs1PrivateKey) GetPublicKey() *RsaSsaPkcs1PublicKey {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *RsaSsaPkcs1PrivateKey) GetD() []byte {
	if m != nil {
		return m.D
	}
	return nil
}

func (m *RsaSsaPkcs1PrivateKey) GetP() []byte {
	if m != nil {
		return m.P
	}
	return nil
}

func (m *RsaSsaPkcs1PrivateKey) GetQ() []byte {
	if m != nil {
		return m.Q
	}
	return nil
}

func (m *RsaSsaPkcs1PrivateKey) GetDp() []byte {
	if m != nil {
		return m.Dp
	}
	return nil
}

func (m *RsaSsaPkcs1PrivateKey) GetDq() []byte {
	if m != nil {
		return m.Dq
	}
	return nil
}

func (m *RsaSsaPkcs1PrivateKey) GetCrt() []byte {
	if m != nil {
		return m.Crt
	}
	return nil
}

type RsaSsaPkcs1KeyFormat struct {
	// Required.
	Params *RsaSsaPkcs1Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
	// Required.
	ModulusSizeInBits uint32 `protobuf:"varint,2,opt,name=modulus_size_in_bits,json=modulusSizeInBits,proto3" json:"modulus_size_in_bits,omitempty"`
	// Required.
	PublicExponent       []byte   `protobuf:"bytes,3,opt,name=public_exponent,json=publicExponent,proto3" json:"public_exponent,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsaSsaPkcs1KeyFormat) Reset()         { *m = RsaSsaPkcs1KeyFormat{} }
func (m *RsaSsaPkcs1KeyFormat) String() string { return proto.CompactTextString(m) }
func (*RsaSsaPkcs1KeyFormat) ProtoMessage()    {}
func (*RsaSsaPkcs1KeyFormat) Descriptor() ([]byte, []int) {
	return fileDescriptor_294a4bbabf091722, []int{3}
}

func (m *RsaSsaPkcs1KeyFormat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsaSsaPkcs1KeyFormat.Unmarshal(m, b)
}
func (m *RsaSsaPkcs1KeyFormat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsaSsaPkcs1KeyFormat.Marshal(b, m, deterministic)
}
func (m *RsaSsaPkcs1KeyFormat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsaSsaPkcs1KeyFormat.Merge(m, src)
}
func (m *RsaSsaPkcs1KeyFormat) XXX_Size() int {
	return xxx_messageInfo_RsaSsaPkcs1KeyFormat.Size(m)
}
func (m *RsaSsaPkcs1KeyFormat) XXX_DiscardUnknown() {
	xxx_messageInfo_RsaSsaPkcs1KeyFormat.DiscardUnknown(m)
}

var xxx_messageInfo_RsaSsaPkcs1KeyFormat proto.InternalMessageInfo

func (m *RsaSsaPkcs1KeyFormat) GetParams() *RsaSsaPkcs1Params {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *RsaSsaPkcs1KeyFormat) GetModulusSizeInBits() uint32 {
	if m != nil {
		return m.ModulusSizeInBits
	}
	return 0
}

func (m *RsaSsaPkcs1KeyFormat) GetPublicExponent() []byte {
	if m != nil {
		return m.PublicExponent
	}
	return nil
}

func init() {
	proto.RegisterType((*RsaSsaPkcs1Params)(nil), "google.crypto.tink.RsaSsaPkcs1Params")
	proto.RegisterType((*RsaSsaPkcs1PublicKey)(nil), "google.crypto.tink.RsaSsaPkcs1PublicKey")
	proto.RegisterType((*RsaSsaPkcs1PrivateKey)(nil), "google.crypto.tink.RsaSsaPkcs1PrivateKey")
	proto.RegisterType((*RsaSsaPkcs1KeyFormat)(nil), "google.crypto.tink.RsaSsaPkcs1KeyFormat")
}

func init() {
	proto.RegisterFile("proto/rsa_ssa_pkcs1.proto", fileDescriptor_294a4bbabf091722)
}

var fileDescriptor_294a4bbabf091722 = []byte{
	// 421 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcf, 0x6e, 0xd4, 0x30,
	0x10, 0xc6, 0xe5, 0x2d, 0x6c, 0xdb, 0x61, 0xbb, 0x50, 0xab, 0x48, 0x16, 0xea, 0x61, 0xb5, 0x08,
	0x11, 0x71, 0x48, 0x44, 0x7b, 0xe2, 0xc0, 0xa5, 0x12, 0xff, 0x54, 0x09, 0xad, 0x52, 0xb8, 0x70,
	0xb1, 0x9c, 0xc4, 0x4a, 0xac, 0xdd, 0xd8, 0x8e, 0xed, 0x54, 0xa4, 0xaf, 0xc0, 0xc3, 0xf0, 0x36,
	0x3c, 0x0f, 0x8a, 0xe3, 0xa0, 0x45, 0xed, 0x82, 0xd4, 0x53, 0xfc, 0xfb, 0xe2, 0x99, 0xf9, 0x66,
	0x3c, 0xf0, 0xca, 0x55, 0xc2, 0x14, 0x54, 0x33, 0xe3, 0xba, 0xc4, 0x09, 0xb9, 0x4e, 0xb4, 0x51,
	0x4e, 0x25, 0xc6, 0x32, 0x6a, 0x2d, 0xa3, 0x7a, 0x9d, 0xdb, 0xd7, 0xb1, 0xd7, 0x30, 0x2e, 0x95,
	0x2a, 0x37, 0x3c, 0xce, 0x4d, 0xa7, 0x9d, 0x8a, 0xfb, 0xdb, 0xcf, 0x9e, 0xef, 0x88, 0xcf, 0x55,
	0x5d, 0x2b, 0x39, 0x04, 0x2e, 0x3f, 0xc3, 0x71, 0x6a, 0xd9, 0x95, 0x65, 0xab, 0x3e, 0xdb, 0x8a,
	0x19, 0x56, 0x5b, 0xfc, 0x06, 0x0e, 0x2b, 0x66, 0x2b, 0xea, 0x3a, 0xcd, 0x09, 0x5a, 0xa0, 0x68,
	0x7e, 0x76, 0x1a, 0xdf, 0xae, 0x10, 0x7f, 0x64, 0xb6, 0xfa, 0xd2, 0x69, 0x9e, 0x1e, 0x54, 0xe1,
	0xb4, 0xfc, 0x81, 0xe0, 0x64, 0x3b, 0x61, 0x9b, 0x6d, 0x44, 0x7e, 0xc9, 0x3b, 0x4c, 0x60, 0xff,
	0x9a, 0x1b, 0x2b, 0x94, 0xf4, 0x19, 0x8f, 0xd2, 0x11, 0xf1, 0x5b, 0x98, 0x6a, 0x5f, 0x97, 0x4c,
	0x16, 0x28, 0x7a, 0x74, 0xf6, 0xe2, 0xae, 0x52, 0xb7, 0x4c, 0xa6, 0x21, 0x08, 0xcf, 0x00, 0x49,
	0xb2, 0xb7, 0x40, 0xd1, 0x2c, 0x45, 0xb2, 0x27, 0x4e, 0x1e, 0x0c, 0xc4, 0x97, 0xbf, 0x10, 0x3c,
	0xdd, 0x8e, 0x34, 0xe2, 0x9a, 0x39, 0xfe, 0x6f, 0x3b, 0x1f, 0x00, 0xb4, 0x77, 0x4d, 0xd7, 0xbc,
	0x0b, 0x96, 0xa2, 0xff, 0x59, 0x1a, 0xdb, 0x4c, 0x0f, 0xf5, 0x9f, 0x8e, 0x67, 0x80, 0x8a, 0xd1,
	0x58, 0xd1, 0x93, 0x1e, 0x8d, 0xe9, 0x9e, 0x1a, 0xf2, 0x70, 0xa0, 0x06, 0xcf, 0x61, 0x52, 0x68,
	0x32, 0xf5, 0x38, 0x29, 0xb4, 0xe7, 0x86, 0xec, 0x07, 0x6e, 0xf0, 0x13, 0xd8, 0xcb, 0x8d, 0x23,
	0x07, 0x5e, 0xe8, 0x8f, 0xcb, 0x9f, 0x7f, 0x8f, 0xf9, 0x92, 0x77, 0xef, 0x95, 0xa9, 0x99, 0xdb,
	0x1a, 0x26, 0xba, 0xcf, 0x30, 0x13, 0x38, 0xa9, 0x55, 0xd1, 0x6e, 0x5a, 0x4b, 0xad, 0xb8, 0xe1,
	0x54, 0x48, 0x9a, 0x09, 0x37, 0xbc, 0xcc, 0x51, 0x7a, 0x1c, 0xfe, 0x5d, 0x89, 0x1b, 0xfe, 0x49,
	0x5e, 0x08, 0x67, 0xf1, 0x4b, 0x78, 0x1c, 0xa6, 0xc5, 0xbf, 0x6b, 0x25, 0xb9, 0x74, 0xa1, 0xe5,
	0xf9, 0x20, 0xbf, 0x0b, 0xea, 0xc5, 0x57, 0x38, 0xcd, 0x55, 0x7d, 0x97, 0x1b, 0xbf, 0x88, 0x2b,
	0xf4, 0xed, 0xbc, 0x14, 0xae, 0x6a, 0xb3, 0x38, 0x57, 0x75, 0x32, 0x5c, 0xdb, 0xb9, 0xf5, 0xb4,
	0x54, 0xd4, 0xcb, 0xd9, 0xd4, 0x7f, 0xce, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x50, 0x44, 0x19,
	0xd3, 0x2d, 0x03, 0x00, 0x00,
}
