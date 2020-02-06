package server

import (
	"io"

	"github.com/mnorbury/grpc-example/calculatorpb"
	log "github.com/sirupsen/logrus"
)

func (calculatorService) Max(stream calculatorpb.CalculatorService_MaxServer) error {
	log.Infof("Got a Max request: %v", stream)
	defer log.Info("Finished Max request")

	var max int64

	for {
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		log.Infof("Received request: %v", req)
		value := req.GetValue()
		if value > max {
			max = value
			resp := &calculatorpb.MaxResponse{Result: max}
			log.Infof("Sending response: %v", resp)
			if err := stream.Send(resp); err != nil {
				return err
			}
		}
	}
	return nil
}
