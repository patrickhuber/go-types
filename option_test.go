package types_test

import (
	"testing"

	"github.com/patrickhuber/go-types"
	"github.com/patrickhuber/go-types/option"
)

func TestNone(t *testing.T) {
	t.Run("deconstruct", func(t *testing.T) {
		_, ok := option.None[int]().Deconstruct()
		if ok {
			t.Fatalf("expected ok to be false")
		}
	})

	t.Run("match", func(t *testing.T) {
		switch option.None[int]().(type) {
		case types.Some[int]:
			t.Fatalf("expected types.None but found types.Some")
		case types.None[int]:
		}
	})

	t.Run("is_none", func(t *testing.T) {
		opt := option.None[int]()
		if opt.IsSome() {
			t.Fatalf("expected IsSome to be false")
		}
		if !opt.IsNone() {
			t.Fatalf("expected IsNone to be true")
		}
	})
}

func TestSome(t *testing.T) {
	t.Run("deconstruct", func(t *testing.T) {
		expected := 1
		v, ok := option.Some[int](expected).Deconstruct()
		if !ok {
			t.Fatalf("expected ok to be true")
		}
		if v != expected {
			t.Fatalf("expected %d to equal %d", v, expected)
		}
	})
	t.Run("match", func(t *testing.T) {
		switch opt := option.Some[int](123).(type) {
		case types.Some[int]:
			if opt.Value != 123 {
				t.Fatalf("expected option value 123 found %d", opt.Value)
			}
		case types.None[int]:
			t.Fatalf("expected types.None but found types.Some")
		}
	})

	t.Run("is_some", func(t *testing.T) {
		opt := option.Some(123)
		if !opt.IsSome() {
			t.Fatalf("expected IsSome to be true")
		}
		if opt.IsNone() {
			t.Fatalf("expected IsNone to be false")
		}
	})

}
