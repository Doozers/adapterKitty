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
