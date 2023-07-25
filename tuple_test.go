// this file was generated by "go generate" please do not edit
package types_test

import (
	"testing"

	"github.com/patrickhuber/go-types"
)

func TestTuple2(t *testing.T) {
	t.Run("deconstruct", func(t *testing.T) {
		expect1 := 0
		expect2 := 1
		t2 := types.NewTuple2(expect1, expect2)
		v1, v2 := t2.Deconstruct()
		
		if v1 != expect1 {
			t.Fatalf("expected %d found %d", expect1, v1)
		}
		if v2 != expect2 {
			t.Fatalf("expected %d found %d", expect2, v2)
		}		
	})

	t.Run("value", func(t *testing.T) {
		expect1 := 0
		expect2 := 1
		tup := types.NewTuple2(expect1, expect2)
		
		if tup.Value1() != expect1 {
			t.Fatalf("expected %d to equal %d", tup.Value1(), expect1)
		}
		if tup.Value2() != expect2 {
			t.Fatalf("expected %d to equal %d", tup.Value2(), expect2)
		}	
	})
}

func TestTuple3(t *testing.T) {
	t.Run("deconstruct", func(t *testing.T) {
		expect1 := 0
		expect2 := 1
		expect3 := 2
		t3 := types.NewTuple3(expect1, expect2, expect3)
		v1, v2, v3 := t3.Deconstruct()
		
		if v1 != expect1 {
			t.Fatalf("expected %d found %d", expect1, v1)
		}
		if v2 != expect2 {
			t.Fatalf("expected %d found %d", expect2, v2)
		}
		if v3 != expect3 {
			t.Fatalf("expected %d found %d", expect3, v3)
		}		
	})

	t.Run("value", func(t *testing.T) {
		expect1 := 0
		expect2 := 1
		expect3 := 2
		tup := types.NewTuple3(expect1, expect2, expect3)
		
		if tup.Value1() != expect1 {
			t.Fatalf("expected %d to equal %d", tup.Value1(), expect1)
		}
		if tup.Value2() != expect2 {
			t.Fatalf("expected %d to equal %d", tup.Value2(), expect2)
		}
		if tup.Value3() != expect3 {
			t.Fatalf("expected %d to equal %d", tup.Value3(), expect3)
		}	
	})
}

func TestTuple4(t *testing.T) {
	t.Run("deconstruct", func(t *testing.T) {
		expect1 := 0
		expect2 := 1
		expect3 := 2
		expect4 := 3
		t4 := types.NewTuple4(expect1, expect2, expect3, expect4)
		v1, v2, v3, v4 := t4.Deconstruct()
		
		if v1 != expect1 {
			t.Fatalf("expected %d found %d", expect1, v1)
		}
		if v2 != expect2 {
			t.Fatalf("expected %d found %d", expect2, v2)
		}
		if v3 != expect3 {
			t.Fatalf("expected %d found %d", expect3, v3)
		}
		if v4 != expect4 {
			t.Fatalf("expected %d found %d", expect4, v4)
		}		
	})

	t.Run("value", func(t *testing.T) {
		expect1 := 0
		expect2 := 1
		expect3 := 2
		expect4 := 3
		tup := types.NewTuple4(expect1, expect2, expect3, expect4)
		
		if tup.Value1() != expect1 {
			t.Fatalf("expected %d to equal %d", tup.Value1(), expect1)
		}
		if tup.Value2() != expect2 {
			t.Fatalf("expected %d to equal %d", tup.Value2(), expect2)
		}
		if tup.Value3() != expect3 {
			t.Fatalf("expected %d to equal %d", tup.Value3(), expect3)
		}
		if tup.Value4() != expect4 {
			t.Fatalf("expected %d to equal %d", tup.Value4(), expect4)
		}	
	})
}

func TestTuple5(t *testing.T) {
	t.Run("deconstruct", func(t *testing.T) {
		expect1 := 0
		expect2 := 1
		expect3 := 2
		expect4 := 3
		expect5 := 4
		t5 := types.NewTuple5(expect1, expect2, expect3, expect4, expect5)
		v1, v2, v3, v4, v5 := t5.Deconstruct()
		
		if v1 != expect1 {
			t.Fatalf("expected %d found %d", expect1, v1)
		}
		if v2 != expect2 {
			t.Fatalf("expected %d found %d", expect2, v2)
		}
		if v3 != expect3 {
			t.Fatalf("expected %d found %d", expect3, v3)
		}
		if v4 != expect4 {
			t.Fatalf("expected %d found %d", expect4, v4)
		}
		if v5 != expect5 {
			t.Fatalf("expected %d found %d", expect5, v5)
		}		
	})

	t.Run("value", func(t *testing.T) {
		expect1 := 0
		expect2 := 1
		expect3 := 2
		expect4 := 3
		expect5 := 4
		tup := types.NewTuple5(expect1, expect2, expect3, expect4, expect5)
		
		if tup.Value1() != expect1 {
			t.Fatalf("expected %d to equal %d", tup.Value1(), expect1)
		}
		if tup.Value2() != expect2 {
			t.Fatalf("expected %d to equal %d", tup.Value2(), expect2)
		}
		if tup.Value3() != expect3 {
			t.Fatalf("expected %d to equal %d", tup.Value3(), expect3)
		}
		if tup.Value4() != expect4 {
			t.Fatalf("expected %d to equal %d", tup.Value4(), expect4)
		}
		if tup.Value5() != expect5 {
			t.Fatalf("expected %d to equal %d", tup.Value5(), expect5)
		}	
	})
}

