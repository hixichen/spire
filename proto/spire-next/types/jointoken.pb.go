// Code generated by protoc-gen-go. DO NOT EDIT.
// source: jointoken.proto

package types

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type JoinToken struct {
	// The value of the token.
	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	// The token expiration (seconds since Unix epoch).
	ExpiresAt            int64    `protobuf:"varint,2,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JoinToken) Reset()         { *m = JoinToken{} }
func (m *JoinToken) String() string { return proto.CompactTextString(m) }
func (*JoinToken) ProtoMessage()    {}
func (*JoinToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f6ded0b1862e9d8, []int{0}
}

func (m *JoinToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinToken.Unmarshal(m, b)
}
func (m *JoinToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinToken.Marshal(b, m, deterministic)
}
func (m *JoinToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinToken.Merge(m, src)
}
func (m *JoinToken) XXX_Size() int {
	return xxx_messageInfo_JoinToken.Size(m)
}
func (m *JoinToken) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinToken.DiscardUnknown(m)
}

var xxx_messageInfo_JoinToken proto.InternalMessageInfo

func (m *JoinToken) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *JoinToken) GetExpiresAt() int64 {
	if m != nil {
		return m.ExpiresAt
	}
	return 0
}

func init() {
	proto.RegisterType((*JoinToken)(nil), "spire.types.JoinToken")
}

func init() { proto.RegisterFile("jointoken.proto", fileDescriptor_1f6ded0b1862e9d8) }

var fileDescriptor_1f6ded0b1862e9d8 = []byte{
	// 149 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcf, 0xca, 0xcf, 0xcc,
	0x2b, 0xc9, 0xcf, 0x4e, 0xcd, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x2e, 0x2e, 0xc8,
	0x2c, 0x4a, 0xd5, 0x2b, 0xa9, 0x2c, 0x48, 0x2d, 0x56, 0x72, 0xe0, 0xe2, 0xf4, 0xca, 0xcf, 0xcc,
	0x0b, 0x01, 0xc9, 0x0b, 0x89, 0x70, 0xb1, 0x96, 0x25, 0xe6, 0x94, 0xa6, 0x4a, 0x30, 0x2a, 0x30,
	0x6a, 0x70, 0x06, 0x41, 0x38, 0x42, 0xb2, 0x5c, 0x5c, 0xa9, 0x15, 0x20, 0x2d, 0xc5, 0xf1, 0x89,
	0x25, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0x9c, 0x50, 0x11, 0xc7, 0x12, 0x27, 0x83, 0x28,
	0xbd, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0xfd, 0xe2, 0x82, 0xcc, 0xb4,
	0xb4, 0x54, 0x7d, 0xb0, 0x15, 0xfa, 0x60, 0xfb, 0x20, 0x6c, 0xdd, 0xbc, 0xd4, 0x8a, 0x12, 0x7d,
	0xb0, 0x9d, 0x49, 0x6c, 0x60, 0x71, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xae, 0x65, 0xf0,
	0x73, 0x9a, 0x00, 0x00, 0x00,
}
