package commands

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/firstthumb/go-hue"
	"github.com/firstthumb/huec/pkg/config"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
	"gopkg.in/yaml.v2"
)

const (
	callbackURL = "http://localhost:8181/callback"
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
		fmt.Println("Unable to discover bridge on the network.")
		fmt.Println("Trying remote api...")

		return remoteInit(cfgFile)
	}

	fmt.Printf("Found => %s\n", host)

	client, err := hue.CreateUser(host, "huec_user", &hue.ClientOptions{})
	if err != nil {
		return fmt.Errorf("unable to create user on bridge: %w", err)
	}

	fmt.Printf("Client(%s) is created.\n", client.GetClientID())

	if err = saveConfig(cfgFile, &config.Config{
		Host:     client.GetHost(),
		ClientID: client.GetClientID(),
	}); err != nil {
		return fmt.Errorf("unable to save configuration: %w", err)
	}

	return nil
}

func remoteInit(cfgFile string) error {
	auth := hue.NewAuthenticator(callbackURL)
	client, err := auth.Authenticate()
	if err != nil {
		return fmt.Errorf("unable to login: %w", err)
	}

	_, err = client.CreateRemoteUser()
	if err != nil {
		return fmt.Errorf("unable to create user: %w", err)
	}

	fmt.Printf("Client(%s) is created.\n", client.GetClientID())

	token := auth.GetToken()

	if err = saveConfig(cfgFile, &config.Config{
		Host:         "",
		ClientID:     client.GetClientID(),
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		Expiry:       token.Expiry.Unix(),
		RedirectURL:  callbackURL,
	}); err != nil {
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

	if len(cfg.Host) == 0 {
		// Remote Api
		auth := hue.NewAuthenticator(cfg.RedirectURL)
		client := auth.NewClient(&oauth2.Token{
			AccessToken:  cfg.AccessToken,
			RefreshToken: cfg.RefreshToken,
			TokenType:    "Bearer",
			Expiry:       time.Unix(cfg.Expiry, 0),
		})
		client.Login(cfg.ClientID)

		return client, nil
	} else {
		// Local Api
		return hue.NewClient(cfg.Host, cfg.ClientID, &hue.ClientOptions{}), nil
	}
}
