package config

import (
	"errors"
	"fmt"
)

var (
	confidInvalid = "config is invalid"

	errHostEmpty            = errors.New("host field is empty")
	errPortEmpty            = errors.New("port field is empty")
	errStorageBasePathEmpty = errors.New("storage_base_path field is empty")
)

type Config struct {
	Server  ServerConfig
	FileMgr FileMgrConfig
}

func New() (*Config, error) {
	cfg := new(Config)
	cfg.Server.parseEnv()
	cfg.FileMgr.parseEnv()

	if err := cfg.validate(); err != nil {
		return nil, err
	}
	return cfg, nil
}

func (c *Config) validate() error {
	if err := c.Server.validate(); err != nil {
		return fmt.Errorf("%v: %v", confidInvalid, err)
	}
	if err := c.FileMgr.validate(); err != nil {
		return fmt.Errorf("%v: %v", confidInvalid, err)
	}
	return nil
}
