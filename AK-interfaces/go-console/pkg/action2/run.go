package action2

import (
	"fmt"

	"github.com/Doozers/adapterKitty/AK-interfaces/go-console/proto"
	p "github.com/golang/protobuf/proto"
)

func RunAction2(payload []byte) {
	data := &proto.Action2Payload{}
	err := p.Unmarshal(payload, data)
	if err != nil {
		fmt.Print("LOGS: SERV ANSWER >> Invalid data unmarshalling\n\n >> ")
		return
	}
	fmt.Print("LOGS: SERV ANSWER >> ", data, "\n\n >> ")
}
