package config

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig_validate(t *testing.T) {
	type fields struct {
		Server  ServerConfig
		FileMgr FileMgrConfig
	}
	tests := []struct {
		name     string
		fields   fields
		expected error
	}{
		{
			name: "valid config",
			fields: fields{
				Server: ServerConfig{
					Host: "localhost",
					Port: 9012,
				},
				FileMgr: FileMgrConfig{
					StorageBasePath: "/opt/easy-storage-light/data",
				},
			},
			expected: nil,
		},
		{
			name: "host field is empty",
			fields: fields{
				Server: ServerConfig{
					Host: "",
					Port: 9012,
				},
				FileMgr: FileMgrConfig{
					StorageBasePath: "/opt/easy-storage-light/data",
				},
			},
			expected: invalidErrorWrap(errHostEmpty),
		},
		{
			name: "port field is empty",
			fields: fields{
				Server: ServerConfig{
					Host: "localhost",
				},
				FileMgr: FileMgrConfig{
					StorageBasePath: "/opt/easy-storage-light/data",
				},
			},
			expected: invalidErrorWrap(errPortEmpty),
		},
		{
			name: "storage_base_path field is empty",
			fields: fields{
				Server: ServerConfig{
					Host: "localhost",
					Port: 9012,
				},
			},
			expected: invalidErrorWrap(errStorageBasePathEmpty),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				Server:  tt.fields.Server,
				FileMgr: tt.fields.FileMgr,
			}
			err := c.validate()
			assert.Equal(t, tt.expected, err)
		})
	}
}

func invalidErrorWrap(err error) error {
	return fmt.Errorf("%v: %v", confidInvalid, err)
}
