package commands

import (
	"fmt"

	hc "github.com/firstthumb/huec/context"
	"github.com/spf13/cobra"
)

func NewVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Prints the version",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Version:\t%v\n", hc.Version)
		},
	}
}
