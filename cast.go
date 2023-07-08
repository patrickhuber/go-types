package types

import "fmt"

func Cast[TSource, TTarget any](source TSource) Result[TTarget] {
	var zero TTarget
	return Castf[TSource, TTarget](source, "unable to cast %T to %T", source, zero)

}

func Castf[TSource, TTarget any](source TSource, format string, args ...any) Result[TTarget] {
	target, ok := any(source).(TTarget)
	if !ok {
		return NewError[TTarget](fmt.Errorf(format, args...))
	}
	return NewOk(target)
}
