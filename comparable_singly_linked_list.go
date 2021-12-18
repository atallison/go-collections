package collection

type ComparableSinglyLinkedList[T comparable] struct {
	*SinglyLinkedList[T]
}

// NewComparableSinglyLinkedList returns an SinglyLinkedList based on the specified type which contains only comparable values.
func NewComparableSinglyLinkedList[T comparable]() *ComparableSinglyLinkedList[T] {
	return &ComparableSinglyLinkedList[T]{SinglyLinkedList: NewSinglyLinkedList[T]()}
}

// Remove removes the same value with given v in the list.
// It uses == operator to make sure if the values are the same.
func (a *ComparableSinglyLinkedList[T]) Remove(v T) {
	if a.SinglyLinkedList.length == 0 {
		return
	}

	if a.SinglyLinkedList.length == 1 {
		if a.SinglyLinkedList.head.v == v {
			a.SinglyLinkedList.head = nil
			a.SinglyLinkedList.tail = nil
			a.SinglyLinkedList.length = 0
		}
		return
	}

	// First, make sure the head value != v
	for {
		if a.SinglyLinkedList.head == nil {
			// all values are removed
			return
		}
		if a.SinglyLinkedList.head.v != v {
			break
		} else {
			a.SinglyLinkedList.head = a.SinglyLinkedList.head.next
			a.SinglyLinkedList.length--
		}
	}

	// At this point, some values are remaining but we are sure the head is not v
	prev := a.SinglyLinkedList.head
	curr := a.SinglyLinkedList.head.next
	for {
		if curr == nil {
			return
		}

		if curr.v == v {
			prev.next = curr.next
			a.SinglyLinkedList.length--
		}
		next := curr.next
		prev = curr
		curr = next
	}
}

// Contains returns if the given value is contained in the list.
func (a *ComparableSinglyLinkedList[T]) Contains(v T) bool {
	if a.SinglyLinkedList.length == 0 {
		return false
	}

	curr := a.SinglyLinkedList.head
	for {
		if curr == nil {
			break
		}
		if curr.v == v {
			return true
		}
		curr = curr.next
	}
	return false
}
