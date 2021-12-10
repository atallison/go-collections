package collection

import (
	"testing"
)

// linkedListMustEqual makes sure the given list is logically the same as the given slice.
// For example, a linkedlist 1 -> 2 -> 3 -> nil is considered to be the same as [1, 2, 3].
func linkedListMustEqual[T any](t *testing.T, list *LinkedList[T], values []T) {
	t.Helper()
	buff := []T{}
	curr := list.head
	for {
		if curr == (*linkedNode[T])(nil) {
			break
		}
		buff = append(buff, curr.v)
		curr = curr.next
	}
	MustEqual(t, values, buff)
}

func TestLinkedList_Add(t *testing.T) {
	l := NewLinkedList[int]()
	MustEqual(t, (*linkedNode[int])(nil), l.head)
	MustEqual(t, (*linkedNode[int])(nil), l.tail)

	l.Add(1)
	MustEqual(t, 1, l.head.v)
	MustEqual(t, 1, l.tail.v)
	MustEqual(t, l.head, l.tail)
	MustEqual(t, 1, l.length)

	l.Add(2)
	MustEqual(t, 1, l.head.v)
	MustEqual(t, l.tail, l.head.next)
	MustEqual(t, 2, l.tail.v)
	MustEqual(t, 2, l.length)
	MustEqual(t, (*linkedNode[int])(nil), l.tail.next)
}

func TestLinkedList_AddHead(t *testing.T) {
	l := NewLinkedList[int]()
	MustEqual(t, (*linkedNode[int])(nil), l.head)
	MustEqual(t, (*linkedNode[int])(nil), l.tail)

	l.Add(1)
	l.Add(2)
	l.AddHead(3)
	linkedListMustEqual(t, l, []int{3, 1, 2})
	MustEqual(t, 3, l.length)

	MustEqual(t, 3, l.head.v)
	MustEqual(t, 2, l.tail.v)
}

func TestLinkedList_AddAt(t *testing.T) {
	l := NewLinkedList[int]()
	l.Add(1)
	l.Add(2)
	l.Add(3)
	linkedListMustEqual(t, l, []int{1, 2, 3})
	MustEqual(t, 3, l.length)

	MustBeErr(t, l.AddAt(-1, 4), ErrInvalidIndex)
	MustBeErr(t, l.AddAt(4, 4), ErrInvalidIndex)

	err := l.AddAt(0, 0)
	MustBeNil(t, err)
	linkedListMustEqual(t, l, []int{0, 1, 2, 3})
	MustEqual(t, 4, l.length)

	MustEqual(t, 0, l.head.v)
	MustEqual(t, 3, l.tail.v)

	err = l.AddAt(4, 4)
	MustBeNil(t, err)
	linkedListMustEqual(t, l, []int{0, 1, 2, 3, 4})
	MustEqual(t, 5, l.length)

	MustEqual(t, 0, l.head.v)
	MustEqual(t, 4, l.tail.v)

	err = l.AddAt(3, 5)
	MustBeNil(t, err)
	linkedListMustEqual(t, l, []int{0, 1, 2, 5, 3, 4})
	MustEqual(t, 6, l.length)

	MustEqual(t, 0, l.head.v)
	MustEqual(t, 4, l.tail.v)
}

func TestLinkedList_AddAll(t *testing.T) {
	l := NewLinkedList[int]()
	l.AddAll([]int{})
	linkedListMustEqual(t, l, []int{})
	MustEqual(t, 0, l.length)

	l.AddAll([]int{0})
	linkedListMustEqual(t, l, []int{0})
	MustEqual(t, 1, l.length)

	MustEqual(t, 0, l.head.v)
	MustEqual(t, 0, l.tail.v)

	l.AddAll([]int{1, 2, 3})
	linkedListMustEqual(t, l, []int{0, 1, 2, 3})
	MustEqual(t, 4, l.length)

	MustEqual(t, 0, l.head.v)
	MustEqual(t, 3, l.tail.v)

	l.AddAll([]int{4, 5, 6})
	linkedListMustEqual(t, l, []int{0, 1, 2, 3, 4, 5, 6})
	MustEqual(t, 7, l.length)

	MustEqual(t, 0, l.head.v)
	MustEqual(t, 6, l.tail.v)

	l.AddAll([]int{7})
	linkedListMustEqual(t, l, []int{0, 1, 2, 3, 4, 5, 6, 7})
	MustEqual(t, 8, l.length)

	MustEqual(t, 0, l.head.v)
	MustEqual(t, 7, l.tail.v)
}

