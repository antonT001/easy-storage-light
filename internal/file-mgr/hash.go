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

func (fm fileMgr) IsEquivalentSHA256Checksum(data []byte, hashSHA256 string) bool {
	return hashSHA256 == fm.SHA256Checksum(data)
}
