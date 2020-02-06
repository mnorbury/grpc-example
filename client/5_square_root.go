package client

import (
	"context"

	"github.com/mnorbury/grpc-example/calculatorpb"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func RunSquareRoot(cc *grpc.ClientConn, value int64) error {
	client := calculatorpb.NewCalculatorServiceClient(cc)

	resp, err := client.SquareRoot(context.Background(), &calculatorpb.SquareRootRequest{Value: value})
	if err != nil {
		log.Errorf("error calculating square root: %v %T", err, err)
		if grpcErr, ok := status.FromError(err); ok {
			log.Fatalf("This was a grpc error with code: %v", grpcErr.Code())
		}
	}

	log.Infof("got a response of :%v", resp)

	return nil
}
