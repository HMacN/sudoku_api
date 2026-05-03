package cmd

import (
	"fmt"
	"net/http"
	"os"
	"sudoku_api/services/logging"
	"sudoku_api/services/server"
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

var rootCmd = &cobra.Command{
	Use:   "sudoku",
	Short: "runs sudoku commands",
	Long:  "runs sudoku commands",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		err := initialiseConfig(cmd)
		if err != nil {
			return errors.Wrap(err, "failed to initialise config")
		}

		err = initialiseLogger()
		if err != nil {
			return errors.Wrap(err, "failed to initialise logger")
		}

		loggingService, fxLogger := logging.NewLogger()
		loggerProvider := func() logging.LogWrapper { return loggingService }
		app = fx.New(
			fx.WithLogger(func() fxevent.Logger { return fxLogger }),
			fx.Provide(
				server.NewServer,
				server.NewServeMux,
				loggerProvider,
			),
			fx.Invoke(func(server *http.Server) {}),
		)

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		initErr := app.Err()
		if initErr != nil {
			return initErr
		}
		app.Run()
		return nil
	},
}

var exampleChildCommand = &cobra.Command{
	Use: "a",
	RunE: func(cmd *cobra.Command, args []string) error {
		logger.Info("start ChildCommand")
		time.Sleep(100 * time.Millisecond)
		logger.Info("end ChildCommand")
		return nil
	},
}

var exampleGrandchildCommand = &cobra.Command{
	Use: "b",
	RunE: func(cmd *cobra.Command, args []string) error {
		logger.Info("start GrandchildCommand")
		time.Sleep(100 * time.Millisecond)
		logger.Info("end GrandchildCommand")
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "An error while executing command '%s': error='%s'\n", rootCmd.Use, err)
		os.Exit(1)
	}
}
