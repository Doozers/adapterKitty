package toolbox

import (
	"fmt"
	"strconv"
	"strings"

	"go.uber.org/zap"

	"github.com/Doozers/adapterKitty/AK-interfaces/go-console/pkg/pipe"
	"github.com/Doozers/adapterKitty/AK-interfaces/go-console/pkg/services"
	"github.com/Doozers/adapterKitty/AK-interfaces/go-console/pkg/utils"
	"github.com/Doozers/adapterKitty/AK-interfaces/go-console/proto"
	pb "github.com/golang/protobuf/proto"
)

func operation(op proto.OperationSign) func(args []string) (*proto.AdapterRequest, services.GrpcType, error) {
	return func(args []string) (*proto.AdapterRequest, services.GrpcType, error) {
		nb1, err := strconv.Atoi(args[0])
		if err != nil {
			return nil, 0, err
		}
		nb2, err := strconv.Atoi(args[1])
		if err != nil {
			return nil, 0, err
		}

		scheme := &proto.Operation{
			Op: op,
			A:  int32(nb1),
			B:  int32(nb2),
		}
		serialized, err := pb.Marshal(scheme)
		if err != nil {
			return nil, 0, err
		}
		return &proto.AdapterRequest{
			Payload: serialized,
			Id:      int32(proto.ActionType_ACTION_OPERATION),
		}, services.Uni, nil
	}
}

func ping(method services.GrpcType) func([]string) (*proto.AdapterRequest, services.GrpcType, error) {
	return func(args []string) (*proto.AdapterRequest, services.GrpcType, error) {
		scheme := &proto.Ping{Message: strings.Join(args, " ")}

		serialized, err := pb.Marshal(scheme)
		if err != nil {
			return nil, 0, err
		}

		return &proto.AdapterRequest{
			Payload: serialized,
			Id:      int32(proto.ActionType_ACTION_PING),
		}, method, nil
	}
}

func errorFunc(errorType proto.ErrErrorType) func([]string) (*proto.AdapterRequest, services.GrpcType, error) {
	return func(args []string) (*proto.AdapterRequest, services.GrpcType, error) {
		scheme := &proto.Err{Error: errorType}

		serialized, err := pb.Marshal(scheme)
		if err != nil {
			return nil, 0, err
		}

		return &proto.AdapterRequest{
			Payload: serialized,
			Id:      int32(proto.ActionType_ACTION_ERROR),
		}, services.Uni, nil
	}
}

func strFunc(args []string) (*proto.AdapterRequest, services.GrpcType, error) {
	if len(args) < 1 {
		return nil, services.Bi, nil
	}
	scheme := &proto.Str{Msg: strings.Join(args, " ")}

	serialized, err := pb.Marshal(scheme)
	if err != nil {
		return nil, 0, err
	}

	return &proto.AdapterRequest{
		Payload: serialized,
		Id:      int32(proto.ActionType_ACTION_STR),
	}, services.Bi, nil
}

func definePipeline() *pipe.Pipeline {
	root := &pipe.Pipeline{}

	// operation
	ope := (&pipe.PipeWay{
		ID:    "OPE",
		Check: &utils.CheckOpt{Min: 3, Max: 3},
		Branch: []*pipe.Pipe{
			(&pipe.PipeEnd{ID: "ADD", F: operation(proto.Operation_PLUS)}).Pipe(),
			(&pipe.PipeEnd{ID: "SUB", F: operation(proto.Operation_MINUS)}).Pipe(),
			(&pipe.PipeEnd{ID: "MUL", F: operation(proto.Operation_MULTIPLY)}).Pipe(),
			(&pipe.PipeEnd{ID: "DIV", F: operation(proto.Operation_DIVIDE)}).Pipe(),
		}}).Pipe()

	// ping
	ping := (&pipe.PipeWay{
		ID:    "PING",
		Check: &utils.CheckOpt{Min: 2},
		Branch: []*pipe.Pipe{
			(&pipe.PipeEnd{ID: "REPEAT", F: ping(services.Ss)}).Pipe(),
			(&pipe.PipeEnd{ID: "ONCE", F: ping(services.Uni)}).Pipe(),
		}}).Pipe()

	// error
	errorPipe := (&pipe.PipeWay{
		ID:    "ERROR",
		Check: &utils.CheckOpt{Min: 1, Max: 1},
		Branch: []*pipe.Pipe{
			(&pipe.PipeEnd{ID: "BASIC", F: errorFunc(proto.Err_ERROR)}).Pipe(),
			(&pipe.PipeEnd{ID: "PANIC", F: errorFunc(proto.Err_PANIC)}).Pipe(),
		}}).Pipe()

	// BI
	bi := (&pipe.PipeEnd{ID: "BI", F: strFunc}).Pipe()

	// other
	other := []*pipe.Pipe{
		(&pipe.PipeEnd{ID: "DOUBLE", Check: &utils.CheckOpt{Min: 1, Max: 1}, F: func(args []string) (*proto.AdapterRequest, services.GrpcType, error) {
			nb, err := strconv.Atoi(args[0])
			if err != nil {
				return nil, 0, err
			}

			scheme := &proto.Operation{
				Op: proto.Operation_MULTIPLY,
				A:  int32(nb),
				B:  2,
			}
			serialized, err := pb.Marshal(scheme)
			if err != nil {
				return nil, 0, err
			}

			return &proto.AdapterRequest{
				Id:      int32(proto.ActionType_ACTION_OPERATION),
				Payload: serialized,
			}, services.Uni, nil
		}}).Pipe(),
		(&pipe.PipeEnd{ID: "RAND", Check: &utils.CheckOpt{Min: 0, Max: 0}, F: func(args []string) (*proto.AdapterRequest, services.GrpcType, error) {
			return &proto.AdapterRequest{Id: int32(proto.ActionType_ACTION_RANDOM)}, services.Uni, nil
		}}).Pipe(),
		(&pipe.PipeEnd{ID: "NORETURN", Check: &utils.CheckOpt{Min: 0, Max: 0}, F: func(args []string) (*proto.AdapterRequest, services.GrpcType, error) {
			return &proto.AdapterRequest{Id: int32(proto.ActionType_ACTION_NORETURN)}, services.Uni, nil
		}}).Pipe(),
	}

	root.Pipes = append(root.Pipes, ope)
	root.Pipes = append(root.Pipes, ping)
	root.Pipes = append(root.Pipes, errorPipe)
	root.Pipes = append(root.Pipes, bi)
	root.Pipes = append(root.Pipes, other...)

	return root
}

func FormatToolbox(b []byte, logger *zap.Logger) (*proto.AdapterRequest, services.GrpcType, error) {
	root := definePipeline()
	return root.Piping(string(b))
}

func ReactToolbox(b []byte, T int32, logger *zap.Logger) (string, error) {
	t := proto.ActionType(T)

	if len(b) == 0 {
		return "NO DATA", nil
	}
	if t == proto.ActionType_ACTION_RESULT {
		res := &proto.Result{}
		err := pb.Unmarshal(b, res)
		if err != nil {
			return "", err
		}

		return fmt.Sprintf("RESULT: %d\n", res.Result), nil
	}
	return string(b), nil
}
