package filemgr

import (
	"io/fs"
	"os"
	"path"

	"github.com/antonT001/easy-storage-light/internal/lib/httplib"
)

func (fm fileMgr) CreateDirectory(relativePath string) error {
	if err := os.MkdirAll(path.Join(fm.storageBasePath, relativePath), os.ModePerm); err != nil {
		restErr := httplib.NewError(err, httplib.GeneralFileMgrError, `"CreateDirectory"`)
		return restErr
	}
	return nil
}

func (fm fileMgr) DeleteDirectory(relativePath string) error {
	if err := os.RemoveAll(path.Join(fm.storageBasePath, relativePath)); err != nil {
		restErr := httplib.NewError(err, httplib.GeneralFileMgrError, `"DeleteDirectory"`)
		return restErr
	}
	return nil
}

func (fm fileMgr) GetListContentsInDirectory(relativePath string) ([]fs.DirEntry, error) {
	list, err := os.ReadDir(path.Join(fm.storageBasePath, relativePath))
	if err != nil {
		restErr := httplib.NewError(err, httplib.GeneralFileMgrError, `"GetListContentsInDirectory"`)
		return nil, restErr
	}
	return list, nil
}
