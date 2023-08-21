package option_test

import (
	"strconv"
	"testing"

	"github.com/patrickhuber/go-types"
	"github.com/patrickhuber/go-types/option"
)

func TestMap(t *testing.T) {
	t.Run("some", func(t *testing.T) {
		some := option.Some(10)
		someString := option.Map(some, func(i int) string { return strconv.Itoa(i) })
		switch s := someString.(type) {
		case types.Some[string]:
			if s.Value != "10" {
				t.Fatalf("expected some(10) found %s", s.Value)
			}
		case types.None[string]:
			t.Fatalf("expected some(10)")
		default:
			t.Fatalf("expected some(10)")
		}
	})

	t.Run("none", func(t *testing.T) {
		none := option.None[int]()
		noneString := option.Map(none, func(i int) string { return "" })
		switch noneString.(type) {
		case types.Some[string]:
			t.Fatalf("expected none")
		case types.None[string]:
		default:
			t.Fatalf("expected none")
		}
	})
}

func TestMapOr(t *testing.T) {
	t.Run("some", func(t *testing.T) {
		some := option.Some(10)
		s := option.MapOr(some, func(i int) string { return strconv.Itoa(i) }, "11")
		if s != "10" {
			t.Fatalf("expected 10 found %s", s)
		}
	})

	t.Run("none", func(t *testing.T) {
		none := option.None[int]()
		s := option.MapOr(none, func(i int) string { return "" }, "11")
		if s != "11" {
			t.Fatalf("expected 11 found %s", s)
		}
	})
}

func TestMapOrElse(t *testing.T) {
	t.Run("some", func(t *testing.T) {
		some := option.Some(10)
		s := option.MapOrElse(some, func(i int) string { return strconv.Itoa(i) }, func() string { return "11" })
		if s != "10" {
			t.Fatalf("expected 10 found %s", s)
		}
	})

	t.Run("none", func(t *testing.T) {
		none := option.None[int]()
		s := option.MapOrElse(none, func(i int) string { return "" }, func() string { return "11" })
		if s != "11" {
			t.Fatalf("expected 11 found %s", s)
		}
	})
}
