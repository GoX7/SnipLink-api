package config

import (
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Path   Path   `yaml:"path"`
	Server Server `yaml:"server"`
	Log    Log    `yaml:"logger"`
}

type Path struct {
	Server string `yaml:"server_log"`
	Sqlite string `yaml:"sqlite_log"`
	MW     string `yaml:"mw_log"`
}

type Server struct {
	Addr string        `yaml:"addr"`
	Wto  time.Duration `yaml:"wto"`
	Rto  time.Duration `yaml:"rto"`
}

type Log struct {
	Level string `yaml:"level"`
}

func Load() (*Config, error) {
	var config Config

	err := cleanenv.ReadConfig("config/config.yaml", &config) //Loading and reading config - main/config/config.yaml
	if err != nil {
		return nil, err
	}

	return &config, nil
}
