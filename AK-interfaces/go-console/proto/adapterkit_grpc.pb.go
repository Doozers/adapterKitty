// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: adapterkit.proto

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

// AdapterKitServiceClient is the client API for AdapterKitService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdapterKitServiceClient interface {
	BiDirectionalAdapter(ctx context.Context, opts ...grpc.CallOption) (AdapterKitService_BiDirectionalAdapterClient, error)
	UniDirectionalAdapter(ctx context.Context, in *AdapterRequest, opts ...grpc.CallOption) (*AdapterResponse, error)
	ServerStreamingAdapter(ctx context.Context, in *AdapterRequest, opts ...grpc.CallOption) (AdapterKitService_ServerStreamingAdapterClient, error)
}

type adapterKitServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdapterKitServiceClient(cc grpc.ClientConnInterface) AdapterKitServiceClient {
	return &adapterKitServiceClient{cc}
}

func (c *adapterKitServiceClient) BiDirectionalAdapter(ctx context.Context, opts ...grpc.CallOption) (AdapterKitService_BiDirectionalAdapterClient, error) {
	stream, err := c.cc.NewStream(ctx, &AdapterKitService_ServiceDesc.Streams[0], "/AdapterKitService/BiDirectionalAdapter", opts...)
	if err != nil {
		return nil, err
	}
	x := &adapterKitServiceBiDirectionalAdapterClient{stream}
	return x, nil
}

type AdapterKitService_BiDirectionalAdapterClient interface {
	Send(*AdapterRequest) error
	Recv() (*AdapterResponse, error)
	grpc.ClientStream
}

type adapterKitServiceBiDirectionalAdapterClient struct {
	grpc.ClientStream
}

func (x *adapterKitServiceBiDirectionalAdapterClient) Send(m *AdapterRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *adapterKitServiceBiDirectionalAdapterClient) Recv() (*AdapterResponse, error) {
	m := new(AdapterResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *adapterKitServiceClient) UniDirectionalAdapter(ctx context.Context, in *AdapterRequest, opts ...grpc.CallOption) (*AdapterResponse, error) {
	out := new(AdapterResponse)
	err := c.cc.Invoke(ctx, "/AdapterKitService/UniDirectionalAdapter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adapterKitServiceClient) ServerStreamingAdapter(ctx context.Context, in *AdapterRequest, opts ...grpc.CallOption) (AdapterKitService_ServerStreamingAdapterClient, error) {
	stream, err := c.cc.NewStream(ctx, &AdapterKitService_ServiceDesc.Streams[1], "/AdapterKitService/ServerStreamingAdapter", opts...)
	if err != nil {
		return nil, err
	}
	x := &adapterKitServiceServerStreamingAdapterClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AdapterKitService_ServerStreamingAdapterClient interface {
	Recv() (*AdapterResponse, error)
	grpc.ClientStream
}

type adapterKitServiceServerStreamingAdapterClient struct {
	grpc.ClientStream
}

func (x *adapterKitServiceServerStreamingAdapterClient) Recv() (*AdapterResponse, error) {
	m := new(AdapterResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AdapterKitServiceServer is the server API for AdapterKitService service.
// All implementations must embed UnimplementedAdapterKitServiceServer
// for forward compatibility
type AdapterKitServiceServer interface {
	BiDirectionalAdapter(AdapterKitService_BiDirectionalAdapterServer) error
	UniDirectionalAdapter(context.Context, *AdapterRequest) (*AdapterResponse, error)
	ServerStreamingAdapter(*AdapterRequest, AdapterKitService_ServerStreamingAdapterServer) error
	mustEmbedUnimplementedAdapterKitServiceServer()
}

// UnimplementedAdapterKitServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAdapterKitServiceServer struct {
}

func (UnimplementedAdapterKitServiceServer) BiDirectionalAdapter(AdapterKitService_BiDirectionalAdapterServer) error {
	return status.Errorf(codes.Unimplemented, "method BiDirectionalAdapter not implemented")
}
func (UnimplementedAdapterKitServiceServer) UniDirectionalAdapter(context.Context, *AdapterRequest) (*AdapterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UniDirectionalAdapter not implemented")
}
func (UnimplementedAdapterKitServiceServer) ServerStreamingAdapter(*AdapterRequest, AdapterKitService_ServerStreamingAdapterServer) error {
	return status.Errorf(codes.Unimplemented, "method ServerStreamingAdapter not implemented")
}
func (UnimplementedAdapterKitServiceServer) mustEmbedUnimplementedAdapterKitServiceServer() {}

// UnsafeAdapterKitServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdapterKitServiceServer will
// result in compilation errors.
type UnsafeAdapterKitServiceServer interface {
	mustEmbedUnimplementedAdapterKitServiceServer()
}

func RegisterAdapterKitServiceServer(s grpc.ServiceRegistrar, srv AdapterKitServiceServer) {
	s.RegisterService(&AdapterKitService_ServiceDesc, srv)
}

func _AdapterKitService_BiDirectionalAdapter_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AdapterKitServiceServer).BiDirectionalAdapter(&adapterKitServiceBiDirectionalAdapterServer{stream})
}

type AdapterKitService_BiDirectionalAdapterServer interface {
	Send(*AdapterResponse) error
	Recv() (*AdapterRequest, error)
	grpc.ServerStream
}

type adapterKitServiceBiDirectionalAdapterServer struct {
	grpc.ServerStream
}

func (x *adapterKitServiceBiDirectionalAdapterServer) Send(m *AdapterResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *adapterKitServiceBiDirectionalAdapterServer) Recv() (*AdapterRequest, error) {
	m := new(AdapterRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _AdapterKitService_UniDirectionalAdapter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdapterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdapterKitServiceServer).UniDirectionalAdapter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AdapterKitService/UniDirectionalAdapter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdapterKitServiceServer).UniDirectionalAdapter(ctx, req.(*AdapterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdapterKitService_ServerStreamingAdapter_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(AdapterRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AdapterKitServiceServer).ServerStreamingAdapter(m, &adapterKitServiceServerStreamingAdapterServer{stream})
}

type AdapterKitService_ServerStreamingAdapterServer interface {
	Send(*AdapterResponse) error
	grpc.ServerStream
}

type adapterKitServiceServerStreamingAdapterServer struct {
	grpc.ServerStream
}

func (x *adapterKitServiceServerStreamingAdapterServer) Send(m *AdapterResponse) error {
	return x.ServerStream.SendMsg(m)
}

// AdapterKitService_ServiceDesc is the grpc.ServiceDesc for AdapterKitService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdapterKitService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "AdapterKitService",
	HandlerType: (*AdapterKitServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UniDirectionalAdapter",
			Handler:    _AdapterKitService_UniDirectionalAdapter_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "BiDirectionalAdapter",
			Handler:       _AdapterKitService_BiDirectionalAdapter_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "ServerStreamingAdapter",
			Handler:       _AdapterKitService_ServerStreamingAdapter_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "adapterkit.proto",
}
