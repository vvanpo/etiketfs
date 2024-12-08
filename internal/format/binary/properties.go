package metadata

import "io/fs"

/*
Binary properties:
- File size (bytes)
- SHA-256 content hash (binary)
*/

func Size(f fs.File) (uint64, error) {
	fi, err := f.Stat()

	if err != nil {
		return 0, err
	}

	size := fi.Size()

	if size < 0 {
		return 0, nil
	}

	return uint64(size), nil
}
