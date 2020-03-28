// Code generated by protoc-gen-go. DO NOT EDIT.
// source: addr.proto

package addr

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type AddrReply struct {
	Addr                 string   `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddrReply) Reset()         { *m = AddrReply{} }
func (m *AddrReply) String() string { return proto.CompactTextString(m) }
func (*AddrReply) ProtoMessage()    {}
func (*AddrReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_eaa67da8ec6ef8e1, []int{0}
}

func (m *AddrReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddrReply.Unmarshal(m, b)
}
func (m *AddrReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddrReply.Marshal(b, m, deterministic)
}
func (m *AddrReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddrReply.Merge(m, src)
}
func (m *AddrReply) XXX_Size() int {
	return xxx_messageInfo_AddrReply.Size(m)
}
func (m *AddrReply) XXX_DiscardUnknown() {
	xxx_messageInfo_AddrReply.DiscardUnknown(m)
}

var xxx_messageInfo_AddrReply proto.InternalMessageInfo

func (m *AddrReply) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func init() {
	proto.RegisterType((*AddrReply)(nil), "addr.AddrReply")
}

func init() {
	proto.RegisterFile("addr.proto", fileDescriptor_eaa67da8ec6ef8e1)
}

var fileDescriptor_eaa67da8ec6ef8e1 = []byte{
	// 129 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x4c, 0x49, 0x29,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0xa5, 0xa4, 0xd3, 0xf3, 0xf3, 0xd3,
	0x73, 0x52, 0xf5, 0xc1, 0x62, 0x49, 0xa5, 0x69, 0xfa, 0xa9, 0xb9, 0x05, 0x25, 0x95, 0x10, 0x25,
	0x4a, 0xf2, 0x5c, 0x9c, 0x8e, 0x29, 0x29, 0x45, 0x41, 0xa9, 0x05, 0x39, 0x95, 0x42, 0x42, 0x5c,
	0x60, 0x1d, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x91, 0x0d, 0x17, 0x0b, 0x48,
	0x81, 0x90, 0x09, 0x17, 0xbb, 0x7b, 0x6a, 0x09, 0x98, 0x29, 0xa6, 0x07, 0x31, 0x51, 0x0f, 0x66,
	0xa2, 0x9e, 0x2b, 0xc8, 0x44, 0x29, 0x7e, 0x3d, 0xb0, 0xdd, 0x70, 0xf3, 0x94, 0x18, 0x92, 0xd8,
	0xc0, 0x4a, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xcd, 0x39, 0xae, 0x4e, 0x96, 0x00, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AddrClient is the client API for Addr service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AddrClient interface {
	GetAddr(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*AddrReply, error)
}

type addrClient struct {
	cc grpc.ClientConnInterface
}

func NewAddrClient(cc grpc.ClientConnInterface) AddrClient {
	return &addrClient{cc}
}

func (c *addrClient) GetAddr(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*AddrReply, error) {
	out := new(AddrReply)
	err := c.cc.Invoke(ctx, "/addr.Addr/GetAddr", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AddrServer is the server API for Addr service.
type AddrServer interface {
	GetAddr(context.Context, *empty.Empty) (*AddrReply, error)
}

// UnimplementedAddrServer can be embedded to have forward compatible implementations.
type UnimplementedAddrServer struct {
}

func (*UnimplementedAddrServer) GetAddr(ctx context.Context, req *empty.Empty) (*AddrReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAddr not implemented")
}

func RegisterAddrServer(s *grpc.Server, srv AddrServer) {
	s.RegisterService(&_Addr_serviceDesc, srv)
}

func _Addr_GetAddr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddrServer).GetAddr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/addr.Addr/GetAddr",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddrServer).GetAddr(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Addr_serviceDesc = grpc.ServiceDesc{
	ServiceName: "addr.Addr",
	HandlerType: (*AddrServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAddr",
			Handler:    _Addr_GetAddr_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "addr.proto",
}
