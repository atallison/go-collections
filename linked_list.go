package collection

// LinkedList is an implementation of singly linked list.
// This is not concurrent safe.
type LinkedList[T any] struct {
	dummy  *linkedNode[T]
	length int
}

type linkedNode[T any] struct {
	v    T
	next *linkedNode[T]
}
