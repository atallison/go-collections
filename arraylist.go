package collection

// ArrayList is a variable-sized list.
// This data structure is not concurrent safe.
// The caller must be responsible to synchronize before accessing from multiple goroutines.
type ArrayList[T any] struct {
	values []T
}

// NewArrayList returns an ArrayList based on the specified type.
func NewArrayList[T any]() *ArrayList[T] {
	return &ArrayList[T]{values: []T{}}
}

// Add appends a given value to the bottom of the list.
func (a *ArrayList[T]) Add(e T) {
	a.values = append(a.values, e)
}

// AddAll appends given values to the bottom of the list in the given order.
func (a *ArrayList[T]) AddAll(es []T) {
	a.values = append(a.values, es...)
}

// AddAt inserts the given element at the given index.
// When the given is less than 0 or greater than the list size,
// ErrInvalidIndex will be responded.
func (a *ArrayList[T]) AddAt(i int, e T) error {
	if i < 0 || len(a.values) < i {
		return ErrInvalidIndex
	}

	if len(a.values) == i {
		a.values = append(a.values, e)
		return nil
	}

	a.values = append(a.values[:i+1], a.values[i:]...)
	a.values[i] = e
	return nil
}

// AddAllAt inserts the given values at the given index.
// Let's say current list looks [a, b, c] then
// calls AddAllAt(1, [d, e]), it results [a, d, e, b, c].
// When the given index is less than 0 or greater than the list size,
// ErrInvalidIndex will be responded.
func (a *ArrayList[T]) AddAllAt(i int, es []T) error {
	if i < 0 || len(a.values) < i {
		return ErrInvalidIndex
	}

	a.values = append(a.values[:i], append(es, a.values[i:]...)...)
	return nil
}

// Clear removes all the data in the list. The list is still usable after clear.
func (a *ArrayList[T]) Clear() {
	a.values = []T{}
}

// Get returns a value which locates at the given index of the list.
// ErrInvalidIndex will be responded if the index < 0 or length <= index.
// greater equal than
func (a *ArrayList[T]) Get(index int) (ret T, err error) {
	if index < 0 || len(a.values) <= index {
		return ret, ErrInvalidIndex
	}

	return a.values[index], nil
}
