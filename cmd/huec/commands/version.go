package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version = "v1"

func NewVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Prints the version",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Version:\t%v\n", version)
		},
	}
}
