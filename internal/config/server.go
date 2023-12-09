package config

import (
	"log"
	"os"
	"time"
)

type ServerConfig struct {
	Host         string
	Port         string
	IdleTimeout  time.Duration
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
}

func (s *ServerConfig) parseEnv() {
	idleTimeout, err := time.ParseDuration(os.Getenv("IDLE_TIMEOUT"))
	if err != nil {
		log.Printf("failed parse duration env.IDLE_TIMEOUT: %v", err) // TODO изменить на log.Errorf когда будет доступно
	}
	writeTimeout, err := time.ParseDuration(os.Getenv("WRITE_TIMEOUT"))
	if err != nil {
		log.Printf("failed parse duration env.WRITE_TIMEOUT: %v", err) // TODO изменить на log.Errorf когда будет доступно
	}
	readTimeout, err := time.ParseDuration(os.Getenv("READ_TIMEOUT"))
	if err != nil {
		log.Printf("failed parse duration env.READ_TIMEOUT: %v", err) // TODO изменить на log.Errorf когда будет доступно
	}

	s.Host = os.Getenv("HOST")
	s.Port = os.Getenv("PORT")
	s.IdleTimeout = idleTimeout
	s.WriteTimeout = writeTimeout
	s.ReadTimeout = readTimeout
}

func (s *ServerConfig) validate() error {
	if s.Host == "" {
		return errHostEmpty
	}
	if s.Port == "" {
		return errPortEmpty
	}
	return nil
}
