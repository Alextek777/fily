package config

import (
	"flag"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env string `yaml:"env" env-default:"local"`
	Db  string `yaml:"database"`
}

type DatabaseConfig struct {
	Storage_ip       string `yaml:"storage_ip" env-required:"true"`
	Storage_port     string `yaml:"storage_port" env-required:"true"`
	Storage_user     string `yaml:"storage_user" env-default:"postgres"`
	Storage_password string `yaml:"storage_password env-default:"postgres""`
}

func MustLoad() *Config {
	path := fetchConfigPath()

	if path == "" {
		panic("config path is empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file does not exist: " + path)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read config: " + err.Error())
	}

	return &cfg
}

func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
