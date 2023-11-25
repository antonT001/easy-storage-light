package filemgr

import (
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"os"

	"github.com/antonT001/easy-storage-light/internal/config"
)

//go:generate mockery --name FileMgr
type FileMgr interface {
	OpenFile(relativePath string) (*os.File, error)
	CreateDirectory(relativePath string) error
	CreateFile(relativePath string) (*os.File, error)
	CopyFile(dst io.Writer, src io.Reader) error
	GetListContentsInDirectory(relativePath string) ([]fs.DirEntry, error)
	DeleteDirectory(relativePath string) error
	DeleteFilesByPaths(relativePaths []string) error
	IsEquivalentSHA256Checksum(data []byte, hashSHA256 string) bool
}

type (
	fileMgr struct {
		storageBasePath string
	}
	ErrorWithArguments struct {
		OriginError error
		Message     string
		Args        []string
	}
)

func (e *ErrorWithArguments) Error() string {
	msgByte, err := json.Marshal(e)
	if err != nil {
		return fmt.Sprintf(
			"failed unmarshal custom error, message: %v, origin_error: %v, args: %v",
			e.Message,
			e.OriginError,
			e.Args,
		)
	}
	return string(msgByte)
}

func NewService(cfg config.FileMgrConfig) (FileMgr, error) {
	err := os.MkdirAll(cfg.StorageBasePath, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("failed to create storage base directory [%s]: %v", cfg.StorageBasePath, err)
	}
	return &fileMgr{storageBasePath: cfg.StorageBasePath}, nil
}
