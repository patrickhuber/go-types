package types

type Option[T any] interface {
	option(t T)
	Deconstruct() (T, bool)
	IsSome() bool
	IsNone() bool
	Unwrap() T
}

type optionImpl[T any] struct {
}

//lint:ignore U1000 method is used to implement Option[T]
func (*optionImpl[T]) option(t T) {}
