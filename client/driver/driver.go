package driver

import "github.com/vvanpo/vind"

type Driver interface {
	List(vind.Filter, vind.Sort, []Identifier) (<-chan []any, error)
}

type Identifier struct {
	Group     string
	Name      string
	Parameter any
}
