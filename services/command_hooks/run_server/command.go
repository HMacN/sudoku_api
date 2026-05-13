package run_server

import (
	"errors"
	"sudoku_api/services/config/config_keys"

	"github.com/spf13/cobra"
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
	cmd.PersistentFlags().IntP(config_keys.Port.String(), "p", 8080, "the port number for the server to listen on")
	return cmd
}
