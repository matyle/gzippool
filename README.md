# gzippool

Using a sync.Pool to cache GZip writers and readers could be a good idea to reduce the overhead of allocating and garbage collecting them.

gzippool is a Go package that provides a container for reusing gzip.Writer and gzip objects.

## Installation

Use `go get` to install gzippool.

```bash
go get github.com/matyle/gzippool
```

## Usage

```go
package main

import (
"fmt"
"compress/gzip"

"github.com/matyle/gzippool"
)

func main() {
// Create a new GzipPool instance
gzipPool := gzippool.NewGzipPool()

// Compress data
data := []byte("Hello, Golang gzip pool!")
compressedData, err := gzipPool.Compress(data)
if err != nil {
panic(err)
}

fmt.Printf("Compressed data: %x\n", compressedData)

// Decompress data
decompressedData, err := gzipPool.Decompress(compressedData)
if err != nil {
panic(err)
}

fmt.Printf("Decompressed data: %s\n", decompressedData)

// Create a new GzipPoolLevel instance
gzipPoolLevel := gzippool.NewGzipPoolLevel(gzip.BestCompression)

// Compress data
data := []byte("Hello, Golang gzip pool!")
compressedData, err := gzipPoolLevel.Compress(data)
if err != nil {
panic(err)
}

fmt.Printf("Compressed data: %x\n", compressedData)

// Decompress data
decompressedData, err := gzipPoolLevel.Decompress(compressedData)
if err != nil {
panic(err)
}

fmt.Printf("Decompressed data: %s\n", decompressedData)

}
```

- Use Test

```bash
go test -v .
```

## License

gzippool is licensed under the MIT License. See LICENSE for more information.
