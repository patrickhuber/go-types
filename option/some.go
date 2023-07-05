package option

import "github.com/patrickhuber/go-types"

func Some[T any](value T) types.Option[T] {
	return types.NewSome(value)
}
