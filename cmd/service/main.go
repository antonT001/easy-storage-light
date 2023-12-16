package main

import (
	"log"

	"github.com/antonT001/easy-storage-light/internal/config"
	filemgr "github.com/antonT001/easy-storage-light/internal/file-mgr"
	fileRepository "github.com/antonT001/easy-storage-light/internal/repository/file"
	"github.com/antonT001/easy-storage-light/internal/rest"
	fileHandler "github.com/antonT001/easy-storage-light/internal/rest/file"
	fileService "github.com/antonT001/easy-storage-light/internal/service/file"
	"github.com/antonT001/easy-storage-light/migrations"
	"github.com/seivanov1986/sql_client/sqlite"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("failed to initialize configuration: %v", err)
	}

	fileMgr, err := filemgr.New(cfg.FileMgr)
	if err != nil {
		log.Fatalf("failed to initialize configuration fileMgr: %v", err)
	}

	db, err := sqlite.NewClient("easy-storage-light.db")
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}
	err = db.RunMigrations(log.Default(), migrations.MigrationFiles)
	if err != nil {
		log.Fatalf("failed database migrations: %v", err)
	}

	fileRepo := fileRepository.New(db)

	fileSvc := fileService.New(fileRepo, fileMgr)

	fHandler := fileHandler.New(fileSvc)

	app := rest.NewServer(cfg.Server, fHandler)

	log.Fatal(app.Run())
}
