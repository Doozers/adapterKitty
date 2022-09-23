package action2

import (
	"fmt"
	"io"
	"strconv"

	"adapterKitty/proto"

	p "github.com/golang/protobuf/proto"
)

func Mod(s proto.Serv_AdapterServer) error {
	tab := []struct {
		Name string
		Num  int
	}{
		{"one", 1111},
		{"two", 2222},
		{"three", 3333},
	}

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
		num, err := strconv.Atoi(string(req.Payload))
		typedPayload := &proto.Action2Payload{
			Name: "zero",
			Num:  0,
		}
		if err != nil || num > 3 || num < 1 {
			marshalledPayload, err := p.Marshal(typedPayload)
			if err != nil {
				return err
			}
			s.Send(&proto.AdapterResponse{Payload: marshalledPayload})
			continue
		}
		typedPayload.Name = tab[num-1].Name
		typedPayload.Num = int64(tab[num-1].Num)
		marshalledPayload, err := p.Marshal(typedPayload)
		if err != nil {
			return err
		}
		s.Send(&proto.AdapterResponse{Payload: marshalledPayload})
	}
}