func TestTuple6(t *testing.T) {
	t.Run("deconstruct", func(t *testing.T) {
		expect1 := 0
		expect2 := 1
		expect3 := 2
		expect4 := 3
		expect5 := 4
		expect6 := 5
		t6 := types.NewTuple6(expect1, expect2, expect3, expect4, expect5, expect6)
		v1, v2, v3, v4, v5, v6 := t6.Deconstruct()
		
		if v1 != expect1 {
			t.Fatalf("expected %d found %d", expect1, v1)
		}
		if v2 != expect2 {
			t.Fatalf("expected %d found %d", expect2, v2)
		}
		if v3 != expect3 {
			t.Fatalf("expected %d found %d", expect3, v3)
		}
		if v4 != expect4 {
			t.Fatalf("expected %d found %d", expect4, v4)
		}
		if v5 != expect5 {
			t.Fatalf("expected %d found %d", expect5, v5)
		}
		if v6 != expect6 {
			t.Fatalf("expected %d found %d", expect6, v6)
		}		
	})

	t.Run("value", func(t *testing.T) {
		expect1 := 0
		expect2 := 1
		expect3 := 2
		expect4 := 3
		expect5 := 4
		expect6 := 5
		tup := types.NewTuple6(expect1, expect2, expect3, expect4, expect5, expect6)
		
		if tup.Value1() != expect1 {
			t.Fatalf("expected %d to equal %d", tup.Value1(), expect1)
		}
		if tup.Value2() != expect2 {
			t.Fatalf("expected %d to equal %d", tup.Value2(), expect2)
		}
		if tup.Value3() != expect3 {
			t.Fatalf("expected %d to equal %d", tup.Value3(), expect3)
		}
		if tup.Value4() != expect4 {
			t.Fatalf("expected %d to equal %d", tup.Value4(), expect4)
		}
		if tup.Value5() != expect5 {
			t.Fatalf("expected %d to equal %d", tup.Value5(), expect5)
		}
		if tup.Value6() != expect6 {
			t.Fatalf("expected %d to equal %d", tup.Value6(), expect6)
		}	
	})
}

func TestTuple7(t *testing.T) {
	t.Run("deconstruct", func(t *testing.T) {
		expect1 := 0
		expect2 := 1
		expect3 := 2
		expect4 := 3
		expect5 := 4
		expect6 := 5
		expect7 := 6
		t7 := types.NewTuple7(expect1, expect2, expect3, expect4, expect5, expect6, expect7)
		v1, v2, v3, v4, v5, v6, v7 := t7.Deconstruct()
		
		if v1 != expect1 {
			t.Fatalf("expected %d found %d", expect1, v1)
		}
		if v2 != expect2 {
			t.Fatalf("expected %d found %d", expect2, v2)
		}
		if v3 != expect3 {
			t.Fatalf("expected %d found %d", expect3, v3)
		}
		if v4 != expect4 {
			t.Fatalf("expected %d found %d", expect4, v4)
		}
		if v5 != expect5 {
			t.Fatalf("expected %d found %d", expect5, v5)
		}
		if v6 != expect6 {
			t.Fatalf("expected %d found %d", expect6, v6)
		}
		if v7 != expect7 {
			t.Fatalf("expected %d found %d", expect7, v7)
		}		
	})

	t.Run("value", func(t *testing.T) {
		expect1 := 0
		expect2 := 1
		expect3 := 2
		expect4 := 3
		expect5 := 4
		expect6 := 5
		expect7 := 6
		tup := types.NewTuple7(expect1, expect2, expect3, expect4, expect5, expect6, expect7)
		
		if tup.Value1() != expect1 {
			t.Fatalf("expected %d to equal %d", tup.Value1(), expect1)
		}
		if tup.Value2() != expect2 {
			t.Fatalf("expected %d to equal %d", tup.Value2(), expect2)
		}
		if tup.Value3() != expect3 {
			t.Fatalf("expected %d to equal %d", tup.Value3(), expect3)
		}
		if tup.Value4() != expect4 {
			t.Fatalf("expected %d to equal %d", tup.Value4(), expect4)
		}
		if tup.Value5() != expect5 {
			t.Fatalf("expected %d to equal %d", tup.Value5(), expect5)
		}
		if tup.Value6() != expect6 {
			t.Fatalf("expected %d to equal %d", tup.Value6(), expect6)
		}
		if tup.Value7() != expect7 {
			t.Fatalf("expected %d to equal %d", tup.Value7(), expect7)
		}	
	})
}

