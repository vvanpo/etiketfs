package test

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vvanpo/vind"
	"github.com/vvanpo/vind/storage/unix"
)

func TestProperty(t *testing.T) {
	dir, _ := os.MkdirTemp("", "vind_test-")
	defer os.RemoveAll(dir)

	assert.Nil(t, unix.New(dir))
	fs, err := vind.Load(unix.Load(dir))
	assert.Nil(t, err)
	content := strings.NewReader("foo bar baz ほげ ふが")
	assert.Nil(t, fs.Add(content))

	files, err := fs.Select(vind.Filter{}, vind.Sort{})
	assert.Nil(t, err)
	f := <-files

	size, err := f.Property("binary", "size")
	assert.Equal(t, 21, size)
	chars, err := f.Property("unicode", "characters")
	assert.Equal(t, 17, chars)
}
