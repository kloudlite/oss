// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.3
// source: messages.proto

package messages

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

const (
	MessageDispatchService_GetAccessToken_FullMethodName           = "/MessageDispatchService/GetAccessToken"
	MessageDispatchService_SendActions_FullMethodName              = "/MessageDispatchService/SendActions"
	MessageDispatchService_ReceiveErrors_FullMethodName            = "/MessageDispatchService/ReceiveErrors"
	MessageDispatchService_ReceiveResourceUpdates_FullMethodName   = "/MessageDispatchService/ReceiveResourceUpdates"
	MessageDispatchService_ReceiveInfraUpdates_FullMethodName      = "/MessageDispatchService/ReceiveInfraUpdates"
	MessageDispatchService_ReceiveBYOCClientUpdates_FullMethodName = "/MessageDispatchService/ReceiveBYOCClientUpdates"
)

// MessageDispatchServiceClient is the client API for MessageDispatchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessageDispatchServiceClient interface {
	GetAccessToken(ctx context.Context, in *GetClusterTokenIn, opts ...grpc.CallOption) (*GetClusterTokenOut, error)
	SendActions(ctx context.Context, in *StreamActionsRequest, opts ...grpc.CallOption) (MessageDispatchService_SendActionsClient, error)
	ReceiveErrors(ctx context.Context, opts ...grpc.CallOption) (MessageDispatchService_ReceiveErrorsClient, error)
	ReceiveResourceUpdates(ctx context.Context, opts ...grpc.CallOption) (MessageDispatchService_ReceiveResourceUpdatesClient, error)
	ReceiveInfraUpdates(ctx context.Context, opts ...grpc.CallOption) (MessageDispatchService_ReceiveInfraUpdatesClient, error)
	ReceiveBYOCClientUpdates(ctx context.Context, opts ...grpc.CallOption) (MessageDispatchService_ReceiveBYOCClientUpdatesClient, error)
}

type messageDispatchServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageDispatchServiceClient(cc grpc.ClientConnInterface) MessageDispatchServiceClient {
	return &messageDispatchServiceClient{cc}
}

