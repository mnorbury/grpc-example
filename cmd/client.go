package cmd

import (
	"log"

	"github.com/mnorbury/grpc-example/client"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// clientCmd represents the server command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Start a client instance",
}

var unaryCmd = &cobra.Command{
	Use:   "unary",
	Short: "Unary client example",
	Run: func(cmd *cobra.Command, args []string) {
		runClient(client.RunUnaryClient)
	},
}

var clientStreamCmd = &cobra.Command{
	Use:   "client_stream",
	Short: "Client stream example",
	Run: func(cmd *cobra.Command, args []string) {
		runClient(client.RunClientStreaming)
	},
}

var serverStreamCmd = &cobra.Command{
	Use:   "server_stream",
	Short: "Server stream example",
	Run: func(cmd *cobra.Command, args []string) {
		runClient(client.RunServerStreaming)
	},
}

var bidirectionalStreamCmd = &cobra.Command{
	Use:   "bidirectional_stream",
	Short: "Bidirectional stream example",
	Run: func(cmd *cobra.Command, args []string) {
		runClient(client.RunBidirectionalStreaming)
	},
}

func runClient(fn func(conn *grpc.ClientConn) error) {
	cc, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("error opening client connection: %v", err)
	}
	defer func() { _ = cc.Close() }()

	err = fn(cc)
	if err != nil {
		log.Fatalf("error running client: %v", err)
	}
}

func init() {
	rootCmd.AddCommand(clientCmd)
	clientCmd.AddCommand(unaryCmd)
	clientCmd.AddCommand(clientStreamCmd)
	clientCmd.AddCommand(serverStreamCmd)
	clientCmd.AddCommand(bidirectionalStreamCmd)
}
