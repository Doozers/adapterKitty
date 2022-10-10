package example

import (
	"fmt"
	"strconv"

	"go.uber.org/zap"

	"github.com/Doozers/adapterKitty/AK/proto"
)

// SsAction server streaming action that take a number x and send x requests back
func SsAction(req *proto.AdapterRequest, server proto.AdapterKitService_ServerStreamingAdapterServer, logger *zap.Logger) error {
	num, err := strconv.Atoi(string(req.Payload))
	if err != nil {
		return err
	}
	for i := 0; i < num; i++ {
		err := server.Send(&proto.AdapterResponse{Payload: []byte(fmt.Sprintf("message: %d", i))})
		if err != nil {
			return err
		}
		logger.Info("Send message: ", zap.Int("index", i))
	}
	return nil
}
