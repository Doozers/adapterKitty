package services

import (
	"context"

	"github.com/Doozers/adapterKitty/AK-interfaces/go-console/proto"
)

type GrpcType int

const (
	Uni = 1
	Ss  = 2
	Bi  = 3
)

type Service interface {
	Format(msg []byte) (*proto.AdapterRequest, GrpcType, error)
	React(b []byte, a int32) (string, error)
	Listener(ctx context.Context, client proto.AdapterKitServiceClient)
}
