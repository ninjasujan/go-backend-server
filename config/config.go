package config

import (
	"log"
	"os"
	"path"
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

type Config struct {
	Server server `yaml:"server"`
}

func LoadConfig() *Config {
	file, err := os.ReadFile(path.Join("config", "env.yaml"))
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
			LoadConfig()
		}
	})
	return configStore
}
