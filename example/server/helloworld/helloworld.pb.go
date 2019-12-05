// Code generated by protoc-gen-go. DO NOT EDIT.
// source: helloworld.proto

package helloworld

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Color int32

const (
	Color_RED   Color = 0
	Color_BLUE  Color = 1
	Color_GREEN Color = 2
)

var Color_name = map[int32]string{
	0: "RED",
	1: "BLUE",
	2: "GREEN",
}

var Color_value = map[string]int32{
	"RED":   0,
	"BLUE":  1,
	"GREEN": 2,
}

func (x Color) String() string {
	return proto.EnumName(Color_name, int32(x))
}

func (Color) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_17b8c58d586b62f2, []int{0}
}

type HelloRequest struct {
	Displayname          string        `protobuf:"bytes,1,opt,name=displayname,proto3" json:"displayname,omitempty"`
	Ex                   []*ComplexMsg `protobuf:"bytes,2,rep,name=ex,proto3" json:"ex,omitempty"`
	Jobs                 []string      `protobuf:"bytes,3,rep,name=jobs,proto3" json:"jobs,omitempty"`
	Color                Color         `protobuf:"varint,4,opt,name=color,proto3,enum=helloworld.Color" json:"color,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_17b8c58d586b62f2, []int{0}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetDisplayname() string {
	if m != nil {
		return m.Displayname
	}
	return ""
}

func (m *HelloRequest) GetEx() []*ComplexMsg {
	if m != nil {
		return m.Ex
	}
	return nil
}

func (m *HelloRequest) GetJobs() []string {
	if m != nil {
		return m.Jobs
	}
	return nil
}

func (m *HelloRequest) GetColor() Color {
	if m != nil {
		return m.Color
	}
	return Color_RED
}

type ComplexMsg struct {
	Displayname          string               `protobuf:"bytes,1,opt,name=displayname,proto3" json:"displayname,omitempty"`
	Foo                  *YetAnotherNestedMsg `protobuf:"bytes,2,opt,name=foo,proto3" json:"foo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ComplexMsg) Reset()         { *m = ComplexMsg{} }
func (m *ComplexMsg) String() string { return proto.CompactTextString(m) }
func (*ComplexMsg) ProtoMessage()    {}
func (*ComplexMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_17b8c58d586b62f2, []int{1}
}

func (m *ComplexMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComplexMsg.Unmarshal(m, b)
}
func (m *ComplexMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComplexMsg.Marshal(b, m, deterministic)
}
func (m *ComplexMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComplexMsg.Merge(m, src)
}
func (m *ComplexMsg) XXX_Size() int {
	return xxx_messageInfo_ComplexMsg.Size(m)
}
func (m *ComplexMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_ComplexMsg.DiscardUnknown(m)
}

var xxx_messageInfo_ComplexMsg proto.InternalMessageInfo

func (m *ComplexMsg) GetDisplayname() string {
	if m != nil {
		return m.Displayname
	}
	return ""
}

func (m *ComplexMsg) GetFoo() *YetAnotherNestedMsg {
	if m != nil {
		return m.Foo
	}
	return nil
}

type YetAnotherNestedMsg struct {
	Grades               []int32  `protobuf:"varint,1,rep,packed,name=grades,proto3" json:"grades,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *YetAnotherNestedMsg) Reset()         { *m = YetAnotherNestedMsg{} }
func (m *YetAnotherNestedMsg) String() string { return proto.CompactTextString(m) }
func (*YetAnotherNestedMsg) ProtoMessage()    {}
func (*YetAnotherNestedMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_17b8c58d586b62f2, []int{2}
}

func (m *YetAnotherNestedMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_YetAnotherNestedMsg.Unmarshal(m, b)
}
func (m *YetAnotherNestedMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_YetAnotherNestedMsg.Marshal(b, m, deterministic)
}
func (m *YetAnotherNestedMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_YetAnotherNestedMsg.Merge(m, src)
}
func (m *YetAnotherNestedMsg) XXX_Size() int {
	return xxx_messageInfo_YetAnotherNestedMsg.Size(m)
}
func (m *YetAnotherNestedMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_YetAnotherNestedMsg.DiscardUnknown(m)
}

var xxx_messageInfo_YetAnotherNestedMsg proto.InternalMessageInfo

func (m *YetAnotherNestedMsg) GetGrades() []int32 {
	if m != nil {
		return m.Grades
	}
	return nil
}

type HelloReply struct {
	Message              string               `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	ReplyAt              *timestamp.Timestamp `protobuf:"bytes,2,opt,name=reply_at,json=replyAt,proto3" json:"reply_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *HelloReply) Reset()         { *m = HelloReply{} }
func (m *HelloReply) String() string { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()    {}
func (*HelloReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_17b8c58d586b62f2, []int{3}
}

func (m *HelloReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloReply.Unmarshal(m, b)
}
func (m *HelloReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloReply.Marshal(b, m, deterministic)
}
func (m *HelloReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloReply.Merge(m, src)
}
func (m *HelloReply) XXX_Size() int {
	return xxx_messageInfo_HelloReply.Size(m)
}
func (m *HelloReply) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloReply.DiscardUnknown(m)
}

var xxx_messageInfo_HelloReply proto.InternalMessageInfo

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *HelloReply) GetReplyAt() *timestamp.Timestamp {
	if m != nil {
		return m.ReplyAt
	}
	return nil
}

func init() {
	proto.RegisterEnum("helloworld.Color", Color_name, Color_value)
	proto.RegisterType((*HelloRequest)(nil), "helloworld.HelloRequest")
	proto.RegisterType((*ComplexMsg)(nil), "helloworld.ComplexMsg")
	proto.RegisterType((*YetAnotherNestedMsg)(nil), "helloworld.YetAnotherNestedMsg")
	proto.RegisterType((*HelloReply)(nil), "helloworld.HelloReply")
}

func init() { proto.RegisterFile("helloworld.proto", fileDescriptor_17b8c58d586b62f2) }

var fileDescriptor_17b8c58d586b62f2 = []byte{
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xd1, 0xab, 0xd3, 0x30,
	0x14, 0xc6, 0xd7, 0x76, 0x5d, 0xb7, 0x33, 0x91, 0x1a, 0x61, 0x84, 0xbd, 0xac, 0x14, 0xd4, 0x22,
	0xd8, 0x61, 0xc5, 0x57, 0x61, 0x6a, 0x99, 0x82, 0xee, 0x21, 0xea, 0x83, 0x0f, 0x22, 0x99, 0x3d,
	0xeb, 0x26, 0xe9, 0xd2, 0x9b, 0x64, 0xdc, 0xf5, 0x0f, 0xb9, 0xff, 0xef, 0xa5, 0xdd, 0xca, 0x7a,
	0x61, 0x0f, 0xf7, 0x2d, 0xe7, 0x7c, 0x5f, 0x4e, 0x7e, 0x39, 0x1f, 0xf8, 0x5b, 0x14, 0x42, 0xde,
	0x4a, 0x25, 0xb2, 0xb8, 0x54, 0xd2, 0x48, 0x02, 0x97, 0xce, 0x74, 0x96, 0x4b, 0x99, 0x0b, 0x9c,
	0x37, 0xca, 0xfa, 0xb0, 0x99, 0x9b, 0x5d, 0x81, 0xda, 0xf0, 0xa2, 0x3c, 0x99, 0xc3, 0x3b, 0x0b,
	0x9e, 0x7c, 0xa9, 0xfd, 0x0c, 0x6f, 0x0e, 0xa8, 0x0d, 0x09, 0x60, 0x9c, 0xed, 0x74, 0x29, 0x78,
	0xb5, 0xe7, 0x05, 0x52, 0x2b, 0xb0, 0xa2, 0x11, 0xeb, 0xb6, 0xc8, 0x4b, 0xb0, 0xf1, 0x48, 0xed,
	0xc0, 0x89, 0xc6, 0xc9, 0x24, 0xee, 0x3c, 0xff, 0x49, 0x16, 0xa5, 0xc0, 0xe3, 0x77, 0x9d, 0x33,
	0x1b, 0x8f, 0x84, 0x40, 0xff, 0xbf, 0x5c, 0x6b, 0xea, 0x04, 0x4e, 0x34, 0x62, 0xcd, 0x99, 0xbc,
	0x02, 0xf7, 0x9f, 0x14, 0x52, 0xd1, 0x7e, 0x60, 0x45, 0x4f, 0x93, 0x67, 0x0f, 0xaf, 0x0b, 0xa9,
	0xd8, 0x49, 0x0f, 0x39, 0xc0, 0x65, 0xdc, 0x23, 0xa0, 0xde, 0x82, 0xb3, 0x91, 0x92, 0xda, 0x81,
	0x15, 0x8d, 0x93, 0x59, 0x77, 0xec, 0x6f, 0x34, 0x8b, 0xbd, 0x34, 0x5b, 0x54, 0x2b, 0xd4, 0x06,
	0xb3, 0x1a, 0xaf, 0xf6, 0x86, 0x6f, 0xe0, 0xf9, 0x15, 0x8d, 0x4c, 0x60, 0x90, 0x2b, 0x9e, 0xa1,
	0xa6, 0x56, 0xe0, 0x44, 0x2e, 0x3b, 0x57, 0xe1, 0x1f, 0x80, 0xf3, 0xa2, 0x4a, 0x51, 0x11, 0x0a,
	0x5e, 0x81, 0x5a, 0xf3, 0xbc, 0xa5, 0x69, 0x4b, 0xf2, 0x1e, 0x86, 0xaa, 0xb6, 0xfc, 0xe5, 0xe6,
	0x8c, 0x33, 0x8d, 0x4f, 0x29, 0xc4, 0x6d, 0x0a, 0xf1, 0xcf, 0x36, 0x05, 0xe6, 0x35, 0xde, 0x85,
	0x79, 0xfd, 0x02, 0xdc, 0x66, 0x01, 0xc4, 0x03, 0x87, 0xa5, 0x9f, 0xfd, 0x1e, 0x19, 0x42, 0xff,
	0xe3, 0xb7, 0x5f, 0xa9, 0x6f, 0x91, 0x11, 0xb8, 0x4b, 0x96, 0xa6, 0x2b, 0xdf, 0x4e, 0xbe, 0x82,
	0xb7, 0x54, 0x88, 0x06, 0x15, 0xf9, 0x00, 0xc3, 0x1f, 0xbc, 0x6a, 0x98, 0x08, 0xed, 0xfe, 0xb8,
	0x9b, 0xe7, 0x74, 0x72, 0x45, 0x29, 0x45, 0x15, 0xf6, 0xd6, 0x83, 0x06, 0xe7, 0xdd, 0x7d, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x7e, 0x48, 0xdc, 0x05, 0x42, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreeterClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/helloworld.Greeter/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServer is the server API for Greeter service.
type GreeterServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
}

// UnimplementedGreeterServer can be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (*UnimplementedGreeterServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "helloworld.proto",
}
