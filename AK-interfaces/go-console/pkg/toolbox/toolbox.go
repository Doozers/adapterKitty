package toolbox

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Doozers/adapterKitty/AK-interfaces/go-console/pkg/services"
	"github.com/Doozers/adapterKitty/AK-interfaces/go-console/proto"
	"github.com/Doozers/adapterKitty/AK/pkg/utils"
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

func FormatToolbox(b []byte) (*proto.AdapterRequest, services.GrpcType, error) {
	args := strings.Split(string(b), " ")
	if len(args) >= 3 {
		switch args[0] {
		case "ADD", "SUB", "MUL", "DIV":
			nb1, err := strconv.Atoi(args[1])
			if err != nil {
				return nil, 0, err
			}
			nb2, err := strconv.Atoi(args[2])
			if err != nil {
				return nil, 0, err
			}

			sign, err := strToOp(args[0])
			if err != nil {
				return nil, 0, err
			}

			scheme := &proto.Operation{
				Op: sign,
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

		case "PING":
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
	}

	return &proto.AdapterRequest{Payload: b}, services.Uni, nil
}

func ReactToolbox(b []byte) (string, error) {
	if len(b) == 0 {
		return "NO DATA", nil
	}
	if body := utils.IsProtoType(b, &proto.Result{}); body != nil {
		p := body.(*proto.Result)

		return fmt.Sprintf("RESULT: %d\n", p.Result), nil
	}
	return string(b), nil
}
