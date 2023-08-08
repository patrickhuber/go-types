package types

import "fmt"

// None represents an Option[T] that does not have a value
type None[T any] struct {
}

// option implements Option[T].option(t T)
//
//lint:ignore U1000 method is used to implement Option[T]
func (None[T]) option(t T) {}

// Deconstruct implements Option[T].Deconstruct
func (None[T]) Deconstruct() (T, bool) {
	var t T
	return t, false
}

// IsSome implements Option[T].IsSome
func (n None[T]) IsSome() bool { return false }

// IsNone implements Option[T].IsNone
func (n None[T]) IsNone() bool { return true }

// Unwrap implements Option[T].Unwrap
func (n None[T]) Unwrap() T {
	Throw(fmt.Errorf("types.%T upwrapped", n))
	var zero T
	return zero
}

// UnwrapOr implements Option[T].UnwrapOr
func (n None[T]) UnwrapOr(other T) T {
	return other
}

// NewNone returns a None[T]
func NewNone[T any]() Option[T] {
	return None[T]{}
}
