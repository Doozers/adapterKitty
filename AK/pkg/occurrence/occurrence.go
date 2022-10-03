package occurrence

import (
	"context"

	"github.com/Doozers/adapterKitty/AK/proto"
)

func (s Srv) BiDirectionalAdapter(server proto.AdapterKitService_BiDirectionalAdapterServer) error {
	//TODO implement me
	panic("implement me")
}

func (s Srv) UniDirectionalAdapter(ctx context.Context, req *proto.AdapterRequest) (*proto.AdapterResponse, error) {
}

func (s Srv) ServerStreamingAdapter(req *proto.AdapterRequest, server proto.AdapterKitService_ServerStreamingAdapterServer) error {
	//TODO implement me
	panic("implement me")
}

type Srv struct {
	proto.UnimplementedAdapterKitServiceServer
}
