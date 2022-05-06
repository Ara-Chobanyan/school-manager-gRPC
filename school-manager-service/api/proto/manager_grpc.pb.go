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

// SchoolManagerClient is the client API for SchoolManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SchoolManagerClient interface {
	Behavior(ctx context.Context, in *Helper, opts ...grpc.CallOption) (*StudentList, error)
	Exceptional(ctx context.Context, in *Helper, opts ...grpc.CallOption) (*StudentList, error)
}

type schoolManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewSchoolManagerClient(cc grpc.ClientConnInterface) SchoolManagerClient {
	return &schoolManagerClient{cc}
}

func (c *schoolManagerClient) Behavior(ctx context.Context, in *Helper, opts ...grpc.CallOption) (*StudentList, error) {
	out := new(StudentList)
	err := c.cc.Invoke(ctx, "/manager.SchoolManager/Behavior", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schoolManagerClient) Exceptional(ctx context.Context, in *Helper, opts ...grpc.CallOption) (*StudentList, error) {
	out := new(StudentList)
	err := c.cc.Invoke(ctx, "/manager.SchoolManager/Exceptional", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SchoolManagerServer is the server API for SchoolManager service.
// All implementations must embed UnimplementedSchoolManagerServer
// for forward compatibility
type SchoolManagerServer interface {
	Behavior(context.Context, *Helper) (*StudentList, error)
	Exceptional(context.Context, *Helper) (*StudentList, error)
	mustEmbedUnimplementedSchoolManagerServer()
}

// UnimplementedSchoolManagerServer must be embedded to have forward compatible implementations.
type UnimplementedSchoolManagerServer struct {
}

func (UnimplementedSchoolManagerServer) Behavior(context.Context, *Helper) (*StudentList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Behavior not implemented")
}
func (UnimplementedSchoolManagerServer) Exceptional(context.Context, *Helper) (*StudentList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Exceptional not implemented")
}
func (UnimplementedSchoolManagerServer) mustEmbedUnimplementedSchoolManagerServer() {}

// UnsafeSchoolManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SchoolManagerServer will
// result in compilation errors.
type UnsafeSchoolManagerServer interface {
	mustEmbedUnimplementedSchoolManagerServer()
}

func RegisterSchoolManagerServer(s grpc.ServiceRegistrar, srv SchoolManagerServer) {
	s.RegisterService(&SchoolManager_ServiceDesc, srv)
}

func _SchoolManager_Behavior_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Helper)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchoolManagerServer).Behavior(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manager.SchoolManager/Behavior",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchoolManagerServer).Behavior(ctx, req.(*Helper))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchoolManager_Exceptional_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Helper)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchoolManagerServer).Exceptional(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manager.SchoolManager/Exceptional",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchoolManagerServer).Exceptional(ctx, req.(*Helper))
	}
	return interceptor(ctx, in, info, handler)
}

// SchoolManager_ServiceDesc is the grpc.ServiceDesc for SchoolManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SchoolManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "manager.SchoolManager",
	HandlerType: (*SchoolManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Behavior",
			Handler:    _SchoolManager_Behavior_Handler,
		},
		{
			MethodName: "Exceptional",
			Handler:    _SchoolManager_Exceptional_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/manager.proto",
}