func TestTuple8(t *testing.T) {
	t.Run("deconstruct", func(t *testing.T) {
		expect1 := 0
		expect2 := 1
		expect3 := 2
		expect4 := 3
		expect5 := 4
		expect6 := 5
		expect7 := 6
		expect8 := 7
		t8 := types.NewTuple8(expect1, expect2, expect3, expect4, expect5, expect6, expect7, expect8)
		v1, v2, v3, v4, v5, v6, v7, v8 := t8.Deconstruct()
		
		if v1 != expect1 {
			t.Fatalf("expected %d found %d", expect1, v1)
		}
		if v2 != expect2 {
			t.Fatalf("expected %d found %d", expect2, v2)
		}
		if v3 != expect3 {
			t.Fatalf("expected %d found %d", expect3, v3)
		}
		if v4 != expect4 {
			t.Fatalf("expected %d found %d", expect4, v4)
		}
		if v5 != expect5 {
			t.Fatalf("expected %d found %d", expect5, v5)
		}
		if v6 != expect6 {
			t.Fatalf("expected %d found %d", expect6, v6)
		}
		if v7 != expect7 {
			t.Fatalf("expected %d found %d", expect7, v7)
		}
		if v8 != expect8 {
			t.Fatalf("expected %d found %d", expect8, v8)
		}		
	})

	t.Run("value", func(t *testing.T) {
		expect1 := 0
		expect2 := 1
		expect3 := 2
		expect4 := 3
		expect5 := 4
		expect6 := 5
		expect7 := 6
		expect8 := 7
		tup := types.NewTuple8(expect1, expect2, expect3, expect4, expect5, expect6, expect7, expect8)
		
		if tup.Value1() != expect1 {
			t.Fatalf("expected %d to equal %d", tup.Value1(), expect1)
		}
		if tup.Value2() != expect2 {
			t.Fatalf("expected %d to equal %d", tup.Value2(), expect2)
		}
		if tup.Value3() != expect3 {
			t.Fatalf("expected %d to equal %d", tup.Value3(), expect3)
		}
		if tup.Value4() != expect4 {
			t.Fatalf("expected %d to equal %d", tup.Value4(), expect4)
		}
		if tup.Value5() != expect5 {
			t.Fatalf("expected %d to equal %d", tup.Value5(), expect5)
		}
		if tup.Value6() != expect6 {
			t.Fatalf("expected %d to equal %d", tup.Value6(), expect6)
		}
		if tup.Value7() != expect7 {
			t.Fatalf("expected %d to equal %d", tup.Value7(), expect7)
		}
		if tup.Value8() != expect8 {
			t.Fatalf("expected %d to equal %d", tup.Value8(), expect8)
		}	
	})
}

func TestTuple9(t *testing.T) {
	t.Run("deconstruct", func(t *testing.T) {
		expect1 := 0
		expect2 := 1
		expect3 := 2
		expect4 := 3
		expect5 := 4
		expect6 := 5
		expect7 := 6
		expect8 := 7
		expect9 := 8
		t9 := types.NewTuple9(expect1, expect2, expect3, expect4, expect5, expect6, expect7, expect8, expect9)
		v1, v2, v3, v4, v5, v6, v7, v8, v9 := t9.Deconstruct()
		
		if v1 != expect1 {
			t.Fatalf("expected %d found %d", expect1, v1)
		}
		if v2 != expect2 {
			t.Fatalf("expected %d found %d", expect2, v2)
		}
		if v3 != expect3 {
			t.Fatalf("expected %d found %d", expect3, v3)
		}
		if v4 != expect4 {
			t.Fatalf("expected %d found %d", expect4, v4)
		}
		if v5 != expect5 {
			t.Fatalf("expected %d found %d", expect5, v5)
		}
		if v6 != expect6 {
			t.Fatalf("expected %d found %d", expect6, v6)
		}
		if v7 != expect7 {
			t.Fatalf("expected %d found %d", expect7, v7)
		}
		if v8 != expect8 {
			t.Fatalf("expected %d found %d", expect8, v8)
		}
		if v9 != expect9 {
			t.Fatalf("expected %d found %d", expect9, v9)
		}		
	})

	t.Run("value", func(t *testing.T) {
		expect1 := 0
		expect2 := 1
		expect3 := 2
		expect4 := 3
		expect5 := 4
		expect6 := 5
		expect7 := 6
		expect8 := 7
		expect9 := 8
		tup := types.NewTuple9(expect1, expect2, expect3, expect4, expect5, expect6, expect7, expect8, expect9)
		
		if tup.Value1() != expect1 {
			t.Fatalf("expected %d to equal %d", tup.Value1(), expect1)
		}
		if tup.Value2() != expect2 {
			t.Fatalf("expected %d to equal %d", tup.Value2(), expect2)
		}
		if tup.Value3() != expect3 {
			t.Fatalf("expected %d to equal %d", tup.Value3(), expect3)
		}
		if tup.Value4() != expect4 {
			t.Fatalf("expected %d to equal %d", tup.Value4(), expect4)
		}
		if tup.Value5() != expect5 {
			t.Fatalf("expected %d to equal %d", tup.Value5(), expect5)
		}
		if tup.Value6() != expect6 {
			t.Fatalf("expected %d to equal %d", tup.Value6(), expect6)
		}
		if tup.Value7() != expect7 {
			t.Fatalf("expected %d to equal %d", tup.Value7(), expect7)
		}
		if tup.Value8() != expect8 {
			t.Fatalf("expected %d to equal %d", tup.Value8(), expect8)
		}
		if tup.Value9() != expect9 {
			t.Fatalf("expected %d to equal %d", tup.Value9(), expect9)
		}	
	})
}

