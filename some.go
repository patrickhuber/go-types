package types

// Some represents an Option[T] that has a value
type Some[T any] interface {
	// must implement Option[T]
	Option[T]
	// Value returns the underlying value
	Value() T
}

type someImpl[T any] struct {
	value T
}

// option implements Option[T].option
//
//lint:ignore U1000 method is used to implement Option[T]
func (*someImpl[T]) option(t T) {}

// Value implements Some[T].Value()
func (s *someImpl[T]) Value() T {
	return s.value
}

// Deconstruct implements Option[T].Deconstruct()
func (s *someImpl[T]) Deconstruct() (T, bool) {
	return s.value, true
}

// IsSome implementes Option[T].IsSome
func (s *someImpl[T]) IsSome() bool { return true }

// IsNone implementes Option[T].IsNone
func (s *someImpl[T]) IsNone() bool { return false }

// Unwrap impelements Option[T].Unwrap
func (s *someImpl[T]) Unwrap() T {
	return s.value
}

// NewSome returns a Some[T] that wraps the value T
func NewSome[T any](value T) Option[T] {
	return &someImpl[T]{
		value: value,
	}
}
