package handle

import (
	"github.com/patrickhuber/go-types"
	"github.com/patrickhuber/go-types/result"
)

func Error[T any](res *types.Result[T]) {
	if r := recover(); r != nil {
		if err, ok := r.(error); ok {
			*res = result.Error[T](err)
		}
	}
}
