package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"

	"github.com/pkg/errors"
)

const (
	size1K   = 1024
	size1M   = 1024 * 1024
	size20M  = 20 * 1024 * 1024
	size200M = 200 * 1024 * 1024

	seekSet = 0
)

func Checksum(name string) (string, error) {
	var size int64
	var sha string
	var err error

	size, err = getFileSize(name)
	if err != nil {
		return "", errors.Wrap(err, "failed to get file size\n")
	}

	if size < size20M {
		sha, err = checksumSmallFile(name)
	} else if size < size200M {
		sha, err = checksumMediumFile(name)
	} else {
		sha, err = checksumLargeFile(name)
	}

	if err != nil {
		return "", errors.Wrap(err, "failed to get file checksum\n")
	}

	return sha, nil
}

func getFileSize(name string) (int64, error) {
	info, err := os.Stat(name)
	if err != nil {
		return 0, errors.Wrap(err, "failed to stat file\n")
	}

	return info.Size(), nil
}

func checksumSmallFile(name string) (string, error) {
	file, err := os.Open(name)
	if err != nil {
		return "", errors.Wrap(err, "failed to open file\n")
	}

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", errors.Wrap(err, "failed to hash file\n")
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}

func checksumMediumFile(name string) (string, error) {
	file, err := os.Open(name)
	if err != nil {
		return "", errors.Wrap(err, "failed to open file\n")
	}

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	chunkSize := int64(size1M)
	hash := sha256.New()

	for {
		buffer := make([]byte, chunkSize)
		n, err := file.Read(buffer)
		if err != nil {
			if err == io.EOF {
				_, _ = hash.Write(buffer[:n])
				break
			}
			return "", errors.Wrap(err, "failed to read file\n")
		}
		_, _ = hash.Write(buffer[:n])
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}

func checksumLargeFile(name string) (string, error) {
	file, err := os.Open(name)
	if err != nil {
		return "", errors.Wrap(err, "failed to open file\n")
	}

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	startBuffer := make([]byte, size1K)
	if _, err := io.ReadFull(file, startBuffer); err != nil {
		return "", errors.Wrap(err, "failed to read file\n")
	}

	hash := sha256.New()
	if _, err := hash.Write(startBuffer); err != nil {
		return "", errors.Wrap(err, "failed to hash file\n")
	}

	fileStat, err := file.Stat()
	if err != nil {
		return "", errors.Wrap(err, "failed to stat file\n")
	}

	if _, err := file.Seek(fileStat.Size()-int64(size1K), seekSet); err != nil {
		return "", errors.Wrap(err, "failed to seek file\n")
	}

	endBuffer := make([]byte, size1K)
	if _, err := io.ReadFull(file, endBuffer); err != nil {
		return "", errors.Wrap(err, "failed to read file\n")
	}

	if _, err := hash.Write(endBuffer); err != nil {
		return "", errors.Wrap(err, "failed to hash file\n")
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}
