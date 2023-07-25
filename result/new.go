package result

import (
	"github.com/patrickhuber/go-types"
)

func New[T any](ok T, err error) types.Result[T] {
	return types.NewResult(ok, err)
}

func New2[T1, T2 any](t1 T1, t2 T2, err error) types.Result[types.Tuple2[T1, T2]] {
	return types.NewResult(
		types.NewTuple2(t1, t2), 
		err,
	)
}

func New3[T1, T2, T3 any](t1 T1, t2 T2, t3 T3, err error) types.Result[types.Tuple3[T1, T2, T3]] {
	return types.NewResult(
		types.NewTuple3(t1, t2, t3), 
		err,
	)
}

func New4[T1, T2, T3, T4 any](t1 T1, t2 T2, t3 T3, t4 T4, err error) types.Result[types.Tuple4[T1, T2, T3, T4]] {
	return types.NewResult(
		types.NewTuple4(t1, t2, t3, t4), 
		err,
	)
}

func New5[T1, T2, T3, T4, T5 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, err error) types.Result[types.Tuple5[T1, T2, T3, T4, T5]] {
	return types.NewResult(
		types.NewTuple5(t1, t2, t3, t4, t5), 
		err,
	)
}

func New6[T1, T2, T3, T4, T5, T6 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, err error) types.Result[types.Tuple6[T1, T2, T3, T4, T5, T6]] {
	return types.NewResult(
		types.NewTuple6(t1, t2, t3, t4, t5, t6), 
		err,
	)
}

func New7[T1, T2, T3, T4, T5, T6, T7 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7, err error) types.Result[types.Tuple7[T1, T2, T3, T4, T5, T6, T7]] {
	return types.NewResult(
		types.NewTuple7(t1, t2, t3, t4, t5, t6, t7), 
		err,
	)
}

func New8[T1, T2, T3, T4, T5, T6, T7, T8 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7, t8 T8, err error) types.Result[types.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]] {
	return types.NewResult(
		types.NewTuple8(t1, t2, t3, t4, t5, t6, t7, t8), 
		err,
	)
}

func New9[T1, T2, T3, T4, T5, T6, T7, T8, T9 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7, t8 T8, t9 T9, err error) types.Result[types.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]] {
	return types.NewResult(
		types.NewTuple9(t1, t2, t3, t4, t5, t6, t7, t8, t9), 
		err,
	)
}

func New10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7, t8 T8, t9 T9, t10 T10, err error) types.Result[types.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]] {
	return types.NewResult(
		types.NewTuple10(t1, t2, t3, t4, t5, t6, t7, t8, t9, t10), 
		err,
	)
}

func New11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7, t8 T8, t9 T9, t10 T10, t11 T11, err error) types.Result[types.Tuple11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11]] {
	return types.NewResult(
		types.NewTuple11(t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, t11), 
		err,
	)
}

func New12[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7, t8 T8, t9 T9, t10 T10, t11 T11, t12 T12, err error) types.Result[types.Tuple12[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12]] {
	return types.NewResult(
		types.NewTuple12(t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, t11, t12), 
		err,
	)
}

func New13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7, t8 T8, t9 T9, t10 T10, t11 T11, t12 T12, t13 T13, err error) types.Result[types.Tuple13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13]] {
	return types.NewResult(
		types.NewTuple13(t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, t11, t12, t13), 
		err,
	)
}

func New14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7, t8 T8, t9 T9, t10 T10, t11 T11, t12 T12, t13 T13, t14 T14, err error) types.Result[types.Tuple14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14]] {
	return types.NewResult(
		types.NewTuple14(t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, t11, t12, t13, t14), 
		err,
	)
}

func New15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7, t8 T8, t9 T9, t10 T10, t11 T11, t12 T12, t13 T13, t14 T14, t15 T15, err error) types.Result[types.Tuple15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15]] {
	return types.NewResult(
		types.NewTuple15(t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, t11, t12, t13, t14, t15), 
		err,
	)
}

func New16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7, t8 T8, t9 T9, t10 T10, t11 T11, t12 T12, t13 T13, t14 T14, t15 T15, t16 T16, err error) types.Result[types.Tuple16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16]] {
	return types.NewResult(
		types.NewTuple16(t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, t11, t12, t13, t14, t15, t16), 
		err,
	)
}
