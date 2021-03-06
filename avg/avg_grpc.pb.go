// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package dproto2

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

// PrimeServiceClient is the client API for PrimeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PrimeServiceClient interface {
	AVG(ctx context.Context, opts ...grpc.CallOption) (PrimeService_AVGClient, error)
}

type primeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPrimeServiceClient(cc grpc.ClientConnInterface) PrimeServiceClient {
	return &primeServiceClient{cc}
}

func (c *primeServiceClient) AVG(ctx context.Context, opts ...grpc.CallOption) (PrimeService_AVGClient, error) {
	stream, err := c.cc.NewStream(ctx, &PrimeService_ServiceDesc.Streams[0], "/avg.PrimeService/AVG", opts...)
	if err != nil {
		return nil, err
	}
	x := &primeServiceAVGClient{stream}
	return x, nil
}

type PrimeService_AVGClient interface {
	Send(*Request) error
	CloseAndRecv() (*Response, error)
	grpc.ClientStream
}

type primeServiceAVGClient struct {
	grpc.ClientStream
}

func (x *primeServiceAVGClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *primeServiceAVGClient) CloseAndRecv() (*Response, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PrimeServiceServer is the server API for PrimeService service.
// All implementations must embed UnimplementedPrimeServiceServer
// for forward compatibility
type PrimeServiceServer interface {
	AVG(PrimeService_AVGServer) error
	mustEmbedUnimplementedPrimeServiceServer()
}

// UnimplementedPrimeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPrimeServiceServer struct {
}

func (UnimplementedPrimeServiceServer) AVG(PrimeService_AVGServer) error {
	return status.Errorf(codes.Unimplemented, "method AVG not implemented")
}
func (UnimplementedPrimeServiceServer) mustEmbedUnimplementedPrimeServiceServer() {}

// UnsafePrimeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PrimeServiceServer will
// result in compilation errors.
type UnsafePrimeServiceServer interface {
	mustEmbedUnimplementedPrimeServiceServer()
}

func RegisterPrimeServiceServer(s grpc.ServiceRegistrar, srv PrimeServiceServer) {
	s.RegisterService(&PrimeService_ServiceDesc, srv)
}

func _PrimeService_AVG_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PrimeServiceServer).AVG(&primeServiceAVGServer{stream})
}

type PrimeService_AVGServer interface {
	SendAndClose(*Response) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type primeServiceAVGServer struct {
	grpc.ServerStream
}

func (x *primeServiceAVGServer) SendAndClose(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *primeServiceAVGServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PrimeService_ServiceDesc is the grpc.ServiceDesc for PrimeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PrimeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "avg.PrimeService",
	HandlerType: (*PrimeServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "AVG",
			Handler:       _PrimeService_AVG_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "avg/avg.proto",
}
