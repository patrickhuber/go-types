package option

import "github.com/patrickhuber/go-types"

// Map returns none if op is none
// applies the function to the value if some and returns some of the result
func Map[T, U any](op types.Option[T], transform func(t T) U) types.Option[U] {
	some, ok := op.(types.Some[T])
	if !ok {
		return None[U]()
	}
	return Some(transform(some.Value))
}

// MapOr returns the default value if none.
// or applies the function to the value if some
func MapOr[T, U any](op types.Option[T], transform func(t T) U, def U) U {
	m := Map[T, U](op, transform)
	if some, ok := m.(types.Some[U]); ok {
		return some.Value
	}
	return def
}

// MapOrElse evaluates the default function if none
// or applies the function to the value if some
func MapOrElse[T, U any](op types.Option[T], transform func(t T) U, def func() U) U {
	m := Map[T, U](op, transform)
	if some, ok := m.(types.Some[U]); ok {
		return some.Value
	}
	return def()
}
