package collection

// LinkedList is an implementation of doubly linked list.
// This is not concurrent safe.
type LinkedList[T any] struct {
	// dummy.next is head, dummy.prev is tail.
	dummy  *linkedNode[T]
	length int
}

type linkedNode[T any] struct {
	v    T
	next *linkedNode[T]
	prev *linkedNode[T]
}

func (n *linkedNode[T]) next() *linkedNode[T] {
	return n.next
}

func (n *linkedNode[T]) previous() *linkedNode[T] {
	return n.prev
}

func NewLinkedList[T any]() *LinkedList[T] {
	var zero T
	dummy := &linkedNode[T]{v: zero}
	dummy.next, dummy.prev = dummy, dummy
	return &LinkedList[T]{
		dummy:  dummy,
		length: 0,
	}
}

func (l *LinkedList[T]) Head() T {
	return l.dummy.next().v
}

func (l *LinkedList[T]) Tail() T {
	return l.dummy.previous().v
}
