package example

import (
	"context"
	"fmt"

	"github.com/Doozers/adapterKitty/AK/proto"
)

func ActionUni(ctx context.Context, req *proto.AdapterRequest) (*proto.AdapterResponse, error) {
	fmt.Println("log : ActionUni")
	return &proto.AdapterResponse{Payload: []byte("boomerang: " + string(req.Payload))}, nil
}
