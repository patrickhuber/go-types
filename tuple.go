package types

type Tuple2[T1, T2 any] interface{
	tuple2(t1 T1, t2 T2)
	Deconstruct() (T1, T2)
	Value1() T1
	Value2() T2
}

type tuple2[T1, T2 any] struct{
	t1 T1
	t2 T2
}

//lint:ignore U1000 method is used to implement Tuple2[T1, T2]
func (*tuple2[T1, T2]) tuple2(t1 T1, t2 T2) {}

func (t *tuple2[T1, T2]) Deconstruct() (T1, T2){
	return t.t1, t.t2
}

func (t *tuple2[T1, T2]) Value1() T1 {
	return t.t1
}

func (t *tuple2[T1, T2]) Value2() T2 {
	return t.t2
}

func NewTuple2[T1, T2 any](t1 T1, t2 T2) Tuple2[T1, T2]{
	return &tuple2[T1, T2]{
		t1: t1, 
		t2: t2, 
	}
}

type Tuple3[T1, T2, T3 any] interface{
	tuple3(t1 T1, t2 T2, t3 T3)
	Deconstruct() (T1, T2, T3)
	Value1() T1
	Value2() T2
	Value3() T3
}

type tuple3[T1, T2, T3 any] struct{
	t1 T1
	t2 T2
	t3 T3
}

//lint:ignore U1000 method is used to implement Tuple3[T1, T2, T3]
func (*tuple3[T1, T2, T3]) tuple3(t1 T1, t2 T2, t3 T3) {}

func (t *tuple3[T1, T2, T3]) Deconstruct() (T1, T2, T3){
	return t.t1, t.t2, t.t3
}

func (t *tuple3[T1, T2, T3]) Value1() T1 {
	return t.t1
}

func (t *tuple3[T1, T2, T3]) Value2() T2 {
	return t.t2
}

func (t *tuple3[T1, T2, T3]) Value3() T3 {
	return t.t3
}

func NewTuple3[T1, T2, T3 any](t1 T1, t2 T2, t3 T3) Tuple3[T1, T2, T3]{
	return &tuple3[T1, T2, T3]{
		t1: t1, 
		t2: t2, 
		t3: t3, 
	}
}

type Tuple4[T1, T2, T3, T4 any] interface{
	tuple4(t1 T1, t2 T2, t3 T3, t4 T4)
	Deconstruct() (T1, T2, T3, T4)
	Value1() T1
	Value2() T2
	Value3() T3
	Value4() T4
}

type tuple4[T1, T2, T3, T4 any] struct{
	t1 T1
	t2 T2
	t3 T3
	t4 T4
}

//lint:ignore U1000 method is used to implement Tuple4[T1, T2, T3, T4]
func (*tuple4[T1, T2, T3, T4]) tuple4(t1 T1, t2 T2, t3 T3, t4 T4) {}

func (t *tuple4[T1, T2, T3, T4]) Deconstruct() (T1, T2, T3, T4){
	return t.t1, t.t2, t.t3, t.t4
}

func (t *tuple4[T1, T2, T3, T4]) Value1() T1 {
	return t.t1
}

func (t *tuple4[T1, T2, T3, T4]) Value2() T2 {
	return t.t2
}

func (t *tuple4[T1, T2, T3, T4]) Value3() T3 {
	return t.t3
}

func (t *tuple4[T1, T2, T3, T4]) Value4() T4 {
	return t.t4
}

func NewTuple4[T1, T2, T3, T4 any](t1 T1, t2 T2, t3 T3, t4 T4) Tuple4[T1, T2, T3, T4]{
	return &tuple4[T1, T2, T3, T4]{
		t1: t1, 
		t2: t2, 
		t3: t3, 
		t4: t4, 
	}
}

type Tuple5[T1, T2, T3, T4, T5 any] interface{
	tuple5(t1 T1, t2 T2, t3 T3, t4 T4, t5 T5)
	Deconstruct() (T1, T2, T3, T4, T5)
	Value1() T1
	Value2() T2
	Value3() T3
	Value4() T4
	Value5() T5
}

