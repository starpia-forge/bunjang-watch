package filter

type Filter[T any] interface {
	Apply(T) bool
}
