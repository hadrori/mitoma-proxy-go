package utils

import (
	"bytes"
	"compress/gzip"
	"io"
	"log"
)

func ReadGzip(gzipByte []byte) []byte {
	gzipReader, err := gzip.NewReader(bytes.NewReader(gzipByte))
	if err != nil {
		log.Printf("Failed to extract gzip")
	}
	defer gzipReader.Close()

	decoded, err := io.ReadAll(gzipReader)
	if err != nil {
		log.Printf("Failed to read gzip")
	}
	return decoded
}
