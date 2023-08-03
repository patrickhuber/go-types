package result

import (
	"github.com/patrickhuber/go-types"
)

func Cast[TSource, TTarget any](source types.Result[TSource]) types.Result[TTarget] {
	switch t := source.(type) {
	case types.Ok[TSource]:
		value := any(t.Value)
		target, ok := value.(TTarget)
		if !ok {
			var t TTarget
			return Errorf[TTarget]("results.cast : unable to match target type %T", t)
		}
		return Ok(target)
	case types.Error[TSource]:
		return Error[TTarget](t.Value)
	default:
		return Errorf[TTarget]("results.cast : unable to match source type %T", source)
	}
}
