package values

type Formatter[T any] interface {
	Format(T) string
}
