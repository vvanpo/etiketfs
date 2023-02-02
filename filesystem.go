package main

// Filesystem ...
type Filesystem struct {
}

// Files returns a reactive selection of all files in the filesystem, and is the
// originating selection of all filter chains. The channel is updated when a
// file is removed or added to the filesystem.
func (fs *Filesystem) Files() <-chan Selection {
	return nil
}
