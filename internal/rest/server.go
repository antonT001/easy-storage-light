package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/antonT001/easy-storage-light/internal/config"
	"github.com/antonT001/easy-storage-light/internal/service"
)

type Server struct {
	svc service.Service
	App *http.Server
}

func NewServer(cfg config.ServerConfig, svc service.Service) *Server {
	app := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		IdleTimeout:  cfg.IdleTimeout,
		WriteTimeout: cfg.WriteTimeout,
		ReadTimeout:  cfg.ReadTimeout,
	}

	srv := &Server{
		svc: svc,
		App: app,
	}

	app.Handler = srv.initRoutes()
	return srv
}

func (s *Server) Run() error {
	return s.App.ListenAndServe()
}

func (s *Server) ping(w http.ResponseWriter, r *http.Request) {
	s.sendResponse(w, http.StatusOK, "Pong")
}

func (s *Server) sendResponse(w http.ResponseWriter, statusCode int, body any) {
	data, err := json.Marshal(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		msg := fmt.Sprintf("error sending request, fail marshal body: %v", err)
		_, err := w.Write([]byte(msg))
		if err != nil {
			log.Print(err) // TODO when the logger is added change to log.Errorf
		}
	}
	w.WriteHeader(statusCode)
	_, err = w.Write(data)
	if err != nil {
		log.Print(err) // TODO when the logger is added change to log.Errorf
	}
}
