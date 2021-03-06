// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Packet.proto

package raft

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

type Packet struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Payload              []byte   `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Packet) Reset()         { *m = Packet{} }
func (m *Packet) String() string { return proto.CompactTextString(m) }
func (*Packet) ProtoMessage()    {}
func (*Packet) Descriptor() ([]byte, []int) {
	return fileDescriptor_365103b9ca9fe667, []int{0}
}

func (m *Packet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Packet.Unmarshal(m, b)
}
func (m *Packet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Packet.Marshal(b, m, deterministic)
}
func (m *Packet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Packet.Merge(m, src)
}
func (m *Packet) XXX_Size() int {
	return xxx_messageInfo_Packet.Size(m)
}
func (m *Packet) XXX_DiscardUnknown() {
	xxx_messageInfo_Packet.DiscardUnknown(m)
}

var xxx_messageInfo_Packet proto.InternalMessageInfo

func (m *Packet) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Packet) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterType((*Packet)(nil), "raft.Packet")
}

func init() { proto.RegisterFile("Packet.proto", fileDescriptor_365103b9ca9fe667) }

var fileDescriptor_365103b9ca9fe667 = []byte{
	// 91 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x09, 0x48, 0x4c, 0xce,
	0x4e, 0x2d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x29, 0x4a, 0x4c, 0x2b, 0x51, 0x32,
	0xe2, 0x62, 0x83, 0x88, 0x0a, 0xf1, 0x71, 0x31, 0x65, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0xf0,
	0x06, 0x31, 0x65, 0xa6, 0x08, 0x49, 0x70, 0xb1, 0x17, 0x24, 0x56, 0xe6, 0xe4, 0x27, 0xa6, 0x48,
	0x30, 0x29, 0x30, 0x6a, 0xf0, 0x04, 0xc1, 0xb8, 0x49, 0x6c, 0x60, 0x03, 0x8c, 0x01, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xfd, 0x06, 0x9d, 0x44, 0x50, 0x00, 0x00, 0x00,
}
