package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	pb "google.golang.org/protobuf/proto"

	"github.com/Doozers/adapterKitty/AK-interfaces/go-console/pkg/client"
	"github.com/Doozers/adapterKitty/AK-interfaces/go-console/pkg/services"
	ff "github.com/peterbourgon/ff/v3"
	"github.com/peterbourgon/ff/v3/ffcli"

	"github.com/Doozers/adapterKitty/AK-interfaces/go-console/proto"
)

var opts = client.Opts{}

func main() {
	err := adapterconsole(os.Args[1:])
	if err != nil {
		log.Fatalf("err: %+v", err)
		os.Exit(1)
	}
}

func adapterconsole(args []string) error {
	rootFlagSet := flag.NewFlagSet("adapterconsole", flag.ExitOnError)
	rootFlagSet.StringVar(&opts.Addr, "addr", "127.0.0.1", "Address to listen on")
	rootFlagSet.StringVar(&opts.GRPCPort, "grpc", ":9314", "gRPC listen port")

	root := ffcli.Command{
		ShortUsage: "adapterconsole [flags] <command>",
		FlagSet:    rootFlagSet,
		Options:    []ff.Option{ff.WithEnvVarNoPrefix()},
		Subcommands: []*ffcli.Command{
			receiver(),
			sender(),
			discord(),
		},
		Exec: func(_ context.Context, _ []string) error {
			return flag.ErrHelp
		},
	}

	return root.ParseAndRun(context.Background(), args)
}

func receiver() *ffcli.Command {
	return &ffcli.Command{
		Name:       "receiver",
		ShortUsage: "adapterconsole receiver [flags]",
		ShortHelp:  "Starts the receiver",
		Options:    []ff.Option{ff.WithEnvVarNoPrefix()},
		Exec: func(_ context.Context, _ []string) error {
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

			return client.Connect(svc, opts)
		},
	}
}

func sender() *ffcli.Command {
	return &ffcli.Command{
		Name:       "sender",
		ShortUsage: "adapterconsole sender [flags]",
		ShortHelp:  "Starts the sender",
		Options:    []ff.Option{ff.WithEnvVarNoPrefix()},
		Exec: func(_ context.Context, _ []string) error {
			svc := &services.CLISvc{
				Type: services.Uni,
				FormatPlug: func(b []byte) ([]byte, error) {
					announce := &proto.Announce{Message: string(b)}
					res, err := pb.Marshal(announce)
					if err != nil {
						return nil, err
					}
					return res, nil
				},
			}

			return client.Connect(svc, opts)
		},
	}
}

func discord() *ffcli.Command {
	return &ffcli.Command{
		Name:       "discord",
		ShortUsage: "adapterconsole discord [flags]",
		ShortHelp:  "Starts the discord adapter",
		Options:    []ff.Option{ff.WithEnvVarNoPrefix()},
		Exec: func(_ context.Context, _ []string) error {
			svc := &services.DiscordSvc{
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
				Token: "OTc3Mjk3ODcxMjY1MjA2Mjky.Gdukbm.3jJNLAWF1BBHFopbbme-884oc50P-OrTErOB5g",
				Type:  services.Ss,
			}

			return client.Connect(svc, opts)
		},
	}
}
