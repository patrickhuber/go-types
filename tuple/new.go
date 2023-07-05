package tuple

import "github.com/patrickhuber/go-types"

func New1[T1 any](t1 T1) types.Tuple1[T1] {
	return types.NewTuple1(t1)
}

func New2[T1, T2 any](t1 T1, t2 T2) types.Tuple2[T1, T2] {
	return types.NewTuple2[T1, T2](t1, t2)
}

func New3[T1, T2, T3 any](t1 T1, t2 T2, t3 T3) types.Tuple3[T1, T2, T3] {
	return types.NewTuple3[T1, T2, T3](t1, t2, t3)
}
func New4[T1, T2, T3, T4 any](t1 T1, t2 T2, t3 T3, t4 T4) types.Tuple4[T1, T2, T3, T4] {
	return types.NewTuple4[T1, T2, T3, T4](t1, t2, t3, t4)
}
