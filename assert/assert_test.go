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

	t.Run("nil", func(t *testing.T) {
		switch run(assert.Nil, nil).(type) {
		case types.Error[any]:
			t.Fatalf("expected Ok")
		}
	})

	t.Run("nil_err", func(t *testing.T) {
		switch run(assert.Nil, 1).(type) {
		case types.Ok[any]:
			t.Fatalf("expected Error")
		}
	})

	t.Run("not_nil", func(t *testing.T) {
		switch run(assert.NotNil, 1).(type) {
		case types.Error[any]:
			t.Fatalf("expected Ok")
		}
	})

	t.Run("not_nil_err", func(t *testing.T) {
		switch run(assert.NotNil, nil).(type) {
		case types.Ok[any]:
			t.Fatalf("expected Error")
		}
	})
}

func run[T any](action func(T), actual T) (res types.Result[any]) {
	defer handle.Error(&res)
	action(actual)
	return
}
