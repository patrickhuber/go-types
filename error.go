package types

type Error[T any] interface {
	Result[T]
	Error() error
}

type err[T any] struct {
	err error
	resultImpl[T]
}

func (e *err[T]) Error() error {
	return e.err
}

func (e *err[T]) Deconstruct() (T, error) {
	var t T
	return t, e.err
}

func (e *err[T]) IsError() bool {
	return true
}

func (e *err[T]) IsOk() bool {
	return false
}

func NewError[T any](value error) Result[T] {
	return &err[T]{
		err: value,
	}
}
