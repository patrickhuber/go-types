package types

type Tuple2[T1, T2 any] interface {
	tuple2(t1 T1, t2 T2)
	Deconstruct() (T1, T2)
	Value1() T1
	Value2() T2
}

type tuple2[T1, T2 any] struct {
	t1 T1
	t2 T2
}

//lint:ignore U1000 method is used to implement Tuple2[T1,T2]
func (*tuple2[T1, T2]) tuple2(t1 T1, t2 T2) {}

func (t *tuple2[T1, T2]) Deconstruct() (T1, T2) {
	return t.t1, t.t2
}

func NewTuple2[T1, T2 any](t1 T1, t2 T2) Tuple2[T1, T2] {
	return &tuple2[T1, T2]{
		t1: t1,
		t2: t2,
	}
}

func (t *tuple2[T1, T2]) Value1() T1 {
	return t.t1
}

func (t *tuple2[T1, T2]) Value2() T2 {
	return t.t2
}

type Tuple3[T1, T2, T3 any] interface {
	tuple3(t1 T1, t2 T2, t3 T3)
	Deconstruct() (T1, T2, T3)
	Value1() T1
	Value2() T2
	Value3() T3
}

type tuple3[T1, T2, T3 any] struct {
	t1 T1
	t2 T2
	t3 T3
}

//lint:ignore U1000 method is used to implement Tuple3[T1,T2,T3]
func (*tuple3[T1, T2, T3]) tuple3(t1 T1, t2 T2, t3 T3) {}

func (t *tuple3[T1, T2, T3]) Deconstruct() (T1, T2, T3) {
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

func NewTuple3[T1, T2, T3 any](t1 T1, t2 T2, t3 T3) Tuple3[T1, T2, T3] {
	return &tuple3[T1, T2, T3]{
		t1: t1,
		t2: t2,
		t3: t3,
	}
}

type Tuple4[T1, T2, T3, T4 any] interface {
	tuple4(t1 T1, t2 T2, t3 T3, t4 T4)
	Deconstruct() (T1, T2, T3, T4)
	Value1() T1
	Value2() T2
	Value3() T3
	Value4() T4
}

type tuple4[T1, T2, T3, T4 any] struct {
	t1 T1
	t2 T2
	t3 T3
	t4 T4
}

//lint:ignore U1000 method is used to implement Tuple4[T1,T2,T3,T4]
func (*tuple4[T1, T2, T3, T4]) tuple4(t1 T1, t2 T2, t3 T3, t4 T4) {}

func (t *tuple4[T1, T2, T3, T4]) Deconstruct() (T1, T2, T3, T4) {
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

func NewTuple4[T1, T2, T3, T4 any](t1 T1, t2 T2, t3 T3, t4 T4) Tuple4[T1, T2, T3, T4] {
	return &tuple4[T1, T2, T3, T4]{
		t1: t1,
		t2: t2,
		t3: t3,
		t4: t4,
	}
}
