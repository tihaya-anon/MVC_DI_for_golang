package stream

type pair[K string | int, V any] struct {
	key K
	val V
}

type IMapStream[K string | int, V any] <-chan pair[K, V]

func NewMapStream[K string | int, V any](m map[K]V) IMapStream[K, V] {
	ch := make(chan pair[K, V])
	go func() {
		defer close(ch)
		for key, val := range m {
			ch <- pair[K, V]{key, val}
		}
	}()
	return ch
}

func (s IMapStream[K, V]) Map(fn func(K, V) (K, any)) IMapStream[K, any] {
	out := make(chan pair[K, any])
	go func() {
		defer close(out)
		for item := range s {
			key, val := fn(item.key, item.val)
			out <- pair[K, any]{key, val}
		}
	}()
	return out
}

func (s IMapStream[K, V]) Filter(fn func(K, V) bool) IMapStream[K, V] {
	out := make(chan pair[K, V])
	go func() {
		defer close(out)
		for item := range s {
			if fn(item.key, item.val) {
				out <- item
			}
		}
	}()
	return out
}

func (s IMapStream[K, V]) ToMap() map[K]V {
	out := make(map[K]V, 0)
	for item := range s {
		out[item.key] = item.val
	}
	return out
}
