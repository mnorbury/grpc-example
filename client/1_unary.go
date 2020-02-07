package client

import (
	"context"

	"github.com/mnorbury/grpc-example/calculatorpb"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func RunUnaryClient(cc *grpc.ClientConn) error {
	client := calculatorpb.NewCalculatorServiceClient(cc)
	resp, err := client.Sum(context.Background(), &calculatorpb.SumRequest{A: 1, B: 2})
	if err != nil {
		return err
	}

	log.Infof("got response: %v", resp)
	return nil
}
