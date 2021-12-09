package collection

import "testing"

func TestRemove(t *testing.T) {
	al := NewComparableArrayList[int]()
	al.AddAll([]int{1, 2, 3, 1, 2, 3})
	al.Remove(1)
	MustEqual(t, []int{2, 3, 2, 3}, al.values)
}

func TestContains(t *testing.T) {
	al := NewComparableArrayList[int]()
	al.AddAll([]int{1, 2, 3, 4, 5, 6})
	MustEqual(t, true, al.Contains(3))
	MustEqual(t, false, al.Contains(0))
}

func TestIndexOf(t *testing.T) {
	al := NewComparableArrayList[int]()
	al.AddAll([]int{1, 2, 3, 1, 2, 3})
	MustEqual(t, 2, al.IndexOf(3))
	MustEqual(t, -1, al.IndexOf(0))
}

func TestLastIndexOf(t *testing.T) {
	al := NewComparableArrayList[int]()
	al.AddAll([]int{1, 2, 3, 1, 2, 3})
	MustEqual(t, 5, al.LastIndexOf(3))
	MustEqual(t, -1, al.LastIndexOf(0))
}
