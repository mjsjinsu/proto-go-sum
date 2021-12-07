// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/sum/sum.proto

package sum

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

type SumArgs struct {
	Value1               int32    `protobuf:"varint,1,opt,name=value1,proto3" json:"value1,omitempty"`
	Value2               int32    `protobuf:"varint,2,opt,name=value2,proto3" json:"value2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumArgs) Reset()         { *m = SumArgs{} }
func (m *SumArgs) String() string { return proto.CompactTextString(m) }
func (*SumArgs) ProtoMessage()    {}
func (*SumArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd2d4a23a39ebd3f, []int{0}
}

func (m *SumArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumArgs.Unmarshal(m, b)
}
func (m *SumArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumArgs.Marshal(b, m, deterministic)
}
func (m *SumArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumArgs.Merge(m, src)
}
func (m *SumArgs) XXX_Size() int {
	return xxx_messageInfo_SumArgs.Size(m)
}
func (m *SumArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_SumArgs.DiscardUnknown(m)
}

var xxx_messageInfo_SumArgs proto.InternalMessageInfo

func (m *SumArgs) GetValue1() int32 {
	if m != nil {
		return m.Value1
	}
	return 0
}

func (m *SumArgs) GetValue2() int32 {
	if m != nil {
		return m.Value2
	}
	return 0
}

type SumReturns struct {
	Value                int32    `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumReturns) Reset()         { *m = SumReturns{} }
func (m *SumReturns) String() string { return proto.CompactTextString(m) }
func (*SumReturns) ProtoMessage()    {}
func (*SumReturns) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd2d4a23a39ebd3f, []int{1}
}

func (m *SumReturns) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumReturns.Unmarshal(m, b)
}
func (m *SumReturns) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumReturns.Marshal(b, m, deterministic)
}
func (m *SumReturns) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumReturns.Merge(m, src)
}
func (m *SumReturns) XXX_Size() int {
	return xxx_messageInfo_SumReturns.Size(m)
}
func (m *SumReturns) XXX_DiscardUnknown() {
	xxx_messageInfo_SumReturns.DiscardUnknown(m)
}

var xxx_messageInfo_SumReturns proto.InternalMessageInfo

func (m *SumReturns) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterType((*SumArgs)(nil), "sum.SumArgs")
	proto.RegisterType((*SumReturns)(nil), "sum.SumReturns")
}

func init() { proto.RegisterFile("protos/sum/sum.proto", fileDescriptor_bd2d4a23a39ebd3f) }

var fileDescriptor_bd2d4a23a39ebd3f = []byte{
	// 181 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x29, 0x28, 0xca, 0x2f,
	0xc9, 0x2f, 0xd6, 0x2f, 0x2e, 0xcd, 0x05, 0x61, 0x3d, 0x30, 0x57, 0x88, 0xb9, 0xb8, 0x34, 0x57,
	0xc9, 0x92, 0x8b, 0x3d, 0xb8, 0x34, 0xd7, 0xb1, 0x28, 0xbd, 0x58, 0x48, 0x8c, 0x8b, 0xad, 0x2c,
	0x31, 0xa7, 0x34, 0xd5, 0x50, 0x82, 0x51, 0x81, 0x51, 0x83, 0x35, 0x08, 0xca, 0x83, 0x8b, 0x1b,
	0x49, 0x30, 0x21, 0x89, 0x1b, 0x29, 0x29, 0x71, 0x71, 0x05, 0x97, 0xe6, 0x06, 0xa5, 0x96, 0x94,
	0x16, 0xe5, 0x15, 0x0b, 0x89, 0x70, 0xb1, 0x82, 0xc5, 0xa1, 0x9a, 0x21, 0x1c, 0x23, 0x53, 0x2e,
	0xde, 0xf4, 0xa0, 0x00, 0x67, 0xb7, 0xd2, 0xbc, 0xe4, 0x62, 0x9f, 0xcc, 0xe2, 0x12, 0x21, 0x15,
	0x2e, 0x90, 0xb5, 0x42, 0x3c, 0x7a, 0x20, 0x77, 0x40, 0x6d, 0x96, 0xe2, 0x87, 0xf1, 0xa0, 0x86,
	0x39, 0xe9, 0x46, 0x69, 0xa7, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0xe7,
	0x66, 0x15, 0x67, 0x65, 0xe6, 0x15, 0x97, 0xea, 0x83, 0xdd, 0xad, 0x9b, 0x9e, 0xaf, 0x0b, 0xf2,
	0x08, 0xc2, 0x4f, 0x49, 0x6c, 0x60, 0xb6, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xac, 0xc1, 0x24,
	0x0d, 0xe8, 0x00, 0x00, 0x00,
}