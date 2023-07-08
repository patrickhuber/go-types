package assert_test

import (
	"testing"

	"github.com/patrickhuber/go-types"
	"github.com/patrickhuber/go-types/assert"
	"github.com/patrickhuber/go-types/handle"
)

func TestAssert(t *testing.T) {
	t.Run("true_when_true", func(t *testing.T) {
		switch run(assert.True, true).(type) {
		case types.Error[any]:
			t.Fatalf("expected Ok")
		}
	})
	t.Run("true_when_false", func(t *testing.T) {
		switch run(assert.True, false).(type) {
		case types.Ok[any]:
			t.Fatalf("expected Error")
		}
	})
	t.Run("false_when_false", func(t *testing.T) {
		switch run(assert.False, false).(type) {
		case types.Error[any]:
			t.Fatalf("expected Ok")
		}
	})
	t.Run("false_when_true", func(t *testing.T) {
		switch run(assert.False, true).(type) {
		case types.Ok[any]:
			t.Fatalf("expected Error")
		}
	})
}

func run(action func(bool), condition bool) (res types.Result[any]) {
	defer handle.Error(&res)
	action(condition)
	return
}
