package collection

// LinkedList is an implementation of doubly linked list.
// This is not concurrent safe.
type LinkedList[T any] struct {
	// dummy.next is head, dummy.prev is tail.
	dummy  *linkedNode[T]
	length int
}

// linkedNode is a node in LinkedList.
type linkedNode[T any] struct {
	v        T
	prevNode *linkedNode[T]
	nextNode *linkedNode[T]
}

func (n *linkedNode[T]) previous() *linkedNode[T] {
	return n.prevNode
}

func (n *linkedNode[T]) next() *linkedNode[T] {
	return n.nextNode
}

// insertAfter inserts given valud to the next to n.
// Newly added node will be returned.
func (n *linkedNode[T]) insertAfter(v T) *linkedNode[T] {
	m := &linkedNode[T]{
		v:        v,
		prevNode: n,
		nextNode: n.nextNode,
	}
	n.nextNode.prevNode = m
	n.nextNode = m
	return m
}

// NewLinkedList returns an empty LinkedList based on the specified type T.
func NewLinkedList[T any]() *LinkedList[T] {
	var zero T // zero value
	dummy := &linkedNode[T]{v: zero}
	dummy.nextNode, dummy.prevNode = dummy, dummy
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
