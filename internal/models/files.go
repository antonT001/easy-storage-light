package models

import (
	"errors"
	"log"
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
	TotalChunks         string
	SHA256ChunkChecksum string
}

func (u UploadChunk) Validate() error {
	const (
		maxLengthNameFile = 255
		lenghtChecksum    = 64

		filenameRegexExp = `^\S[\S ]*\S$`
		checksumRegexExp = `^[a-z0-9][a-z0-9]*[a-z0-9]$`
	)
	// UUID
	_, err := uuid.Parse(u.UUID)
	if err != nil {
		log.Printf("uuid field is invalid, parse error: %v", err) // TODO log.Errorf
		return errors.New("uuid field is invalid, parse error")
	}

	// Name
	if len(u.Name) < 1 || len(u.Name) > maxLengthNameFile {
		log.Printf("name field has an invalid length, expected: 1 - %v, actual: %v", maxLengthNameFile, len(u.Name)) // TODO log.Errorf
		return errors.New("name field has an invalid length")
	}
	filenameRegexp := regexp.MustCompile(filenameRegexExp)
	if !filenameRegexp.MatchString(u.Name) {
		log.Printf("name field does not satisfy the regular expression: %s, actual value: %v", filenameRegexExp, u.Name) // TODO log.Errorf
		return errors.New("name field does not satisfy the regular expression")
	}

	// ChunkNum
	_, err = strconv.ParseUint(u.ChunkNum, 10, 64)
	if err != nil {
		log.Printf("chunk_num field has an invalid, expected: positive number, actual: %v", u.ChunkNum) // TODO log.Errorf
		return errors.New("chunk_num field has an invalid")
	}

	// TotalChunks
	_, err = strconv.ParseUint(u.TotalChunks, 10, 64)
	if err != nil {
		log.Printf("total_chunks field has an invalid, expected: positive number, actual: %v", u.TotalChunks) // TODO log.Errorf
		return errors.New("total_chunks field has an invalid")
	}

	// ChunkChecksum
	if len(u.SHA256ChunkChecksum) != lenghtChecksum {
		log.Printf("checksum_chunk field has an invalid length, expected: %v, actual: %v", lenghtChecksum, len(u.SHA256ChunkChecksum)) // TODO log.Errorf
		return errors.New("checksum_chunk field has an invalid length")
	}
	checksumRegexp := regexp.MustCompile(checksumRegexExp)
	if !checksumRegexp.MatchString(u.SHA256ChunkChecksum) {
		log.Printf("checksum_chunk field does not satisfy the regular expression: %s, actual value: %v", checksumRegexExp, u.SHA256ChunkChecksum) // TODO log.Errorf
		return errors.New("checksum_chunk field does not satisfy the regular expression")
	}

	// FileChecksum
	if len(u.SHA256FileChecksum) != lenghtChecksum {
		log.Printf("checksum_file field has an invalid length, expected: %v, actual: %v", lenghtChecksum, len(u.SHA256FileChecksum)) // TODO log.Errorf
		return errors.New("checksum_file field has an invalid length")
	}
	if !checksumRegexp.MatchString(u.SHA256FileChecksum) {
		log.Printf("checksum_file field does not satisfy the regular expression: %s, actual value: %v", checksumRegexExp, u.SHA256FileChecksum) // TODO log.Errorf
		return errors.New("checksum_file field does not satisfy the regular expression")
	}

	return nil
}
