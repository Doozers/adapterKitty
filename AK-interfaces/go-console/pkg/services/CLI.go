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
	FormatPlug  func([]byte) (*proto.AdapterRequest, GrpcType, error)
	ReactPlug   func([]byte, proto.ActionType) (string, error)
	DefaultType GrpcType
}

func (svc *CLISvc) Format(msg []byte) (*proto.AdapterRequest, GrpcType, error) {
	if svc.FormatPlug != nil {
		res, t, err := svc.FormatPlug(msg)
		if err != nil {
			return nil, 0, err
		}
		return res, t, nil
	}

	return &proto.AdapterRequest{Payload: msg}, svc.DefaultType, nil
}

func (svc *CLISvc) React(b []byte, a proto.ActionType) (string, error) {
	if svc.ReactPlug != nil {
		res, err := svc.ReactPlug(b, a)
		if err != nil {
			return "", err
		}
		return res, nil
	}

	// default reaction
	return "LOGS: SERV ANSWER >> " + string(b) + "\n\n >> ", nil
}

// Deprecated
func (svc *CLISvc) BiListener(client proto.AdapterKitService_BiDirectionalAdapterClient) {
	var input string
	Reader := bufio.NewReader(os.Stdin)
	fmt.Print(" >> ")
	for {
		input, _ = Reader.ReadString('\n')
		fmt.Print(" >> ")

		if len(input) > 1 {
			res, _, err := svc.Format([]byte(input[:len(input)-1]))
			if err != nil {
				fmt.Println("Error1: ", err)
				return
			}
			if err := client.Send(res); err != nil {
				fmt.Println("Error1: ", err)
				return
			}
		} else {
			client.CloseSend()
		}
	}
}

func (svc *CLISvc) UniSsListener(ctx context.Context, client proto.AdapterKitServiceClient) {
	var input string
	Reader := bufio.NewReader(os.Stdin)
	fmt.Print(" >> ")
	for {
		input, _ = Reader.ReadString('\n')
		fmt.Print(" >> ")

		if len(input) > 0 {
			res, t, err := svc.Format([]byte(input[:len(input)-1]))
			if err != nil {
				fmt.Println("Error1: ", err)
				return
			}

			switch t {
			case Uni:
				resp, err := client.UniDirectionalAdapter(ctx, res)
				if err != nil {
					fmt.Println("Error1 Uni: ", err)
					return
				}
				fmt.Println(svc.React(resp.Payload, proto.ActionType(resp.Id)))
			case Ss:
				resp, err := client.ServerStreamingAdapter(ctx, res)
				if err != nil {
					fmt.Println("Error1 Ss: ", err)
					return
				}
				go func() {
					for {
						resp, err := resp.Recv()
						if err == io.EOF {
							break
						}
						if err != nil {
							fmt.Println("Error23: ", err)
							return
						}
						fmt.Println(svc.React(resp.Payload, proto.ActionType(resp.Id)))
					}
				}()
			}
		}
	}
}

/*func (svc *CLISvc) SsListener(ctx context.Context, client proto.AdapterKitServiceClient) {
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
				fmt.Println(svc.React(resp.Payload))
			}
		}
	}
}*/

func (svc *CLISvc) GetType() GrpcType {
	return svc.DefaultType
}
