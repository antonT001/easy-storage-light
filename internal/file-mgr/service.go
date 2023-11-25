package filemgr

import (
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"

	"github.com/antonT001/easy-storage-light/internal/config"
)

//go:generate mockery --name FileMgr
type FileMgr interface {
	OpenFile(relativePath string) (*os.File, error)
	CreateDirectory(relativePath string) error
	CreateFile(relativePath string) (*os.File, error)
	CopyFile(dstRelativePath, srcRelativePath string) (err error)
	CopyAllFiles(dstRelativePath, srcRelativePath string) error
	GetListContentsInDirectory(relativePath string) ([]fs.DirEntry, error)
	DeleteDirectory(relativePath string) error
	DeleteFilesByPaths(relativePaths []string) error
	IsEquivalentSHA256Checksum(data []byte, hashSHA256 string) bool
}

type fileMgr struct {
	storageBasePath string
}

func NewService(cfg config.FileMgrConfig) (FileMgr, error) {
	err := os.MkdirAll(cfg.StorageBasePath, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("failed to create storage base directory [%s]: %v", cfg.StorageBasePath, err)
	}
	return &fileMgr{storageBasePath: cfg.StorageBasePath}, nil
}

func (fm fileMgr) OpenFile(relativePath string) (*os.File, error) {
	return os.Open(path.Join(fm.storageBasePath, relativePath))
}

func (fm fileMgr) CreateDirectory(relativePath string) error {
	return os.MkdirAll(path.Join(fm.storageBasePath, relativePath), os.ModePerm)
}

func (fm fileMgr) CreateFile(relativePath string) (*os.File, error) {
	dir := path.Dir(relativePath)
	if dir != "." {
		if err := fm.CreateDirectory(dir); err != nil {
			return nil, err
		}
	}
	return os.Create(path.Join(fm.storageBasePath, relativePath))
}

func (fm fileMgr) DeleteDirectory(relativePath string) error {
	return os.RemoveAll(path.Join(fm.storageBasePath, relativePath))
}

func (fm fileMgr) DeleteFilesByPaths(relativePaths []string) error {
	delErr := new(ErrorWithArguments)

	for i := range relativePaths {
		path := path.Join(fm.storageBasePath, relativePaths[i])
		if err := os.Remove(path); err != nil {
			delErr.OriginError = err
			delErr.Args = append(delErr.Args, path)
		}
	}

	if delErr.OriginError != nil {
		delErr.Message = "failed delete files"
		return delErr
	}

	return nil
}

func (fm fileMgr) GetListContentsInDirectory(relativePath string) ([]fs.DirEntry, error) {
	return os.ReadDir(path.Join(fm.storageBasePath, relativePath))
}

func (fm fileMgr) CopyFile(dstRelativePath, srcRelativePath string) (err error) {
	srcF, err := fm.OpenFile(srcRelativePath)
	if err != nil {
		return err
	}
	dstF, err := fm.CreateFile(dstRelativePath)
	if err != nil {
		srcF.Close()
		return err
	}
	_, err = io.Copy(dstF, srcF)
	srcF.Close()
	dstF.Close()
	return err
}

func (fm fileMgr) CopyAllFiles(dstRelativePath, srcRelativePath string) error {
	contents, err := fm.GetListContentsInDirectory(srcRelativePath)
	if err != nil {
		return err
	}

	err = fm.CreateDirectory(dstRelativePath)
	if err != nil {
		return err
	}

	copyAllErr := new(ErrorWithArguments)
	for i := range contents {
		if contents[i].IsDir() {
			continue
		}

		dst := path.Join(dstRelativePath, contents[i].Name())
		src := path.Join(srcRelativePath, contents[i].Name())
		if err := fm.CopyFile(dst, src); err != nil {
			copyAllErr.OriginError = err
			copyAllErr.Args = append(copyAllErr.Args, fmt.Sprintf("%s to %s", src, dst))
		}
	}
	if copyAllErr.OriginError != nil {
		copyAllErr.Message = "failed copy files"
		return copyAllErr
	}
	return nil
}

func (fm fileMgr) IsEquivalentSHA256Checksum(data []byte, hashSHA256 string) bool {
	return hashSHA256 == fm.SHA256Checksum(data)
}

type ErrorWithArguments struct {
	OriginError error
	Message     string
	Args        []string
}

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
