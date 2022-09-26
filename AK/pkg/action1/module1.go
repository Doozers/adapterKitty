package action1

import (
	"fmt"
	"io"
	"strconv"

	"adapterKitty/proto"
)

func Mod(s proto.AdapterKitService_BiDirectionalAdapterServer) error {
	var num int
	for {
		req, err := s.Recv()
		if err == io.EOF {
			fmt.Println("Error: ", err)
			for i := 0; i < num; i++ {
				fmt.Println("Sending response: ", i)
				err := s.Send(&proto.AdapterResponse{Payload: []byte(fmt.Sprintf("Hello from action1: %d", i))})
				if err != nil {
					fmt.Println("Error send: ", err)
					return err
				}
			}
			s.Send(&proto.AdapterResponse{Payload: []byte(fmt.Sprintf("Hello from action1: %d", num))})
			fmt.Println("response sent")
			return nil
		}
		if err != nil {
			fmt.Println("Error: ", err)
			return err
		}
		fmt.Println(len(req.Payload))
		i, err := strconv.Atoi(string(req.Payload))
		if err != nil {
			fmt.Println("Error: enter a valid number")
			continue
		}
		num = i
	}
}
