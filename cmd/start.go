package cmd

import (
	"github.com/souravbiswassanto/prometheus/api"

	"github.com/spf13/cobra"
)

var Port string

// startCmd represents the start command
var (
	// Port stores port number for starting a connection
	startCmd = &cobra.Command{
		Use:   "start",
		Short: "start cmd starts the server on a port",
		Long: `It starts the server on a given port number, 
				Port number will be given in the cmd`,
		Run: func(cmd *cobra.Command, args []string) {
			api.RunServer(Port)
		},
	}
)

func init() {

	startCmd.PersistentFlags().StringVarP(&Port, "port", "p", "8081", "Port number for starting server")
	rootCmd.AddCommand(startCmd)
}

// https://github.com/yangshun/tech-interview-handbook