type tuple5[T1, T2, T3, T4, T5 any] struct{
	t1 T1
	t2 T2
	t3 T3
	t4 T4
	t5 T5
}

//lint:ignore U1000 method is used to implement Tuple5[T1, T2, T3, T4, T5]
func (*tuple5[T1, T2, T3, T4, T5]) tuple5(t1 T1, t2 T2, t3 T3, t4 T4, t5 T5) {}

func (t *tuple5[T1, T2, T3, T4, T5]) Deconstruct() (T1, T2, T3, T4, T5){
	return t.t1, t.t2, t.t3, t.t4, t.t5
}

func (t *tuple5[T1, T2, T3, T4, T5]) Value1() T1 {
	return t.t1
}

func (t *tuple5[T1, T2, T3, T4, T5]) Value2() T2 {
	return t.t2
}

func (t *tuple5[T1, T2, T3, T4, T5]) Value3() T3 {
	return t.t3
}

func (t *tuple5[T1, T2, T3, T4, T5]) Value4() T4 {
	return t.t4
}

func (t *tuple5[T1, T2, T3, T4, T5]) Value5() T5 {
	return t.t5
}

func NewTuple5[T1, T2, T3, T4, T5 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5) Tuple5[T1, T2, T3, T4, T5]{
	return &tuple5[T1, T2, T3, T4, T5]{
		t1: t1, 
		t2: t2, 
		t3: t3, 
		t4: t4, 
		t5: t5, 
	}
}

type Tuple6[T1, T2, T3, T4, T5, T6 any] interface{
	tuple6(t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6)
	Deconstruct() (T1, T2, T3, T4, T5, T6)
	Value1() T1
	Value2() T2
	Value3() T3
	Value4() T4
	Value5() T5
	Value6() T6
}

type tuple6[T1, T2, T3, T4, T5, T6 any] struct{
	t1 T1
	t2 T2
	t3 T3
	t4 T4
	t5 T5
	t6 T6
}

//lint:ignore U1000 method is used to implement Tuple6[T1, T2, T3, T4, T5, T6]
func (*tuple6[T1, T2, T3, T4, T5, T6]) tuple6(t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6) {}

func (t *tuple6[T1, T2, T3, T4, T5, T6]) Deconstruct() (T1, T2, T3, T4, T5, T6){
	return t.t1, t.t2, t.t3, t.t4, t.t5, t.t6
}

func (t *tuple6[T1, T2, T3, T4, T5, T6]) Value1() T1 {
	return t.t1
}

func (t *tuple6[T1, T2, T3, T4, T5, T6]) Value2() T2 {
	return t.t2
}

func (t *tuple6[T1, T2, T3, T4, T5, T6]) Value3() T3 {
	return t.t3
}

func (t *tuple6[T1, T2, T3, T4, T5, T6]) Value4() T4 {
	return t.t4
}

func (t *tuple6[T1, T2, T3, T4, T5, T6]) Value5() T5 {
	return t.t5
}

func (t *tuple6[T1, T2, T3, T4, T5, T6]) Value6() T6 {
	return t.t6
}

func NewTuple6[T1, T2, T3, T4, T5, T6 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6) Tuple6[T1, T2, T3, T4, T5, T6]{
	return &tuple6[T1, T2, T3, T4, T5, T6]{
		t1: t1, 
		t2: t2, 
		t3: t3, 
		t4: t4, 
		t5: t5, 
		t6: t6, 
	}
}

type Tuple7[T1, T2, T3, T4, T5, T6, T7 any] interface{
	tuple7(t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7)
	Deconstruct() (T1, T2, T3, T4, T5, T6, T7)
	Value1() T1
	Value2() T2
	Value3() T3
	Value4() T4
	Value5() T5
	Value6() T6
	Value7() T7
}

type tuple7[T1, T2, T3, T4, T5, T6, T7 any] struct{
	t1 T1
	t2 T2
	t3 T3
	t4 T4
	t5 T5
	t6 T6
	t7 T7
}

