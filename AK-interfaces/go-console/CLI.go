package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/Doozers/adapterKitty/AK-interfaces/go-console/proto"
)

type CLISvc struct {
	Plugin func([]byte)
	Type   grpcType
}

func (svc *CLISvc) react(b []byte) {
	if svc.Plugin != nil {
		svc.Plugin(b)
		return
	}

	// default reaction
	fmt.Print("LOGS: SERV ANSWER >> ", string(b), "\n\n >> ")
}

func (svc *CLISvc) biListener(client proto.AdapterKitService_BiDirectionalAdapterClient) {
	var input string
	Reader := bufio.NewReader(os.Stdin)
	fmt.Print(" >> ")
	for {
		input, _ = Reader.ReadString('\n')
		fmt.Print(" >> ")

		if len(input) > 1 {
			if err := client.Send(&proto.AdapterRequest{Payload: []byte(input[:len(input)-1])}); err != nil {
				fmt.Println("Error1: ", err)
				return
			}
		} else {
			client.CloseSend()
			os.Exit(0)
		}
	}
}

func (svc *CLISvc) uniListener(ctx context.Context, client proto.AdapterKitServiceClient) {
	var input string
	Reader := bufio.NewReader(os.Stdin)
	fmt.Print(" >> ")
	for {
		input, _ = Reader.ReadString('\n')
		fmt.Print(" >> ")

		if len(input) > 0 {
			resp, err := client.UniDirectionalAdapter(ctx, &proto.AdapterRequest{Payload: []byte(input[:len(input)-1])})
			if err != nil {
				fmt.Println("Error1: ", err)
				return
			}
			svc.react(resp.Payload)
		}
	}
}

func (svc *CLISvc) getType() grpcType {
	return svc.Type
}
