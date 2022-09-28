package main

import (
	"flag"

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
		/*FormatPlug: func(b []byte) ([]byte, error) {
			announce := &proto.Announce{Message: string(b)}
			res, err := pb.Marshal(announce)
			if err != nil {
				return nil, err
			}
			return res, nil
		},*/
	}

	if err := client.Connect(svc, opts); err != nil {
		panic(err)
	}
	return
}
