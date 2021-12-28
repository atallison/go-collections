package collection

// SinglyLinkedListIterator is the iteratable made of SinglyLinkedList.
type SinglyLinkedListIterator[T any] struct {
	list  *SinglyLinkedList[T]
	curr  *singlyLinkedNode[T]
	index int
}

// Next returns if the iterator has the next value.
func (i *SinglyLinkedListIterator[T]) Next() bool {
	return i.curr != nil
}

// Value returns the next value. This must be called when Next() is true.
func (i *SinglyLinkedListIterator[T]) Value() T {
	ret := i.curr.v
	i.curr = i.curr.next
	i.index++
	return ret
}

// Set sets the value which is returned by the last call of Value() from the original SinglyLinkedList.
// Note that this must be called at most once per a Value() call. Unless, it might lead to list/iterator inconsistent or invalid state and
// that case is not tested.
func (i *SinglyLinkedListIterator[T]) Set(value T) {
	_ = i.list.Set(i.index-1, value) // ignore error because index must be valid
}

// Remove removes the value which is returned by the last call of Value() from the original SinglyLinkedList.
// Note that this must be called at most once per a Value() call. Unless, it might lead to list/iterator inconsistent or invalid state and
// that case is not tested.
func (i *SinglyLinkedListIterator[T]) Remove() {
	_ = i.list.RemoveAt(i.index - 1)
	i.index--
}
