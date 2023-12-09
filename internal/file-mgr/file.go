package filemgr

import (
	"io"
	"os"
	"path"

	"github.com/antonT001/easy-storage-light/internal/lib/httplib"
)

func (fm fileMgrImpl) OpenFile(relativePath string) (*os.File, error) {
	f, err := os.Open(path.Join(fm.storageBasePath, relativePath))
	if err != nil {
		restErr := httplib.NewError(err, httplib.GeneralFileMgrError, `"OpenFile"`)
		return nil, restErr
	}
	return f, nil
}

func (fm fileMgrImpl) CreateFile(relativePath string) (*os.File, error) {
	if err := fm.CreateDirectory(path.Dir(relativePath)); err != nil {
		return nil, err
	}

	f, err := os.Create(path.Join(fm.storageBasePath, relativePath))
	if err != nil {
		restErr := httplib.NewError(err, httplib.GeneralFileMgrError, `"CreateFile"`)
		return nil, restErr
	}
	return f, nil
}

func (fm fileMgrImpl) CopyFile(dst io.Writer, src io.Reader) error {
	if _, err := io.Copy(dst, src); err != nil {
		restErr := httplib.NewError(err, httplib.GeneralFileMgrError, `"CopyFile"`)
		return restErr
	}
	return nil
}

func (fm fileMgrImpl) DeleteFilesByPaths(relativePaths []string) error {
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
		restErr := httplib.NewError(delErr, httplib.GeneralFileMgrError, `"DeleteFilesByPaths"`)
		return restErr

	}
	return nil
}
