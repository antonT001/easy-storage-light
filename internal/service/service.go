package service

import (
	"fmt"
	"io"
	"path"

	filemgr "github.com/antonT001/easy-storage-light/internal/file-mgr"
	"github.com/antonT001/easy-storage-light/internal/models"
)

type Service interface {
	UploadChunk(upload *models.UploadChunk, body io.ReadCloser) error
}

type service struct {
	fileMgr filemgr.FileMgr
}

func NewService(
	filemgr filemgr.FileMgr,

) Service {
	return &service{
		fileMgr: filemgr,
	}
}

func (svc service) UploadChunk(upload *models.UploadChunk, body io.ReadCloser) (err error) {
	nameChunk := fmt.Sprintf("%s<chunk>%s", upload.ChunkNum, upload.Name)
	file, err := svc.fileMgr.CreateFile(path.Join(upload.UUID, nameChunk))
	if err != nil {
		return
	}
	defer file.Close()

	return svc.fileMgr.CopyFile(file, body)
}
