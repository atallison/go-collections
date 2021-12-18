package collection

import (
	"testing"
)

func TestSinglyLinkedList_Add(t *testing.T) {
	l := NewSinglyLinkedList[int]()
	MustEqual(t, (*singlyLinkedNode[int])(nil), l.head)
	MustEqual(t, (*singlyLinkedNode[int])(nil), l.tail)

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
	MustEqual(t, (*singlyLinkedNode[int])(nil), l.tail.next)
}

func TestSinglyLinkedList_AddHead(t *testing.T) {
	l := NewSinglyLinkedList[int]()
	MustEqual(t, (*singlyLinkedNode[int])(nil), l.head)
	MustEqual(t, (*singlyLinkedNode[int])(nil), l.tail)

	l.Add(1)
	l.Add(2)
	l.AddHead(3)
	iteratorMustEqual[int](t, l.Iterator(), []int{3, 1, 2})
	MustEqual(t, 3, l.length)

	MustEqual(t, 3, l.head.v)
	MustEqual(t, 2, l.tail.v)
}

func TestSinglyLinkedList_AddAt(t *testing.T) {
	l := NewSinglyLinkedList[int]()
	l.Add(1)
	l.Add(2)
	l.Add(3)
	iteratorMustEqual[int](t, l.Iterator(), []int{1, 2, 3})
	MustEqual(t, 3, l.length)

	MustBeErr(t, l.AddAt(-1, 4), ErrInvalidIndex)
	MustBeErr(t, l.AddAt(4, 4), ErrInvalidIndex)

	err := l.AddAt(0, 0)
	MustBeNil(t, err)
	iteratorMustEqual[int](t, l.Iterator(), []int{0, 1, 2, 3})
	MustEqual(t, 4, l.length)

	MustEqual(t, 0, l.head.v)
	MustEqual(t, 3, l.tail.v)

	err = l.AddAt(4, 4)
	MustBeNil(t, err)
	iteratorMustEqual[int](t, l.Iterator(), []int{0, 1, 2, 3, 4})
	MustEqual(t, 5, l.length)

	MustEqual(t, 0, l.head.v)
	MustEqual(t, 4, l.tail.v)

	err = l.AddAt(3, 5)
	MustBeNil(t, err)
	iteratorMustEqual[int](t, l.Iterator(), []int{0, 1, 2, 5, 3, 4})
	MustEqual(t, 6, l.length)

	MustEqual(t, 0, l.head.v)
	MustEqual(t, 4, l.tail.v)
}

func TestSinglyLinkedList_AddAll(t *testing.T) {
	l := NewSinglyLinkedList[int]()
	l.AddAll([]int{})
	iteratorMustEqual[int](t, l.Iterator(), []int{})
	MustEqual(t, 0, l.length)

	l.AddAll([]int{0})
	iteratorMustEqual[int](t, l.Iterator(), []int{0})
	MustEqual(t, 1, l.length)

	MustEqual(t, 0, l.head.v)
	MustEqual(t, 0, l.tail.v)

	l.AddAll([]int{1, 2, 3})
	iteratorMustEqual[int](t, l.Iterator(), []int{0, 1, 2, 3})
	MustEqual(t, 4, l.length)

	MustEqual(t, 0, l.head.v)
	MustEqual(t, 3, l.tail.v)

	l.AddAll([]int{4, 5, 6})
	iteratorMustEqual[int](t, l.Iterator(), []int{0, 1, 2, 3, 4, 5, 6})
	MustEqual(t, 7, l.length)

	MustEqual(t, 0, l.head.v)
	MustEqual(t, 6, l.tail.v)

	l.AddAll([]int{7})
	iteratorMustEqual[int](t, l.Iterator(), []int{0, 1, 2, 3, 4, 5, 6, 7})
	MustEqual(t, 8, l.length)

	MustEqual(t, 0, l.head.v)
	MustEqual(t, 7, l.tail.v)
}

