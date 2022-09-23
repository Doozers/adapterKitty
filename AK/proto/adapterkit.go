package proto

type AdapterServ struct {
	UnimplementedServServer

	Mod func(s Serv_AdapterServer) error
}

func (a AdapterServ) Adapter(s Serv_AdapterServer) error {
	return a.Mod(s)
}
