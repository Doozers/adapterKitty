package services

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"

	"github.com/Doozers/adapterKitty/AK-interfaces/go-console/proto"
)

type CLISvc struct {
	FormatPlug func([]byte) []byte
	ReactPlug  func([]byte)
	Type       GrpcType
}

func (svc *CLISvc) Format(msg []byte) []byte {
	if svc.FormatPlug != nil {
		return svc.FormatPlug(msg)
	}
	return msg
}

func (svc *CLISvc) React(b []byte) {
	if svc.ReactPlug != nil {
		svc.ReactPlug(b)
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
			if err := client.Send(&proto.AdapterRequest{Payload: svc.Format([]byte(input[:len(input)-1]))}); err != nil {
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
			resp, err := client.UniDirectionalAdapter(ctx, &proto.AdapterRequest{Payload: svc.Format([]byte(input[:len(input)-1]))})
			if err != nil {
				fmt.Println("Error1: ", err)
				return
			}
			svc.React(resp.Payload)
		}
	}
}

func (svc *CLISvc) SsListener(ctx context.Context, client proto.AdapterKitServiceClient) {
	var input string
	Reader := bufio.NewReader(os.Stdin)
	fmt.Print(" >> ")
	for {
		input, _ = Reader.ReadString('\n')
		fmt.Print(" >> ")

		if len(input) > 0 {
			resp, err := client.ServerStreamingAdapter(ctx, &proto.AdapterRequest{Payload: svc.Format([]byte(input[:len(input)-1]))})
			if err != nil {
				fmt.Println("Error1: ", err)
				return
			}
			for {
				resp, err := resp.Recv()
				if err == io.EOF {
					break
				}
				if err != nil {
					fmt.Println("Error2: ", err)
					return
				}
				svc.React(resp.Payload)
			}
		}
	}
}

func (svc *CLISvc) GetType() GrpcType {
	return svc.Type
}
