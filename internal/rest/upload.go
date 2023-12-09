package rest

import (
	"errors"
	"net/http"

	"github.com/antonT001/easy-storage-light/internal/lib/httplib"
	"github.com/antonT001/easy-storage-light/internal/models"
)

func (s *Server) upload(w http.ResponseWriter, r *http.Request) {
	upload := new(models.UploadChunk)
	upload.UUID = r.Header.Get(httplib.UUIDHeaderKey)
	upload.Name = r.Header.Get(httplib.NameHeaderKey)
	upload.ChunkNum = r.Header.Get(httplib.ChunkNumHeaderKey)
	upload.SHA256FileChecksum = r.Header.Get(httplib.SHA256FileChecksumHeaderKey)
	upload.SHA256ChunkChecksum = r.Header.Get(httplib.SHA256ChunkChecksumHeaderKey)
	if err := upload.Valdate(); err != nil {
		restErr := httplib.NewError(err, httplib.InvalidParam, "header")
		s.sendResponse(w, restErr.HTTPStatus, restErr)
		return
	}

	if r.ContentLength <= 0 {
		err := errors.New("empty request body")
		restErr := httplib.NewError(err, httplib.PayloadError, "body")
		s.sendResponse(w, restErr.HTTPStatus, restErr)
		return
	}

	err := s.svc.UploadChunk(upload, r.Body)
	if restErr := httplib.AsRestErr(err); restErr != nil {
		s.sendResponse(w, restErr.HTTPStatus, restErr)
		return
	}

	s.sendResponse(w, http.StatusOK, nil)
}
