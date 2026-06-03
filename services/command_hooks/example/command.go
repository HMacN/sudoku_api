package example

import (
	"sudoku_api/services/logging"

	"github.com/spf13/cobra"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "example",
		Short: "Example command",
		Long:  "Example command",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
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
}
