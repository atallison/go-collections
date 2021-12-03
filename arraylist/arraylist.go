package arraylist

type ArrayList[T any] struct {
    values []T
}

func New[T any]() *ArrayList[T] {
	return &ArrayList[T]{}
}

func(a *ArrayList[T]) Append(e T) {
    a.values = append(a.values, e)
}

func(a *ArrayList[T]) AppendAll(es []T) {
    a.values = append(a.values, es...)
}

func(a *ArrayList[T]) Insert(i int, e T) {
    a.values = append(a.values[:i+1], a.values[i:]...)
    a.values[i] = e
}
