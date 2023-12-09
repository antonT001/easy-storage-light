package filemgr

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/antonT001/easy-storage-light/internal/config"
)

type (
	fileMgrImpl struct {
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

func New(cfg config.FileMgrConfig) (FileMgr, error) {
	err := os.MkdirAll(cfg.StorageBasePath, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("failed to create storage base directory [%s]: %v", cfg.StorageBasePath, err)
	}
	return &fileMgrImpl{storageBasePath: cfg.StorageBasePath}, nil
}
