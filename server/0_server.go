package server

import (
	"net"

	"github.com/mnorbury/grpc-example/calculatorpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// calculatorService A struct used to implement the generated service interface
type calculatorService struct {

	// Embed this to allow partial implementations without the compiler complaining
	calculatorpb.UnimplementedCalculatorServiceServer
}

// StartServer starts a gRPC server and listens forever
func StartServer() error {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		return err
	}

	// Create a gRPC server
	server := grpc.NewServer()

	// Register our service
	calculatorpb.RegisterCalculatorServiceServer(server, &calculatorService{})

	// Register reflection service on gRPC server.
	reflection.Register(server)

	// Start serving
	return server.Serve(l)
}
