// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc.proto

/*
Package grpc is a generated protocol buffer package.

It is generated from these files:
	grpc.proto

It has these top-level messages:
	VersionRequest
	VersionResponse
*/
package grpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc1 "google.golang.org/grpc"
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

type VersionRequest struct {
}

func (m *VersionRequest) Reset()                    { *m = VersionRequest{} }
func (m *VersionRequest) String() string            { return proto.CompactTextString(m) }
func (*VersionRequest) ProtoMessage()               {}
func (*VersionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type VersionResponse struct {
	Name      string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Version   string `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
	Build     string `protobuf:"bytes,3,opt,name=build" json:"build,omitempty"`
	BuildDate string `protobuf:"bytes,4,opt,name=build_date" json:"build_date,omitempty"`
	Uptime    uint64 `protobuf:"varint,5,opt,name=uptime" json:"uptime,omitempty"`
}

func (m *VersionResponse) Reset()                    { *m = VersionResponse{} }
func (m *VersionResponse) String() string            { return proto.CompactTextString(m) }
func (*VersionResponse) ProtoMessage()               {}
func (*VersionResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *VersionResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *VersionResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *VersionResponse) GetBuild() string {
	if m != nil {
		return m.Build
	}
	return ""
}

func (m *VersionResponse) GetBuildDate() string {
	if m != nil {
		return m.BuildDate
	}
	return ""
}

func (m *VersionResponse) GetUptime() uint64 {
	if m != nil {
		return m.Uptime
	}
	return 0
}

func init() {
	proto.RegisterType((*VersionRequest)(nil), "grpc.VersionRequest")
	proto.RegisterType((*VersionResponse)(nil), "grpc.VersionResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc1.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc1.SupportPackageIsVersion4

// Client API for Grpc service

type GrpcClient interface {
	Version(ctx context.Context, in *VersionRequest, opts ...grpc1.CallOption) (*VersionResponse, error)
}

type grpcClient struct {
	cc *grpc1.ClientConn
}

func NewGrpcClient(cc *grpc1.ClientConn) GrpcClient {
	return &grpcClient{cc}
}

func (c *grpcClient) Version(ctx context.Context, in *VersionRequest, opts ...grpc1.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := grpc1.Invoke(ctx, "/grpc.Grpc/Version", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Grpc service

type GrpcServer interface {
	Version(context.Context, *VersionRequest) (*VersionResponse, error)
}

func RegisterGrpcServer(s *grpc1.Server, srv GrpcServer) {
	s.RegisterService(&_Grpc_serviceDesc, srv)
}

func _Grpc_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcServer).Version(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Grpc/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcServer).Version(ctx, req.(*VersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Grpc_serviceDesc = grpc1.ServiceDesc{
	ServiceName: "grpc.Grpc",
	HandlerType: (*GrpcServer)(nil),
	Methods: []grpc1.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _Grpc_Version_Handler,
		},
	},
	Streams:  []grpc1.StreamDesc{},
	Metadata: "grpc.proto",
}

func init() { proto.RegisterFile("grpc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 170 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x2f, 0x2a, 0x48,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0x04, 0xb8, 0xf8, 0xc2, 0x52,
	0x8b, 0x8a, 0x33, 0xf3, 0xf3, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x94, 0x92, 0xb9, 0xf8,
	0xe1, 0x22, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x3c, 0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9,
	0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x42, 0xfc, 0x5c, 0xec, 0x65, 0x10, 0x05, 0x12, 0x4c, 0x60,
	0x01, 0x5e, 0x2e, 0xd6, 0xa4, 0xd2, 0xcc, 0x9c, 0x14, 0x09, 0x66, 0x30, 0x57, 0x88, 0x8b, 0x0b,
	0xcc, 0x8d, 0x4f, 0x49, 0x2c, 0x49, 0x95, 0x60, 0x01, 0x8b, 0xf1, 0x71, 0xb1, 0x95, 0x16, 0x94,
	0x64, 0xe6, 0xa6, 0x4a, 0xb0, 0x2a, 0x30, 0x6a, 0xb0, 0x18, 0x39, 0x70, 0xb1, 0xb8, 0x17, 0x15,
	0x24, 0x0b, 0x59, 0x70, 0xb1, 0x43, 0x2d, 0x13, 0x12, 0xd1, 0x03, 0x3b, 0x0e, 0xd5, 0x35, 0x52,
	0xa2, 0x68, 0xa2, 0x10, 0x17, 0x29, 0x31, 0x24, 0xb1, 0x81, 0x7d, 0x61, 0x0c, 0x08, 0x00, 0x00,
	0xff, 0xff, 0x82, 0x87, 0x66, 0x0a, 0xd3, 0x00, 0x00, 0x00,
}