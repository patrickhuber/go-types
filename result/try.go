package result

import (
	"fmt"

	"github.com/patrickhuber/go-types"
)

func Try[T any](res types.Result[T]) T {
	switch t := res.(type) {
	case types.Error[T]:
		panic(t.Error())
	case types.Ok[T]:
		return t.Ok()
	}
	panic(fmt.Errorf("unable to match %T", res))
}

func Handle[T any](res *types.Result[T]) {
	if r := recover(); r != nil {
		*res = Error[T](r.(error))
	}
}
