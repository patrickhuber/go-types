package assert

import (
	"fmt"

	"github.com/patrickhuber/go-types"
)

func Some[T any](opt types.Option[T]) {
	Somef(opt, "expected some")
}

func Somef[T any](opt types.Option[T], format string, args ...any) {
	switch opt.(type) {
	case types.Some[T]:
		return
	default:
		panic(fmt.Errorf(format, args...))
	}
}

func None[T any](opt types.Option[T]) {
	Nonef(opt, "expected none")
}

func Nonef[T any](opt types.Option[T], format string, args ...any) {
	switch opt.(type) {
	case types.None[T]:
		return
	default:
		panic(fmt.Errorf(format, args...))
	}
}
