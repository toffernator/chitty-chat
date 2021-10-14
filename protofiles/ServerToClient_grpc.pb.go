// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protofiles

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

// ServerToClientClient is the client API for ServerToClient service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServerToClientClient interface {
	Broadcast(ctx context.Context, in *Message, opts ...grpc.CallOption) (*StatusOk, error)
}

type serverToClientClient struct {
	cc grpc.ClientConnInterface
}

func NewServerToClientClient(cc grpc.ClientConnInterface) ServerToClientClient {
	return &serverToClientClient{cc}
}

func (c *serverToClientClient) Broadcast(ctx context.Context, in *Message, opts ...grpc.CallOption) (*StatusOk, error) {
	out := new(StatusOk)
	err := c.cc.Invoke(ctx, "/ServerToClient/Broadcast", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServerToClientServer is the server API for ServerToClient service.
// All implementations must embed UnimplementedServerToClientServer
// for forward compatibility
type ServerToClientServer interface {
	Broadcast(context.Context, *Message) (*StatusOk, error)
	mustEmbedUnimplementedServerToClientServer()
}

// UnimplementedServerToClientServer must be embedded to have forward compatible implementations.
type UnimplementedServerToClientServer struct {
}

func (UnimplementedServerToClientServer) Broadcast(context.Context, *Message) (*StatusOk, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Broadcast not implemented")
}
func (UnimplementedServerToClientServer) mustEmbedUnimplementedServerToClientServer() {}

// UnsafeServerToClientServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServerToClientServer will
// result in compilation errors.
type UnsafeServerToClientServer interface {
	mustEmbedUnimplementedServerToClientServer()
}

func RegisterServerToClientServer(s grpc.ServiceRegistrar, srv ServerToClientServer) {
	s.RegisterService(&ServerToClient_ServiceDesc, srv)
}

func _ServerToClient_Broadcast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerToClientServer).Broadcast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ServerToClient/Broadcast",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerToClientServer).Broadcast(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

// ServerToClient_ServiceDesc is the grpc.ServiceDesc for ServerToClient service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServerToClient_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ServerToClient",
	HandlerType: (*ServerToClientServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Broadcast",
			Handler:    _ServerToClient_Broadcast_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ServerToClient.proto",
}
