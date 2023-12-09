package filemgr

import (
	"io"
	"io/fs"
	"os"
)

//go:generate mockgen -source=service.go -destination=service_mock.go -package=filemgr
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