func (c *messageDispatchServiceClient) GetAccessToken(ctx context.Context, in *GetClusterTokenIn, opts ...grpc.CallOption) (*GetClusterTokenOut, error) {
	out := new(GetClusterTokenOut)
	err := c.cc.Invoke(ctx, MessageDispatchService_GetAccessToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageDispatchServiceClient) SendActions(ctx context.Context, in *StreamActionsRequest, opts ...grpc.CallOption) (MessageDispatchService_SendActionsClient, error) {
	stream, err := c.cc.NewStream(ctx, &MessageDispatchService_ServiceDesc.Streams[0], MessageDispatchService_SendActions_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &messageDispatchServiceSendActionsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MessageDispatchService_SendActionsClient interface {
	Recv() (*Action, error)
	grpc.ClientStream
}

type messageDispatchServiceSendActionsClient struct {
	grpc.ClientStream
}

func (x *messageDispatchServiceSendActionsClient) Recv() (*Action, error) {
	m := new(Action)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *messageDispatchServiceClient) ReceiveErrors(ctx context.Context, opts ...grpc.CallOption) (MessageDispatchService_ReceiveErrorsClient, error) {
	stream, err := c.cc.NewStream(ctx, &MessageDispatchService_ServiceDesc.Streams[1], MessageDispatchService_ReceiveErrors_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &messageDispatchServiceReceiveErrorsClient{stream}
	return x, nil
}

type MessageDispatchService_ReceiveErrorsClient interface {
	Send(*ErrorData) error
	CloseAndRecv() (*Empty, error)
	grpc.ClientStream
}

type messageDispatchServiceReceiveErrorsClient struct {
	grpc.ClientStream
}

func (x *messageDispatchServiceReceiveErrorsClient) Send(m *ErrorData) error {
	return x.ClientStream.SendMsg(m)
}

func (x *messageDispatchServiceReceiveErrorsClient) CloseAndRecv() (*Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *messageDispatchServiceClient) ReceiveResourceUpdates(ctx context.Context, opts ...grpc.CallOption) (MessageDispatchService_ReceiveResourceUpdatesClient, error) {
	stream, err := c.cc.NewStream(ctx, &MessageDispatchService_ServiceDesc.Streams[2], MessageDispatchService_ReceiveResourceUpdates_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &messageDispatchServiceReceiveResourceUpdatesClient{stream}
	return x, nil
}

type MessageDispatchService_ReceiveResourceUpdatesClient interface {
	Send(*ResourceUpdate) error
	CloseAndRecv() (*Empty, error)
	grpc.ClientStream
}

type messageDispatchServiceReceiveResourceUpdatesClient struct {
	grpc.ClientStream
}

func (x *messageDispatchServiceReceiveResourceUpdatesClient) Send(m *ResourceUpdate) error {
	return x.ClientStream.SendMsg(m)
}

func (x *messageDispatchServiceReceiveResourceUpdatesClient) CloseAndRecv() (*Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *messageDispatchServiceClient) ReceiveInfraUpdates(ctx context.Context, opts ...grpc.CallOption) (MessageDispatchService_ReceiveInfraUpdatesClient, error) {
	stream, err := c.cc.NewStream(ctx, &MessageDispatchService_ServiceDesc.Streams[3], MessageDispatchService_ReceiveInfraUpdates_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &messageDispatchServiceReceiveInfraUpdatesClient{stream}
	return x, nil
}

type MessageDispatchService_ReceiveInfraUpdatesClient interface {
	Send(*InfraUpdate) error
	CloseAndRecv() (*Empty, error)
	grpc.ClientStream
}

type messageDispatchServiceReceiveInfraUpdatesClient struct {
	grpc.ClientStream
}

func (x *messageDispatchServiceReceiveInfraUpdatesClient) Send(m *InfraUpdate) error {
	return x.ClientStream.SendMsg(m)
}

func (x *messageDispatchServiceReceiveInfraUpdatesClient) CloseAndRecv() (*Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *messageDispatchServiceClient) ReceiveBYOCClientUpdates(ctx context.Context, opts ...grpc.CallOption) (MessageDispatchService_ReceiveBYOCClientUpdatesClient, error) {
	stream, err := c.cc.NewStream(ctx, &MessageDispatchService_ServiceDesc.Streams[4], MessageDispatchService_ReceiveBYOCClientUpdates_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &messageDispatchServiceReceiveBYOCClientUpdatesClient{stream}
	return x, nil
}

type MessageDispatchService_ReceiveBYOCClientUpdatesClient interface {
	Send(*BYOCClientUpdate) error
	CloseAndRecv() (*Empty, error)
	grpc.ClientStream
}

type messageDispatchServiceReceiveBYOCClientUpdatesClient struct {
	grpc.ClientStream
}

func (x *messageDispatchServiceReceiveBYOCClientUpdatesClient) Send(m *BYOCClientUpdate) error {
	return x.ClientStream.SendMsg(m)
}

func (x *messageDispatchServiceReceiveBYOCClientUpdatesClient) CloseAndRecv() (*Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MessageDispatchServiceServer is the server API for MessageDispatchService service.
// All implementations must embed UnimplementedMessageDispatchServiceServer
// for forward compatibility
type MessageDispatchServiceServer interface {
	GetAccessToken(context.Context, *GetClusterTokenIn) (*GetClusterTokenOut, error)
	SendActions(*StreamActionsRequest, MessageDispatchService_SendActionsServer) error
	ReceiveErrors(MessageDispatchService_ReceiveErrorsServer) error
	ReceiveResourceUpdates(MessageDispatchService_ReceiveResourceUpdatesServer) error
	ReceiveInfraUpdates(MessageDispatchService_ReceiveInfraUpdatesServer) error
	ReceiveBYOCClientUpdates(MessageDispatchService_ReceiveBYOCClientUpdatesServer) error
	mustEmbedUnimplementedMessageDispatchServiceServer()
}

// UnimplementedMessageDispatchServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMessageDispatchServiceServer struct {
}

func (UnimplementedMessageDispatchServiceServer) GetAccessToken(context.Context, *GetClusterTokenIn) (*GetClusterTokenOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccessToken not implemented")
}
func (UnimplementedMessageDispatchServiceServer) SendActions(*StreamActionsRequest, MessageDispatchService_SendActionsServer) error {
	return status.Errorf(codes.Unimplemented, "method SendActions not implemented")
}
func (UnimplementedMessageDispatchServiceServer) ReceiveErrors(MessageDispatchService_ReceiveErrorsServer) error {
	return status.Errorf(codes.Unimplemented, "method ReceiveErrors not implemented")
}
func (UnimplementedMessageDispatchServiceServer) ReceiveResourceUpdates(MessageDispatchService_ReceiveResourceUpdatesServer) error {
	return status.Errorf(codes.Unimplemented, "method ReceiveResourceUpdates not implemented")
}
func (UnimplementedMessageDispatchServiceServer) ReceiveInfraUpdates(MessageDispatchService_ReceiveInfraUpdatesServer) error {
	return status.Errorf(codes.Unimplemented, "method ReceiveInfraUpdates not implemented")
}
func (UnimplementedMessageDispatchServiceServer) ReceiveBYOCClientUpdates(MessageDispatchService_ReceiveBYOCClientUpdatesServer) error {
	return status.Errorf(codes.Unimplemented, "method ReceiveBYOCClientUpdates not implemented")
}
func (UnimplementedMessageDispatchServiceServer) mustEmbedUnimplementedMessageDispatchServiceServer() {
}

// UnsafeMessageDispatchServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessageDispatchServiceServer will
// result in compilation errors.
type UnsafeMessageDispatchServiceServer interface {
	mustEmbedUnimplementedMessageDispatchServiceServer()
}

func RegisterMessageDispatchServiceServer(s grpc.ServiceRegistrar, srv MessageDispatchServiceServer) {
	s.RegisterService(&MessageDispatchService_ServiceDesc, srv)
}

func _MessageDispatchService_GetAccessToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClusterTokenIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageDispatchServiceServer).GetAccessToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageDispatchService_GetAccessToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageDispatchServiceServer).GetAccessToken(ctx, req.(*GetClusterTokenIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageDispatchService_SendActions_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamActionsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MessageDispatchServiceServer).SendActions(m, &messageDispatchServiceSendActionsServer{stream})
}

type MessageDispatchService_SendActionsServer interface {
	Send(*Action) error
	grpc.ServerStream
}

type messageDispatchServiceSendActionsServer struct {
	grpc.ServerStream
}

func (x *messageDispatchServiceSendActionsServer) Send(m *Action) error {
	return x.ServerStream.SendMsg(m)
}

func _MessageDispatchService_ReceiveErrors_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MessageDispatchServiceServer).ReceiveErrors(&messageDispatchServiceReceiveErrorsServer{stream})
}

type MessageDispatchService_ReceiveErrorsServer interface {
	SendAndClose(*Empty) error
	Recv() (*ErrorData, error)
	grpc.ServerStream
}

type messageDispatchServiceReceiveErrorsServer struct {
	grpc.ServerStream
}

func (x *messageDispatchServiceReceiveErrorsServer) SendAndClose(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *messageDispatchServiceReceiveErrorsServer) Recv() (*ErrorData, error) {
	m := new(ErrorData)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _MessageDispatchService_ReceiveResourceUpdates_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MessageDispatchServiceServer).ReceiveResourceUpdates(&messageDispatchServiceReceiveResourceUpdatesServer{stream})
}

type MessageDispatchService_ReceiveResourceUpdatesServer interface {
	SendAndClose(*Empty) error
	Recv() (*ResourceUpdate, error)
	grpc.ServerStream
}

type messageDispatchServiceReceiveResourceUpdatesServer struct {
	grpc.ServerStream
}

func (x *messageDispatchServiceReceiveResourceUpdatesServer) SendAndClose(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *messageDispatchServiceReceiveResourceUpdatesServer) Recv() (*ResourceUpdate, error) {
	m := new(ResourceUpdate)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _MessageDispatchService_ReceiveInfraUpdates_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MessageDispatchServiceServer).ReceiveInfraUpdates(&messageDispatchServiceReceiveInfraUpdatesServer{stream})
}

type MessageDispatchService_ReceiveInfraUpdatesServer interface {
	SendAndClose(*Empty) error
	Recv() (*InfraUpdate, error)
	grpc.ServerStream
}

type messageDispatchServiceReceiveInfraUpdatesServer struct {
	grpc.ServerStream
}

func (x *messageDispatchServiceReceiveInfraUpdatesServer) SendAndClose(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *messageDispatchServiceReceiveInfraUpdatesServer) Recv() (*InfraUpdate, error) {
	m := new(InfraUpdate)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _MessageDispatchService_ReceiveBYOCClientUpdates_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MessageDispatchServiceServer).ReceiveBYOCClientUpdates(&messageDispatchServiceReceiveBYOCClientUpdatesServer{stream})
}

