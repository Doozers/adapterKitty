package utils

import (
	"github.com/Doozers/adapterKitty/AK/proto"
	pb "github.com/golang/protobuf/proto"
)

func IsProtoType(b []byte, t *proto.Test) bool {
	// TODO: try to make this with proto.reflect
	if pb.Unmarshal(b, t) == nil {
		return true
	}
	return false
}
