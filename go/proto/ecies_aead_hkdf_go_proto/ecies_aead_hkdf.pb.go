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
// source: third_party/tink/proto/ecies_aead_hkdf.proto

package ecies_aead_hkdf_go_proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	common_go_proto "github.com/nandawangsa/tink/go/proto/common_go_proto"
	tink_go_proto "github.com/nandawangsa/tink/go/proto/tink_go_proto"
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

// Parameters of KEM (Key Encapsulation Mechanism)
type EciesHkdfKemParams struct {
	// Required.
	CurveType common_go_proto.EllipticCurveType `protobuf:"varint,1,opt,name=curve_type,json=curveType,proto3,enum=google.crypto.tink.EllipticCurveType" json:"curve_type,omitempty"`
	// Required.
	HkdfHashType common_go_proto.HashType `protobuf:"varint,2,opt,name=hkdf_hash_type,json=hkdfHashType,proto3,enum=google.crypto.tink.HashType" json:"hkdf_hash_type,omitempty"`
	// Optional.
	HkdfSalt             []byte   `protobuf:"bytes,11,opt,name=hkdf_salt,json=hkdfSalt,proto3" json:"hkdf_salt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EciesHkdfKemParams) Reset()         { *m = EciesHkdfKemParams{} }
func (m *EciesHkdfKemParams) String() string { return proto.CompactTextString(m) }
func (*EciesHkdfKemParams) ProtoMessage()    {}
func (*EciesHkdfKemParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8ddf7d5f8978761, []int{0}
}

func (m *EciesHkdfKemParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EciesHkdfKemParams.Unmarshal(m, b)
}
func (m *EciesHkdfKemParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EciesHkdfKemParams.Marshal(b, m, deterministic)
}
func (m *EciesHkdfKemParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EciesHkdfKemParams.Merge(m, src)
}
func (m *EciesHkdfKemParams) XXX_Size() int {
	return xxx_messageInfo_EciesHkdfKemParams.Size(m)
}
func (m *EciesHkdfKemParams) XXX_DiscardUnknown() {
	xxx_messageInfo_EciesHkdfKemParams.DiscardUnknown(m)
}

var xxx_messageInfo_EciesHkdfKemParams proto.InternalMessageInfo

func (m *EciesHkdfKemParams) GetCurveType() common_go_proto.EllipticCurveType {
	if m != nil {
		return m.CurveType
	}
	return common_go_proto.EllipticCurveType_UNKNOWN_CURVE
}

func (m *EciesHkdfKemParams) GetHkdfHashType() common_go_proto.HashType {
	if m != nil {
		return m.HkdfHashType
	}
	return common_go_proto.HashType_UNKNOWN_HASH
}

func (m *EciesHkdfKemParams) GetHkdfSalt() []byte {
	if m != nil {
		return m.HkdfSalt
	}
	return nil
}

// Parameters of AEAD DEM (Data Encapsulation Mechanism).
type EciesAeadDemParams struct {
	// Required.
	AeadDem              *tink_go_proto.KeyTemplate `protobuf:"bytes,2,opt,name=aead_dem,json=aeadDem,proto3" json:"aead_dem,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *EciesAeadDemParams) Reset()         { *m = EciesAeadDemParams{} }
func (m *EciesAeadDemParams) String() string { return proto.CompactTextString(m) }
func (*EciesAeadDemParams) ProtoMessage()    {}
func (*EciesAeadDemParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8ddf7d5f8978761, []int{1}
}

func (m *EciesAeadDemParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EciesAeadDemParams.Unmarshal(m, b)
}
func (m *EciesAeadDemParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EciesAeadDemParams.Marshal(b, m, deterministic)
}
func (m *EciesAeadDemParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EciesAeadDemParams.Merge(m, src)
}
func (m *EciesAeadDemParams) XXX_Size() int {
	return xxx_messageInfo_EciesAeadDemParams.Size(m)
}
func (m *EciesAeadDemParams) XXX_DiscardUnknown() {
	xxx_messageInfo_EciesAeadDemParams.DiscardUnknown(m)
}

var xxx_messageInfo_EciesAeadDemParams proto.InternalMessageInfo

func (m *EciesAeadDemParams) GetAeadDem() *tink_go_proto.KeyTemplate {
	if m != nil {
		return m.AeadDem
	}
	return nil
}

type EciesAeadHkdfParams struct {
	// Key Encapsulation Mechanism.
	// Required.
	KemParams *EciesHkdfKemParams `protobuf:"bytes,1,opt,name=kem_params,json=kemParams,proto3" json:"kem_params,omitempty"`
	// Data Encapsulation Mechanism.
	// Required.
	DemParams *EciesAeadDemParams `protobuf:"bytes,2,opt,name=dem_params,json=demParams,proto3" json:"dem_params,omitempty"`
	// EC point format.
	// Required.
	EcPointFormat        common_go_proto.EcPointFormat `protobuf:"varint,3,opt,name=ec_point_format,json=ecPointFormat,proto3,enum=google.crypto.tink.EcPointFormat" json:"ec_point_format,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *EciesAeadHkdfParams) Reset()         { *m = EciesAeadHkdfParams{} }
func (m *EciesAeadHkdfParams) String() string { return proto.CompactTextString(m) }
func (*EciesAeadHkdfParams) ProtoMessage()    {}
func (*EciesAeadHkdfParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8ddf7d5f8978761, []int{2}
}

func (m *EciesAeadHkdfParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EciesAeadHkdfParams.Unmarshal(m, b)
}
func (m *EciesAeadHkdfParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EciesAeadHkdfParams.Marshal(b, m, deterministic)
}
func (m *EciesAeadHkdfParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EciesAeadHkdfParams.Merge(m, src)
}
func (m *EciesAeadHkdfParams) XXX_Size() int {
	return xxx_messageInfo_EciesAeadHkdfParams.Size(m)
}
func (m *EciesAeadHkdfParams) XXX_DiscardUnknown() {
	xxx_messageInfo_EciesAeadHkdfParams.DiscardUnknown(m)
}

var xxx_messageInfo_EciesAeadHkdfParams proto.InternalMessageInfo

func (m *EciesAeadHkdfParams) GetKemParams() *EciesHkdfKemParams {
	if m != nil {
		return m.KemParams
	}
	return nil
}

func (m *EciesAeadHkdfParams) GetDemParams() *EciesAeadDemParams {
	if m != nil {
		return m.DemParams
	}
	return nil
}

func (m *EciesAeadHkdfParams) GetEcPointFormat() common_go_proto.EcPointFormat {
	if m != nil {
		return m.EcPointFormat
	}
	return common_go_proto.EcPointFormat_UNKNOWN_FORMAT
}

// EciesAeadHkdfPublicKey represents HybridEncryption primitive.
// key_type: type.googleapis.com/google.crypto.tink.EciesAeadHkdfPublicKey
type EciesAeadHkdfPublicKey struct {
	// Required.
	Version uint32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	// Required.
	Params *EciesAeadHkdfParams `protobuf:"bytes,2,opt,name=params,proto3" json:"params,omitempty"`
	// Affine coordinates of the public key in bigendian representation.
	// The public key is a point (x, y) on the curve defined by params.kem_params.curve.
	// Required.
	X []byte `protobuf:"bytes,3,opt,name=x,proto3" json:"x,omitempty"`
	// Required.
	Y                    []byte   `protobuf:"bytes,4,opt,name=y,proto3" json:"y,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EciesAeadHkdfPublicKey) Reset()         { *m = EciesAeadHkdfPublicKey{} }
func (m *EciesAeadHkdfPublicKey) String() string { return proto.CompactTextString(m) }
func (*EciesAeadHkdfPublicKey) ProtoMessage()    {}
func (*EciesAeadHkdfPublicKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8ddf7d5f8978761, []int{3}
}

func (m *EciesAeadHkdfPublicKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EciesAeadHkdfPublicKey.Unmarshal(m, b)
}
func (m *EciesAeadHkdfPublicKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EciesAeadHkdfPublicKey.Marshal(b, m, deterministic)
}
func (m *EciesAeadHkdfPublicKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EciesAeadHkdfPublicKey.Merge(m, src)
}
func (m *EciesAeadHkdfPublicKey) XXX_Size() int {
	return xxx_messageInfo_EciesAeadHkdfPublicKey.Size(m)
}
func (m *EciesAeadHkdfPublicKey) XXX_DiscardUnknown() {
	xxx_messageInfo_EciesAeadHkdfPublicKey.DiscardUnknown(m)
}

var xxx_messageInfo_EciesAeadHkdfPublicKey proto.InternalMessageInfo

func (m *EciesAeadHkdfPublicKey) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *EciesAeadHkdfPublicKey) GetParams() *EciesAeadHkdfParams {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *EciesAeadHkdfPublicKey) GetX() []byte {
	if m != nil {
		return m.X
	}
	return nil
}

func (m *EciesAeadHkdfPublicKey) GetY() []byte {
	if m != nil {
		return m.Y
	}
	return nil
}

// EciesKdfAeadPrivateKey represents HybridDecryption primitive.
// key_type: type.googleapis.com/google.crypto.tink.EciesAeadHkdfPrivateKey
type EciesAeadHkdfPrivateKey struct {
	// Required.
	Version uint32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	// Required.
	PublicKey *EciesAeadHkdfPublicKey `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	// Required.
	KeyValue             []byte   `protobuf:"bytes,3,opt,name=key_value,json=keyValue,proto3" json:"key_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EciesAeadHkdfPrivateKey) Reset()         { *m = EciesAeadHkdfPrivateKey{} }
func (m *EciesAeadHkdfPrivateKey) String() string { return proto.CompactTextString(m) }
func (*EciesAeadHkdfPrivateKey) ProtoMessage()    {}
func (*EciesAeadHkdfPrivateKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8ddf7d5f8978761, []int{4}
}

func (m *EciesAeadHkdfPrivateKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EciesAeadHkdfPrivateKey.Unmarshal(m, b)
}
func (m *EciesAeadHkdfPrivateKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EciesAeadHkdfPrivateKey.Marshal(b, m, deterministic)
}
func (m *EciesAeadHkdfPrivateKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EciesAeadHkdfPrivateKey.Merge(m, src)
}
func (m *EciesAeadHkdfPrivateKey) XXX_Size() int {
	return xxx_messageInfo_EciesAeadHkdfPrivateKey.Size(m)
}
func (m *EciesAeadHkdfPrivateKey) XXX_DiscardUnknown() {
	xxx_messageInfo_EciesAeadHkdfPrivateKey.DiscardUnknown(m)
}

var xxx_messageInfo_EciesAeadHkdfPrivateKey proto.InternalMessageInfo

func (m *EciesAeadHkdfPrivateKey) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *EciesAeadHkdfPrivateKey) GetPublicKey() *EciesAeadHkdfPublicKey {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *EciesAeadHkdfPrivateKey) GetKeyValue() []byte {
	if m != nil {
		return m.KeyValue
	}
	return nil
}

type EciesAeadHkdfKeyFormat struct {
	// Required.
	Params               *EciesAeadHkdfParams `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *EciesAeadHkdfKeyFormat) Reset()         { *m = EciesAeadHkdfKeyFormat{} }
func (m *EciesAeadHkdfKeyFormat) String() string { return proto.CompactTextString(m) }
func (*EciesAeadHkdfKeyFormat) ProtoMessage()    {}
func (*EciesAeadHkdfKeyFormat) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8ddf7d5f8978761, []int{5}
}

func (m *EciesAeadHkdfKeyFormat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EciesAeadHkdfKeyFormat.Unmarshal(m, b)
}
func (m *EciesAeadHkdfKeyFormat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EciesAeadHkdfKeyFormat.Marshal(b, m, deterministic)
}
func (m *EciesAeadHkdfKeyFormat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EciesAeadHkdfKeyFormat.Merge(m, src)
}
func (m *EciesAeadHkdfKeyFormat) XXX_Size() int {
	return xxx_messageInfo_EciesAeadHkdfKeyFormat.Size(m)
}
func (m *EciesAeadHkdfKeyFormat) XXX_DiscardUnknown() {
	xxx_messageInfo_EciesAeadHkdfKeyFormat.DiscardUnknown(m)
}

var xxx_messageInfo_EciesAeadHkdfKeyFormat proto.InternalMessageInfo

func (m *EciesAeadHkdfKeyFormat) GetParams() *EciesAeadHkdfParams {
	if m != nil {
		return m.Params
	}
	return nil
}

func init() {
	proto.RegisterType((*EciesHkdfKemParams)(nil), "google.crypto.tink.EciesHkdfKemParams")
	proto.RegisterType((*EciesAeadDemParams)(nil), "google.crypto.tink.EciesAeadDemParams")
	proto.RegisterType((*EciesAeadHkdfParams)(nil), "google.crypto.tink.EciesAeadHkdfParams")
	proto.RegisterType((*EciesAeadHkdfPublicKey)(nil), "google.crypto.tink.EciesAeadHkdfPublicKey")
	proto.RegisterType((*EciesAeadHkdfPrivateKey)(nil), "google.crypto.tink.EciesAeadHkdfPrivateKey")
	proto.RegisterType((*EciesAeadHkdfKeyFormat)(nil), "google.crypto.tink.EciesAeadHkdfKeyFormat")
}

func init() {
	proto.RegisterFile("proto/ecies_aead_hkdf.proto", fileDescriptor_a8ddf7d5f8978761)
}

var fileDescriptor_a8ddf7d5f8978761 = []byte{
	// 520 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xd1, 0x6a, 0x13, 0x41,
	0x14, 0x86, 0x99, 0x28, 0x69, 0xf6, 0x24, 0xad, 0xb0, 0x82, 0x86, 0xb6, 0x60, 0xbb, 0xa2, 0x16,
	0x91, 0x0d, 0x44, 0xbc, 0xf1, 0x46, 0x8c, 0x8d, 0x34, 0x04, 0x24, 0xac, 0x41, 0xd0, 0x9b, 0x71,
	0xb2, 0x7b, 0x92, 0x5d, 0x76, 0x27, 0x33, 0xec, 0x4e, 0x42, 0xe7, 0x29, 0x7c, 0x00, 0xef, 0x7c,
	0x0d, 0xdf, 0xc9, 0x67, 0x90, 0x99, 0x4d, 0xd2, 0xc6, 0x6e, 0x53, 0xf0, 0x6e, 0xce, 0xe4, 0x3f,
	0xdf, 0xfc, 0xff, 0x0f, 0x59, 0x78, 0xa5, 0xe2, 0x24, 0x8f, 0xa8, 0x64, 0xb9, 0xd2, 0x1d, 0x95,
	0xcc, 0xd3, 0x8e, 0xcc, 0x85, 0x12, 0x1d, 0x0c, 0x13, 0x2c, 0x28, 0x43, 0x16, 0xd1, 0x38, 0x8d,
	0xa6, 0xbe, 0xbd, 0x75, 0xdd, 0x99, 0x10, 0xb3, 0x0c, 0xfd, 0x30, 0xd7, 0x52, 0x09, 0xdf, 0xe8,
	0x0f, 0x9f, 0xde, 0x42, 0x08, 0x05, 0xe7, 0x62, 0x5e, 0x2e, 0x1e, 0x9e, 0xde, 0x22, 0x32, 0xc7,
	0x52, 0xe2, 0xfd, 0x26, 0xe0, 0xf6, 0xcd, 0xab, 0x17, 0x69, 0x34, 0x1d, 0x22, 0x1f, 0xb1, 0x9c,
	0xf1, 0xc2, 0x3d, 0x07, 0x08, 0x17, 0xf9, 0x12, 0xa9, 0xd2, 0x12, 0xdb, 0xe4, 0x84, 0x9c, 0x1d,
	0x74, 0x9f, 0xf9, 0x37, 0x7d, 0xf8, 0xfd, 0x2c, 0x4b, 0xa4, 0x4a, 0xc2, 0x0f, 0x46, 0x3d, 0xd6,
	0x12, 0x03, 0x27, 0x5c, 0x1f, 0xdd, 0x1e, 0x1c, 0x98, 0x18, 0x34, 0x66, 0x45, 0x5c, 0x92, 0x6a,
	0x96, 0x74, 0x5c, 0x45, 0xba, 0x60, 0x45, 0x6c, 0x01, 0x2d, 0xb3, 0xb3, 0x9e, 0xdc, 0x23, 0x70,
	0x2c, 0xa3, 0x60, 0x99, 0x6a, 0x37, 0x4f, 0xc8, 0x59, 0x2b, 0x68, 0x98, 0x8b, 0xcf, 0x2c, 0x53,
	0xde, 0x68, 0x65, 0xfe, 0x3d, 0xb2, 0xe8, 0x7c, 0x63, 0xfe, 0x2d, 0x34, 0x6c, 0x85, 0x11, 0x72,
	0xfb, 0x60, 0xb3, 0xfb, 0xa4, 0xea, 0xc1, 0x21, 0xea, 0x31, 0x72, 0x99, 0x31, 0x85, 0xc1, 0x1e,
	0x2b, 0x09, 0xde, 0x1f, 0x02, 0x0f, 0x37, 0x48, 0xd3, 0xc9, 0x8a, 0xd9, 0x07, 0x48, 0x91, 0x9b,
	0x2a, 0x19, 0x2f, 0x6c, 0x21, 0xcd, 0xee, 0xf3, 0xca, 0x42, 0x6e, 0x94, 0x19, 0x38, 0xe9, 0xc6,
	0x5a, 0x1f, 0x20, 0xba, 0xc2, 0xd4, 0xee, 0xc0, 0x6c, 0xc5, 0x0a, 0x9c, 0x68, 0x83, 0x19, 0xc0,
	0x03, 0x0c, 0xa9, 0x14, 0xc9, 0x5c, 0xd1, 0xa9, 0xc8, 0x39, 0x53, 0xed, 0x7b, 0xb6, 0xd9, 0xd3,
	0x6a, 0xd6, 0xc8, 0x28, 0x3f, 0x5a, 0x61, 0xb0, 0x8f, 0xd7, 0x47, 0xef, 0x07, 0x81, 0x47, 0xdb,
	0x81, 0x17, 0x93, 0x2c, 0x09, 0x87, 0xa8, 0xdd, 0x36, 0xec, 0x2d, 0x31, 0x2f, 0x12, 0x31, 0xb7,
	0x81, 0xf7, 0x83, 0xf5, 0xe8, 0xbe, 0x83, 0xfa, 0x56, 0x84, 0x17, 0x3b, 0x23, 0x5c, 0xd5, 0x18,
	0xac, 0xd6, 0xdc, 0x16, 0x90, 0x4b, 0x6b, 0xb9, 0x15, 0x90, 0x4b, 0x33, 0xe9, 0xf6, 0xfd, 0x72,
	0xd2, 0xde, 0x4f, 0x02, 0x8f, 0xb7, 0x77, 0xf3, 0x64, 0xc9, 0x14, 0xee, 0xb6, 0x34, 0x00, 0x90,
	0xd6, 0x39, 0x4d, 0x51, 0xaf, 0x6c, 0xbd, 0xbc, 0xdb, 0xd6, 0x3a, 0x6c, 0xe0, 0xc8, 0x4d, 0xee,
	0x23, 0x70, 0x52, 0xd4, 0x74, 0xc9, 0xb2, 0x05, 0xae, 0x4c, 0x36, 0x52, 0xd4, 0x5f, 0xcc, 0xec,
	0x7d, 0xfd, 0xa7, 0xae, 0x21, 0xea, 0xb2, 0xc9, 0x6b, 0xa5, 0x90, 0xff, 0x2a, 0xa5, 0xf7, 0x1d,
	0x8e, 0x43, 0xc1, 0xab, 0xb6, 0xec, 0x7f, 0x75, 0x44, 0xbe, 0xbd, 0x99, 0x25, 0x2a, 0x5e, 0x4c,
	0xfc, 0x50, 0xf0, 0x4e, 0x29, 0xdb, 0xf1, 0xf5, 0xa0, 0x33, 0x41, 0xed, 0x0f, 0xbf, 0x6a, 0xf5,
	0xf1, 0xe0, 0xd3, 0x70, 0xd4, 0x9b, 0xd4, 0xed, 0xfc, 0xfa, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x80, 0x6f, 0xd1, 0xc8, 0x80, 0x04, 0x00, 0x00,
}
