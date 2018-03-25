// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	grpc.proto

It has these top-level messages:
	VersionRequest
	VersionResponse
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
import google_protobuf1 "github.com/golang/protobuf/ptypes/duration"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type VersionRequest struct {
}

func (m *VersionRequest) Reset()                    { *m = VersionRequest{} }
func (m *VersionRequest) String() string            { return proto1.CompactTextString(m) }
func (*VersionRequest) ProtoMessage()               {}
func (*VersionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type VersionResponse struct {
	Name          string                     `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Version       string                     `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
	Build         string                     `protobuf:"bytes,3,opt,name=build" json:"build,omitempty"`
	BuildDatetime *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=build_datetime,json=buildDatetime" json:"build_datetime,omitempty"`
	StartDatetime *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=start_datetime,json=startDatetime" json:"start_datetime,omitempty"`
	Uptime        *google_protobuf1.Duration `protobuf:"bytes,6,opt,name=uptime" json:"uptime,omitempty"`
}

func (m *VersionResponse) Reset()                    { *m = VersionResponse{} }
func (m *VersionResponse) String() string            { return proto1.CompactTextString(m) }
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

func (m *VersionResponse) GetBuildDatetime() *google_protobuf.Timestamp {
	if m != nil {
		return m.BuildDatetime
	}
	return nil
}

func (m *VersionResponse) GetStartDatetime() *google_protobuf.Timestamp {
	if m != nil {
		return m.StartDatetime
	}
	return nil
}

func (m *VersionResponse) GetUptime() *google_protobuf1.Duration {
	if m != nil {
		return m.Uptime
	}
	return nil
}

func init() {
	proto1.RegisterType((*VersionRequest)(nil), "shadow.grpc.VersionRequest")
	proto1.RegisterType((*VersionResponse)(nil), "shadow.grpc.VersionResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Grpc service

type GrpcClient interface {
	Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error)
}

type grpcClient struct {
	cc *grpc.ClientConn
}

func NewGrpcClient(cc *grpc.ClientConn) GrpcClient {
	return &grpcClient{cc}
}

func (c *grpcClient) Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := grpc.Invoke(ctx, "/shadow.grpc.Grpc/Version", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Grpc service

type GrpcServer interface {
	Version(context.Context, *VersionRequest) (*VersionResponse, error)
}

func RegisterGrpcServer(s *grpc.Server, srv GrpcServer) {
	s.RegisterService(&_Grpc_serviceDesc, srv)
}

func _Grpc_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shadow.grpc.Grpc/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcServer).Version(ctx, req.(*VersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Grpc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "shadow.grpc.Grpc",
	HandlerType: (*GrpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _Grpc_Version_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc.proto",
}

func init() { proto1.RegisterFile("grpc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xbb, 0x4e, 0xc3, 0x30,
	0x18, 0x85, 0x69, 0xe9, 0x45, 0xb8, 0xa2, 0x20, 0x8b, 0x21, 0x04, 0x04, 0x55, 0xa7, 0x4e, 0x36,
	0x94, 0x27, 0x00, 0x55, 0xb0, 0x31, 0x44, 0x88, 0x81, 0x05, 0x39, 0x89, 0x49, 0x2c, 0x6a, 0xff,
	0xc6, 0x17, 0x78, 0x0b, 0x9e, 0x19, 0xc5, 0x4e, 0xca, 0x55, 0x62, 0x8a, 0x7d, 0xfe, 0xf3, 0x1d,
	0xfd, 0x39, 0x46, 0xa8, 0x32, 0xba, 0x20, 0xda, 0x80, 0x03, 0x3c, 0xb1, 0x35, 0x2b, 0xe1, 0x8d,
	0x34, 0x52, 0x7a, 0x5a, 0x01, 0x54, 0x6b, 0x4e, 0xc3, 0x28, 0xf7, 0x4f, 0xd4, 0x09, 0xc9, 0xad,
	0x63, 0x52, 0x47, 0x77, 0x7a, 0xf2, 0xd3, 0x50, 0x7a, 0xc3, 0x9c, 0x00, 0x15, 0xe7, 0xf3, 0x7d,
	0x34, 0xbd, 0xe7, 0xc6, 0x0a, 0x50, 0x19, 0x7f, 0xf1, 0xdc, 0xba, 0xf9, 0x7b, 0x1f, 0xed, 0x6d,
	0x24, 0xab, 0x41, 0x59, 0x8e, 0x31, 0x1a, 0x28, 0x26, 0x79, 0xd2, 0x9b, 0xf5, 0x16, 0x3b, 0x59,
	0x38, 0xe3, 0x04, 0x8d, 0x5f, 0xa3, 0x2d, 0xe9, 0x07, 0xb9, 0xbb, 0xe2, 0x03, 0x34, 0xcc, 0xbd,
	0x58, 0x97, 0xc9, 0x76, 0xd0, 0xe3, 0x05, 0x5f, 0xa2, 0x69, 0x38, 0x3c, 0x96, 0xcc, 0xf1, 0x66,
	0xcd, 0x64, 0x30, 0xeb, 0x2d, 0x26, 0xcb, 0x94, 0xc4, 0x15, 0x49, 0xb7, 0x22, 0xb9, 0xeb, 0xfe,
	0x21, 0xdb, 0x0d, 0xc4, 0xaa, 0x05, 0x9a, 0x08, 0xeb, 0x98, 0x71, 0x9f, 0x11, 0xc3, 0xff, 0x23,
	0x02, 0xb1, 0x89, 0x38, 0x47, 0x23, 0xaf, 0x03, 0x3a, 0x0a, 0xe8, 0xe1, 0x2f, 0x74, 0xd5, 0x16,
	0x94, 0xb5, 0xc6, 0xe5, 0x2d, 0x1a, 0xdc, 0x18, 0x5d, 0xe0, 0x6b, 0x34, 0x6e, 0x7b, 0xc1, 0x47,
	0xe4, 0xcb, 0x23, 0x90, 0xef, 0x05, 0xa6, 0xc7, 0x7f, 0x0f, 0x63, 0x95, 0xf3, 0xad, 0xab, 0xb3,
	0x07, 0x52, 0x09, 0x57, 0xfb, 0x9c, 0x14, 0x20, 0xe9, 0xb3, 0xa8, 0x99, 0x04, 0x1a, 0x11, 0x5a,
	0x80, 0xd4, 0xa0, 0xb8, 0x72, 0x96, 0x36, 0x74, 0xfb, 0x6c, 0xa3, 0xf0, 0xb9, 0xf8, 0x08, 0x00,
	0x00, 0xff, 0xff, 0xf4, 0x2e, 0x32, 0xd6, 0x07, 0x02, 0x00, 0x00,
}