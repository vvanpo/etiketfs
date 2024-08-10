package local

import (
	"database/sql"
	"io"
	"io/fs"
	"os"
	pathpkg "path"
)

const dbFile = "etiketfs.db"

// Storage ...
type Storage struct {
	fs fs.FS
	*sql.DB
}

// Load ...
func Load(path string) (*Storage, error) {
	fs := os.DirFS(path)
	dbPath := pathpkg.Join(path, dbFile)
	db, err := sql.Open("sqlite3", dbPath)

	if err != nil {
		return nil, err
	}

	return &Storage{fs, db}, nil
}

func (s *Storage) Read(hash string) io.Reader {
	f, err := s.fs.Open(hash)

	if err != nil {
		panic(err)
	}

	return f
}

// New creates a Storage instance at the passed path.
func New(path string) error {
	return os.MkdirAll(path, 0755)
}
