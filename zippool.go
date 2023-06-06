package gzippool

import (
	"bytes"
	"compress/gzip"
	"io"
	"sync"
)

// ZipPool is a container for reusing gzip.Writer and gzip.Reader objects.
type ZipPool struct {
	writerPool sync.Pool
	readerPool sync.Pool
}

// NewZipPool creates a new ZipPool instance.
func NewZipPool() *ZipPool {
	return &ZipPool{
		writerPool: sync.Pool{
			New: func() interface{} {
				return gzip.NewWriter(nil)
			},
		},
		readerPool: sync.Pool{
			New: func() interface{} {
				return new(gzip.Reader)
			},
		},
	}
}

// GetWriter retrieves a gzip.Writer from thePool.
func (zp *ZipPool) GetWriter() *gzip.Writer {
	return zp.writerPool.Get().(*gzip.Writer)
}

// PutWriter returns a gzip.Writer to the ZipPool.
func (zp *ZipPool) PutWriter(w *gzip.Writer) {
	w.Reset(nil)
	zp.writerPool.Put(w)
}

// GetReader retrieves a gzip.Reader from the ZipPool.
func (zp *ZipPool) GetReader() *gzip.Reader {
	return zp.readerPool.Get().(*gzip.Reader)
}

// PutReader returns a gzip.Reader to the ZipPool.
func (zp *ZipPool) PutReader(r *gzip.Reader) {
	zp.readerPool.Put(r)
}

// Compress compresses the input data using gzip.
func (zp *ZipPool) Compress(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	writer := zp.GetWriter()
	writer.Reset(&buf)

	_, err := writer.Write(data)
	if err != nil {
		return nil, err
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	zp.PutWriter(writer)
	return buf.Bytes(), nil
}

// Decompress decompresses the input data using gzip.
func (zp *ZipPool) Decompress(data []byte) ([]byte, error) {
	reader := zp.GetReader()
	err := reader.Reset(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, reader)
	if err != nil {
		return nil, err
	}

	err = reader.Close()
	if err != nil {
		return nil, err
	}

	zp.PutReader(reader)
	return buf.Bytes(), nil
}
