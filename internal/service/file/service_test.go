package fileservice

import (
	"io"
	"testing"

	filemgr "github.com/antonT001/easy-storage-light/internal/file-mgr"
	"github.com/antonT001/easy-storage-light/internal/models"
	filerepository "github.com/antonT001/easy-storage-light/internal/repository/file"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type Mocks struct {
	FileRepo *filerepository.MockRepository
	FileMgr  *filemgr.MockFileMgr
}

func Test_service_UploadChunk(t *testing.T) {
	type args struct {
		upload models.UploadChunk
		body   io.ReadCloser
	}
	tests := []struct {
		name     string
		args     args
		prepare  func(m *Mocks)
		expected error
	}{
		{
			name: "ok",
			args: args{
				upload: models.UploadChunk{},
			},
			prepare: func(m *Mocks) {
				createFile := m.FileMgr.EXPECT().CreateFile(gomock.Any()).Return(nil, nil)
				copyFile := m.FileMgr.EXPECT().CopyFile(gomock.Any(), gomock.Any()).Return(nil).After(createFile)
				addChunk := m.FileRepo.EXPECT().AddChunk(gomock.Any()).Return(nil).After(copyFile)
				m.FileRepo.EXPECT().AllChunksUploadedForUUID("").Return(true, nil).After(addChunk)
			},
			expected: nil,
		},
	}

	ctrl := gomock.NewController(t)
	mockFilemgr := filemgr.NewMockFileMgr(ctrl)
	mockFileRepo := filerepository.NewMockRepository(ctrl)
	svc := serviceImpl{
		fileRepo: mockFileRepo,
		fileMgr:  mockFilemgr,
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare(&Mocks{
				FileRepo: mockFileRepo,
				FileMgr:  mockFilemgr,
			})
			err := svc.UploadChunk(tt.args.upload, tt.args.body)
			assert.Equal(t, tt.expected, err)
		})
	}
}
