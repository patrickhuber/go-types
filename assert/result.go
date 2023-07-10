package assert

import (
	"fmt"

	"github.com/patrickhuber/go-types"
)

func Okf[T any](res types.Result[T], format string, args ...any) {
	switch res.(type) {
	case types.Ok[T]:
		return
	default:
		panic(fmt.Errorf(format, args...))
	}
}

func Ok[T any](res types.Result[T]) {
	Okf(res, "unable to math Ok[T]")
}

func Errorf[T any](res types.Result[T], format string, args ...any) {
	switch res.(type) {
	case types.Error[T]:
		return
	default:
		panic(fmt.Errorf(format, args...))
	}
}

func Error[T any](res types.Result[T]) {
	Errorf(res, "unable to match Error[T]")
}
