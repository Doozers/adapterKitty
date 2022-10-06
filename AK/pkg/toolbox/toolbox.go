package toolbox

import (
	"context"
	"fmt"
	"time"

	"github.com/Doozers/adapterKitty/AK/proto"
	pb "github.com/golang/protobuf/proto"
)

func Uni(ctx context.Context, req *proto.AdapterRequest) (*proto.AdapterResponse, error) {
	/*
	 *
	 *  this could probably be done with a "switch-like" statement
	 * 	but I'm not sure how to do that with protobuf (yet)
	 *
	 *  switch req.Payload.(type) {
	 *  case *proto.Operation:
	 *  	// do something
	 *  case *proto.Ping:
	 *  	// do something
	 *  }
	 *
	 */
	if body := utils.IsProtoType(req.Payload, &proto.Ping{}); body != nil {
		p := body.(*proto.Ping)

		return &proto.AdapterResponse{Payload: []byte(p.Message)}, nil
	}
	if body := utils.IsProtoType(req.Payload, &proto.Operation{}); body != nil {
		p := body.(*proto.Operation)
		scheme := &proto.Result{Result: calc(p.Op, p.A, p.B)}
		serialized, err := pb.Marshal(scheme)
		if err != nil {
			return nil, err
		}

		return &proto.AdapterResponse{Payload: serialized}, nil
	}

	return nil, fmt.Errorf("unknown request")
}

func calc(sign proto.OperationSign, a int32, b int32) int32 {
	switch sign {
	case proto.Operation_PLUS:
		return a + b
	case proto.Operation_MINUS:
		return a - b
	case proto.Operation_MULTIPLY:
		return a * b
	case proto.Operation_DIVIDE:
		return a / b
	}
	return 0
}

func Ss(req *proto.AdapterRequest, server proto.AdapterKitService_ServerStreamingAdapterServer) error {
	/*
	 * same as above
	 */
	if body := utils.IsProtoType(req.Payload, &proto.Ping{}); body != nil {
		p := body.(*proto.Ping)
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			err := server.Send(&proto.AdapterResponse{Payload: []byte(p.Message)})
			if err != nil {
				return err
			}
		}
		return nil
	}

	return fmt.Errorf("unknown request")
}
