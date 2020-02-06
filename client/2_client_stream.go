package client

import (
	"context"

	"github.com/mnorbury/grpc-example/calculatorpb"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func RunClientStreaming(cc *grpc.ClientConn) error {
	client := calculatorpb.NewCalculatorServiceClient(cc)

	stream, err := client.SumStream(context.Background())
	if err != nil {
		return err
	}

	for i := 0; i < 10; i++ {
		err := stream.Send(&calculatorpb.SumStreamRequest{Value: int64(i)})
		if err != nil {
			return err
		}
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		return err
	}

	log.Infof("Got a response of :%v", resp)

	return nil
}
