package result_test

import (
	"fmt"
	"testing"

	"github.com/patrickhuber/go-types"
	"github.com/patrickhuber/go-types/result"
)

func TestCanTryHandle(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		switch Use(result.Ok(1)).(type) {
		case types.Error[int]:
			t.Fatalf("expected types.Ok")
		}
	})

	t.Run("error", func(t *testing.T) {
		switch Use(result.Error[int](fmt.Errorf("failed"))).(type) {
		case types.Ok[int]:
			t.Fatalf("expected types.Error")
		}
	})
}

func Use[T any](in types.Result[T]) (res types.Result[T]) {
	defer result.Handle(&res)
	t := result.Try(in)
	fmt.Println(t)
	return
}
