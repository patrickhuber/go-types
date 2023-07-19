package types

import "fmt"

// None represents an Option[T] that does not have a value
type None[T any] interface {
	Option[T]
	none(t T)
}

type noneImpl[T any] struct{}

// option implements Option[T].option(t T)
//
//lint:ignore U1000 method is used to implement Option[T]
func (*noneImpl[T]) option(t T) {}

// option implements None[T].none(t T)
//
//lint:ignore U1000 method is used to implement None[T]
func (*noneImpl[T]) none(t T) {}

// Deconstruct implements Option[T].Deconstruct
func (*noneImpl[T]) Deconstruct() (T, bool) {
	var t T
	return t, false
}

// IsSome implements Option[T].IsSome
func (n *noneImpl[T]) IsSome() bool { return false }

// IsNone implements Option[T].IsNone
func (n *noneImpl[T]) IsNone() bool { return true }

// Unwrap implements Option[T].Unwrap
func (n *noneImpl[T]) Unwrap() T {
	Throw(fmt.Errorf("types.%T upwrapped", n))
	var zero T
	return zero
}

// NewNone returns a None[T]
func NewNone[T any]() Option[T] {
	return &noneImpl[T]{}
}
