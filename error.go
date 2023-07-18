package types

import (
	"errors"
	"fmt"
)

type Error[T any] interface {
	Result[T]
	Error() error
}

type err[T any] struct {
	err error
	//lint:ignore U1000 field is used to implement Result[T]
	resultImpl[T]
}

func (e *err[T]) Error() error {
	return e.err
}

func (e *err[T]) Deconstruct() (T, error) {
	var t T
	return t, e.err
}

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

func (e *err[T]) IsOk() bool {
	return false
}

func (e *err[T]) Unwrap() T {
	panic(fmt.Errorf("unable to match %T", e))
}

func NewError[T any](value error) Result[T] {
	return &err[T]{
		err: value,
	}
}
