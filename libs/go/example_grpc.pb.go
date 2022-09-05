// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: example.proto

package _go

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

// ExampleServiceClient is the client API for ExampleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExampleServiceClient interface {
	Query(ctx context.Context, in *Request, opts ...grpc.CallOption) (ExampleService_QueryClient, error)
}

type exampleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExampleServiceClient(cc grpc.ClientConnInterface) ExampleServiceClient {
	return &exampleServiceClient{cc}
}

func (c *exampleServiceClient) Query(ctx context.Context, in *Request, opts ...grpc.CallOption) (ExampleService_QueryClient, error) {
	stream, err := c.cc.NewStream(ctx, &ExampleService_ServiceDesc.Streams[0], "/example.ExampleService/Query", opts...)
	if err != nil {
		return nil, err
	}
	x := &exampleServiceQueryClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ExampleService_QueryClient interface {
	Recv() (*Result, error)
	grpc.ClientStream
}

type exampleServiceQueryClient struct {
	grpc.ClientStream
}

func (x *exampleServiceQueryClient) Recv() (*Result, error) {
	m := new(Result)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ExampleServiceServer is the server API for ExampleService service.
// All implementations must embed UnimplementedExampleServiceServer
// for forward compatibility
type ExampleServiceServer interface {
	Query(*Request, ExampleService_QueryServer) error
	mustEmbedUnimplementedExampleServiceServer()
}

// UnimplementedExampleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedExampleServiceServer struct {
}

func (UnimplementedExampleServiceServer) Query(*Request, ExampleService_QueryServer) error {
	return status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (UnimplementedExampleServiceServer) mustEmbedUnimplementedExampleServiceServer() {}

// UnsafeExampleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExampleServiceServer will
// result in compilation errors.
type UnsafeExampleServiceServer interface {
	mustEmbedUnimplementedExampleServiceServer()
}

func RegisterExampleServiceServer(s grpc.ServiceRegistrar, srv ExampleServiceServer) {
	s.RegisterService(&ExampleService_ServiceDesc, srv)
}

func _ExampleService_Query_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ExampleServiceServer).Query(m, &exampleServiceQueryServer{stream})
}

type ExampleService_QueryServer interface {
	Send(*Result) error
	grpc.ServerStream
}

type exampleServiceQueryServer struct {
	grpc.ServerStream
}

func (x *exampleServiceQueryServer) Send(m *Result) error {
	return x.ServerStream.SendMsg(m)
}

// ExampleService_ServiceDesc is the grpc.ServiceDesc for ExampleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExampleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "example.ExampleService",
	HandlerType: (*ExampleServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Query",
			Handler:       _ExampleService_Query_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "example.proto",
}
