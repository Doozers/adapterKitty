package client

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"

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
	switch svc.GetType() {
	case services.Bi:
		adapter, err := client.BiDirectionalAdapter(context.Background())
		if err = runBi(adapter, svc); err != nil {
			return err
		}
		break
	case services.Ss, services.Uni:
		svc.UniSsListener(context.Background(), client)
		break
	default:
		return fmt.Errorf("unknown gRPC type")
	}

	if err != nil {
		return err
	}
	defer conn.Close()
	return nil
}

// runBi runs the bidirectional client-server service
func runBi(client proto.AdapterKitService_BiDirectionalAdapterClient, svc services.Service) error {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go svc.BiListener(client)

	go func(client proto.AdapterKitService_BiDirectionalAdapterClient, svc services.Service) {
		for {
			resp, err := client.Recv()
			if err == io.EOF {
				return
			}
			if err != nil {
				fmt.Println("Error2: ", err)
				return
			}
			svc.React(resp.Payload, proto.ActionType(resp.Id))
		}
	}(client, svc)

	<-c
	return nil
}
