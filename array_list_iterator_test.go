package collection

import "testing"

func TestArrayListIterator_Next(t *testing.T) {
	al := NewArrayList[int]()
	al.AddAll([]int{1, 2, 3, 4, 5})

	i := al.Iterator()
	cnt := 0
	for i.Next() {
		cnt++
		_ = i.Value() // consume the value to iterate
	}
	MustEqual(t, 5, cnt)
}

func TestArrayListIterator_Value(t *testing.T) {
	al := NewArrayList[int]()
	al.AddAll([]int{1, 2, 3, 4, 5})

	i := al.Iterator()
	got := []int{}
	for i.Next() {
		got = append(got, i.Value())
	}
	MustEqual(t, []int{1, 2, 3, 4, 5}, got)
}

func TestArrayListIterator_Set(t *testing.T) {
	al := NewArrayList[int]()
	al.AddAll([]int{1, 2, 3, 4, 5})

	i := al.Iterator()
	for i.Next() {
		v := i.Value()
		i.Set(v * 2)
	}
	MustEqual(t, []int{2, 4, 6, 8, 10}, al.values)

	i2 := al.Iterator()
	for i2.Next() {
		v := i2.Value()
		i2.Set(v * 2)
	}
	MustEqual(t, []int{4, 8, 12, 16, 20}, al.values)
}

func TestArrayListIterator_Remove(t *testing.T) {
	al := NewArrayList[int]()
	al.AddAll([]int{1, 2, 3, 4, 5})

	i := al.Iterator()
	for i.Next() {
		v := i.Value()
		if v%2 == 0 {
			i.Remove()
		}
	}
	MustEqual(t, []int{1, 3, 5}, al.values)

	i2 := al.Iterator()
	for i2.Next() {
		v := i2.Value()
		if v == 3 {
			i2.Remove()
		}
	}
	MustEqual(t, []int{1, 5}, al.values)
}
