package config

import "os"

type FileMgrConfig struct {
	StorageBasePath string `yaml:"storage_base_path"`
}

func (fm *FileMgrConfig) parseEnv() {
	fm.StorageBasePath = os.Getenv("STORAGE_BASE_PATH")
}

func (fm *FileMgrConfig) validate() error {
	if fm.StorageBasePath == "" {
		return errStorageBasePathEmpty
	}
	return nil
}
