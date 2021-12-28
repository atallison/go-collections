package arraylist

import (
	"testing"

	"github.com/hidetatz/collection/testutil"
)

func TestArrayList_Iterator(t *testing.T) {
	al := New[int]()

	pushBackAll(t, al, []int{1, 2, 3, 4, 5})

	i := al.Iterator()
	buff := []int{}
	cnt := 0
	for i.Next() {
		buff = append(buff, i.Value())
		cnt++
	}

	testutil.MustEqual(t, []int{1, 2, 3, 4, 5}, buff)
	testutil.MustEqual(t, 5, cnt)
}
