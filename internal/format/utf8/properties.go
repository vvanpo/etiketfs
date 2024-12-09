package utf8

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"unicode"
	"unicode/utf8"

	"github.com/vvanpo/vind/metadata"
)

func Characters(content io.ReadSeeker) (metadata.Integer, error) {
	buffered := bufio.NewReader(content)
	var count metadata.Integer

	for {
		r, n, err := buffered.ReadRune()

		if n == 0 {
			break
		}

		if err != nil {
			return 0, err
		}

		if r == utf8.RuneError {
			return 0, fmt.Errorf("Invalid UTF-8")
		}

		if unicode.IsGraphic(r) {
			count++
		}

	}

	return count, nil
}

func Contains(content io.Reader, match string) (bool, error) {
	r := bufio.NewReader(content)

	return regexp.MatchReader(match, r)
}
