package services

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"

	"github.com/Doozers/adapterKitty/AK-interfaces/go-console/proto"
)

const DEFAULT_TYPE = Uni

type CLISvc struct {
	FormatPlug func([]byte) (*proto.AdapterRequest, GrpcType, error)
	ReactPlug  func([]byte, int32) (string, error)
}

func (svc *CLISvc) Format(msg []byte) (*proto.AdapterRequest, GrpcType, error) {
	if svc.FormatPlug != nil {
		res, t, err := svc.FormatPlug(msg)
		if err != nil {
			return nil, 0, err
		}
		return res, t, nil
	}

	// default format
	return &proto.AdapterRequest{Payload: msg}, DEFAULT_TYPE, nil
}

func (svc *CLISvc) React(b []byte, a int32) (string, error) {
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

func (svc *CLISvc) Listener(ctx context.Context, client proto.AdapterKitServiceClient) {
	var bi struct {
		running bool
		stream  proto.AdapterKitService_BiDirectionalAdapterClient
	}
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
				fmt.Println(svc.React(resp.Payload, resp.Id))

			case Ss:
				s, err := client.ServerStreamingAdapter(ctx, res)
				if err != nil {
					fmt.Println("Error1 Ss: ", err)
					return
				}
				for {
					resp, err := s.Recv()
					if err == io.EOF {
						break
					}
					if err != nil {
						fmt.Println("Error23: ", err)
						return
					}
					fmt.Println(svc.React(resp.Payload, resp.Id))
				}

			case Bi:
				if !bi.running {
					bi.stream, err = client.BiDirectionalAdapter(ctx)
					if err != nil {
						fmt.Println("Error1 Bi: ", err)
						return
					}

					go func() {
						for {
							resp, err := bi.stream.Recv()
							if err == io.EOF {
								bi.running = false
								bi.stream = nil
								break
							}
							if err != nil {
								fmt.Println("Error23: ", err)
								return
							}
							fmt.Println(svc.React(resp.Payload, resp.Id))
						}
					}()

					bi.running = true
				}

				if res != nil {
					if err := bi.stream.Send(res); err != nil {
						fmt.Println("Error1: ", err)
						return
					}
				} else {
					bi.stream.CloseSend()
				}
			}
		}
	}
}
