package commands

import (
	"context"
	"fmt"
	"os"

	"github.com/firstthumb/go-hue"
	"github.com/spf13/cobra"
)

func NewTurnOffLightCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "off",
		Short: "Turn off lights",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return fmt.Errorf("at least one light id is required, e.g.: `%s 1`", cmd.CommandPath())
			}

			return nil
		},
		Run: func(_ *cobra.Command, args []string) {
			err := runLightTurnOffCmd(args)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		},
	}
}

func runLightTurnOffCmd(args []string) error {
	client, err := setupClient()
	if err != nil {
		return fmt.Errorf("unable to setup Hue client: %w", err)
	}

	for _, arg := range args {
		status := false
		if _, _, err = client.Light.SetState(context.Background(), arg, &hue.SetStateRequest{On: &status}); err != nil {
			fmt.Fprintf(os.Stderr, "unable to turn on light %q: %v\n", arg, err)
			continue
		}
	}

	return nil
}
