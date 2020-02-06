package server

import (
	"io"

	"github.com/mnorbury/grpc-example/calculatorpb"
)

func (calculatorService) Max(stream calculatorpb.CalculatorService_MaxServer) error {
	var max int64

	for {
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		value := req.GetValue()
		if value > max {
			max = value
			if err := stream.Send(&calculatorpb.MaxResponse{Result: max}); err != nil {
				return err
			}
		}
	}
	return nil
}
