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

// TODO use something akin to github.com/rivo/uniseg to calculate character count, since this just counts the number of graphical code points, which is not the same.
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

// TODO feed as input to the unicode format group's 'code points' property
// TODO feed this into the text/characters property and get rid of "characters" above
func Runes(content io.Reader) io.RuneReader {
	return bufio.NewReader(content)
}
