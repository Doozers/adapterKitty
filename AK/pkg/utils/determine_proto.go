package utils

import (
	pb "github.com/golang/protobuf/proto"
)

func IsProtoType(b []byte, t pb.Message) pb.Message {
	// TODO: try to make this with proto.reflect
	if pb.Unmarshal(b, t) != nil {
		return nil
	}

	return t
}
