package mocks

import (
	"github.com/golang/mock/gomock"
)

func Upload(m *Mocks) {
	m.Service.EXPECT().UploadChunk(gomock.Any(), gomock.Any()).Return(nil)
}
