package types

// A Ok[T] is an Result[T] representing a success
type Ok[T any] interface {
	// Ok[T] is a Result[T]
	Result[T]
	// Returns the value in the Ok[T] instance
	Ok() T
}

type ok[T any] struct {
	ok T
}

//lint:ignore U1000 method is used to implement Result[T]
func (*ok[T]) result(t T) {}

// Ok implements Ok[T].Ok
func (o *ok[T]) Ok() T {
	return o.ok
}

// Deconstruct implements Result[T].Deconstruct
func (o *ok[T]) Deconstruct() (T, error) {
	return o.ok, nil
}

// IsOk implements Result[T].IsOk
func (o *ok[T]) IsOk() bool {
	return true
}

// IsError implements Result[T].IsError
func (o *ok[T]) IsError(err ...error) bool {
	return false
}

// Unwrap implements Result[T].Unwrap
func (o *ok[T]) Unwrap() T {
	return o.ok
}

// NewOk returns an Ok[T] that wraps the given value
func NewOk[T any](value T) Result[T] {
	return &ok[T]{
		ok: value,
	}
}
