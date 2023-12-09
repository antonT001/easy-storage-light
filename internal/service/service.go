package service

import (
	"fmt"
	"io"
	"path"

	filemgr "github.com/antonT001/easy-storage-light/internal/file-mgr"
	"github.com/antonT001/easy-storage-light/internal/models"
)

//go:generate mockgen -source=service.go -destination=service_mock.go -package=service
type Service interface {
	UploadChunk(upload *models.UploadChunk, body io.ReadCloser) error
}

type serviceImpl struct {
	fileMgr filemgr.FileMgr
}

func New(filemgr filemgr.FileMgr) Service {
	return &serviceImpl{
		fileMgr: filemgr,
	}
}

func (svc serviceImpl) UploadChunk(upload *models.UploadChunk, body io.ReadCloser) error {
	nameChunk := fmt.Sprintf("%s<chunk>%s", upload.ChunkNum, upload.Name)
	file, err := svc.fileMgr.CreateFile(path.Join(upload.UUID, nameChunk))
	if err != nil {
		return err
	}
	defer file.Close()

	return svc.fileMgr.CopyFile(file, body)
}
