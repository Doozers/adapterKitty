package server

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/Doozers/adapterKitty/AK/pkg/announcement"
	"github.com/Doozers/adapterKitty/AK/proto"

	"google.golang.org/grpc"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
)

type Opts struct {
	ExposeWeb bool
	Addr      string
	GRPCPort  string
	HTTPPort  string
}

func Expose(opt Opts) error {
	/*return runGRPCServers(AdapterServ{
		BiAction:  example.ActionBi,
		UniAction: example.ActionUni,
		SsAction:  example.SsAction,
	}, opt)*/
	return runGRPCServers(&announcement.Srv{
		Mutx: &sync.Mutex{},
	}, opt)
}

func runGRPCServers(service proto.AdapterKitServiceServer, opts Opts) error {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s%s", opts.Addr, opts.GRPCPort))
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()

	proto.RegisterAdapterKitServiceServer(grpcServer, service)

	fmt.Println("Server started on: ", lis.Addr())

	// expose http server to serve grpc-web requests
	if opts.ExposeWeb {
		go func() {
			log.Fatalln(grpcServer.Serve(lis))
		}()

		grpcWebServer := grpcweb.WrapServer(
			grpcServer,
			grpcweb.WithOriginFunc(func(origin string) bool { return true }),
		)

		srv := &http.Server{
			Handler: grpcWebServer,
			Addr:    fmt.Sprintf("%s%s", opts.Addr, opts.HTTPPort),
		}

		return srv.ListenAndServe()
	} else {
		return grpcServer.Serve(lis)
	}
}
