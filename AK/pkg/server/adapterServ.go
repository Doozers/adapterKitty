package server

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	"github.com/Doozers/adapterKitty/AK/proto"
)

type AdapterServ struct {
	proto.UnimplementedAdapterKitServiceServer

	BiAction  func(s proto.AdapterKitService_BiDirectionalAdapterServer, logger *zap.Logger) error
	UniAction func(ctx context.Context, request *proto.AdapterRequest, logger *zap.Logger) (*proto.AdapterResponse, error)
	SsAction  func(request *proto.AdapterRequest, server proto.AdapterKitService_ServerStreamingAdapterServer, logger *zap.Logger) error

	Logger *zap.Logger
}

func (a AdapterServ) BiDirectionalAdapter(server proto.AdapterKitService_BiDirectionalAdapterServer) error {
	if a.BiAction == nil {
		return fmt.Errorf("log : BiDirectionalAction is nil")
	}
	return a.BiAction(server, a.Logger)
}

func (a AdapterServ) UniDirectionalAdapter(ctx context.Context, request *proto.AdapterRequest) (*proto.AdapterResponse, error) {
	if a.UniAction == nil {
		return nil, fmt.Errorf("log : UniDirectionalAction is nil")
	}
	return a.UniAction(ctx, request, a.Logger)
}

func (a AdapterServ) ServerStreamingAdapter(request *proto.AdapterRequest, server proto.AdapterKitService_ServerStreamingAdapterServer) error {
	if a.SsAction == nil {
		return fmt.Errorf("log : ServerStreamingAction is nil")
	}
	return a.SsAction(request, server, a.Logger)
}

func (a AdapterServ) GetLogger() *zap.Logger {
	return a.Logger
}
