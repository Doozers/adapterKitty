package action2

import (
	"context"
	"strconv"

	"adapterKitty/proto"

	p "github.com/golang/protobuf/proto"
)

func ModUni(ctx context.Context, req *proto.AdapterRequest) (*proto.AdapterResponse, error) {
	tab := []struct {
		Name string
		Num  int
	}{
		{"one", 1111},
		{"two", 2222},
		{"three", 3333},
	}

	num, err := strconv.Atoi(string(req.Payload))
	typedPayload := &proto.Action2Payload{
		Name: "zero",
		Num:  0,
	}
	if err != nil || num > 3 || num < 1 {
		marshalledPayload, err := p.Marshal(typedPayload)
		if err != nil {
			return &proto.AdapterResponse{Payload: marshalledPayload}, err
		}
	}
	typedPayload.Name = tab[num-1].Name
	typedPayload.Num = int64(tab[num-1].Num)
	marshalledPayload, err := p.Marshal(typedPayload)
	if err != nil {
		return &proto.AdapterResponse{Payload: marshalledPayload}, err
	}
	return &proto.AdapterResponse{Payload: marshalledPayload}, nil
}
