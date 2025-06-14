package option

import "github.com/patrickhuber/go-types"

func Get[TKey comparable, TValue any](m map[TKey]TValue, key TKey) types.Option[TValue] {
	v, ok := m[key]
	return New(v, ok)
}
