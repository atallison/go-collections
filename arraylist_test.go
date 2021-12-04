package collection

import (
	"testing"
)

func TestAdd(t *testing.T) {
	al := NewArrayList[int]()
	MustEqual(t, []int{}, al.values)
	al.Add(1)
	MustEqual(t, []int{1}, al.values)
	al.Add(2)
	MustEqual(t, []int{1, 2}, al.values)
}

func TestAddAll(t *testing.T) {
	al := NewArrayList[int]()
	al.AddAll([]int{1})
	MustEqual(t, []int{1}, al.values)

	al.AddAll([]int{2, 3, 4})
	MustEqual(t, []int{1, 2, 3, 4}, al.values)

	al.AddAll([]int{5})
	MustEqual(t, []int{1, 2, 3, 4, 5}, al.values)

	al.AddAll([]int{})
	MustEqual(t, []int{1, 2, 3, 4, 5}, al.values)
}

func TestAddAt(t *testing.T) {
	al := NewArrayList[int]()

	err := al.AddAt(1, 1)
	MustBeErr(t, err, ErrInvalidIndex)

	MustBeNil(t, al.AddAt(0, 1))
	MustEqual(t, []int{1}, al.values)

	err = al.AddAt(3, 2)
	MustBeErr(t, err, ErrInvalidIndex)

	err = al.AddAt(2, 2)
	MustBeErr(t, err, ErrInvalidIndex)

	MustBeNil(t, al.AddAt(1, 2))
	MustEqual(t, []int{1, 2}, al.values)

	MustBeNil(t, al.AddAt(1, 3))
	MustEqual(t, []int{1, 3, 2}, al.values)

	MustBeNil(t, al.AddAt(0, 4))
	MustEqual(t, []int{4, 1, 3, 2}, al.values)
}

func TestAddAllAt(t *testing.T) {
	al := NewArrayList[int]()

	err := al.AddAllAt(1, []int{1})
	MustBeErr(t, err, ErrInvalidIndex)

	MustBeNil(t, al.AddAllAt(0, []int{1}))
	MustEqual(t, []int{1}, al.values)

	err = al.AddAllAt(2, []int{2, 3, 4})
	MustBeErr(t, err, ErrInvalidIndex)

	MustBeNil(t, al.AddAllAt(1, []int{2, 3, 4}))
	MustEqual(t, []int{1, 2, 3, 4}, al.values)

	MustBeNil(t, al.AddAllAt(2, []int{5, 6}))
	MustEqual(t, []int{1, 2, 5, 6, 3, 4}, al.values)
}
