package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	pb "google.golang.org/protobuf/proto"

	"github.com/Doozers/adapterKitty/AK-interfaces/go-console/pkg/client"
	"github.com/Doozers/adapterKitty/AK-interfaces/go-console/pkg/services"
	"github.com/Doozers/adapterKitty/AK-interfaces/go-console/proto"
	"github.com/peterbourgon/ff/v3"
	"github.com/peterbourgon/ff/v3/ffcli"
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
			cliWiki(),
		},
		Exec: func(_ context.Context, _ []string) error {
			return flag.ErrHelp
		},
	}

	return root.ParseAndRun(context.Background(), args)
}

func cliWiki() *ffcli.Command {
	return &ffcli.Command{
		Name:       "cliWiki",
		ShortUsage: "adapterconsole cliWiki [flags]",
		ShortHelp:  "Starts the client",
		Options:    []ff.Option{ff.WithEnvVarNoPrefix()},
		Exec: func(_ context.Context, _ []string) error {
			svc := &services.CLISvc{
				Type: services.Uni,
				FormatPlug: func(b []byte) ([]byte, error) {
					keyword, needle, ok := strings.Cut(strings.TrimSpace(string(b)), " ")
					if !ok {
						return nil, fmt.Errorf("bad format, expected \"{keyword} {needle}\"")
					}
					f := &proto.WikiRequest{
						Keyword: keyword,
						Needle:  needle,
					}
					b, err := pb.Marshal(f)
					if err != nil {
						return nil, err
					}
					return b, nil
				},
				ReactPlug: func(b []byte) {
					f := &proto.WikiResponse{}
					if err := pb.Unmarshal(b, f); err != nil {
						fmt.Println("bad format response")
						return
					}
					if f.GetError() {
						fmt.Println("internal server error")
					}
					fmt.Println("srv >> occurrence: ", f.Occurrence)
				},
			}

			return client.Connect(svc, opts)
		},
	}
}

func discordWiki() *ffcli.Command {
	var token string
	discordFS := flag.NewFlagSet("discordWiki", flag.ExitOnError)
	discordFS.StringVar(&token, "token", "0", "discord bot token")

	return &ffcli.Command{
		Name:       "discordWiki",
		ShortUsage: "adapterconsole discordWiki [flags]",
		ShortHelp:  "Starts the client",
		Options:    []ff.Option{ff.WithEnvVarNoPrefix()},
		FlagSet:    discordFS,
		Exec: func(_ context.Context, _ []string) error {
			svc := &services.DiscordSvc{
				FormatPlug: func(b []byte) ([]byte, error) {
					keyword, needle, ok := strings.Cut(strings.TrimSpace(string(b)), " ")
					if !ok {
						return nil, fmt.Errorf("bad format, expected \"{keyword} {needle}\"")
					}
					f := &proto.WikiRequest{
						Keyword: keyword,
						Needle:  needle,
					}
					b, err := pb.Marshal(f)
					if err != nil {
						return nil, err
					}
					return b, nil
				},
				ReactPlug: nil,
				Token:     token,
				Type:      services.Uni,
			}

			return client.Connect(svc, opts)
		},
	}
}
