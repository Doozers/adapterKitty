package services

import (
	"context"

	"github.com/Doozers/adapterKitty/AK-interfaces/go-console/proto"
)

type grpcType int

const (
	Uni = 1
	Bi  = 2
)

type Service interface {
	React(b []byte)
	BiListener(client proto.AdapterKitService_BiDirectionalAdapterClient)
	UniListener(ctx context.Context, client proto.AdapterKitServiceClient)

	GetType() grpcType
}
