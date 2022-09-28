package server

import (
	"context"
	"fmt"

	"github.com/Doozers/adapterKitty/AK/proto"
)

type AdapterServ struct {
	proto.UnimplementedAdapterKitServiceServer

	BiAction  func(s proto.AdapterKitService_BiDirectionalAdapterServer) error
	UniAction func(ctx context.Context, request *proto.AdapterRequest) (*proto.AdapterResponse, error)
	SsAction  func(request *proto.AdapterRequest, server proto.AdapterKitService_ServerStreamingAdapterServer) error
}

func (a AdapterServ) BiDirectionalAdapter(server proto.AdapterKitService_BiDirectionalAdapterServer) error {
	if a.BiAction == nil {
		return fmt.Errorf("log : BiDirectionalAction is nil")
	}
	return a.BiAction(server)
}

func (a AdapterServ) UniDirectionalAdapter(ctx context.Context, request *proto.AdapterRequest) (*proto.AdapterResponse, error) {
	if a.UniAction == nil {
		return nil, fmt.Errorf("log : UniDirectionalAction is nil")
	}
	return a.UniAction(ctx, request)
}

func (a AdapterServ) ServerStreamingAdapter(request *proto.AdapterRequest, server proto.AdapterKitService_ServerStreamingAdapterServer) error {
	if a.SsAction == nil {
		return fmt.Errorf("log : ServerStreamingAction is nil")
	}
	return a.SsAction(request, server)
}
