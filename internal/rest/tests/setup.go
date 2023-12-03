package tests

import (
	"testing"

	"github.com/antonT001/easy-storage-light/internal/config"
	"github.com/antonT001/easy-storage-light/internal/rest"
	"github.com/antonT001/easy-storage-light/internal/rest/tests/mocks"
	"github.com/antonT001/easy-storage-light/internal/service"
	"github.com/golang/mock/gomock"
)

func setup(t *testing.T, mock mockFn) *rest.Server {
	cfg := config.ServerConfig{Host: "localhost", Port: 8080}

	ctrl := gomock.NewController(t)

	svc := service.NewMockService(ctrl)
	testServer := rest.NewServer(cfg, svc)

	if mock != nil {
		mock(&mocks.Mocks{
			Service: svc,
		})
	}

	return testServer
}
