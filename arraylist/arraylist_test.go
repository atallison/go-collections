package arraylist

import (
	"reflect"
	"testing"
)

func TestAppend(t *testing.T) {
	al := New[int]()
	al.Append(1)
	al.Append(2)
	if !reflect.DeepEqual(al.values, []int{1, 2}) {
		t.Errorf("expected: %v but got %v", []int{1, 2}, al.values)
	}
}

func TestAppendAll(t *testing.T) {
	al := New[int]()
	al.Append(1)
	al.AppendAll([]int{2, 3, 4})
	al.AppendAll([]int{5})
	if !reflect.DeepEqual(al.values, []int{1, 2, 3, 4, 5}) {
		t.Errorf("expected: %v but got %v", []int{1, 2, 3, 4, 5}, al.values)
	}
}
