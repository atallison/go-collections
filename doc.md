# package collection // import "github.com/hidetatz/collection"


## VARIABLES

```go
var (
	// ErrInvalidIndex indicates the given index is invalid for the operation.
	ErrInvalidIndex = errors.New("invalid index")

	// ErrHeadNotFound indicates the head node is not found.
	ErrHeadNotFound = errors.New("head not found")
)
```

## FUNCTIONS

### func MustBeErr(t *testing.T, expected, actual error)
    MustBeErr let the test fail if an underlying error of the given error is not
    the same with the expected error.

### func MustBeNil(t *testing.T, v interface{})
    MustBeNil let the test fail if the given value is non nil.

### func MustEqual(t *testing.T, expected, actual interface{})
    MustEqual let the test fail if the given values are not the same.


## TYPES

```go
type ArrayList[T any] struct {
	// Has unexported fields.
}
```
    ArrayList is a variable-sized list. This data structure is not concurrent
    safe. The caller must be responsible to synchronize before accessing from
    multiple goroutines.

### func NewArrayList[T any]() *ArrayList[T]
    NewArrayList returns an ArrayList based on the specified type.

### func (a *ArrayList[T]) Add(value T)
    Add appends a given value to the bottom of the list.

### func (a *ArrayList[T]) AddAll(es []T)
    AddAll appends given values to the bottom of the list in the given order.

### func (a *ArrayList[T]) AddAllAt(i int, es []T) error
    AddAllAt inserts the given values at the given index. Let's say current list
    looks [a, b, c] then calls AddAllAt(1, [d, e]), it results [a, d, e, b, c].
    When the given index is less than 0 or greater than the list size,
    ErrInvalidIndex will be responded.

### func (a *ArrayList[T]) AddAt(i int, e T) error
    AddAt inserts the given element at the given index. When the given is less
    than 0 or greater than the list size, ErrInvalidIndex will be responded.

### func (a *ArrayList[T]) Clear()
    Clear makes the list empty. The list after Clear must be reusable.

### func (a *ArrayList[T]) Clone() *ArrayList[T]
    Clone clones the arraylist and return it. Modifying cloned list does not
    affect to the original one and vice versa.

### func (a *ArrayList[T]) Get(index int) (ret T, err error)
    Get returns a value which locates at the given index of the list.
    ErrInvalidIndex will be responded if the index < 0 or length <= index.

### func (a *ArrayList[T]) IsEmpty() bool
    IsEmpty returns true if the list contains nothing.

### func (a *ArrayList[T]) Iterator() Iterator[T]
    Iterator returns iteratable data structure based on the list.

### func (a *ArrayList[T]) Len() int
    Len returns the size of the list.

### func (a *ArrayList[T]) RemoveAt(index int) error
    RemoveAt removes a value at the given index in the list. ErrInvalidIndex
    will be responded if the index < 0 or length <= index.

### func (a *ArrayList[T]) RemoveRange(from, to int) error
    RemoveRange removes values whose index is between from (inclusive) and to
    (exclusive). The same values as from and to can be passed but it has no
    effect and no error will be responded. ErrInvalidIndex will be responded if
    the given index is invalid.

### func (a *ArrayList[T]) Set(index int, v T) error
    Set replaces the value at the given index in the list with the given value.
    ErrInvalidIndex will be responded if the index < 0 or length <= index.

### func (a *ArrayList[T]) String() string
    String shows the list in the string form.

### func (a *ArrayList[T]) SubList(from, to int) (*ArrayList[T], error)
    SubList returns the new ArrayList which contains the elements in the list
    between the specified fromIndex (inclusive) and toIndex (exclusive).

```go
type ArrayListIterator[T any] struct {
	// Has unexported fields.
}
```
    ArrayListIterator is an implementation of Iterator for ArrayList.

### func (ai *ArrayListIterator[T]) Next() bool
    Next returns if the next value exists in the iterator.

### func (ai *ArrayListIterator[T]) Remove()
    Remove removes the value which is returned by the last call of Value() from
    the original ArrayList. Note that this must be called at most once per a
    Value() call. Unless, it might lead to list/iterator inconsistent or invalid
    state and that case is not tested.

