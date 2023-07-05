package types_test

import (
	"testing"

	"github.com/patrickhuber/go-types"
	"github.com/patrickhuber/go-types/option"
)

func TestCanMatchNone(t *testing.T) {
	switch option.None[int]().(type) {
	case types.Some[int]:
		t.Fatalf("expected types.None but found types.Some")
	case types.None[int]:
	}
}

func TestCanMatchSomeOption(t *testing.T) {
	switch opt := option.Some[int](123).(type) {
	case types.Some[int]:
		if opt.Value() != 123 {
			t.Fatalf("expected option value 123 found %d", opt.Value())
		}
	case types.None[int]:
		t.Fatalf("expected types.None but found types.Some")
	}
}

func TestSomeOptionIsSomeTrue(t *testing.T) {
	opt := option.Some(123)
	if !opt.IsSome() {
		t.Fatalf("expected IsSome to be true")
	}
	if opt.IsNone() {
		t.Fatalf("expected IsNone to be false")
	}
}

func TestNoneOptionIsNoneTrue(t *testing.T) {
	opt := option.None[int]()
	if opt.IsSome() {
		t.Fatalf("expected IsSome to be false")
	}
	if !opt.IsNone() {
		t.Fatalf("expected IsNone to be true")
	}
}