//lint:ignore U1000 method is used to implement Tuple7[T1, T2, T3, T4, T5, T6, T7]
func (*tuple7[T1, T2, T3, T4, T5, T6, T7]) tuple7(t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7) {}

func (t *tuple7[T1, T2, T3, T4, T5, T6, T7]) Deconstruct() (T1, T2, T3, T4, T5, T6, T7){
	return t.t1, t.t2, t.t3, t.t4, t.t5, t.t6, t.t7
}

func (t *tuple7[T1, T2, T3, T4, T5, T6, T7]) Value1() T1 {
	return t.t1
}

func (t *tuple7[T1, T2, T3, T4, T5, T6, T7]) Value2() T2 {
	return t.t2
}

func (t *tuple7[T1, T2, T3, T4, T5, T6, T7]) Value3() T3 {
	return t.t3
}

func (t *tuple7[T1, T2, T3, T4, T5, T6, T7]) Value4() T4 {
	return t.t4
}

func (t *tuple7[T1, T2, T3, T4, T5, T6, T7]) Value5() T5 {
	return t.t5
}

func (t *tuple7[T1, T2, T3, T4, T5, T6, T7]) Value6() T6 {
	return t.t6
}

func (t *tuple7[T1, T2, T3, T4, T5, T6, T7]) Value7() T7 {
	return t.t7
}

func NewTuple7[T1, T2, T3, T4, T5, T6, T7 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7) Tuple7[T1, T2, T3, T4, T5, T6, T7]{
	return &tuple7[T1, T2, T3, T4, T5, T6, T7]{
		t1: t1, 
		t2: t2, 
		t3: t3, 
		t4: t4, 
		t5: t5, 
		t6: t6, 
		t7: t7, 
	}
}

type Tuple8[T1, T2, T3, T4, T5, T6, T7, T8 any] interface{
	tuple8(t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7, t8 T8)
	Deconstruct() (T1, T2, T3, T4, T5, T6, T7, T8)
	Value1() T1
	Value2() T2
	Value3() T3
	Value4() T4
	Value5() T5
	Value6() T6
	Value7() T7
	Value8() T8
}

type tuple8[T1, T2, T3, T4, T5, T6, T7, T8 any] struct{
	t1 T1
	t2 T2
	t3 T3
	t4 T4
	t5 T5
	t6 T6
	t7 T7
	t8 T8
}

//lint:ignore U1000 method is used to implement Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]
func (*tuple8[T1, T2, T3, T4, T5, T6, T7, T8]) tuple8(t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7, t8 T8) {}

func (t *tuple8[T1, T2, T3, T4, T5, T6, T7, T8]) Deconstruct() (T1, T2, T3, T4, T5, T6, T7, T8){
	return t.t1, t.t2, t.t3, t.t4, t.t5, t.t6, t.t7, t.t8
}

func (t *tuple8[T1, T2, T3, T4, T5, T6, T7, T8]) Value1() T1 {
	return t.t1
}

func (t *tuple8[T1, T2, T3, T4, T5, T6, T7, T8]) Value2() T2 {
	return t.t2
}

func (t *tuple8[T1, T2, T3, T4, T5, T6, T7, T8]) Value3() T3 {
	return t.t3
}

func (t *tuple8[T1, T2, T3, T4, T5, T6, T7, T8]) Value4() T4 {
	return t.t4
}

func (t *tuple8[T1, T2, T3, T4, T5, T6, T7, T8]) Value5() T5 {
	return t.t5
}

func (t *tuple8[T1, T2, T3, T4, T5, T6, T7, T8]) Value6() T6 {
	return t.t6
}

func (t *tuple8[T1, T2, T3, T4, T5, T6, T7, T8]) Value7() T7 {
	return t.t7
}

func (t *tuple8[T1, T2, T3, T4, T5, T6, T7, T8]) Value8() T8 {
	return t.t8
}

func NewTuple8[T1, T2, T3, T4, T5, T6, T7, T8 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7, t8 T8) Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]{
	return &tuple8[T1, T2, T3, T4, T5, T6, T7, T8]{
		t1: t1, 
		t2: t2, 
		t3: t3, 
		t4: t4, 
		t5: t5, 
		t6: t6, 
		t7: t7, 
		t8: t8, 
	}
}

type Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9 any] interface{
	tuple9(t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7, t8 T8, t9 T9)
	Deconstruct() (T1, T2, T3, T4, T5, T6, T7, T8, T9)
	Value1() T1
	Value2() T2
	Value3() T3
	Value4() T4
	Value5() T5
	Value6() T6
	Value7() T7
	Value8() T8
	Value9() T9
}

type tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9 any] struct{
	t1 T1
	t2 T2
	t3 T3
	t4 T4
	t5 T5
	t6 T6
	t7 T7
	t8 T8
	t9 T9
}

//lint:ignore U1000 method is used to implement Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]
func (*tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) tuple9(t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7, t8 T8, t9 T9) {}

func (t *tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) Deconstruct() (T1, T2, T3, T4, T5, T6, T7, T8, T9){
	return t.t1, t.t2, t.t3, t.t4, t.t5, t.t6, t.t7, t.t8, t.t9
}

func (t *tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) Value1() T1 {
	return t.t1
}

func (t *tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) Value2() T2 {
	return t.t2
}

func (t *tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) Value3() T3 {
	return t.t3
}

func (t *tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) Value4() T4 {
	return t.t4
}

func (t *tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) Value5() T5 {
	return t.t5
}

func (t *tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) Value6() T6 {
	return t.t6
}

func (t *tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) Value7() T7 {
	return t.t7
}

func (t *tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) Value8() T8 {
	return t.t8
}

func (t *tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) Value9() T9 {
	return t.t9
}

func NewTuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7, t8 T8, t9 T9) Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]{
	return &tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]{
		t1: t1, 
		t2: t2, 
		t3: t3, 
		t4: t4, 
		t5: t5, 
		t6: t6, 
		t7: t7, 
		t8: t8, 
		t9: t9, 
	}
}

type Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10 any] interface{
	tuple10(t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7, t8 T8, t9 T9, t10 T10)
	Deconstruct() (T1, T2, T3, T4, T5, T6, T7, T8, T9, T10)
	Value1() T1
	Value2() T2
	Value3() T3
	Value4() T4
	Value5() T5
	Value6() T6
	Value7() T7
	Value8() T8
	Value9() T9
	Value10() T10
}

type tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10 any] struct{
	t1 T1
	t2 T2
	t3 T3
	t4 T4
	t5 T5
	t6 T6
	t7 T7
	t8 T8
	t9 T9
	t10 T10
}

//lint:ignore U1000 method is used to implement Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]
func (*tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) tuple10(t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7, t8 T8, t9 T9, t10 T10) {}

func (t *tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Deconstruct() (T1, T2, T3, T4, T5, T6, T7, T8, T9, T10){
	return t.t1, t.t2, t.t3, t.t4, t.t5, t.t6, t.t7, t.t8, t.t9, t.t10
}

func (t *tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Value1() T1 {
	return t.t1
}

func (t *tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Value2() T2 {
	return t.t2
}

func (t *tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Value3() T3 {
	return t.t3
}

func (t *tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Value4() T4 {
	return t.t4
}

func (t *tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Value5() T5 {
	return t.t5
}

func (t *tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Value6() T6 {
	return t.t6
}

func (t *tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Value7() T7 {
	return t.t7
}

func (t *tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Value8() T8 {
	return t.t8
}

func (t *tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Value9() T9 {
	return t.t9
}

func (t *tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Value10() T10 {
	return t.t10
}

func NewTuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7, t8 T8, t9 T9, t10 T10) Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]{
	return &tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]{
		t1: t1, 
		t2: t2, 
		t3: t3, 
		t4: t4, 
		t5: t5, 
		t6: t6, 
		t7: t7, 
		t8: t8, 
		t9: t9, 
		t10: t10, 
	}
}

type Tuple11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11 any] interface{
	tuple11(t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7, t8 T8, t9 T9, t10 T10, t11 T11)
	Deconstruct() (T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11)
	Value1() T1
	Value2() T2
	Value3() T3
	Value4() T4
	Value5() T5
	Value6() T6
	Value7() T7
	Value8() T8
	Value9() T9
	Value10() T10
	Value11() T11
}

type tuple11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11 any] struct{
	t1 T1
	t2 T2
	t3 T3
	t4 T4
	t5 T5
	t6 T6
	t7 T7
	t8 T8
	t9 T9
	t10 T10
	t11 T11
}

//lint:ignore U1000 method is used to implement Tuple11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11]
func (*tuple11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11]) tuple11(t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7, t8 T8, t9 T9, t10 T10, t11 T11) {}

func (t *tuple11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11]) Deconstruct() (T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11){
	return t.t1, t.t2, t.t3, t.t4, t.t5, t.t6, t.t7, t.t8, t.t9, t.t10, t.t11
}

func (t *tuple11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11]) Value1() T1 {
	return t.t1
}

func (t *tuple11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11]) Value2() T2 {
	return t.t2
}

func (t *tuple11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11]) Value3() T3 {
	return t.t3
}

func (t *tuple11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11]) Value4() T4 {
	return t.t4
}

