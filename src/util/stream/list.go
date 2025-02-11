package stream

type IListStream[T any] <-chan T

func NewListStream[T any](items []T) IListStream[T] {
	ch := make(chan T)
	go func() {
		defer close(ch)
		for _, item := range items {
			ch <- item
		}
	}()
	return ch
}

func (s IListStream[T]) Map(fn func(T) any) IListStream[any] {
	out := make(chan any)
	go func() {
		defer close(out)
		for item := range s {
			out <- fn(item)
		}
	}()
	return out
}

func (s IListStream[T]) Filter(fn func(T) bool) IListStream[T] {
	out := make(chan T)
	go func() {
		defer close(out)
		for item := range s {
			if fn(item) {
				out <- item
			}
		}
	}()
	return out
}

func (s IListStream[T]) ToList() []T {
	out := make([]T, 0)
	for item := range s {
		out = append(out, item)
	}
	return out
}