func TestSinglyLinkedList_Clear(t *testing.T) {
	l := NewSinglyLinkedList[int]()
	l.AddAll([]int{1, 2, 3})
	iteratorMustEqual[int](t, l.Iterator(), []int{1, 2, 3})
	MustEqual(t, 3, l.length)

	l.Clear()
	iteratorMustEqual[int](t, l.Iterator(), []int{})
	MustEqual(t, 0, l.length)

	// makes sure the cleared list is still usable
	l.AddAll([]int{1, 2, 3})
	iteratorMustEqual[int](t, l.Iterator(), []int{1, 2, 3})
	MustEqual(t, 3, l.length)
}

func TestSinglyLinkedList_Clone(t *testing.T) {
	l := NewSinglyLinkedList[int]()
	l.AddAll([]int{1, 2, 3})
	iteratorMustEqual[int](t, l.Iterator(), []int{1, 2, 3})
	MustEqual(t, 3, l.length)

	nl := l.Clone()
	iteratorMustEqual[int](t, nl.Iterator(), []int{1, 2, 3})
	MustEqual(t, 3, nl.length)

	// add values into l
	l.AddAll([]int{4, 5, 6})
	iteratorMustEqual[int](t, l.Iterator(), []int{1, 2, 3, 4, 5, 6})
	MustEqual(t, 6, l.length)

	// makes sure the cloned list is not changed
	iteratorMustEqual[int](t, nl.Iterator(), []int{1, 2, 3})
	MustEqual(t, 3, nl.length)

	// add values into nl
	nl.AddAll([]int{7, 8, 9})
	iteratorMustEqual[int](t, nl.Iterator(), []int{1, 2, 3, 7, 8, 9})
	MustEqual(t, 6, nl.length)

	// make srue the original list is not changed
	iteratorMustEqual[int](t, l.Iterator(), []int{1, 2, 3, 4, 5, 6})
	MustEqual(t, 6, l.length)
}

func TestSinglyLinkedList_GetHead(t *testing.T) {
	l := NewSinglyLinkedList[int]()
	v, ok := l.GetHead()
	MustEqual(t, false, ok)
	MustEqual(t, 0, v)

	l.AddAll([]int{1, 2, 3})
	v, ok = l.GetHead()
	MustEqual(t, true, ok)
	MustEqual(t, 1, v)
	iteratorMustEqual[int](t, l.Iterator(), []int{1, 2, 3})
	MustEqual(t, 3, l.length)
}

