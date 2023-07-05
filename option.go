package types

type Option[T any] interface {
	option(t T)
	Deconstruct() (T, bool)
	IsSome() bool
	IsNone() bool
}

type optionImpl[T any] struct {
}

func (*optionImpl[T]) option(t T) {}
