package main

import (
	"flag"
	"log"

	"go.uber.org/zap"
	"moul.io/zapconfig"

	"github.com/Doozers/adapterKitty/AK-interfaces/go-console/pkg/client"
	"github.com/Doozers/adapterKitty/AK-interfaces/go-console/pkg/services"
	"github.com/Doozers/adapterKitty/AK-interfaces/go-console/pkg/toolbox"
)

var opts = client.Opts{}

func init() {
	flag.StringVar(&opts.Addr, "addr", "127.0.0.1", "Address to listen on")
	flag.StringVar(&opts.GRPCPort, "grpc", ":9314", "gRPC listen port")
	flag.BoolVar(&opts.Verbose, "v", false, "Verbose mode")
	flag.Parse()
}

func main() {
	logger, err := newLogger(opts.Verbose)
	if err != nil {
		log.Fatalln(err)
	}

	svc := &services.CLISvc{
		FormatPlug: toolbox.FormatToolbox,
		ReactPlug:  toolbox.ReactToolbox,
		Logger:     logger,
	}

	if err := client.Connect(svc, opts); err != nil {
		panic(err)
	}
	return
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