func (t *tuple11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11]) Value5() T5 {
	return t.t5
}

func (t *tuple11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11]) Value6() T6 {
	return t.t6
}

func (t *tuple11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11]) Value7() T7 {
	return t.t7
}

func (t *tuple11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11]) Value8() T8 {
	return t.t8
}

func (t *tuple11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11]) Value9() T9 {
	return t.t9
}

func (t *tuple11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11]) Value10() T10 {
	return t.t10
}

func (t *tuple11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11]) Value11() T11 {
	return t.t11
}

func NewTuple11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7, t8 T8, t9 T9, t10 T10, t11 T11) Tuple11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11]{
	return &tuple11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11]{
		t1: t1, 
		t2: t2, 
		t3: t3, 
		t4: t4, 
		t5: t5, 
		t6: t6, 
		t7: t7, 
		t8: t8, 
		t9: t9, 
		t10: t10, 
		t11: t11, 
	}
}

type Tuple12[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12 any] interface{
	tuple12(t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7, t8 T8, t9 T9, t10 T10, t11 T11, t12 T12)
	Deconstruct() (T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12)
	Value1() T1
	Value2() T2
	Value3() T3
	Value4() T4
	Value5() T5
	Value6() T6
	Value7() T7
	Value8() T8
	Value9() T9
	Value10() T10
	Value11() T11
	Value12() T12
}

type tuple12[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12 any] struct{
	t1 T1
	t2 T2
	t3 T3
	t4 T4
	t5 T5
	t6 T6
	t7 T7
	t8 T8
	t9 T9
	t10 T10
	t11 T11
	t12 T12
}

//lint:ignore U1000 method is used to implement Tuple12[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12]
func (*tuple12[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12]) tuple12(t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7, t8 T8, t9 T9, t10 T10, t11 T11, t12 T12) {}

func (t *tuple12[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12]) Deconstruct() (T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12){
	return t.t1, t.t2, t.t3, t.t4, t.t5, t.t6, t.t7, t.t8, t.t9, t.t10, t.t11, t.t12
}

func (t *tuple12[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12]) Value1() T1 {
	return t.t1
}

func (t *tuple12[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12]) Value2() T2 {
	return t.t2
}

func (t *tuple12[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12]) Value3() T3 {
	return t.t3
}

func (t *tuple12[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12]) Value4() T4 {
	return t.t4
}

func (t *tuple12[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12]) Value5() T5 {
	return t.t5
}

func (t *tuple12[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12]) Value6() T6 {
	return t.t6
}

func (t *tuple12[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12]) Value7() T7 {
	return t.t7
}

func (t *tuple12[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12]) Value8() T8 {
	return t.t8
}

func (t *tuple12[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12]) Value9() T9 {
	return t.t9
}

func (t *tuple12[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12]) Value10() T10 {
	return t.t10
}

func (t *tuple12[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12]) Value11() T11 {
	return t.t11
}

func (t *tuple12[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12]) Value12() T12 {
	return t.t12
}

