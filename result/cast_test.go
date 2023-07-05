package result_test

import (
	"testing"

	"github.com/patrickhuber/go-types/result"
)

func TestCanCast(t *testing.T) {
	expected := 123
	res := result.Ok[any](expected)
	casted := result.Cast[any, int](res)
	val, err := casted.Deconstruct()
	if err != nil {
		t.Fatalf("expected err to be nil")
	}
	if val != expected {
		t.Fatalf("expected %d to be %d", val, expected)
	}
}
