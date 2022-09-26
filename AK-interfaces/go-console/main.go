package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc/credentials/insecure"

	"google.golang.org/grpc"

	"github.com/Doozers/adapterKitty/AK-interfaces/go-console/proto"
)

type grpcType int

const (
	Uni = 1
	Bi  = 2
)

type service interface {
	react(b []byte)
	biListener(client proto.AdapterKitService_BiDirectionalAdapterClient)
	uniListener(ctx context.Context, client proto.AdapterKitServiceClient)

	getType() grpcType
}

func main() {
	o := &CLISvc{
		//Plugin: action2.RunAction2,
		Type: Bi,
	}
	/*o := &discordSvc{
		//Plugin: action2.RunAction2,
		Token: "OTc3Mjk3ODcxMjY1MjA2Mjky.GKipSp.J03_0MXFOWqQ6IU2iWu9B3wf1ALb6TyEuCuYvk",
	}*/
	if err := expose(o); err != nil {
		panic(err)
	}
	return
}

func expose(svc service) error {
	conn, err := grpc.Dial("127.0.0.1:9314", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	client := proto.NewAdapterKitServiceClient(conn)
	switch svc.getType() {
	case Bi:
		adapter, err := client.BiDirectionalAdapter(context.Background())
		if err = runBi(adapter, svc); err != nil {
			return err
		}
		break
	case Uni:
		svc.uniListener(context.Background(), client)
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

func runBi(client proto.AdapterKitService_BiDirectionalAdapterClient, opt service) error {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go opt.biListener(client)

	go func() {
		for {
			resp, err := client.Recv()
			if err == io.EOF {
				//fmt.Println("Error2: ", err)
				return
			}
			if err != nil {
				fmt.Println("Error2: ", err)
				return
			}
			opt.react(resp.Payload)
		}
	}()

	<-c
	return nil
}
