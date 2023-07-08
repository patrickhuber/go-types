package result_test

import (
	"testing"

	"github.com/patrickhuber/go-types"
	"github.com/patrickhuber/go-types/result"
)

func TestNew(t *testing.T) {
	t.Run("new", func(t *testing.T) {
		switch result.New(1, nil).(type) {
		case types.Error[int]:
			t.Fatalf("expected types.Ok[int]")
		}
	})
	t.Run("func", func(t *testing.T) {
		f := func() (int, error) {
			return 1, nil
		}
		switch result.New(f()).(type) {
		case types.Error[int]:
			t.Fatalf("expected types.Ok[int]")
		}
	})
}
