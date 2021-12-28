package collection

// collection interface is just an idiom to make sure that
// each types in the package has uniformed method signature for library developers.
// You don't need to depend on this type from outside of this library.
type collection[T any] interface {
	// Add adds the given value to the list.
	Add(value T)

	// Clear makes the list empty. The list after Clear must be reusable.
	Clear()

	// IsEmpty returns true if the list contains nothing.
	IsEmpty() bool

	// Iterator returns the iteratable form of the list.
	Iterator() Iterator[T]

	// Len returns the size of the list.
	Len() int

	// String shows the list in the string form.
	String() string
}

var (
	// make sure method form
	_ []collection[any] = []collection[any]{
		(*ArrayList[any])(nil),
		(*SinglyLinkedList[any])(nil),
		// (*LinkedList)(nil),
	}
)
