package main

func render(fmts formatters, files <-chan []any) <-chan []string {
	out := make(chan []string)

	go func() {
		for f := range files {
			rendered := make([]string, len(f))

			for i, val := range f {
				rendered[i] = fmts.format(val)
			}

			out <- rendered
		}

		close(out)
	}()

	return out
}