func TestTuple10(t *testing.T) {
	t.Run("deconstruct", func(t *testing.T) {
		expect1 := 0
		expect2 := 1
		expect3 := 2
		expect4 := 3
		expect5 := 4
		expect6 := 5
		expect7 := 6
		expect8 := 7
		expect9 := 8
		expect10 := 9
		t10 := types.NewTuple10(expect1, expect2, expect3, expect4, expect5, expect6, expect7, expect8, expect9, expect10)
		v1, v2, v3, v4, v5, v6, v7, v8, v9, v10 := t10.Deconstruct()
		
		if v1 != expect1 {
			t.Fatalf("expected %d found %d", expect1, v1)
		}
		if v2 != expect2 {
			t.Fatalf("expected %d found %d", expect2, v2)
		}
		if v3 != expect3 {
			t.Fatalf("expected %d found %d", expect3, v3)
		}
		if v4 != expect4 {
			t.Fatalf("expected %d found %d", expect4, v4)
		}
		if v5 != expect5 {
			t.Fatalf("expected %d found %d", expect5, v5)
		}
		if v6 != expect6 {
			t.Fatalf("expected %d found %d", expect6, v6)
		}
		if v7 != expect7 {
			t.Fatalf("expected %d found %d", expect7, v7)
		}
		if v8 != expect8 {
			t.Fatalf("expected %d found %d", expect8, v8)
		}
		if v9 != expect9 {
			t.Fatalf("expected %d found %d", expect9, v9)
		}
		if v10 != expect10 {
			t.Fatalf("expected %d found %d", expect10, v10)
		}		
	})

	t.Run("value", func(t *testing.T) {
		expect1 := 0
		expect2 := 1
		expect3 := 2
		expect4 := 3
		expect5 := 4
		expect6 := 5
		expect7 := 6
		expect8 := 7
		expect9 := 8
		expect10 := 9
		tup := types.NewTuple10(expect1, expect2, expect3, expect4, expect5, expect6, expect7, expect8, expect9, expect10)
		
		if tup.Value1() != expect1 {
			t.Fatalf("expected %d to equal %d", tup.Value1(), expect1)
		}
		if tup.Value2() != expect2 {
			t.Fatalf("expected %d to equal %d", tup.Value2(), expect2)
		}
		if tup.Value3() != expect3 {
			t.Fatalf("expected %d to equal %d", tup.Value3(), expect3)
		}
		if tup.Value4() != expect4 {
			t.Fatalf("expected %d to equal %d", tup.Value4(), expect4)
		}
		if tup.Value5() != expect5 {
			t.Fatalf("expected %d to equal %d", tup.Value5(), expect5)
		}
		if tup.Value6() != expect6 {
			t.Fatalf("expected %d to equal %d", tup.Value6(), expect6)
		}
		if tup.Value7() != expect7 {
			t.Fatalf("expected %d to equal %d", tup.Value7(), expect7)
		}
		if tup.Value8() != expect8 {
			t.Fatalf("expected %d to equal %d", tup.Value8(), expect8)
		}
		if tup.Value9() != expect9 {
			t.Fatalf("expected %d to equal %d", tup.Value9(), expect9)
		}
		if tup.Value10() != expect10 {
			t.Fatalf("expected %d to equal %d", tup.Value10(), expect10)
		}	
	})
}

func TestTuple11(t *testing.T) {
	t.Run("deconstruct", func(t *testing.T) {
		expect1 := 0
		expect2 := 1
		expect3 := 2
		expect4 := 3
		expect5 := 4
		expect6 := 5
		expect7 := 6
		expect8 := 7
		expect9 := 8
		expect10 := 9
		expect11 := 10
		t11 := types.NewTuple11(expect1, expect2, expect3, expect4, expect5, expect6, expect7, expect8, expect9, expect10, expect11)
		v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11 := t11.Deconstruct()
		
		if v1 != expect1 {
			t.Fatalf("expected %d found %d", expect1, v1)
		}
		if v2 != expect2 {
			t.Fatalf("expected %d found %d", expect2, v2)
		}
		if v3 != expect3 {
			t.Fatalf("expected %d found %d", expect3, v3)
		}
		if v4 != expect4 {
			t.Fatalf("expected %d found %d", expect4, v4)
		}
		if v5 != expect5 {
			t.Fatalf("expected %d found %d", expect5, v5)
		}
		if v6 != expect6 {
			t.Fatalf("expected %d found %d", expect6, v6)
		}
		if v7 != expect7 {
			t.Fatalf("expected %d found %d", expect7, v7)
		}
		if v8 != expect8 {
			t.Fatalf("expected %d found %d", expect8, v8)
		}
		if v9 != expect9 {
			t.Fatalf("expected %d found %d", expect9, v9)
		}
		if v10 != expect10 {
			t.Fatalf("expected %d found %d", expect10, v10)
		}
		if v11 != expect11 {
			t.Fatalf("expected %d found %d", expect11, v11)
		}		
	})

	t.Run("value", func(t *testing.T) {
		expect1 := 0
		expect2 := 1
		expect3 := 2
		expect4 := 3
		expect5 := 4
		expect6 := 5
		expect7 := 6
		expect8 := 7
		expect9 := 8
		expect10 := 9
		expect11 := 10
		tup := types.NewTuple11(expect1, expect2, expect3, expect4, expect5, expect6, expect7, expect8, expect9, expect10, expect11)
		
		if tup.Value1() != expect1 {
			t.Fatalf("expected %d to equal %d", tup.Value1(), expect1)
		}
		if tup.Value2() != expect2 {
			t.Fatalf("expected %d to equal %d", tup.Value2(), expect2)
		}
		if tup.Value3() != expect3 {
			t.Fatalf("expected %d to equal %d", tup.Value3(), expect3)
		}
		if tup.Value4() != expect4 {
			t.Fatalf("expected %d to equal %d", tup.Value4(), expect4)
		}
		if tup.Value5() != expect5 {
			t.Fatalf("expected %d to equal %d", tup.Value5(), expect5)
		}
		if tup.Value6() != expect6 {
			t.Fatalf("expected %d to equal %d", tup.Value6(), expect6)
		}
		if tup.Value7() != expect7 {
			t.Fatalf("expected %d to equal %d", tup.Value7(), expect7)
		}
		if tup.Value8() != expect8 {
			t.Fatalf("expected %d to equal %d", tup.Value8(), expect8)
		}
		if tup.Value9() != expect9 {
			t.Fatalf("expected %d to equal %d", tup.Value9(), expect9)
		}
		if tup.Value10() != expect10 {
			t.Fatalf("expected %d to equal %d", tup.Value10(), expect10)
		}
		if tup.Value11() != expect11 {
			t.Fatalf("expected %d to equal %d", tup.Value11(), expect11)
		}	
	})
}

