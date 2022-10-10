package server

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"go.uber.org/zap"

	"github.com/Doozers/adapterKitty/AK/proto"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"

	"google.golang.org/grpc"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
)

type Opts struct {
	ExposeWeb bool
	Addr      string
	GRPCPort  string
	HTTPPort  string
	Verbose   bool
}

type Server interface {
	proto.AdapterKitServiceServer
	GetLogger() *zap.Logger
}

func RunGRPCServers(service Server, opts Opts) error {
	logger := service.GetLogger()
	grpcLogger := logger.Named("grpc")
	grpc_zap.ReplaceGrpcLoggerV2(grpcLogger)

	serverStreamOpts := []grpc.StreamServerInterceptor{
		grpc_zap.StreamServerInterceptor(grpcLogger),
	}
	serverUnaryOpts := []grpc.UnaryServerInterceptor{
		grpc_zap.UnaryServerInterceptor(grpcLogger),
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("%s%s", opts.Addr, opts.GRPCPort))
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(serverStreamOpts...)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(serverUnaryOpts...)),
	)

	proto.RegisterAdapterKitServiceServer(grpcServer, service)
	logger.Info("Server started on: ", zap.String("addr: ", lis.Addr().String()))

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

		logger.Info("gRPC web server started on: ", zap.String("addr: ", srv.Addr))
		return srv.ListenAndServe()
	} else {
		return grpcServer.Serve(lis)
	}
}
