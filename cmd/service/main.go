package main

import (
	"log"

	"github.com/antonT001/easy-storage-light/internal/config"
	filemgr "github.com/antonT001/easy-storage-light/internal/file-mgr"
	"github.com/antonT001/easy-storage-light/internal/rest"
	"github.com/antonT001/easy-storage-light/internal/service"
)

func main() {
	cfg, err := config.NewConfig("./config.yaml")
	if err != nil {
		log.Fatalf("failed to initialize configuration: %v", err)
	}

	fileMgr, err := filemgr.NewService(cfg.FileMgr)
	if err != nil {
		log.Fatalf("failed to initialize configuration fileMgr: %v", err)
	}

	svc := service.NewService(fileMgr)
	app := rest.NewServer(cfg.Server, svc)

	log.Fatal(app.Run())
}
