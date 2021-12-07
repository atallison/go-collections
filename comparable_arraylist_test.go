package collection

import "testing"

func TestRemove(t *testing.T) {
	al := NewComparableArrayList[int]()
	MustEqual(t, []int{}, al.values)
	al.AddAll([]int{1, 2, 3, 1, 2, 3})
	al.Remove(1)
	MustEqual(t, []int{2, 3, 2, 3}, al.values)
}
