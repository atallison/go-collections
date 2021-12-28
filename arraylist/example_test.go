package arraylist_test

import (
	"fmt"

	"github.com/hidetatz/collection/arraylist"
)

func ExampleArrayList() {
	l := arraylist.New[int]()
	// Pushes a value
	l.PushBack(1)
	l.PushFront(2)
	l.PushBack(3)

	// To iterate over the list, use iterator
	i := l.Iterator()
	for i.Next() {
		fmt.Println(i.Value())
	}
}
