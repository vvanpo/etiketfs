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
// Calling Read or Delete on a non-existent hash should panic, as should a Write
// on a hash that is already present.
type Storage interface {
	sql.DB
	Read(hash string) io.ReadSeeker
	Write(hash string, content io.Reader)
	Delete(hash string)
}
