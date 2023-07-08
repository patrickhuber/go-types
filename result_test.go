package types_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/patrickhuber/go-types"
	"github.com/patrickhuber/go-types/option"
	"github.com/patrickhuber/go-types/result"
)

func TestOk(t *testing.T) {
	t.Run("match", func(t *testing.T) {
		switch res := result.Ok(123).(type) {
		case types.Ok[int]:
			if res.Ok() != 123 {
				t.Fatalf("expected 123 but found %d", res.Ok())
			}
		case types.Error[int]:
			t.Fatalf("expected match types.Ok[int] but matched types.Error[int]")
		default:
			t.Fatalf("unable to match %v", res)
		}
	})
	t.Run("deconstruct", func(t *testing.T) {
		expect := 123
		result := result.Ok(expect)
		ok, err := result.Deconstruct()
		if err != nil {
			t.Fatal(err)
		}
		if ok != expect {
			t.Fatalf("expected %d found %d", expect, ok)
		}
	})
	t.Run("is_ok", func(t *testing.T) {
		res := result.Ok(123)
		if !res.IsOk() {
			t.Fatalf("expected IsOk to be true")
		}
		if res.IsError() {
			t.Fatalf("expected IsErr to be false")
		}
	})
}

func TestError(t *testing.T) {
	t.Run("match", func(t *testing.T) {
		myerror := fmt.Errorf("my error")
		switch res := result.Error[int](myerror).(type) {
		case types.Error[int]:
			if !errors.Is(res.Error(), myerror) {
				t.Fatalf("expected %v but found %v", myerror, res.Error())
			}
		case types.Ok[int]:
			t.Fatalf("expected types.Error[int] but found types.Ok[int]")
		default:
			t.Fatalf("unable to match %v", res)
		}
	})
	t.Run("deconstruct", func(t *testing.T) {
		myerr := fmt.Errorf("err")
		result := result.Error[int](myerr)
		_, err := result.Deconstruct()
		if err == nil {
			t.Fatalf("expected err to not be nil")
		}
		if !errors.Is(err, myerr) {
			t.Fatalf("expected err to be myerr")
		}
	})
}

func TestResult(t *testing.T) {
	t.Run("match", func(t *testing.T) {
		expected1 := 123
		switch res := result.New(option.Some(expected1), nil).(type) {
		case types.Ok[types.Option[int]]:
			switch opt := res.Ok().(type) {
			case types.Some[int]:
				found := opt.Value()
				if found != expected1 {
					t.Fatalf("expected %d found %d", expected1, found)
				}
			case types.None[int]:
				t.Fatalf("expected types.Some[int] found types.None[int]")
			}
		case types.Error[types.Option[int]]:
			t.Fatal(res.Error())
		default:
			t.Fatalf("unable to match result")
		}
	})

	t.Run("error", func(t *testing.T) {
		res := result.New[int](0, fmt.Errorf("err"))
		switch res.(type) {
		case types.Ok[int]:
			t.Fatalf("expected types.Error[int]")
		case types.Error[int]:
		default:
			t.Fatalf("expected types.Error[int]")
		}
	})

	t.Run("deconstruct", func(t *testing.T) {
		expected1 := 123
		expected2 := "test"
		success := func() (int, string, error) {
			return expected1, expected2, nil
		}
		result := result.New2(success())
		tup, err := result.Deconstruct()
		if err != nil {
			t.Fatal(err)
		}
		v1, v2 := tup.Deconstruct()
		if v1 != expected1 {
			t.Fatalf("expected %d found %d", expected1, v1)
		}
		if v2 != expected2 {
			t.Fatalf("expected %s found %s", expected2, v2)
		}
	})

	t.Run("is_error", func(t *testing.T) {
		res := result.Error[int](fmt.Errorf("fail"))
		if res.IsOk() {
			t.Fatalf("expected IsOk to be false")
		}
		if !res.IsError() {
			t.Fatalf("expected IsErr to be true")
		}
	})

	t.Run("unwrap", func(t *testing.T) {
		test := func() (res types.Result[int]) {
			defer result.Handle(&res)
			err := result.Error[int](fmt.Errorf("fail"))
			_ = err.Unwrap() // should fail
			return err
		}
		switch test().(type) {
		case types.Ok[int]:
			t.Fatalf("expected types.Error[int]")
		}
	})
}
