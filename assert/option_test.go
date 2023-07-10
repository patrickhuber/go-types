package assert_test

import (
	"testing"

	"github.com/patrickhuber/go-types/assert"
	"github.com/patrickhuber/go-types/option"
)

func TestOption(t *testing.T) {

	t.Run("some", func(t *testing.T) {
		run(assert.Some[int], option.Some(1))
	})

	t.Run("none", func(t *testing.T) {
		run(assert.None[int], option.None[int]())
	})
}
