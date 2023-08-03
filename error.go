package types

import (
	"errors"
)

type Error[T any] struct {
	Value error
}

//lint:ignore U1000 method is used to implement Result[T]
func (Error[T]) result(t T) {}

// Deconstruct implements Result[T].Deconstruct()
func (e Error[T]) Deconstruct() (T, error) {
	var t T
	return t, e.Value
}

// IsError implements Result[T].IsError()
func (e Error[T]) IsError(err ...error) bool {
	// no filters so match any error
	if len(err) == 0 {
		return true
	}

	// must match one of the filters
	for _, target := range err {
		if errors.Is(e.Value, target) {
			return true
		}
	}
	return false
}

// IsOk implements Result[T].IsOk()
func (e Error[T]) IsOk() bool {
	return false
}

// Unwrap() implements Result[T].Unwrap()
func (e Error[T]) Unwrap() T {
	Throw(e.Value)
	var zero T
	return zero
}

// NewError returns an error over the supplied error
func NewError[T any](value error) Result[T] {
	return Error[T]{
		Value: value,
	}
}
