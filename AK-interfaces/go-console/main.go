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

type Exposer int

const (
	CLI        = 1
	DiscordBot = 2
)

type service interface {
	react(b []byte)
	listener(client proto.Serv_AdapterClient)
}

func main() {
	/*o := &CLISvc{
		//Plugin: action2.RunAction2,
	}*/
	o := &discordSvc{
		//Plugin: action2.RunAction2,
		Token: "OTc3Mjk3ODcxMjY1MjA2Mjky.G7AP2L.LADYO_R5zc6MZOQ0E2asXY2yU7aT9F7c_DWp-0",
	}
	if err := expose(o); err != nil {
		panic(err)
	}
	return
}

func expose(o service) error {
	conn, err := grpc.Dial("127.0.0.1:9314", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	client := proto.NewServClient(conn)
	adapter, err := client.Adapter(context.Background())
	if err != nil {
		return err
	}
	if err := run(adapter, o); err != nil {
		return err
	}
	defer conn.Close()
	return nil
}

func run(client proto.Serv_AdapterClient, opt service) error {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go opt.listener(client)

	go func() {
		for {
			resp, err := client.Recv()
			if err == io.EOF {
				fmt.Println("Error2: ", err)
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
