// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package message

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RequestArgs struct {
	Args1                int32    `protobuf:"varint,1,opt,name=args1,proto3" json:"args1,omitempty"`
	Args2                int32    `protobuf:"varint,2,opt,name=args2,proto3" json:"args2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestArgs) Reset()         { *m = RequestArgs{} }
func (m *RequestArgs) String() string { return proto.CompactTextString(m) }
func (*RequestArgs) ProtoMessage()    {}
func (*RequestArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_8fda6feb58de70bd, []int{0}
}
func (m *RequestArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestArgs.Unmarshal(m, b)
}
func (m *RequestArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestArgs.Marshal(b, m, deterministic)
}
func (dst *RequestArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestArgs.Merge(dst, src)
}
func (m *RequestArgs) XXX_Size() int {
	return xxx_messageInfo_RequestArgs.Size(m)
}
func (m *RequestArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestArgs.DiscardUnknown(m)
}

var xxx_messageInfo_RequestArgs proto.InternalMessageInfo

func (m *RequestArgs) GetArgs1() int32 {
	if m != nil {
		return m.Args1
	}
	return 0
}

func (m *RequestArgs) GetArgs2() int32 {
	if m != nil {
		return m.Args2
	}
	return 0
}

type Response struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_8fda6feb58de70bd, []int{1}
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

func (m *Response) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Response) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*RequestArgs)(nil), "message.RequestArgs")
	proto.RegisterType((*Response)(nil), "message.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MathServiceClient is the client API for MathService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MathServiceClient interface {
	// 服务
	AddMethod(ctx context.Context, in *RequestArgs, opts ...grpc.CallOption) (*Response, error)
}

type mathServiceClient struct {
	cc *grpc.ClientConn
}

func NewMathServiceClient(cc *grpc.ClientConn) MathServiceClient {
	return &mathServiceClient{cc}
}

func (c *mathServiceClient) AddMethod(ctx context.Context, in *RequestArgs, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/message.MathService/AddMethod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MathServiceServer is the server API for MathService service.
type MathServiceServer interface {
	// 服务
	AddMethod(context.Context, *RequestArgs) (*Response, error)
}

func RegisterMathServiceServer(s *grpc.Server, srv MathServiceServer) {
	s.RegisterService(&_MathService_serviceDesc, srv)
}

func _MathService_AddMethod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MathServiceServer).AddMethod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.MathService/AddMethod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MathServiceServer).AddMethod(ctx, req.(*RequestArgs))
	}
	return interceptor(ctx, in, info, handler)
}

var _MathService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "message.MathService",
	HandlerType: (*MathServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddMethod",
			Handler:    _MathService_AddMethod_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message.proto",
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_message_8fda6feb58de70bd) }

var fileDescriptor_message_8fda6feb58de70bd = []byte{
	// 167 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0x2c, 0xb9,
	0xb8, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x1c, 0x8b, 0xd2, 0x8b, 0x85, 0x44, 0xb8, 0x58,
	0x13, 0x8b, 0xd2, 0x8b, 0x0d, 0x25, 0x18, 0x15, 0x18, 0x35, 0x58, 0x83, 0x20, 0x1c, 0x98, 0xa8,
	0x91, 0x04, 0x13, 0x42, 0xd4, 0x48, 0xc9, 0x82, 0x8b, 0x23, 0x28, 0xb5, 0xb8, 0x20, 0x3f, 0xaf,
	0x38, 0x55, 0x48, 0x88, 0x8b, 0x25, 0x39, 0x3f, 0x25, 0x15, 0xaa, 0x0d, 0xcc, 0x16, 0x92, 0xe0,
	0x82, 0xd9, 0x02, 0xd6, 0xc7, 0x19, 0x04, 0xe3, 0x1a, 0xb9, 0x72, 0x71, 0xfb, 0x26, 0x96, 0x64,
	0x04, 0xa7, 0x16, 0x95, 0x65, 0x26, 0xa7, 0x0a, 0x99, 0x71, 0x71, 0x3a, 0xa6, 0xa4, 0xf8, 0xa6,
	0x96, 0x64, 0xe4, 0xa7, 0x08, 0x89, 0xe8, 0xc1, 0x5c, 0x8a, 0xe4, 0x2e, 0x29, 0x41, 0x24, 0x51,
	0x88, 0x95, 0x4a, 0x0c, 0x49, 0x6c, 0x60, 0xbf, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xdf,
	0x5b, 0x93, 0xde, 0xdc, 0x00, 0x00, 0x00,
}