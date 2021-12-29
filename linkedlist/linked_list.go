package linkedlist

// LinkedList is an implementation of double ended, doubly linked list.
type LinkedList[T any] struct {
	// dummy.next is head, dummy.prev is tail.
	dummy  *Node[T]
	length int
}

// Node is a node in LinkedList.
type Node[T any] struct {
	Value T
	prev  *Node[T]
	next  *Node[T]
}

// New returns an empty LinkedList based on the specified type T.
func New[T any]() *LinkedList[T] {
	dummy := &Node[T]{}
	dummy.next, dummy.prev = dummy, dummy
	return &LinkedList[T]{
		dummy:  dummy,
		length: 0,
	}
}

func (l *LinkedList[T]) insertAfter(n, at *Node[T]) *Node[T] {
	n.prev = at
	n.next = at.next
	n.prev.next = n
	n.next.prev = n
	l.length++
	return n
}

// PushBack pushes the value at the bottom and return the pushed node.
// This operations is O(1) because this LinkedList is double-ended.
func (l *LinkedList[T]) PushBack(value T) *Node[T] {
	return l.insertAfter(&Node[T]{Value: value}, l.dummy.prev)
}

// PushFron pushes the value at the head and return the pushed node.
func (l *LinkedList[T]) PushFront(value T) *Node[T] {
	return l.insertAfter(&Node[T]{Value: value}, l.dummy)
}

// Len returns the size of the list.
func (l *LinkedList[T]) Len() int {
	return l.length
}
