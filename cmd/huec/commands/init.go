package commands

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/firstthumb/go-hue"
	"github.com/firstthumb/huec/pkg/config"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

func NewInitCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Initializes huec",
		Args:  cobra.NoArgs,
		Run: func(*cobra.Command, []string) {
			err := runInitCmd()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		},
	}
}

func runInitCmd() error {
	cfgFile, err := config.AbsolutePath()
	if err != nil {
		return err
	}

	fmt.Println("Searching for a Hue bridge on your local network...")

	host, err := hue.Discover()
	if err != nil {
		return fmt.Errorf("unable to discover bridge on the network: %w", err)
	}

	fmt.Printf("Found => %s\n", host)

	client, err := hue.CreateUser(host, "huec_1", &hue.ClientOptions{Verbose: false})
	if err != nil {
		return fmt.Errorf("unable to create user on bridge: %w", err)
	}

	fmt.Printf("Client(%s) is created.\n", client.GetClientID())

	if err = saveConfig(cfgFile, &config.Config{Host: client.GetHost(), ClientID: client.GetClientID()}); err != nil {
		return fmt.Errorf("unable to save configuration: %w", err)
	}

	return nil
}

func saveConfig(path string, cfg *config.Config) error {
	if err := os.MkdirAll(filepath.Dir(path), 0700); err != nil {
		return fmt.Errorf("unable to create configuration directory: %w", err)
	}

	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return fmt.Errorf("unable to create configuration file: %w", err)
	}

	if err = yaml.NewEncoder(f).Encode(cfg); err != nil {
		return fmt.Errorf("unable to serialize configuration: %w", err)
	}

	return nil
}

func setupClient() (*hue.Client, error) {
	cfg, err := config.Read()
	if err != nil {
		if os.IsNotExist(err) {
			return nil, errors.New("huec is not initialized, please run `huec init` first")
		}

		return nil, err
	}

	client := hue.NewClient(cfg.Host, cfg.ClientID, &hue.ClientOptions{Verbose: false})

	return client, nil
}
