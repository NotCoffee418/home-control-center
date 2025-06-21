package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

var loadedConfig *Config = nil

func GetConfig() *Config {
	if loadedConfig == nil {
		var err error
		loadedConfig, err = loadConfig()
		if err != nil {
			log.Fatalf("Failed to load config: %v", err)
		}
	}
	return loadedConfig
}

func loadConfig() (*Config, error) {
	var config Config

	// Try dev config first (in parent directory from cmd/)
	devConfigPath := filepath.Join("..", "config-dev.toml")
	if _, err := os.Stat(devConfigPath); err == nil {
		_, err := toml.DecodeFile(devConfigPath, &config)
		return &config, err
	}

	// Fall back to production config
	prodConfigPath := "/etc/home-control-center/config.toml"
	_, err := toml.DecodeFile(prodConfigPath, &config)
	return &config, err
}