### func (ai *ArrayListIterator[T]) Set(value T)
    Set sets the value which is returned by the last call of Value() from the
    original ArrayList safely. Note that this must be called at most once per a
    Value() call. Unless, it might lead to list/iterator inconsistent or invalid
    state and that case is not tested.

### func (ai *ArrayListIterator[T]) Value() T
    Value returns the next value in the iterator.

```go
type Iterator[T any] interface {
	Next() bool
	Value() T
	Set(value T)
	Remove()
}
```

```go
type LinkedList[T any] struct {
	// Has unexported fields.
}
```
    LinkedList is an implementation of doubly linked list. This is not
    concurrent safe.

### func NewLinkedList[T any]() *LinkedList[T]
    NewLinkedList returns an empty LinkedList based on the specified type T.

### func (l *LinkedList[T]) AddHead(v T) T

### func (l *LinkedList[T]) Head() T

### func (l *LinkedList[T]) Tail() T

```go
type LinkedNode[T any] struct {
	Value T

	// Has unexported fields.
}
```
    LinkedNode is a node in LinkedList.

```go
type SinglyLinkedList[T any] struct {
	// Has unexported fields.
}
```
    SinglyLinkedList is an implementation of singly linked list. This is not
    concurrent safe.

### func NewSinglyLinkedList[T any]() *SinglyLinkedList[T]
    NewSinglyLinkedList returns an ArrayList based on the specified type.

### func (l *SinglyLinkedList[T]) Add(v T)
    Add appends a given value to the bottom of the list. This is O(1) because
    SinglyLinkedList internally has the pointer to the tail node.

### func (l *SinglyLinkedList[T]) AddAll(vs []T)
    AddAll appends all the given value at the bottom.

### func (l *SinglyLinkedList[T]) AddAt(index int, v T) error
    AddAt appends a given value at the given index position in the list.

### func (l *SinglyLinkedList[T]) AddHead(v T)
    AddHead inserts the given value at the head of the list.

### func (l *SinglyLinkedList[T]) Clear()
    Clear removes all the data in the list. The list is still usable after
    clear.

### func (l *SinglyLinkedList[T]) Clone() *SinglyLinkedList[T]
    Clone returns the new SinglyLinkedList which the same with l.

### func (l *SinglyLinkedList[T]) GetAt(index int) (v T, err error)
    GetAt returns the value at the given index in the list.

### func (l *SinglyLinkedList[T]) GetHead() (v T, ok bool)
    GetHead returns the head value. If the value is not found since the list is
    empty, the second returned value will be false.

### func (l *SinglyLinkedList[T]) IsEmpty() bool
    IsEmpty returns true if the list contains no values.

### func (l *SinglyLinkedList[T]) Iterator() Iterator[T]
    Iterator returns iteratable struct. Note that the iterator has only a
    snapshot of list data as of this method is called, and any modification to
    the list won't be reflected to the iterator.

### func (l *SinglyLinkedList[T]) Len() int
    Len returns the length of the list.

### func (l *SinglyLinkedList[T]) RemoveAt(index int) error
    RemoveAt removes a value at the given index in the list. ErrInvalidIndex
    will be responded if the index < 0 or length <= index.

### func (l *SinglyLinkedList[T]) RemoveHead() error
    RemoveHead removes the head value.

### func (l *SinglyLinkedList[T]) Set(index int, v T) error
    Set replaces the value at the given index in the list with the given value.
    ErrInvalidIndex will be responded if the index < 0 or length <= index.

### func (l *SinglyLinkedList[T]) String() string
    String returns string form of the list.

```go
type SinglyLinkedListIterator[T any] struct {
	// Has unexported fields.
}
```
    SinglyLinkedListIterator is the iteratable made of SinglyLinkedList.

### func (i *SinglyLinkedListIterator[T]) Next() bool

### func (i *SinglyLinkedListIterator[T]) Remove()

### func (i *SinglyLinkedListIterator[T]) Set(value T)

### func (i *SinglyLinkedListIterator[T]) Value() T
