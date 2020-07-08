// Code generated by protoc-gen-go. DO NOT EDIT.
// source: UserService.proto

package Services

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RegResponse struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegResponse) Reset()         { *m = RegResponse{} }
func (m *RegResponse) String() string { return proto.CompactTextString(m) }
func (*RegResponse) ProtoMessage()    {}
func (*RegResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_UserService_50164160402450a0, []int{0}
}
func (m *RegResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegResponse.Unmarshal(m, b)
}
func (m *RegResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegResponse.Marshal(b, m, deterministic)
}
func (dst *RegResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegResponse.Merge(dst, src)
}
func (m *RegResponse) XXX_Size() int {
	return xxx_messageInfo_RegResponse.Size(m)
}
func (m *RegResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegResponse proto.InternalMessageInfo

func (m *RegResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *RegResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*RegResponse)(nil), "Services.RegResponse")
}

func init() { proto.RegisterFile("UserService.proto", fileDescriptor_UserService_50164160402450a0) }

var fileDescriptor_UserService_50164160402450a0 = []byte{
	// 145 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x0c, 0x2d, 0x4e, 0x2d,
	0x0a, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80,
	0x72, 0x8b, 0xa5, 0x78, 0x7c, 0xf3, 0x53, 0x52, 0x73, 0x8a, 0x21, 0xe2, 0x4a, 0xf6, 0x5c, 0xdc,
	0x41, 0xa9, 0xe9, 0x41, 0xa9, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x62, 0x5c, 0x6c, 0xc5,
	0x25, 0x89, 0x25, 0xa5, 0xc5, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x50, 0x9e, 0x90, 0x04,
	0x17, 0x7b, 0x6e, 0x6a, 0x71, 0x71, 0x62, 0x7a, 0xaa, 0x04, 0x13, 0x58, 0x02, 0xc6, 0x35, 0x72,
	0xe1, 0xe2, 0x46, 0xb2, 0x4d, 0xc8, 0x94, 0x8b, 0x1d, 0xc4, 0x0d, 0x4a, 0x4d, 0x17, 0x12, 0xd6,
	0x83, 0xd9, 0xa9, 0x07, 0x12, 0x02, 0x5b, 0x2b, 0x25, 0x8a, 0x10, 0x44, 0xb2, 0x37, 0x89, 0x0d,
	0xec, 0x1a, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x95, 0xe5, 0xc0, 0xbf, 0xba, 0x00, 0x00,
	0x00,
}
