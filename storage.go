package vind

import (
	"database/sql"
	"io"
)

// Storage provides persistence for the file server.
type Storage interface {
	DB() (*sql.DB, error)
	Open(id string) (io.ReadSeekCloser, error)
	Add(id string, content io.Reader) error
	Delete(id string) error
}