func TestLinkedList_Clear(t *testing.T) {
	l := NewLinkedList[int]()
	l.AddAll([]int{1, 2, 3})
	linkedListMustEqual(t, l, []int{1, 2, 3})
	MustEqual(t, 3, l.length)

	l.Clear()
	linkedListMustEqual(t, l, []int{})
	MustEqual(t, 0, l.length)

	// makes sure the cleared list is still usable
	l.AddAll([]int{1, 2, 3})
	linkedListMustEqual(t, l, []int{1, 2, 3})
	MustEqual(t, 3, l.length)
}

func TestLinkedList_Clone(t *testing.T) {
	l := NewLinkedList[int]()
	l.AddAll([]int{1, 2, 3})
	linkedListMustEqual(t, l, []int{1, 2, 3})
	MustEqual(t, 3, l.length)

	nl := l.Clone()
	linkedListMustEqual(t, nl, []int{1, 2, 3})
	MustEqual(t, 3, nl.length)

	// add values into l
	l.AddAll([]int{4, 5, 6})
	linkedListMustEqual(t, l, []int{1, 2, 3, 4, 5, 6})
	MustEqual(t, 6, l.length)

	// makes sure the cloned list is not changed
	linkedListMustEqual(t, nl, []int{1, 2, 3})
	MustEqual(t, 3, nl.length)

	// add values into nl
	nl.AddAll([]int{7, 8, 9})
	linkedListMustEqual(t, nl, []int{1, 2, 3, 7, 8, 9})
	MustEqual(t, 6, nl.length)

	// make srue the original list is not changed
	linkedListMustEqual(t, l, []int{1, 2, 3, 4, 5, 6})
	MustEqual(t, 6, l.length)
}

func TestLinkedList_GetHead(t *testing.T) {
	l := NewLinkedList[int]()
	v, ok := l.GetHead()
	MustEqual(t, false, ok)
	MustEqual(t, 0, v)

	l.AddAll([]int{1, 2, 3})
	v, ok = l.GetHead()
	MustEqual(t, true, ok)
	MustEqual(t, 1, v)
	linkedListMustEqual(t, l, []int{1, 2, 3})
	MustEqual(t, 3, l.length)
}

func TestLinkedList_GetAt(t *testing.T) {
	l := NewLinkedList[int]()
	l.AddAll([]int{1, 2, 3})

	v, err := l.GetAt(-1)
	MustBeErr(t, ErrInvalidIndex, err)

	v, err = l.GetAt(3)
	MustBeErr(t, ErrInvalidIndex, err)

	v, err = l.GetAt(0)
	MustBeNil(t, err)
	MustEqual(t, 1, v)

	v, err = l.GetAt(1)
	MustBeNil(t, err)
	MustEqual(t, 2, v)

	v, err = l.GetAt(2)
	MustBeNil(t, err)
	MustEqual(t, 3, v)
}

func TestLinkedList_GetTail(t *testing.T) {
	l := NewLinkedList[int]()
	v, ok := l.GetTail()
	MustEqual(t, false, ok)
	MustEqual(t, 0, v)

	l.AddAll([]int{1})
	v, ok = l.GetTail()
	MustEqual(t, true, ok)
	MustEqual(t, 1, v)

	l.AddAll([]int{2})
	v, ok = l.GetTail()

	linkedListMustEqual(t, l, []int{1, 2})
	MustEqual(t, true, ok)
	MustEqual(t, 2, v)
}
