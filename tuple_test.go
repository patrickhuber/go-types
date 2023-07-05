package types_test

import (
	"testing"

	"github.com/patrickhuber/go-types/tuple"
)

func TestTuple1(t *testing.T) {

	t.Run("deconstruct", func(t *testing.T) {
		expected := 1
		t1 := tuple.New1(expected)
		v := t1.Deconstruct()
		if v != expected {
			t.Fatalf("expected %d found %d", expected, v)
		}
	})

	t.Run("value", func(t *testing.T) {
		expected := 1
		tup := tuple.New1(expected)
		if tup.Value1() != expected {
			t.Fatalf("expected value to equal %d", expected)
		}
	})
}

func TestTuple2(t *testing.T) {
	t.Run("deconstruct", func(t *testing.T) {
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
	})

	t.Run("value", func(t *testing.T) {
		expected1 := 1
		expected2 := 2
		tup := tuple.New2(expected1, expected2)
		if tup.Value1() != expected1 {
			t.Fatalf("expected value to equal %d", expected1)
		}
		if tup.Value2() != expected2 {
			t.Fatalf("expected value to equal %d", expected2)
		}
	})

}

func TestTuple3(t *testing.T) {
	t.Run("deconstruct", func(t *testing.T) {
		expect1 := 1
		expect2 := "expect"
		expect3 := 1.23
		t2 := tuple.New3(expect1, expect2, expect3)
		v1, v2, v3 := t2.Deconstruct()
		if v1 != expect1 {
			t.Fatalf("expected %d found %d", expect1, v1)
		}
		if v2 != expect2 {
			t.Fatalf("expected %s found %s", expect2, v2)
		}
		if v3 != expect3 {
			t.Fatalf("expected %f found %f", expect3, v3)
		}
	})
	t.Run("value", func(t *testing.T) {
		expected := []int{1, 2, 3}
		tup := tuple.New3(expected[0], expected[1], expected[2])
		if tup.Value1() != expected[0] {
			t.Fatalf("expected value to equal %d", expected[0])
		}
		if tup.Value2() != expected[1] {
			t.Fatalf("expected value to equal %d", expected[1])
		}
		if tup.Value3() != expected[2] {
			t.Fatalf("expected value to equal %d", expected[2])
		}
	})
}

func TestTuple4(t *testing.T) {
	t.Run("deconstruct", func(t *testing.T) {
		expect1 := 1
		expect2 := "expect"
		expect3 := 1.23
		expect4 := false
		t4 := tuple.New4(expect1, expect2, expect3, expect4)
		v1, v2, v3, v4 := t4.Deconstruct()
		if v1 != expect1 {
			t.Fatalf("expected %d found %d", expect1, v1)
		}
		if v2 != expect2 {
			t.Fatalf("expected %s found %s", expect2, v2)
		}
		if v3 != expect3 {
			t.Fatalf("expected %f found %f", expect3, v3)
		}
		if v4 != expect4 {
			t.Fatalf("expected %t found %t", expect4, v4)
		}
	})
	t.Run("value", func(t *testing.T) {
		expected := []int{1, 2, 3, 4}
		tup := tuple.New4(expected[0], expected[1], expected[2], expected[3])
		if tup.Value1() != expected[0] {
			t.Fatalf("expected value to equal %d", expected[0])
		}
		if tup.Value2() != expected[1] {
			t.Fatalf("expected value to equal %d", expected[1])
		}
		if tup.Value3() != expected[2] {
			t.Fatalf("expected value to equal %d", expected[2])
		}
		if tup.Value4() != expected[3] {
			t.Fatalf("expected value to equal %d", expected[3])
		}
	})
}
