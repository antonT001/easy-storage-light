package fileRepository

import "github.com/seivanov1986/sql_client"

//go:generate mockgen -source=repo.go -destination=repo_mock.go -package=fileRepository
type Repository interface {
	AddChunk(arg Chunk) error
	AllChunksUploadedForUUID(deliveryUUID string) (bool, error)
}

type repository struct {
	db sql_client.DataBase
}

func New(db sql_client.DataBase) Repository {
	return &repository{db: db}
}
