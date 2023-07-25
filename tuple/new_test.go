package tuple_test

import (
	"github.com/patrickhuber/go-types/tuple"

	"testing"
)

func TestNew2(t *testing.T) {
	t.Run("value", func(t *testing.T) {
		expect1 := 0
		expect2 := 1
		tup := tuple.New2(expect1, expect2)
		
		if tup.Value1() != expect1 {
			t.Fatalf("expected %d to equal %d", tup.Value1(), expect1)
		}
		if tup.Value2() != expect2 {
			t.Fatalf("expected %d to equal %d", tup.Value2(), expect2)
		}	
	})
}

func TestNew3(t *testing.T) {
	t.Run("value", func(t *testing.T) {
		expect1 := 0
		expect2 := 1
		expect3 := 2
		tup := tuple.New3(expect1, expect2, expect3)
		
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

func TestNew4(t *testing.T) {
	t.Run("value", func(t *testing.T) {
		expect1 := 0
		expect2 := 1
		expect3 := 2
		expect4 := 3
		tup := tuple.New4(expect1, expect2, expect3, expect4)
		
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

func TestNew5(t *testing.T) {
	t.Run("value", func(t *testing.T) {
		expect1 := 0
		expect2 := 1
		expect3 := 2
		expect4 := 3
		expect5 := 4
		tup := tuple.New5(expect1, expect2, expect3, expect4, expect5)
		
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

func TestNew6(t *testing.T) {
	t.Run("value", func(t *testing.T) {
		expect1 := 0
		expect2 := 1
		expect3 := 2
		expect4 := 3
		expect5 := 4
		expect6 := 5
		tup := tuple.New6(expect1, expect2, expect3, expect4, expect5, expect6)
		
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

func TestNew7(t *testing.T) {
	t.Run("value", func(t *testing.T) {
		expect1 := 0
		expect2 := 1
		expect3 := 2
		expect4 := 3
		expect5 := 4
		expect6 := 5
		expect7 := 6
		tup := tuple.New7(expect1, expect2, expect3, expect4, expect5, expect6, expect7)
		
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

func TestNew8(t *testing.T) {
	t.Run("value", func(t *testing.T) {
		expect1 := 0
		expect2 := 1
		expect3 := 2
		expect4 := 3
		expect5 := 4
		expect6 := 5
		expect7 := 6
		expect8 := 7
		tup := tuple.New8(expect1, expect2, expect3, expect4, expect5, expect6, expect7, expect8)
		
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

func TestNew9(t *testing.T) {
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
		tup := tuple.New9(expect1, expect2, expect3, expect4, expect5, expect6, expect7, expect8, expect9)
		
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

func TestNew10(t *testing.T) {
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
		tup := tuple.New10(expect1, expect2, expect3, expect4, expect5, expect6, expect7, expect8, expect9, expect10)
		
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

func TestNew11(t *testing.T) {
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
		tup := tuple.New11(expect1, expect2, expect3, expect4, expect5, expect6, expect7, expect8, expect9, expect10, expect11)
		
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

func TestNew12(t *testing.T) {
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
		tup := tuple.New12(expect1, expect2, expect3, expect4, expect5, expect6, expect7, expect8, expect9, expect10, expect11, expect12)
		
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

func TestNew13(t *testing.T) {
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
		tup := tuple.New13(expect1, expect2, expect3, expect4, expect5, expect6, expect7, expect8, expect9, expect10, expect11, expect12, expect13)
		
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

func TestNew14(t *testing.T) {
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
		tup := tuple.New14(expect1, expect2, expect3, expect4, expect5, expect6, expect7, expect8, expect9, expect10, expect11, expect12, expect13, expect14)
		
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

func TestNew15(t *testing.T) {
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
		tup := tuple.New15(expect1, expect2, expect3, expect4, expect5, expect6, expect7, expect8, expect9, expect10, expect11, expect12, expect13, expect14, expect15)
		
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

func TestNew16(t *testing.T) {
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
		tup := tuple.New16(expect1, expect2, expect3, expect4, expect5, expect6, expect7, expect8, expect9, expect10, expect11, expect12, expect13, expect14, expect15, expect16)
		
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

