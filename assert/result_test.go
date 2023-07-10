package assert_test

import (
	"testing"

	"github.com/patrickhuber/go-types/assert"
	"github.com/patrickhuber/go-types/result"
)

func TestResult(t *testing.T) {

	t.Run("ok", func(t *testing.T) {
		run(assert.Ok[int], result.Ok(1))
	})

	t.Run("error", func(t *testing.T) {
		run(assert.Error[int], result.Errorf[int]("error"))
	})
}
