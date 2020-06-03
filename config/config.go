package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/jinzhu/configor"
)

// Config is a config :)
type Config struct {
	LogLevel     string `env:"LOG_LEVEL"`
	Addr         string `env:"ADDR"`
	PostgresHost string `env:"POSTGRES_HOST"`
	PostgresPort string `env:"POSTGRES_PORT"`
	PostgresDB   string `env:"POSTGRES_DB"`
	PostgresUser string `env:"POSTGRES_USER"`
	PostgresPass string `env:"POSTGRES_PASS"`
}

var (
	config Config
	once   sync.Once
)

// Get reads config from environment
func Get() *Config {
	once.Do(func() {
		envType := os.Getenv("SERVICE_ENV")
		if envType == "" {
			envType = "dev"
		}
		if err := configor.New(&configor.Config{Environment: envType}).Load(&config, "config/config.json"); err != nil {
			log.Fatal(err)
		}
		configBytes, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Configuration:", string(configBytes))
	})
	return &config
}
