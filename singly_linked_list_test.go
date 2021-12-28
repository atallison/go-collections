package collection

import (
	"testing"
)

func TestSinglyLinkedList_Add(t *testing.T) {
	l := NewSinglyLinkedList[int]()
	MustEqual(t, (*singlyLinkedNode[int])(nil), l.head)

	l.Add(1)
	MustEqual(t, 1, l.head.v)
	MustEqual(t, 1, l.length)
	collectionMustEqual[int](t, []int{1}, l)

	l.Add(2)
	MustEqual(t, 1, l.head.v)
	collectionMustEqual[int](t, []int{1, 2}, l)
	MustEqual(t, 2, l.length)
}

func TestSinglyLinkedList_Clear(t *testing.T) {
	l := NewSinglyLinkedList[int]()
	l.AddAll([]int{1, 2, 3})
	collectionMustEqual[int](t, []int{1, 2, 3}, l)
	MustEqual(t, 3, l.length)

	l.Clear()
	collectionMustEqual[int](t, []int{}, l)
	MustEqual(t, 0, l.length)

	// makes sure the cleared list is still usable
	l.AddAll([]int{1, 2, 3})
	collectionMustEqual[int](t, []int{1, 2, 3}, l)
	MustEqual(t, 3, l.length)
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

func TestSinglyLinkedList_String(t *testing.T) {
}

func TestSinglyLinkedList_AddHead(t *testing.T) {
	l := NewSinglyLinkedList[int]()
	MustEqual(t, (*singlyLinkedNode[int])(nil), l.head)

	l.Add(1)
	l.Add(2)
	l.AddHead(3)
	collectionMustEqual[int](t, []int{3, 1, 2}, l)
	MustEqual(t, 3, l.length)

	MustEqual(t, 3, l.head.v)
}

func TestSinglyLinkedList_AddAll(t *testing.T) {
	l := NewSinglyLinkedList[int]()
	l.AddAll([]int{})
	collectionMustEqual[int](t, []int{}, l)
	MustEqual(t, 0, l.length)

	l.AddAll([]int{0})
	collectionMustEqual[int](t, []int{0}, l)
	MustEqual(t, 1, l.length)

	MustEqual(t, 0, l.head.v)

	l.AddAll([]int{1, 2, 3})
	collectionMustEqual[int](t, []int{0, 1, 2, 3}, l)
	MustEqual(t, 4, l.length)

	MustEqual(t, 0, l.head.v)

	l.AddAll([]int{4, 5, 6})
	collectionMustEqual[int](t, []int{0, 1, 2, 3, 4, 5, 6}, l)
	MustEqual(t, 7, l.length)

	MustEqual(t, 0, l.head.v)

	l.AddAll([]int{7})
	collectionMustEqual[int](t, []int{0, 1, 2, 3, 4, 5, 6, 7}, l)
	MustEqual(t, 8, l.length)

	MustEqual(t, 0, l.head.v)
}

func TestSinglyLinkedList_Clone(t *testing.T) {
	l := NewSinglyLinkedList[int]()
	l.AddAll([]int{1, 2, 3})
	collectionMustEqual[int](t, []int{1, 2, 3}, l)
	MustEqual(t, 3, l.length)

	nl := l.Clone()
	collectionMustEqual[int](t, []int{1, 2, 3}, l)
	MustEqual(t, 3, nl.length)

	// add values into l
	l.AddAll([]int{4, 5, 6})
	collectionMustEqual[int](t, []int{1, 2, 3, 4, 5, 6}, l)
	MustEqual(t, 6, l.length)

	// makes sure the cloned list is not changed
	collectionMustEqual[int](t, []int{1, 2, 3}, nl)
	MustEqual(t, 3, nl.length)

	// add values into nl
	nl.AddAll([]int{7, 8, 9})
	collectionMustEqual[int](t, []int{1, 2, 3, 7, 8, 9}, nl)
	MustEqual(t, 6, nl.length)

	// make srue the original list is not changed
	collectionMustEqual[int](t, []int{1, 2, 3, 4, 5, 6}, l)
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
	collectionMustEqual[int](t, []int{1, 2, 3}, l)
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

func TestSinglyLinkedList_Set(t *testing.T) {
	l := NewSinglyLinkedList[int]()
	l.AddAll([]int{1, 2, 3, 4, 5})

	err := l.Set(-1, 6)
	MustBeErr(t, ErrInvalidIndex, err)

	err = l.Set(5, 6)
	MustBeErr(t, ErrInvalidIndex, err)

	err = l.Set(0, 6)
	MustBeNil(t, err)
	collectionMustEqual[int](t, []int{6, 2, 3, 4, 5}, l)

	err = l.Set(4, 7)
	MustBeNil(t, err)
	collectionMustEqual[int](t, []int{6, 2, 3, 4, 7}, l)

	err = l.Set(1, 8)
	MustBeNil(t, err)
	collectionMustEqual[int](t, []int{6, 8, 3, 4, 7}, l)

	err = l.Set(3, 9)
	MustBeNil(t, err)
	collectionMustEqual[int](t, []int{6, 8, 3, 9, 7}, l)
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
	collectionMustEqual[int](t, []int{1, 2, 4, 5}, l)

	err = l.RemoveAt(3)
	MustBeNil(t, err)
	collectionMustEqual[int](t, []int{1, 2, 4}, l)

	err = l.RemoveAt(0)
	MustBeNil(t, err)
	collectionMustEqual[int](t, []int{2, 4}, l)
}
