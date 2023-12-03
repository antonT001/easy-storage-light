package models

import (
	"errors"
	"regexp"
	"strconv"

	"github.com/google/uuid"
)

type File struct {
	UUID               string
	Name               string
	SHA256FileChecksum string
}

type UploadChunk struct {
	File
	ChunkNum            string
	SHA256ChunkChecksum string
}

func (u UploadChunk) Valdate() error {
	const (
		maxLengthNameFile = 255
		lenghtChecksum    = 64

		filenameRegexExp = `^\S[\S ]*\S$`
		checksumRegexExp = `^[a-z0-9][a-z0-9]*[a-z0-9]$`
	)
	// UUID
	_, err := uuid.Parse(u.UUID)
	if err != nil {
		// TODO log.Errorf("uuid field is invalid, parse error: %v", err)
		return errors.New("uuid field is invalid, parse error")
	}

	// Name
	if len(u.Name) < 1 || len(u.Name) > maxLengthNameFile {
		// TODO log.Errorf("name field has an invalid length, expected: 1 - %v, actual: %v", maxLengthNameFile, len(u.Name))
		return errors.New("name field has an invalid length")
	}
	filenameRegexp := regexp.MustCompile(filenameRegexExp)
	if !filenameRegexp.MatchString(u.Name) {
		// TODO log.Errorf("name field does not satisfy the regular expression: %s, actual value: %v", filenameRegexExp, u.Name)
		return errors.New("name field does not satisfy the regular expression")
	}

	// ChunkNum
	_, err = strconv.ParseUint(u.ChunkNum, 10, 64)
	if err != nil {
		// TODO log.Errorf("chunk_num field has an invalid, expected: positive number, actual: %v", u.ChunkNum)
		return errors.New("chunk_num field has an invalid")
	}

	// ChunkChecksum
	if len(u.SHA256ChunkChecksum) != lenghtChecksum {
		// TODO log.Errorf("checksum_chunk field has an invalid length, expected: %v, actual: %v", lenghtChecksum, len(u.SHA256ChunkChecksum))
		return errors.New("checksum_chunk field has an invalid length")
	}
	checksumRegexp := regexp.MustCompile(checksumRegexExp)
	if !checksumRegexp.MatchString(u.SHA256ChunkChecksum) {
		// TODO log.Errorf("checksum_chunk field does not satisfy the regular expression: %s, actual value: %v", checksumRegexExp, u.SHA256ChunkChecksum)
		return errors.New("checksum_chunk field does not satisfy the regular expression")
	}

	// FileChecksum
	if len(u.SHA256FileChecksum) != lenghtChecksum {
		// TODO log.Errorf("checksum_file field has an invalid length, expected: %v, actual: %v", lenghtChecksum, len(u.SHA256FileChecksum))
		return errors.New("checksum_file field has an invalid length")
	}
	if !checksumRegexp.MatchString(u.SHA256FileChecksum) {
		// TODO log.Errorf("checksum_file field does not satisfy the regular expression: %s, actual value: %v", checksumRegexExp, u.SHA256FileChecksum)
		return errors.New("checksum_file field does not satisfy the regular expression")
	}

	return nil
}
