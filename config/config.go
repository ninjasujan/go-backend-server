package config

import (
	"log"
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

var (
	configStore *Config
	once        sync.Once
)

type server struct {
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
	Server   server   `yaml:"server"`
	Postgres Postgres `yaml:"postgres"`
}

func LoadConfig(path string) *Config {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("[unable to load config file]:", err)
	}

	var config Config
	if err := yaml.Unmarshal(file, &config); err != nil {
		log.Fatal("[unable to unmarshal config file]:", err)
	}
	configStore = &config
	return configStore
}

func GetConfig() *Config {
	once.Do(func() {
		if configStore == nil {
			LoadConfig("local.yaml")
		}
	})
	return configStore
}
