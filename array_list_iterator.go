package collection

// ArrayListIterator is an implementation of Iterator for ArrayList.
type ArrayListIterator[T any] struct {
	arrayList *ArrayList[T]
	cursor    int
}

// Next returns if the next value exists in the iterator.
func (ai *ArrayListIterator[T]) Next() bool {
	return ai.cursor < len(ai.arrayList.values)
}

// Value returns the next value in the iterator.
func (ai *ArrayListIterator[T]) Value() T {
	v := ai.arrayList.values[ai.cursor]
	ai.cursor++
	return v
}

// Set sets the value which is returned by the last call of Value() from the original ArrayList safely.
// Note that this must be called at most once per a Value() call. Unless, it might lead to list/iterator inconsistent or invalid state and
// that case is not tested.
func (ai *ArrayListIterator[T]) Set(value T) {
	// ignore error because cursor must be a valid value as index.
	_ = ai.arrayList.Set(ai.cursor-1, value)
}

// Remove removes the value which is returned by the last call of Value() from the original ArrayList.
// Note that this must be called at most once per a Value() call. Unless, it might lead to list/iterator inconsistent or invalid state and
// that case is not tested.
func (ai *ArrayListIterator[T]) Remove() {
	ai.arrayList.RemoveAt(ai.cursor - 1)
	ai.cursor--
}
