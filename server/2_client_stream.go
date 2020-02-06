package server

import (
	"io"

	"github.com/mnorbury/grpc-example/calculatorpb"
	log "github.com/sirupsen/logrus"
)

func (calculatorService) SumStream(stream calculatorpb.CalculatorService_SumStreamServer) error {
	log.Infof("Got a SumStream request: %v", stream)
	defer log.Info("Finished SumStream request")

	var result int64

	for {
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		log.Infof("Received request: %v", req)
		result += req.GetValue()
	}

	resp := &calculatorpb.SumStreamResponse{Result: result}
	log.Infof("Sending response: %v", resp)
	return stream.SendAndClose(resp)
}
