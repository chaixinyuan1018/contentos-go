// Code generated by protoc-gen-go. DO NOT EDIT.
// source: app/table/so_extDailyTrx.proto

package table

import (
	fmt "fmt"
	prototype "github.com/coschain/contentos-go/prototype"
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

type SoExtDailyTrx struct {
	Date                 *prototype.TimePointSec `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Count                uint64                  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *SoExtDailyTrx) Reset()         { *m = SoExtDailyTrx{} }
func (m *SoExtDailyTrx) String() string { return proto.CompactTextString(m) }
func (*SoExtDailyTrx) ProtoMessage()    {}
func (*SoExtDailyTrx) Descriptor() ([]byte, []int) {
	return fileDescriptor_a71b07e7e9721030, []int{0}
}

func (m *SoExtDailyTrx) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoExtDailyTrx.Unmarshal(m, b)
}
func (m *SoExtDailyTrx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoExtDailyTrx.Marshal(b, m, deterministic)
}
func (m *SoExtDailyTrx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoExtDailyTrx.Merge(m, src)
}
func (m *SoExtDailyTrx) XXX_Size() int {
	return xxx_messageInfo_SoExtDailyTrx.Size(m)
}
func (m *SoExtDailyTrx) XXX_DiscardUnknown() {
	xxx_messageInfo_SoExtDailyTrx.DiscardUnknown(m)
}

var xxx_messageInfo_SoExtDailyTrx proto.InternalMessageInfo

func (m *SoExtDailyTrx) GetDate() *prototype.TimePointSec {
	if m != nil {
		return m.Date
	}
	return nil
}

func (m *SoExtDailyTrx) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type SoListExtDailyTrxByDate struct {
	Date                 *prototype.TimePointSec `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *SoListExtDailyTrxByDate) Reset()         { *m = SoListExtDailyTrxByDate{} }
func (m *SoListExtDailyTrxByDate) String() string { return proto.CompactTextString(m) }
func (*SoListExtDailyTrxByDate) ProtoMessage()    {}
func (*SoListExtDailyTrxByDate) Descriptor() ([]byte, []int) {
	return fileDescriptor_a71b07e7e9721030, []int{1}
}

func (m *SoListExtDailyTrxByDate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoListExtDailyTrxByDate.Unmarshal(m, b)
}
func (m *SoListExtDailyTrxByDate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoListExtDailyTrxByDate.Marshal(b, m, deterministic)
}
func (m *SoListExtDailyTrxByDate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoListExtDailyTrxByDate.Merge(m, src)
}
func (m *SoListExtDailyTrxByDate) XXX_Size() int {
	return xxx_messageInfo_SoListExtDailyTrxByDate.Size(m)
}
func (m *SoListExtDailyTrxByDate) XXX_DiscardUnknown() {
	xxx_messageInfo_SoListExtDailyTrxByDate.DiscardUnknown(m)
}

var xxx_messageInfo_SoListExtDailyTrxByDate proto.InternalMessageInfo

func (m *SoListExtDailyTrxByDate) GetDate() *prototype.TimePointSec {
	if m != nil {
		return m.Date
	}
	return nil
}

type SoListExtDailyTrxByCount struct {
	Count                uint64                  `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Date                 *prototype.TimePointSec `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *SoListExtDailyTrxByCount) Reset()         { *m = SoListExtDailyTrxByCount{} }
func (m *SoListExtDailyTrxByCount) String() string { return proto.CompactTextString(m) }
func (*SoListExtDailyTrxByCount) ProtoMessage()    {}
func (*SoListExtDailyTrxByCount) Descriptor() ([]byte, []int) {
	return fileDescriptor_a71b07e7e9721030, []int{2}
}

func (m *SoListExtDailyTrxByCount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoListExtDailyTrxByCount.Unmarshal(m, b)
}
func (m *SoListExtDailyTrxByCount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoListExtDailyTrxByCount.Marshal(b, m, deterministic)
}
func (m *SoListExtDailyTrxByCount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoListExtDailyTrxByCount.Merge(m, src)
}
func (m *SoListExtDailyTrxByCount) XXX_Size() int {
	return xxx_messageInfo_SoListExtDailyTrxByCount.Size(m)
}
func (m *SoListExtDailyTrxByCount) XXX_DiscardUnknown() {
	xxx_messageInfo_SoListExtDailyTrxByCount.DiscardUnknown(m)
}

var xxx_messageInfo_SoListExtDailyTrxByCount proto.InternalMessageInfo

func (m *SoListExtDailyTrxByCount) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *SoListExtDailyTrxByCount) GetDate() *prototype.TimePointSec {
	if m != nil {
		return m.Date
	}
	return nil
}

