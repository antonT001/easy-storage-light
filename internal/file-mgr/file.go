package filemgr

import (
	"io"
	"os"
	"path"
)

func (fm fileMgr) OpenFile(relativePath string) (*os.File, error) {
	return os.Open(path.Join(fm.storageBasePath, relativePath))
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

func (fm fileMgr) CopyFile(dst io.Writer, src io.Reader) error {
	_, err := io.Copy(dst, src)
	return err
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
