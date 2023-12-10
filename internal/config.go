package internal

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env      string         `yaml:"env" env-default:"local"  env-required:"true"`
	DbConfig DatabaseConfig `yaml:"db"`
}

type DatabaseConfig struct {
	Driver string `yaml:"driver"`
	Path   string `yaml:"path"`
}

var cfg Config

func initConfig() Config {
	configPath := os.Getenv("CONFIG_PATH")
	
	if configPath == "" {
		panic("no config path specified")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return cfg
}

func ObtainConfig() Config {
	if (cfg == Config{}) {
		cfg = initConfig()
	}
	return cfg
}
