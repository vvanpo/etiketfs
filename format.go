package main

import "io"

type Format interface {
	// Identify determines whether file content can be described by this format.
	Identify(io.ReadSeeker) bool
	// Enumerates intrinsic properties associated with files in this format.
	Properties() []PropertyName
}
