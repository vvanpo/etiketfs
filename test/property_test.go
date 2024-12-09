package test

import (
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/vvanpo/vind"
	"github.com/vvanpo/vind/metadata"
	"github.com/vvanpo/vind/storage/unix"
)

func TestProperty(t *testing.T) {
	dir, _ := os.MkdirTemp("", "vind_test-")
	defer os.RemoveAll(dir)

	assert.Nil(t, unix.New(dir))
	fs, err := vind.Load(unix.Load(dir))
	assert.Nil(t, err)
	content := strings.NewReader("foo bar baz ほげ ふが")
	ts := time.Now()
	assert.Nil(t, fs.Add(content))

	files, err := fs.Select(vind.Filter{}, vind.Sort{})
	assert.Nil(t, err)
	f := <-files

	size, err := vind.Property[metadata.Bytes](fs, f, "binary", "size")
	assert.Nil(t, err)
	assert.Equal(t, metadata.Bytes(25), *size)
	added, err := vind.Property[metadata.Time](fs, f, "", "added")
	assert.Nil(t, err)
	assert.Equal(t, metadata.Time(ts.Unix()), added)
	chars, err := vind.Property[metadata.Integer](fs, f, "unicode", "characters")
	assert.Nil(t, err)
	assert.Equal(t, metadata.Integer(17), *chars)

	_, ok := <-files
	assert.False(t, ok)
}
