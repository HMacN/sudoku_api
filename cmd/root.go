package cmd

import (
	"fmt"
	"os"
	"sudoku_api/app"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
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

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		app, err := app.NewApp()
		if err != nil {
			return errors.Wrap(err, "failed to initialise app")
		}

		return app.Run()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "An error while executing command '%s': error='%s'\n", rootCmd.Use, err)
		os.Exit(1)
	}
}
