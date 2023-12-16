package fileService

import (
	"fmt"
	"io"
	"log"
	"path"

	filemgr "github.com/antonT001/easy-storage-light/internal/file-mgr"
	"github.com/antonT001/easy-storage-light/internal/models"
	fileRepository "github.com/antonT001/easy-storage-light/internal/repository/file"
)

//go:generate mockgen -source=service.go -destination=service_mock.go -package=fileService
type Service interface {
	UploadChunk(upload models.UploadChunk, body io.ReadCloser) error
}

type serviceImpl struct {
	fileRepo fileRepository.Repository
	fileMgr  filemgr.FileMgr
}

func New(fileRepo fileRepository.Repository, filemgr filemgr.FileMgr) Service {
	return &serviceImpl{
		fileRepo: fileRepo,
		fileMgr:  filemgr,
	}
}

func (svc serviceImpl) UploadChunk(upload models.UploadChunk, body io.ReadCloser) error {
	nameChunk := fmt.Sprintf("%s<chunk>%s", upload.ChunkNum, upload.Name)
	file, err := svc.fileMgr.CreateFile(path.Join(upload.UUID, nameChunk))
	if err != nil {
		return err
	}
	defer file.Close()

	err = svc.fileMgr.CopyFile(file, body)
	if err != nil {
		return err
	}
	empty := fileRepository.Chunk{}
	empty.FromModel(upload)
	err = svc.fileRepo.AddChunk(empty) // TODO перед сохранением файла проверяем на его наличие
	if err != nil {
		return err
	}

	allUploaded, err := svc.fileRepo.AllChunksUploadedForUUID(empty.UUID)
	if err != nil {
		return err
	}

	if allUploaded {
		log.Printf("all chunks uploaded, run collect uuid: %v", empty.UUID)
		// run collect
	}

	return nil
}
