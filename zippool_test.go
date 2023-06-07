// zippool_test.go
package gzippool

import (
	"bytes"
	"compress/gzip"
	"testing"
)

var zipPool = NewZipPool(gzip.BestCompression)

func TestZipPool(t *testing.T) {

	data := []byte("Hello, Golang zip pool!")
	for i := 0; i < 10; i++ {
		compressedData, err := zipPool.Compress(data)
		if err != nil {
			t.Fatalf("Failed to compress data: %v", err)
		}

		decompressedData, err := zipPool.Decompress(compressedData)
		if err != nil {
			t.Fatalf("Failed to decompress data: %v", err)
		}

		if !bytes.Equal(data, decompressedData) {
			t.Fatalf("Decompressed data does not match original data: got %s, want %s", decompressedData, data)
		}
	}
}
