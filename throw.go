package types

import (
	"errors"
	"fmt"
)

// ErrRecoverable is a marker used when handling panics.
// If the error is wrapped in ErrRecoverable defer handle.Error will recover
// The most common way to use ErrRecoverable is to call the Throw method
var ErrRecoverable = errors.New("")

// Throw wraps the given error with ErrRecoverable and then panics
func Throw(err error) {
	if !errors.Is(err, ErrRecoverable) {
		err = fmt.Errorf("%w %w", err, ErrRecoverable)
	}
	panic(err)
}
