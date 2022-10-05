package announcement

import (
	"context"
	"fmt"
	"sync"

	pb "google.golang.org/protobuf/proto"

	"github.com/Doozers/adapterKitty/AK/proto"
)

type Srv struct {
	proto.UnimplementedAdapterKitServiceServer

	Mutx             *sync.Mutex
	connectedClients []chan *proto.Announce
}

func (s *Srv) BiDirectionalAdapter(server proto.AdapterKitService_BiDirectionalAdapterServer) error {
	//TODO implement me
	panic("implement me")
}

func (s *Srv) UniDirectionalAdapter(_ context.Context, req *proto.AdapterRequest) (*proto.AdapterResponse, error) {
	announce := &proto.Announce{}

	err := pb.Unmarshal(req.Payload, announce)
	if err != nil {
		return nil, err
	}

	err = s.broadcast(announce)
	if err != nil {
		return nil, err
	}
	return &proto.AdapterResponse{Payload: []byte(">>>log: message sent")}, nil
}

func (s *Srv) ServerStreamingAdapter(req *proto.AdapterRequest, srv proto.AdapterKitService_ServerStreamingAdapterServer) error {
	ask := &proto.ConnectionRequest{}
	err := pb.Unmarshal(req.Payload, ask)
	if err != nil || !ask.AskToConnect {
		return fmt.Errorf("not a ask")
	}

	stream := make(chan *proto.Announce, 100)

	s.Mutx.Lock()
	s.connectedClients = append(s.connectedClients, stream)
	s.Mutx.Unlock()

	err = s.broadcast(&proto.Announce{Message: "new client connected"})
	if err != nil {
		return err
	}

	for {
		res, err := pb.Marshal(<-stream)
		if err != nil {
			return err
		}
		if err = srv.Send(&proto.AdapterResponse{Payload: res}); err != nil {
			return err
		}
	}
}

func (s *Srv) broadcast(announce *proto.Announce) error {
	fmt.Printf("log: sent to %d peer\n", len(s.connectedClients))
	for _, stream := range s.connectedClients {
		stream <- announce
	}
	return nil
}
