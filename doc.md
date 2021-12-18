# package collection // import "github.com/hidetatz/collection"


## VARIABLES

```go
var (
	// ErrInvalidIndex indicates the given index is invalid for the operation.
	ErrInvalidIndex = errors.New("invalid index")

	// ErrHeadNotFound indicates the head node is not found.
	ErrHeadNotFound = errors.New("head not found")

	// ErrTailNotFound indicates the tail node is not found.
	ErrTailNotFound = errors.New("tail not found")
)
```

## FUNCTIONS

### func MustBeErr(t *testing.T, expected, got error)
    MustBeErr let the test fail if an underlying error of the given error is not
    the same with the expected error.

### func MustBeNil(t *testing.T, v interface{})
    MustBeNil let the test fail if the given value is non nil.

### func MustEqual(t *testing.T, expected, got interface{})
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

### func (a *ArrayList[T]) Add(e T)
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
    Clear removes all the data in the list. The list is still usable after
    clear.

### func (a *ArrayList[T]) Clone() *ArrayList[T]
    Clone clones the arraylist and return it. Modifying cloned list does not
    affect to the original one and vice versa.

### func (a *ArrayList[T]) Filter(f func(index int, v T) bool) *ArrayList[T]
    Filter returns a new ArrayList which only contain filtered value by the
    given f. Filter does not break the original ArrayList.

### func (a *ArrayList[T]) ForEach(f func(index int, v T))
    ForEach runs the given f to each value in the list.

### func (a *ArrayList[T]) Get(index int) (ret T, err error)
    Get returns a value which locates at the given index of the list.
    ErrInvalidIndex will be responded if the index < 0 or length <= index.

### func (a *ArrayList[T]) IsEmpty() bool
    IsEmpty returns true if the list is empty.

### func (a *ArrayList[T]) Iterator() *ArrayListIterator[T]
    Iterator returns iteratable struct based on current Arraylist. Note that the
    iterator has only a snapshot of list data as of this method is called, and
    any modification to the list won't be reflected to the iterator.

### func (a *ArrayList[T]) Len() int
    Len returns the number of the elements in the list.

### func (a *ArrayList[T]) Map(f func(v T) T) *ArrayList[T]
    Map applies f to the list values then return it as a new ArrayList. Map does
    not break the original arraylist. If you want to do that, you can use
    ReplaceAll.

### func (a *ArrayList[T]) RemoveAt(index int) error
    RemoveAt removes a value at the given index in the list. ErrInvalidIndex
    will be responded if the index < 0 or length <= index.

### func (a *ArrayList[T]) RemoveIf(f func(index int, v T) bool)
    RemoveIf removes values if the f(value) returns true.

### func (a *ArrayList[T]) RemoveRange(from, to int) error
    RemoveRange removes values whose index is between from (inclusive) and to
    (exclusive). The same values as from and to can be passed but it has no
    effect and no error will be responded. ErrInvalidIndex will be responded if
    the given index is invalid.

### func (a *ArrayList[T]) ReplaceAll(f func(v T) T)
    ReplaceAll applies the given f to all the values in the list.

### func (a *ArrayList[T]) Set(index int, v T) error
    Set replaces the value at the given index in the list with the given value.
    ErrInvalidIndex will be responded if the index < 0 or length <= index.

### func (a *ArrayList[T]) Slice() []T
    Slice returns underlying slice in the arraylist.

### func (a *ArrayList[T]) SubList(from, to int) (*ArrayList[T], error)
    SubList returns the new ArrayList which contains the elements in the list
    between the specified fromIndex (inclusive) and toIndex (exclusive).

```go
type ArrayListIterator[T any] struct {
	// Has unexported fields.
}
```
    ArrayListIterator is an implementation of Iterator. Because ArrayList is not
    concurrent safe, ArrayListIterator is not coucurrent safe as well.

### func (ai *ArrayListIterator[T]) Next() bool

### func (ai *ArrayListIterator[T]) Value() T

```go
type ComparableArrayList[T comparable] struct {
	*ArrayList[T]
}
```

### func NewComparableArrayList[T comparable]() *ComparableArrayList[T]
    NewComparableArrayList returns an ArrayList based on the specified type
    which is comparable.

### func (a *ComparableArrayList[T]) Contains(v T) bool
    Contains returns if the given value is contained in the list.

### func (a *ComparableArrayList[T]) IndexOf(v T) int
    IndexOf returns the first index of an value which is the same as given v. It
    returns negative value if v is not found in the list.

### func (a *ComparableArrayList[T]) LastIndexOf(v T) int
    LastIndexOf returns the last index of an value which is the same as given v.
    It returns negative value if v is not found in the list.

### func (a *ComparableArrayList[T]) Remove(v T)
    Remove removes the same value with given v in the list. It uses == operator
    to make sure if the values are the same.

```go
type Iterator[T any] interface {
	Next() bool
	Value() T
}
```

```go
type LinkedList[T any] struct {
	// Has unexported fields.
}
```
    LinkedList is an implementation of singly linked list. This is not
    concurrent safe.

### func NewLinkedList[T any]() *LinkedList[T]
    NewLinkedList returns an ArrayList based on the specified type.

### func (l *LinkedList[T]) Add(v T)
    Add appends a given value to the bottom of the list. This is O(1) because
    LinkedList internally has the pointer to the tail node.

### func (l *LinkedList[T]) AddAll(vs []T)
    AddAll appends all the given value at the bottom.

### func (l *LinkedList[T]) AddAt(index int, v T) error
    AddAt appends a given value at the given index position in the list.

### func (l *LinkedList[T]) AddHead(v T)
    AddHead inserts the given value at the head of the list.

### func (l *LinkedList[T]) Clear()
    Clear removes all the data in the list. The list is still usable after
    clear.

### func (l *LinkedList[T]) Clone() *LinkedList[T]
    Clone returns the new LinkedList which the same with l.

### func (l *LinkedList[T]) GetAt(index int) (v T, err error)
    GetAt returns the value at the given index in the list.

### func (l *LinkedList[T]) GetHead() (v T, ok bool)
    GetHead returns the head value. If the value is not found since the list is
    empty, the second returned value will be false.

### func (l *LinkedList[T]) GetTail() (v T, ok bool)
    GetTail returns the tail value. If the value is not found since the list is
    empty, the second returned value will be false.

### func (l *LinkedList[T]) IsEmpty() bool
    IsEmpty returns true if the list contains no values.

### func (l *LinkedList[T]) Iterator() *LinkedListIterator[T]
    Iterator returns iteratable struct. Note that the iterator has only a
    snapshot of list data as of this method is called, and any modification to
    the list won't be reflected to the iterator.

### func (l *LinkedList[T]) Len() int
    Len returns the length of the list.

### func (l *LinkedList[T]) RemoveAt(index int) error
    RemoveAt removes a value at the given index in the list. ErrInvalidIndex
    will be responded if the index < 0 or length <= index.

### func (l *LinkedList[T]) RemoveHead() error
    RemoveHead removes the head value.

### func (l *LinkedList[T]) RemoveTail() error
    RemoveTail removes the litailvalude.

### func (l *LinkedList[T]) Set(index int, v T) error
    Set replaces the value at the given index in the list with the given value.
    ErrInvalidIndex will be responded if the index < 0 or length <= index.

```go
type LinkedListIterator[T any] struct {
	// Has unexported fields.
}
```
    LinkedListIterator is the iteratable made of LinkedList.

### func (i *LinkedListIterator[T]) Next() bool

### func (i *LinkedListIterator[T]) Value() T
