package result

import "github.com/patrickhuber/go-types"

func Ok[T any](value T) types.Result[T] {
	return types.NewOk[T](value)
}
