package arraylist

import (
	"errors"
)

var (
	InvalidIndexErr = errors.New("invalid index")
)

type ArrayList[T any] struct {
	values []T
}

func New[T any]() *ArrayList[T] {
	return &ArrayList[T]{}
}

func(a *ArrayList[T]) Add(e T) {
	a.values = append(a.values, e)
}

func(a *ArrayList[T]) AddAll(es []T) {
	a.values = append(a.values, es...)
}

func(a *ArrayList[T]) AddAt(i int, e T) error {
	if i < 0 || len(a.values) < i {
		return InvalidIndexErr
	}

	if len(a.values) == i {
		a.values = append(a.values, e)
		return nil
	}

	a.values = append(a.values[:i+1], a.values[i:]...)
	a.values[i] = e
	return nil
}

func(a *ArrayList[T]) AddAllAt(i int, es []T) error {
	if i < 0 || len(a.values) < i {
		return InvalidIndexErr
	}

	a.values = append(a.values[:i], append(es, a.values[i:]...)...)
	return nil
}
