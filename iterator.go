package collection

type Iterator[T any] interface {
	Next() bool
	Value() T
	Set(value T)
	Remove()
}

var _ []Iterator[any] = []Iterator[any]{
	(*ArrayListIterator[any])(nil),
	(*SinglyLinkedListIterator[any])(nil),
}
