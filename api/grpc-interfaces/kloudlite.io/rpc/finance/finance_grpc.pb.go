// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: finance.proto

package finance

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

// FinanceClient is the client API for Finance service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FinanceClient interface {
	StartBillable(ctx context.Context, in *StartBillableIn, opts ...grpc.CallOption) (*StartBillableOut, error)
	StopBillable(ctx context.Context, in *StopBillableIn, opts ...grpc.CallOption) (*StopBillableOut, error)
}

type financeClient struct {
	cc grpc.ClientConnInterface
}

func NewFinanceClient(cc grpc.ClientConnInterface) FinanceClient {
	return &financeClient{cc}
}

func (c *financeClient) StartBillable(ctx context.Context, in *StartBillableIn, opts ...grpc.CallOption) (*StartBillableOut, error) {
	out := new(StartBillableOut)
	err := c.cc.Invoke(ctx, "/Finance/startBillable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financeClient) StopBillable(ctx context.Context, in *StopBillableIn, opts ...grpc.CallOption) (*StopBillableOut, error) {
	out := new(StopBillableOut)
	err := c.cc.Invoke(ctx, "/Finance/stopBillable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FinanceServer is the server API for Finance service.
// All implementations must embed UnimplementedFinanceServer
// for forward compatibility
type FinanceServer interface {
	StartBillable(context.Context, *StartBillableIn) (*StartBillableOut, error)
	StopBillable(context.Context, *StopBillableIn) (*StopBillableOut, error)
	mustEmbedUnimplementedFinanceServer()
}

// UnimplementedFinanceServer must be embedded to have forward compatible implementations.
type UnimplementedFinanceServer struct {
}

func (UnimplementedFinanceServer) StartBillable(context.Context, *StartBillableIn) (*StartBillableOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartBillable not implemented")
}
func (UnimplementedFinanceServer) StopBillable(context.Context, *StopBillableIn) (*StopBillableOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopBillable not implemented")
}
func (UnimplementedFinanceServer) mustEmbedUnimplementedFinanceServer() {}

// UnsafeFinanceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FinanceServer will
// result in compilation errors.
type UnsafeFinanceServer interface {
	mustEmbedUnimplementedFinanceServer()
}

func RegisterFinanceServer(s grpc.ServiceRegistrar, srv FinanceServer) {
	s.RegisterService(&Finance_ServiceDesc, srv)
}

func _Finance_StartBillable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartBillableIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinanceServer).StartBillable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Finance/startBillable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinanceServer).StartBillable(ctx, req.(*StartBillableIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _Finance_StopBillable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopBillableIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinanceServer).StopBillable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Finance/stopBillable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinanceServer).StopBillable(ctx, req.(*StopBillableIn))
	}
	return interceptor(ctx, in, info, handler)
}

// Finance_ServiceDesc is the grpc.ServiceDesc for Finance service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Finance_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Finance",
	HandlerType: (*FinanceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "startBillable",
			Handler:    _Finance_StartBillable_Handler,
		},
		{
			MethodName: "stopBillable",
			Handler:    _Finance_StopBillable_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "finance.proto",
}
