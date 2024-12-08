package unix

import (
	"os"
	"path"
	"strings"
	"testing"
)

func TestAddRead(t *testing.T) {
	dir := mkdirTemp()
	defer os.RemoveAll(dir)

	New(dir)
	s := Load(dir)

	id := "foo"
	s.Add(id, strings.NewReader("bar"))

	readContent := make([]byte, 3)
	content, _ := s.Open(id)
	content.Read(readContent)

	if string(readContent) != "bar" {
		t.Errorf("Failed read correct file contents")
	}
}

func TestAddDelete(t *testing.T) {
	dir := mkdirTemp()
	defer os.RemoveAll(dir)

	New(dir)
	s := Load(dir)

	id := "foo"
	s.Add(id, strings.NewReader("bar"))
	s.Delete(id)

	if _, err := os.Stat(path.Join(dir, "foo")); err == nil {
		t.Error("File not deleted")
	}
}
