package collection

import (
	"testing"
)

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

	err = l.AddAt(4, 4)
	MustBeNil(t, err)
	linkedListMustEqual(t, l, []int{0, 1, 2, 3, 4})
	MustEqual(t, 5, l.length)

	err = l.AddAt(3, 5)
	MustBeNil(t, err)
	linkedListMustEqual(t, l, []int{0, 1, 2, 5, 3, 4})
	MustEqual(t, 6, l.length)
}
