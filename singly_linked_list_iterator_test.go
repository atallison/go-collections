package collection

import (
	"testing"
)

func TestSinglyLinkedListIterator(t *testing.T) {
	l := NewSinglyLinkedList[int]()
	l.AddAll([]int{1, 2, 3, 4, 5})
	i := l.Iterator()

	cnt := 0
	got := []int{}
	for i.Next() {
		got = append(got, i.Value())
		cnt++
	}

	MustEqual(t, got, []int{1, 2, 3, 4, 5})
	MustEqual(t, cnt, 5)
}