func NewTuple12[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7, t8 T8, t9 T9, t10 T10, t11 T11, t12 T12) Tuple12[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12]{
	return &tuple12[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12]{
		t1: t1, 
		t2: t2, 
		t3: t3, 
		t4: t4, 
		t5: t5, 
		t6: t6, 
		t7: t7, 
		t8: t8, 
		t9: t9, 
		t10: t10, 
		t11: t11, 
		t12: t12, 
	}
}

type Tuple13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13 any] interface{
	tuple13(t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7, t8 T8, t9 T9, t10 T10, t11 T11, t12 T12, t13 T13)
	Deconstruct() (T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13)
	Value1() T1
	Value2() T2
	Value3() T3
	Value4() T4
	Value5() T5
	Value6() T6
	Value7() T7
	Value8() T8
	Value9() T9
	Value10() T10
	Value11() T11
	Value12() T12
	Value13() T13
}

type tuple13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13 any] struct{
	t1 T1
	t2 T2
	t3 T3
	t4 T4
	t5 T5
	t6 T6
	t7 T7
	t8 T8
	t9 T9
	t10 T10
	t11 T11
	t12 T12
	t13 T13
}

//lint:ignore U1000 method is used to implement Tuple13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13]
func (*tuple13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13]) tuple13(t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7, t8 T8, t9 T9, t10 T10, t11 T11, t12 T12, t13 T13) {}

func (t *tuple13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13]) Deconstruct() (T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13){
	return t.t1, t.t2, t.t3, t.t4, t.t5, t.t6, t.t7, t.t8, t.t9, t.t10, t.t11, t.t12, t.t13
}

func (t *tuple13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13]) Value1() T1 {
	return t.t1
}

func (t *tuple13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13]) Value2() T2 {
	return t.t2
}

func (t *tuple13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13]) Value3() T3 {
	return t.t3
}

func (t *tuple13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13]) Value4() T4 {
	return t.t4
}

func (t *tuple13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13]) Value5() T5 {
	return t.t5
}

func (t *tuple13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13]) Value6() T6 {
	return t.t6
}

func (t *tuple13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13]) Value7() T7 {
	return t.t7
}

func (t *tuple13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13]) Value8() T8 {
	return t.t8
}

func (t *tuple13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13]) Value9() T9 {
	return t.t9
}

func (t *tuple13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13]) Value10() T10 {
	return t.t10
}

func (t *tuple13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13]) Value11() T11 {
	return t.t11
}

func (t *tuple13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13]) Value12() T12 {
	return t.t12
}

func (t *tuple13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13]) Value13() T13 {
	return t.t13
}

func NewTuple13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7, t8 T8, t9 T9, t10 T10, t11 T11, t12 T12, t13 T13) Tuple13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13]{
	return &tuple13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13]{
		t1: t1, 
		t2: t2, 
		t3: t3, 
		t4: t4, 
		t5: t5, 
		t6: t6, 
		t7: t7, 
		t8: t8, 
		t9: t9, 
		t10: t10, 
		t11: t11, 
		t12: t12, 
		t13: t13, 
	}
}

type Tuple14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14 any] interface{
	tuple14(t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7, t8 T8, t9 T9, t10 T10, t11 T11, t12 T12, t13 T13, t14 T14)
	Deconstruct() (T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14)
	Value1() T1
	Value2() T2
	Value3() T3
	Value4() T4
	Value5() T5
	Value6() T6
	Value7() T7
	Value8() T8
	Value9() T9
	Value10() T10
	Value11() T11
	Value12() T12
	Value13() T13
	Value14() T14
}

type tuple14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14 any] struct{
	t1 T1
	t2 T2
	t3 T3
	t4 T4
	t5 T5
	t6 T6
	t7 T7
	t8 T8
	t9 T9
	t10 T10
	t11 T11
	t12 T12
	t13 T13
	t14 T14
}

//lint:ignore U1000 method is used to implement Tuple14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14]
func (*tuple14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14]) tuple14(t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7, t8 T8, t9 T9, t10 T10, t11 T11, t12 T12, t13 T13, t14 T14) {}

