package result_test

import (
	"testing"

	"github.com/patrickhuber/go-types"
	"github.com/patrickhuber/go-types/result"
)

func TestNew(t *testing.T) {
	t.Run("new", func(t *testing.T) {
		switch result.New(1, nil).(type) {
		case types.Error[int]:
			t.Fatalf("expected types.Ok[int]")
		}
	})
	t.Run("func", func(t *testing.T) {
		f := func() (int, error) {
			return 1, nil
		}
		switch result.New(f()).(type) {
		case types.Error[int]:
			t.Fatalf("expected types.Ok[int]")
		}
	})
}

func TestNew2(t *testing.T) {
	t.Run("new", func(t *testing.T) {
		switch result.New2(1, 2, nil).(type) {
		case types.Error[types.Tuple2[int, int]]:
			t.Fatalf("expected types.Ok[types.Tuple2[int, int]]")
		}
	})
	t.Run("func", func(t *testing.T) {
		f := func() (int, int, error) {
			return 1, 2, nil
		}
		switch result.New2(f()).(type) {
		case types.Error[types.Tuple2[int, int]]:
			t.Fatalf("expected types.Ok[Tuple2[int, int]]")
		}
	})
}

func TestNew3(t *testing.T) {
	t.Run("new", func(t *testing.T) {
		switch result.New3(1, 2, 3, nil).(type) {
		case types.Error[types.Tuple3[int, int, int]]:
			t.Fatalf("expected types.Ok[types.Tuple3[int, int, int]]")
		}
	})
	t.Run("func", func(t *testing.T) {
		f := func() (int, int, int, error) {
			return 1, 2, 3, nil
		}
		switch result.New3(f()).(type) {
		case types.Error[types.Tuple3[int, int, int]]:
			t.Fatalf("expected types.Ok[Tuple3[int, int, int]]")
		}
	})
}

func TestNew4(t *testing.T) {
	t.Run("new", func(t *testing.T) {
		switch result.New4(1, 2, 3, 4, nil).(type) {
		case types.Error[types.Tuple4[int, int, int, int]]:
			t.Fatalf("expected types.Ok[types.Tuple4[int, int, int, int]]")
		}
	})
	t.Run("func", func(t *testing.T) {
		f := func() (int, int, int, int, error) {
			return 1, 2, 3, 4, nil
		}
		switch result.New4(f()).(type) {
		case types.Error[types.Tuple4[int, int, int, int]]:
			t.Fatalf("expected types.Ok[Tuple4[int, int, int, int]]")
		}
	})
}

func TestNew5(t *testing.T) {
	t.Run("new", func(t *testing.T) {
		switch result.New5(1, 2, 3, 4, 5, nil).(type) {
		case types.Error[types.Tuple5[int, int, int, int, int]]:
			t.Fatalf("expected types.Ok[types.Tuple5[int, int, int, int, int]]")
		}
	})
	t.Run("func", func(t *testing.T) {
		f := func() (int, int, int, int, int, error) {
			return 1, 2, 3, 4, 5, nil
		}
		switch result.New5(f()).(type) {
		case types.Error[types.Tuple5[int, int, int, int, int]]:
			t.Fatalf("expected types.Ok[Tuple5[int, int, int, int, int]]")
		}
	})
}

func TestNew6(t *testing.T) {
	t.Run("new", func(t *testing.T) {
		switch result.New6(1, 2, 3, 4, 5, 6, nil).(type) {
		case types.Error[types.Tuple6[int, int, int, int, int, int]]:
			t.Fatalf("expected types.Ok[types.Tuple6[int, int, int, int, int, int]]")
		}
	})
	t.Run("func", func(t *testing.T) {
		f := func() (int, int, int, int, int, int, error) {
			return 1, 2, 3, 4, 5, 6, nil
		}
		switch result.New6(f()).(type) {
		case types.Error[types.Tuple6[int, int, int, int, int, int]]:
			t.Fatalf("expected types.Ok[Tuple6[int, int, int, int, int, int]]")
		}
	})
}

