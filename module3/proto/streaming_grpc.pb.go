// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: proto/streaming.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	StreamingService_StreamServerTime_FullMethodName = "/streaming.StreamingService/StreamServerTime"
	StreamingService_LogStream_FullMethodName        = "/streaming.StreamingService/LogStream"
	StreamingService_Echo_FullMethodName             = "/streaming.StreamingService/Echo"
)

// StreamingServiceClient is the client API for StreamingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StreamingServiceClient interface {
	StreamServerTime(ctx context.Context, in *StreamServerTimeRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[StreamServerTimeResponse], error)
	LogStream(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[LogStreamRequest, LogStreamResponse], error)
	Echo(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[EchoRequest, EchoResponse], error)
}

type streamingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamingServiceClient(cc grpc.ClientConnInterface) StreamingServiceClient {
	return &streamingServiceClient{cc}
}

func (c *streamingServiceClient) StreamServerTime(ctx context.Context, in *StreamServerTimeRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[StreamServerTimeResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &StreamingService_ServiceDesc.Streams[0], StreamingService_StreamServerTime_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[StreamServerTimeRequest, StreamServerTimeResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type StreamingService_StreamServerTimeClient = grpc.ServerStreamingClient[StreamServerTimeResponse]

func (c *streamingServiceClient) LogStream(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[LogStreamRequest, LogStreamResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &StreamingService_ServiceDesc.Streams[1], StreamingService_LogStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[LogStreamRequest, LogStreamResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type StreamingService_LogStreamClient = grpc.ClientStreamingClient[LogStreamRequest, LogStreamResponse]

func (c *streamingServiceClient) Echo(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[EchoRequest, EchoResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &StreamingService_ServiceDesc.Streams[2], StreamingService_Echo_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[EchoRequest, EchoResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type StreamingService_EchoClient = grpc.BidiStreamingClient[EchoRequest, EchoResponse]

// StreamingServiceServer is the server API for StreamingService service.
// All implementations must embed UnimplementedStreamingServiceServer
// for forward compatibility.
type StreamingServiceServer interface {
	StreamServerTime(*StreamServerTimeRequest, grpc.ServerStreamingServer[StreamServerTimeResponse]) error
	LogStream(grpc.ClientStreamingServer[LogStreamRequest, LogStreamResponse]) error
	Echo(grpc.BidiStreamingServer[EchoRequest, EchoResponse]) error
	mustEmbedUnimplementedStreamingServiceServer()
}

// UnimplementedStreamingServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedStreamingServiceServer struct{}

func (UnimplementedStreamingServiceServer) StreamServerTime(*StreamServerTimeRequest, grpc.ServerStreamingServer[StreamServerTimeResponse]) error {
	return status.Errorf(codes.Unimplemented, "method StreamServerTime not implemented")
}
func (UnimplementedStreamingServiceServer) LogStream(grpc.ClientStreamingServer[LogStreamRequest, LogStreamResponse]) error {
	return status.Errorf(codes.Unimplemented, "method LogStream not implemented")
}
func (UnimplementedStreamingServiceServer) Echo(grpc.BidiStreamingServer[EchoRequest, EchoResponse]) error {
	return status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (UnimplementedStreamingServiceServer) mustEmbedUnimplementedStreamingServiceServer() {}
func (UnimplementedStreamingServiceServer) testEmbeddedByValue()                          {}

// UnsafeStreamingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StreamingServiceServer will
// result in compilation errors.
type UnsafeStreamingServiceServer interface {
	mustEmbedUnimplementedStreamingServiceServer()
}

func RegisterStreamingServiceServer(s grpc.ServiceRegistrar, srv StreamingServiceServer) {
	// If the following call pancis, it indicates UnimplementedStreamingServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&StreamingService_ServiceDesc, srv)
}

func _StreamingService_StreamServerTime_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamServerTimeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamingServiceServer).StreamServerTime(m, &grpc.GenericServerStream[StreamServerTimeRequest, StreamServerTimeResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type StreamingService_StreamServerTimeServer = grpc.ServerStreamingServer[StreamServerTimeResponse]

func _StreamingService_LogStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamingServiceServer).LogStream(&grpc.GenericServerStream[LogStreamRequest, LogStreamResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type StreamingService_LogStreamServer = grpc.ClientStreamingServer[LogStreamRequest, LogStreamResponse]

func _StreamingService_Echo_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamingServiceServer).Echo(&grpc.GenericServerStream[EchoRequest, EchoResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type StreamingService_EchoServer = grpc.BidiStreamingServer[EchoRequest, EchoResponse]

// StreamingService_ServiceDesc is the grpc.ServiceDesc for StreamingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StreamingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "streaming.StreamingService",
	HandlerType: (*StreamingServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamServerTime",
			Handler:       _StreamingService_StreamServerTime_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "LogStream",
			Handler:       _StreamingService_LogStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Echo",
			Handler:       _StreamingService_Echo_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/streaming.proto",
}