func TestSinglyLinkedList_GetAt(t *testing.T) {
	l := NewSinglyLinkedList[int]()
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

func TestSinglyLinkedList_GetTail(t *testing.T) {
	l := NewSinglyLinkedList[int]()
	v, ok := l.GetTail()
	MustEqual(t, false, ok)
	MustEqual(t, 0, v)

	l.AddAll([]int{1})
	v, ok = l.GetTail()
	MustEqual(t, true, ok)
	MustEqual(t, 1, v)

	l.AddAll([]int{2})
	v, ok = l.GetTail()

	iteratorMustEqual[int](t, l.Iterator(), []int{1, 2})
	MustEqual(t, true, ok)
	MustEqual(t, 2, v)
}

func TestSinglyLinkedList_IsEmpty(t *testing.T) {
	l := NewSinglyLinkedList[int]()
	MustEqual(t, true, l.IsEmpty())
	l.AddAll([]int{1})
	MustEqual(t, false, l.IsEmpty())
	l.Clear()
	MustEqual(t, true, l.IsEmpty())
}

func TestSinglyLinkedList_Len(t *testing.T) {
	l := NewSinglyLinkedList[int]()
	MustEqual(t, 0, l.Len())
	l.Add(1)
	MustEqual(t, 1, l.Len())
	l.AddAll([]int{2, 3})
	MustEqual(t, 3, l.Len())
	l.Clear()
	MustEqual(t, 0, l.Len())
}

func TestSinglyLinkedList_RemoveAt(t *testing.T) {
	l := NewSinglyLinkedList[int]()
	l.AddAll([]int{1, 2, 3, 4, 5})
	err := l.RemoveAt(-1)
	MustBeErr(t, err, ErrInvalidIndex)

	err = l.RemoveAt(5)
	MustBeErr(t, err, ErrInvalidIndex)

	err = l.RemoveAt(2)
	MustBeNil(t, err)
	iteratorMustEqual[int](t, l.Iterator(), []int{1, 2, 4, 5})

	err = l.RemoveAt(3)
	MustBeNil(t, err)
	iteratorMustEqual[int](t, l.Iterator(), []int{1, 2, 4})

	err = l.RemoveAt(0)
	MustBeNil(t, err)
	iteratorMustEqual[int](t, l.Iterator(), []int{2, 4})
}

func TestSinglyLinkedList_RemoveHead(t *testing.T) {
	l := NewSinglyLinkedList[int]()
	err := l.RemoveHead()
	MustBeErr(t, err, ErrHeadNotFound)

	l.AddAll([]int{1, 2, 3, 4, 5})

	err = l.RemoveHead()
	MustBeNil(t, err)
	iteratorMustEqual[int](t, l.Iterator(), []int{2, 3, 4, 5})

	err = l.RemoveHead()
	MustBeNil(t, err)
	iteratorMustEqual[int](t, l.Iterator(), []int{3, 4, 5})

	err = l.RemoveHead()
	MustBeNil(t, err)
	iteratorMustEqual[int](t, l.Iterator(), []int{4, 5})

	err = l.RemoveHead()
	MustBeNil(t, err)
	iteratorMustEqual[int](t, l.Iterator(), []int{5})

	err = l.RemoveHead()
	MustBeNil(t, err)
	iteratorMustEqual[int](t, l.Iterator(), []int{})

	err = l.RemoveHead()
	MustBeErr(t, err, ErrHeadNotFound)
}

func TestSinglyLinkedList_RemoveTail(t *testing.T) {
	l := NewSinglyLinkedList[int]()
	err := l.RemoveTail()
	MustBeErr(t, err, ErrTailNotFound)

	l.AddAll([]int{1, 2, 3, 4, 5})

	err = l.RemoveTail()
	MustBeNil(t, err)
	iteratorMustEqual[int](t, l.Iterator(), []int{1, 2, 3, 4})

	err = l.RemoveTail()
	MustBeNil(t, err)
	iteratorMustEqual[int](t, l.Iterator(), []int{1, 2, 3})

	err = l.RemoveTail()
	MustBeNil(t, err)
	iteratorMustEqual[int](t, l.Iterator(), []int{1, 2})

	err = l.RemoveTail()
	MustBeNil(t, err)
	iteratorMustEqual[int](t, l.Iterator(), []int{1})

	err = l.RemoveTail()
	MustBeNil(t, err)
	iteratorMustEqual[int](t, l.Iterator(), []int{})

	err = l.RemoveTail()
	MustBeErr(t, err, ErrTailNotFound)
}

func TestSinglyLinkedList_Set(t *testing.T) {
	l := NewSinglyLinkedList[int]()
	l.AddAll([]int{1, 2, 3, 4, 5})

	err := l.Set(-1, 6)
	MustBeErr(t, ErrInvalidIndex, err)

	err = l.Set(5, 6)
	MustBeErr(t, ErrInvalidIndex, err)

	err = l.Set(0, 6)
	MustBeNil(t, err)
	iteratorMustEqual[int](t, l.Iterator(), []int{6, 2, 3, 4, 5})

	err = l.Set(4, 7)
	MustBeNil(t, err)
	iteratorMustEqual[int](t, l.Iterator(), []int{6, 2, 3, 4, 7})

	err = l.Set(1, 8)
	MustBeNil(t, err)
	iteratorMustEqual[int](t, l.Iterator(), []int{6, 8, 3, 4, 7})

	err = l.Set(3, 9)
	MustBeNil(t, err)
	iteratorMustEqual[int](t, l.Iterator(), []int{6, 8, 3, 9, 7})
}
