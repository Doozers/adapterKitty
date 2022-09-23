package main

import (
	"adapterKitty/pkg/action2"
	"adapterKitty/proto"
)

func main() {
	server := &proto.AdapterServ{
		Mod: action2.Mod,
	}
	if err := Srv(server); err != nil {
		panic(err)
	}
}