type MessageDispatchService_ReceiveBYOCClientUpdatesServer interface {
	SendAndClose(*Empty) error
	Recv() (*BYOCClientUpdate, error)
	grpc.ServerStream
}

type messageDispatchServiceReceiveBYOCClientUpdatesServer struct {
	grpc.ServerStream
}

func (x *messageDispatchServiceReceiveBYOCClientUpdatesServer) SendAndClose(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *messageDispatchServiceReceiveBYOCClientUpdatesServer) Recv() (*BYOCClientUpdate, error) {
	m := new(BYOCClientUpdate)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MessageDispatchService_ServiceDesc is the grpc.ServiceDesc for MessageDispatchService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MessageDispatchService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "MessageDispatchService",
	HandlerType: (*MessageDispatchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAccessToken",
			Handler:    _MessageDispatchService_GetAccessToken_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendActions",
			Handler:       _MessageDispatchService_SendActions_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ReceiveErrors",
			Handler:       _MessageDispatchService_ReceiveErrors_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "ReceiveResourceUpdates",
			Handler:       _MessageDispatchService_ReceiveResourceUpdates_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "ReceiveInfraUpdates",
			Handler:       _MessageDispatchService_ReceiveInfraUpdates_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "ReceiveBYOCClientUpdates",
			Handler:       _MessageDispatchService_ReceiveBYOCClientUpdates_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "messages.proto",
}
