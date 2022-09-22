package server

import (
	"fmt"
	"io"

	"adapterKitty/proto"
)

type Adapter struct {
	proto.UnimplementedServiceServer
}

func (a Adapter) Adapter(s proto.Service_AdapterServer) error {
	i := 0
	for {
		req, err := s.Recv()
		if err == io.EOF {
			fmt.Println("Error: ", err)
			return nil
		}
		if err != nil {
			fmt.Println("Error: ", err)
			return err
		}
		i++
		fmt.Println(string(req.Payload))
		if i == 3 {
			fmt.Println("YA 3 MSG")
			if err := s.Send(&proto.AdapterResponse{Payload: []byte("Five messages\n")}); err != nil {
				return err
			}
			i = 0
		}
	}
}
