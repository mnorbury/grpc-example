package client

import (
	"context"
	"io"
	"math/rand"

	"github.com/mnorbury/grpc-example/calculatorpb"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func RunBidirectionalStreaming(cc *grpc.ClientConn) error {
	client := calculatorpb.NewCalculatorServiceClient(cc)
	stream, err := client.Max(context.Background())
	if err != nil {
		return err
	}

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			resp, rErr := stream.Recv()
			if rErr != nil {
				if rErr == io.EOF {
					break
				}
				log.Errorf("error receiving message: %v", err)
			}
			log.Infof("got response: %v", resp)
		}
	}()

	for i := 0; i < 10; i++ {
		req := &calculatorpb.MaxRequest{
			Value: rand.Int63n(100),
		}
		log.Infof("sending :%v", req)

		err := stream.Send(req)
		if err != nil {
			return err
		}
	}
	err = stream.CloseSend()
	if err != nil {
		return err
	}

	<-done
	return nil
}
