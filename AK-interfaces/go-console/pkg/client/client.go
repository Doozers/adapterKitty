package client

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/Doozers/adapterKitty/AK-interfaces/go-console/pkg/services"
	"github.com/Doozers/adapterKitty/AK-interfaces/go-console/proto"
)

type Opts struct {
	Addr     string
	GRPCPort string
}

// Connect to the server and chose the grpc service to use
func Connect(svc services.Service, opts Opts) error {
	conn, err := grpc.Dial(fmt.Sprintf("%s%s", opts.Addr, opts.GRPCPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	client := proto.NewAdapterKitServiceClient(conn)
	svc.Listener(context.Background(), client)

	if err != nil {
		return err
	}
	defer conn.Close()

	return nil
}
