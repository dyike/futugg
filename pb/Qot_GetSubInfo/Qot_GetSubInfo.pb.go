// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Qot_GetSubInfo/Qot_GetSubInfo.proto

package Qot_GetSubInfo

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "futugg/pb/Common"
import Qot_Common "futugg/pb/Qot_Common"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type C2S struct {
	IsReqAllConn         *bool    `protobuf:"varint,1,opt,name=isReqAllConn" json:"isReqAllConn,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C2S) Reset()         { *m = C2S{} }
func (m *C2S) String() string { return proto.CompactTextString(m) }
func (*C2S) ProtoMessage()    {}
func (*C2S) Descriptor() ([]byte, []int) {
	return fileDescriptor_Qot_GetSubInfo_508eb49e88e9ec90, []int{0}
}
func (m *C2S) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C2S.Unmarshal(m, b)
}
func (m *C2S) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C2S.Marshal(b, m, deterministic)
}
func (dst *C2S) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2S.Merge(dst, src)
}
func (m *C2S) XXX_Size() int {
	return xxx_messageInfo_C2S.Size(m)
}
func (m *C2S) XXX_DiscardUnknown() {
	xxx_messageInfo_C2S.DiscardUnknown(m)
}

var xxx_messageInfo_C2S proto.InternalMessageInfo

func (m *C2S) GetIsReqAllConn() bool {
	if m != nil && m.IsReqAllConn != nil {
		return *m.IsReqAllConn
	}
	return false
}

type S2C struct {
	ConnSubInfoList      []*Qot_Common.ConnSubInfo `protobuf:"bytes,1,rep,name=connSubInfoList" json:"connSubInfoList,omitempty"`
	TotalUsedQuota       *int32                    `protobuf:"varint,2,req,name=totalUsedQuota" json:"totalUsedQuota,omitempty"`
	RemainQuota          *int32                    `protobuf:"varint,3,req,name=remainQuota" json:"remainQuota,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *S2C) Reset()         { *m = S2C{} }
func (m *S2C) String() string { return proto.CompactTextString(m) }
func (*S2C) ProtoMessage()    {}
func (*S2C) Descriptor() ([]byte, []int) {
	return fileDescriptor_Qot_GetSubInfo_508eb49e88e9ec90, []int{1}
}
func (m *S2C) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S2C.Unmarshal(m, b)
}
func (m *S2C) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S2C.Marshal(b, m, deterministic)
}
func (dst *S2C) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2C.Merge(dst, src)
}
func (m *S2C) XXX_Size() int {
	return xxx_messageInfo_S2C.Size(m)
}
func (m *S2C) XXX_DiscardUnknown() {
	xxx_messageInfo_S2C.DiscardUnknown(m)
}

var xxx_messageInfo_S2C proto.InternalMessageInfo

func (m *S2C) GetConnSubInfoList() []*Qot_Common.ConnSubInfo {
	if m != nil {
		return m.ConnSubInfoList
	}
	return nil
}

func (m *S2C) GetTotalUsedQuota() int32 {
	if m != nil && m.TotalUsedQuota != nil {
		return *m.TotalUsedQuota
	}
	return 0
}

func (m *S2C) GetRemainQuota() int32 {
	if m != nil && m.RemainQuota != nil {
		return *m.RemainQuota
	}
	return 0
}

