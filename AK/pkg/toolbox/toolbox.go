package toolbox

import (
	"context"
	"fmt"
	"io"
	"math/rand"
	"time"

	"go.uber.org/zap"

	"github.com/Doozers/adapterKitty/AK/proto"
	pb "github.com/golang/protobuf/proto"
)

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

func errorFunc(e proto.ErrErrorType) (*proto.AdapterResponse, error) {
	switch e {
	case proto.Err_PANIC:
		panic("motivated panic")

	default:
		return nil, fmt.Errorf("basic error (motivated)")
	}
}

func Uni(_ context.Context, req *proto.AdapterRequest, _ *zap.Logger) (*proto.AdapterResponse, error) {
	switch req.Id {
	case int32(proto.ActionType_ACTION_OPERATION):
		p := &proto.Operation{}
		if err := pb.Unmarshal(req.Payload, p); err != nil {
			return nil, err
		}

		scheme := &proto.Result{Result: calc(p.Op, p.A, p.B)}
		serialized, err := pb.Marshal(scheme)
		if err != nil {
			return nil, err
		}

		return &proto.AdapterResponse{
			Payload: serialized,
			Id:      int32(proto.ActionType_ACTION_RESULT),
		}, nil

	case int32(proto.ActionType_ACTION_PING):
		p := &proto.Ping{}
		if err := pb.Unmarshal(req.Payload, p); err != nil {
			return nil, err
		}

		return &proto.AdapterResponse{Payload: []byte(p.Message)}, nil

	case int32(proto.ActionType_ACTION_ERROR):
		p := &proto.Err{}
		if err := pb.Unmarshal(req.Payload, p); err != nil {
			return nil, err
		}

		return errorFunc(p.Error)

	case int32(proto.ActionType_ACTION_RANDOM):
		scheme := &proto.Result{Result: rand.Int31n(424242)}
		serialized, err := pb.Marshal(scheme)
		if err != nil {
			return nil, err
		}

		return &proto.AdapterResponse{
			Payload: serialized,
			Id:      int32(proto.ActionType_ACTION_RESULT),
		}, nil

	case int32(proto.ActionType_ACTION_NORETURN):
		return nil, nil

	default:
		return nil, fmt.Errorf("unknown request")
	}
}

func Ss(req *proto.AdapterRequest, server proto.AdapterKitService_ServerStreamingAdapterServer, _ *zap.Logger) error {
	switch req.Id {
	case int32(proto.ActionType_ACTION_PING):
		p := &proto.Ping{}
		if err := pb.Unmarshal(req.Payload, p); err != nil {
			return err
		}

		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			err := server.Send(&proto.AdapterResponse{Payload: []byte(p.Message)})
			if err != nil {
				return err
			}
		}
		return nil

	default:
		return fmt.Errorf("unknown request")
	}
}

func Bi(s proto.AdapterKitService_BiDirectionalAdapterServer, _gi *zap.Logger) error {
	var count int
	for {
		req, err := s.Recv()
		if err == io.EOF {
			fmt.Println("EOF")
			return nil
		}
		if err != nil {
			return err
		}

		switch req.Id {
		case int32(proto.ActionType_ACTION_STR):
			p := &proto.Str{}
			if err := pb.Unmarshal(req.Payload, p); err != nil {
				return err
			}

			err := s.Send(&proto.AdapterResponse{Payload: []byte(fmt.Sprintf("%s -> is the number %d", p.Msg, count))})
			if err != nil {
				return err
			}
			count++

		default:
			return fmt.Errorf("unknown request")
		}
	}
}
