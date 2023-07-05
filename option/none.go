package option

import "github.com/patrickhuber/go-types"

func None[T any]() types.Option[T] {
	return types.NewNone[T]()
}
