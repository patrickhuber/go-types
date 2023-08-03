package types

type Some[T any] struct {
	Value T
}

// option implements Option[T].option
//
//lint:ignore U1000 method is used to implement Option[T]
func (Some[T]) option(t T) {}

// Deconstruct implements Option[T].Deconstruct()
func (s Some[T]) Deconstruct() (T, bool) {
	return s.Value, true
}

// IsSome implementes Option[T].IsSome
func (s Some[T]) IsSome() bool { return true }

// IsNone implementes Option[T].IsNone
func (s Some[T]) IsNone() bool { return false }

// Unwrap impelements Option[T].Unwrap
func (s Some[T]) Unwrap() T {
	return s.Value
}

// NewSome returns a Some[T] that wraps the value T
func NewSome[T any](value T) Option[T] {
	return Some[T]{
		Value: value,
	}
}
