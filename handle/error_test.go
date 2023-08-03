package handle_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/patrickhuber/go-types"
	"github.com/patrickhuber/go-types/handle"
)

func TestError(t *testing.T) {
	t.Run("handle_recoverable", func(t *testing.T) {
		RunError(t, func() (res types.Result[any]) {
			defer handle.Error(&res)
			err := types.NewError[any](fmt.Errorf("test"))
			err.Unwrap() // unwrap is recoverable
			return types.NewOk[any](nil)
		})
	})
	t.Run("panic_unrecoverable", func(t *testing.T) {
		RunError(t, func() (res types.Result[any]) {
			// defers are unraveled in reverse order so this defer
			// will catch the panic of the following defers
			defer func() {
				r := recover()
				if r == nil {
					t.Fatalf("recover returned nil")
				}
				err, ok := r.(error)
				if !ok {
					t.Fatalf("recover returned non error")
				}
				if !errors.Is(err, types.ErrRecoverable) {
					t.Log("error was not recoverable")
				} else {
					t.Log("error was recoverable")
				}
			}()

			// this will re-throw because the error does not contain ErrRecoverable
			defer handle.Error(&res)

			panic(fmt.Errorf("panic"))
		})
	})
}

func RunError[T any](t *testing.T, action func() types.Result[T]) {
	switch action().(type) {
	case types.Error[T]:
	case types.Ok[T]:
		t.Fatal("expected error")
	}
}

func RunOk[T any](t *testing.T, action func() types.Result[T]) {
	switch res := action().(type) {
	case types.Error[T]:
		t.Fatal(res.Value)
	case types.Ok[T]:
	}
}
