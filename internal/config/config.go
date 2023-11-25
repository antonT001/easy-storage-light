package config

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

var (
	confidInvalid = "config is invalid"

	errHostEmpty            = errors.New("host field is empty")
	errPortEmpty            = errors.New("port field is empty")
	errStorageBasePathEmpty = errors.New("storage_base_path field is empty")
)

type Config struct {
	Server  ServerConfig  `yaml:"server"`
	FileMgr FileMgrConfig `yaml:"file_mgr"`
}

func NewConfig(path string) (*Config, error) {
	if path == "" {
		return nil, errors.New("empty path")
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failure to read config file: %s", path)
	}
	cfg := &Config{}
	if err = yaml.Unmarshal(data, cfg); err != nil {
		return nil, fmt.Errorf("failure to parse config: %v", err)
	}

	if err = cfg.validate(); err != nil {
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
