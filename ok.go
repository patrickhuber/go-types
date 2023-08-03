package types

type Ok[T any] struct {
	Value T
}

//lint:ignore U1000 method is used to implement Result[T]
func (Ok[T]) result(t T) {}

// Deconstruct implements Result[T].Deconstruct
func (o Ok[T]) Deconstruct() (T, error) {
	return o.Value, nil
}

// IsOk implements Result[T].IsOk
func (o Ok[T]) IsOk() bool {
	return true
}

// IsError implements Result[T].IsError
func (o Ok[T]) IsError(err ...error) bool {
	return false
}

// Unwrap implements Result[T].Unwrap
func (o Ok[T]) Unwrap() T {
	return o.Value
}

// NewOk returns an Ok[T] that wraps the given value
func NewOk[T any](value T) Result[T] {
	return Ok[T]{
		Value: value,
	}
}
