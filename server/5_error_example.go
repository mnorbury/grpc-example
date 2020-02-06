package server

import (
	"context"
	"fmt"
	"math"

	"github.com/mnorbury/grpc-example/calculatorpb"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (calculatorService) SquareRoot(ctx context.Context, req *calculatorpb.SquareRootRequest) (*calculatorpb.SquareRootResponse, error) {
	log.Infof("Called SquareRoot: %v", req)
	defer log.Infof("Finished SquareRoot")

	value := req.GetValue()
	if value < 0 {
		return nil, status.Errorf(codes.InvalidArgument,
			fmt.Sprintf("negative numbers are not supported: received %v", value),
		)
	}

	result := math.Sqrt(float64(value))

	return &calculatorpb.SquareRootResponse{
		Result: result,
	}, nil
}
