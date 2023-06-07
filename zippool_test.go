// zippool_test.go
package gzippool

import (
	"bytes"
	"compress/gzip"
	"testing"
)

var zipPool = NewZipPoolLevel(gzip.BestCompression)

type Data struct {
	Str string
	I   int64
}

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

	// dataD := &Data{
	// 	Str: "Hello, Golang zip pool!",
	// 	I:   1234567890,
	// }
	//
	// for i := 0; i < 10; i++ {
	// 	// struct to []byte
	//
	// 	compressedData1, err := zipPool.Compress(dataB)
	// 	if err != nil {
	// 		t.Fatalf("Failed to compress data: %v", err)
	// 	}
	//
	// 	decompressedData1, err := zipPool.Decompress(compressedData1)
	// 	if err != nil {
	// 		t.Fatalf("Failed to decompress data: %v", err)
	// 	}
	//
	// 	if !bytes.Equal(dataB, decompressedData1) {
	// 		t.Fatalf("Decompressed data does not match original data: got %s, want %s", decompressedData1, compressedData1)
	// 	}
	// }

}
