package collection

import "testing"

func TestComparableLinkedList_Remove(t *testing.T) {
	al := NewComparableLinkedList[int]()
	al.Remove(1)
	iteratorMustEqual[int](t, al.Iterator(), []int{})

	al.AddAll([]int{2})
	al.Remove(1)
	iteratorMustEqual[int](t, al.Iterator(), []int{2})

	al.Remove(2)
	iteratorMustEqual[int](t, al.Iterator(), []int{})

	al.AddAll([]int{2, 2, 2})
	al.Remove(2)
	iteratorMustEqual[int](t, al.Iterator(), []int{})

	al.AddAll([]int{2, 2, 3, 3})
	al.Remove(2)
	iteratorMustEqual[int](t, al.Iterator(), []int{3, 3})

	al.AddAll([]int{2, 3, 2, 3})
	al.Remove(2)
	iteratorMustEqual[int](t, al.Iterator(), []int{3, 3, 3, 3})
}

func TestComparableLinkedList_Contains(t *testing.T) {
	al := NewComparableLinkedList[int]()
	MustEqual(t, false, al.Contains(1))

	al.AddAll([]int{1})
	MustEqual(t, true, al.Contains(1))
	MustEqual(t, false, al.Contains(2))

	al.AddAll([]int{2, 3})
	MustEqual(t, true, al.Contains(1))
	MustEqual(t, true, al.Contains(2))
	MustEqual(t, true, al.Contains(3))
	MustEqual(t, false, al.Contains(4))
}
