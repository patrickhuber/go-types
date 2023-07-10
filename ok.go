package types

type Ok[T any] interface {
	Ok() T
}

type ok[T any] struct {
	ok T
	resultImpl[T]
}

func (o *ok[T]) Ok() T {
	return o.ok
}

func (o *ok[T]) Deconstruct() (T, error) {
	return o.ok, nil
}

func (o *ok[T]) IsOk() bool {
	return true
}

func (o *ok[T]) IsError(err ...error) bool {
	return false
}

func (o *ok[T]) Unwrap() T {
	return o.ok
}

func NewOk[T any](value T) Result[T] {
	return &ok[T]{
		ok: value,
	}
}
