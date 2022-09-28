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

	mutx             *sync.Mutex
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

	for _, stream := range s.connectedClients {
		stream <- announce
	}
	return &proto.AdapterResponse{Payload: []byte("message sent")}, nil
}

func (s *Srv) ServerStreamingAdapter(req *proto.AdapterRequest, srv proto.AdapterKitService_ServerStreamingAdapterServer) error {
	ask := &proto.ConnectionRequest{}
	err := pb.Unmarshal(req.Payload, ask)
	if err != nil || !ask.AskToConnect {
		return fmt.Errorf("not a ask")
	}

	stream := make(chan *proto.Announce, 100)

	s.mutx.Lock()
	s.connectedClients = append(s.connectedClients, stream)
	s.mutx.Unlock()

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
