package example

import (
	"errors"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use: "ExampleCommand",
		RunE: func(cmd *cobra.Command, args []string) error {
			if singleton == nil {
				return errors.New("service is not initialized")
			}

			singleton.Foo()
			return nil
		},
	}
}
