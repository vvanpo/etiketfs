package main

// Selection ...
type Selection []File

func Files(s <-chan Selection) <-chan []File {
	out := make(chan []File)

	return out
}
