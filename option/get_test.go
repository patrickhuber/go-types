package option_test

import (
	"testing"

	"github.com/patrickhuber/go-types/option"
)

func TestGet(t *testing.T) {
	m := map[string]int{
		"hello": 1,
	}
	t.Run("some", func(t *testing.T) {
		o := option.Get(m, "hello")
		if o.IsNone() {
			t.Fatalf("expected some")
		}
	})
	t.Run("none", func(t *testing.T) {
		o := option.Get(m, "world")
		if o.IsSome() {
			t.Fatalf("expected none")
		}
	})
}
