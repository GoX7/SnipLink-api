package config

import (
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Path   Path   `yaml:"path"`
	Server Server `yaml:"server"`
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

func Load() (*Config, error) {
	var config Config

	err := cleanenv.ReadConfig("config/config.yaml", &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
