// Code generated by protoc-gen-go. DO NOT EDIT.
// source: application/model/grpc_streaming_inputter.proto

/*
Package model is a generated protocol buffer package.

It is generated from these files:
	application/model/grpc_streaming_inputter.proto
	application/model/grpc_unary_inputter.proto
	application/model/types.proto

It has these top-level messages:
	Request
	Empty
*/
package model

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for GrpcStreamingInputterService service

type GrpcStreamingInputterServiceClient interface {
	// Accepts a message but doesn't return anything back
	MakeRequest(ctx context.Context, opts ...grpc.CallOption) (GrpcStreamingInputterService_MakeRequestClient, error)
}

type grpcStreamingInputterServiceClient struct {
	cc *grpc.ClientConn
}

func NewGrpcStreamingInputterServiceClient(cc *grpc.ClientConn) GrpcStreamingInputterServiceClient {
	return &grpcStreamingInputterServiceClient{cc}
}

func (c *grpcStreamingInputterServiceClient) MakeRequest(ctx context.Context, opts ...grpc.CallOption) (GrpcStreamingInputterService_MakeRequestClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_GrpcStreamingInputterService_serviceDesc.Streams[0], c.cc, "/model.GrpcStreamingInputterService/MakeRequest", opts...)
	if err != nil {
		return nil, err
	}
	x := &grpcStreamingInputterServiceMakeRequestClient{stream}
	return x, nil
}

type GrpcStreamingInputterService_MakeRequestClient interface {
	Send(*Request) error
	CloseAndRecv() (*Empty, error)
	grpc.ClientStream
}

type grpcStreamingInputterServiceMakeRequestClient struct {
	grpc.ClientStream
}

func (x *grpcStreamingInputterServiceMakeRequestClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *grpcStreamingInputterServiceMakeRequestClient) CloseAndRecv() (*Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for GrpcStreamingInputterService service

type GrpcStreamingInputterServiceServer interface {
	// Accepts a message but doesn't return anything back
	MakeRequest(GrpcStreamingInputterService_MakeRequestServer) error
}

func RegisterGrpcStreamingInputterServiceServer(s *grpc.Server, srv GrpcStreamingInputterServiceServer) {
	s.RegisterService(&_GrpcStreamingInputterService_serviceDesc, srv)
}

func _GrpcStreamingInputterService_MakeRequest_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GrpcStreamingInputterServiceServer).MakeRequest(&grpcStreamingInputterServiceMakeRequestServer{stream})
}

type GrpcStreamingInputterService_MakeRequestServer interface {
	SendAndClose(*Empty) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type grpcStreamingInputterServiceMakeRequestServer struct {
	grpc.ServerStream
}

func (x *grpcStreamingInputterServiceMakeRequestServer) SendAndClose(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *grpcStreamingInputterServiceMakeRequestServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _GrpcStreamingInputterService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "model.GrpcStreamingInputterService",
	HandlerType: (*GrpcStreamingInputterServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "MakeRequest",
			Handler:       _GrpcStreamingInputterService_MakeRequest_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "application/model/grpc_streaming_inputter.proto",
}

func init() { proto.RegisterFile("application/model/grpc_streaming_inputter.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 150 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4f, 0x2c, 0x28, 0xc8,
	0xc9, 0x4c, 0x4e, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0xcf, 0xcd, 0x4f, 0x49, 0xcd, 0xd1, 0x4f, 0x2f,
	0x2a, 0x48, 0x8e, 0x2f, 0x2e, 0x29, 0x4a, 0x4d, 0xcc, 0xcd, 0xcc, 0x4b, 0x8f, 0xcf, 0xcc, 0x2b,
	0x28, 0x2d, 0x29, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x2b, 0x92,
	0x92, 0xc5, 0xd4, 0x57, 0x52, 0x59, 0x90, 0x5a, 0x0c, 0x51, 0x65, 0xe4, 0xcf, 0x25, 0xe3, 0x5e,
	0x54, 0x90, 0x1c, 0x0c, 0x33, 0xc5, 0x13, 0x6a, 0x48, 0x70, 0x6a, 0x51, 0x59, 0x66, 0x72, 0xaa,
	0x90, 0x3e, 0x17, 0xb7, 0x6f, 0x62, 0x76, 0x6a, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x10,
	0x9f, 0x1e, 0xd8, 0x08, 0x3d, 0x28, 0x5f, 0x8a, 0x07, 0xca, 0x77, 0xcd, 0x2d, 0x28, 0xa9, 0x54,
	0x62, 0xd0, 0x60, 0x4c, 0x62, 0x03, 0x9b, 0x6b, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x47, 0x4b,
	0xf1, 0x28, 0xb0, 0x00, 0x00, 0x00,
}
