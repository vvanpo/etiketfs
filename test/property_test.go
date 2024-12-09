package test

import (
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/vvanpo/vind"
	"github.com/vvanpo/vind/storage/unix"
)

func TestProperty(t *testing.T) {
	dir, _ := os.MkdirTemp("", "vind_test-")
	// defer os.RemoveAll(dir)

	assert.Nil(t, unix.New(dir))
	fs, err := vind.Load(unix.Load(dir))
	assert.Nil(t, err)
	content := strings.NewReader("foo bar baz ほげ ふが")
	ts := time.Now()
	assert.Nil(t, fs.Add(content))

	files, err := fs.Select(vind.Filter{}, vind.Sort{})
	assert.Nil(t, err)
	f := <-files

	size, err := fs.Property(f, "binary", "size")
	assert.Equal(t, 21, size)
	added, err := fs.Property(f, "", "added")
	assert.Equal(t, ts.Unix(), added)
	chars, err := fs.Property(f, "unicode", "characters")
	assert.Equal(t, 17, chars)

	_, ok := <-files
	assert.False(t, ok)
}
