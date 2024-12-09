package metadata

import "io"

/*
Binary properties:
- File size (bytes)
- SHA-256 content hash (binary)
*/

func Size(content io.Seeker) (uint64, error) {
	size, err := content.Seek(0, io.SeekEnd)

	return uint64(size), err
}
