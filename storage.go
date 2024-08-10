package main

import (
	"database/sql"
	"io"
)

// Storage provides persistence for the file server.
//
// Because the database keeps track of file identity, Storage implementations
// can be ignorant of files and operate only on contents. Hash calculation is
// the caller's responsibility, as is tracking what is being stored.
//
// Calling Read on a non-existent hash should panic.
type Storage interface {
	sql.DB
	Read(hash string) io.ReadSeeker
	Write(hash string, content io.Reader)
	Delete(hash string)
}
