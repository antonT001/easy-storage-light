package filemgr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_fileMgr_SHA256Checksum(t *testing.T) {
	type fields struct {
		storageBasePath string
	}
	type args struct {
		data []byte
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		expected string
	}{
		{
			name: "valid hash 1",
			args: args{
				data: []byte("He loved solitude so much that he could sit in an empty chat all day…"),
			},
			expected: "cbb756eb255316279a3e09cb7342c38754060a5b4bd6560e14f51d85cbd745e6",
		},
		{
			name: "valid hash 2",
			args: args{
				data: []byte(""),
			},
			expected: "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
		},
		{
			name: "valid hash 3",
			args: args{
				data: []byte("e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"),
			},
			expected: "cd372fb85148700fa88095e3492d3f9f5beb43e555e5ff26d95f5a6adc36f8e6",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fm := fileMgr{
				storageBasePath: tt.fields.storageBasePath,
			}
			actual := fm.SHA256Checksum(tt.args.data)
			assert.Equal(t, tt.expected, actual)

		})
	}
}
