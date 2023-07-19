package types

// Option represents a value or nothing
type Option[T any] interface {

	// option allows for Some and None to inherit this interface
	option(t T)

	// Deconstruct returns Value, true if the underlying type is Some
	// returns Default, false if the underlying type is None
	Deconstruct() (T, bool)

	// IsSome returns true if the underlying type is Some
	IsSome() bool

	// IsNone returns true if the underlying type is None
	IsNone() bool

	// Unwrap unwraps the Option to its underlying value if the type is Some and
	// panics if the type is None
	Unwrap() T
}