type SoUniqueExtDailyTrxByDate struct {
	Date                 *prototype.TimePointSec `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *SoUniqueExtDailyTrxByDate) Reset()         { *m = SoUniqueExtDailyTrxByDate{} }
func (m *SoUniqueExtDailyTrxByDate) String() string { return proto.CompactTextString(m) }
func (*SoUniqueExtDailyTrxByDate) ProtoMessage()    {}
func (*SoUniqueExtDailyTrxByDate) Descriptor() ([]byte, []int) {
	return fileDescriptor_a71b07e7e9721030, []int{3}
}

func (m *SoUniqueExtDailyTrxByDate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoUniqueExtDailyTrxByDate.Unmarshal(m, b)
}
func (m *SoUniqueExtDailyTrxByDate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoUniqueExtDailyTrxByDate.Marshal(b, m, deterministic)
}
func (m *SoUniqueExtDailyTrxByDate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoUniqueExtDailyTrxByDate.Merge(m, src)
}
func (m *SoUniqueExtDailyTrxByDate) XXX_Size() int {
	return xxx_messageInfo_SoUniqueExtDailyTrxByDate.Size(m)
}
func (m *SoUniqueExtDailyTrxByDate) XXX_DiscardUnknown() {
	xxx_messageInfo_SoUniqueExtDailyTrxByDate.DiscardUnknown(m)
}

var xxx_messageInfo_SoUniqueExtDailyTrxByDate proto.InternalMessageInfo

func (m *SoUniqueExtDailyTrxByDate) GetDate() *prototype.TimePointSec {
	if m != nil {
		return m.Date
	}
	return nil
}

func init() {
	proto.RegisterType((*SoExtDailyTrx)(nil), "table.so_extDailyTrx")
	proto.RegisterType((*SoListExtDailyTrxByDate)(nil), "table.so_list_extDailyTrx_by_date")
	proto.RegisterType((*SoListExtDailyTrxByCount)(nil), "table.so_list_extDailyTrx_by_count")
	proto.RegisterType((*SoUniqueExtDailyTrxByDate)(nil), "table.so_unique_extDailyTrx_by_date")
}

func init() { proto.RegisterFile("app/table/so_extDailyTrx.proto", fileDescriptor_a71b07e7e9721030) }

var fileDescriptor_a71b07e7e9721030 = []byte{
	// 232 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x91, 0x4f, 0x4b, 0xc4, 0x30,
	0x14, 0xc4, 0xc9, 0xb2, 0xeb, 0x21, 0x82, 0x87, 0xb2, 0x87, 0xf5, 0x2f, 0x4b, 0x4f, 0x8b, 0xb8,
	0x09, 0xe8, 0x37, 0x10, 0x8f, 0xe2, 0x61, 0xd1, 0x8b, 0x97, 0x90, 0xc4, 0x47, 0x1b, 0x68, 0xf3,
	0x62, 0xf3, 0x02, 0xed, 0xb7, 0x17, 0x53, 0xa8, 0xf5, 0x20, 0x28, 0x78, 0x09, 0x0c, 0x93, 0x99,
	0xdf, 0xc0, 0xe3, 0x57, 0x3a, 0x04, 0x49, 0xda, 0x34, 0x20, 0x23, 0x2a, 0xe8, 0xe9, 0x41, 0xbb,
	0x66, 0x78, 0xee, 0x7a, 0x11, 0x3a, 0x24, 0x2c, 0x56, 0xd9, 0x3b, 0x5b, 0x67, 0x45, 0x43, 0x00,
	0xf9, 0xf9, 0x8c, 0x66, 0xf9, 0xc2, 0x4f, 0xbe, 0x87, 0x8a, 0x3d, 0x5f, 0xbe, 0x69, 0x82, 0x0d,
	0xdb, 0xb2, 0xdd, 0xf1, 0xed, 0xa9, 0x98, 0x62, 0x82, 0x5c, 0x0b, 0x2a, 0xa0, 0xf3, 0xa4, 0x22,
	0xd8, 0x43, 0xfe, 0x56, 0xac, 0xf9, 0xca, 0x62, 0xf2, 0xb4, 0x59, 0x6c, 0xd9, 0x6e, 0x79, 0x18,
	0x45, 0xf9, 0xc8, 0xcf, 0x23, 0xaa, 0xc6, 0x45, 0x9a, 0x77, 0x2b, 0x33, 0xa8, 0x1c, 0xfa, 0x1b,
	0xa3, 0xb4, 0xfc, 0xe2, 0x87, 0xb6, 0x4c, 0xfb, 0xda, 0xc0, 0x66, 0x1b, 0x26, 0xc8, 0xe2, 0x77,
	0x90, 0x27, 0x7e, 0x19, 0x51, 0x25, 0xef, 0xde, 0x13, 0xfc, 0xc3, 0xe8, 0xfb, 0x9b, 0xd7, 0xeb,
	0xca, 0x51, 0x9d, 0x8c, 0xb0, 0xd8, 0x4a, 0x8b, 0xd1, 0xd6, 0xda, 0x79, 0x69, 0xd1, 0x13, 0x78,
	0xc2, 0xb8, 0xaf, 0x50, 0x4e, 0x97, 0x33, 0x47, 0xb9, 0xed, 0xee, 0x23, 0x00, 0x00, 0xff, 0xff,
	0xd9, 0xf4, 0x43, 0x2a, 0xcd, 0x01, 0x00, 0x00,
}
