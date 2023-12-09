package rest

import (
	"net/http"

	"github.com/antonT001/easy-storage-light/internal/middlewares"
	"github.com/gorilla/mux"
)

func (s *Server) initRoutes() http.Handler {
	router := mux.NewRouter()

	v1 := router.PathPrefix("/api/v1/").Subrouter()
	v1.Use(middlewares.CorsMiddleware)
	v1.HandleFunc("/ping", s.ping).Methods(http.MethodGet)

	{ // files
		files := v1.PathPrefix("/files").Subrouter()
		files.HandleFunc("/upload", s.upload).Methods(http.MethodPost, http.MethodOptions)
	}

	return router
}
