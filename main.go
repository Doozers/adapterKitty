package main

import (
	"adapterKitty/AK"
)

func main() {
	if err := AK.Srv(); err != nil {
		panic(err)
	}
}
