package fileRepository

import (
	"time"

	"github.com/antonT001/easy-storage-light/internal/models"
)

type Chunk struct {
	UUID                string `db:"uuid"`
	Name                string `db:"name"`
	ChunkNum            string `db:"chunk_num"`
	TotalChunks         string `db:"total_chunks"`
	SHA256FileChecksum  string `db:"sha_256_file_checksum"`
	SHA256ChunkChecksum string `db:"sha_256_chunk_checksum"`
	Created             int64  `db:"created"`
}

func (c *Chunk) FromModel(in models.UploadChunk) {
	c.UUID = in.UUID
	c.Name = in.Name
	c.ChunkNum = in.ChunkNum
	c.TotalChunks = in.TotalChunks
	c.SHA256FileChecksum = in.SHA256FileChecksum
	c.SHA256ChunkChecksum = in.SHA256ChunkChecksum
	c.Created = time.Now().Unix()
}

func (c *Chunk) ToModel() models.UploadChunk {
	return models.UploadChunk{
		File: models.File{
			UUID:               c.UUID,
			Name:               c.Name,
			SHA256FileChecksum: c.SHA256FileChecksum,
		},
		ChunkNum:            c.ChunkNum,
		TotalChunks:         c.TotalChunks,
		SHA256ChunkChecksum: c.SHA256ChunkChecksum,
	}
}
