package app

import (
	"fmt"
	"ocp-client/util/config"
	"os"
	"path/filepath"
	"strings"
)

func (app *App) initConfig() {
	const (
		configFileName    = "config.toml"
		defaultConfigPath = "./conf"
		configPathKey     = "SERVICE_CONFIG_FILE_PATH"
	)
	configPath := os.Getenv(configPathKey)
	if configPath == "" {
		configPath = defaultConfigPath
	}
	config.SetConfigPath(filepath.Join(configPath, configFileName))
	if err := config.ReadInConfig(); err != nil {
		fmt.Printf("config read error: %s", err)
	}
	config.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}
