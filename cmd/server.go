package cmd

import (
	"log"

	"github.com/mnorbury/grpc-example/server"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start a server instance",
	Run: func(cmd *cobra.Command, args []string) {
		if err := server.StartServer(); err != nil {
			log.Fatalf("error starting server: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
