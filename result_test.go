package types_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/patrickhuber/go-types"
	"github.com/patrickhuber/go-types/option"
	"github.com/patrickhuber/go-types/result"
)

func TestCanMatchOk(t *testing.T) {
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
}

func TestCanMatchError(t *testing.T) {
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
}

func TestCanDeconstructResult(t *testing.T) {
	expect := 123
	result := result.Ok(expect)
	ok, err := result.Deconstruct()
	if err != nil {
		t.Fatal(err)
	}
	if ok != expect {
		t.Fatalf("expected %d found %d", expect, ok)
	}
}

func TestCanDeconstructResultTuple(t *testing.T) {
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
}

func TestCanMatchResultTuple(t *testing.T) {
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
}

func TestOkResultIsOkTrue(t *testing.T) {
	res := result.Ok(123)
	if !res.IsOk() {
		t.Fatalf("expected IsOk to be true")
	}
	if res.IsError() {
		t.Fatalf("expected IsErr to be false")
	}
}

func TestErrorResultIsErrorTrue(t *testing.T) {
	res := result.Error[int](fmt.Errorf("fail"))
	if res.IsOk() {
		t.Fatalf("expected IsOk to be false")
	}
	if !res.IsError() {
		t.Fatalf("expected IsErr to be true")
	}
}
