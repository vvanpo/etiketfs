package main

import "io"

type Content interface {
	io.ReadWriteSeeker
}

// File ...
type File struct {
	id      string
	content Content
}

func (f File) Property(i PropertyIdentifier) <-chan PropertyValue {
	return nil
}
