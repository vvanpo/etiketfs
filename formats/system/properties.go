package metadata

import "io/fs"

/*
System properties:
  Intrinsic:
   - File size in bytes (uint)
   - Content hash
  Extrinsic:
   - Added (datetime)
   - Modified (datetime)
   - Accessed (datetime)
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
