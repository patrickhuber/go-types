package types_test

import (
	"testing"

	"github.com/patrickhuber/go-types/tuple"
)

func TestCanDeconstructTuple1(t *testing.T) {
	expected := 1
	t1 := tuple.New1(expected)
	v := t1.Deconstruct()
	if v != expected {
		t.Fatalf("expected %d found %d", expected, v)
	}
}

func TestCanDeconstructTuple2(t *testing.T) {
	expect1 := 1
	expect2 := "expect"
	t2 := tuple.New2(expect1, expect2)
	v1, v2 := t2.Deconstruct()
	if v1 != expect1 {
		t.Fatalf("expected %d found %d", expect1, v1)
	}
	if v2 != expect2 {
		t.Fatalf("expected %s found %s", expect2, v2)
	}
}
