package example

import (
	"context"

	"go.uber.org/zap"

	"github.com/Doozers/adapterKitty/AK/proto"
)

// ActionUni send back the request
func ActionUni(ctx context.Context, req *proto.AdapterRequest, logger *zap.Logger) (*proto.AdapterResponse, error) {
	return &proto.AdapterResponse{Payload: []byte("boomerang: " + string(req.Payload))}, nil
}
