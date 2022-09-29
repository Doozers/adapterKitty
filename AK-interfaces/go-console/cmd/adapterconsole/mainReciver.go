package main

import (
	"flag"
	"fmt"

	pb "google.golang.org/protobuf/proto"

	"github.com/Doozers/adapterKitty/AK-interfaces/go-console/pkg/client"
	"github.com/Doozers/adapterKitty/AK-interfaces/go-console/pkg/services"
	"github.com/Doozers/adapterKitty/AK-interfaces/go-console/proto"
)

var opts = client.Opts{}

func init() {
	flag.StringVar(&opts.Addr, "addr", "127.0.0.1", "Address to listen on")
	flag.StringVar(&opts.GRPCPort, "grpc", ":9314", "gRPC listen port")
	flag.Parse()
}

func main() {
	svc := &services.CLISvc{
		Type: services.Ss,
		FormatPlug: func(b []byte) ([]byte, error) {
			if string(b) == "connect" {
				ask := &proto.ConnectionRequest{AskToConnect: true}
				res, err := pb.Marshal(ask)
				if err != nil {
					return nil, err
				}
				return res, nil
			}
			ask := &proto.ConnectionRequest{AskToConnect: false}
			res, err := pb.Marshal(ask)
			if err != nil {
				return nil, err
			}
			return res, nil
		},
		ReactPlug: func(b []byte) {
			fmt.Println("msg: ", string(b))
		},
	}

	if err := client.Connect(svc, opts); err != nil {
		panic(err)
	}
	return
}
