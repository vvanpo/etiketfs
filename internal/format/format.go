package format

import "io"

type Format interface {
	// Identify determines whether file content can be described by this format.
	Identify(io.ReadSeeker) bool
}
