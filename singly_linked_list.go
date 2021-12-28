package collection

import (
	"fmt"
	"strings"
	"sync"
)

// SinglyLinkedList is an implementation of singly linked list.
type SinglyLinkedList[T any] struct {
	sync.Mutex

	head   *singlyLinkedNode[T]
	length int
}

type singlyLinkedNode[T any] struct {
	v    T
	next *singlyLinkedNode[T]
}

// NewSinglyLinkedList returns an SinglyLinkedList based on the specified type.
func NewSinglyLinkedList[T any]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{}
}

// Add appends a given value to the bottom of the list.
func (l *SinglyLinkedList[T]) Add(v T) {
	l.Lock()
	defer l.Unlock()
	if l.head == nil {
		// if head is nil, the list has no elements.
		n := &singlyLinkedNode[T]{v: v, next: nil}
		l.head = n
		l.length++
		return
	}

	var curr *singlyLinkedNode[T]
	for i := 0; i < l.length; i++ {
		if i == 0 {
			curr = l.head
		} else {
			curr = curr.next
		}

		if i == l.length-1 {
			n := &singlyLinkedNode[T]{v: v, next: nil}
			curr.next = n
			l.length++
			break
		}
	}
}

// Clear removes all the data in the list. The list is still usable after clear.
func (l *SinglyLinkedList[T]) Clear() {
	l.Lock()
	defer l.Unlock()
	l.head = nil
	l.length = 0
}

// IsEmpty returns true if the list contains no values.
func (l *SinglyLinkedList[T]) IsEmpty() bool {
	return l.head == nil
}

// Iterator returns iteratable struct.
// Note that the iterator has only a snapshot of list data as of this method is called,
// and any modification to the list won't be reflected to the iterator.
func (l *SinglyLinkedList[T]) Iterator() Iterator[T] {
	l.Lock()
	defer l.Unlock()
	return &SinglyLinkedListIterator[T]{
		list:  l,
		curr:  l.head,
		index: 0,
	}

}

// Len returns the length of the list.
func (l *SinglyLinkedList[T]) Len() int {
	return l.length
}

// String returns string form of the list.
func (l *SinglyLinkedList[T]) String() string {
	l.Lock()
	defer l.Unlock()
	if l.IsEmpty() {
		return "[]"
	}

	var sb strings.Builder
	sb.WriteString("[")
	sb.WriteString(fmt.Sprintf("%v, ", l.head))
	curr := l.head
	for {
		if curr == nil {
			break
		}

		curr = curr.next
		sb.WriteString(fmt.Sprintf("%v, ", curr))
	}

	sb.WriteString("]")
	return sb.String()
}

// AddHead inserts the given value at the head of the list.
func (l *SinglyLinkedList[T]) AddHead(v T) {
	l.Lock()
	defer l.Unlock()
	if l.head == nil {
		// if head is nil, the list has no elements.
		n := &singlyLinkedNode[T]{v: v, next: nil}
		l.head = n
		l.length++
		return
	}

	curHead := l.head
	n := &singlyLinkedNode[T]{v: v, next: curHead}
	l.head = n
	l.length++
}

// AddAll appends all the given value at the bottom.
func (l *SinglyLinkedList[T]) AddAll(vs []T) {
	l.Lock()
	defer l.Unlock()
	if len(vs) == 0 {
		return
	}

	var vhead *singlyLinkedNode[T] // vhead is always the first node which contains the head of the given values
	var curr *singlyLinkedNode[T]
	for i, v := range vs {
		n := &singlyLinkedNode[T]{v: v}
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
		l.length += len(vs)
		return
	}

	var tail *singlyLinkedNode[T]
	// else, append to the tail value
	for i := 0; i < l.length; i++ {
		if i == 0 {
			tail = l.head
		} else {
			tail = tail.next
		}

		if i == l.length-1 {
			tail.next = vhead
		}

	}
	l.length += len(vs)
}

// Clone returns the new SinglyLinkedList which the same with l.
func (l *SinglyLinkedList[T]) Clone() *SinglyLinkedList[T] {
	l.Lock()
	defer l.Unlock()
	nl := NewSinglyLinkedList[T]()

	if l.head == nil {
		return nl
	}

	var nhead *singlyLinkedNode[T]
	var ncurr *singlyLinkedNode[T]
	curr := l.head
	index := 0
	for {
		if curr == nil {
			break
		}
		n := &singlyLinkedNode[T]{v: curr.v}
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
	nl.length = l.length
	return nl
}

// GetHead returns the head value.
// If the value is not found since the list is empty, the second returned value will be false.
func (l *SinglyLinkedList[T]) GetHead() (v T, ok bool) {
	l.Lock()
	defer l.Unlock()
	if l.head == nil {
		return v, false
	}
	return l.head.v, true
}

// GetAt returns the value at the given index in the list.
func (l *SinglyLinkedList[T]) GetAt(index int) (v T, err error) {
	l.Lock()
	defer l.Unlock()
	if index < 0 || l.length <= index {
		return v, ErrInvalidIndex
	}

	if index == 0 {
		return l.head.v, nil
	}

	curr := l.head
	for i := 0; i < index; i++ {
		curr = curr.next
	}

	return curr.v, nil
}

// RemoveAt removes a value at the given index in the list.
// ErrInvalidIndex will be responded if the index < 0 or length <= index.
func (l *SinglyLinkedList[T]) RemoveAt(index int) error {
	if index < 0 || l.length <= index {
		return ErrInvalidIndex
	}

	l.length--

	// we must take care of the case index == 0 because
	// l.head must be changed in the case
	if index == 0 {
		l.head = l.head.next
		return nil
	}

	// else, we will need to take care of changing next, instead of head
	curr := l.head
	// After this loop, curr will be list[index-1]
	// Let's say l is [1, 2, 3, 4, 5], if index is 2, curr will be 2 after the loop
	for i := 0; i < index-1; i++ {
		curr = curr.next
	}

	removed := curr.next
	curr.next = removed.next
	return nil
}

// Set replaces the value at the given index in the list with the given value.
// ErrInvalidIndex will be responded if the index < 0 or length <= index.
func (l *SinglyLinkedList[T]) Set(index int, v T) error {
	if index < 0 || l.length <= index {
		return ErrInvalidIndex
	}

	if index == 0 {
		l.head.v = v
		return nil
	}

	var curr *singlyLinkedNode[T]
	for i := 0; i < index+1; i++ {
		if i == 0 {
			curr = l.head
		} else {
			curr = curr.next
		}
	}
	curr.v = v
	return nil
}
