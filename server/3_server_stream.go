package server

import (
	"github.com/mnorbury/grpc-example/calculatorpb"
	log "github.com/sirupsen/logrus"
)

func (calculatorService) Fibonacci(req *calculatorpb.FibonacciRequest, stream calculatorpb.CalculatorService_FibonacciServer) error {
	log.Infof("Got a Fibonacci request: %v", req)
	defer log.Info("Finished Fibonacci request")

	f := fib()

	max := req.GetValue()
	for {
		v := f()
		if v > max {
			break
		}
		resp := &calculatorpb.FibonacciResponse{Result: v}
		log.Infof("Sending response: %v", resp)
		if err := stream.Send(resp); err != nil {
			return err
		}
	}
	return nil
}

func fib() func() int64 {
	a, b := int64(0), int64(1)
	return func() int64 {
		a, b = b, a+b
		return a
	}
}
