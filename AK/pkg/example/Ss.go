package example

import (
	"fmt"
	"strconv"

	"github.com/Doozers/adapterKitty/AK/proto"
)

func SsAction(req *proto.AdapterRequest, server proto.AdapterKitService_ServerStreamingAdapterServer) error {
	num, err := strconv.Atoi(string(req.Payload))
	if err != nil {
		return err
	}
	for i := 0; i < num; i++ {
		err := server.Send(&proto.AdapterResponse{Payload: []byte(fmt.Sprintf("message: %d", i))})
		if err != nil {
			return err
		}
		fmt.Println("log : Send message: ", i)
	}
	return nil
}
