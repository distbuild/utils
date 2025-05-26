package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	nameGetFileSize = "../test/checksum/small.txt"
	sizeGetFileSize = 19922944

	nameChecksumSmallFile = "../test/checksum/small.txt"
	sumChecksumSmallFile  = "2454ff43c4aacb28c69a65ec1ab3b674e5d77bf42fcfbd844b0b320bdb6776f5"

	nameChecksumMediumFile = "../test/checksum/medium.txt"
	sumChecksumMediumFile  = "bcb044fa06ad5f878f03a6249fd0649df7c344fc79c23b08b0a134b0ac334e51"

	nameChecksumLargeFile = "../test/checksum/large.txt"
	sumChecksumLargeFile  = "2125208dde1bc7866fd48ed8f6a42a1c36113c0a24c7557504d67f938045f9a4"
)

func TestGetFileSize(t *testing.T) {
	size, err := getFileSize(nameGetFileSize)
	assert.Equal(t, nil, err)
	assert.Equal(t, int64(sizeGetFileSize), size)
}

func TestChecksumSmallFile(t *testing.T) {
	sum, err := checksumSmallFile(nameChecksumSmallFile)
	assert.Equal(t, nil, err)
	assert.Equal(t, sumChecksumSmallFile, sum)
}

func TestChecksumMediumFile(t *testing.T) {
	sum, err := checksumMediumFile(nameChecksumMediumFile)
	assert.Equal(t, nil, err)
	assert.Equal(t, sumChecksumMediumFile, sum)
}

func TestChecksumLargeFile(t *testing.T) {
	sum, err := checksumLargeFile(nameChecksumLargeFile)
	assert.Equal(t, nil, err)
	assert.Equal(t, sumChecksumLargeFile, sum)
}
