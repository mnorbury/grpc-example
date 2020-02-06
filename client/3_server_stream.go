package client

import (
	"context"
	"io"

	"github.com/mnorbury/grpc-example/calculatorpb"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func RunServerStreaming(cc *grpc.ClientConn) error {
	client := calculatorpb.NewCalculatorServiceClient(cc)

	stream, err := client.Fibonacci(context.Background(), &calculatorpb.FibonacciRequest{Value: 10})
	if err != nil {
		return err
	}

	for {
		resp, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		log.Infof("Got a response of :%v", resp)
	}

	return nil
}
