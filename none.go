package types

type None[T any] interface {
	Option[T]
	none(t T)
}

type noneImpl[T any] struct {
	optionImpl[T]
}

func (*noneImpl[T]) none(t T) {}

func (*noneImpl[T]) Deconstruct() (T, bool) {
	var t T
	return t, false
}

func (s *noneImpl[T]) IsSome() bool { return false }

func (s *noneImpl[T]) IsNone() bool { return true }

func NewNone[T any]() Option[T] {
	return &noneImpl[T]{}
}
