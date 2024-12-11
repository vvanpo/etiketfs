package metadata

import "io"

type Registry struct {
	properties map[string]map[string]Evaluator[any]
}

func NewRegistry() Registry {
	return Registry{map[string]map[string]Evaluator[any]{}}
}

func SetEval[V Value](r Registry, group, name string, eval Evaluator[V]) {
	if r.properties[group] == nil {
		r.properties[group] = map[string]Evaluator[any]{}
	}
	e := func(c io.ReadSeeker) (any, error) { return eval(c) }

	r.properties[group][name] = e
}

func Lookup[V Value](r Registry, group, name string) Evaluator[V] {
	if r.properties[group] == nil {
		return nil
	}

	e := r.properties[group][name]

	return func(c io.ReadSeeker) (V, error) {
		v, err := e(c)

		return v.(V), err
	}
}

type Evaluator[V any] func(io.ReadSeeker) (V, error)
