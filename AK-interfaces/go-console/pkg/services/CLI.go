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
	FormatPlug func([]byte) ([]byte, error)
	ReactPlug  func([]byte)
	Type       GrpcType
}

func (svc *CLISvc) Format(msg []byte) ([]byte, error) {
	if svc.FormatPlug != nil {
		res, err := svc.FormatPlug(msg)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
	return msg, nil
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
			res, err := svc.Format([]byte(input[:len(input)-1]))
			if err != nil {
				fmt.Println("Error1: ", err)
				return
			}
			if err := client.Send(&proto.AdapterRequest{Payload: res}); err != nil {
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
			res, err := svc.Format([]byte(input[:len(input)-1]))
			if err != nil {
				fmt.Println("Error1: ", err)
				return
			}
			resp, err := client.UniDirectionalAdapter(ctx, &proto.AdapterRequest{Payload: res})
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
			res, err := svc.Format([]byte(input[:len(input)-1]))
			if err != nil {
				fmt.Println("Error1: ", err)
				return
			}

			resp, err := client.ServerStreamingAdapter(ctx, &proto.AdapterRequest{Payload: res})
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
					fmt.Println("Error23: ", err)
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
