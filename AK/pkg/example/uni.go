package example

import (
	"context"

	"adapterKitty/proto"
)

func ActionUni(ctx context.Context, req *proto.AdapterRequest) (*proto.AdapterResponse, error) {
	return &proto.AdapterResponse{Payload: []byte("boomerang: " + string(req.Payload))}, nil
}
