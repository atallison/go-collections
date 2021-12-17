package collection

import (
	"testing"
)

func TestArrayListIterator(t *testing.T) {
	al := NewArrayList[int]()
	al.AddAll([]int{1, 2, 3, 4, 5})

	i := al.Iterator()
	cnt := 0
	got := []int{}
	for i.Next() {
		got = append(got, i.Value())
		cnt++
	}
	MustEqual(t, got, []int{1, 2, 3, 4, 5})
	MustEqual(t, cnt, 5)
}
