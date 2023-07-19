package types

import "fmt"

// Cast pereforms a type assertion on the given source to the target type.
// If the type assertion succeeds, Ok[Target] is returned.
// If the type assertion fails, Error[TTarget] is returned.
func Cast[TSource, TTarget any](source TSource) Result[TTarget] {
	var zero TTarget
	return Castf[TSource, TTarget](source, "unable to cast %T to %T", source, zero)

}

// Castf pereforms a type assertion on the given source to the target type.
// If the type assertion succeeds, Ok[Target] is returned.
// If the type assertion fails, Error[TTarget] with the giiven message is returned.
func Castf[TSource, TTarget any](source TSource, format string, args ...any) Result[TTarget] {
	target, ok := any(source).(TTarget)
	if !ok {
		return NewError[TTarget](fmt.Errorf(format, args...))
	}
	return NewOk(target)
}