func (t *tuple14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14]) Deconstruct() (T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14){
	return t.t1, t.t2, t.t3, t.t4, t.t5, t.t6, t.t7, t.t8, t.t9, t.t10, t.t11, t.t12, t.t13, t.t14
}

func (t *tuple14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14]) Value1() T1 {
	return t.t1
}

func (t *tuple14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14]) Value2() T2 {
	return t.t2
}

func (t *tuple14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14]) Value3() T3 {
	return t.t3
}

func (t *tuple14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14]) Value4() T4 {
	return t.t4
}

func (t *tuple14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14]) Value5() T5 {
	return t.t5
}

func (t *tuple14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14]) Value6() T6 {
	return t.t6
}

func (t *tuple14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14]) Value7() T7 {
	return t.t7
}

func (t *tuple14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14]) Value8() T8 {
	return t.t8
}

func (t *tuple14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14]) Value9() T9 {
	return t.t9
}

func (t *tuple14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14]) Value10() T10 {
	return t.t10
}

func (t *tuple14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14]) Value11() T11 {
	return t.t11
}

func (t *tuple14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14]) Value12() T12 {
	return t.t12
}

func (t *tuple14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14]) Value13() T13 {
	return t.t13
}

func (t *tuple14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14]) Value14() T14 {
	return t.t14
}

func NewTuple14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7, t8 T8, t9 T9, t10 T10, t11 T11, t12 T12, t13 T13, t14 T14) Tuple14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14]{
	return &tuple14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14]{
		t1: t1, 
		t2: t2, 
		t3: t3, 
		t4: t4, 
		t5: t5, 
		t6: t6, 
		t7: t7, 
		t8: t8, 
		t9: t9, 
		t10: t10, 
		t11: t11, 
		t12: t12, 
		t13: t13, 
		t14: t14, 
	}
}

type Tuple15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15 any] interface{
	tuple15(t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7, t8 T8, t9 T9, t10 T10, t11 T11, t12 T12, t13 T13, t14 T14, t15 T15)
	Deconstruct() (T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15)
	Value1() T1
	Value2() T2
	Value3() T3
	Value4() T4
	Value5() T5
	Value6() T6
	Value7() T7
	Value8() T8
	Value9() T9
	Value10() T10
	Value11() T11
	Value12() T12
	Value13() T13
	Value14() T14
	Value15() T15
}

type tuple15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15 any] struct{
	t1 T1
	t2 T2
	t3 T3
	t4 T4
	t5 T5
	t6 T6
	t7 T7
	t8 T8
	t9 T9
	t10 T10
	t11 T11
	t12 T12
	t13 T13
	t14 T14
	t15 T15
}

//lint:ignore U1000 method is used to implement Tuple15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15]
func (*tuple15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15]) tuple15(t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7, t8 T8, t9 T9, t10 T10, t11 T11, t12 T12, t13 T13, t14 T14, t15 T15) {}

func (t *tuple15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15]) Deconstruct() (T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15){
	return t.t1, t.t2, t.t3, t.t4, t.t5, t.t6, t.t7, t.t8, t.t9, t.t10, t.t11, t.t12, t.t13, t.t14, t.t15
}

func (t *tuple15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15]) Value1() T1 {
	return t.t1
}

func (t *tuple15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15]) Value2() T2 {
	return t.t2
}

func (t *tuple15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15]) Value3() T3 {
	return t.t3
}

func (t *tuple15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15]) Value4() T4 {
	return t.t4
}

func (t *tuple15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15]) Value5() T5 {
	return t.t5
}

func (t *tuple15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15]) Value6() T6 {
	return t.t6
}

func (t *tuple15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15]) Value7() T7 {
	return t.t7
}

func (t *tuple15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15]) Value8() T8 {
	return t.t8
}

func (t *tuple15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15]) Value9() T9 {
	return t.t9
}

func (t *tuple15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15]) Value10() T10 {
	return t.t10
}

func (t *tuple15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15]) Value11() T11 {
	return t.t11
}

