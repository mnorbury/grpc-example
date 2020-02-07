package server

import (
	"context"

	"github.com/mnorbury/grpc-example/calculatorpb"
	log "github.com/sirupsen/logrus"
)

func (calculatorService) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	log.Infof("Got a Sum request: %v", req)
	defer log.Info("Finished Sum request")

	a := req.GetA()
	b := req.GetB()

	return &calculatorpb.SumResponse{Result: a + b}, nil
}
