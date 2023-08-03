package option

import (
	"github.com/patrickhuber/go-types"
	"github.com/patrickhuber/go-types/result"
)

// Cast attempts to cast a source option type of TSource to a result option type of TTarget
// if the cast fails, Error is returned
// if the cast succeeds, Ok is returned wrapping the Option[TTarget]
func Cast[TSource, TTarget any](source types.Option[TSource]) types.Result[types.Option[TTarget]] {
	switch t := source.(type) {

	case types.None[TSource]:
		return result.Ok(None[TTarget]())

	case types.Some[TSource]:
		value := any(t.Value)
		target, ok := value.(TTarget)
		if !ok {
			return result.Errorf[types.Option[TTarget]]("options.cast : unable to match type target type %T", value)
		}
		return result.Ok(Some(target))

	default:
		return result.Errorf[types.Option[TTarget]](
			"options.cast : unable to match type source type %T", source)
	}
}
