package services

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

func (svc *CLISvc) React(b []byte) {
	if svc.Plugin != nil {
		svc.Plugin(b)
		return
	}

	// default reaction
	fmt.Print("LOGS: SERV ANSWER >> ", string(b), "\n\n >> ")
}

func (svc *CLISvc) BiListener(client proto.AdapterKitService_BiDirectionalAdapterClient) {
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
		}
	}
}

func (svc *CLISvc) UniListener(ctx context.Context, client proto.AdapterKitServiceClient) {
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
			svc.React(resp.Payload)
		}
	}
}

func (svc *CLISvc) GetType() grpcType {
	return svc.Type
}
