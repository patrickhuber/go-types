package handle

import (
	"errors"

	"github.com/patrickhuber/go-types"
	"github.com/patrickhuber/go-types/result"
)

// Error[T] recovers from panic only where the error wraps types.ErrRecoverable.
// this function will not handle system panics, and allows them to propigate upward
func Error[T any](res *types.Result[T]) {

	// attempt to recover from a panic
	r := recover()

	// the function was deferred but there was no panic
	// exit early
	if r == nil {
		return
	}

	// if recovery is not an error, panic
	// this may be a result of the system panicing in which case we want to
	// let the panic propigate upward
	err, ok := r.(error)
	if !ok {
		panic(r)
	}

	// panic if the error does not wrap types.ErrRecoverable
	// this may be a result of the system panicing in which case we want to
	// let the panic propigate upward
	if !errors.Is(err, types.ErrRecoverable) {
		panic(err)
	}

	// set the result to the error
	*res = result.Error[T](err)
}
