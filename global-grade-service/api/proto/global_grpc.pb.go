// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GlobalClient is the client API for Global service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GlobalClient interface {
	GetGlobalAverage(ctx context.Context, in *GlobalParams, opts ...grpc.CallOption) (*GlobalAverage, error)
}

type globalClient struct {
	cc grpc.ClientConnInterface
}

func NewGlobalClient(cc grpc.ClientConnInterface) GlobalClient {
	return &globalClient{cc}
}

func (c *globalClient) GetGlobalAverage(ctx context.Context, in *GlobalParams, opts ...grpc.CallOption) (*GlobalAverage, error) {
	out := new(GlobalAverage)
	err := c.cc.Invoke(ctx, "/parent.Global/GetGlobalAverage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GlobalServer is the server API for Global service.
// All implementations must embed UnimplementedGlobalServer
// for forward compatibility
type GlobalServer interface {
	GetGlobalAverage(context.Context, *GlobalParams) (*GlobalAverage, error)
	mustEmbedUnimplementedGlobalServer()
}

// UnimplementedGlobalServer must be embedded to have forward compatible implementations.
type UnimplementedGlobalServer struct {
}

func (UnimplementedGlobalServer) GetGlobalAverage(context.Context, *GlobalParams) (*GlobalAverage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGlobalAverage not implemented")
}
func (UnimplementedGlobalServer) mustEmbedUnimplementedGlobalServer() {}

// UnsafeGlobalServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GlobalServer will
// result in compilation errors.
type UnsafeGlobalServer interface {
	mustEmbedUnimplementedGlobalServer()
}

func RegisterGlobalServer(s grpc.ServiceRegistrar, srv GlobalServer) {
	s.RegisterService(&Global_ServiceDesc, srv)
}

func _Global_GetGlobalAverage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GlobalParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GlobalServer).GetGlobalAverage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/parent.Global/GetGlobalAverage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GlobalServer).GetGlobalAverage(ctx, req.(*GlobalParams))
	}
	return interceptor(ctx, in, info, handler)
}

// Global_ServiceDesc is the grpc.ServiceDesc for Global service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Global_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "parent.Global",
	HandlerType: (*GlobalServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGlobalAverage",
			Handler:    _Global_GetGlobalAverage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/global.proto",
}