func (t *tuple15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15]) Value12() T12 {
	return t.t12
}

func (t *tuple15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15]) Value13() T13 {
	return t.t13
}

func (t *tuple15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15]) Value14() T14 {
	return t.t14
}

func (t *tuple15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15]) Value15() T15 {
	return t.t15
}

func NewTuple15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7, t8 T8, t9 T9, t10 T10, t11 T11, t12 T12, t13 T13, t14 T14, t15 T15) Tuple15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15]{
	return &tuple15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15]{
		t1: t1, 
		t2: t2, 
		t3: t3, 
		t4: t4, 
		t5: t5, 
		t6: t6, 
		t7: t7, 
		t8: t8, 
		t9: t9, 
		t10: t10, 
		t11: t11, 
		t12: t12, 
		t13: t13, 
		t14: t14, 
		t15: t15, 
	}
}

type Tuple16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16 any] interface{
	tuple16(t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7, t8 T8, t9 T9, t10 T10, t11 T11, t12 T12, t13 T13, t14 T14, t15 T15, t16 T16)
	Deconstruct() (T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16)
	Value1() T1
	Value2() T2
	Value3() T3
	Value4() T4
	Value5() T5
	Value6() T6
	Value7() T7
	Value8() T8
	Value9() T9
	Value10() T10
	Value11() T11
	Value12() T12
	Value13() T13
	Value14() T14
	Value15() T15
	Value16() T16
}

type tuple16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16 any] struct{
	t1 T1
	t2 T2
	t3 T3
	t4 T4
	t5 T5
	t6 T6
	t7 T7
	t8 T8
	t9 T9
	t10 T10
	t11 T11
	t12 T12
	t13 T13
	t14 T14
	t15 T15
	t16 T16
}

//lint:ignore U1000 method is used to implement Tuple16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16]
func (*tuple16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16]) tuple16(t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7, t8 T8, t9 T9, t10 T10, t11 T11, t12 T12, t13 T13, t14 T14, t15 T15, t16 T16) {}

func (t *tuple16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16]) Deconstruct() (T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16){
	return t.t1, t.t2, t.t3, t.t4, t.t5, t.t6, t.t7, t.t8, t.t9, t.t10, t.t11, t.t12, t.t13, t.t14, t.t15, t.t16
}

func (t *tuple16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16]) Value1() T1 {
	return t.t1
}

func (t *tuple16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16]) Value2() T2 {
	return t.t2
}

func (t *tuple16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16]) Value3() T3 {
	return t.t3
}

func (t *tuple16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16]) Value4() T4 {
	return t.t4
}

func (t *tuple16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16]) Value5() T5 {
	return t.t5
}

func (t *tuple16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16]) Value6() T6 {
	return t.t6
}

func (t *tuple16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16]) Value7() T7 {
	return t.t7
}

func (t *tuple16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16]) Value8() T8 {
	return t.t8
}

func (t *tuple16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16]) Value9() T9 {
	return t.t9
}

func (t *tuple16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16]) Value10() T10 {
	return t.t10
}

func (t *tuple16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16]) Value11() T11 {
	return t.t11
}

func (t *tuple16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16]) Value12() T12 {
	return t.t12
}

func (t *tuple16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16]) Value13() T13 {
	return t.t13
}

func (t *tuple16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16]) Value14() T14 {
	return t.t14
}

func (t *tuple16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16]) Value15() T15 {
	return t.t15
}

func (t *tuple16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16]) Value16() T16 {
	return t.t16
}

func NewTuple16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7, t8 T8, t9 T9, t10 T10, t11 T11, t12 T12, t13 T13, t14 T14, t15 T15, t16 T16) Tuple16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16]{
	return &tuple16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16]{
		t1: t1, 
		t2: t2, 
		t3: t3, 
		t4: t4, 
		t5: t5, 
		t6: t6, 
		t7: t7, 
		t8: t8, 
		t9: t9, 
		t10: t10, 
		t11: t11, 
		t12: t12, 
		t13: t13, 
		t14: t14, 
		t15: t15, 
		t16: t16, 
	}
}
