package binary

import (
	"io"

	"github.com/vvanpo/vind/metadata"
)

/*
Binary properties:
- File size (bytes)
- SHA-256 content hash (binary)
*/

func Size(content io.ReadSeeker) (metadata.Bytes, error) {
	size, err := content.Seek(0, io.SeekEnd)

	return metadata.Bytes(size), err
}
