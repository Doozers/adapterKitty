package main

import (
	"flag"
	"log"

	"go.uber.org/zap"
	"moul.io/zapconfig"

	"github.com/Doozers/adapterKitty/AK/pkg/server"
	"github.com/Doozers/adapterKitty/AK/pkg/toolbox"
)

var opts = server.Opts{}

func init() {
	flag.BoolVar(&opts.ExposeWeb, "web", false, "Expose web interface")
	flag.StringVar(&opts.Addr, "addr", "127.0.0.1", "Address to listen on")
	flag.StringVar(&opts.GRPCPort, "grpc", ":9314", "gRPC listen port")
	flag.StringVar(&opts.HTTPPort, "http", ":9315", "HTTP listen port")
	flag.BoolVar(&opts.Verbose, "v", false, "Verbose mode")
	flag.Parse()
}

func main() {
	logger, err := newLogger(opts.Verbose)
	if err != nil {
		log.Fatalln(err)
	}

	if err := server.RunGRPCServers(&server.AdapterServ{
		UniAction: toolbox.Uni,
		SsAction:  toolbox.Ss,
		BiAction:  toolbox.Bi,
		Logger:    logger,
	}, opts); err != nil {
		panic(err)
	}
}

func newLogger(verbose bool) (*zap.Logger, error) {
	config := zapconfig.Configurator{}

	if verbose {
		config.SetLevel(zap.DebugLevel)
	} else {
		config.SetLevel(zap.InfoLevel)
	}

	return config.Build()
}
