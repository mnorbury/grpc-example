package server

import "github.com/mnorbury/grpc-example/calculatorpb"

func (calculatorService) Fibonacci(req *calculatorpb.FibonacciRequest, stream calculatorpb.CalculatorService_FibonacciServer) error {
	f := fib()

	max := req.GetValue()
	for {
		v := f()
		if v > max {
			break
		}
		if err := stream.Send(&calculatorpb.FibonacciResponse{Result: v}); err != nil {
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
