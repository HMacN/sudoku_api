package run_server

import (
	"sudoku_api/config/config_keys"
	"sudoku_api/services/logging"

	"github.com/spf13/cobra"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "as-server",
		Short: "Run as a server",
		Long:  "Run as a server",
		RunE: func(cmd *cobra.Command, args []string) error {
			loggingService, fxLogger := logging.NewLogger()
			loggerProvider := func() logging.LogWrapper { return loggingService }
			fx.New(
				fx.WithLogger(func() fxevent.Logger { return fxLogger }),
				fx.Provide(
					loggerProvider,
				),
				fx.Invoke(invoke),
			).Run()

			return nil
		},
	}
	cmd.PersistentFlags().IntP(config_keys.Port.String(), "p", 8080, "the port number for the server to listen on")
	return cmd
}