func TestTuple12(t *testing.T) {
	t.Run("deconstruct", func(t *testing.T) {
		expect1 := 0
		expect2 := 1
		expect3 := 2
		expect4 := 3
		expect5 := 4
		expect6 := 5
		expect7 := 6
		expect8 := 7
		expect9 := 8
		expect10 := 9
		expect11 := 10
		expect12 := 11
		t12 := types.NewTuple12(expect1, expect2, expect3, expect4, expect5, expect6, expect7, expect8, expect9, expect10, expect11, expect12)
		v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12 := t12.Deconstruct()
		
		if v1 != expect1 {
			t.Fatalf("expected %d found %d", expect1, v1)
		}
		if v2 != expect2 {
			t.Fatalf("expected %d found %d", expect2, v2)
		}
		if v3 != expect3 {
			t.Fatalf("expected %d found %d", expect3, v3)
		}
		if v4 != expect4 {
			t.Fatalf("expected %d found %d", expect4, v4)
		}
		if v5 != expect5 {
			t.Fatalf("expected %d found %d", expect5, v5)
		}
		if v6 != expect6 {
			t.Fatalf("expected %d found %d", expect6, v6)
		}
		if v7 != expect7 {
			t.Fatalf("expected %d found %d", expect7, v7)
		}
		if v8 != expect8 {
			t.Fatalf("expected %d found %d", expect8, v8)
		}
		if v9 != expect9 {
			t.Fatalf("expected %d found %d", expect9, v9)
		}
		if v10 != expect10 {
			t.Fatalf("expected %d found %d", expect10, v10)
		}
		if v11 != expect11 {
			t.Fatalf("expected %d found %d", expect11, v11)
		}
		if v12 != expect12 {
			t.Fatalf("expected %d found %d", expect12, v12)
		}		
	})

	t.Run("value", func(t *testing.T) {
		expect1 := 0
		expect2 := 1
		expect3 := 2
		expect4 := 3
		expect5 := 4
		expect6 := 5
		expect7 := 6
		expect8 := 7
		expect9 := 8
		expect10 := 9
		expect11 := 10
		expect12 := 11
		tup := types.NewTuple12(expect1, expect2, expect3, expect4, expect5, expect6, expect7, expect8, expect9, expect10, expect11, expect12)
		
		if tup.Value1() != expect1 {
			t.Fatalf("expected %d to equal %d", tup.Value1(), expect1)
		}
		if tup.Value2() != expect2 {
			t.Fatalf("expected %d to equal %d", tup.Value2(), expect2)
		}
		if tup.Value3() != expect3 {
			t.Fatalf("expected %d to equal %d", tup.Value3(), expect3)
		}
		if tup.Value4() != expect4 {
			t.Fatalf("expected %d to equal %d", tup.Value4(), expect4)
		}
		if tup.Value5() != expect5 {
			t.Fatalf("expected %d to equal %d", tup.Value5(), expect5)
		}
		if tup.Value6() != expect6 {
			t.Fatalf("expected %d to equal %d", tup.Value6(), expect6)
		}
		if tup.Value7() != expect7 {
			t.Fatalf("expected %d to equal %d", tup.Value7(), expect7)
		}
		if tup.Value8() != expect8 {
			t.Fatalf("expected %d to equal %d", tup.Value8(), expect8)
		}
		if tup.Value9() != expect9 {
			t.Fatalf("expected %d to equal %d", tup.Value9(), expect9)
		}
		if tup.Value10() != expect10 {
			t.Fatalf("expected %d to equal %d", tup.Value10(), expect10)
		}
		if tup.Value11() != expect11 {
			t.Fatalf("expected %d to equal %d", tup.Value11(), expect11)
		}
		if tup.Value12() != expect12 {
			t.Fatalf("expected %d to equal %d", tup.Value12(), expect12)
		}	
	})
}

