package types

type Type[T any] interface {
	Set(T)
	Get() T
}
