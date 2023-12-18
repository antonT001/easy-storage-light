package tests

import (
	"testing"

	"github.com/antonT001/easy-storage-light/internal/config"
	"github.com/antonT001/easy-storage-light/internal/rest"
	filehandler "github.com/antonT001/easy-storage-light/internal/rest/file"
	"github.com/antonT001/easy-storage-light/internal/rest/tests/mocks"
	fileservice "github.com/antonT001/easy-storage-light/internal/service/file"

	"github.com/golang/mock/gomock"
)

func setup(t *testing.T, mock mockFn) *rest.Server {
	cfg := config.ServerConfig{Host: "localhost", Port: "8080"}

	ctrl := gomock.NewController(t)

	fileSvc := fileservice.NewMockService(ctrl)

	file := filehandler.New(fileSvc)

	testServer := rest.NewServer(cfg, file)

	if mock != nil {
		mock(&mocks.Mocks{
			FileService: fileSvc,
		})
	}

	return testServer
}
