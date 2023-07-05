package result

import (
	"github.com/patrickhuber/go-types"
	tuples "github.com/patrickhuber/go-types/tuple"
)

func New[T any](ok T, err error) types.Result[T] {
	return types.NewResult(ok, err)
}

func New2[T1, T2 any](t1 T1, t2 T2, err error) types.Result[types.Tuple2[T1, T2]] {
	return types.NewResult[types.Tuple2[T1, T2]](
		tuples.New2(t1, t2),
		err,
	)
}

func New3[T1, T2, T3 any](t1 T1, t2 T2, t3 T3, err error) types.Result[types.Tuple3[T1, T2, T3]] {
	return types.NewResult[types.Tuple3[T1, T2, T3]](
		tuples.New3(t1, t2, t3),
		err)
}

func New4[T1, T2, T3, T4 any](t1 T1, t2 T2, t3 T3, t4 T4, err error) types.Result[types.Tuple4[T1, T2, T3, T4]] {
	return types.NewResult[types.Tuple4[T1, T2, T3, T4]](
		tuples.New4(t1, t2, t3, t4),
		err)
}
