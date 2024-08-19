package local

import (
	"database/sql"
	"io"
	"os"
	pathpkg "path"
)

const dbFile = "etiketfs.db"

// New creates a Storage instance at the passed path.
func New(path string) error {
	return os.MkdirAll(path, 0755)
}

// Storage ...
type Storage struct {
	path string
}

func Load(path string) *Storage {
	return &Storage{path}
}

func (s *Storage) DB() (*sql.DB, error) {
	return sql.Open("sqlite3", s.absPath(dbFile))
}

func (s *Storage) Open(id string) (*os.File, error) {
	return os.Open(s.absPath(id))
}

func (s *Storage) Add(id string, content io.Reader) error {
	f, err := os.Create(s.absPath(id))

	if err != nil {
		return err
	}

	_, err = io.Copy(f, content)

	return err
}

func (s *Storage) Delete(id string) error {
	return os.Remove(s.absPath(id))
}

func (s *Storage) absPath(id string) string {
	return pathpkg.Join(s.path, id)
}
