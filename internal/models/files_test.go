package models

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUploadChunk_Valdate(t *testing.T) {

	tests := []struct {
		name     string
		model    UploadChunk
		expected error
	}{
		{
			name: "valid model",
			model: UploadChunk{
				File: File{
					UUID:               "39c4739c-91cf-11ee-b9d1-0242ac120002",
					Name:               "gopher.jpg",
					SHA256FileChecksum: "cd372fb85148700fa88095e3492d3f9f5beb43e555e5ff26d95f5a6adc36f8e6",
				},
				ChunkNum:            "0",
				TotalChunks:         "11",
				SHA256ChunkChecksum: "cbb756eb255316279a3e09cb7342c38754060a5b4bd6560e14f51d85cbd745e6",
			},
			expected: nil,
		},
		{
			name: "invalid UUID",
			model: UploadChunk{
				File: File{
					UUID: "39c4739c-91cf-11 ",
				},
			},
			expected: errors.New("uuid field is invalid, parse error"),
		},
		{
			name: "invalid name (length)",
			model: UploadChunk{
				File: File{
					UUID: "39c4739c-91cf-11ee-b9d1-0242ac120002",
					Name: "",
				},
			},
			expected: errors.New("name field has an invalid length"),
		},
		{
			name: "invalid name (content)",
			model: UploadChunk{
				File: File{
					UUID: "39c4739c-91cf-11ee-b9d1-0242ac120002",
					Name: "gopher.jpg   ",
				},
			},
			expected: errors.New("name field does not satisfy the regular expression"),
		},
		{
			name: "invalid ChunkNum (not a number)",
			model: UploadChunk{
				File: File{
					UUID: "39c4739c-91cf-11ee-b9d1-0242ac120002",
					Name: "gopher.jpg",
				},
				ChunkNum: "x",
			},
			expected: errors.New("chunk_num field has an invalid"),
		},
		{
			name: "invalid ChunkNum (not a positive number)",
			model: UploadChunk{
				File: File{
					UUID: "39c4739c-91cf-11ee-b9d1-0242ac120002",
					Name: "gopher.jpg",
				},
				ChunkNum: "-13",
			},
			expected: errors.New("chunk_num field has an invalid"),
		},
		{
			name: "invalid TotalChunks (not a number)",
			model: UploadChunk{
				File: File{
					UUID: "39c4739c-91cf-11ee-b9d1-0242ac120002",
					Name: "gopher.jpg",
				},
				ChunkNum:    "0",
				TotalChunks: "x",
			},
			expected: errors.New("total_chunks field has an invalid"),
		},
		{
			name: "invalid TotalChunks (not a positive number)",
			model: UploadChunk{
				File: File{
					UUID: "39c4739c-91cf-11ee-b9d1-0242ac120002",
					Name: "gopher.jpg",
				},
				ChunkNum:    "0",
				TotalChunks: "-13",
			},
			expected: errors.New("total_chunks field has an invalid"),
		},
		{
			name: "invalid ChunkChecksum (length)",
			model: UploadChunk{
				File: File{
					UUID: "39c4739c-91cf-11ee-b9d1-0242ac120002",
					Name: "gopher.jpg",
				},
				ChunkNum:            "0",
				TotalChunks:         "11",
				SHA256ChunkChecksum: "cbb756",
			},
			expected: errors.New("checksum_chunk field has an invalid length"),
		},
		{
			name: "invalid ChunkChecksum (content)",
			model: UploadChunk{
				File: File{
					UUID: "39c4739c-91cf-11ee-b9d1-0242ac120002",
					Name: "gopher.jpg",
				},
				ChunkNum:            "0",
				TotalChunks:         "11",
				SHA256ChunkChecksum: "CBB756eb255316279a3e09cb7342c38754060a5b4bd6560e14f51d85cbd745e6", // use capital letters
			},
			expected: errors.New("checksum_chunk field does not satisfy the regular expression"),
		},
		{
			name: "invalid FileChecksum (length)",
			model: UploadChunk{
				File: File{
					UUID:               "39c4739c-91cf-11ee-b9d1-0242ac120002",
					Name:               "gopher.jpg",
					SHA256FileChecksum: "cd372",
				},
				ChunkNum:            "0",
				TotalChunks:         "11",
				SHA256ChunkChecksum: "cbb756eb255316279a3e09cb7342c38754060a5b4bd6560e14f51d85cbd745e6",
			},
			expected: errors.New("checksum_file field has an invalid length"),
		},
		{
			name: "invalid FileChecksum (content)",
			model: UploadChunk{
				File: File{
					UUID:               "39c4739c-91cf-11ee-b9d1-0242ac120002",
					Name:               "gopher.jpg",
					SHA256FileChecksum: "cd372f.85148700fa88095 3492d3f9f5beb43e555-5ff26d95f5a6adc36f8e6",
				},
				ChunkNum:            "0",
				TotalChunks:         "11",
				SHA256ChunkChecksum: "cbb756eb255316279a3e09cb7342c38754060a5b4bd6560e14f51d85cbd745e6",
			},
			expected: errors.New("checksum_file field does not satisfy the regular expression"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.model.Validate()
			assert.Equal(t, tt.expected, err)

		})
	}
}
