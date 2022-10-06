package main

import (
	"flag"
	"fmt"
	"strconv"

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
		DefaultType: services.UniSs,
		FormatPlug: func(b []byte) ([]byte, services.GrpcType, error) {
			_, err := strconv.Atoi(string(b))
			if err == nil {
				return b, services.Ss, nil
			}

			return b, services.UniSs, nil
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
