package example

import (
	"io"

	"go.uber.org/zap"

	"github.com/Doozers/adapterKitty/AK/proto"
)

// ActionBi save request and retrieve it when payload is empty
func ActionBi(s proto.AdapterKitService_BiDirectionalAdapterServer, logger *zap.Logger) error {
	var keep [][]byte
	for {
		req, err := s.Recv()
		if err == io.EOF {
			for _, v := range keep {
				err := s.Send(&proto.AdapterResponse{Payload: []byte("Retrieve from keep: " + string(v))})
				if err != nil {
					logger.Error("ActionBi: ", zap.Error(err))
					return err
				}
				logger.Info("Retrieve from keep: ", zap.ByteString("kept", v))
			}
			logger.Info("All message retrieved")
			return nil
		}

		if err != nil {
			logger.Error("ActionBi: ", zap.Error(err))
			return err
		}

		logger.Info("Received message: ", zap.ByteString("", req.Payload))
		keep = append(keep, req.Payload)
	}
}
