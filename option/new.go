package option

import "github.com/patrickhuber/go-types"

func New[T any](t T, ok bool) types.Option[T] {
	if ok {
		return Some(t)
	}
	return None[T]()
}
