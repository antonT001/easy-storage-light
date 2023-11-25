package config

import (
	"time"
)

type ServerConfig struct {
	Host         string        `yaml:"host"`
	Port         int           `yaml:"port"`
	IdleTimeout  time.Duration `yaml:"idle_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
}

func (s *ServerConfig) validate() error {
	if s.Host == "" {
		return errHostEmpty
	}
	if s.Port == 0 {
		return errPortEmpty
	}
	return nil
}
