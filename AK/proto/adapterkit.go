package proto

import (
	"context"
)

type AdapterServ struct {
	UnimplementedAdapterKitServiceServer

	BiMod  func(s AdapterKitService_BiDirectionalAdapterServer) error
	UniMod func(ctx context.Context, request *AdapterRequest) (*AdapterResponse, error)
}

func (a AdapterServ) BiDirectionalAdapter(server AdapterKitService_BiDirectionalAdapterServer) error {
	return a.BiMod(server)
}

func (a AdapterServ) UniDirectionalAdapter(ctx context.Context, request *AdapterRequest) (*AdapterResponse, error) {
	return a.UniMod(ctx, request)
}
