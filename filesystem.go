package main

// Filesystem ...
type Filesystem struct {
	storage Storage
}

// Load ...
func Load(s Storage) Filesystem {
	return Filesystem{s}
}
