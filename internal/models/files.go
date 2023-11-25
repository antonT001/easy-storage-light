package models

import "errors"

type File struct {
	UUID               string
	Name               string
	SHA256FileChecksum string
}

type UploadChunk struct {
	File
	ChunkNum            string
	SHA256ChunkChecksum string
}

func (u UploadChunk) Valdate() error {
	if u.UUID == "" {
		return errors.New("uuid field is empty")
	}
	if u.Name == "" {
		return errors.New("name field is empty")
	}
	if u.ChunkNum == "" {
		return errors.New("chunk_num field is empty")
	}
	if u.SHA256ChunkChecksum == "" {
		return errors.New("check_sum_chunk field is empty")
	}
	if u.SHA256FileChecksum == "" {
		return errors.New("check_sum_file field is empty")
	}
	return nil
}
