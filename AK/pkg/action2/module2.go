package action2

import (
	"fmt"
	"io"
	"strconv"

	"adapterKitty/proto"
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
	list := []string{"1", "2", "3"}

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
		if err != nil {
			s.Send(&proto.AdapterResponse{Payload: []byte("Error: " + err.Error() + "\n")})
			return nil
		}
		if num > 3 || num < 1 {
			s.Send(&proto.AdapterResponse{Payload: []byte("Error: number must be 1, 2 or 3\n")})
			return nil
		}
		s.Send(&proto.AdapterResponse{Payload: []byte("You choose " + list[num-1] + " " + tab[num-1].Name + " " + strconv.Itoa(tab[num-1].Num) + "\n")})
	}
}
