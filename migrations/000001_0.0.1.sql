-- +goose Up
-- +goose StatementBegin

CREATE TABLE chunks 
(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	uuid varchar(36) NOT NULL,
	name varchar(255) NOT NULL,
	sha_256_file_checksum varchar(64) NOT NULL,
	sha_256_chunk_checksum varchar(64) NOT NULL,
	chunk_num varchar(19) NOT NULL,
  	total_chunks varchar(19) NOT NULL,
	created INTEGER NOT NULL
)

-- +goose StatementEnd
-- +goose Downs