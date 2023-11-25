package filemgr

import (
	"crypto/sha256"
	"fmt"
)

func (fm fileMgr) SHA256Checksum(data []byte) string {
	h := sha256.New()
	h.Write(data)
	return fmt.Sprintf("%x", h.Sum(nil))
}
