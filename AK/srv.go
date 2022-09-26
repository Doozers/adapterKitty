package main

import (
	"context"
	"fmt"
	"net"

	"adapterKitty/pkg/example"
	"adapterKitty/proto"

	"google.golang.org/grpc"
)

type AdapterServ struct {
	proto.UnimplementedAdapterKitServiceServer

	BiAction  func(s proto.AdapterKitService_BiDirectionalAdapterServer) error
	UniAction func(ctx context.Context, request *proto.AdapterRequest) (*proto.AdapterResponse, error)
}

func (a AdapterServ) BiDirectionalAdapter(server proto.AdapterKitService_BiDirectionalAdapterServer) error {
	if a.BiAction == nil {
		return fmt.Errorf("log : BiDirectionalAction is nil")
	}
	return a.BiAction(server)
}

func (a AdapterServ) UniDirectionalAdapter(ctx context.Context, request *proto.AdapterRequest) (*proto.AdapterResponse, error) {
	if a.UniAction == nil {
		return nil, fmt.Errorf("log : UniDirectionalAction is nil")
	}
	return a.UniAction(ctx, request)
}

func Expose() error {
	return runGrpcServ(&AdapterServ{
		BiAction:  example.ActionBi,
		UniAction: example.ActionUni,
	})
}

func runGrpcServ(service *AdapterServ) error {
	var opts []grpc.ServerOption
	lis, err := net.Listen("tcp", "127.0.0.1:9314")
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer(opts...)

	proto.RegisterAdapterKitServiceServer(grpcServer, service)

	fmt.Println("Server started on: ", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		return err
	}
	return nil
}
