package arraylist

import "github.com/hidetatz/collection"

// ArrayList is a variable-sized list.
type ArrayList[T any] struct {
	values []T
}

// New returns an ArrayList based on the specified type.
func New[T any]() *ArrayList[T] {
	return &ArrayList[T]{values: []T{}}
}

// PushBack appends the given value to the bottom.
func (a *ArrayList[T]) PushBack(value T) {
	a.values = append(a.values, value)
}

// PushFront inserts the given value to the head.
func (a *ArrayList[T]) PushFront(value T) {
	a.values = append([]T{value}, a.values...)
}

// Len returns the size of the list.
func (a *ArrayList[T]) Len() int {
	return len(a.values)
}

// Init inits (clears) the list. Initiated list is still usable.
func (a *ArrayList[T]) Init() {
	a.values = []T{}
}

// GetAt returns the value at the given index.
// When the given index is less than 0 or greater than the list size,
// ErrInvalidIndex will be responded.
func (a *ArrayList[T]) GetAt(i int) (T, error) {
	var zero T
	if i < 0 || len(a.values) <= i {
		return zero, collection.ErrInvalidIndex
	}

	return a.values[i], nil
}

// InsertAt inserts the given element at the given index.
// When the given index is less than 0 or greater than the list size,
// ErrInvalidIndex will be responded.
func (a *ArrayList[T]) InsertAt(i int, e T) error {
	if i < 0 || len(a.values) < i {
		return collection.ErrInvalidIndex
	}

	if len(a.values) == i {
		a.values = append(a.values, e)
		return nil
	}

	a.values = append(a.values[:i+1], a.values[i:]...)
	a.values[i] = e
	return nil
}

// RemoveAt removes a value at the given index in the list.
// ErrInvalidIndex will be responded if the index < 0 or length <= index.
func (a *ArrayList[T]) RemoveAt(index int) error {
	if index < 0 || len(a.values) <= index {
		return collection.ErrInvalidIndex
	}

	a.values = append(a.values[:index], a.values[index+1:]...)
	return nil
}

// SetAt replaces the value at the given index in the list with the given value.
// ErrInvalidIndex will be responded if the index < 0 or length <= index.
func (a *ArrayList[T]) SetAt(index int, v T) error {
	if index < 0 || len(a.values) <= index {
		return collection.ErrInvalidIndex
	}

	a.values[index] = v
	return nil
}
