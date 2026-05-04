package run_server

import (
	"errors"

	"github.com/spf13/cobra"
)

const (
	flagPort = "port"
)

func NewAsServerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Run as a server",
		Long:  "Run as a server",
		RunE: func(cmd *cobra.Command, args []string) error {
			if singleton == nil {
				return errors.New("service is not initialized")
			}

			return singleton.RunServer()
		},
	}
	cmd.PersistentFlags().IntP(flagPort, "p", 8080, "the port number for the server to listen on")
	return cmd
}
