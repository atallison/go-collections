package collection

import "testing"

func TestSinglyLinkedListIterator_Next(t *testing.T) {
	l := NewSinglyLinkedList[int]()
	l.AddAll([]int{1, 2, 3, 4, 5})
	i := l.Iterator()

	cnt := 0
	for i.Next() {
		_ = i.Value()
		cnt++
	}

	MustEqual(t, 5, cnt)
}

func TestSinglyLinkedListIterator_Value(t *testing.T) {
	l := NewSinglyLinkedList[int]()
	l.AddAll([]int{1, 2, 3, 4, 5})
	i := l.Iterator()

	buff := []int{}
	for i.Next() {
		buff = append(buff, i.Value())
	}

	MustEqual(t, []int{1, 2, 3, 4, 5}, buff)
}

func TestSinglyLinkedListIterator_Set(t *testing.T) {
	l := NewSinglyLinkedList[int]()
	l.AddAll([]int{1, 2, 3, 4, 5})

	i := l.Iterator()
	for i.Next() {
		v := i.Value()
		i.Set(v * 2)
	}
	collectionMustEqual[int](t, []int{2, 4, 6, 8, 10}, l)

	i2 := l.Iterator()
	for i2.Next() {
		v := i2.Value()
		i2.Set(v * 2)
	}
	collectionMustEqual[int](t, []int{4, 8, 12, 16, 20}, l)
}

func TestSinglyLinkedListIterator_Remove(t *testing.T) {
	l := NewSinglyLinkedList[int]()
	l.AddAll([]int{1, 2, 3, 4, 5})

	i := l.Iterator()
	for i.Next() {
		v := i.Value()
		if v%2 == 0 {
			i.Remove()
		}
	}
	collectionMustEqual[int](t, []int{1, 3, 5}, l)

	i2 := l.Iterator()
	for i2.Next() {
		v := i2.Value()
		if v == 3 {
			i2.Remove()
		}
	}
	collectionMustEqual[int](t, []int{1, 5}, l)
}
