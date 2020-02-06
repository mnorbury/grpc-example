package server

import (
	"context"

	"github.com/mnorbury/grpc-example/calculatorpb"
)

func (calculatorService) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	a := req.GetA()
	b := req.GetB()

	return &calculatorpb.SumResponse{Result: a + b}, nil
}
