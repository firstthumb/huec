package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Host         string `yaml:"host"`
	ClientID     string `yaml:"clientId"`
	AccessToken  string `yaml:"accessToken"`
	RefreshToken string `yaml:"refreshToken"`
	RedirectURL  string `yaml:"redirectUrl"`
	Expiry       int64  `yaml:"expiry"`
}

func Read() (*Config, error) {
	cfgPath, err := AbsolutePath()
	if err != nil {
		return nil, err
	}

	if _, err = os.Stat(cfgPath); os.IsNotExist(err) {
		return nil, err
	}

	f, err := os.Open(cfgPath)
	if err != nil {
		return nil, fmt.Errorf("unable to open config file: %w", err)
	}

	var cfg Config
	if err = yaml.NewDecoder(f).Decode(&cfg); err != nil {
		return nil, fmt.Errorf("unable to decode config file: %w", err)
	}

	return &cfg, nil
}

func AbsolutePath() (string, error) {
	cfgDir, err := os.UserConfigDir()
	if err != nil {
		return "", fmt.Errorf("user config dir not found: %w", err)
	}

	return filepath.Join(cfgDir, "huec", "config.yml"), nil
}
