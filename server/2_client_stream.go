package server

import (
	"io"

	"github.com/mnorbury/grpc-example/calculatorpb"
)

func (calculatorService) SumStream(stream calculatorpb.CalculatorService_SumStreamServer) error {
	var result int64

	for {
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		result += req.GetValue()
	}

	return stream.SendAndClose(&calculatorpb.SumStreamResponse{Result: result})
}
