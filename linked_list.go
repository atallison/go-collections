package collection

// LinkedList is an implementation of singly linked list.
// This is not concurrent safe.
type LinkedList[T any] struct {
	head   *linkedNode[T]
	tail   *linkedNode[T]
	length int
}

type linkedNode[T any] struct {
	v    T
	next *linkedNode[T]
}

// NewLinkedList returns an ArrayList based on the specified type.
func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

// Add appends a given value to the bottom of the list.
// This is O(1) because LinkedList internally has the pointer to the tail node.
func (l *LinkedList[T]) Add(v T) {
	if l.head == nil {
		// if head is nil, the list has no elements.
		n := &linkedNode[T]{v: v, next: nil}
		l.head = n
		l.tail = n
		l.length++
		return
	}

	n := &linkedNode[T]{v: v, next: nil}
	l.tail.next = n
	l.tail = n
	l.length++
}

// AddHead inserts the given value at the head of the list.
func (l *LinkedList[T]) AddHead(v T) {
	if l.head == nil {
		// if head is nil, the list has no elements.
		n := &linkedNode[T]{v: v, next: nil}
		l.head = n
		l.tail = n
		l.length++
		return
	}

	curHead := l.head
	n := &linkedNode[T]{v: v, next: curHead}
	l.head = n
	l.length++
}

// AddAt appends a given value at the given index position in the list.
func (l *LinkedList[T]) AddAt(index int, v T) error {
	if index < 0 || l.length < index {
		return ErrInvalidIndex
	}

	if index == l.length {
		l.Add(v)
		return nil
	}

	if index == 0 {
		l.AddHead(v)
		return nil
	}

	n := &linkedNode[T]{v: v}

	var curr *linkedNode[T]
	for i := 0; i < index; i++ {
		if i == 0 {
			curr = l.head
		} else {
			curr = curr.next
		}
	}

	next := curr.next
	curr.next = n
	n.next = next
	l.length++
	return nil
}

// AddAll appends all the given value at the bottom.
func (l *LinkedList[T]) AddAll(vs []T) {
	if len(vs) == 0 {
		return
	}

	var vhead *linkedNode[T] // vhead is always the first node which contains the head of the given values
	var curr *linkedNode[T]  // curr is eventually the last value
	for i, v := range vs {
		n := &linkedNode[T]{v: v}
		if i == 0 {
			curr = n
			vhead = curr
		} else {
			curr.next = n
			curr = n
		}
	}

	if l.head == nil {
		// if list is empty
		l.head = vhead
		l.tail = curr
		l.length += len(vs)
		return
	}

	// else, append
	l.tail.next = vhead
	l.tail = curr
	l.length += len(vs)
}

// Clear removes all the data in the list. The list is still usable after clear.
func (l *LinkedList[T]) Clear() {
	l.head = nil
	l.tail = nil
	l.length = 0
}

// Clone returns the new LinkedList which the same with l.
func (l *LinkedList[T]) Clone() *LinkedList[T] {
	nl := NewLinkedList[T]()

	if l.head == nil {
		return nl
	}

	var nhead *linkedNode[T]
	var ncurr *linkedNode[T]
	curr := l.head
	index := 0
	for {
		if curr == nil {
			break
		}
		n := &linkedNode[T]{v: curr.v}
		if index == 0 {
			nhead = n
			ncurr = n
		} else {
			ncurr.next = n
			ncurr = n
		}
		curr = curr.next
		index++
	}

	nl.head = nhead
	nl.tail = ncurr
	nl.length = l.length
	return nl
}

// GetHead returns the head value.
// If the value is not found since the list is empty, the second returned value will be false.
func (l *LinkedList[T]) GetHead() (v T, ok bool) {
	if l.head == nil {
		return v, false
	}
	return l.head.v, true
}

// GetAt returns the value at the given index in the list.
func (l *LinkedList[T]) GetAt(index int) (v T, err error) {
	if index < 0 || l.length <= index {
		return v, ErrInvalidIndex
	}

	if index == 0 {
		return l.head.v, nil
	}

	if index == l.length-1 {
		return l.tail.v, nil
	}

	curr := l.head
	for i := 0; i < index; i++ {
		curr = curr.next
	}

	return curr.v, nil
}

// GetTail returns the tail value.
// If the value is not found since the list is empty, the second returned value will be false.
func (l *LinkedList[T]) GetTail() (v T, ok bool) {
	if l.tail == nil {
		return v, false
	}
	return l.tail.v, true
}
