package option_test

import (
	"testing"

	"github.com/patrickhuber/go-types/option"
)

func TestCast(t *testing.T) {
	expect := 123
	opt := option.Some[any](expect)
	result := option.Cast[any, int](opt)
	o, err := result.Deconstruct()
	if err != nil {
		t.Fatal(err)
	}
	v, ok := o.Deconstruct()
	if !ok {
		t.Fatalf("expected option to contain value")
	}
	if v != expect {
		t.Fatalf("expected %d to equal %d", v, expect)
	}
}