func TestNew7(t *testing.T) {
	t.Run("new", func(t *testing.T) {
		switch result.New7(1, 2, 3, 4, 5, 6, 7, nil).(type) {
		case types.Error[types.Tuple7[int, int, int, int, int, int, int]]:
			t.Fatalf("expected types.Ok[types.Tuple7[int, int, int, int, int, int, int]]")
		}
	})
	t.Run("func", func(t *testing.T) {
		f := func() (int, int, int, int, int, int, int, error) {
			return 1, 2, 3, 4, 5, 6, 7, nil
		}
		switch result.New7(f()).(type) {
		case types.Error[types.Tuple7[int, int, int, int, int, int, int]]:
			t.Fatalf("expected types.Ok[Tuple7[int, int, int, int, int, int, int]]")
		}
	})
}

func TestNew8(t *testing.T) {
	t.Run("new", func(t *testing.T) {
		switch result.New8(1, 2, 3, 4, 5, 6, 7, 8, nil).(type) {
		case types.Error[types.Tuple8[int, int, int, int, int, int, int, int]]:
			t.Fatalf("expected types.Ok[types.Tuple8[int, int, int, int, int, int, int, int]]")
		}
	})
	t.Run("func", func(t *testing.T) {
		f := func() (int, int, int, int, int, int, int, int, error) {
			return 1, 2, 3, 4, 5, 6, 7, 8, nil
		}
		switch result.New8(f()).(type) {
		case types.Error[types.Tuple8[int, int, int, int, int, int, int, int]]:
			t.Fatalf("expected types.Ok[Tuple8[int, int, int, int, int, int, int, int]]")
		}
	})
}

func TestNew9(t *testing.T) {
	t.Run("new", func(t *testing.T) {
		switch result.New9(1, 2, 3, 4, 5, 6, 7, 8, 9, nil).(type) {
		case types.Error[types.Tuple9[int, int, int, int, int, int, int, int, int]]:
			t.Fatalf("expected types.Ok[types.Tuple9[int, int, int, int, int, int, int, int, int]]")
		}
	})
	t.Run("func", func(t *testing.T) {
		f := func() (int, int, int, int, int, int, int, int, int, error) {
			return 1, 2, 3, 4, 5, 6, 7, 8, 9, nil
		}
		switch result.New9(f()).(type) {
		case types.Error[types.Tuple9[int, int, int, int, int, int, int, int, int]]:
			t.Fatalf("expected types.Ok[Tuple9[int, int, int, int, int, int, int, int, int]]")
		}
	})
}

func TestNew10(t *testing.T) {
	t.Run("new", func(t *testing.T) {
		switch result.New10(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, nil).(type) {
		case types.Error[types.Tuple10[int, int, int, int, int, int, int, int, int, int]]:
			t.Fatalf("expected types.Ok[types.Tuple10[int, int, int, int, int, int, int, int, int, int]]")
		}
	})
	t.Run("func", func(t *testing.T) {
		f := func() (int, int, int, int, int, int, int, int, int, int, error) {
			return 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, nil
		}
		switch result.New10(f()).(type) {
		case types.Error[types.Tuple10[int, int, int, int, int, int, int, int, int, int]]:
			t.Fatalf("expected types.Ok[Tuple10[int, int, int, int, int, int, int, int, int, int]]")
		}
	})
}

func TestNew11(t *testing.T) {
	t.Run("new", func(t *testing.T) {
		switch result.New11(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, nil).(type) {
		case types.Error[types.Tuple11[int, int, int, int, int, int, int, int, int, int, int]]:
			t.Fatalf("expected types.Ok[types.Tuple11[int, int, int, int, int, int, int, int, int, int, int]]")
		}
	})
	t.Run("func", func(t *testing.T) {
		f := func() (int, int, int, int, int, int, int, int, int, int, int, error) {
			return 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, nil
		}
		switch result.New11(f()).(type) {
		case types.Error[types.Tuple11[int, int, int, int, int, int, int, int, int, int, int]]:
			t.Fatalf("expected types.Ok[Tuple11[int, int, int, int, int, int, int, int, int, int, int]]")
		}
	})
}

