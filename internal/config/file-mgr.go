package config

type FileMgrConfig struct {
	StorageBasePath string `yaml:"storage_base_path"`
}

func (fm *FileMgrConfig) validate() error {
	if fm.StorageBasePath == "" {
		return errStorageBasePathEmpty
	}
	return nil
}
