package fileRepository

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/antonT001/easy-storage-light/internal/lib/httplib"
)

const (
	allFields = "uuid, name, sha_256_file_checksum, sha_256_chunk_checksum, chunk_num, total_chunks, created"
)

func (r repository) AddChunk(arg Chunk) error {
	query := `
	INSERT INTO chunks (` + allFields + `) 
		VALUES (:uuid, :name, :sha_256_file_checksum, :sha_256_chunk_checksum, :chunk_num, :total_chunks, :created);
	`
	_, err := r.db.NamedExecContext(context.TODO(), query, arg)
	if err != nil {
		log.Printf("db error: %v", err)
		return httplib.NewError(err, httplib.GeneralDBError, fmt.Sprintf("uuid: %s, chunk_num: %s", arg.UUID, arg.ChunkNum))
	}
	return nil
}

func (r repository) AllChunksUploadedForUUID(deliveryUUID string) (bool, error) {
	query := `SELECT ` + allFields + ` FROM chunks WHERE uuid = ?`
	var chunks []Chunk
	err := r.db.SelectContext(context.TODO(), &chunks, query, deliveryUUID)
	if err != nil {
		return false, httplib.NewError(err, httplib.GeneralDBError, fmt.Sprintf("uuid: %s", deliveryUUID))
	}

	loadedChunks := strconv.Itoa(len(chunks))
	if len(chunks) == 0 || loadedChunks != chunks[0].TotalChunks {
		log.Printf("loaded chunks: %v", loadedChunks)
		return false, nil
	}

	return true, nil
}
