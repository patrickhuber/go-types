package types

type Result[T any] interface {
	result(t T)
	Deconstruct() (T, error)
	IsOk() bool
	IsError() bool
	Unwrap() T
}

type resultImpl[T any] struct {
}

func (*resultImpl[T]) result(t T) {}

func NewResult[T any](ok T, err error) Result[T] {
	if err != nil {
		return NewError[T](err)
	}
	return NewOk(ok)
}
