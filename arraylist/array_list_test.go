package arraylist

import (
	"testing"

	"github.com/hidetatz/collection"
	"github.com/hidetatz/collection/testutil"
)

func pushBackAll[T any](t *testing.T, l *ArrayList[T], values []T) {
	t.Helper()
	for _, v := range values {
		l.PushBack(v)
	}
}

func TestArrayList_PushBack(t *testing.T) {
	al := New[int]()
	testutil.MustEqual(t, []int{}, al.values)
	al.PushBack(1)
	testutil.MustEqual(t, []int{1}, al.values)
	al.PushBack(2)
	testutil.MustEqual(t, []int{1, 2}, al.values)
}

func TestArrayList_PushFront(t *testing.T) {
	al := New[int]()
	testutil.MustEqual(t, []int{}, al.values)
	al.PushFront(1)
	testutil.MustEqual(t, []int{1}, al.values)
	al.PushFront(2)
	testutil.MustEqual(t, []int{2, 1}, al.values)
}

func TestArrayList_Len(t *testing.T) {
	al := New[int]()
	testutil.MustEqual(t, 0, al.Len())

	al.PushBack(1)
	testutil.MustEqual(t, 1, al.Len())

	pushBackAll(t, al, []int{2, 3, 4})
	testutil.MustEqual(t, 4, al.Len())
}

func TestArrayList_Init(t *testing.T) {
	al := New[int]()

	pushBackAll(t, al, []int{1, 2})
	al.Init()
	testutil.MustEqual(t, []int{}, al.values)

	pushBackAll(t, al, []int{3, 4})
	testutil.MustEqual(t, []int{3, 4}, al.values)

	al.Init()
	testutil.MustEqual(t, []int{}, al.values)
}

func TestArrayList_GetAt(t *testing.T) {
	al := New[int]()

	pushBackAll(t, al, []int{1, 2, 3})

	_, err := al.GetAt(-1)
	testutil.MustBeErr(t, err, collection.ErrInvalidIndex)

	_, err = al.GetAt(3)
	testutil.MustBeErr(t, err, collection.ErrInvalidIndex)

	ret, err := al.GetAt(2)
	testutil.MustBeNil(t, err)
	testutil.MustEqual(t, ret, 3)
}

func TestArrayList_InsertAt(t *testing.T) {
	al := New[int]()

	err := al.InsertAt(1, 1)
	testutil.MustBeErr(t, err, collection.ErrInvalidIndex)

	testutil.MustBeNil(t, al.InsertAt(0, 1))
	testutil.MustEqual(t, []int{1}, al.values)

	err = al.InsertAt(3, 2)
	testutil.MustBeErr(t, err, collection.ErrInvalidIndex)

	err = al.InsertAt(2, 2)
	testutil.MustBeErr(t, err, collection.ErrInvalidIndex)

	testutil.MustBeNil(t, al.InsertAt(1, 2))
	testutil.MustEqual(t, []int{1, 2}, al.values)

	testutil.MustBeNil(t, al.InsertAt(1, 3))
	testutil.MustEqual(t, []int{1, 3, 2}, al.values)

	testutil.MustBeNil(t, al.InsertAt(0, 4))
	testutil.MustEqual(t, []int{4, 1, 3, 2}, al.values)
}

func TestArrayList_RemoveAt(t *testing.T) {
	al := New[int]()

	pushBackAll(t, al, []int{1, 2, 3})

	err := al.RemoveAt(-1)
	testutil.MustBeErr(t, err, collection.ErrInvalidIndex)

	err = al.RemoveAt(3)
	testutil.MustBeErr(t, err, collection.ErrInvalidIndex)

	err = al.RemoveAt(1)
	testutil.MustBeNil(t, err)
	testutil.MustEqual(t, []int{1, 3}, al.values)
}

func TestArrayList_SetAt(t *testing.T) {
	al := New[int]()

	err := al.SetAt(-1, 1)
	testutil.MustBeErr(t, err, collection.ErrInvalidIndex)

	err = al.SetAt(1, 1)
	testutil.MustBeErr(t, err, collection.ErrInvalidIndex)

	pushBackAll(t, al, []int{1, 2, 3})

	err = al.SetAt(3, 4)
	testutil.MustBeErr(t, err, collection.ErrInvalidIndex)

	err = al.SetAt(2, 4)
	testutil.MustBeNil(t, err)
	testutil.MustEqual(t, []int{1, 2, 4}, al.values)
}
