package collection

// LinkedListIterator is the iteratable made of LinkedList.
type LinkedListIterator[T any] struct {
	curr  *linkedNode[T]
	index int
}

func (i *LinkedListIterator[T]) Next() bool {
	return i.curr != nil
}

func (i *LinkedListIterator[T]) Value() T {
	ret := i.curr.v
	i.curr = i.curr.next
	return ret
}
