package config

import (
	"log"
	"path/filepath"
	"strings"

	"Babe-Piya/tamboo/appconfig"

	"github.com/spf13/viper"
)

func LoadConfig(configPath string) *appconfig.AppConfig {
	dir := filepath.Dir(configPath)
	fileBase := filepath.Base(configPath)
	filename := strings.TrimSuffix(fileBase, filepath.Ext(fileBase))

	viper.SetConfigName(filename)
	viper.AddConfigPath(dir)
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")

	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("fatal error config file: %+v", err)
	}

	var cfg appconfig.AppConfig
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("error unmarshaling config: %+v", err)
	}

	return &cfg
}
