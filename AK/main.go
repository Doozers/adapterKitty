package AK

import (
	"fmt"
	"net"

	"adapterKitty/AK/server"
	"adapterKitty/proto"

	"google.golang.org/grpc"
)

func Srv() error {
	var opts []grpc.ServerOption
	lis, err := net.Listen("tcp", "127.0.0.1:9314")
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer(opts...)

	proto.RegisterServiceServer(grpcServer, &server.Adapter{})

	fmt.Println("Server started")
	if err := grpcServer.Serve(lis); err != nil {
		return err
	}
	return nil
}
