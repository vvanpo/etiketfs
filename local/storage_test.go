package local

import (
	"os"
	"path"
	"strings"
	"testing"
)

func TestWriteRead(t *testing.T) {
	dir := mkdirTemp()
	defer os.RemoveAll(dir)

	New(dir)
	s, _ := Load(dir)

	hash := "foo"
	content := strings.NewReader("bar")
	s.Write(hash, content)

	readContent := make([]byte, 3)
	s.Open(hash).Read(readContent)

	if string(readContent) != "bar" {
		t.Errorf("Failed read correct file contents")
	}
}

func TestWriteDelete(t *testing.T) {
	dir := mkdirTemp()
	defer os.RemoveAll(dir)

	New(dir)
	s, _ := Load(dir)

	hash := "foo"
	content := strings.NewReader("bar")
	s.Write(hash, content)
	s.Delete(hash)

	if _, err := os.Stat(path.Join(dir, "foo")); err == nil {
		t.Error("File not deleted")
	}
}
