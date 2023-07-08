package result

import (
	"errors"
	"fmt"

	"github.com/patrickhuber/go-types"
)

func Try[T any](res types.Result[T], filters ...error) types.Result[T] {
	switch t := res.(type) {
	case types.Error[T]:
		err := t.Error()
		for _, filter := range filters {
			if errors.Is(err, filter) {
				return res
			}
		}
		panic(err)
	case types.Ok[T]:
		return res
	}
	panic(fmt.Errorf("unable to match %T", res))
}

func Handle[T any](res *types.Result[T]) {
	if r := recover(); r != nil {
		*res = Error[T](r.(error))
	}
}
