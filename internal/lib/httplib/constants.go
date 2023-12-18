package httplib

const (
	UUIDHeaderKey                = "X-Project-Uuid"
	NameHeaderKey                = "X-Project-Name"
	ChunkNumHeaderKey            = "X-Project-Chunk-Num"
	TotalChunksHeaderKey         = "X-Project-Total-Chunks"
	SHA256FileChecksumHeaderKey  = "X-Project-Sha256-File-Checksum"
	SHA256ChunkChecksumHeaderKey = "X-Project-Sha256-Chunk-Checksum"

	AccessControlAllowOrigin  = "Access-Control-Allow-Origin"
	AccessControlAllowHeaders = "Access-Control-Allow-Headers"
	AccessControlAllowMethod  = "Access-Control-Allow-Method"
)
