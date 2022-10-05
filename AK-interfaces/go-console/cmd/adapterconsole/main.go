package main

import (
	"flag"
	"fmt"

	"github.com/Doozers/adapterKitty/AK-interfaces/go-console/pkg/client"
	"github.com/Doozers/adapterKitty/AK-interfaces/go-console/pkg/services"
)

var opts = client.Opts{}

func init() {
	flag.StringVar(&opts.Addr, "addr", "127.0.0.1", "Address to listen on")
	flag.StringVar(&opts.GRPCPort, "grpc", ":9314", "gRPC listen port")
	flag.Parse()
}

func main() {
	svc := &services.CLISvc{
		Type: services.Uni,
		FormatPlug: func(b []byte) ([]byte, error) {
			return b, nil
		},
		ReactPlug: func(bytes []byte) (string, error) {
			return fmt.Sprintf("bytes: %v\n", bytes), nil
		},
	}

	if err := client.Connect(svc, opts); err != nil {
		panic(err)
	}
	return
}
