package utf8

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCharacters(t *testing.T) {
	r := strings.NewReader("foo")

	count, err := Characters(r)

	if err != nil {
		t.Fatalf("Rune-read error: %v", err)
	} else if count != 3 {
		t.Fatalf("Incorrect count: %d", count)
	}
}

func TestCharactersInvalidUTF8(t *testing.T) {
	r := bytes.NewReader([]byte{'f', 'o', 'o', ' ', 128})

	_, err := Characters(r)

	if err == nil {
		t.Fatal("Invalid UTF-8 not detected")
	} else if err.Error() != "Invalid UTF-8" {
		t.Fatalf("Unexpected error: %v", err)
	}
}

func TestContains(t *testing.T) {
	r := strings.NewReader("foo barbaz")

	c, err := Contains(r, "arb")

	assert.True(t, c)
	assert.Nil(t, err)
}

func TestContainsFalse(t *testing.T) {
	r := strings.NewReader("foo barbaz")

	c, err := Contains(r, "qux")

	assert.False(t, c)
	assert.Nil(t, err)
}
