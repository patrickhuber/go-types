package result

import "github.com/patrickhuber/go-types"

func Error[T any](err error) types.Result[T] {
	return types.NewError[T](err)
}
