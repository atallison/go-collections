package collection

// LinkedList is an implementation of doubly linked list.
type LinkedList[T any] struct {
	// dummy.next is head, dummy.prev is tail.
	dummy  *LinkedNode[T]
	length int
}

// LinkedNode is a node in LinkedList.
type LinkedNode[T any] struct {
	Value    T
	prevNode *LinkedNode[T]
	nextNode *LinkedNode[T]
}

func (n *LinkedNode[T]) previous() *LinkedNode[T] {
	return n.prevNode
}

func (n *LinkedNode[T]) next() *LinkedNode[T] {
	return n.nextNode
}

// NewLinkedList returns an empty LinkedList based on the specified type T.
func NewLinkedList[T any]() *LinkedList[T] {
	var zero T // zero value
	dummy := &LinkedNode[T]{Value: zero}
	dummy.nextNode, dummy.prevNode = dummy, dummy
	return &LinkedList[T]{
		dummy:  dummy,
		length: 0,
	}
}

// insertAfter inserts given valud to the next to n.
// Newly added node will be returned.
func (l *LinkedList[T]) insertAfter(n *LinkedNode[T], v T) *LinkedNode[T] {
	m := &LinkedNode[T]{
		Value:    v,
		prevNode: n,
		nextNode: n.nextNode,
	}
	n.nextNode.prevNode = m
	n.nextNode = m
	return m
}

// insertBefore inserts given valud to the previous to n.
// Newly added node will be returned.
func (l *LinkedList[T]) insertBefore(n *LinkedNode[T], v T) *LinkedNode[T] {
	m := &LinkedNode[T]{
		Value:    v,
		prevNode: n.prevNode,
		nextNode: n,
	}
	n.prevNode.nextNode = m
	n.prevNode = m
	return m
}

// delete deletes the node n.
// The node next to the deleted node will be returned.
func (l *LinkedList[T]) delete(n *LinkedNode[T]) *LinkedNode[T] {
	if n == l.dummy {
		return l.dummy
	}
	n.prevNode.nextNode = n.nextNode
	n.nextNode.prevNode = n.prevNode
	return n.nextNode
}

func (l *LinkedList[T]) Head() T {
	return l.dummy.next().Value
}

func (l *LinkedList[T]) Tail() T {
	return l.dummy.previous().Value
}

func (l *LinkedList[T]) AddHead(v T) T {
	return l.dummy.previous().Value
}
