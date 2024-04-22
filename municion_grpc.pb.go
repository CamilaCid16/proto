// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: proto/municion.proto

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

// NotifyServiceClient is the client API for NotifyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotifyServiceClient interface {
	Exchange(ctx context.Context, in *Notify, opts ...grpc.CallOption) (*Notify, error)
	SendUsers(ctx context.Context, in *RegUsers, opts ...grpc.CallOption) (*RegUsers, error)
}

type notifyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNotifyServiceClient(cc grpc.ClientConnInterface) NotifyServiceClient {
	return &notifyServiceClient{cc}
}

func (c *notifyServiceClient) Exchange(ctx context.Context, in *Notify, opts ...grpc.CallOption) (*Notify, error) {
	out := new(Notify)
	err := c.cc.Invoke(ctx, "/grpc.NotifyService/Exchange", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifyServiceClient) SendUsers(ctx context.Context, in *RegUsers, opts ...grpc.CallOption) (*RegUsers, error) {
	out := new(RegUsers)
	err := c.cc.Invoke(ctx, "/grpc.NotifyService/SendUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotifyServiceServer is the server API for NotifyService service.
// All implementations must embed UnimplementedNotifyServiceServer
// for forward compatibility
type NotifyServiceServer interface {
	Exchange(context.Context, *Notify) (*Notify, error)
	SendUsers(context.Context, *RegUsers) (*RegUsers, error)
	mustEmbedUnimplementedNotifyServiceServer()
}

// UnimplementedNotifyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNotifyServiceServer struct {
}

func (UnimplementedNotifyServiceServer) Exchange(context.Context, *Notify) (*Notify, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Exchange not implemented")
}
func (UnimplementedNotifyServiceServer) SendUsers(context.Context, *RegUsers) (*RegUsers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendUsers not implemented")
}
func (UnimplementedNotifyServiceServer) mustEmbedUnimplementedNotifyServiceServer() {}

// UnsafeNotifyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotifyServiceServer will
// result in compilation errors.
type UnsafeNotifyServiceServer interface {
	mustEmbedUnimplementedNotifyServiceServer()
}

func RegisterNotifyServiceServer(s grpc.ServiceRegistrar, srv NotifyServiceServer) {
	s.RegisterService(&NotifyService_ServiceDesc, srv)
}

func _NotifyService_Exchange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Notify)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifyServiceServer).Exchange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.NotifyService/Exchange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifyServiceServer).Exchange(ctx, req.(*Notify))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotifyService_SendUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegUsers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifyServiceServer).SendUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.NotifyService/SendUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifyServiceServer).SendUsers(ctx, req.(*RegUsers))
	}
	return interceptor(ctx, in, info, handler)
}

// NotifyService_ServiceDesc is the grpc.ServiceDesc for NotifyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NotifyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.NotifyService",
	HandlerType: (*NotifyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Exchange",
			Handler:    _NotifyService_Exchange_Handler,
		},
		{
			MethodName: "SendUsers",
			Handler:    _NotifyService_SendUsers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/municion.proto",
}
