package utils

import (
	"fmt"

	pb "github.com/golang/protobuf/proto"
)

func IsProtoType(b []byte, t pb.Message) pb.Message {
	// TODO: try to make this with proto.reflect
	fmt.Println(string(b))
	if pb.Unmarshal(b, t) != nil {
		return nil
	}

	return t
}
