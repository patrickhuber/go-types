package types

import "fmt"

// None represents an Option[T] that does not have a value
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

func (n *noneImpl[T]) IsSome() bool { return false }

func (n *noneImpl[T]) IsNone() bool { return true }

func (n *noneImpl[T]) Unwrap() T {
	panic(fmt.Errorf("unable to match %T", n))
}

func NewNone[T any]() Option[T] {
	return &noneImpl[T]{}
}
