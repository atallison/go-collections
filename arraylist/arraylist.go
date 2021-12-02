package arraylist

type ArrayList[T any] struct {
    values []T
}

func New[T any]() *ArrayList[T] {
	return &ArrayList[T]{}
}

func(a *ArrayList[T]) Add(e T) {
    a.values = append(a.values, e)
}

