package main

import (
	"golang.org/x/exp/maps"
)

type Sorter struct {
	i   PropertyIdentifier
	cmp func(a, b PropertyValue) bool
}

func (s Sorter) Reverse() Sorter {
	s.cmp = func(a, b PropertyValue) bool { return s.cmp(b, a) }

	return s
}

func NewSorter(identifier PropertyIdentifier, comparator func(a, b PropertyValue) bool) Sorter {
	return Sorter{identifier, comparator}
}

func Sort(in <-chan Selection, sorters ...Sorter) <-chan []File {
	out := make(chan []File)

	go func() {
		for selection := range in {
			out <- sortSelection(selection, sorters...)
		}
	}()

	return out
}

func sortSelection(selection Selection, sorters ...Sorter) []File {
	files := maps.Values(selection.files)
	properties := make(map[PropertyIdentifier][]PropertyValue)

	for sorter := range sorters {
	}

	//
	return files
}
