package commands

import (
	"context"
	"fmt"
	"math"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

func NewLightsCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "lights",
		Aliases: []string{"light", "l"},
		Short:   "Manage Hue light bulbs",
		Args:    cobra.NoArgs,
		Run: func(*cobra.Command, []string) {
			err := runListLightsCmd()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		},
	}
}

func NewListLightsCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List available lights",
		Args:    cobra.NoArgs,
		Run:     func(*cobra.Command, []string) { runListLightsCmd() },
	}
}

func runListLightsCmd() error {
	client, err := setupClient()
	if err != nil {
		return fmt.Errorf("unable to setup Hue client: %w", err)
	}

	lights, _, err := client.Light.GetAll(context.Background())
	if err != nil {
		return fmt.Errorf("unable to list lights: %w", err)
	}

	tw := tabwriter.NewWriter(os.Stdout, 0, 4, 4, ' ', 0)
	defer tw.Flush()

	fmt.Fprintln(tw, "ID\tNAME\tON\tBRIGHTNESS (%)")

	for _, light := range lights {
		bri := math.Round(float64(light.GetBri()) / 254 * 100)

		fmt.Fprintf(tw, "%v\t%v\t%v\t%v\n", light.GetID(), light.GetName(), light.IsOn(), int(bri))
	}

	return nil
}
