package essentials

import (
	"io"
)

// CountingReader allows monitoring of read func
// contributed by Jakob Borg
type CountingReader struct {
	Reader    io.Reader
	BytesRead int64 // bytes
}

// Read func of *CountingReader allows monitoring of read func
// contributed by Jakob Borg
func (c *CountingReader) Read(bs []byte) (int, error) {
	n, err := c.Reader.Read(bs)
	c.BytesRead += int64(n)
	return n, err
}
