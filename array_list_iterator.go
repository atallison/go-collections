package collection

// ArrayListIterator is an implementation of Iterator.
// Because ArrayList is not concurrent safe,
// ArrayListIterator is not coucurrent safe as well.
type ArrayListIterator[T any] struct {
	values []T
	index  int
}

func (ai *ArrayListIterator[T]) Next() bool {
	return ai.index < len(ai.values)
}

func (ai *ArrayListIterator[T]) Value() T {
	ret := ai.values[ai.index]
	ai.index++
	return ret
}
