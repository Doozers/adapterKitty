package main

import (
	"github.com/Doozers/adapterKitty/AK/pkg/example"
	"github.com/Doozers/adapterKitty/AK/pkg/server"

	"flag"
)

var opts = server.Opts{}

func init() {
	flag.BoolVar(&opts.ExposeWeb, "web", false, "Expose web interface")
	flag.StringVar(&opts.Addr, "addr", "127.0.0.1", "Address to listen on")
	flag.StringVar(&opts.GRPCPort, "grpc", ":9314", "gRPC listen port")
	flag.StringVar(&opts.HTTPPort, "http", ":9315", "HTTP listen port")
	flag.Parse()
}

func main() {
	if err := server.RunGRPCServers(&server.AdapterServ{
		BiAction:  example.ActionBi,
		UniAction: example.ActionUni,
		SsAction:  example.SsAction,
	}, opts); err != nil {
		panic(err)
	}
}