func TestTuple13(t *testing.T) {
	t.Run("deconstruct", func(t *testing.T) {
		expect1 := 0
		expect2 := 1
		expect3 := 2
		expect4 := 3
		expect5 := 4
		expect6 := 5
		expect7 := 6
		expect8 := 7
		expect9 := 8
		expect10 := 9
		expect11 := 10
		expect12 := 11
		expect13 := 12
		t13 := types.NewTuple13(expect1, expect2, expect3, expect4, expect5, expect6, expect7, expect8, expect9, expect10, expect11, expect12, expect13)
		v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13 := t13.Deconstruct()
		
		if v1 != expect1 {
			t.Fatalf("expected %d found %d", expect1, v1)
		}
		if v2 != expect2 {
			t.Fatalf("expected %d found %d", expect2, v2)
		}
		if v3 != expect3 {
			t.Fatalf("expected %d found %d", expect3, v3)
		}
		if v4 != expect4 {
			t.Fatalf("expected %d found %d", expect4, v4)
		}
		if v5 != expect5 {
			t.Fatalf("expected %d found %d", expect5, v5)
		}
		if v6 != expect6 {
			t.Fatalf("expected %d found %d", expect6, v6)
		}
		if v7 != expect7 {
			t.Fatalf("expected %d found %d", expect7, v7)
		}
		if v8 != expect8 {
			t.Fatalf("expected %d found %d", expect8, v8)
		}
		if v9 != expect9 {
			t.Fatalf("expected %d found %d", expect9, v9)
		}
		if v10 != expect10 {
			t.Fatalf("expected %d found %d", expect10, v10)
		}
		if v11 != expect11 {
			t.Fatalf("expected %d found %d", expect11, v11)
		}
		if v12 != expect12 {
			t.Fatalf("expected %d found %d", expect12, v12)
		}
		if v13 != expect13 {
			t.Fatalf("expected %d found %d", expect13, v13)
		}		
	})

	t.Run("value", func(t *testing.T) {
		expect1 := 0
		expect2 := 1
		expect3 := 2
		expect4 := 3
		expect5 := 4
		expect6 := 5
		expect7 := 6
		expect8 := 7
		expect9 := 8
		expect10 := 9
		expect11 := 10
		expect12 := 11
		expect13 := 12
		tup := types.NewTuple13(expect1, expect2, expect3, expect4, expect5, expect6, expect7, expect8, expect9, expect10, expect11, expect12, expect13)
		
		if tup.Value1() != expect1 {
			t.Fatalf("expected %d to equal %d", tup.Value1(), expect1)
		}
		if tup.Value2() != expect2 {
			t.Fatalf("expected %d to equal %d", tup.Value2(), expect2)
		}
		if tup.Value3() != expect3 {
			t.Fatalf("expected %d to equal %d", tup.Value3(), expect3)
		}
		if tup.Value4() != expect4 {
			t.Fatalf("expected %d to equal %d", tup.Value4(), expect4)
		}
		if tup.Value5() != expect5 {
			t.Fatalf("expected %d to equal %d", tup.Value5(), expect5)
		}
		if tup.Value6() != expect6 {
			t.Fatalf("expected %d to equal %d", tup.Value6(), expect6)
		}
		if tup.Value7() != expect7 {
			t.Fatalf("expected %d to equal %d", tup.Value7(), expect7)
		}
		if tup.Value8() != expect8 {
			t.Fatalf("expected %d to equal %d", tup.Value8(), expect8)
		}
		if tup.Value9() != expect9 {
			t.Fatalf("expected %d to equal %d", tup.Value9(), expect9)
		}
		if tup.Value10() != expect10 {
			t.Fatalf("expected %d to equal %d", tup.Value10(), expect10)
		}
		if tup.Value11() != expect11 {
			t.Fatalf("expected %d to equal %d", tup.Value11(), expect11)
		}
		if tup.Value12() != expect12 {
			t.Fatalf("expected %d to equal %d", tup.Value12(), expect12)
		}
		if tup.Value13() != expect13 {
			t.Fatalf("expected %d to equal %d", tup.Value13(), expect13)
		}	
	})
}

func TestTuple14(t *testing.T) {
	t.Run("deconstruct", func(t *testing.T) {
		expect1 := 0
		expect2 := 1
		expect3 := 2
		expect4 := 3
		expect5 := 4
		expect6 := 5
		expect7 := 6
		expect8 := 7
		expect9 := 8
		expect10 := 9
		expect11 := 10
		expect12 := 11
		expect13 := 12
		expect14 := 13
		t14 := types.NewTuple14(expect1, expect2, expect3, expect4, expect5, expect6, expect7, expect8, expect9, expect10, expect11, expect12, expect13, expect14)
		v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14 := t14.Deconstruct()
		
		if v1 != expect1 {
			t.Fatalf("expected %d found %d", expect1, v1)
		}
		if v2 != expect2 {
			t.Fatalf("expected %d found %d", expect2, v2)
		}
		if v3 != expect3 {
			t.Fatalf("expected %d found %d", expect3, v3)
		}
		if v4 != expect4 {
			t.Fatalf("expected %d found %d", expect4, v4)
		}
		if v5 != expect5 {
			t.Fatalf("expected %d found %d", expect5, v5)
		}
		if v6 != expect6 {
			t.Fatalf("expected %d found %d", expect6, v6)
		}
		if v7 != expect7 {
			t.Fatalf("expected %d found %d", expect7, v7)
		}
		if v8 != expect8 {
			t.Fatalf("expected %d found %d", expect8, v8)
		}
		if v9 != expect9 {
			t.Fatalf("expected %d found %d", expect9, v9)
		}
		if v10 != expect10 {
			t.Fatalf("expected %d found %d", expect10, v10)
		}
		if v11 != expect11 {
			t.Fatalf("expected %d found %d", expect11, v11)
		}
		if v12 != expect12 {
			t.Fatalf("expected %d found %d", expect12, v12)
		}
		if v13 != expect13 {
			t.Fatalf("expected %d found %d", expect13, v13)
		}
		if v14 != expect14 {
			t.Fatalf("expected %d found %d", expect14, v14)
		}		
	})

	t.Run("value", func(t *testing.T) {
		expect1 := 0
		expect2 := 1
		expect3 := 2
		expect4 := 3
		expect5 := 4
		expect6 := 5
		expect7 := 6
		expect8 := 7
		expect9 := 8
		expect10 := 9
		expect11 := 10
		expect12 := 11
		expect13 := 12
		expect14 := 13
		tup := types.NewTuple14(expect1, expect2, expect3, expect4, expect5, expect6, expect7, expect8, expect9, expect10, expect11, expect12, expect13, expect14)
		
		if tup.Value1() != expect1 {
			t.Fatalf("expected %d to equal %d", tup.Value1(), expect1)
		}
		if tup.Value2() != expect2 {
			t.Fatalf("expected %d to equal %d", tup.Value2(), expect2)
		}
		if tup.Value3() != expect3 {
			t.Fatalf("expected %d to equal %d", tup.Value3(), expect3)
		}
		if tup.Value4() != expect4 {
			t.Fatalf("expected %d to equal %d", tup.Value4(), expect4)
		}
		if tup.Value5() != expect5 {
			t.Fatalf("expected %d to equal %d", tup.Value5(), expect5)
		}
		if tup.Value6() != expect6 {
			t.Fatalf("expected %d to equal %d", tup.Value6(), expect6)
		}
		if tup.Value7() != expect7 {
			t.Fatalf("expected %d to equal %d", tup.Value7(), expect7)
		}
		if tup.Value8() != expect8 {
			t.Fatalf("expected %d to equal %d", tup.Value8(), expect8)
		}
		if tup.Value9() != expect9 {
			t.Fatalf("expected %d to equal %d", tup.Value9(), expect9)
		}
		if tup.Value10() != expect10 {
			t.Fatalf("expected %d to equal %d", tup.Value10(), expect10)
		}
		if tup.Value11() != expect11 {
			t.Fatalf("expected %d to equal %d", tup.Value11(), expect11)
		}
		if tup.Value12() != expect12 {
			t.Fatalf("expected %d to equal %d", tup.Value12(), expect12)
		}
		if tup.Value13() != expect13 {
			t.Fatalf("expected %d to equal %d", tup.Value13(), expect13)
		}
		if tup.Value14() != expect14 {
			t.Fatalf("expected %d to equal %d", tup.Value14(), expect14)
		}	
	})
}

