package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/davecgh/go-spew/spew"
	"github.com/imdario/mergo"
	yaml "gopkg.in/yaml.v2"
)

// Config is a config :)
type Config struct {
	LogLevel             string `yaml:"log_level,omitempty"`
	Addr                 string `yaml:"addr,omitempty"`
	PostgresHost         string `yaml:"postgres_host,omitempty"`
	PostgresPort         string `yaml:"postgres_port,omitempty"`
	PostgresUser         string `yaml:"postgres_user,omitempty"`
	PostgresDB           string `yaml:"postgres_db,omitempty"`
	PostgresPass         string `yaml:"postgres_pass,omitempty"`
	PostgresMaxIdleConns int    `yaml:"postgres_max_idle_conns"`
	PostgresMaxOpenConns int    `yaml:"postgres_max_open_conns"`
	AllowedOrigins       string `yaml:"allowed_origins,omitempty"`
	Debug                bool   `yaml:"debug,omitempty"`
	CORS                 bool   `yaml:"cors"`
}

// CNFG is a Config singletone
var CNFG *Config

func init() {
	CNFG = loadConfig()
}

// Get returns config
func Get() *Config {
	return CNFG
}

// getEnv return env variable or default value provided
func getEnv(name, defaultVal string) string {
	val := os.Getenv(name)
	if val != "" {
		return val
	}
	return defaultVal
}

// loadConfig loads config from YAML files
func loadConfig() *Config {
	configPath := getEnv("CONFIG_PATH", "config")
	stage := getEnv("STAGE", "development")

	fmt.Printf("Config path: %s\n", configPath)
	fmt.Printf("Stage: %s\n", stage)

	yamlFileList := []string{}
	err := filepath.Walk(configPath, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error accessing a path %q: %v\n", configPath, err)
			return nil
		}
		if f.IsDir() {
			return nil
		}
		if filepath.Ext(path) != ".yaml" {
			return nil
		}
		fmt.Printf("Found YAML file %s\n", path)
		yamlFileList = append(yamlFileList, path)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	loadedConfigs := map[string]Config{}
	for _, yamlFilePath := range yamlFileList {
		fmt.Printf("Processing YAML file %s\n", yamlFilePath)
		yamlFileBytes, err := ioutil.ReadFile(yamlFilePath)
		fmt.Printf("%s contents:\n%s\n", yamlFilePath, yamlFileBytes)
		if err != nil {
			log.Fatal(err)
		}
		fileConfig := map[string]Config{}
		err = yaml.Unmarshal(yamlFileBytes, fileConfig)
		if err != nil {
			log.Fatal(err)
		}
		mergo.Merge(&loadedConfigs, fileConfig)
	}

	fmt.Println("Loaded configs:")
	spew.Dump(loadedConfigs)

	_, stageExists := loadedConfigs[stage]
	defaultConfig, defaultExists := loadedConfigs["defaults"]
	if !stageExists {
		fmt.Printf("Stage %s doesn't exist. Using default config", stage)
		if !defaultExists {
			panic(`No "defaults" config found`)
		}
		return &defaultConfig
	}
	CONFIG := defaultConfig
	mergo.Merge(&CONFIG, loadedConfigs[stage], mergo.WithOverride)

	return &CONFIG
}
