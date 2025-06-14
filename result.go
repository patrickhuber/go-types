package types

// Result represents a value or an error
type Result[T any] interface {
	// result allows inheritance modeling
	result(t T)

	// Deconstruct expands the result to its underlying values.
	// if the result is an error, T contains the zero value.
	Deconstruct() (T, error)

	// IsOk returns true for Ok results, false for Error results.
	IsOk() bool

	// IsError returns false for Ok results, true for Error results.
	// IsError returns true for Error results where the err list is empty
	// IsError returns false for Error results where the err list has at least one error that doesn't match errors.Is
	// IsError returns true for Error results where the err list has at least one error that matches errors.Is
	IsError(err ...error) bool

	// Unwrap attempts to unwrap the result to its value.
	// If the result is an Error, Unwrap will panic.
	// Unwrap is meant to be used with handle.Error(*types.Result)
	Unwrap() T

	// MapError maps the error to another error
	// The function is meant to reduce if err != nil boilerplate and allows the user to add additioanl error information
	MapError(func(error) error) Result[T]
}

// NewResult returns a Error if err is not nil and Ok if err is nil.
// Most consumers of the module will use the `result` package `Ok` and `Error` functions instead.
func NewResult[T any](ok T, err error) Result[T] {
	if err != nil {
		return NewError[T](err)
	}
	return NewOk(ok)
}
