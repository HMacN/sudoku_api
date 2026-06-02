package example

import (
	"errors"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "example",
		Short: "Example command",
		Long:  "Example command",
		RunE: func(cmd *cobra.Command, args []string) error {
			if singleton == nil {
				return errors.New("service is not initialised")
			}

			singleton.Foo()
			return nil
		},
	}
}
