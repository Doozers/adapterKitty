package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc/credentials/insecure"

	"github.com/Doozers/adapterKitty/AK-interfaces/go-console/proto"

	"google.golang.org/grpc"
)

func main() {
	if err := exposeToCli(); err != nil {
		panic(err)
	}
	return
}

func exposeToCli() error {
	conn, err := grpc.Dial("127.0.0.1:9314", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	client := proto.NewServiceClient(conn)
	adapter, err := client.Adapter(context.Background())
	if err != nil {
		return err
	}
	if err := run(adapter); err != nil {
		return err
	}
	defer conn.Close()
	return nil
}

func run(client proto.Service_AdapterClient) error {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		var input string
		Reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Print(" >> ")
			input, _ = Reader.ReadString('\n')
			fmt.Print(input)

			if err := client.Send(&proto.AdapterRequest{Payload: []byte(input)}); err != nil {
				fmt.Println("Error: ", err)
				return
			}

			fmt.Println("---------------------------------------------------------------------------------------------------------------------")
		}
	}()

	go func() {
		for {
			resp, err := client.Recv()
			if err == io.EOF {
				return
			}
			if err != nil {
				return
			}
			fmt.Print("SERV ANSWER >> ", string(resp.Payload), "\n >> ")
		}
	}()

	<-c
	return nil
}
