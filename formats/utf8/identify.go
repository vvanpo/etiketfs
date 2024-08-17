package utf8

import (
	"io"
	"unicode/utf8"
)

func Identify(file io.Reader) (bool, error) {
	const bufSize = 1024
	buf := make([]byte, bufSize)

	for {
		n, err := file.Read(buf)

		if err != nil {
			return false, err
		}

		if !utf8.Valid(buf[:n]) {
			return false, nil
		}

		if n < 1024 {
			break
		}
	}

	return true, nil
}
