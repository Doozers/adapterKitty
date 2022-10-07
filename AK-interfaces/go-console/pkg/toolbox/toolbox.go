package toolbox

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Doozers/adapterKitty/AK-interfaces/go-console/pkg/services"
	"github.com/Doozers/adapterKitty/AK-interfaces/go-console/pkg/utils"
	"github.com/Doozers/adapterKitty/AK-interfaces/go-console/proto"
	pb "github.com/golang/protobuf/proto"
)

func strToOp(s string) (proto.OperationSign, error) {
	switch s {
	case "ADD":
		return proto.Operation_PLUS, nil
	case "SUB":
		return proto.Operation_MINUS, nil
	case "MUL":
		return proto.Operation_MULTIPLY, nil
	case "DIV":
		return proto.Operation_DIVIDE, nil
	}
	return 0, fmt.Errorf("unknown operation")
}

func operation(args []string) (*proto.AdapterRequest, error) {
	nb1, err := strconv.Atoi(args[1])
	if err != nil {
		return nil, err
	}
	nb2, err := strconv.Atoi(args[2])
	if err != nil {
		return nil, err
	}

	sign, err := strToOp(args[0])
	if err != nil {
		return nil, err
	}

	scheme := &proto.Operation{
		Op: sign,
		A:  int32(nb1),
		B:  int32(nb2),
	}
	serialized, err := pb.Marshal(scheme)
	if err != nil {
		return nil, err
	}
	return &proto.AdapterRequest{
		Payload: serialized,
		Id:      int32(proto.ActionType_ACTION_OPERATION),
	}, nil
}

func ping(args []string) (*proto.AdapterRequest, services.GrpcType, error) {
	scheme := &proto.Ping{Message: strings.Join(args[2:], " ")}

	var grpcMethod services.GrpcType

	switch args[1] {
	case "REPEAT":
		grpcMethod = services.Ss
	case "ONCE":
		grpcMethod = services.Uni
	default:
		return nil, 0, fmt.Errorf("unknown ping method")
	}

	serialized, err := pb.Marshal(scheme)
	if err != nil {
		return nil, 0, err
	}

	return &proto.AdapterRequest{
		Payload: serialized,
		Id:      int32(proto.ActionType_ACTION_PING),
	}, grpcMethod, nil
}

func errorFunc(args []string) (*proto.AdapterRequest, error) {
	var errorType proto.ErrErrorType
	if len(args) > 1 && args[1] == "PANIC" {
		errorType = proto.Err_PANIC
	} else {
		errorType = proto.Err_ERROR
	}
	scheme := &proto.Err{Error: errorType}

	serialized, err := pb.Marshal(scheme)
	if err != nil {
		return nil, err
	}

	return &proto.AdapterRequest{
		Payload: serialized,
		Id:      int32(proto.ActionType_ACTION_ERROR),
	}, nil
}

func strFunc(args []string) (*proto.AdapterRequest, error) {
	if len(args) < 2 {
		return nil, nil
	}
	scheme := &proto.Str{Msg: strings.Join(args[1:], " ")}

	serialized, err := pb.Marshal(scheme)
	if err != nil {
		return nil, err
	}

	return &proto.AdapterRequest{
		Payload: serialized,
		Id:      int32(proto.ActionType_ACTION_STR),
	}, nil
}

func FormatToolbox(b []byte) (*proto.AdapterRequest, services.GrpcType, error) {
	args := strings.Split(strings.TrimSpace(string(b)), " ")

	switch args[0] {
	case "ADD", "SUB", "MUL", "DIV":
		if !utils.CheckArgs(args, &utils.CheckOpt{Min: 3, Max: 3}) {
			return nil, 0, fmt.Errorf("invalid args")
		}

		res, err := operation(args)
		if err != nil {
			return nil, 0, err
		}
		return res, services.Uni, nil

	case "PING":
		if !utils.CheckArgs(args, &utils.CheckOpt{Min: 3}) {
			return nil, 0, fmt.Errorf("invalid args")
		}

		res, grpcMethod, err := ping(args)
		if err != nil {
			return nil, 0, err
		}
		return res, grpcMethod, nil

	case "ERROR":
		if !utils.CheckArgs(args, &utils.CheckOpt{Min: 1}) {
			return nil, 0, fmt.Errorf("invalid args")
		}

		res, err := errorFunc(args)
		if err != nil {
			return nil, 0, err
		}
		return res, services.Uni, nil

	case "BI":
		if !utils.CheckArgs(args, &utils.CheckOpt{Min: 1}) {
			return nil, 0, fmt.Errorf("invalid args")
		}

		res, err := strFunc(args)
		if err != nil {
			return nil, 0, err
		}
		return res, services.Bi, nil

	case "DOUBLE":
		if !utils.CheckArgs(args, &utils.CheckOpt{Min: 2, Max: 2}) {
			return nil, 0, fmt.Errorf("invalid args")
		}

		res, err := operation([]string{"MUL", args[1], "2"})
		if err != nil {
			return nil, 0, err
		}
		return res, services.Uni, nil

	case "RAND":
		return &proto.AdapterRequest{Id: int32(proto.ActionType_ACTION_RANDOM)}, services.Uni, nil

	case "NORETURN":
		return &proto.AdapterRequest{Id: int32(proto.ActionType_ACTION_NORETURN)}, services.Uni, nil

	default:
		return &proto.AdapterRequest{Payload: b}, services.Uni, nil
	}

}

func ReactToolbox(b []byte, T int32) (string, error) {
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
