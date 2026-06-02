package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sudoku",
	Short: "runs sudoku commands",
	Long:  "runs sudoku commands",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO: Does anything want to go in here?
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "An error while executing command '%s': error='%s'\n", rootCmd.Use, err)
		os.Exit(1)
	}
}
