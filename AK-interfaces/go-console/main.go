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
	client := proto.NewServClient(conn)
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

func run(client proto.Serv_AdapterClient) error {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		var input string
		Reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Print(" >> ")
			input, _ = Reader.ReadString('\n')
			fmt.Print(input)

			if len(input) > 0 {
				if err := client.Send(&proto.AdapterRequest{Payload: []byte(input[:len(input)-1])}); err != nil {
					fmt.Println("Error1: ", err)
					return
				}
			}
			fmt.Println("---------------------------------------------------------------------------------------------------------------------")
		}
	}()

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
			fmt.Print("SERV ANSWER >> ", string(resp.Payload), "\n >> ")
		}
	}()

	<-c
	return nil
}