func TestTuple15(t *testing.T) {
	t.Run("deconstruct", func(t *testing.T) {
		expect1 := 0
		expect2 := 1
		expect3 := 2
		expect4 := 3
		expect5 := 4
		expect6 := 5
		expect7 := 6
		expect8 := 7
		expect9 := 8
		expect10 := 9
		expect11 := 10
		expect12 := 11
		expect13 := 12
		expect14 := 13
		expect15 := 14
		t15 := types.NewTuple15(expect1, expect2, expect3, expect4, expect5, expect6, expect7, expect8, expect9, expect10, expect11, expect12, expect13, expect14, expect15)
		v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15 := t15.Deconstruct()
		
		if v1 != expect1 {
			t.Fatalf("expected %d found %d", expect1, v1)
		}
		if v2 != expect2 {
			t.Fatalf("expected %d found %d", expect2, v2)
		}
		if v3 != expect3 {
			t.Fatalf("expected %d found %d", expect3, v3)
		}
		if v4 != expect4 {
			t.Fatalf("expected %d found %d", expect4, v4)
		}
		if v5 != expect5 {
			t.Fatalf("expected %d found %d", expect5, v5)
		}
		if v6 != expect6 {
			t.Fatalf("expected %d found %d", expect6, v6)
		}
		if v7 != expect7 {
			t.Fatalf("expected %d found %d", expect7, v7)
		}
		if v8 != expect8 {
			t.Fatalf("expected %d found %d", expect8, v8)
		}
		if v9 != expect9 {
			t.Fatalf("expected %d found %d", expect9, v9)
		}
		if v10 != expect10 {
			t.Fatalf("expected %d found %d", expect10, v10)
		}
		if v11 != expect11 {
			t.Fatalf("expected %d found %d", expect11, v11)
		}
		if v12 != expect12 {
			t.Fatalf("expected %d found %d", expect12, v12)
		}
		if v13 != expect13 {
			t.Fatalf("expected %d found %d", expect13, v13)
		}
		if v14 != expect14 {
			t.Fatalf("expected %d found %d", expect14, v14)
		}
		if v15 != expect15 {
			t.Fatalf("expected %d found %d", expect15, v15)
		}		
	})

	t.Run("value", func(t *testing.T) {
		expect1 := 0
		expect2 := 1
		expect3 := 2
		expect4 := 3
		expect5 := 4
		expect6 := 5
		expect7 := 6
		expect8 := 7
		expect9 := 8
		expect10 := 9
		expect11 := 10
		expect12 := 11
		expect13 := 12
		expect14 := 13
		expect15 := 14
		tup := types.NewTuple15(expect1, expect2, expect3, expect4, expect5, expect6, expect7, expect8, expect9, expect10, expect11, expect12, expect13, expect14, expect15)
		
		if tup.Value1() != expect1 {
			t.Fatalf("expected %d to equal %d", tup.Value1(), expect1)
		}
		if tup.Value2() != expect2 {
			t.Fatalf("expected %d to equal %d", tup.Value2(), expect2)
		}
		if tup.Value3() != expect3 {
			t.Fatalf("expected %d to equal %d", tup.Value3(), expect3)
		}
		if tup.Value4() != expect4 {
			t.Fatalf("expected %d to equal %d", tup.Value4(), expect4)
		}
		if tup.Value5() != expect5 {
			t.Fatalf("expected %d to equal %d", tup.Value5(), expect5)
		}
		if tup.Value6() != expect6 {
			t.Fatalf("expected %d to equal %d", tup.Value6(), expect6)
		}
		if tup.Value7() != expect7 {
			t.Fatalf("expected %d to equal %d", tup.Value7(), expect7)
		}
		if tup.Value8() != expect8 {
			t.Fatalf("expected %d to equal %d", tup.Value8(), expect8)
		}
		if tup.Value9() != expect9 {
			t.Fatalf("expected %d to equal %d", tup.Value9(), expect9)
		}
		if tup.Value10() != expect10 {
			t.Fatalf("expected %d to equal %d", tup.Value10(), expect10)
		}
		if tup.Value11() != expect11 {
			t.Fatalf("expected %d to equal %d", tup.Value11(), expect11)
		}
		if tup.Value12() != expect12 {
			t.Fatalf("expected %d to equal %d", tup.Value12(), expect12)
		}
		if tup.Value13() != expect13 {
			t.Fatalf("expected %d to equal %d", tup.Value13(), expect13)
		}
		if tup.Value14() != expect14 {
			t.Fatalf("expected %d to equal %d", tup.Value14(), expect14)
		}
		if tup.Value15() != expect15 {
			t.Fatalf("expected %d to equal %d", tup.Value15(), expect15)
		}	
	})
}