func TestNew12(t *testing.T) {
	t.Run("new", func(t *testing.T) {
		switch result.New12(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, nil).(type) {
		case types.Error[types.Tuple12[int, int, int, int, int, int, int, int, int, int, int, int]]:
			t.Fatalf("expected types.Ok[types.Tuple12[int, int, int, int, int, int, int, int, int, int, int, int]]")
		}
	})
	t.Run("func", func(t *testing.T) {
		f := func() (int, int, int, int, int, int, int, int, int, int, int, int, error) {
			return 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, nil
		}
		switch result.New12(f()).(type) {
		case types.Error[types.Tuple12[int, int, int, int, int, int, int, int, int, int, int, int]]:
			t.Fatalf("expected types.Ok[Tuple12[int, int, int, int, int, int, int, int, int, int, int, int]]")
		}
	})
}

func TestNew13(t *testing.T) {
	t.Run("new", func(t *testing.T) {
		switch result.New13(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, nil).(type) {
		case types.Error[types.Tuple13[int, int, int, int, int, int, int, int, int, int, int, int, int]]:
			t.Fatalf("expected types.Ok[types.Tuple13[int, int, int, int, int, int, int, int, int, int, int, int, int]]")
		}
	})
	t.Run("func", func(t *testing.T) {
		f := func() (int, int, int, int, int, int, int, int, int, int, int, int, int, error) {
			return 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, nil
		}
		switch result.New13(f()).(type) {
		case types.Error[types.Tuple13[int, int, int, int, int, int, int, int, int, int, int, int, int]]:
			t.Fatalf("expected types.Ok[Tuple13[int, int, int, int, int, int, int, int, int, int, int, int, int]]")
		}
	})
}

func TestNew14(t *testing.T) {
	t.Run("new", func(t *testing.T) {
		switch result.New14(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, nil).(type) {
		case types.Error[types.Tuple14[int, int, int, int, int, int, int, int, int, int, int, int, int, int]]:
			t.Fatalf("expected types.Ok[types.Tuple14[int, int, int, int, int, int, int, int, int, int, int, int, int, int]]")
		}
	})
	t.Run("func", func(t *testing.T) {
		f := func() (int, int, int, int, int, int, int, int, int, int, int, int, int, int, error) {
			return 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, nil
		}
		switch result.New14(f()).(type) {
		case types.Error[types.Tuple14[int, int, int, int, int, int, int, int, int, int, int, int, int, int]]:
			t.Fatalf("expected types.Ok[Tuple14[int, int, int, int, int, int, int, int, int, int, int, int, int, int]]")
		}
	})
}

func TestNew15(t *testing.T) {
	t.Run("new", func(t *testing.T) {
		switch result.New15(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, nil).(type) {
		case types.Error[types.Tuple15[int, int, int, int, int, int, int, int, int, int, int, int, int, int, int]]:
			t.Fatalf("expected types.Ok[types.Tuple15[int, int, int, int, int, int, int, int, int, int, int, int, int, int, int]]")
		}
	})
	t.Run("func", func(t *testing.T) {
		f := func() (int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, error) {
			return 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, nil
		}
		switch result.New15(f()).(type) {
		case types.Error[types.Tuple15[int, int, int, int, int, int, int, int, int, int, int, int, int, int, int]]:
			t.Fatalf("expected types.Ok[Tuple15[int, int, int, int, int, int, int, int, int, int, int, int, int, int, int]]")
		}
	})
}

func TestNew16(t *testing.T) {
	t.Run("new", func(t *testing.T) {
		switch result.New16(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, nil).(type) {
		case types.Error[types.Tuple16[int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int]]:
			t.Fatalf("expected types.Ok[types.Tuple16[int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int]]")
		}
	})
	t.Run("func", func(t *testing.T) {
		f := func() (int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, error) {
			return 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, nil
		}
		switch result.New16(f()).(type) {
		case types.Error[types.Tuple16[int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int]]:
			t.Fatalf("expected types.Ok[Tuple16[int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int]]")
		}
	})
}
