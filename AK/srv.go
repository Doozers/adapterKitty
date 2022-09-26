package main

import (
	"fmt"
	"net"

	"adapterKitty/pkg/action1"
	"adapterKitty/proto"

	"google.golang.org/grpc"
)

func Expose() error {
	return runGrpcServ(&proto.AdapterServ{BiMod: action1.Mod})
}

func runGrpcServ(service *proto.AdapterServ) error {
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
