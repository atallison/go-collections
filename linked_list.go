package collection

// LinkedList is an implementation of doubly linked list.
// This is not concurrent safe.
type LinkedList[T any] struct {
	dummy  *singlyLinkedNode[T]
	length int
}

type linkedNode[T any] struct {
	v    T
	next *linkedNode[T]
	prev *linkedNode[T]
}
