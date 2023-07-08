package types_test

import (
	"testing"

	"github.com/patrickhuber/go-types"
)

func TestCast(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		var a any = 123
		switch res := types.Cast[any, int](a).(type) {
		case types.Ok[int]:
			if res.Ok() != a.(int) {
				t.Fatalf("expected %d to equal %d", res.Ok(), a)
			}
		case types.Error[int]:
			t.Fatalf("expected types.Ok")
		default:
			t.Fatalf("unable to match type")
		}
	})
}
