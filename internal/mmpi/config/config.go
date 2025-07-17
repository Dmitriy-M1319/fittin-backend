package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

var cfg *Config

func GetConfigInstance() Config {
	if cfg != nil {
		return *cfg
	}

	return Config{}
}

type Llm struct {
	ApiKey string `yaml:"apiKey"`
}

// Config - contains all configuration parameters in config package.
type Config struct {
	Llm Llm `yaml:"llm"`
}

// ReadConfigYML - read configurations from file and init instance Config.
func ReadConfigYML(configYML string) error {
	if cfg != nil {
		return nil
	}

	file, err := os.Open(configYML)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return err
	}

	return nil
}
