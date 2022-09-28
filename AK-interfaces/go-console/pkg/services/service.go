package services

import (
	"context"

	"github.com/Doozers/adapterKitty/AK-interfaces/go-console/proto"
)

type GrpcType int

const (
	Uni = 1
	Bi  = 2
	Ss  = 3
)

type Service interface {
	Format(msg []byte) ([]byte, error)
	React(b []byte)
	BiListener(client proto.AdapterKitService_BiDirectionalAdapterClient)
	UniListener(ctx context.Context, client proto.AdapterKitServiceClient)
	SsListener(ctx context.Context, client proto.AdapterKitServiceClient)

	GetType() GrpcType
}
