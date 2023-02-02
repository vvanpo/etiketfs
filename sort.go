package main

import "sort"

func Sort(in <-chan []File, cmp func(i, j File) bool) <-chan []File {
	out := make(chan []File)

	go func() {
		for files := range in {
			sort.Slice(files, func(i, j int) bool {
				return cmp(files[i], files[j])
			})

			out <- files
		}

		close(out)
	}()

	return out
}
