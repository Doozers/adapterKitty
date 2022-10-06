package main

import (
	"flag"

	"go.uber.org/zap"
	"moul.io/zapconfig"

	"github.com/Doozers/adapterKitty/AK-interfaces/go-console/pkg/client"
	"github.com/Doozers/adapterKitty/AK-interfaces/go-console/pkg/services"
	"github.com/Doozers/adapterKitty/AK-interfaces/go-console/proto"
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
		FormatPlug: func(b []byte, logger *zap.Logger) (*proto.AdapterRequest, services.GrpcType, error) {
			if len(b) >= 3 && string(b[:3]) == "bi " {
				if string(b) == "bi " {
					return nil, services.Bi, nil
				}
				return &proto.AdapterRequest{Payload: b[3:]}, services.Bi, nil
			}

			_, err := strconv.Atoi(string(b))
			if err == nil {
				return &proto.AdapterRequest{Payload: b}, services.Ss, nil
			}

			return &proto.AdapterRequest{Payload: b}, services.Uni, nil
		},
		ReactPlug: func(b []byte, _ int32, logger *zap.Logger) (string, error) {
			return fmt.Sprintf("server sent: %s\n", b), nil
		},
		Logger: logger,
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
