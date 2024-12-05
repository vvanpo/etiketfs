package vind

// Filesystem ...
type Filesystem struct{}

func (fs Filesystem) Select(filter Filter, sort Sort) (<-chan File, error)

type Filter struct{}

type Sort struct{}

type File struct{}

// Property ...
func (f File) Property(group, name string, params ...any) (any, error)
