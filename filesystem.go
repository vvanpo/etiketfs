package vind

import (
	"io"

	"github.com/google/uuid"
	"github.com/vvanpo/vind/internal/state"
)

// Filesystem ...
type Filesystem struct {
	storage Storage
	state   state.State
}

func Load(so Storage) (Filesystem, error) {
	db, err := so.DB()

	if err != nil {
		return Filesystem{}, err
	}

	st, err := state.Init(db)

	if err != nil {
		return Filesystem{}, err
	}

	return Filesystem{so, st}, nil
}

func (fs Filesystem) Add(content io.Reader) error {
	id := uuid.New()

	if err := fs.storage.Add(id.String(), content); err != nil {
		return err
	}

	if err := fs.state.AddFile(id); err != nil {
		fs.storage.Delete(id.String())

		return err
	}

	return nil
}

func (fs Filesystem) Select(filter Filter, sort Sort) (<-chan File, error) {
	ids, err := fs.state.FileIds()

	if err != nil {
		return nil, err
	}

	out := make(chan File)

	go func() {
		for _, id := range ids {
			out <- File{id}
		}

		close(out)
	}()

	return out, nil
}

// Property ...
func (fs Filesystem) Property(file File, group, name string, params ...any) (any, error) {
	content, err := fs.storage.Open(file.id.String())

	if err != nil {
		return nil, err
	}

	// lookup metadata source

	return nil, nil
}

type Filter struct{}

type Sort struct{}

type File struct {
	id uuid.UUID
}
