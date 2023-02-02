package main

type Filter func(File) <-chan bool

func (f Filter) Apply(input <-chan Selection) (output <-chan Selection) {
	return
}

type Predicate func(PropertyValue) bool

func MatchPredicate(match PropertyValue) Predicate {
	return func(value PropertyValue) bool {
		return match == value
	}
}

func MakeFilter(i PropertyIdentifier, pred Predicate) Filter {
	return func(f File) <-chan bool {
		out := make(chan bool)
		prop := f.Property(i)

		go func() {
			for value := range prop {
				out <- pred(value)
			}

			close(out)
		}()

		return out
	}
}
