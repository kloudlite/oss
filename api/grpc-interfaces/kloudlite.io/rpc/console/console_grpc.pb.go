// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: console.proto

package console

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

// ConsoleClient is the client API for Console service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConsoleClient interface {
	CreateDefaultCluster(ctx context.Context, in *CreateClusterIn, opts ...grpc.CallOption) (*CreateClusterOut, error)
}

type consoleClient struct {
	cc grpc.ClientConnInterface
}

func NewConsoleClient(cc grpc.ClientConnInterface) ConsoleClient {
	return &consoleClient{cc}
}

func (c *consoleClient) CreateDefaultCluster(ctx context.Context, in *CreateClusterIn, opts ...grpc.CallOption) (*CreateClusterOut, error) {
	out := new(CreateClusterOut)
	err := c.cc.Invoke(ctx, "/Console/CreateDefaultCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConsoleServer is the server API for Console service.
// All implementations must embed UnimplementedConsoleServer
// for forward compatibility
type ConsoleServer interface {
	CreateDefaultCluster(context.Context, *CreateClusterIn) (*CreateClusterOut, error)
	mustEmbedUnimplementedConsoleServer()
}

// UnimplementedConsoleServer must be embedded to have forward compatible implementations.
type UnimplementedConsoleServer struct {
}

func (UnimplementedConsoleServer) CreateDefaultCluster(context.Context, *CreateClusterIn) (*CreateClusterOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDefaultCluster not implemented")
}
func (UnimplementedConsoleServer) mustEmbedUnimplementedConsoleServer() {}

// UnsafeConsoleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConsoleServer will
// result in compilation errors.
type UnsafeConsoleServer interface {
	mustEmbedUnimplementedConsoleServer()
}

func RegisterConsoleServer(s grpc.ServiceRegistrar, srv ConsoleServer) {
	s.RegisterService(&Console_ServiceDesc, srv)
}

func _Console_CreateDefaultCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateClusterIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsoleServer).CreateDefaultCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Console/CreateDefaultCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsoleServer).CreateDefaultCluster(ctx, req.(*CreateClusterIn))
	}
	return interceptor(ctx, in, info, handler)
}

// Console_ServiceDesc is the grpc.ServiceDesc for Console service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Console_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Console",
	HandlerType: (*ConsoleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDefaultCluster",
			Handler:    _Console_CreateDefaultCluster_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "console.proto",
}
