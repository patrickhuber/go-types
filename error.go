package types

import (
	"errors"
)

// An Error[T] is an Result[T] representing an error
type Error[T any] interface {
	// Error[T] is a Result[T]
	Result[T]
	// Error returns the error contained in the Error[T] instance
	Error() error
}

type err[T any] struct {
	err error
}

//lint:ignore U1000 method is used to implement Result[T]
func (*err[T]) result(t T) {}

// Error implements Result[T].Error()
func (e *err[T]) Error() error {
	return e.err
}

// Deconstruct implements Result[T].Deconstruct()
func (e *err[T]) Deconstruct() (T, error) {
	var t T
	return t, e.err
}

// IsError implements Result[T].IsError()
func (e *err[T]) IsError(err ...error) bool {
	// no filters so match any error
	if len(err) == 0 {
		return true
	}

	// must match one of the filters
	for _, target := range err {
		if errors.Is(e.err, target) {
			return true
		}
	}
	return false
}

// IsOk implements Result[T].IsOk()
func (e *err[T]) IsOk() bool {
	return false
}

// Unwrap() implements Result[T].Unwrap()
func (e *err[T]) Unwrap() T {
	Throw(e.err)
	var zero T
	return zero
}

// NewError returns an error over the supplied error
func NewError[T any](value error) Result[T] {
	return &err[T]{
		err: value,
	}
}
