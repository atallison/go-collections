package collection

type ComparableLinkedList[T comparable] struct {
	*LinkedList[T]
}

// NewComparableLinkedList returns an LinkedList based on the specified type which contains only comparable values.
func NewComparableLinkedList[T comparable]() *ComparableLinkedList[T] {
	return &ComparableLinkedList[T]{LinkedList: NewLinkedList[T]()}
}

// Remove removes the same value with given v in the list.
// It uses == operator to make sure if the values are the same.
func (a *ComparableLinkedList[T]) Remove(v T) {
	if a.LinkedList.length == 0 {
		return
	}

	if a.LinkedList.length == 1 {
		if a.LinkedList.head.v == v {
			a.LinkedList.head = nil
			a.LinkedList.tail = nil
			a.LinkedList.length = 0
		}
		return
	}

	// First, make sure the head value != v
	for {
		if a.LinkedList.head == nil {
			// all values are removed
			return
		}
		if a.LinkedList.head.v != v {
			break
		} else {
			a.LinkedList.head = a.LinkedList.head.next
			a.LinkedList.length--
		}
	}

	// At this point, some values are remaining but we are sure the head is not v
	prev := a.LinkedList.head
	curr := a.LinkedList.head.next
	for {
		if curr == nil {
			return
		}

		if curr.v == v {
			prev.next = curr.next
			a.LinkedList.length--
		}
		next := curr.next
		prev = curr
		curr = next
	}
}

// Contains returns if the given value is contained in the list.
func (a *ComparableLinkedList[T]) Contains(v T) bool {
	if a.LinkedList.length == 0 {
		return false
	}

	curr := a.LinkedList.head
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

// IndexOf returns the first index of an value which is the same as given v.
// It returns negative value if v is not found in the list.
func (a *ComparableLinkedList[T]) IndexOf(v T) int {
	return -1
}

// LastIndexOf returns the last index of an value which is the same as given v.
// It returns negative value if v is not found in the list.
func (a *ComparableLinkedList[T]) LastIndexOf(v T) int {
	return -1
}
