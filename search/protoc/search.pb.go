// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protoc/search.proto

package search

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type Request struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_80e94d5afc8c5090, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Response struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Age                  int64    `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	School               string   `protobuf:"bytes,4,opt,name=school,proto3" json:"school,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_80e94d5afc8c5090, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Response) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Response) GetAge() int64 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *Response) GetSchool() string {
	if m != nil {
		return m.School
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "search.Request")
	proto.RegisterType((*Response)(nil), "search.Response")
}

func init() { proto.RegisterFile("protoc/search.proto", fileDescriptor_80e94d5afc8c5090) }

var fileDescriptor_80e94d5afc8c5090 = []byte{
	// 176 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0x4f, 0xd6, 0x2f, 0x4e, 0x4d, 0x2c, 0x4a, 0xce, 0xd0, 0x03, 0xf3, 0x84, 0xd8, 0x20, 0x3c,
	0x25, 0x49, 0x2e, 0xf6, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x3e, 0x2e, 0xa6, 0xcc,
	0x14, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0xa6, 0xcc, 0x14, 0xa5, 0x08, 0x2e, 0x8e, 0xa0,
	0xd4, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x24, 0x39, 0x66, 0x90, 0x9c, 0x90, 0x10, 0x17, 0x4b,
	0x5e, 0x62, 0x6e, 0xaa, 0x04, 0x13, 0x58, 0x35, 0x98, 0x2d, 0x24, 0xc0, 0xc5, 0x9c, 0x98, 0x9e,
	0x2a, 0xc1, 0x0c, 0x56, 0x04, 0x62, 0x0a, 0x89, 0x71, 0xb1, 0x15, 0x27, 0x67, 0xe4, 0xe7, 0xe7,
	0x48, 0xb0, 0x80, 0xd5, 0x41, 0x79, 0x46, 0x56, 0x5c, 0xec, 0xee, 0x45, 0xa9, 0xa9, 0x25, 0xa9,
	0x45, 0x42, 0xfa, 0x5c, 0x1c, 0x6e, 0x99, 0x79, 0x29, 0x4e, 0x95, 0x9e, 0x29, 0x42, 0xfc, 0x7a,
	0x50, 0x27, 0x42, 0x5d, 0x24, 0x25, 0x80, 0x10, 0x80, 0xb8, 0x43, 0x89, 0x21, 0x89, 0x0d, 0xec,
	0x7e, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6b, 0xc6, 0x28, 0xa5, 0xd6, 0x00, 0x00, 0x00,
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
	FindById(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) FindById(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/search.Greeter/FindById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServer is the server API for Greeter service.
type GreeterServer interface {
	FindById(context.Context, *Request) (*Response, error)
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_FindById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).FindById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/search.Greeter/FindById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).FindById(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "search.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindById",
			Handler:    _Greeter_FindById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protoc/search.proto",
}
