package unix

import (
	"os"
	"path"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func mkdirTemp() string {
	tmpdir, _ := os.MkdirTemp("", "vind-storage-unix_test-")

	return tmpdir
}

func TestNewExistingDir(t *testing.T) {
	tmpdir := mkdirTemp()
	defer os.RemoveAll(tmpdir)

	t.Logf("Created %s", tmpdir)

	if err := New(tmpdir); err != nil {
		t.Error(err)
	}
}

func TestNewNestedDir(t *testing.T) {
	tmpdir := mkdirTemp()
	defer os.RemoveAll(tmpdir)

	nesteddir := path.Join(tmpdir, "foo/bar")

	if err := New(nesteddir); err != nil {
		t.Fatalf("Failed to create %s: %v", nesteddir, err)
	}

	fi, err := os.Stat(nesteddir)

	if err != nil {
		t.Fatalf("Failed to stat %s: %v", nesteddir, err)
	}

	if !fi.IsDir() {
		t.Fatalf("%s is not a directory", nesteddir)
	}

	t.Logf("Created %s", nesteddir)
}

// func TestNewNoPermission()

// func TestNewPathNotDir()
