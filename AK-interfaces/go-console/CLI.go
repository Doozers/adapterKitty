package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Doozers/adapterKitty/AK-interfaces/go-console/proto"
)

type CLISvc struct {
	Plugin func([]byte)
}

func (opt *CLISvc) react(b []byte) {
	if opt.Plugin != nil {
		opt.Plugin(b)
		return
	}

	// default reaction
	fmt.Print("LOGS: SERV ANSWER >> ", string(b), "\n\n >> ")
}

func (opt *CLISvc) listener(client proto.Serv_AdapterClient) {
	var input string
	Reader := bufio.NewReader(os.Stdin)
	fmt.Print(" >> ")
	for {
		input, _ = Reader.ReadString('\n')
		fmt.Print(" >> ")

		if len(input) > 0 {
			if err := client.Send(&proto.AdapterRequest{Payload: []byte(input[:len(input)-1])}); err != nil {
				fmt.Println("Error1: ", err)
				return
			}
		}
	}
}