type Request struct {
	C2S                  *C2S     `protobuf:"bytes,1,req,name=c2s" json:"c2s,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_Qot_GetSubInfo_508eb49e88e9ec90, []int{2}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetC2S() *C2S {
	if m != nil {
		return m.C2S
	}
	return nil
}

type Response struct {
	RetType              *int32   `protobuf:"varint,1,req,name=retType,def=-400" json:"retType,omitempty"`
	RetMsg               *string  `protobuf:"bytes,2,opt,name=retMsg" json:"retMsg,omitempty"`
	ErrCode              *int32   `protobuf:"varint,3,opt,name=errCode" json:"errCode,omitempty"`
	S2C                  *S2C     `protobuf:"bytes,4,opt,name=s2c" json:"s2c,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_Qot_GetSubInfo_508eb49e88e9ec90, []int{3}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

const Default_Response_RetType int32 = -400

func (m *Response) GetRetType() int32 {
	if m != nil && m.RetType != nil {
		return *m.RetType
	}
	return Default_Response_RetType
}

func (m *Response) GetRetMsg() string {
	if m != nil && m.RetMsg != nil {
		return *m.RetMsg
	}
	return ""
}

func (m *Response) GetErrCode() int32 {
	if m != nil && m.ErrCode != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *Response) GetS2C() *S2C {
	if m != nil {
		return m.S2C
	}
	return nil
}

func init() {
	proto.RegisterType((*C2S)(nil), "Qot_GetSubInfo.C2S")
	proto.RegisterType((*S2C)(nil), "Qot_GetSubInfo.S2C")
	proto.RegisterType((*Request)(nil), "Qot_GetSubInfo.Request")
	proto.RegisterType((*Response)(nil), "Qot_GetSubInfo.Response")
}

func init() {
	proto.RegisterFile("Qot_GetSubInfo/Qot_GetSubInfo.proto", fileDescriptor_Qot_GetSubInfo_508eb49e88e9ec90)
}

var fileDescriptor_Qot_GetSubInfo_508eb49e88e9ec90 = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xdf, 0x4b, 0x42, 0x31,
	0x14, 0xc7, 0xd9, 0xbd, 0x9a, 0x76, 0x0c, 0x83, 0x09, 0x35, 0x0c, 0x62, 0xdc, 0x28, 0xd6, 0x43,
	0x2a, 0xa3, 0xa7, 0xde, 0x64, 0x0f, 0x11, 0xd4, 0x83, 0xbb, 0xf5, 0x1c, 0xa6, 0xa7, 0x10, 0x74,
	0x47, 0xb7, 0xf9, 0xd0, 0x73, 0xff, 0x41, 0x7f, 0x71, 0x5c, 0xaf, 0xe2, 0x0f, 0x7a, 0xda, 0xbe,
	0x9f, 0xef, 0x67, 0xdb, 0x61, 0x70, 0x35, 0xa0, 0xf8, 0xfe, 0x88, 0x31, 0x5f, 0x7e, 0x3c, 0xb9,
	0x4f, 0xea, 0xee, 0xc7, 0xce, 0xdc, 0x53, 0x24, 0xde, 0xdc, 0xa7, 0xed, 0x96, 0xa1, 0xd9, 0x8c,
	0x5c, 0xb7, 0x5c, 0x4a, 0xa9, 0x7d, 0x51, 0x48, 0xeb, 0x62, 0xbb, 0x2d, 0xcb, 0xec, 0x16, 0x52,
	0xa3, 0x73, 0x9e, 0xc1, 0xc9, 0x24, 0x58, 0x5c, 0xf4, 0xa7, 0x53, 0x43, 0xce, 0x09, 0x26, 0x99,
	0xaa, 0xdb, 0x3d, 0x96, 0xfd, 0x32, 0x48, 0x73, 0x6d, 0x78, 0x1f, 0x4e, 0x47, 0xe4, 0xdc, 0xfa,
	0xcd, 0xe7, 0x49, 0x88, 0x82, 0xc9, 0x54, 0x35, 0xf4, 0x79, 0x67, 0xe7, 0x7a, 0xb3, 0x55, 0xec,
	0xa1, 0xcf, 0x6f, 0xa0, 0x19, 0x29, 0x0e, 0xa7, 0x6f, 0x01, 0xc7, 0x83, 0x25, 0xc5, 0xa1, 0x48,
	0x64, 0xa2, 0xaa, 0xf6, 0x80, 0x72, 0x09, 0x0d, 0x8f, 0xb3, 0xe1, 0xc4, 0x95, 0x52, 0xba, 0x92,
	0x76, 0x51, 0xd6, 0x83, 0x9a, 0xc5, 0xc5, 0x12, 0x43, 0xe4, 0xd7, 0x90, 0x8e, 0x74, 0x10, 0x4c,
	0x26, 0xaa, 0xa1, 0x5b, 0x9d, 0x83, 0x0f, 0x33, 0x3a, 0xb7, 0x45, 0x9f, 0xfd, 0x30, 0xa8, 0x5b,
	0x0c, 0x73, 0x72, 0x01, 0xf9, 0x25, 0xd4, 0x3c, 0xc6, 0xd7, 0xef, 0x39, 0xae, 0xce, 0x55, 0x1f,
	0x2a, 0x77, 0xf7, 0xbd, 0x9e, 0xdd, 0x40, 0x7e, 0x06, 0x47, 0x1e, 0xe3, 0x4b, 0xf8, 0x12, 0x89,
	0x64, 0xea, 0xd8, 0xae, 0x13, 0x17, 0x50, 0x43, 0xef, 0x0d, 0x8d, 0x51, 0xa4, 0x92, 0xa9, 0xaa,
	0xdd, 0xc4, 0x62, 0x8a, 0xa0, 0x47, 0xa2, 0x22, 0xd9, 0x7f, 0x53, 0xe4, 0xda, 0xd8, 0xa2, 0xff,
	0x0b, 0x00, 0x00, 0xff, 0xff, 0x9a, 0x5e, 0x9e, 0xb2, 0xdf, 0x01, 0x00, 0x00,
}