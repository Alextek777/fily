package config

import (
	"flag"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env string         `yaml:"env" env-default:"local"`
	Db  DatabaseConfig `yaml:"database"`
	Web WebConfig      `yaml:"web"`
}

type DatabaseConfig struct {
	Ip       string `yaml:"ip" env-required:"true"`
	Port     string `yaml:"port" env-required:"true"`
	User     string `yaml:"user" env-default:"postgres"`
	Password string `yaml:"password" env-default:"postgres"`
	SslMode  string `yaml:"sslMode" env-default:"disable"`
	DbName   string `yaml:"dbName" env-default:"fily"`
}

type WebConfig struct {
	Ip   string `yaml:"ip" env-required:"true"`
	Port string `yaml:"port" env-required:"true"`
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
