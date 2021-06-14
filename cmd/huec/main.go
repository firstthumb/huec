package main

import (
	"fmt"
	"os"

	"github.com/firstthumb/huec/cmd/huec/commands"
	"github.com/spf13/cobra"
)

var (
	hueClientId     = ""
	hueClientSecret = ""
	hueAppId        = ""
)

func init() {
	os.Setenv("HUE_CLIENT_ID", hueClientId)
	os.Setenv("HUE_CLIENT_SECRET", hueClientSecret)
	os.Setenv("HUE_APP_ID", hueAppId)
}

func Huec() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "huec",
		Short: "huec controls a Philips Hue",
	}

	rootCmd.AddCommand(commands.NewVersionCmd())
	rootCmd.AddCommand(commands.NewInitCmd())

	lightsCmd := commands.NewLightsCmd()
	rootCmd.AddCommand(lightsCmd)

	lightsCmd.AddCommand(commands.NewListLightsCmd())
	lightsCmd.AddCommand(commands.NewTurnOnLightCmd())
	lightsCmd.AddCommand(commands.NewTurnOffLightCmd())

	return rootCmd
}

func main() {
	if err := Huec().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
