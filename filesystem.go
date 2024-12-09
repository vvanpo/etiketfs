package vind

import (
	"fmt"
	"io"

	"github.com/google/uuid"
	"github.com/vvanpo/vind/internal/format/binary"
	"github.com/vvanpo/vind/internal/format/utf8"
	"github.com/vvanpo/vind/internal/state"
	"github.com/vvanpo/vind/metadata"
)

// Filesystem ...
type Filesystem struct {
	storage  Storage
	state    state.State
	registry metadata.Registry
}

func initRegistry() metadata.Registry {
	r := metadata.NewRegistry()
	metadata.SetEval(r, "binary", "size", binary.Size)
	metadata.SetEval(r, "unicode", "characters", utf8.Characters)

	return r
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

	return Filesystem{so, st, initRegistry()}, nil
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
func Property[V metadata.Value](fs Filesystem, file File, group, name string, params ...any) (*V, error) {
	content, err := fs.storage.Open(file.id.String())

	if err != nil {
		return nil, err
	}

	eval := metadata.Lookup[V](fs.registry, group, name)

	if eval == nil {
		return nil, fmt.Errorf("vind: invalid property (%s/%s)", group, name)
	}

	val, err := eval(content)

	if err != nil {
		return nil, err
	}

	return &val, nil
}

type Filter struct{}

type Sort struct{}

type File struct {
	id uuid.UUID
}
