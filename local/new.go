package local

import (
	"database/sql"
	pathpkg "path"

	_ "github.com/mattn/go-sqlite3"
)

// New takes an empty directory path and creates a LocalStorage instance.
func New(path string) error {
	dbPath := pathpkg.Join(path, dbFile)
	_, err := sql.Open("sqlite3", dbPath)

	if err != nil {
		return err
	}

	return nil
}
