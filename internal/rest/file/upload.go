package filehandler

import (
	"errors"
	"log"
	"net/http"

	"github.com/antonT001/easy-storage-light/internal/lib/httplib"
	"github.com/antonT001/easy-storage-light/internal/models"
	fileservice "github.com/antonT001/easy-storage-light/internal/service/file"
)

type Handler interface {
	Upload(w http.ResponseWriter, r *http.Request)
}

type fileHandler struct {
	fileSrc fileservice.Service
}

func New(fileSrc fileservice.Service) Handler {
	return &fileHandler{fileSrc: fileSrc}
}

func (s *fileHandler) Upload(w http.ResponseWriter, r *http.Request) {
	upload := models.UploadChunk{
		File: models.File{
			UUID:               r.Header.Get(httplib.UUIDHeaderKey),
			Name:               r.Header.Get(httplib.NameHeaderKey),
			SHA256FileChecksum: r.Header.Get(httplib.SHA256FileChecksumHeaderKey),
		},
		ChunkNum:            r.Header.Get(httplib.ChunkNumHeaderKey),
		TotalChunks:         r.Header.Get(httplib.TotalChunksHeaderKey),
		SHA256ChunkChecksum: r.Header.Get(httplib.SHA256ChunkChecksumHeaderKey),
	}

	if err := upload.Validate(); err != nil {
		log.Printf("failed validate upload: %v", err) // TODO заменить на log.Errorf
		restErr := httplib.NewError(err, httplib.InvalidParam, "header")
		httplib.SendResponse(w, restErr.HTTPStatus, restErr)
		return
	}

	if r.ContentLength <= 0 {
		err := errors.New("empty request body")
		restErr := httplib.NewError(err, httplib.PayloadError, "body")
		httplib.SendResponse(w, restErr.HTTPStatus, restErr)
		return
	}

	err := s.fileSrc.UploadChunk(upload, r.Body)
	if restErr := httplib.AsRestErr(err); restErr != nil {
		httplib.SendResponse(w, restErr.HTTPStatus, restErr)
		return
	}

	httplib.SendResponse(w, http.StatusOK, nil)
}
