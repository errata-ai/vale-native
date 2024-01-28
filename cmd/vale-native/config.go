package main

import (
	"encoding/json"
	"os"

	"github.com/adrg/xdg"
)

// CLIConfig is a struct that stores configuration information set by the Vale
// CLI during the installation process.
type CLIConfig struct {
	// The path to the Vale binary.
	Path string `json:"path"`
}

func newLogFile() (*os.File, error) {
	lf, err := xdg.ConfigFile("vale/native/host.log")
	if err != nil {
		return nil, err
	}
	return os.Create(lf)
}

func readConfig() (*CLIConfig, error) {
	cfg := &CLIConfig{}

	cf, err := xdg.ConfigFile("vale/native/config.json")
	if err != nil {
		return nil, err
	}

	b, err := os.ReadFile(cf)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
