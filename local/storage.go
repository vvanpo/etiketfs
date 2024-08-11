package local

import (
	"database/sql"
	"io"
	"os"
	pathpkg "path"
)

const dbFile = "etiketfs.db"

// Storage ...
type Storage struct {
	path string
	db   *sql.DB
}

// Load ...
func Load(path string) (*Storage, error) {
	dbPath := pathpkg.Join(path, dbFile)
	db, err := sql.Open("sqlite3", dbPath)

	if err != nil {
		return nil, err
	}

	return &Storage{path, db}, nil
}

func (s *Storage) Open(hash string) *os.File {
	f, err := os.Open(s.filepath(hash))

	if err != nil {
		panic(err)
	}

	return f
}

func (s *Storage) Write(hash string, content io.Reader) {
	f, err := os.Create(s.filepath(hash))

	if err != nil {
		panic(err)
	}

	if _, err := io.Copy(f, content); err != nil {
		panic(err)
	}
}

func (s *Storage) Delete(hash string) {
	if err := os.Remove(s.filepath(hash)); err != nil {
		panic(err)
	}
}

func (s *Storage) filepath(hash string) string {
	return pathpkg.Join(s.path, hash)
}

// New creates a Storage instance at the passed path.
func New(path string) error {
	return os.MkdirAll(path, 0755)
}
