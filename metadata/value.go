package metadata

type Value interface {
	string | Bytes | Integer | Set[string] | Time
}

type Bytes uint64
type Integer int64
type Set[E string] map[E]struct{}
type Time int64
