package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Server struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Postgres struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type Config struct {
	Server   Server   `yaml:"server"`
	Postgres Postgres `yaml:"postgres"`
}

func LoadConfig(path string) (Config, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("unable to load config file: %w", err)
	}

	var config Config
	if err := yaml.Unmarshal(file, &config); err != nil {
		return Config{}, fmt.Errorf("unable to unmarshal config file: %w", err)
	}

	return config, nil
}
