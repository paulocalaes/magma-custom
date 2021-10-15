// Code generated by protoc-gen-go. DO NOT EDIT.
// source: provider.proto

package protos

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protos "magma/orc8r/lib/go/protos"
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

func init() { proto.RegisterFile("provider.proto", fileDescriptor_c6a9f3c02af3d1c8) }

var fileDescriptor_c6a9f3c02af3d1c8 = []byte{
	// 140 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x28, 0xca, 0x2f,
	0xcb, 0x4c, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xc9, 0x4d, 0x4c, 0xcf,
	0x4d, 0xd4, 0xcb, 0x2f, 0x4a, 0xb6, 0x28, 0xd2, 0x2b, 0x2e, 0x29, 0x4a, 0x4d, 0xcc, 0x4d, 0x2d,
	0x92, 0x92, 0x06, 0xf3, 0xf5, 0xc1, 0x4a, 0x8a, 0xf5, 0x61, 0xc2, 0x10, 0x2d, 0x46, 0x51, 0x5c,
	0x7c, 0xc1, 0x60, 0x91, 0x00, 0xa8, 0x51, 0x42, 0x1e, 0x5c, 0x5c, 0xee, 0xa9, 0x25, 0xa1, 0x05,
	0x29, 0x89, 0x25, 0xa9, 0xc5, 0x42, 0x52, 0x7a, 0xc8, 0x66, 0x42, 0x94, 0x06, 0xa5, 0x16, 0x96,
	0xa6, 0x16, 0x97, 0x48, 0xc9, 0xa0, 0xc8, 0xb9, 0x24, 0x96, 0x24, 0x42, 0x74, 0x39, 0x25, 0x96,
	0x24, 0x67, 0x28, 0x31, 0x38, 0x71, 0x44, 0xb1, 0x41, 0x2c, 0x4d, 0x82, 0xd0, 0xc6, 0x80, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xb6, 0x56, 0x6b, 0x59, 0xb1, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// StreamProviderClient is the client API for StreamProvider service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StreamProviderClient interface {
	GetUpdates(ctx context.Context, in *protos.StreamRequest, opts ...grpc.CallOption) (*protos.DataUpdateBatch, error)
}

type streamProviderClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamProviderClient(cc grpc.ClientConnInterface) StreamProviderClient {
	return &streamProviderClient{cc}
}

func (c *streamProviderClient) GetUpdates(ctx context.Context, in *protos.StreamRequest, opts ...grpc.CallOption) (*protos.DataUpdateBatch, error) {
	out := new(protos.DataUpdateBatch)
	err := c.cc.Invoke(ctx, "/magma.orc8r.streamer.StreamProvider/GetUpdates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StreamProviderServer is the server API for StreamProvider service.
type StreamProviderServer interface {
	GetUpdates(context.Context, *protos.StreamRequest) (*protos.DataUpdateBatch, error)
}

// UnimplementedStreamProviderServer can be embedded to have forward compatible implementations.
type UnimplementedStreamProviderServer struct {
}

func (*UnimplementedStreamProviderServer) GetUpdates(ctx context.Context, req *protos.StreamRequest) (*protos.DataUpdateBatch, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUpdates not implemented")
}

func RegisterStreamProviderServer(s *grpc.Server, srv StreamProviderServer) {
	s.RegisterService(&_StreamProvider_serviceDesc, srv)
}

func _StreamProvider_GetUpdates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protos.StreamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamProviderServer).GetUpdates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.streamer.StreamProvider/GetUpdates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamProviderServer).GetUpdates(ctx, req.(*protos.StreamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StreamProvider_serviceDesc = grpc.ServiceDesc{
	ServiceName: "magma.orc8r.streamer.StreamProvider",
	HandlerType: (*StreamProviderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUpdates",
			Handler:    _StreamProvider_GetUpdates_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "provider.proto",
}
