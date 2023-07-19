package assert

import (
	"fmt"

	"github.com/patrickhuber/go-types"
)

// Okf[T] casts the the types.Result[T] to types.Ok[T].
// If the cast fails, it panics with an ErrRecoverable formatted with fmt.Errorf
func Okf[T any](res types.Result[T], format string, args ...any) {
	switch res.(type) {
	case types.Ok[T]:
		return
	default:
		types.Throw(fmt.Errorf(format, args...))
	}
}

// Ok[T] casts the the types.Result[T] to types.Ok[T].
// If the cast fails, it panics with an ErrRecoverable
func Ok[T any](res types.Result[T]) {
	Okf(res, "unable to math Ok[T]")
}

// Errorf[T] casts the the types.Result[T] to types.Error[T].
// If the cast fails, it panics with an ErrRecoverable formatted with fmt.Errorf
func Errorf[T any](res types.Result[T], format string, args ...any) {
	switch res.(type) {
	case types.Error[T]:
		return
	default:
		types.Throw(fmt.Errorf(format, args...))
	}
}

// Error[T] casts the the types.Result[T] to types.Error[T].
// If the cast fails, it panics with an ErrRecoverable
func Error[T any](res types.Result[T]) {
	Errorf(res, "unable to match Error[T]")
}
