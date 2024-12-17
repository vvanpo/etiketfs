package format

import (
	"fmt"
	"io"
	"unicode"
)

func Codepoints(runes io.RuneReader) (uint64, error) {
	var count uint64

	for {
		r, n, err := runes.ReadRune()

		if n == 0 {
			break
		}

		if err != nil {
			return 0, err
		}

		if r == unicode.ReplacementChar {
			return 0, fmt.Errorf("unicode: invalid code point")
		}

		count++
	}

	return count, nil
}
