package metadata

import "io"

type Registry struct {
	bytes     registry[Bytes]
	integer   registry[Integer]
	stringset registry[Set[string]]
	time      registry[Time]
}

func NewRegistry() Registry {
	return Registry{
		bytes:     registry[Bytes]{},
		integer:   registry[Integer]{},
		stringset: registry[Set[string]]{},
		time:      registry[Time]{},
	}
}

func SetEval[V Value](rs Registry, group, name string, eval Evaluator[V]) {
	r := get[V](rs).(registry[V])

	if r[group] == nil {
		r[group] = map[string]Evaluator[V]{}
	}

	r[group][name] = eval
}

func Lookup[V Value](rs Registry, group, name string) Evaluator[V] {
	r := get[V](rs).(registry[V])

	if r[group] == nil {
		return nil
	}

	return r[group][name]
}

func get[V Value](rs Registry) any {
	var val V

	switch any(val).(type) {
	case Bytes:
		return rs.bytes
	case Integer:
		return rs.integer
	case Set[string]:
		return rs.stringset
	case Time:
		return rs.time
	}

	return nil
}

type registry[V Value] map[string]map[string]Evaluator[V]

type Evaluator[V Value] func(io.ReadSeeker) (V, error)
