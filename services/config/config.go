package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env            string `yaml:"env" env-default:"local"` //env-required:"true"
	RepositoryPath string `yaml:"repository_path" env-required:"true"`
	HTTPServer     `yaml:"http_server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8082"`
	TimeOut     time.Duration `yaml:"timeout" env-default:"5s"`
	IdleTimeOut time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

// load config
func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatalln("config_path is not set")
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config_path is not exist: %s", configPath)
	}
	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalln("cannot read config")
	}
	return &cfg
}
