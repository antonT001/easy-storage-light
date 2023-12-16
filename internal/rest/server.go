package rest

import (
	"fmt"
	"net/http"

	"github.com/antonT001/easy-storage-light/internal/config"
	"github.com/antonT001/easy-storage-light/internal/lib/httplib"
	fileHandler "github.com/antonT001/easy-storage-light/internal/rest/file"
)

type Server struct {
	fileHandler fileHandler.Handler
	App         *http.Server
}

func NewServer(cfg config.ServerConfig, file fileHandler.Handler) *Server {
	app := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		IdleTimeout:  cfg.IdleTimeout,
		WriteTimeout: cfg.WriteTimeout,
		ReadTimeout:  cfg.ReadTimeout,
	}

	srv := &Server{
		fileHandler: file,
		App:         app,
	}

	app.Handler = srv.initRoutes()
	return srv
}

func (s *Server) Run() error {
	return s.App.ListenAndServe()
}

func (s *Server) ping(w http.ResponseWriter, r *http.Request) {
	httplib.SendResponse(w, http.StatusOK, "Pong")
}
