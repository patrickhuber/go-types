package types

import "fmt"

type None[T any] interface {
	Option[T]
	none(t T)
}

type noneImpl[T any] struct {
	//lint:ignore U1000 field is used to implement Option[T]
	optionImpl[T]
}

//lint:ignore U1000 method is used to implement None[T]
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
