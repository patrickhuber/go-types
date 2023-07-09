package types

type Some[T any] interface {
	Value() T
}

type someImpl[T any] struct {
	optionImpl[T]
	value T
}

func (s *someImpl[T]) Value() T {
	return s.value
}

func (s *someImpl[T]) Deconstruct() (T, bool) {
	return s.value, true
}

func (s *someImpl[T]) IsSome() bool { return true }
func (s *someImpl[T]) IsNone() bool { return false }
func (s *someImpl[T]) Unwrap() T {
	return s.value
}

func NewSome[T any](value T) Option[T] {
	return &someImpl[T]{
		value: value,
	}
}
