package example

import (
	"fmt"
	"io"

	"github.com/Doozers/adapterKitty/AK/proto"
)

// ActionBi save request and retrieve it when payload is empty
func ActionBi(s proto.AdapterKitService_BiDirectionalAdapterServer) error {
	var keep [][]byte
	for {
		req, err := s.Recv()
		if err == io.EOF {
			for _, v := range keep {
				err := s.Send(&proto.AdapterResponse{Payload: []byte("Retrieve from keep: " + string(v))})
				if err != nil {
					fmt.Println("Error send: ", err)
					return err
				}
				fmt.Println("log : Retrieve from keep: ", string(v))
			}
			fmt.Println("log : All message retrieved")
			return nil
		}

		if err != nil {
			fmt.Println("Error: ", err)
			return err
		}

		fmt.Println("log : Received message: ", string(req.Payload))
		keep = append(keep, req.Payload)
	}
}