func TestTuple16(t *testing.T) {
	t.Run("deconstruct", func(t *testing.T) {
		expect1 := 0
		expect2 := 1
		expect3 := 2
		expect4 := 3
		expect5 := 4
		expect6 := 5
		expect7 := 6
		expect8 := 7
		expect9 := 8
		expect10 := 9
		expect11 := 10
		expect12 := 11
		expect13 := 12
		expect14 := 13
		expect15 := 14
		expect16 := 15
		t16 := types.NewTuple16(expect1, expect2, expect3, expect4, expect5, expect6, expect7, expect8, expect9, expect10, expect11, expect12, expect13, expect14, expect15, expect16)
		v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15, v16 := t16.Deconstruct()
		
		if v1 != expect1 {
			t.Fatalf("expected %d found %d", expect1, v1)
		}
		if v2 != expect2 {
			t.Fatalf("expected %d found %d", expect2, v2)
		}
		if v3 != expect3 {
			t.Fatalf("expected %d found %d", expect3, v3)
		}
		if v4 != expect4 {
			t.Fatalf("expected %d found %d", expect4, v4)
		}
		if v5 != expect5 {
			t.Fatalf("expected %d found %d", expect5, v5)
		}
		if v6 != expect6 {
			t.Fatalf("expected %d found %d", expect6, v6)
		}
		if v7 != expect7 {
			t.Fatalf("expected %d found %d", expect7, v7)
		}
		if v8 != expect8 {
			t.Fatalf("expected %d found %d", expect8, v8)
		}
		if v9 != expect9 {
			t.Fatalf("expected %d found %d", expect9, v9)
		}
		if v10 != expect10 {
			t.Fatalf("expected %d found %d", expect10, v10)
		}
		if v11 != expect11 {
			t.Fatalf("expected %d found %d", expect11, v11)
		}
		if v12 != expect12 {
			t.Fatalf("expected %d found %d", expect12, v12)
		}
		if v13 != expect13 {
			t.Fatalf("expected %d found %d", expect13, v13)
		}
		if v14 != expect14 {
			t.Fatalf("expected %d found %d", expect14, v14)
		}
		if v15 != expect15 {
			t.Fatalf("expected %d found %d", expect15, v15)
		}
		if v16 != expect16 {
			t.Fatalf("expected %d found %d", expect16, v16)
		}		
	})

	t.Run("value", func(t *testing.T) {
		expect1 := 0
		expect2 := 1
		expect3 := 2
		expect4 := 3
		expect5 := 4
		expect6 := 5
		expect7 := 6
		expect8 := 7
		expect9 := 8
		expect10 := 9
		expect11 := 10
		expect12 := 11
		expect13 := 12
		expect14 := 13
		expect15 := 14
		expect16 := 15
		tup := types.NewTuple16(expect1, expect2, expect3, expect4, expect5, expect6, expect7, expect8, expect9, expect10, expect11, expect12, expect13, expect14, expect15, expect16)
		
		if tup.Value1() != expect1 {
			t.Fatalf("expected %d to equal %d", tup.Value1(), expect1)
		}
		if tup.Value2() != expect2 {
			t.Fatalf("expected %d to equal %d", tup.Value2(), expect2)
		}
		if tup.Value3() != expect3 {
			t.Fatalf("expected %d to equal %d", tup.Value3(), expect3)
		}
		if tup.Value4() != expect4 {
			t.Fatalf("expected %d to equal %d", tup.Value4(), expect4)
		}
		if tup.Value5() != expect5 {
			t.Fatalf("expected %d to equal %d", tup.Value5(), expect5)
		}
		if tup.Value6() != expect6 {
			t.Fatalf("expected %d to equal %d", tup.Value6(), expect6)
		}
		if tup.Value7() != expect7 {
			t.Fatalf("expected %d to equal %d", tup.Value7(), expect7)
		}
		if tup.Value8() != expect8 {
			t.Fatalf("expected %d to equal %d", tup.Value8(), expect8)
		}
		if tup.Value9() != expect9 {
			t.Fatalf("expected %d to equal %d", tup.Value9(), expect9)
		}
		if tup.Value10() != expect10 {
			t.Fatalf("expected %d to equal %d", tup.Value10(), expect10)
		}
		if tup.Value11() != expect11 {
			t.Fatalf("expected %d to equal %d", tup.Value11(), expect11)
		}
		if tup.Value12() != expect12 {
			t.Fatalf("expected %d to equal %d", tup.Value12(), expect12)
		}
		if tup.Value13() != expect13 {
			t.Fatalf("expected %d to equal %d", tup.Value13(), expect13)
		}
		if tup.Value14() != expect14 {
			t.Fatalf("expected %d to equal %d", tup.Value14(), expect14)
		}
		if tup.Value15() != expect15 {
			t.Fatalf("expected %d to equal %d", tup.Value15(), expect15)
		}
		if tup.Value16() != expect16 {
			t.Fatalf("expected %d to equal %d", tup.Value16(), expect16)
		}	
	})
}
