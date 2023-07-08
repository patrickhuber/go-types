package result

import (
	"fmt"

	"github.com/patrickhuber/go-types"
)

// Error wraps the error as a types.Error[T]
func Error[T any](err error) types.Result[T] {
	return types.NewError[T](err)
}

// Errorf wraps fmt.Errorf and returns a types.Result[T]
func Errorf[T any](format string, args ...any) types.Result[T] {
	return Error[T](fmt.Errorf(format, args...))
}
