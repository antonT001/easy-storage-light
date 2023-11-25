package filemgr

import (
	"io/fs"
	"os"
	"path"
)

func (fm fileMgr) CreateDirectory(relativePath string) error {
	return os.MkdirAll(path.Join(fm.storageBasePath, relativePath), os.ModePerm)
}

func (fm fileMgr) DeleteDirectory(relativePath string) error {
	return os.RemoveAll(path.Join(fm.storageBasePath, relativePath))
}

func (fm fileMgr) GetListContentsInDirectory(relativePath string) ([]fs.DirEntry, error) {
	return os.ReadDir(path.Join(fm.storageBasePath, relativePath))
}
