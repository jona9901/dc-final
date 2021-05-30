// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// FilterClient is the client API for Filter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FilterClient interface {
	// Sends a greeting
	ApplyFilter(ctx context.Context, in *FilterRequest, opts ...grpc.CallOption) (*FilterReply, error)
}

type filterClient struct {
	cc grpc.ClientConnInterface
}

func NewFilterClient(cc grpc.ClientConnInterface) FilterClient {
	return &filterClient{cc}
}

func (c *filterClient) ApplyFilter(ctx context.Context, in *FilterRequest, opts ...grpc.CallOption) (*FilterReply, error) {
	out := new(FilterReply)
	err := c.cc.Invoke(ctx, "/proto.Filter/ApplyFilter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FilterServer is the server API for Filter service.
// All implementations must embed UnimplementedFilterServer
// for forward compatibility
type FilterServer interface {
	// Sends a greeting
	ApplyFilter(context.Context, *FilterRequest) (*FilterReply, error)
	mustEmbedUnimplementedFilterServer()
}

// UnimplementedFilterServer must be embedded to have forward compatible implementations.
type UnimplementedFilterServer struct {
}

func (UnimplementedFilterServer) ApplyFilter(context.Context, *FilterRequest) (*FilterReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplyFilter not implemented")
}
func (UnimplementedFilterServer) mustEmbedUnimplementedFilterServer() {}

// UnsafeFilterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FilterServer will
// result in compilation errors.
type UnsafeFilterServer interface {
	mustEmbedUnimplementedFilterServer()
}

func RegisterFilterServer(s grpc.ServiceRegistrar, srv FilterServer) {
	s.RegisterService(&Filter_ServiceDesc, srv)
}

func _Filter_ApplyFilter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilterServer).ApplyFilter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Filter/ApplyFilter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilterServer).ApplyFilter(ctx, req.(*FilterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Filter_ServiceDesc is the grpc.ServiceDesc for Filter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Filter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Filter",
	HandlerType: (*FilterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ApplyFilter",
			Handler:    _Filter_ApplyFilter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "filters.proto",
}