package collection

// SinglyLinkedListIterator is the iteratable made of SinglyLinkedList.
type SinglyLinkedListIterator[T any] struct {
	curr  *singlyLinkedNode[T]
	index int
}

func (i *SinglyLinkedListIterator[T]) Next() bool {
	return i.curr != nil
}

func (i *SinglyLinkedListIterator[T]) Value() T {
	ret := i.curr.v
	i.curr = i.curr.next
	return ret
}
