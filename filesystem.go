package vind

import "io"

// Filesystem ...
type Filesystem struct {
	storage Storage
}

func Load(s Storage) Filesystem {
	return Filesystem{s}
}

func (fs Filesystem) Add(content io.Reader) error

func (fs Filesystem) Select(filter Filter, sort Sort) (<-chan File, error)

type Filter struct{}

type Sort struct{}

type File struct{}

// Property ...
func (f File) Property(group, name string, params ...any) (any, error)
